# Running Tests

Testing is currently done against an existing standing cluster, either via
[Codeready Containers](https://github.com/code-ready/crc) or using an already
standing OpenShift cluster. Using a local control plane provisioned by EnvTest
has not been tested.

With your cluster ready and kubeconfig in your environment, indicate to EnvTest
that you're using an existing cluster.

```bash
export USE_EXISTING_CLUSTER=true
```

In a separate window, start your controller with the podrefresher controller
enabled.

```bash
make OPFLAGS=--enable-podrefresher run
```

Clean up from any previous executions if necessary (or skip otherwise).

```bash
./scripts/bin/test-cleanup.sh
```

Run the tests.

```bash
./scripts/bin/test-run.sh
```
