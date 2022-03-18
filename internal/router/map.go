package router

import (
	"github.com/CommercialManagementSystem/back/internal/dao"
	"github.com/CommercialManagementSystem/back/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterAPI 路由列表
func (a *Router) registerAPI(app *gin.Engine) {
	g := app.Group("/api")

	v1 := g.Group("/v1")
	{
		v1.POST("/login", a.LoginAPI.Login)

		v1.POST("/user", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.UserAPI.AddUser)
		v1.PUT("/user", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.UserAPI.UpdateUser)
		v1.GET("/user", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.UserAPI.QueryUser)
		v1.DELETE("/user", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.UserAPI.DeleteUser)

		v1.POST("/product", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.ProductAPI.AddProduct)
		v1.PUT("/product", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.ProductAPI.UpdateProduct)
		v1.GET("/product", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.ProductAPI.QueryProduct)
		v1.DELETE("/product", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.ProductAPI.DeleteProduct)

		v1.POST("/product_user", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.ProductUserAPI.AddProductUser)
		v1.GET("/product_user", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.ProductUserAPI.QueryProduct)
		v1.DELETE("/product_user", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.ProductUserAPI.DeleteProduct)

		v1.POST("/product_appendix", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.ProductAppendixAPI.AddAppendix)
		v1.GET("/product_appendix/download", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.ProductAppendixAPI.DownloadProductAppendix)
		v1.DELETE("/product_appendix", middleware.AuthMiddleware(a.UserDao, dao.AdminAuth|dao.CompanyAuth|dao.LeaderAuth|dao.NormalAuth), a.ProductAppendixAPI.DeleteAppendix)
	}
}
