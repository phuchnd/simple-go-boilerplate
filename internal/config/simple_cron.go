package config

import "fmt"

const (
	JobsConfigName              = "jobs"
	CronSimpleExampleConfigName = "simple_example"
)

type CronConfig struct {
	CronSpec string
	Enable   bool
}

func GetCronSimpleExampleConfig() *CronConfig {
	v := cfgProvider.viper
	cfgPath := fmt.Sprintf("%s.%s", JobsConfigName, CronSimpleExampleConfigName)
	return &CronConfig{
		CronSpec: v.GetString(cfgPath + ".cron_spec"),
		Enable:   v.GetBool(cfgPath + ".enable"),
	}
}

func initCronSimpleExampleConfig() {
	cfgProvider.viper.SetDefault(fmt.Sprintf("%s.%s", JobsConfigName, CronSimpleExampleConfigName), map[string]interface{}{
		"cron_spec": "@every 5s",
		"enable":    false,
	})
}
