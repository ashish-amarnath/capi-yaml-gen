/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"io"
	"strings"

	"github.com/ashish-amarnath/capi-yaml-gen/cmd/cabpk"
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/capa"
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/capd"
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/capi"
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/constants"
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/generator"
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/serialize"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type printMachineParams struct {
	count             int
	infraProvider     generator.InfrastructureProvider
	bootstrapProvider generator.BootstrapProvider
	namePrefix        string
	clusterName       string
	clusterNamespace  string
	k8sVersion        string
	isControlPlane    bool
}

func getInfraProvider(provider string) (generator.InfrastructureProvider, error) {
	switch strings.ToLower(provider) {
	case "docker":
		return capd.Provider{}, nil
	case "aws":
		return capa.Provider{}, nil
	default:
		return nil, fmt.Errorf("Unsupported infrastructure provider %q", provider)
	}
}

func getBootstrapProvider(provider string) (generator.BootstrapProvider, error) {
	switch strings.ToLower(provider) {
	case "kubeadm":
		return cabpk.Provider{}, nil
	default:
		return nil, fmt.Errorf("Unsupported bootstrap provider %q", provider)
	}
}

func configuredMachines(p printMachineParams) ([]runtime.Object, error) {
	out := make([]runtime.Object, 0)
	for i := 0; i < p.count; i++ {
		machineName := fmt.Sprintf("%s-%d", p.namePrefix, i)

		bsConfigName := fmt.Sprintf("%s-config", strings.ToLower(machineName))
		bsConfig := p.bootstrapProvider.GetConfig(bsConfigName, p.clusterNamespace, p.isControlPlane, i)
		infraMachine := p.infraProvider.GetInfraMachine(machineName, p.clusterNamespace)
		p.infraProvider.SetBootstrapConfigInfraValues(bsConfig)

		out = append(out, infraMachine)
		// TODO get rid of if/else
		if p.isControlPlane {
			out = append(out, capi.GetCoreControlPlaneMachine(machineName, p.clusterNamespace, p.clusterName, p.k8sVersion, bsConfig, infraMachine))
		} else {
			out = append(out, capi.GetCoreMachine(machineName, p.clusterNamespace, p.clusterName, p.k8sVersion, bsConfig, infraMachine))
		}
		out = append(out, bsConfig)
	}
	return out, nil
}

func configuredMachineDeployment(p printMachineParams) []runtime.Object {
	machineTemplate := p.infraProvider.GetInfraMachineTemplate(p.namePrefix, p.clusterNamespace)
	configTemplate := p.bootstrapProvider.GetConfigTemplate(p.namePrefix, p.clusterNamespace)
	p.infraProvider.SetBootstrapConfigTemplateInfraValues(configTemplate)
	md := capi.GetCoreMachineDeployment(p.clusterName, p.namePrefix, p.clusterNamespace, p.k8sVersion, int32(p.count), machineTemplate, configTemplate)
	return []runtime.Object{machineTemplate, configTemplate, md}
}

func getNamespace(name string) runtime.Object {
	ns := v1.Namespace{}
	ns.Name = name

	return ns.DeepCopyObject()
}

func runGenerateCommand(opts generateOptions, stdout io.Writer) error {
	items := make([]runtime.Object, 0)
	ip, err := getInfraProvider(opts.infraProvider)
	if err != nil {
		return err
	}
	bp, err := getBootstrapProvider(opts.bsProvider)
	if err != nil {
		return err
	}

	if opts.clusterNamespace != defaultNamespace {
		items = append(items, getNamespace(opts.clusterNamespace))
	}

	infraCluster := ip.GetInfraCluster(opts.clusterName, opts.clusterNamespace)

	coreCluster := capi.GetCoreCluster(opts.clusterName, opts.clusterNamespace, infraCluster)

	pcmControlplane := printMachineParams{
		count:             opts.controlplaneMachineCount,
		infraProvider:     ip,
		bootstrapProvider: bp,
		namePrefix:        "controlplane",
		clusterName:       opts.clusterName,
		clusterNamespace:  opts.clusterNamespace,
		k8sVersion:        opts.k8sVersion,
		isControlPlane:    true,
	}

	pmcWorker := printMachineParams{
		count:             opts.workerMachineCount,
		infraProvider:     ip,
		bootstrapProvider: bp,
		namePrefix:        "worker",
		clusterName:       opts.clusterName,
		clusterNamespace:  opts.clusterNamespace,
		k8sVersion:        opts.k8sVersion,
		isControlPlane:    false,
	}

	controlPlanes, err := configuredMachines(pcmControlplane)
	if err != nil {
		return err
	}
	workers, err := configuredMachines(pmcWorker)
	if err != nil {
		return err
	}
	if opts.machineDeployment {
		pmcWorker.namePrefix += "-md"
		workers = configuredMachineDeployment(pmcWorker)
	}
	items = append(items, infraCluster, coreCluster)
	items = append(items, controlPlanes...)
	items = append(items, workers...)
	for _, item := range items {
		fmt.Fprintln(stdout, constants.YAMLSeperator)
		b, err := serialize.MarshalToYAML(item)
		if err != nil {
			return err
		}
		fmt.Fprintln(stdout, string(b))
	}
	return nil
}
