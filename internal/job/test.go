package job111

import (
	"context"
	"fmt"
	"kratos-cron/internal/biz"
	"time"
)

func Test(s *ExampleJob) {
	s.uc.CreateGreeter(context.Background(), &biz.Greeter{})
	fmt.Printf("当前时间 %v \n", time.Now().Unix())
}
