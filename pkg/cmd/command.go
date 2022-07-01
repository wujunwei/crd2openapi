package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wujunwei/crd2openapi/pkg/cmd/options"
	"io"
)

func NewCommand(out, errOut io.Writer, stopCh <-chan struct{}) *cobra.Command {
	o := options.NewCRDConvertOptions(out, errOut)

	cmd := &cobra.Command{
		Short: "Launch an API extensions API server",
		Long:  "Launch an API extensions API server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := Run(o, stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	fs := cmd.Flags()
	o.AddFlags(fs)
	return cmd
}

func Run(o *options.CRDConvertOptions, stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New(genericapiserver.NewEmptyDelegate())
	if err != nil {
		return err
	}
	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}
