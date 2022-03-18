package router

import (
	"github.com/CommercialManagementSystem/back/internal/controller"
	"github.com/CommercialManagementSystem/back/internal/dao"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// RouterSet 路由注入
var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"))

// Router 路由管理器
type Router struct {
	UserDao            *dao.UserDao
	LoginAPI           *controller.Login
	UserAPI            *controller.UserController
	ProductAPI         *controller.ProductController
	ProductUserAPI     *controller.ProductUserController
	ProductAppendixAPI *controller.AppendixController
}

// Register 注册路由
func (a *Router) Register(app *gin.Engine) {
	a.registerAPI(app)
}
