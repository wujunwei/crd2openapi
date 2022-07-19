package build

import (
	extensionv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/controller/openapi/builder"
	"k8s.io/kube-openapi/pkg/validation/spec"
)

type Converter interface {
	Convert(crd *extensionv1.CustomResourceDefinition) ([]*spec.Swagger, error)
}

type SwaggerV2 struct {
}

func (s SwaggerV2) Convert(crd *extensionv1.CustomResourceDefinition) ([]*spec.Swagger, error) {
	var crdSpecs []*spec.Swagger
	for _, v := range crd.Spec.Versions {
		if !v.Served {
			continue
		}
		// Defaults are not pruned here, but before being served.
		sw, err := builder.BuildOpenAPIV2(crd, v.Name, builder.Options{V2: true, SkipFilterSchemaForKubectlOpenAPIV2Validation: true, StripValueValidation: true, StripNullable: true, AllowNonStructural: false})
		if err != nil {
			return nil, err
		}
		crdSpecs = append(crdSpecs, sw)
	}
	//builder.MergeSpecs(c.staticSpec, crdSpecs...)
	return crdSpecs, nil
}

func NewSwaggerV2Converter() Converter {
	return SwaggerV2{}
}
