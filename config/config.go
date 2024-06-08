package config

import "fmt"

type Config struct {
	XXLJobConfig XXLJobConfig `toml:"xxljob"`
}

// XXLJobConfig 定时任务配置
type XXLJobConfig struct {
	ServerAddr  string `toml:"ServerAddr"`  // xxjob调度器地址
	AppName     string `toml:"AppName"`     // 执行器名称
	IP          string `toml:"IP"`          // 一般不需要配置 自动获取
	Port        int    `toml:"PORT"`        // 执行器端口 默认9999
	AccessToken string `toml:"AccessToken"` // 请求令牌(默认为空) 需要配置
}

func (c *XXLJobConfig) DefaultConfig() *XXLJobConfig {
	return &XXLJobConfig{
		ServerAddr:  "",
		AppName:     "",
		Port:        9999,
		AccessToken: "",
	}
}

func (c *XXLJobConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.IP, c.Port)
}
