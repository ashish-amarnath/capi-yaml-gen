module github.com/ashish-amarnath/capi-yaml-gen

go 1.12

require (
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/google/go-cmp v0.3.1
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45 // indirect
	google.golang.org/appengine v1.5.0 // indirect
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/cluster-api v0.2.3-0.20190920220726-2b0b75ef2cc5
	sigs.k8s.io/cluster-api-bootstrap-provider-kubeadm v0.1.1-0.20190919144501-b7c1ea3d99f2
	sigs.k8s.io/cluster-api-provider-aws v0.4.1-0.20190916225024-9dc57fdea283
	sigs.k8s.io/cluster-api-provider-docker v0.2.1-0.20190918183731-8dc38744ec02
)
