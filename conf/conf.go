package conf

import (
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	// 从本地读取环境变量
	err := godotenv.Load()
	if err != nil {
		os.Exit(0)
	}
}
