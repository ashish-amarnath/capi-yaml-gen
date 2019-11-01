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
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
)

type object interface {
	GetObjectKind() schema.ObjectKind
	GetName() string
	GetNamespace() string
}

func referenceToObjectRef(r object) *v1.ObjectReference {
	return &v1.ObjectReference{
		Kind:       r.GetObjectKind().GroupVersionKind().GroupKind().Kind,
		APIVersion: r.GetObjectKind().GroupVersionKind().GroupVersion().String(),
		Name:       r.GetName(),
		Namespace:  r.GetNamespace(),
	}
}

// GetCoreCluster returns a CAPI cluster object
func GetCoreCluster(name, namespace string, infraCluster object) *clusterv1.Cluster {
	coreCluster := &clusterv1.Cluster{}
	coreCluster.Name = name
	coreCluster.Namespace = namespace
	coreCluster.APIVersion = clusterv1.GroupVersion.String()
	coreCluster.Spec = clusterv1.ClusterSpec{
		InfrastructureRef: referenceToObjectRef(infraCluster),
		ClusterNetwork: &clusterv1.ClusterNetwork{
			Pods: &clusterv1.NetworkRanges{
				CIDRBlocks: []string{"192.168.0.0/16"},
			},
			Services: &clusterv1.NetworkRanges{
				CIDRBlocks: []string{},
			},
		},
	}
	return coreCluster
}

// GetCoreMachine returns a CAPI machine worker object
func GetCoreMachine(name, namespace, clusterName, version string, bootstrapConfig, infraMachine object) *clusterv1.Machine {
	coreMachine := &clusterv1.Machine{}
	coreMachine.APIVersion = clusterv1.GroupVersion.String()
	coreMachine.Name = name
	coreMachine.Namespace = namespace
	labels := map[string]string{
		clusterv1.ClusterLabelName: clusterName,
	}
	coreMachine.SetLabels(labels)
	coreMachine.Spec.Bootstrap.ConfigRef = referenceToObjectRef(bootstrapConfig)
	coreMachine.Spec.InfrastructureRef = *referenceToObjectRef(infraMachine)
	coreMachine.Spec.Version = &version
	coreMachine.Spec.ClusterName = clusterName
	return coreMachine
}

// GetCoreControlPlaneMachine returns a cluster-api machine that identifies as a
// control plane node
func GetCoreControlPlaneMachine(name, namespace, clusterName, version string, bootstrapConfig, infraMachine object) *clusterv1.Machine {
	machine := GetCoreMachine(name, namespace, clusterName, version, bootstrapConfig, infraMachine)
	machine.Labels[clusterv1.MachineControlPlaneLabelName] = "true"
	return machine
}

// GetCoreMachineDeployment returns a cluster-api machine deployment object
func GetCoreMachineDeployment(clusterName, name, namespace, version string, replicas int32, machineTemplate, bootstrapConfigTemplate object) *clusterv1.MachineDeployment {
	dep := &clusterv1.MachineDeployment{}
	dep.Name = name
	dep.Namespace = namespace
	dep.Spec.Replicas = &replicas
	dep.Spec.ClusterName = clusterName

	labels := map[string]string{
		clusterv1.ClusterLabelName: clusterName,
	}
	dep.Spec.Selector.MatchLabels = labels
	dep.Spec.Template.ObjectMeta.Labels = labels
	dep.Spec.Template.Spec.Version = &version
	dep.Spec.Template.Spec.ClusterName = clusterName

	// Set the machine template
	dep.Spec.Template.Spec.InfrastructureRef.Name = machineTemplate.GetName()
	dep.Spec.Template.Spec.InfrastructureRef.Namespace = machineTemplate.GetNamespace()
	dep.Spec.Template.Spec.InfrastructureRef.Kind = machineTemplate.GetObjectKind().GroupVersionKind().Kind
	dep.Spec.Template.Spec.InfrastructureRef.APIVersion = machineTemplate.GetObjectKind().GroupVersionKind().GroupVersion().String()

	// Set the bootstrap config template
	configTemplate := v1.ObjectReference{
		Kind:       bootstrapConfigTemplate.GetObjectKind().GroupVersionKind().Kind,
		Namespace:  bootstrapConfigTemplate.GetNamespace(),
		Name:       bootstrapConfigTemplate.GetName(),
		APIVersion: bootstrapConfigTemplate.GetObjectKind().GroupVersionKind().GroupVersion().String(),
	}
	dep.Spec.Template.Spec.Bootstrap.ConfigRef = &configTemplate

	return dep
}
