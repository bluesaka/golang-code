/**
github.com/robfig/cron
*/
package cron

import (
	"github.com/robfig/cron"
	"log"
)

func Cron() {
	// 标准cron格式
	// 秒 分 时 日 月 周
	spec := "* * * * * ?"
	c := cron.New()
	c.AddJob(spec, Job1{})
	c.Run()
	defer c.Stop()
}

type Job1 struct{}

func (j Job1) Run() {
	log.Println("job1")
}