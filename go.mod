module github.com/ashish-amarnath/capi-yaml-gen

go 1.12

require (
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/google/go-cmp v0.3.1
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/cluster-api v0.2.6
	sigs.k8s.io/cluster-api-bootstrap-provider-kubeadm v0.1.5
	sigs.k8s.io/cluster-api-provider-aws v0.4.3
	sigs.k8s.io/cluster-api-provider-docker v0.2.0
)
