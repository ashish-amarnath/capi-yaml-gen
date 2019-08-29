module github.com/ashish-amarnath/capi-yaml-gen

go 1.12

require (
	github.com/google/go-cmp v0.3.0
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	k8s.io/api v0.0.0-20190711103429-37c3b8b1ca65
	k8s.io/apimachinery v0.0.0-20190711103026-7bf792636534
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/cluster-api v0.0.0-20190826205838-47413a956cbf
	sigs.k8s.io/cluster-api-bootstrap-provider-kubeadm v0.0.0-20190826125020-ad27480a10e7
	sigs.k8s.io/cluster-api-provider-aws v0.3.4-0.20190827200310-09c3c213c198
	sigs.k8s.io/cluster-api-provider-docker v0.1.1-0.20190827143441-33201d009aa8
)
