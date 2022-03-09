package dao

import (
	"context"
	"errors"
	"github.com/CommercialManagementSystem/back/internal/entity"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProductUserDaoSet = wire.NewSet(wire.Struct(new(ProductUserDao), "*"))

type ProductUserDao struct {
	DB *gorm.DB
}

func (p *ProductUserDao) Create(ctx context.Context, params *[]entity.ProductUser) error {
	db := p.DB.Create(params)
	err := db.Error
	if err != nil {
		return err
	}

	return db.Save(params).Error

}

func (p *ProductUserDao) Delete(ctx context.Context, params *entity.ProductUser) error {
	return p.DB.Where(map[string]interface{}{
		"product_id": params.ProductID,
		"user_id":    params.UserID,
	}).Delete(&entity.ProductUser{}).Error
}

func (p *ProductUserDao) QueryProductByUser(ctx context.Context, params *entity.ProductUser, options *QueryOption) (*[]entity.Product, int64, error) {
	res := new([]entity.ProductUser)
	sql := p.DB.Model(&entity.ProductUser{})
	option := getOption(*options)

	if params.UserID <= 0 {
		return nil, 0, errors.New("ID must be positive integer")
	}
	sql = sql.Where("user_id = ?", params.UserID).Joins("Product")

	txDB := sql.WithContext(ctx)
	var count int64
	err := txDB.Find(&[]entity.Product{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	sql = sql.Scopes(Paginate(option))
	err = sql.Find(res).Error
	if err != nil {
		return nil, 0, err
	}
	result := make([]entity.Product, 0)
	for _, item := range *res {
		result = append(result, item.Product)
	}
	return &result, count, err
}
