package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"xxljob-demo/config"
	"xxljob-demo/xxljob"
)

var tomlFile string

func init() {
	flag.StringVar(&tomlFile, "tomlFile", "./doc/local.toml", "server TOML config file")
}

func main() {
	flag.Parse()
	cfg := config.Config{}
	_, err := toml.DecodeFile(tomlFile, &cfg)
	if err != nil {
		fmt.Printf("配置文件加载失败:%s\n", err.Error())
		return
	}

	handler := xxljob.NewXxlJob(cfg.XXLJobConfig)
	handler.XxlJob.Run()
}
