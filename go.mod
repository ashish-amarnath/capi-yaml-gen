module github.com/ashish-amarnath/capi-yaml-gen

go 1.12

require (
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/google/go-cmp v0.5.9
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/cluster-api v0.2.6-0.20191031173438-6f77ca5bd696
	sigs.k8s.io/cluster-api-provider-aws v0.4.3-0.20191024155312-af5c36445248
	sigs.k8s.io/cluster-api/test/infrastructure/docker v0.0.0-20191028171043-098155a9501c
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
