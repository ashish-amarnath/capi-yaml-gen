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

package capa

import (
	"fmt"
	"strings"

	"github.com/ashish-amarnath/capi-yaml-gen/cmd/constants"
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/generator"
	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha2"
)

// Provider CAPA implementation of InfrastructureProvider
type Provider struct{}

// GetInfraCluster generates an AWS cluster
func (p Provider) GetInfraCluster(name, namespace string) generator.Object {
	awsCluster := &infrav1.AWSCluster{}
	awsCluster.Kind = constants.AWSClusterKind
	awsCluster.APIVersion = infrav1.GroupVersion.String()
	awsCluster.Name = name
	awsCluster.Namespace = namespace
	return awsCluster
}

// GetInfraMachine generates an AWS machine
func (p Provider) GetInfraMachine(name, namespace string) generator.Object {
	awsMachine := &infrav1.AWSMachine{}
	awsMachine.Kind = constants.AWSMachineKind
	awsMachine.APIVersion = infrav1.GroupVersion.String()
	awsMachine.Name = name
	awsMachine.Namespace = namespace
	// TODO (ashish-amarnath) lookup these values from an map, also avoid per machine values
	awsMachine.Spec = infrav1.AWSMachineSpec{
		InstanceType:       strings.ToUpper(fmt.Sprintf("${%s_%s_AWS_EC2_INSTANCE_TYPE}", namespace, name)),
		IAMInstanceProfile: strings.ToUpper(fmt.Sprintf("${%s_%s_AWS_IAM_INSTANCE_PROFILE}", namespace, name)),
		SSHKeyName:         "${SSH_KEY_NAME}",
	}
	return awsMachine
}

// GetInfraMachineTemplate generates an AWS machine template
func (p Provider) GetInfraMachineTemplate(name, namespace string) generator.Object {
	template := &infrav1.AWSMachineTemplate{}
	template.Name = name
	template.Namespace = namespace
	template.Kind = constants.AWSMachineKind + "Template"
	template.APIVersion = infrav1.GroupVersion.String()
	return template
}
