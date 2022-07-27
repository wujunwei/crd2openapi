package options

import (
	"errors"
	"github.com/spf13/pflag"
	"github.com/wujunwei/crd2openapi/pkg/openapi"
	"github.com/wujunwei/crd2openapi/pkg/util"
	"os"
	"path/filepath"
)

type CRDConvertOptions struct {
	file        string
	output      string
	title       string
	version     string
	pretty      bool
	indent      int
	InputFIle   []*os.File
	Out         *os.File
	description string
}

// AddFlags adds the apiextensions-apiserver flags to the flagset.
func (o *CRDConvertOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&o.file, "file", "f", "-", "filename or path to the CRD to be converted.")
	fs.StringVarP(&o.output, "output", "o", "-", "out openapi json file.")
	fs.BoolVarP(&o.pretty, "pretty", "p", false, "print the json pretty.")
	fs.StringVarP(&o.title, "title", "t", "kubernetes crd", "the tile of the swagger json.")
	fs.IntVarP(&o.indent, "indent", "i", 4, "the indent of json line , only enable when pretty is true.")
	fs.StringVarP(&o.version, "version", "v", "1.0.0", "the version of the swagger json.")
	fs.StringVarP(&o.description, "description", "d", "kubernetes crd doc", "the description of the swagger json.")
}

// Validate validates the apiextensions-apiserver options.
func (o CRDConvertOptions) Validate() error {
	if o.indent < 0 {
		return errors.New("indent can not less than zero")
	}
	if o.version == "" {
		return errors.New("invalid version")
	}
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
		open, err := os.OpenFile(o.output, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
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
	conf := &openapi.Config{
		CRDFiles:    o.InputFIle,
		Out:         o.Out,
		Title:       o.title,
		Indent:      o.indent,
		Pretty:      o.pretty,
		Version:     o.version,
		Description: o.description,
	}
	return conf, nil
}

// NewCRDConvertOptions creates default options of this command.
func NewCRDConvertOptions() *CRDConvertOptions {
	o := &CRDConvertOptions{}
	return o
}
