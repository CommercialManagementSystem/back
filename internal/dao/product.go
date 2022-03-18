package dao

import (
	"context"
	"errors"
	"github.com/CommercialManagementSystem/back/internal/entity"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProductDaoSet = wire.NewSet(wire.Struct(new(ProductDao), "*"))

type ProductDao struct {
	DB *gorm.DB
}

func (p *ProductDao) Create(ctx context.Context, params *entity.Product) error {
	db := p.DB.Model(&entity.Product{})
	db = db.Create(params)
	err := db.Error
	if err != nil {
		return err
	}
	return db.Save(params).Error
}

func (p *ProductDao) Update(ctx context.Context, params *entity.Product) error {
	columns := make(map[string]interface{})
	sql := p.DB.Model(&entity.Product{})

	if params.ID <= 0 {
		return errors.New("ID must be positive integer")
	}
	sql = sql.Where("id = ?", params.ID)
	if params.Name != "" {
		columns["name"] = params.Name
	}

	if params.ExpenseRatio != "" {
		columns["expense_ratio"] = params.ExpenseRatio
	}

	if params.Company != 0 {
		columns["company"] = params.Company
	}

	if params.Plan != "" {
		columns["plan"] = params.Plan
	}

	if params.Rule != "" {
		columns["rule"] = params.Rule
	}

	if params.Steps != "" {
		columns["steps"] = params.Steps
	}
	return sql.Updates(columns).Error
}

func (p *ProductDao) Delete(ctx context.Context, params *[]int) error {
	return p.DB.Where("id in ?", *params).Delete(&entity.Product{}).Error
}

func (p *ProductDao) Query(ctx context.Context, params *entity.Product, options *QueryOption) (*[]entity.Product, int64, error) {
	result := new([]entity.Product)
	db := p.DB.Model(&entity.Product{})

	option := getOption(*options)

	if v := params.Name; v != "" {
		db = db.Or(GetLikeSQL("name", params.Name))
	}

	txDB := db.WithContext(ctx)
	var count int64
	err := txDB.Find(&[]entity.Product{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	db = db.Scopes(Paginate(option))
	err = db.Find(result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, count, nil
}
