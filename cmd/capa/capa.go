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

	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha2"
)

// GetAWSCluster generates an AWS cluster
func GetAWSCluster(name, namespace string) *infrav1.AWSCluster {
	awsCluster := &infrav1.AWSCluster{}
	awsCluster.Kind = constants.AWSClusterKind
	awsCluster.APIVersion = infrav1.GroupVersion.String()
	awsCluster.Name = name
	awsCluster.Namespace = namespace
	return awsCluster
}

// GetAWSMachine generates an AWS machine
func GetAWSMachine(name, namespace string) *infrav1.AWSMachine {
	awsMachine := &infrav1.AWSMachine{}
	awsMachine.Kind = constants.AWSMachineKind
	awsMachine.APIVersion = infrav1.GroupVersion.String()
	awsMachine.Name = name
	awsMachine.Namespace = namespace
	return awsMachine
}
