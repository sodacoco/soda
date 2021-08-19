package main

import (
	"os"
	"soda/conf"

	"github.com/gin-gonic/gin"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		os.Exit(0)
	}
}
