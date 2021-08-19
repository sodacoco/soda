package server

import (
	"soda/api"

	"github.com/gin-gonic/gin"
)

// 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)
	}

	return r
}
