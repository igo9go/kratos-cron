package service

import (
	"context"
	"fmt"
	"kratos-cron/internal/biz"
	"time"
)

var DefaultJobs map[string]JobFunc

type JobFunc func()

type JobService struct {
	uc *biz.GreeterUsecase
}

func NewJobService(uc *biz.GreeterUsecase) *JobService {
	job := &JobService{
		uc: uc,
	}
	return job
}

func (s *JobService) Init() {
	DefaultJobs = map[string]JobFunc{
		"one": s.DoMyWork,
		"two": s.DoOtherWork,
	}
}

func (s *JobService) DoMyWork() {
	s.uc.CreateGreeter(context.Background(), &biz.Greeter{})
	fmt.Printf("当前时间 %v \n", time.Now().Unix())
}

func (s *JobService) DoOtherWork() {
	fmt.Printf("当前时间2 %v \n", time.Now().Unix())
}
