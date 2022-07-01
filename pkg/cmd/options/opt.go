package options

import (
	"fmt"
	"github.com/spf13/pflag"
	"io"
	"k8s.io/apiextensions-apiserver/pkg/apiserver"
	"net"
)

type CRDConvertOptions struct {
	StdOut io.Writer
	StdErr io.Writer
}

// AddFlags adds the apiextensions-apiserver flags to the flagset.
func (o CRDConvertOptions) AddFlags(fs *pflag.FlagSet) {
	fs.String()
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
func (o CRDConvertOptions) Config() (*apiserver.Config, error) {
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{netutils.ParseIPSloppy("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	return config, nil
}

// NewCRDConvertOptions creates default options of this command.
func NewCRDConvertOptions(out, errOut io.Writer) *CRDConvertOptions {
	o := &CRDConvertOptions{

		StdOut: out,
		StdErr: errOut,
	}

	return o
}
