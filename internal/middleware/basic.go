package middleware

import (
	"github.com/CommercialManagementSystem/back/pkg/warpper"
	"github.com/gin-gonic/gin"
)

// NoMethodHandler 未找到请求方法的处理函数
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		warpper.ResError(c, warpper.ErrMethodNotAllow)
	}
}

// NoRouteHandler 未找到请求路由的处理函数
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		warpper.ResError(c, warpper.ErrNotFound)
	}
}
