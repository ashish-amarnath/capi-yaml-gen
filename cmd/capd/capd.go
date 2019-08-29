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
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/constants"
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/serialize"
	dockerv1 "sigs.k8s.io/cluster-api-provider-docker/api/v1alpha2"
)

// GetDockerClusterYaml generates yaml for a docker cluster
func GetDockerClusterYaml(name, namespace string) (string, string, string, error) {
	dockerCluster := &dockerv1.DockerCluster{}
	dockerCluster.Kind = dockerCluster.Kind
	dockerCluster.APIVersion = dockerv1.GroupVersion.String()
	dockerCluster.Name = name
	dockerCluster.Namespace = namespace

	yamlBytes, err := serialize.MarshalToYAML(dockerCluster)
	if err != nil {
		return "", "", "", err
	}

	return string(yamlBytes), constants.DockerClusterKind, dockerCluster.APIVersion, nil
}

// GetDockerMachineYaml generates yaml for a docker controlplane machine
func GetDockerMachineYaml(name, namespace string) (string, string, string, error) {
	dockerMachine := &dockerv1.DockerMachine{}
	dockerMachine.Kind = constants.DockerMachineKind
	dockerMachine.APIVersion = dockerv1.GroupVersion.String()
	dockerMachine.Name = name
	dockerMachine.Namespace = namespace

	yamlBytes, err := serialize.MarshalToYAML(dockerMachine)
	if err != nil {
		return "", "", "", err
	}

	return string(yamlBytes), constants.DockerMachineKind, dockerMachine.APIVersion, nil
}
