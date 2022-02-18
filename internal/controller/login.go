package controller

import (
	"github.com/CommercialManagementSystem/back/internal/model"
	"github.com/CommercialManagementSystem/back/internal/schema"
	"github.com/CommercialManagementSystem/back/pkg/logger"
	"github.com/CommercialManagementSystem/back/pkg/warpper"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// LoginSet Login DI
var LoginSet = wire.NewSet(wire.Struct(new(Login), "*"))

// Login 登录结构体
type Login struct {
	LoginModel *model.LoginModel
}

// Login 登录方法
func (l *Login) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.LoginReqBodySchema
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := l.LoginModel.Login(ctx, data.Username, data.Password)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	ctx = logger.NewUserIDContext(ctx, res.UID)
	ctx = logger.NewTagContext(ctx, "__login__")
	logger.WithContext(ctx).Info("登入系统")

	warpper.ResSuccess(c, res)
}
