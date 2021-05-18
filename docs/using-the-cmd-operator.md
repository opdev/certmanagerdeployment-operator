# Using the Cert-Manager Deployment Operator

As mentioned, the Cert-Manager Deployment Operator does not re-implement the
Cert-Manager project, but rather aims to help operator Cert-Manager clusters
over time.

In addition, this project aims to enable the inclusion of additional
functionality that can benefit workloads consuming TLS Certificates.

## Operator Configuration

Operator configuration is currently handled using flags passed to the binary
execution at runtime. Most available flags are standard flags scaffolded by
Golang operators created with Operator SDK. #TODO link

The Cert-Manager Deployment Operator includes an optional PodRefresh Controller
that will restart a controller-backed workload (e.g. Deployments, StatefulSets)
on the change of a TLS secret.

To enable this additional, optional controller, you will need to pass the
`--enable-pod-refresher` flag to the controller manager, which will then enable
the controller on start up.

This optional configuration may be reconfigured as an environment variable or
ConfigMap entry in the future.

## Unsupported Configurations

When instantiating an instance of Cert-Manager by creating a
CertManagerDeployment resource in your cluster, you can influence the resulting
deployment manifests in several ways if the default values do not suite your
needs.

This is done by adding values to the `CertManagerDeployment.spec.dangerZone`.

Each DangerZone spec item contains a mapping of keys to data, where the keys
must be one of `controller`, `webhook`, or `cainjector` - such that the change
corresponds with the given component.

### Overriding Component Images

The `CertManagerDeployment.spec.dangerZone.imageOverrides` configuration allows
you to override a deployment image for your particular deployment manifest. This
is potentially a dangerous change, and the CertManagerDeployment Operator is
unable to provide a guarantee that it will function appropriately with a custom
image. With that said, it's possible that an adminsitrator may want to
temporarily override an image for testing purposes.

To replace an image for a given component, simply assign that component's key in
`CertManagerDeployment.spec.dangerZone.imageOverrides.<component>` to the
replacement image.

Example:

```yaml
spec:
  version: "v1.2.0"
  dangerZone:
    imageOverrides:
        cainjector: "my-registry.example.com/my-namespace/cert-manager-cainjector:v0.0.1"
```

This image string must be a complete reference to an
accessible container image, and as such must include things such as registry,
namespace, and tag.

### Overriding Component Runtime Parameters

The `CertManagerDeployment.spec.dangerZone.containerArgOverrides` configuration
allows you to override the component binary runtime arguments (e.g. the
`--flags` you pass to each component).

To replace the arguments for a given component's container, simply assign that
component's key in
`CertManagerDeployment.spec.dangerZone.containerArgOverrides.<component>` to a
map of keys and values that are then filtered, and if appropriate for the
specified version, passed to the deployment manifest for the component.

To filter these arbitrary keys and values, the CertManagerDeployment Operator
stores a mapping of possible parameters for a given version of the component.

So as an example, the v1.1.0 version of the Cert-Manager Controller supports the
`--renew-before-expiry-duration`. To set this key with an appropriate value, you
would add it to the containerArgOverrides as below:

```yaml
spec:
  version: "v1.1.0"
  dangerZone:
    containerArgOverrides:
        controller:
            renew-before-expiry-duration: <yourvalue>
```

For boolean flags, (e.g. things that would simply be enabled by passing the
parameter itself without a value such as `--enable-feature`), explicitly passing
the appropriate boolean value is required.

If you attempt to add an unsupported key, the CertManagerDeployment Operator
will filter the key and value out, preventing the controller binary from being
passed an incorrect flag.

As a result of this filter process, adjusting the version of Cert-Manager to
another version where a component does not support the flag that is being
overridden can cause unexpected behavior as the operator attempts to filter out
keys that do not exist in the currently newly requested version of the
controller. For this reason, it is recommended that all unsupported
configuration overrides are evaluated when changing the requested version.

## Using the Pod Refresher

The pod refresher controller, once enabled, will restart a workload that meets
the following criteria:

* The workload has opted in to being refreshed through the use of appropriate
  annotation.
* The workload has a TLS certificate that has been issued by Cert Manager
  mounted into the workload.
  
The following annotation is required on the workload that is opting in to be
refreshed.

```yaml
    “certmanagerdeployment.opdev.io/allow-restart”: “true”
```

When the operator restarts a workload, it subsequently labels the workload with
a timestamp indicating when it was last restarted.

```yaml
    “certmanagerdeployment.opdev.io/time-restarted”: “2020-7-9.1640”
```
