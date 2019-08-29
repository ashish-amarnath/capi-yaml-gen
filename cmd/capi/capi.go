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

package capi

import (
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/constants"
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/serialize"
	v1 "k8s.io/api/core/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha2"
)

// GetCoreMachineYamlParams is the parameters to generate core machine yaml
type GetCoreMachineYamlParams struct {
	Name                    string
	Namespace               string
	BsConfigName            string
	BsConfigKind            string
	BsProviderAPIVersion    string
	K8sVersion              string
	ClusterOwner            string
	InfraMachineKind        string
	InfraProviderAPIVersion string
	IsControlPlane          bool
}

// GetCoreClusterYaml returns yaml for CAPI cluster objects
func GetCoreClusterYaml(name, namespace, infraClusterKind, infraProviderAPIVersion string) (string, error) {
	coreCluster := &clusterv1.Cluster{}
	coreCluster.Kind = constants.CoreClusterKind
	coreCluster.Name = name
	coreCluster.Namespace = namespace
	coreCluster.APIVersion = clusterv1.GroupVersion.String()

	coreCluster.Spec = clusterv1.ClusterSpec{
		InfrastructureRef: &v1.ObjectReference{
			Kind:       infraClusterKind,
			APIVersion: infraProviderAPIVersion,
			Name:       name,
			Namespace:  namespace,
		},
	}
	yamlBytes, err := serialize.MarshalToYAML(coreCluster)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}

// GetCoreMachineYaml returns yaml for CAPI machine object configured to be a controlplane or not
func GetCoreMachineYaml(p GetCoreMachineYamlParams) (string, error) {
	coreMachine := &clusterv1.Machine{}
	coreMachine.Kind = constants.CoreMachineKind
	coreMachine.APIVersion = clusterv1.GroupVersion.String()
	coreMachine.Name = p.Name
	coreMachine.Namespace = p.Namespace
	lables := map[string]string{
		clusterv1.MachineClusterLabelName: p.ClusterOwner,
	}
	if p.IsControlPlane {
		lables[clusterv1.MachineControlPlaneLabelName] = "true"
	}
	coreMachine.SetLabels(lables)

	coreMachine.Spec = clusterv1.MachineSpec{
		Bootstrap: clusterv1.Bootstrap{
			ConfigRef: &v1.ObjectReference{
				Kind:       p.BsConfigKind,
				APIVersion: p.BsProviderAPIVersion,
				Name:       p.BsConfigName,
				Namespace:  p.Namespace,
			},
		},
		InfrastructureRef: v1.ObjectReference{
			Kind:       p.InfraMachineKind,
			APIVersion: p.InfraProviderAPIVersion,
			Namespace:  p.Namespace,
			Name:       p.Name,
		},
	}

	yamlBytes, err := serialize.MarshalToYAML(coreMachine)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}
