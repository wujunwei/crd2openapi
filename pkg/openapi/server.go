package openapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wujunwei/crd2openapi/pkg/openapi/build"
	"io"
	extensionv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	swagggerbuilder "k8s.io/apiextensions-apiserver/pkg/controller/openapi/builder"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/kube-openapi/pkg/builder"
	"k8s.io/kube-openapi/pkg/validation/spec"
	"os"
	"strings"
)

var schema *runtime.Scheme

func init() {
	schema = runtime.NewScheme()
	extensionv1.AddToScheme(schema)
}

type Config struct {
	Out         *os.File
	CRDFiles    []*os.File
	Err         io.Writer
	Title       string
	Version     string
	Pretty      bool
	Indent      int
	Description string
}

func (c *Config) Complete() *Config {
	if c.Out == nil {
		c.Out = os.Stdout
	}
	c.Err = os.Stderr
	return c
}
func (c Config) New() (Converter, error) {
	return Converter{
		Out:         c.Out,
		Err:         c.Err,
		CRDReader:   c.CRDFiles,
		Title:       c.Title,
		Indent:      c.Indent,
		Version:     c.Version,
		Pretty:      c.Pretty,
		Description: c.Description,
	}, nil
}

//Converter crd to openapi json file
type Converter struct {
	Out         *os.File
	Err         io.Writer
	CRDReader   []*os.File
	Title       string
	Version     string
	Pretty      bool
	Indent      int
	Description string
}

func (c *Converter) analyzeCRD() []*extensionv1.CustomResourceDefinition {

	var crds []*extensionv1.CustomResourceDefinition
	for _, file := range c.CRDReader {
		decode := yaml.NewYAMLOrJSONDecoder(file, 4096)
		defer func() { _ = file.Close() }()
		for {
			raw := &runtime.RawExtension{}
			crd := &extensionv1.CustomResourceDefinition{}
			err := decode.Decode(raw)
			_ = json.Unmarshal(raw.Raw, crd)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					_, _ = fmt.Fprintf(c.Err, fmt.Sprintf("decode yaml file %s error : %s", file.Name(), err.Error()))
					continue
				}
			}

			crds = append(crds, crd)
		}
	}
	return crds
}

//Complete fill some fields
func (c *Converter) Complete(swagger *spec.Swagger) {
	if swagger.Info == nil {
		swagger.Info = &spec.Info{
			VendorExtensible: spec.VendorExtensible{Extensions: map[string]interface{}{"buildBy": "convert tool"}},
			InfoProps:        spec.InfoProps{Description: c.Description, Title: c.Title, Version: c.Version},
		}
	}
	swagger.Schemes = []string{"https"}

	swagger.Swagger = builder.OpenAPIVersion
}

func (c *Converter) Do() error {
	CRDs := c.analyzeCRD()
	if len(CRDs) == 0 {
		return errors.New("no available crd found")
	}
	var staticSpec = &spec.Swagger{}
	var allSpecs []*spec.Swagger
	for _, resourceDefinition := range CRDs {
		specs, err := build.NewSwaggerV2Converter().Convert(resourceDefinition)
		if err != nil {
			return err
		}
		allSpecs = append(allSpecs, specs...)
	}
	mergeSpecs, err := swagggerbuilder.MergeSpecs(staticSpec, allSpecs...)
	c.Complete(mergeSpecs)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(c.Out)
	defer c.Out.Close()
	if c.Pretty {
		enc.SetIndent("", strings.Repeat(" ", c.Indent))
	}

	err = enc.Encode(mergeSpecs)
	if err != nil {
		return err
	}
	return nil
}
