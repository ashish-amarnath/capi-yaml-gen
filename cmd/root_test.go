package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDefault(t *testing.T) {
	defaultOptions := generateOptions{
		infraProvider:            defaultInfrastructureProvider,
		clusterName:              defaultClusterName,
		clusterNamespace:         defaultNamespace,
		bsProvider:               defaultBootstrapProvider,
		k8sVersion:               defaultVersion,
		controlplaneMachineCount: defaultControlPlaneCount,
		workerMachineCount:       defaultWorkerCount,
	}
	// default generate options
	var stdout, stderr bytes.Buffer
	if err := runGenerateCommand(defaultOptions, &stdout, &stderr); err != nil {
		t.Fatal(err)
	}
	if stderr.Len() > 0 {
		t.Fatal(stderr.String())
	}
	golden, err := ioutil.ReadFile("testdata/default.golden")
	if err != nil {
		t.Fatal(err)
	}
	diff := cmp.Diff(string(golden), stdout.String())
	if diff != "" {
		t.Fatal(diff)
	}
}
