package config

import "github.com/spf13/viper"

type SetDefaultConfigFunc func(v *viper.Viper)
type GetConfigFunc func(v *viper.Viper) interface{}

// ConfigOpt is a configuration on a config.
type ConfigOpt func(c *configImpl)

//go:generate mockery --name=Config --case=snake
type IConfig interface {
	SetDefault(v *viper.Viper)
	Get(v *viper.Viper) interface{}
}

// configImpl implements Config.
type configImpl struct {
	setDefaultFn SetDefaultConfigFunc
	getFn        GetConfigFunc
}

func NewConfig(getFn GetConfigFunc, opts ...ConfigOpt) IConfig {
	c := &configImpl{
		getFn: getFn,
	}

	for _, o := range opts {
		o(c)
	}

	return c
}

func (c *configImpl) SetDefault(v *viper.Viper) {
	if c.setDefaultFn != nil {
		c.setDefaultFn(v)
	}
}

func (c *configImpl) Get(v *viper.Viper) interface{} {
	return c.getFn(v)
}

func WithSetDefault(setDefaultFn SetDefaultConfigFunc) ConfigOpt {
	return func(c *configImpl) {
		c.setDefaultFn = setDefaultFn
	}
}
