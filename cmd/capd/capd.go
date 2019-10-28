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
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/generator"
	bootstrapv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha2"
	infrav1 "sigs.k8s.io/cluster-api/test/infrastructure/docker/api/v1alpha2"
)

// Provider CAPD implementation of InfrastructureProvider
type Provider struct{}

// GetInfraCluster generates a docker cluster
func (p Provider) GetInfraCluster(name, namespace string) generator.Object {
	dockerCluster := &infrav1.DockerCluster{}
	dockerCluster.Kind = constants.DockerClusterKind
	dockerCluster.APIVersion = infrav1.GroupVersion.String()
	dockerCluster.Name = name
	dockerCluster.Namespace = namespace
	return dockerCluster
}

// GetInfraMachine generates a docker machine
func (p Provider) GetInfraMachine(name, namespace string) generator.Object {
	dockerMachine := &infrav1.DockerMachine{}
	dockerMachine.Kind = constants.DockerMachineKind
	dockerMachine.APIVersion = infrav1.GroupVersion.String()
	dockerMachine.Name = name
	dockerMachine.Namespace = namespace
	return dockerMachine
}

// GetInfraMachineTemplate returns a docker machine template
func (p Provider) GetInfraMachineTemplate(name, namespace string) generator.Object {
	template := &infrav1.DockerMachineTemplate{}
	template.Name = name
	template.Namespace = namespace
	template.Kind = constants.DockerMachineKind + "Template"
	template.APIVersion = infrav1.GroupVersion.String()
	return template
}

// SetBootstrapConfigInfraValues does nothing atm
func (p Provider) SetBootstrapConfigInfraValues(*bootstrapv1.KubeadmConfig) {
	// do nothing
	return
}

// SetBootstrapConfigTemplateInfraValues does nothing atm
func (p Provider) SetBootstrapConfigTemplateInfraValues(*bootstrapv1.KubeadmConfigTemplate) {
	// do nothing
	return
}

func (p Provider) GetEnvironmentVariables() map[string]string {
	return map[string]string{}
}
