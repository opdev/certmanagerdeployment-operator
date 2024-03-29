module github.com/opdev/certmanagerdeployment-operator

go 1.13

require (
	github.com/go-logr/logr v0.3.0
	github.com/imdario/mergo v0.3.10
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/openshift/library-go v0.0.0-20200930190915-f7cb85f605db
	k8s.io/api v0.20.2
	k8s.io/apiextensions-apiserver v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	sigs.k8s.io/controller-runtime v0.8.3
)
