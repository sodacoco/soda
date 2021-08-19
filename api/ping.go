package api

import (
	"net/http"
	"soda/serializer"

	"github.com/gin-gonic/gin"
)

// 状态返回
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, serializer.Response{
		Code:    http.StatusOK,
		Message: "Pong",
	})
}
