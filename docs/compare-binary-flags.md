# Comparing Binary Flags

Flags for the Cert-Manager Controller, Webhook, and CAInjector components may
change with every release, and we will need to know when they do so that we can
make sure users can use those flags when requesting that release.

As an example, this for loop spits out CLI flags seen when passing the
`--help` to the controller, webhook, or cainjector directly, and then tools like
`diff` can help determine any differences or changes that need to be accounted
for in the operator's `configs` package.

```bash
for z in 0 1.0.1 1.0.2 1.0.3; do 
    for component in controller webhook cainjector; do 
        docker run "quay.io/jetstack/cert-manager-${component}:v${z}" --help | grep -- "--[a-z]" | awk ' { print $1 }' | tee -a "args-${component}-v${z}".txt
    done
done
```