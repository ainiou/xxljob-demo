package xxljob

import (
	"fmt"
	"github.com/xxl-job/xxl-job-executor-go"
	"xxljob-demo/config"
	"xxljob-demo/job"
)

type XxlOption struct {
	HelloJob *job.HelloJob
}

type XxlHandler struct {
	*XxlOption

	XxlJob xxl.Executor
}

func NewXxlJob(cfg config.XXLJobConfig) *XxlHandler {
	option := &XxlOption{
		HelloJob: job.NewHelloJob(),
	}
	executor := initXXLJob(cfg)
	h := &XxlHandler{
		XxlJob:    executor,
		XxlOption: option,
	}
	// 注册任务
	executor.RegTask("hello", h.HelloJob.SayHello)
	return h
}

func initXXLJob(config config.XXLJobConfig) xxl.Executor {
	var opts []xxl.Option

	if config.ServerAddr != "" {
		opts = append(opts, xxl.ServerAddr(config.ServerAddr))
	}
	if config.AccessToken != "" {
		opts = append(opts, xxl.AccessToken(config.AccessToken))
	}

	if config.AppName != "" {
		opts = append(opts, xxl.RegistryKey(config.AppName))
	}

	if config.IP != "" {
		opts = append(opts, xxl.ExecutorIp(config.IP))
	}

	if config.Port > 0 {
		opts = append(opts, xxl.ExecutorPort(fmt.Sprintf("%d", config.Port)))
	}

	//opts = append(opts,xxl.SetLogger())

	exec := xxl.NewExecutor(opts...)
	exec.Init()
	return exec

}
