package model

import (
	"context"
	"github.com/CommercialManagementSystem/back/internal/dao"
	"github.com/CommercialManagementSystem/back/internal/entity"
	"github.com/CommercialManagementSystem/back/internal/schema"
	"github.com/CommercialManagementSystem/back/pkg/logger"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserModelSet = wire.NewSet(wire.Struct(new(UserModel), "*"))

type UserModel struct {
	UserDao *dao.UserDao
}

func (u *UserModel) QueryUser(ctx context.Context, params *schema.QueryUserRequestBody) (*schema.QueryResponseBody, error) {
	options := &dao.QueryOption{
		Offset:   params.Offset,
		PageSize: params.PageSize,
		Order:    params.Order,
	}

	res, count, err := u.UserDao.Query(
		ctx,
		&entity.User{
			Name:  params.Name,
			Phone: params.Phone,
		},
		options,
	)
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

func (u *UserModel) UpdateUser(ctx context.Context, params *schema.UpdateUserRequestBody) error {
	data := &entity.User{
		Model:     gorm.Model{ID: uint(params.ID)},
		UID:       params.UID,
		Password:  params.Password,
		Authority: params.Authority,
		Name:      params.Name,
		Sex:       params.Sex,
		Phone:     params.Phone,
		Email:     params.Email,
		RoleID:    uint(params.Role),
	}
	err := u.UserDao.Update(ctx, data)
	if err != nil {
		logger.WithContext(ctx).Errorf("user update error:%v", err)
		return err
	}
	return nil
}

func (u *UserModel) CreateUser(ctx context.Context, params *schema.CreateUserRequestBody) error {
	data := &entity.User{
		UID:       params.UID,
		Password:  params.Password,
		Authority: params.Authority,
		Name:      params.Name,
		Sex:       params.Sex,
		Phone:     params.Phone,
		Email:     params.Email,
		RoleID:    uint(params.Role),
	}
	err := u.UserDao.Create(ctx, data)
	if err != nil {
		logger.WithContext(ctx).Errorf("user create error:%v", err)
		return err
	}
	return nil
}

func (u *UserModel) DeleteUser(ctx context.Context, params *schema.DeleteUserRequestBody) error {
	err := u.UserDao.Delete(ctx, &params.IDs)
	if err != nil {
		logger.WithContext(ctx).Errorf("user delete error:%v", err)
		return err
	}
	return nil
}
