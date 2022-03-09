package dao

import (
	"context"
	"errors"
	"strings"

	"github.com/CommercialManagementSystem/back/internal/entity"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// UserDaoSet 注入 DI
var UserDaoSet = wire.NewSet(wire.Struct(new(UserDao), "*"))

// UserDao users 表相关的数据库操作
type UserDao struct {
	DB *gorm.DB
}

// UserQueryParams 查询用户的参数
type UserQueryParams struct {
	UID   string
	Phone string
}

// Get 根据给定条件查询用户
func (u *UserDao) Get(ctx context.Context, params UserQueryParams) (*[]entity.User, error) {
	result := new([]entity.User)
	db := u.DB.Model(&entity.User{})

	if v := params.UID; v != "" {
		db = db.Where("uid = ?", v)
	}

	if v := params.Phone; v != "" {
		db = db.Where("phone = ?", v)
	}

	err := db.Find(result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserDao) Update(ctx context.Context, params *entity.User) error {
	columns := make(map[string]interface{})
	sql := u.DB.Model(&entity.User{})
	if v := params.ID; v <= 0 {
		return errors.New("ID must be positive integer")
	} else {
		sql = sql.Where("id = ?", v)

		if params.UID != "" {
			columns["uid"] = strings.TrimSpace(params.UID)
		}
		if params.Password != "" {
			columns["password"] = strings.TrimSpace(params.Password)
		}
		if params.Phone != "" {
			columns["phone"] = strings.TrimSpace(params.Phone)
		}
		if params.Email != "" {
			columns["email"] = strings.TrimSpace(params.Email)
		}

		columns["sex"] = params.Sex

		if params.RoleID != 0 {
			columns["uid"] = params.RoleID
		}
		if params.Authority != 0 {
			columns["authority"] = params.Authority
		}
		return sql.Updates(columns).Error
	}
}

func (u *UserDao) Delete(ctx context.Context, params *[]int) error {
	return u.DB.Where("id in ?", *params).Delete(&entity.User{}).Error
}

func (u *UserDao) Create(ctx context.Context, params *entity.User) error {
	db := u.DB.Model(&entity.User{})
	db = db.Create(params)
	err := db.Error
	if err != nil {
		return err
	}
	return db.Save(params).Error
}

func (u *UserDao) Query(ctx context.Context, params *entity.User, options *QueryOption) (*[]entity.User, int64, error) {
	result := new([]entity.User)
	db := u.DB.Model(&entity.User{})

	option := getOption(*options)

	if v := params.Name; v != "" {
		db = db.Or(GetLikeSQL("name", params.Name))
	}

	if v := params.Phone; v != "" {
		db = db.Or(GetLikeSQL("phone", params.Phone))
	}

	txDB := db.WithContext(ctx)
	var count int64
	err := txDB.Find(&[]entity.User{}).Count(&count).Error
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
