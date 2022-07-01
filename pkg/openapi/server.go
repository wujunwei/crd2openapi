package openapi

import "github.com/wujunwei/crd2openapi/pkg/openapi/build"

type Config struct {
}

func (c *Config) Complete() *Config {
	return c
}
func (c Config) New() (build.Converter, error) {
	return nil, nil
}
