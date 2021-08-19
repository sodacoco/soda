package main

import (
	"os"
	"soda/conf"
	"soda/server"
	"soda/util"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	err := r.Run()
	if err != nil {
		util.Log().Panic("", err)
		os.Exit(0)
	}
}
