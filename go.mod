module github.com/ashish-amarnath/capiyaml

go 1.12

require (
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/cluster-api v0.0.0-20190826205838-47413a956cbf
	sigs.k8s.io/cluster-api-bootstrap-provider-kubeadm v0.0.0-20190826125020-ad27480a10e7
	sigs.k8s.io/cluster-api-provider-docker v0.1.1-0.20190827143441-33201d009aa8
)
