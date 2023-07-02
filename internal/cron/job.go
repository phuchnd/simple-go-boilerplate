package cron

//go:generate mockery --name=IJob --case=snake --disable-version-string
type IJob interface {
	CronSpec() string
	IsEnable() bool
	Run()
}
