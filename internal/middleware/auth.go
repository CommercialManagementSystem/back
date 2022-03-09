package middleware

import (
	"errors"
	"github.com/CommercialManagementSystem/back/internal/config"
	"github.com/CommercialManagementSystem/back/internal/dao"
	"github.com/CommercialManagementSystem/back/pkg/cryptox"
	"github.com/CommercialManagementSystem/back/pkg/logger"
	"github.com/CommercialManagementSystem/back/pkg/warpper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func wrapUserAuthContext(c *gin.Context, userID string, authority int) {
	warpper.SetUserID(c, userID)
	warpper.SetUserAuthority(c, authority)
	ctx := c.Request.Context()
	ctx = logger.NewUserIDContext(ctx, userID)
	c.Request = c.Request.WithContext(ctx)
}

func AuthMiddleware(userDao *dao.UserDao, authority int) gin.HandlerFunc {
	if !config.C.JWT.Enable {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		claims, err := cryptox.ParseToken(warpper.GetToken(c))
		if err != nil {
			warpper.ResError(c, warpper.ErrInvalidToken)
			return
		}

		res, err := userDao.Get(ctx, dao.UserQueryParams{
			UID: claims.UID,
		})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			warpper.ResError(c, warpper.ErrInvalidUser)
			return
		}

		user := (*res)[0]
		if !dao.CheckAuth(user.Authority, authority) {
			warpper.ResError(c, warpper.ErrNoPerm)
			return
		}

		wrapUserAuthContext(c, claims.UID, user.Authority)
		c.Next()
	}
}
