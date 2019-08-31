package serialize

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/versioning"
	"k8s.io/client-go/kubernetes/scheme"
	kubeadmv1 "sigs.k8s.io/cluster-api-bootstrap-provider-kubeadm/api/v1alpha2"
	awsv1 "sigs.k8s.io/cluster-api-provider-aws/api/v1alpha2"
	dockerv1 "sigs.k8s.io/cluster-api-provider-docker/api/v1alpha2"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha2"
)

func Scheme() *runtime.Scheme {
	myScheme := runtime.NewScheme()
	if err := clusterv1.AddToScheme(myScheme); err != nil {
		panic(err)
	}
	if err := dockerv1.AddToScheme(myScheme); err != nil {
		panic(err)
	}
	if err := kubeadmv1.AddToScheme(myScheme); err != nil {
		panic(err)
	}
	if err := awsv1.AddToScheme(myScheme); err != nil {
		panic(err)
	}

	return myScheme
}

func MarshalToYAML(obj runtime.Object) ([]byte, error) {
	mediaType := "application/yaml"
	info, ok := runtime.SerializerInfoForMediaType(scheme.Codecs.SupportedMediaTypes(), mediaType)
	if !ok {
		return []byte{}, fmt.Errorf("unsupported media type %q", mediaType)
	}
	codec := versioning.NewDefaultingCodecForScheme(Scheme(), info.Serializer, nil, nil, nil)
	var buf bytes.Buffer
	if err := codec.Encode(obj, &buf); err != nil {
		return nil, errors.WithStack(err)
	}
	return bytes.TrimSpace(buf.Bytes()), nil
}
