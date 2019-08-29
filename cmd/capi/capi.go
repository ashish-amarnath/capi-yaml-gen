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

// GetCoreClusterYaml returns yaml for CAPI cluster objects
func GetCoreClusterYaml(name, namespace, infraClusterKind string) (string, error) {
	coreCluster := &clusterv1.Cluster{}
	coreCluster.Kind = constants.CoreClusterKind
	coreCluster.Name = name
	coreCluster.Namespace = namespace
	coreCluster.APIVersion = constants.CoreAPIVersion

	coreCluster.Spec = clusterv1.ClusterSpec{
		InfrastructureRef: &v1.ObjectReference{
			Kind:       infraClusterKind,
			APIVersion: constants.InfrastructureProviderAPIVersion,
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
func GetCoreMachineYaml(name, namespace, bsConfigName, bsConfigKind, version, clusterOwner,
	infraMachineKind string, controlPlane bool) (string, error) {
	coreMachine := &clusterv1.Machine{}
	coreMachine.Kind = constants.CoreMachineKind
	coreMachine.APIVersion = constants.CoreAPIVersion
	coreMachine.Name = name
	coreMachine.Namespace = namespace
	lables := map[string]string{
		clusterv1.MachineClusterLabelName: clusterOwner,
	}
	if controlPlane {
		lables[clusterv1.MachineControlPlaneLabelName] = "true"
	}
	coreMachine.SetLabels(lables)

	coreMachine.Spec = clusterv1.MachineSpec{
		Bootstrap: clusterv1.Bootstrap{
			ConfigRef: &v1.ObjectReference{
				Kind:       bsConfigKind,
				APIVersion: constants.BootstrapProviderAPIVersion,
				Name:       bsConfigName,
				Namespace:  namespace,
			},
		},
		InfrastructureRef: v1.ObjectReference{
			Kind:       infraMachineKind,
			APIVersion: constants.InfrastructureProviderAPIVersion,
			Namespace:  namespace,
			Name:       name,
		},
	}

	yamlBytes, err := serialize.MarshalToYAML(coreMachine)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}
