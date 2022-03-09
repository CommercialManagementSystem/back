package model

import (
	"context"
	"github.com/CommercialManagementSystem/back/internal/dao"
	"github.com/CommercialManagementSystem/back/internal/entity"
	"github.com/CommercialManagementSystem/back/internal/schema"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

var ProductModelSet = wire.NewSet(wire.Struct(new(ProductModel), "*"))

type ProductModel struct {
	ProductDao *dao.ProductDao
}

func (p *ProductModel) QueryProduct(ctx context.Context, params *schema.QueryProductRequestBody) (*schema.QueryResponseBody, error) {
	options := &dao.QueryOption{
		Offset:   params.Offset,
		PageSize: params.PageSize,
		Order:    params.Order,
	}
	res, count, err := p.ProductDao.Query(ctx, &entity.Product{Name: params.Name}, options)
	if err != nil {
		return nil, err
	}
	option := dao.GetResOption(*options)
	return &schema.QueryResponseBody{
		QueryResponse: schema.QueryResponse{
			PageSize:    option.PageSize,
			CurrentPage: option.CurrentPage,
			Count:       int(count),
		},
		Data: res,
	}, nil
}

func (p *ProductModel) UpdateProduct(ctx context.Context, params *schema.UpdateProductRequestBody) error {
	data := new(entity.Product)
	err := copier.Copy(data, params)
	if err != nil {
		return err
	}
	err = p.ProductDao.Update(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductModel) CreateProduct(ctx context.Context, params *schema.CreateProductRequestBody) error {
	data := new(entity.Product)
	err := copier.Copy(data, params)
	if err != nil {
		return err
	}
	err = p.ProductDao.Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductModel) DeleteProduct(ctx context.Context, params *schema.DeleteProductRequestBody) error {
	return p.ProductDao.Delete(ctx, &params.IDs)
}
