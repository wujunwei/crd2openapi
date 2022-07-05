package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wujunwei/crd2openapi/pkg/cmd/options"
)

func NewRootCommand() *cobra.Command {
	o := options.NewCRDConvertOptions()

	cmd := &cobra.Command{
		Short: "Convert crd to openapi",
		Long:  "Convert crd to openapi",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			if err := Run(o); err != nil {
				return err
			}
			return nil
		},
	}

	fs := cmd.Flags()
	o.AddFlags(fs)
	return cmd
}

func Run(o *options.CRDConvertOptions) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	cmd, err := config.Complete().New()
	if err != nil {
		return err
	}
	return cmd.Convert(nil, "")
}
