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
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/constants"
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/generator"
	bootstrapv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha2"
	bootstrapv1beta1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/types/v1beta1"
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
	awsCluster.Spec.Region = "${REGION}"
	awsCluster.Spec.SSHKeyName = "${SSH_KEY_NAME}"
	return awsCluster
}

// GetInfraMachine generates an AWS machine
func (p Provider) GetInfraMachine(name, namespace string) generator.Object {
	awsMachine := &infrav1.AWSMachine{}
	awsMachine.Kind = constants.AWSMachineKind
	awsMachine.APIVersion = infrav1.GroupVersion.String()
	awsMachine.Name = name
	awsMachine.Namespace = namespace
	awsMachine.Spec = infrav1.AWSMachineSpec{
		InstanceType:       "${CONTROL_PLANE_INSTANCE_TYPE}",
		IAMInstanceProfile: "control-plane.cluster-api-provider-aws.sigs.k8s.io",
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
	template.Spec.Template.Spec.InstanceType = "${MACHINE_DEPLOYMENT_INSTANCE_TYPE}"
	template.Spec.Template.Spec.IAMInstanceProfile = "nodes.cluster-api-provider-aws.sigs.k8s.io"
	template.Spec.Template.Spec.SSHKeyName = "${SSH_KEY_NAME}"
	return template
}

// SetBootstrapConfigInfraValues fills in InfraProvider specific values into the bootstrap config
func (p Provider) SetBootstrapConfigInfraValues(c *bootstrapv1.KubeadmConfig) {
	extraArgs := map[string]string{
		"cloud-provider": "aws",
	}
	if c.Spec.InitConfiguration != nil {
		c.Spec.InitConfiguration.NodeRegistration = bootstrapv1beta1.NodeRegistrationOptions{
			Name:             "{{ ds.meta_data.hostname }}",
			KubeletExtraArgs: extraArgs,
		}
	} else if c.Spec.JoinConfiguration != nil {
		c.Spec.JoinConfiguration.NodeRegistration = bootstrapv1beta1.NodeRegistrationOptions{
			Name:             "{{ ds.meta_data.hostname }}",
			KubeletExtraArgs: extraArgs,
		}
	}

	if c.Spec.ClusterConfiguration != nil {
		c.Spec.ClusterConfiguration.APIServer = bootstrapv1beta1.APIServer{
			ControlPlaneComponent: bootstrapv1beta1.ControlPlaneComponent{
				ExtraArgs: extraArgs,
			},
		}

		c.Spec.ClusterConfiguration.ControllerManager = bootstrapv1beta1.ControlPlaneComponent{
			ExtraArgs: extraArgs,
		}
	}
}

// SetBootstrapConfigTemplateInfraValues fills in InfraProvider specific values into the join configuration
func (p Provider) SetBootstrapConfigTemplateInfraValues(t *bootstrapv1.KubeadmConfigTemplate) {
	extraArgs := map[string]string{
		"cloud-provider": "aws",
	}
	if t.Spec.Template.Spec.JoinConfiguration != nil {
		t.Spec.Template.Spec.JoinConfiguration.NodeRegistration = bootstrapv1beta1.NodeRegistrationOptions{
			Name:             "{{ ds.meta_data.hostname }}",
			KubeletExtraArgs: extraArgs,
		}
	}
}

func (p Provider) GetEnvironmentVariables() map[string]string {
	return map[string]string{
		"SSH_KEY_NAME":                     "default",
		"CONTROL_PLANE_INSTANCE_TYPE":      "t2.medium",
		"MACHINE_DEPLOYMENT_INSTANCE_TYPE": "t2.medium",
		"REGION":                           "us-west-2",
	}
}
