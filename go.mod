module github.com/ashish-amarnath/capi-yaml-gen

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/google/go-cmp v0.3.1
	github.com/grpc-ecosystem/grpc-gateway v1.8.5 // indirect
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	golang.org/x/lint v0.0.0-20190301231843-5614ed5bae6f // indirect
	google.golang.org/genproto v0.0.0-20190307195333-5fe7a883aa19 // indirect
	google.golang.org/grpc v1.19.1 // indirect
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/cluster-bootstrap v0.0.0-20190711112844-b7409fb13d1b // indirect
	sigs.k8s.io/cluster-api v0.2.6-0.20191028180044-4e878af0247a
	sigs.k8s.io/cluster-api-provider-aws v0.4.3-0.20191026015723-d8deb7eed0a8
	sigs.k8s.io/cluster-api/test/infrastructure/docker v0.0.0-20191028180044-4e878af0247a
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
