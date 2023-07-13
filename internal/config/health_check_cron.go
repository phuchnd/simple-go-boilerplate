package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	CronHealthCheckConfigName = "health_check"
)

func (c *configImpl) GetCronHealthCheckConfig() *CronConfig {
	cfgPath := fmt.Sprintf("%s.%s", JobsConfigName, CronHealthCheckConfigName)
	return &CronConfig{
		CronSpec: c.viper.GetString(cfgPath + ".cron_spec"),
		Enable:   c.viper.GetBool(cfgPath + ".enable"),
	}
}

func initHealthCheckConfig(v *viper.Viper) {
	v.SetDefault(fmt.Sprintf("%s.%s", JobsConfigName, CronHealthCheckConfigName), map[string]interface{}{
		"cron_spec": "@every 1m",
		"enable":    false,
	})
}
