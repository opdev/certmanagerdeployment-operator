# Implementation

This document serves as an introductory text detailing  the "how" and the "why"
behind this operator's functionality. It is intended to help users and
developers familiarize themselves with the behavior of the operator.

## Version Compatibility

**Cert-Manager Version Compatibility**

The Cert Manager Deployment operator does not currently have any releases. In an
effort to allow for upgrade scenarios across versions of the Cert-Manager, the
proposed release strategy intends to support "Y-1" release of Cert-Manager itself.

As an example, if Cert-Manager v1.3.0 is the "latest" release supported by the
operator, then v1.2.0 would also be installable for that version of the operator.
If Cert-Manager v1.4.0 is the latest release, then v1.3.0 would also be
installable for that version of the operator.

**Kubernetes Version Compatibility**

As of Cert-Manager
[v1.2.0](https://github.com/jetstack/cert-manager/releases/tag/v1.2.0),
the minimum supported version of Kubernetes is v1.16.0. It should be assumed that
this operator only supports the same minimum version of Kubernetes.

## Additional Controllers

### PodRefresher

The PodRefresher controller allow for users to opt-in to having their workloads
restarted by their respective workload controllers based on a change event to an
issued certificate secret. This allows workloads that are tolerant to arbitrary
restarts to ensure they're constantly referring to updated certificates once the
state of the secret has changed (e.g. due to a renewal).

**NOTE:** Additional controllers only intend to cover complementary functionality
that may not make sense to ship in the Cert-Manager project directly (e.g.
because it does not fit the core goals of the project, or any other reason why
it may not make sense for inclusion). If a feature makes sense to include in
Cert-Manager itself, it is always recommended to approach the Cert-Manager project
with a feature request via their existing procedures.

## Implementations

### CertManagerDeployment Controller

This controller introduces a new cluster-scoped custom resource, called
CertManagerDeployment, which represents an installation of Cert-Manager in a
cluster. Only a single instance of this resource can exist in a cluster, and
it must exist with the name "cluster". This aims at protecting the cluster from
the installation of multiple Cert-Manager instances by this operator.

Managing an installation of Cert-Manager requires the installation of various
resource in your cluster. The list of resources managed by the operator mirror
those made available via installation of Cert-Manager through Helm, or those made
available in release assets for a given
[release](https://github.com/jetstack/cert-manager/releases).

* Service Accounts
* Roles
* RoleBindings
* ClusterRoles
* ClusterRoleBindings
* Deployments
* Services
* MutatingWebhookConfigurations
* ValidatingWebhookConfigurations
* CustomResourceDefinitions

As such, with few exceptions (mentioned later), these resources are watched by
the controller and reconciled on change.

#### High-Level Reconciliation Flow

The reconciliation of CertManagerDeployment resources follows the following flow:

* Preflights Checks
  * Ensure the reconcile requests is for a CertManagerDeployment resource named
    "cluster".
  * Ensure the associated resource exists in the API and has not been deleted.
  * Ensure the resource spec defines a desired installation of Cert-Manager in a
    version supported by the operator.
* Secondary Resource Management
  * Update the status of the CertManagerDeployment resource.
  * Ensure CustomResourceDefinitions exist and are in the proper state.
  * Ensure Namespace exists and is in the proper state.
  * Ensure Service Accounts exist and are in the proper state.
  * Ensure Roles exist and are in the proper state.
  * Ensure Role Bindings exist and are in the proper state.
  * Ensure Cluster Roles exist and are in the proper state.
  * Ensure Cluster Role Bindings exist and are in the proper state.
  * Ensure Deployments exist and are in the proper state.
  * Ensure Services exist and are in the proper state.

Once this workflow has completed, the reconciler will end the reconciliation
with a successful result. This workflow can be seen in code in
[certmanagerdeployment_controller.go](../controllers/certmanagerdeployment/certmanagerdeployment_controller.go)
and directly correlates with the
`*CertManagerDeploymentReconciler).Reconcile(req ctrl.Request)` method.

#### Reconciling Secondary Resources

**Create**

* On reconciliation of a CertManagerDeployment, the reconciler will ensure that
  all secondary resources exist in appropriate states.

* The reconciler will handle creation of the install namespace, "cert-manager",
  and will also create all subresources ("service accounts", "deployments", etc.)
  in that namespace.

* All subresources are created with fixed resource names to ensure better
  consistency of the installation of Cert-Manager. These fixed resource names
  should align with those found in
  [release](https://github.com/jetstack/cert-manager/releases).

* The reconciler will also handle creation of Cert-Manager APIs ("certificates",
  "issuers", etc.), via their CRDs, as are found in
  [release](https://github.com/jetstack/cert-manager/releases) assets.

**Update**

Reconciliations after the initial creation of a subresource will include a
comparison to check to see if further reconciliation of the subresource is
required. This comparison may vary depending on the subresource.

* **Status** is always updated to ensure it is up to date.

* **Custom Resource Definitions** are updated if the spec, annotations, or
  labels do not match. This is asserted by walking the JSON representation* of
  each section in the existing instance and comparing it to the desired instance.
  * If labels do not match, the existing labels are replaced with the desired
    labels.
  * If annotations do not match, the existing annotations are replaced with
    desired labels.
  * If spec does not match, the existing spec is replaced with the desired spec.

* **Namespace** is not updated in any case.  There are no pre-populated fields
  passed to created namespaces other than their identifiable elements (i.e.
  group, version, kind) so there is no reconciliation that can take place to
  correct drift.

* **Service Accounts** are not updated in any case.  There are no pre-populated
  fields passed to created service accounts other than their identifiable
  elements (i.e. group, version, kind) so there is no reconciliation that can
  take place to correct drift.

* **Roles** and **Cluster Roles** are updated if the rules or labels do not
  match. This is asserted by *walking the JSON representation of each section
  in the existing instance and comparing it to the desired instance.
  * if rules do not match, the existing rules are replaced with the desired
    rules.
  * If labels do not match, the existing labels are replaced with the desired
    labels.

* **Role Bindings** and **Cluster Role Bindings** are updated if the subjects
  or labels do not match. This is asserted by *walking the JSON representation
  of each section in the existing instance and comparing it to the desired
  instance. The RoleRef cannot be updated, and drift of the RoleRef value is
  currently not handled as it would require deleting and replacing the resource
  * If subjects do not match, the existing subjects are replaced with the
    desired subjects.
  * If labels do not match, the existing labels are replaced with the desired
    labels.

* **Deployments** are updated if the specs, annotations, or labels do not match.
  This is asserted by *walking the JSON representation of each section in the
  existing instance and comparing it to the desired instance.
  * If specs do not match, the desired spec is merged into the existing spec to
    preventing overwriting any spec values defaulted by the APIServer with empty
    values.
  * If annotations do not match, the existing annotations are replaced with
    desired labels.
  * If spec does not match, the existing spec is replaced with the desired spec.

* **Services** are updated if the specs, annotations, or labels do not match.
  This is asserted by *walking the JSON representation of each section in the
  existing instance and comparing it to the desired instance.
  * If specs do not match, the desired spec is merged into the existing spec to
    preventing overwriting any spec values defaulted by the APIServer with empty
    values.
  * If annotations do not match, the existing annotations are replaced with
    desired labels.
  * If spec does not match, the existing spec is replaced with the desired spec.

_*Walking the JSON representation of the associated types refers to serializing
them to `map[string]interface{}` and recursively asserting that values match.
In all cases, this approach checks to ensure that the value in the desired
representation of the resource exists and matches the value of the existing.
That is to say that, for resources such as maps, other key:value pairs may exist
in the map - and that's acceptable as long as the desired keys:values exist.
See [cmdoputils/cmdoptuils.go.ObjectsMatch(...)](../cmdoputils/cmdopoutils.go)_

**Delete**

All resources *except* for custom resource definitions are created with owner
references set on the CertManagerDeployment resource. Deletion of the
CertManagerDeployment resource triggers garbage collection on all remaining
resources.

**Custom resource definitions** managed by the CertManagerDeployment operator
(i.e. "certificates", "issuers", etc.) must be cleaned up manually. These are
not managed with owner reference to prevent garbage collection from taking
place, as the operator must allow for uninstallation of the Cert-Manager
instance and/or the operator without impacting (a) existing workloads, or (b)
existing issued certificates/resources. This is intentional, and should allow
for the adoption of existing resources via another compatible installation of
Cert-Manager (installed via any method) if required.

#### Source Layout

The reconciler source exists in the `controllers/certmanagerdeployment`
directory, with the high-level reconciliation flows existing in
[certmanagerdeployment_controller.go](../controllers/certmanagerdeployment/certmanagerdeployment_controller.go).
Workflows for secondary resources are are stored in additional files,
organizedby resource name, in the same package (e.g. if a developer needs to
make a change to the way a role is reconciled, they would look in the
[roles.go](../controllers/certmanagerdeployment/roles.go) file).

There are three core functions stored in each respective secondary resource
source file.

* **Reconciler method** `*CertManagerDeploymentReconciler.reconcile<Resources>(...)`

This describes how the existing state of this resource is brought closer to the
desired state. These are methods bound to the `CertManagerDeployment` type which
gives them access to clients, loggers, and eventers to perform their task.

* **Accessor function** `Get<Resources>(...)`

This describes what exact resources of this kind should the reconciler need to
create/reconcile. These functions accept the current state of the given
`CertManagerDeployment` and return the requested `<Resources>` for all
Cert-Manager components.

* **Generator functions** `new<Resource>(...)`

This describes how a resource of this kind look like based on the specification.
These functions are informed by the custom resource and the particular
Cert-Manager *component* (e.g. webhook, cainjector, etc.) is in scope for the
request, as the secondary resource may take on different characteristics based
on the component.

#### Packages

**componentry**

[controllers/certmanagerdeployment/componentry](../controllers/certmanagerdeployment/componentry/componentry.go)

Referenced previously, a concept of a "Component" (type: `CertManagerComponent`)
is used to represent a given component of the Cert-Manager application. At the
time of writing, this refers to the Controller, the Webhook, and CAInjector, and
generally parallels the deployments that are included with a given Cert-Manager
installation.

Each component has an associated set of metadata that informs the creation of
secondary resources (*roles,*bindings, etc.) necessary for that component to
function.

The secondary resource reconcilers are responsible for generating all instances
of the secondary resource for all components. As an example, there are three
service accounts generated during calls to `reconcileServiceAccounts` -
one service account be component.

**configs**

[controllers/certmanagerdeployment/configs](../controllers/certmanagerdeployment/configs/doc.go)

Package "configs" codifies the parameters necessary to run the Cert-Manager
component binaries. Because the operator aims to support the installation of
multiple "Y" (context: semantic versioning, X.Y.Z) versions of the Cert-Manager
application, as well as multitudes of "Z" releases, the operator must allow for
an end user to specify a different configuration of flags to pass to the
binaries while still ensuring that (a) the binary will not fail to execute due
to "unknown parameters", and (b) the user is able to take advantage of new flags
as they become available.

The "configs" package makes versioned representations of the flags a user can
pass to Cert-Manager to use as molds, and passes user configuration (supplied
via `CertManagerDeployment.Spec.DangerZone`) through this mold to determine
validity of arguments. Finally, the supplied flags are passed to the associated
component's Deployment arguments to be executed.

**cmdoputils**

[cmdoputils](../cmdoputils/cmdoputils.go)

Package "cmdoputils" contains a series of utility functions used throughout the
operator codebase. This package depends on very little from the codebase itself
to prevent cyclical dependencies. Along with other functions, this package
contains the function `ObjectsMatch` which is used through out the code base to
determine if the desired state and existing state of a given resource match.

### PodRefresher Controller

This minimal controller doers not introduce any custom resources, and instead
watches Secrets. Cert-Manager stores issued certificates in secrets within the
same namespace. When secrets containing Cert-Manager -issued certificates are
modified, this controller will determine if the associated workload needs to be
restarted to ensure mounted secrets are up to date with the latest certificate.

This controller only responds to controller-backed workloads (i.e. Deployments,
Daemonsets, Statefulsets) that have opted int through an annotation present on
the workload itself.

The controller is enabled via a CLI flag to the `manager` binary built from this
project.

#### High-Level Reconciliation Flow

Watches are established on Secrets.

Predicates are used to filter out:

* Secrets that are not Cert-Manager -issued TLS Certificates.
* Secrets for which the Resource Version has not changed.

The reconciliation workflow looks like the following:

* Preflight
  * Ensure the associated resource exists in the API and has not been deleted.
  * Ensure the secret resource contains annotations implying it is a
    Cert-Manager-issued Certificate
    * NOTE: this is not strictly required due to predicate filtering, but is
      left in as a preflight check as a precaution
* Reconciliation
  * Queries all Deployments, Daemonsets, and Statefulsets in the same namespace
    as the secret
  * For each Deplopyment/Daemonset/Statefulset:
    * Confirms the resource has opted in to be refreshed
    * Confirms the resource is using an outdated secret
    * Updates the workload:
      * By adding a label to the workload indicating the time of the restart
      * By adding a label to the workload template indicating the time of the
        restart
      * By adding an annotation indicating the resource version of the secret
        driving the restart

It's important to prevent the controller from hot-looping, causing a workload to
continuously be inaccessible. For this reason, the controller does not currently
re-queue reconciliation for errors while trying to trigger a workload restart.

_**Workloads that have opted in will have the resource version of the secret
annotated on the workload. The controller observes this annotation to determine
if the last resource version that triggered a restart of the workload is the
same as the resource version associated with the event. This also helps prevent
the controller from triggering continuous restarts unnecesarily._

#### Source Layout

The controller operations are minimal, and are contained in a single package
"[podrefresher](../controllers/podrefresher/podrefresher.go)".
