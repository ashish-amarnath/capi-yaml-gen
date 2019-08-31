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
	infrav1 "sigs.k8s.io/cluster-api-provider-docker/api/v1alpha2"
)

// GetDockerCluster generates a docker cluster
func GetDockerCluster(name, namespace string) *infrav1.DockerCluster {
	dockerCluster := &infrav1.DockerCluster{}
	dockerCluster.Kind = constants.DockerClusterKind
	dockerCluster.APIVersion = infrav1.GroupVersion.String()
	dockerCluster.Name = name
	dockerCluster.Namespace = namespace
	return dockerCluster
}

// GetDockerMachine generates a docker machine
func GetDockerMachine(name, namespace string) *infrav1.DockerMachine {
	dockerMachine := &infrav1.DockerMachine{}
	dockerMachine.Kind = constants.DockerMachineKind
	dockerMachine.APIVersion = infrav1.GroupVersion.String()
	dockerMachine.Name = name
	dockerMachine.Namespace = namespace
	return dockerMachine
}
