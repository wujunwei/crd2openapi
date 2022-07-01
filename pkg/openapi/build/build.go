package build

import (
	extensionv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/controller/openapi/builder"
)

func test() {
	crd := &extensionv1.CustomResourceDefinition{}
	v3, err := builder.BuildOpenAPIV3(crd, crd.APIVersion, builder.Options{V2: false})
}
