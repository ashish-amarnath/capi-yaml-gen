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
	"github.com/spf13/cobra"
)

// RootCmd returns the root command for capi-yaml-gen tool
func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "capi-yaml-gen",
		Short: "Yaml generating tool for CAPI and CAPI providers",
		Long:  "Yaml generating tool for CAPI and CAPI providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return rootCmd
}
