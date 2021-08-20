package conf

import (
	"os"
	"soda/model"
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

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
}
