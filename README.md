# capi-yaml

Developer tool to generate yaml for Cluster-API and necessary provider resources.

Read more about this [here](https://docs.google.com/document/d/1Tzx6IXOoQnUxaVSYA2I8IdcNrFE4zgOnzkk55KHOU20/edit)

## Usage Examples

### Docker Infrastructure and Kubeadm Bootstrap Providerrs

- InfraProvider: [Docker](https://github.com/kubernetes-sigs/cluster-api-provider-docker)
- BootstrapProvider: [Kubeadm](https://github.com/kubernetes-sigs/cluster-api-bootstrap-provider-kubeadm)
- KubernetesVersion: v1.14.2
- ControlPlaneMachineCount: 3
- WorkerMachineCountt: 1

```(bash)
$ go run main.go generate --controlplane-count 3
# yaml written to stdout
```

### AWS Infrastructure and Kubeadm Bootstrap Providers

- InfraProvider: [AWS](https://github.com/kubernetes-sigs/cluster-api-provider-aws)
- BootstrapProvider: [Kubeadm](https://github.com/kubernetes-sigs/cluster-api-bootstrap-provider-kubeadm)
- KubernetesVersion: v1.14.2
- ControlPlaneMachineCount: 3
- WorkerMachineCountt: 1

```(bash)
$ go run main.go generate --control-plane-count 3 --infrastructure-provider aws
# yaml written to stdout
```

### MachineDeployments

By default workers will be managed by a MachineDeployment. If you do not want this behavior set the
`--generate-machined-deployment` flag to false like this:

```bash
go run main.go generate --generate-machine-deployment=false
```
