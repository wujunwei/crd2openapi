package openapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wujunwei/crd2openapi/pkg/openapi/build"
	"gopkg.in/yaml.v3"
	"io"
	extensionv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	swagggerbuilder "k8s.io/apiextensions-apiserver/pkg/controller/openapi/builder"
	"k8s.io/kube-openapi/pkg/validation/spec"
	"os"
)

type Config struct {
	Out      *os.File
	CRDFiles []*os.File
	Err      io.Writer
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
		Out:       c.Out,
		Err:       c.Err,
		CRDReader: c.CRDFiles,
	}, nil
}

//Converter crd to openapi json file
type Converter struct {
	Out       *os.File
	Err       io.Writer
	CRDReader []*os.File
}

func (c *Converter) analyzeCRD() []*extensionv1.CustomResourceDefinition {
	var crds []*extensionv1.CustomResourceDefinition
	for _, file := range c.CRDReader {
		decode := yaml.NewDecoder(file)
		defer func() { _ = file.Close() }()
		decode.KnownFields(false)
		crd := &extensionv1.CustomResourceDefinition{}
		for {
			err := decode.Decode(crds)
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
func (c *Converter) Do() error {
	CRDs := c.analyzeCRD()
	if len(CRDs) == 0 {
		return errors.New("no available crd found")
	}
	var stisticSpec = &spec.Swagger{}
	var allSpecs []*spec.Swagger
	for _, resourceDefinition := range CRDs {
		specs, err := build.NewSwaggerV2Converter().Convert(resourceDefinition)
		if err != nil {
			return err
		}
		allSpecs = append(allSpecs, specs...)
	}
	mergeSpecs, err := swagggerbuilder.MergeSpecs(stisticSpec, allSpecs...)
	if err != nil {
		return err
	}
	// todo 支持格式化输出json
	enc := json.NewEncoder(c.Out)
	defer c.Out.Close()
	err = enc.Encode(mergeSpecs)
	if err != nil {
		return err
	}
	return nil
}
