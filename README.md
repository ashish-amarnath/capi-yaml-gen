# capi-yaml
Developer tool to generate yaml for Cluster-API and necessary provider resources.

Read more about this [here](https://docs.google.com/document/d/1Tzx6IXOoQnUxaVSYA2I8IdcNrFE4zgOnzkk55KHOU20/edit)

## Building the tool

Building this tool requires `go 1.12+` for go module support.

To build the binary `capi-yaml`,  run the below command from root of the repo

```(bash)
$ go build -o ./capi-yaml .
```

## Usage Examples

### Docker Infrastructure and Kubeadm Bootstrap Providerrs

- InfraProvider: [Docker](https://github.com/kubernetes-sigs/cluster-api-provider-docker)
- BootstrapProvider: [Kubeadm](https://github.com/kubernetes-sigs/cluster-api-bootstrap-provider-kubeadm)
- KubernetesVersion: v1.14.2
- ControlPlaneMachineCount: 3
- WorkerMachineCountt: 3

```(bash)
$ ./capi-yaml generate -b kubeadm -c test-cluster -m 3 -i docker -k v1.14.2 -n awesome-ns -w 3 

```
yaml for this is available here [examples/capd-kubeadm.yaml](examples/capd-kubeadm.yaml)
