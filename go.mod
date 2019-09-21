module github.com/ashish-amarnath/capi-yaml-gen

go 1.12

require (
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/google/go-cmp v0.3.0
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	golang.org/x/sys v0.0.0-20190710143415-6ec70d6a5542 // indirect
	k8s.io/api v0.0.0-20190711103429-37c3b8b1ca65
	k8s.io/apimachinery v0.0.0-20190711103026-7bf792636534
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/cluster-api v0.0.0-20190830205410-af7f7d74e93b
	sigs.k8s.io/cluster-api-bootstrap-provider-kubeadm v0.0.0-20190830202805-e3b2bcf70c84
	sigs.k8s.io/cluster-api-provider-aws v0.3.4-0.20190830203004-f748de96b794
	sigs.k8s.io/cluster-api-provider-docker v0.1.2-0.20190830184406-5931c16142d3
)
