package router

import "github.com/gin-gonic/gin"

// RegisterAPI 路由列表
func (a *Router) registerAPI(app *gin.Engine) {
	g := app.Group("/api")

	v1 := g.Group("/v1")
	{
		v1.POST("/login", a.LoginAPI.Login)
	}
}
