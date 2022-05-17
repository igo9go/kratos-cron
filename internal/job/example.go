package job111

import (
	"context"
	"fmt"
	"kratos-cron/internal/biz"
	"time"
)

type ExampleJob struct {
	uc *biz.GreeterUsecase
}

func NewExampleJob(uc *biz.GreeterUsecase) *ExampleJob {
	job := &ExampleJob{
		uc: uc,
	}
	return job
}

func (s *ExampleJob) Init() {
	DefaultJobs = map[string]JobFunc{
		"one": s.DoMyWork,
		"two": s.DoOtherWork,
	}
}

func (s *ExampleJob) DoMyWork() {
	s.uc.CreateGreeter(context.Background(), &biz.Greeter{})
	fmt.Printf("当前时间 %v \n", time.Now().Unix())
}

func (s *ExampleJob) DoOtherWork() {
	fmt.Printf("当前时间2 %v \n", time.Now().Unix())
}
