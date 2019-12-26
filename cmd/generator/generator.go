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

package generator

import (
	"k8s.io/apimachinery/pkg/runtime"
	bootstrapv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha2"
)

// Object is a generic type
type Object interface {
	runtime.Object
	GetName() string
	GetNamespace() string
}

// InfrastructureProvider interface for infrastructure providers
type InfrastructureProvider interface {
	GetInfraMachine(name, namespace string) Object
	GetInfraCluster(name, namespace string) Object
	GetInfraMachineTemplate(name, namespace string) Object
	SetBootstrapConfigInfraValues(*bootstrapv1.KubeadmConfig)
	SetBootstrapConfigTemplateInfraValues(*bootstrapv1.KubeadmConfigTemplate)
	GetEnvironmentVariables() map[string]string
}

// BootstrapProvider interface for bootstrap providers
type BootstrapProvider interface {
	GetConfig(name, namespace string, isControlplane bool, count int) *bootstrapv1.KubeadmConfig
	GetConfigTemplate(name, namespace string) *bootstrapv1.KubeadmConfigTemplate
}
