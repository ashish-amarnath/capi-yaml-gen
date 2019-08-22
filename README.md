# capi-yaml-gen
Tool to generate yaml for Cluster-API and necessary provider resources.

Read more about this [here](https://docs.google.com/document/d/1Tzx6IXOoQnUxaVSYA2I8IdcNrFE4zgOnzkk55KHOU20/edit)

## Building the tool

Building this tool requires `go 1.12+` for go module support.

To build the binary `capi-yaml-gen`,  run the below command from root of the repo

```(bash)
$ go build -o ./capi-yaml-gen .
```

## Usage Examples

### Generating yaml for CAPD cluster
The below command generate yaml for CAPI using CAPD as the infrastructure provider
```(bash)
$ ./capi-yaml-gen alpha get-cluster-yaml -p docker --name test-cluster --namespace my-awesome-namespace

---
apiVersion: infrastructure.cluster.x-k8s.io/v1alpha2
kind: DockerCluster
metadata:
  creationTimestamp: null
  name: test-cluster
  namespace: my-awesome-namespace
spec: {}
status:
  ready: false

---
apiVersion: cluster.x-k8s.io/v1alpha2
kind: Cluster
metadata:
  creationTimestamp: null
  name: test-cluster
  namespace: my-awesome-namespace
spec:
  infrastructureRef:
    apiVersion: infrastructure.cluster.x-k8s.io/v1alpha2
    kind: DockerCluster
    name: test-cluster
    namespace: my-awesome-namespace
status:
  infrastructureReady: false

---

```
