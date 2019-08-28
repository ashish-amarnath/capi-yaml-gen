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
	"os"
	"strings"

	"github.com/ashish-amarnath/capiyaml/cmd/cabpk"
	"github.com/ashish-amarnath/capiyaml/cmd/capa"
	"github.com/ashish-amarnath/capiyaml/cmd/capd"
	"github.com/ashish-amarnath/capiyaml/cmd/capi"
	"github.com/ashish-amarnath/capiyaml/cmd/constants"
)

func getInfraClusterYaml(infraProvider, cName, cNamespace string) (string, string, error) {
	var err error
	var infraClusterYaml, infraClusterKind string
	switch strings.ToLower(infraProvider) {
	case "docker":
		infraClusterYaml, infraClusterKind, err = capd.GetDockerClusterYaml(cName, cNamespace)
	case "aws":
		infraClusterYaml, infraClusterKind, err = capa.GetAWSClusterYaml(cName, cNamespace)
	default:
		return "", "", fmt.Errorf("Unsupported cluster infrastructure provider %q", infraProvider)
	}

	return infraClusterYaml, infraClusterKind, err
}

func getBoostrapProviderConfigYaml(bsProvider, bsConfigName, cNamespace, k8sVersion string, isControlPlane bool, itemNumber int) (string, string, error) {
	switch strings.ToLower((bsProvider)) {
	case "kubeadm":
		// TODO: use k8sversion but have to figure out if we need ClusterConfig/InitConfig or JoinConfig
		return cabpk.GetBootstrapProviderConfig(bsConfigName, cNamespace, isControlPlane, itemNumber)
	default:
		return "", "", fmt.Errorf("Unsupported bootstrap provider %q", bsProvider)
	}
}

func getInfraMachineYaml(infraProvider, mName, mNamespace string) (string, string, error) {
	var err error
	var infraCPMachineYaml, infraCPMachineKind string

	switch strings.ToLower(infraProvider) {
	case "docker":
		infraCPMachineYaml, infraCPMachineKind, err = capd.GetDockerMachineYaml(mName, mNamespace)
	case "aws":
		infraCPMachineYaml, infraCPMachineKind, err = capa.GetAWSMachineYaml(mName, mNamespace)
	default:
		return "", "", fmt.Errorf("Unsupported machine infrastructure provider %q", infraProvider)
	}

	return infraCPMachineYaml, infraCPMachineKind, err
}

func printMachineYaml(p printMachineParams) {
	for i := 0; i < p.count; i++ {
		machineName := fmt.Sprintf("%s-%d", p.namePrefix, i)

		bsConfigName := fmt.Sprintf("%s-config", strings.ToLower(machineName))
		bsConfigYAML, bsConfigKind, err := getBoostrapProviderConfigYaml(p.bootstrapProvider, bsConfigName, p.clusterNamespace, p.k8sVersion, p.isControlPlane, i)

		infraMachineYaml, infraMachineKind, err := getInfraMachineYaml(p.infraProvider,
			machineName, p.clusterNamespace)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to generate yaml for infrastructure machine, %v\n", err)
			os.Exit(1)
		}

		coreMachineYaml, err := capi.GetCoreMachineYaml(
			machineName, p.clusterNamespace, bsConfigName, bsConfigKind, p.k8sVersion,
			p.clusterName, infraMachineKind, p.isControlPlane)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to generate yaml for core machine, %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, strings.TrimSpace(infraMachineYaml))
		fmt.Fprintf(os.Stdout, constants.YAMLSeperator)
		fmt.Fprintf(os.Stdout, strings.TrimSpace(coreMachineYaml))
		fmt.Fprintf(os.Stdout, constants.YAMLSeperator)
		fmt.Fprintf(os.Stdout, strings.TrimSpace(bsConfigYAML))
		fmt.Fprintf(os.Stdout, constants.YAMLSeperator)
	}
}

func runGenerateCommand(opts generateOptions) {
	infraClusterYaml, infraClusterKind, err := getInfraClusterYaml(opts.infraProvider, opts.clusterName, opts.clusterNamespace)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate yaml for infrastructure cluster, %v\n", err)
		os.Exit(1)
	}

	coreClusterYaml, err := capi.GetCoreClusterYaml(opts.clusterName, opts.clusterNamespace, infraClusterKind)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate yaml for core cluster, %v\n", err)
		os.Exit(1)
	}

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

	fmt.Fprintf(os.Stdout, constants.YAMLSeperator)
	fmt.Fprintf(os.Stdout, "%s", strings.TrimSpace(infraClusterYaml))
	fmt.Fprintf(os.Stdout, constants.YAMLSeperator)
	fmt.Fprintf(os.Stdout, "%s", strings.TrimSpace(coreClusterYaml))
	fmt.Fprintf(os.Stdout, constants.YAMLSeperator)
	printMachineYaml(pcmControlplane)
	printMachineYaml(pmcWorker)
}
