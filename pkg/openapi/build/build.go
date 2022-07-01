package build

import (
	extensionv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/controller/openapi/builder"
)

type Converter interface {
	Convert(crd *extensionv1.CustomResourceDefinition, version string) error
}

func test() {
	crd := &extensionv1.CustomResourceDefinition{}
	builder.BuildOpenAPIV3(crd, crd.APIVersion, builder.Options{V2: false})
}
