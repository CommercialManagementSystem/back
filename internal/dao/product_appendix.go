package dao

import (
	"context"
	"errors"
	"github.com/CommercialManagementSystem/back/internal/entity"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProductAppendixDaoSet = wire.NewSet(wire.Struct(new(ProductAppendixDao), "*"))

type ProductAppendixDao struct {
	DB *gorm.DB
}

func (p *ProductAppendixDao) Create(ctx context.Context, params *[]entity.ProductAppendix) error {
	db := p.DB.Create(params)
	err := db.Error
	if err != nil {
		return err
	}

	return db.Save(params).Error
}

func (p *ProductAppendixDao) Delete(ctx context.Context, params *[]int) error {
	return p.DB.Where("id in ?", *params).Delete(&entity.ProductAppendix{}).Error
}

func (p *ProductAppendixDao) QueryByProduct(ctx context.Context, params *entity.ProductAppendix) (*[]entity.ProductAppendix, error) {
	res := new([]entity.ProductAppendix)
	sql := p.DB.Model(&entity.ProductAppendix{})
	if params.ProductID <= 0 {
		return nil, errors.New("ID must be positive integer")
	}
	err := sql.Where("id = ?", params.ID).Find(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
