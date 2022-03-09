package controller

import (
	"github.com/CommercialManagementSystem/back/internal/model"
	"github.com/CommercialManagementSystem/back/internal/schema"
	"github.com/CommercialManagementSystem/back/pkg/warpper"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProductControllerSet = wire.NewSet(wire.Struct(new(ProductController), "*"))

type ProductController struct {
	ProductModel *model.ProductModel
}

func (p *ProductController) QueryProduct(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.QueryProductRequestBody
	if err := warpper.ParseQuery(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := p.ProductModel.QueryProduct(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, res)
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.UpdateProductRequestBody
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	err := p.ProductModel.UpdateProduct(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, nil)
}

func (p *ProductController) AddProduct(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.CreateProductRequestBody
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	err := p.ProductModel.CreateProduct(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, nil)
}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.DeleteProductRequestBody
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	err := p.ProductModel.DeleteProduct(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, nil)
}
