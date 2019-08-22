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

package alpha

import (
	"fmt"
	"os"

	"github.com/ashish-amarnath/capiyaml/cmd/alpha/capd"
	"github.com/ashish-amarnath/capiyaml/cmd/alpha/capi"
	"github.com/ashish-amarnath/capiyaml/cmd/alpha/constants"
	"github.com/spf13/cobra"
)

// Cmd returns the set alpha of commands
func Cmd() *cobra.Command {
	alphaCmd := &cobra.Command{
		Use:   "alpha",
		Short: "capi-yaml-gen alpha command set",
		Long:  "capi-yaml-gen alpha command set",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	alphaCmd.AddCommand(getClusterYamlCommand())
	return alphaCmd
}

func getClusterYamlForInfraProvider(infraProvider, cName, cNamespace string) (string, string, error) {
	var err error
	var infraClusterYaml, infraClusterKind string
	switch infraProvider {
	case "docker":
		infraClusterYaml, infraClusterKind, err = capd.GetDockerClusterYaml(cName, cNamespace)
	default:
		return "", "", fmt.Errorf("Unsupported provider %q", infraProvider)
	}

	return infraClusterYaml, infraClusterKind, err
}

func getClusterYamlCommand() *cobra.Command {
	var infraProvider, clusterName, clusterNamespace string

	cmd := &cobra.Command{
		Use:   "get-cluster-yaml",
		Short: "generate yaml for cluster",
		Long:  "generate yaml for cluster",
		Run: func(cmd *cobra.Command, args []string) {
			infraClusterYaml, infraClusterKind, err := getClusterYamlForInfraProvider(infraProvider, clusterName, clusterNamespace)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to generate yaml for infrastructure cluster, %v", err)
				os.Exit(1)
			}

			coreClusterYaml, err := capi.GetCoreClusterYaml(clusterName, clusterNamespace, infraClusterKind)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to generate yaml for core cluster, %v\n", err)
				os.Exit(1)
			}

			fmt.Fprintf(os.Stdout, constants.YAMLSeperator)
			fmt.Fprintf(os.Stdout, "%s", infraClusterYaml)
			fmt.Fprintf(os.Stdout, constants.YAMLSeperator)
			fmt.Fprintf(os.Stdout, "%s", coreClusterYaml)
			fmt.Fprintf(os.Stdout, constants.YAMLSeperator)
		},
	}

	cmd.Flags().StringVarP(&infraProvider, "provider", "p", "", "Infrastructure provider for which yaml needs to be generated")
	cmd.Flags().StringVarP(&clusterName, "name", "n", "", "Name for the cluster")
	cmd.Flags().StringVarP(&clusterNamespace, "namespace", "", "default", "Namespace where the cluster will be created")

	return cmd
}
