# Comparing Manifests

When a new release of Cert-Manager is available, the release includes a YAML
manifest that contains all resources necessary to deploy Cert-Manager in a
cluster.

For example,
[here](https://github.com/jetstack/cert-manager/releases/download/v1.1.1/cert-manager.yaml)
is the install manifest available from the release page of Cert-Manager 1.2.0.

Assuming this operator supports 1.2.0 as the latest release of Cert-Manager, the
operator would also need to deploy the latest 1.1.z release (1.1.0 at the time
of this writing). For this reason, maintainers will need to compare v1.1.0's
release manifests against 1.2.0 to ensure that any difference are accounted for
when generating the resource types programmatically.

This process could use automation, but here is the manual explanation of how
that is done today.

The workflow is as follows:

* Download both sets of release for versions that are being compared (e.g. 1.1.0
and 1.2.0) into separate directories.

* The manifests contain multiple YAML documents. For each version's manifest,
split the documents into individual files using a tool like
[YAMLSplit](https://github.com/komish/yamlsplit).

* If using [YAMLSplit](https://github.com/komish/yamlsplit), the individual
documents will have been split into enumerated files. The goal
at this point is to compare the corresponding files for the two versions to see
if any changes have been released to the default manifests. An easy way to do
this is with `git diff`.

* In order for the files to be properly diffed, they'll need to be sorted and
be matching (otherwise, we're comparing manifests for different resources which,
of course, will not match). An easy to to do this is to use
`./scripts/compare-yaml-names.sh` to perform the separate directories.

  * E.g. `./scripts/compare-yaml-names.sh cert-manager-v1.2.0/ cert-manager-v1.1.0`

* Sort the files that are out of order so that matching resources are in
matching enmumerated files.

* Initialize the older version as a git repository and commit all files. We do
not actually care about the git history so the commit message is irrelevant.

* Copy over the newer version's files over the older version's files.

* Use `git diff` to evaluate changes that need to be persisted to code.

Note: Changes to CRDs (other than GVK-related changes aren't overly important as
the CRDs are copied directly into place and persisted to the API from disk. That
said, it makes sense to pay special attention to test executions with regards to
CRDs if CRD changes have taken place to ensure that version-to-version upgrades
do not pose any issues.
