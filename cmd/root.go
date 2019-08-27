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

package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type printMachineParams struct {
	count            int16
	infraProvider    string
	namePrefix       string
	clusterName      string
	clusterNamespace string
	bsConfigName     string
	bsConfigKind     string
	k8sVersion       string
	isControlPlane   bool
}

type generateOptions struct {
	infraProvider            string
	clusterName              string
	clusterNamespace         string
	bsProvider               string
	k8sVersion               string
	controlplaneMachineCount int16
	workerMachineCount       int16
}

// RootCmd returns the root command for capi-yaml-gen tool
func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "capi-yaml",
		Short: "Devtool to help with YAML for CAPI and CAPI providers",
		Long:  "Devtool to help with YAML for CAPI and CAPI providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Help(); err != nil {
				return err
			}
			return nil
		},
	}
	rootCmd.AddCommand(getGenerateCommand())

	return rootCmd
}

func getGenerateCommand() *cobra.Command {
	opts := generateOptions{}

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "generate yaml for CAPI and its providers",
		Long:  "generate yaml for CAPI and its providers",
		Run: func(cmd *cobra.Command, args []string) {
			runGenerateCommand(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.clusterName, "cluster-name", "c", "my-cluster", "Name for the cluster")
	cmd.Flags().StringVarP(&opts.clusterNamespace, "namespace", "n", "default", "Namespace where the cluster will be created")
	cmd.Flags().StringVarP(&opts.infraProvider, "infra-provider", "i", "docker", "Infrastructure provider for the cluster")
	cmd.Flags().StringVarP(&opts.bsProvider, "boostrap-provider", "b", "kubeadm", "Bootstrap provider for the cluster")
	cmd.Flags().StringVarP(&opts.k8sVersion, "k8s-version", "k", "v1.14.2", "Version of kubernetes for the cluster")

	cmd.Flags().Int16VarP(&opts.controlplaneMachineCount, "controlplane-count", "m", 1, "Number of controlplane machines in the cluster")
	cmd.Flags().Int16VarP(&opts.workerMachineCount, "worker-count", "w", 1, "Number of worker machines in the cluster")

	return cmd
}

// Execute starts the process
func Execute() {
	if err := flag.CommandLine.Parse([]string{}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := RootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}
