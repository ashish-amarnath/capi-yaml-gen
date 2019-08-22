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
	"sigs.k8s.io/yaml"

	"github.com/ashish-amarnath/capiyaml/cmd/alpha/constants"
	v1 "k8s.io/api/core/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha2"
)

// GetCoreClusterYaml returns yaml for CAPI  cluster objects
func GetCoreClusterYaml(name, namespace, infraClusterKind string) (string, error) {
	coreCluster := clusterv1.Cluster{}
	coreCluster.Kind = "Cluster"
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

	yamlBytes, err := yaml.Marshal(coreCluster)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}
