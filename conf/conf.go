package conf

import (
	"github.com/joho/godotenv"
)

func Init() {
	// 从本地读取环境变量
	godotenv.Load()
}
