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
	"github.com/ashish-amarnath/capi-yaml-gen/cmd/serialize"
	infrav1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha2"
)

// GetAWSClusterYaml generates yaml for an AWS cluster
func GetAWSClusterYaml(name, namespace string) (string, string, error) {
	awsCluster := &infrav1.AWSCluster{}

	awsCluster.Kind = constants.AWSClusterKind
	awsCluster.APIVersion = infrav1.GroupVersion.String()
	awsCluster.Name = name
	awsCluster.Namespace = namespace

	yamlBytes, err := serialize.MarshalToYAML(awsCluster)
	if err != nil {
		return "", "", err
	}

	return string(yamlBytes), awsCluster.Kind, nil
}

// GetAWSMachineYaml generates yaml for a docker controlplane machine
func GetAWSMachineYaml(name, namespace string) (string, string, error) {
	awsMachine := &infrav1.AWSMachine{}
	awsMachine.Kind = constants.AWSMachineKind
	awsMachine.APIVersion = infrav1.GroupVersion.String()
	awsMachine.Name = name
	awsMachine.Namespace = namespace

	yamlBytes, err := serialize.MarshalToYAML(awsMachine)
	if err != nil {
		return "", "", err
	}

	return string(yamlBytes), awsMachine.Kind, nil
}
