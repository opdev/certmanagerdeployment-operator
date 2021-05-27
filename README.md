# Cert Manager Deployment Operator

The Cert-Manager Deployment operator aims to help OpenShift platform
administrators run [Cert-Manager](https://github.com/jetstack/cert-manager)
alongside additional controllers to provide complementary feature-sets for
managing workloads consuming TLS certificates.

This operator is written in Golang using the
[OperatorSDK](https://github.com/operator-framework/operator-sdk).
This project is considered **alpha** maturity and is in early stages of
development. Use at your own risk.

## Usage

The easiest way to run this operator is to install the operator-specific catalog
in your OpenShift cluster.

```bash
oc apply -f - <<EOF
apiVersion: operators.coreos.com/v1alpha1
kind: CatalogSource
metadata:
  name: certmanagerdeployment-operator
spec:
  displayName: CertManagerDeployment Operator Index
  image: quay.io/opdev/certmanagerdeployment-operator-index:latest
  priority: -200
  publisher: The OpDev Team
  sourceType: grpc
  updateStrategy:
    registryPoll:
      interval: 30m0s
EOF
```

With this catalog installed, you can then navigate to the embedded Operator
Hub and subscribe to the operator.

To test using a local copy of the source code, use the following Make target:

```bash
make install
OPFLAGS=--enable-pod-refresher make run
```

## Important

This project does not change the functionality of Cert-Manager itself, and has
**no direct association** with Cert-Manager project
(i.e. please don't open issues for jetstack/cert-manager unless it's *confirmed*
to be a cert-manager specific issue, we do not want to create extra noise for
those folks).

For those that desire to forgo additional controllers and strictly use the
Cert-Manager project itself, delivered as an operator without additional
controllers, please see
[jetstack/cert-manager-olm](https://github.com/jetstack/cert-manager-olm).
