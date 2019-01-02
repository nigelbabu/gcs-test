# GCS-Tests

This is a temporary repo to host the first few GCS tests until we decide where
these tests should live. After that decision, these tests can move over to that
repo and live there.

## How to run the tests
1. Setup a Kubernetes Cluster and setup GCS inside it as the [GCS deployment
   scripts recommend][1].
2. Make sure `kubectl` points to that cluster by doing the following:

    KUBECONFIG=/home/nigelb/code/gcs/deploy/kubeconfig

3. Clone this repository in the right folder in GOPATH and cd into the e2e
   folder.
4. Install dependencies with `dep ensure`.
4. Run the tests with `go test`.


[1]: https://github.com/gluster/gcs/tree/master/deploy
