# Supporting New Versions of Cert-Manager

When Cert-Manager release a new version, this operator codebase will require
several changes to ensure it's able to deploy that release successfully.

## Meta & Structure

**Diff the previously supported version's release assets against the latest
release assets to look for changes.**
[See example workflow](diff-manifests.md)

* The operator rolls out these assets programmatically, so any changes across
eassets from different versions need to be accounted for programmatically.

* A common example is a given component's container image which changes with
each release of Cert-Manager.

* Another example might be the Deployment spec.

* These changes influence changes in the codebase (described below)

**Create a new top-level directory containing cert-manager CRDs and add the
`managed-by: operator` label.**

* These are the CRDs of interest, such as Certificates, Issuers, etc.
  
* The label addition is purely semantic.

* These resource should also be compared. If there is no change, then creating a
symlink to an already existing directory of CRDs cuts down on extra files.

**Update `Dockerfile` such that any new CRD directories on disk are copied
over to the resulting container image.**

* When building the operator container image, the image needs to contain the
CRDs, so update the Dockerfile to include them (e.g. `COPY v1.2.0/ v1.2.0/`)

## Codebase

**Update `SupportedVersions` Map in `controllers/componentry/componentry.go`**

* This designates in code what versions of Cert-Manager the operator supports,
and directly restricts a user's ability to request an unsupported version.

**Update `CertManagerDefaultVersion` Variable in
`controllers/componentry/constants.go`**

* This is always the latest version, and is what's deployed if the user does not
specify a version.

* Update `GetComponentFor*` Func for each component.
  
* Each component type contains metadata informing how the various secondary
resources should be created. When a default version change, the default metadata
may also change.

* Most of these changes come from diffing the previously supported version's
resources (e.g roles, deploymnents, against the one you want to deploy.

* If anything changed between the new default version and the previous, those
changes required to deploy the previous version will still need to be accounted
for.

**Update `getCRDListForCertManagerVersion` Func in
`controllers/customresourcedefinitions.go`**

* This needs to include an entry for the newly supported version.

**Update CertManagerDeployment type markers such that they reflect new supported
versions in `api/v1alpha1/certmanagerdeployment_types.go`**

* This ensures that the generated CustomResourceDefinition has proper validation
preventing creation of a CertManagerDeployment with an incorrect version.

**Update `config/samples/operators_v1alpha1_certmanagerdeployment.yaml` such
that it deploys the latest version**

* This is informative, and also guides users using the OpenShift Console have an
example available when creating resources.

**Update objects in `controllers/configs/` to ensure that it serves up the right
empty and default configuration objects for the version in `getter.go`.**

* These configs represents CLI flags for each of the Cert-Manager components.
When a new version of Cert-Manager is made available, maintainers must check
the flags of the new version to ensure they're accurately reflected by these
types, as these types help enforce validation when users attempt to change the
runtime configuration of the Cert-Manager controller binaries.
[See example workflow](compare-binary-flags.md)
