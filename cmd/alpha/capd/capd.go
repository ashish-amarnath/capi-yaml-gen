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

package capd

import (
	"github.com/ashish-amarnath/capiyaml/cmd/alpha/constants"
	caapdv1 "sigs.k8s.io/cluster-api-provider-docker/api/v1alpha2"
	"sigs.k8s.io/yaml"
)

// GetDockerClusterYaml generates yaml for a docker cluster
func GetDockerClusterYaml(name, namespace string) (string, string, error) {
	dockerCluster := caapdv1.DockerCluster{}
	dockerCluster.Kind = constants.DockerClusterKind
	dockerCluster.APIVersion = constants.InfrastructureProviderAPIVersion
	dockerCluster.Name = name
	dockerCluster.Namespace = namespace

	yamlBytes, err := yaml.Marshal(dockerCluster)
	if err != nil {
		return "", "", err
	}

	return string(yamlBytes), constants.DockerClusterKind, nil
}
