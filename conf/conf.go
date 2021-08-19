package conf

import (
	"os"
	"soda/util"

	"github.com/joho/godotenv"
)

func Init() {
	// 从本地读取环境变量
	err := godotenv.Load()
	if err != nil {
		util.Log().Panic("", err)
		os.Exit(0)
	}

	util.BuildLogger(os.Getenv("LOG_LEVEL"))
}
