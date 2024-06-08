package job

import (
	"context"
	"fmt"
	"github.com/xxl-job/xxl-job-executor-go"
	"time"
)

type HelloJob struct {
}

func NewHelloJob() *HelloJob {
	return &HelloJob{}
}

func (h *HelloJob) SayHello(ctx context.Context, param *xxl.RunReq) string {
	fmt.Println("hello world", time.Now().Format(time.DateTime))
	return "hello success"
}
