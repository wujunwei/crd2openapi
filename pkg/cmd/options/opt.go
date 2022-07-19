package options

import (
	"github.com/spf13/pflag"
	"github.com/wujunwei/crd2openapi/pkg/openapi"
	"github.com/wujunwei/crd2openapi/pkg/util"
	"os"
	"path/filepath"
)

type CRDConvertOptions struct {
	file      string
	output    string
	title     string
	InputFIle []*os.File
	Out       *os.File
}

// AddFlags adds the apiextensions-apiserver flags to the flagset.
func (o CRDConvertOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&o.file, "file", "f", "-", "filename or path to the CRD to be converted.")
	fs.StringVarP(&o.output, "output", "o", "-", "out openapi json file.")
	fs.StringVarP(&o.title, "title", "t", "kubernetes crd", "the tile of the swagger json.")
}

// Validate validates the apiextensions-apiserver options.
func (o CRDConvertOptions) Validate() error {
	return nil
}

// Complete fills in missing options.
func (o *CRDConvertOptions) Complete() error {
	if o.file != "-" {
		recursive := false
		_, name := filepath.Split(o.file)
		if name == "..." {
			recursive = true
		}
		file, err := util.ReadAllFile(o.file, recursive)
		if err != nil {
			return err
		}
		o.InputFIle = file
	} else {
		o.InputFIle = append(o.InputFIle, os.Stdin)
	}

	if o.output != "-" {
		open, err := os.Open(o.output)
		if err != nil {
			return err
		}
		o.Out = open
	} else {
		o.Out = os.Stdout
	}
	return nil
}

// Config returns an convert configuration.
func (o CRDConvertOptions) Config() (*openapi.Config, error) {
	conf := &openapi.Config{CRDFiles: o.InputFIle, Out: o.Out}
	return conf, nil
}

// NewCRDConvertOptions creates default options of this command.
func NewCRDConvertOptions() *CRDConvertOptions {
	o := &CRDConvertOptions{}

	return o
}
