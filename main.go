package main

import "k8s.io/apiextensions-apiserver/pkg/controller/openapi/builder"

func main() {
	v3, err := builder.BuildOpenAPIV3(crd, versionName, builder.Options{V2: false})
}
