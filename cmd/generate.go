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
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/serialize"

	"k8s.io/apimachinery/pkg/runtime"
)

type printMachineParams struct {
	count             int
	infraProvider     string
	bootstrapProvider string
	namePrefix        string
	clusterName       string
	clusterNamespace  string
	k8sVersion        string
	isControlPlane    bool
}

type object interface {
	runtime.Object
	GetName() string
	GetNamespace() string
}

func getInfraCluster(provider, name, namespace string) (object, error) {
	switch strings.ToLower(provider) {
	case "docker":
		return capd.GetDockerCluster(name, namespace), nil
	case "aws":
		return capa.GetAWSCluster(name, namespace), nil
	default:
		return nil, fmt.Errorf("Unsupported cluster infrastructure provider %q", provider)
	}
}

func getBoostrapProviderConfig(provider, name, namespace string, isControlPlane bool, itemNumber int) (object, error) {
	switch strings.ToLower(provider) {
	case "kubeadm":
		return cabpk.GetBootstrapProviderConfig(name, namespace, isControlPlane, itemNumber), nil
	default:
		return nil, fmt.Errorf("Unsupported bootstrap provider %q", provider)
	}
}

func getInfraMachine(infraProvider, mName, mNamespace string) (object, error) {
	switch strings.ToLower(infraProvider) {
	case "docker":
		return capd.GetDockerMachine(mName, mNamespace), nil
	case "aws":
		return capa.GetAWSMachine(mName, mNamespace), nil
	default:
		return nil, fmt.Errorf("Unsupported machine infrastructure provider %q", infraProvider)
	}
}

func configuredMachines(p printMachineParams) ([]runtime.Object, error) {
	out := make([]runtime.Object, 0)
	for i := 0; i < p.count; i++ {
		machineName := fmt.Sprintf("%s-%d", p.namePrefix, i)

		bsConfigName := fmt.Sprintf("%s-config", strings.ToLower(machineName))
		bsConfig, err := getBoostrapProviderConfig(p.bootstrapProvider, bsConfigName, p.clusterNamespace, p.isControlPlane, i)
		if err != nil {
			return nil, err
		}
		infraMachine, err := getInfraMachine(p.infraProvider, machineName, p.clusterNamespace)
		if err != nil {
			return nil, err
		}
		out = append(out, infraMachine)
		// TODO get rid of if/else
		if p.isControlPlane {
			out = append(out, capi.GetCoreControlPlaneMachine(machineName, p.clusterNamespace, p.clusterName, bsConfig, infraMachine))
		} else {
			out = append(out, capi.GetCoreMachine(machineName, p.clusterNamespace, p.clusterName, bsConfig, infraMachine))
		}
		out = append(out, bsConfig)
	}
	return out, nil
}

func runGenerateCommand(opts generateOptions, stdout, stderr io.Writer) error {
	items := make([]runtime.Object, 0)
	infraCluster, err := getInfraCluster(opts.infraProvider, opts.clusterName, opts.clusterNamespace)
	if err != nil {
		return err
	}

	coreCluster := capi.GetCoreCluster(opts.clusterName, opts.clusterNamespace, infraCluster)

	pcmControlplane := printMachineParams{
		count:             opts.controlplaneMachineCount,
		infraProvider:     opts.infraProvider,
		bootstrapProvider: opts.bsProvider,
		namePrefix:        "controlplane",
		clusterName:       opts.clusterName,
		clusterNamespace:  opts.clusterNamespace,
		k8sVersion:        opts.k8sVersion,
		isControlPlane:    true,
	}

	pmcWorker := printMachineParams{
		count:             opts.workerMachineCount,
		infraProvider:     opts.infraProvider,
		bootstrapProvider: opts.bsProvider,
		namePrefix:        "worker",
		clusterName:       opts.clusterName,
		clusterNamespace:  opts.clusterNamespace,
		k8sVersion:        opts.k8sVersion,
		isControlPlane:    false,
	}

	controlPlanes, err := configuredMachines(pcmControlplane)
	workers, err := configuredMachines(pmcWorker)
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
