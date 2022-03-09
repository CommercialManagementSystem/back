package cryptox

import (
	"github.com/CommercialManagementSystem/back/internal/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Claims jwt 中的其他信息
type Claims struct {
	jwt.StandardClaims

	UID string `json:"uid"`
}

// GenerateToken 生成 jwt
func GenerateToken(uid string) (string, error) {
	cfg := config.C.JWT

	jwtSecret := []byte(cfg.Secret)
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(cfg.Expires * int(time.Hour)))

	claims := Claims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    cfg.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 格式化 JWT
func ParseToken(token string) (*Claims, error) {
	cfg := config.C.JWT

	jwtSecret := []byte(cfg.Secret)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, err
		}
	}

	return nil, err
}
