package controller

import (
	"github.com/CommercialManagementSystem/back/internal/model"
	"github.com/CommercialManagementSystem/back/internal/schema"
	"github.com/CommercialManagementSystem/back/pkg/warpper"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"strconv"
)

var ProductUserControllerSet = wire.NewSet(wire.Struct(new(ProductUserController), "*"))

type ProductUserController struct {
	ProductUserModel *model.ProductUserModel
}

func (p *ProductUserController) QueryProduct(c *gin.Context) {
	ctx := c.Request.Context()

	var option schema.QueryOption

	if err := warpper.ParseQuery(c, &option); err != nil {
		warpper.ResError(c, err)
		return
	}

	user := warpper.GetUserID(c)
	userID, err := strconv.Atoi(user)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := p.ProductUserModel.QueryBindProductUser(ctx, &schema.QueryProductUserRequestBody{
		QueryOption: option,
		UserID:      userID,
	})
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, res)
}

func (p *ProductUserController) AddProductUser(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.AddProductUserRequestBody
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	err := p.ProductUserModel.AddBindProductUser(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, nil)
}

func (p *ProductUserController) DeleteProduct(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.DeleteProductUserRequestBody
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	err := p.ProductUserModel.DeleteBindProductUser(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, nil)
}
