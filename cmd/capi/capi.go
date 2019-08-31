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
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha2"
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
	coreCluster.Spec.InfrastructureRef = referenceToObjectRef(infraCluster)
	return coreCluster
}

// GetCoreMachine returns a CAPI machine worker object
func GetCoreMachine(name, namespace, clusterName string, bootstrapConfig, infraMachine object) *clusterv1.Machine {
	coreMachine := &clusterv1.Machine{}
	coreMachine.APIVersion = clusterv1.GroupVersion.String()
	coreMachine.Name = name
	coreMachine.Namespace = namespace
	labels := map[string]string{
		clusterv1.MachineClusterLabelName: clusterName,
	}
	coreMachine.SetLabels(labels)
	coreMachine.Spec.Bootstrap.ConfigRef = referenceToObjectRef(bootstrapConfig)
	coreMachine.Spec.InfrastructureRef = *referenceToObjectRef(infraMachine)
	return coreMachine
}

// GetCoreControlPlaneMachine returns a cluster-api machine that identifies as a
// control plane node
func GetCoreControlPlaneMachine(name, namespace, clusterName string, bootstrapConfig, infraMachine object) *clusterv1.Machine {
	machine := GetCoreMachine(name, namespace, clusterName, bootstrapConfig, infraMachine)
	machine.Labels[clusterv1.MachineControlPlaneLabelName] = "true"
	return machine
}
