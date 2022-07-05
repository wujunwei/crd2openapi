package options

import (
	"github.com/spf13/pflag"
	"github.com/wujunwei/crd2openapi/pkg/openapi"
	"io"
)

type CRDConvertOptions struct {
	StdOut io.Writer
	StdErr io.Writer
	StdIn  io.Writer
	file   string
}

// AddFlags adds the apiextensions-apiserver flags to the flagset.
func (o CRDConvertOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.file, "file", "", "filename or path to the CRD to be converted.")
}

// Validate validates the apiextensions-apiserver options.
func (o CRDConvertOptions) Validate() error {

	return nil
}

// Complete fills in missing options.
func (o *CRDConvertOptions) Complete() error {
	return nil
}

// Config returns an apiextensions-apiserver configuration.
func (o CRDConvertOptions) Config() (*openapi.Config, error) {

	return nil, nil
}

// NewCRDConvertOptions creates default options of this command.
func NewCRDConvertOptions() *CRDConvertOptions {
	o := &CRDConvertOptions{}

	return o
}
