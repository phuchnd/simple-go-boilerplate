package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	JobsConfigName              = "jobs"
	CronSimpleExampleConfigName = "simple_example"
)

type CronConfig struct {
	CronSpec string
	Enable   bool
}

func (c *configImpl) GetCronSimpleExampleConfig() *CronConfig {
	cfgPath := fmt.Sprintf("%s.%s", JobsConfigName, CronSimpleExampleConfigName)
	return &CronConfig{
		CronSpec: c.viper.GetString(cfgPath + ".cron_spec"),
		Enable:   c.viper.GetBool(cfgPath + ".enable"),
	}
}

func initCronSimpleExampleConfig(v *viper.Viper) {
	v.SetDefault(fmt.Sprintf("%s.%s", JobsConfigName, CronSimpleExampleConfigName), map[string]interface{}{
		"cron_spec": "@every 5s",
		"enable":    false,
	})
}
