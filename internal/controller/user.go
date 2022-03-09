package controller

import (
	"github.com/CommercialManagementSystem/back/internal/model"
	"github.com/CommercialManagementSystem/back/internal/schema"
	"github.com/CommercialManagementSystem/back/pkg/warpper"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var UserControllerSet = wire.NewSet(wire.Struct(new(UserController), "*"))

type UserController struct {
	UserModel *model.UserModel
}

func (u *UserController) QueryUser(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.QueryUserRequestBody
	if err := warpper.ParseQuery(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := u.UserModel.QueryUser(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, res)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.UpdateUserRequestBody
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	err := u.UserModel.UpdateUser(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, nil)
}

func (u *UserController) AddUser(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.UpdateUserRequestBody
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	err := u.UserModel.UpdateUser(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, nil)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.DeleteUserRequestBody
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	err := u.UserModel.DeleteUser(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, nil)
}
