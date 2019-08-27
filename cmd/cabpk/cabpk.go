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

package cabpk

import (
	"github.com/ashish-amarnath/capiyaml/cmd/constants"
	"github.com/ashish-amarnath/capiyaml/cmd/serialize"
	infrav1 "sigs.k8s.io/cluster-api-bootstrap-provider-kubeadm/api/v1alpha2"
)

// GetBootstrapProviderConfig generates kubeadm bootstrap provider config
func GetBootstrapProviderConfig(name, namespace string) (string, string, error) {
	bsConfig := &infrav1.KubeadmConfig{}
	bsConfig.Name = name
	bsConfig.Namespace = namespace
	bsConfig.APIVersion = constants.BootstrapProviderAPIVersion

	yamlBytes, err := serialize.MarshalToYAML(bsConfig)
	if err != nil {
		return "", "", err
	}

	return string(yamlBytes), constants.KubeadmConfigKind, nil
}
