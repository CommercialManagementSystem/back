package model

import (
	"context"
	"github.com/CommercialManagementSystem/back/internal/dao"
	"github.com/CommercialManagementSystem/back/internal/entity"
	"github.com/CommercialManagementSystem/back/internal/schema"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

var ProductUserModelSet = wire.NewSet(wire.Struct(new(ProductUserModel), "*"))

type ProductUserModel struct {
	ProductUserDao *dao.ProductUserDao
}

func (p *ProductUserModel) AddBindProductUser(ctx context.Context, params *schema.AddProductUserRequestBody) error {
	data := make([]entity.ProductUser, 0)
	for _, product := range params.PIDs {
		data = append(data, entity.ProductUser{
			ProductID: uint(product.PID),
			UserID:    uint(params.UID),
		})
	}
	err := p.ProductUserDao.Create(ctx, &data)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductUserModel) DeleteBindProductUser(ctx context.Context, params *schema.DeleteProductUserRequestBody) error {
	data := &entity.ProductUser{
		ProductID: uint(params.PID),
		UserID:    uint(params.UID),
	}
	err := p.ProductUserDao.Delete(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductUserModel) QueryBindProductUser(ctx context.Context, params *schema.QueryProductUserRequestBody) (*schema.QueryResponseBody, error) {
	data := &entity.ProductUser{UserID: uint(params.UserID)}
	options := &dao.QueryOption{
		Offset:   params.Offset,
		PageSize: params.PageSize,
		Order:    params.Order,
	}

	res, count, err := p.ProductUserDao.QueryProductByUser(ctx, data, options)
	if err != nil {
		return nil, err
	}

	option := dao.GetResOption(*options)

	items := new([]schema.QueryProductUserResponseItem)
	err = copier.Copy(items, res)
	if err != nil {
		return nil, err
	}

	return &schema.QueryResponseBody{
		QueryResponse: schema.QueryResponse{
			PageSize:    option.PageSize,
			CurrentPage: option.CurrentPage,
			Count:       int(count),
		},
		Data: items,
	}, nil
}
