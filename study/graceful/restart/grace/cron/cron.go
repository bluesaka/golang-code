package cron

import (
	"github.com/robfig/cron"
	"go-code/study/log"
)

func Cron() {
	// 标准cron格式
	// 秒 分 时 日 月 周
	spec := "* * * * * ?"
	c := cron.New()
	c.AddJob(spec, Job1{})
	c.Start()
}

type Job1 struct{}

func (j Job1) Run() {
	log.ZapLogger.Info("job1")
}
