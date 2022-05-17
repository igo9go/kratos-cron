package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
	"kratos-cron/internal/conf"
	"kratos-cron/internal/service"
)

type CronWorker struct {
	c    *conf.Job
	sche *cron.Cron
}

func NewCronWorker(c *conf.Job, jobService *service.JobService) (s *CronWorker) {
	jobService.Init()
	s = &CronWorker{
		c:    c,
		sche: cron.New(),
	}
	for _, j := range c.Jobs {
		job, ok := service.DefaultJobs[j.Name]
		if !ok {
			log.Warnf("can not find job: %s", j.Name)

			continue
		}
		s.sche.AddFunc(j.Schedule, job)
	}
	log.Info("加载job数量:", len(c.Jobs))
	return s
}

func (s *CronWorker) Start(c context.Context) error {
	s.sche.Start()
	return nil
}

func (s *CronWorker) Stop(c context.Context) error {
	s.sche.Stop()
	return nil
}

func (s *CronWorker) RunSrv(name string) {
	log.Info("run job{%s}", name)
	//switch name {
	//case s.c.ExampleJob.JonName:
	//	s.job.DoMyWork()
	//default:
	//	s.HeartBeat()
	//}
}

// HeartBeat .
func (s *CronWorker) HeartBeat() {
	log.Info("alive...")
}
