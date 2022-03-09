package controller

import (
	"fmt"
	"github.com/CommercialManagementSystem/back/internal/model"
	"github.com/CommercialManagementSystem/back/internal/schema"
	"github.com/CommercialManagementSystem/back/pkg/warpper"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"time"
)

var AppendixControllerSet = wire.NewSet(wire.Struct(new(AppendixController), "*"))

type AppendixController struct {
	AppendixModel *model.ProductAppendixModel
}

func (a *AppendixController) DownloadProductAppendix(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.DownloadProductAppendixRequestBody
	if err := warpper.ParseQuery(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	res, err := a.AppendixModel.DownloadProductAppendix(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s_product_appendix_export.zip",
		time.Now().Format("2006-01-02 15:04")))
	c.Header("Content-Transfer-Encoding", "binary")

	_, err = res.WriteTo(c.Writer)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, nil)
}

func (a *AppendixController) AddAppendix(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.AddProductAppendixRequestBody
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	err := a.AppendixModel.AddProductAppendix(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, nil)
}

func (a *AppendixController) DeleteAppendix(c *gin.Context) {
	ctx := c.Request.Context()

	var data schema.DeleteProductAppendixRequestBody
	if err := warpper.ParseJSON(c, &data); err != nil {
		warpper.ResError(c, err)
		return
	}

	err := a.AppendixModel.DeleteProductAppendix(ctx, &data)
	if err != nil {
		warpper.ResError(c, err)
		return
	}

	warpper.ResSuccess(c, nil)
}
