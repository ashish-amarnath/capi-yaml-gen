package cmd

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var update = flag.Bool("update", false, "update golden files")

func TestGoldenFiles(t *testing.T) {
	testcases := []struct {
		name       string
		goldenfile string
		options    generateOptions
	}{
		{
			"./capi-yaml-gen generate",
			"default-capd",
			generateOptions{
				infraProvider:            defaultInfrastructureProvider,
				clusterName:              defaultClusterName,
				clusterNamespace:         defaultNamespace,
				bsProvider:               defaultBootstrapProvider,
				k8sVersion:               defaultVersion,
				controlplaneMachineCount: defaultControlPlaneCount,
				workerMachineCount:       defaultWorkerCount,
			},
		},
		{
			"./capi-yaml-gen generate --infra-provider aws",
			"default-capa",
			generateOptions{
				infraProvider:            "aws",
				clusterName:              defaultClusterName,
				clusterNamespace:         defaultNamespace,
				bsProvider:               defaultBootstrapProvider,
				k8sVersion:               defaultVersion,
				controlplaneMachineCount: defaultControlPlaneCount,
				workerMachineCount:       defaultWorkerCount,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			if err := runGenerateCommand(tc.options, &stdout, &stderr); err != nil {
				t.Fatal(err)
			}
			if stderr.Len() > 0 {
				t.Fatal(stderr.String())
			}

			if *update {
				if err := ioutil.WriteFile(goldenFileName(tc.goldenfile), stdout.Bytes(), 0644); err != nil {
					t.Fatal(err)
				}
				return
			}

			golden, err := ioutil.ReadFile(goldenFileName(tc.goldenfile))
			if err != nil {
				t.Fatal(err)
			}
			diff := cmp.Diff(string(golden), stdout.String())
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func goldenFileName(name string) string {
	return fmt.Sprintf("testdata/%s.golden", name)
}
