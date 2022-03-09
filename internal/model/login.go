package model

import (
	"context"
	"errors"
	"github.com/CommercialManagementSystem/back/internal/config"
	"github.com/CommercialManagementSystem/back/internal/dao"
	"github.com/CommercialManagementSystem/back/internal/schema"
	"github.com/CommercialManagementSystem/back/pkg/cryptox"
	"github.com/CommercialManagementSystem/back/pkg/warpper"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// LoginModelSet LoginModel 注入 DI
var LoginModelSet = wire.NewSet(wire.Struct(new(LoginModel), "*"))

// LoginModel 处理登录的主要逻辑
type LoginModel struct {
	UserDao *dao.UserDao
}

// Login 登录方法
func (l *LoginModel) Login(ctx context.Context, username, password string) (*schema.LoginResBodySchema, error) {
	res, err := l.UserDao.Get(ctx, dao.UserQueryParams{
		UID: username,
	})
	// 区分未找到错误
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, warpper.ErrInvalidUserName
	}
	if err != nil {
		return nil, err
	}

	user := (*res)[0]
	temp := cryptox.MD5(password)
	if temp != user.Password {
		return nil, warpper.ErrInvalidPassword
	}

	token := ""
	if config.C.JWT.Enable {
		token, err = cryptox.GenerateToken(user.UID)
		if err != nil {
			return nil, warpper.ErrCanNotGenerateToken
		}
	}

	return &schema.LoginResBodySchema{
		UID:       user.UID,
		Authority: user.Authority,
		Token:     token,
	}, nil
}
