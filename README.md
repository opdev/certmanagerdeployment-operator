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

TODO

## Upstream

This project does not change the functionality of Cert-Manager itself, and has
**no direct association** with Cert-Manager project
(i.e. please don't open issues for jetstack/cert-manager unless it's *confirmed*
to be a cert-manager specific issue, we do not want to create extra noise for
those folks).

For those that desire to forgo additional controllers and strictly use the
Cert-Manager project itself, delivered as an operator without additional
controllers, please see
[jetstack/cert-manager-olm](https://github.com/jetstack/cert-manager-olm).
