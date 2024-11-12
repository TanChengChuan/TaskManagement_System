package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

var JwtSecret = []byte("#HswACG10100100010") // 替换为你的密钥
var jwtExpirationDelta = time.Hour * 720     // token 有效期，这里设置为 30 天

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 创建 JWT token 返回给用户登录后的。
func GenerateToken(username string) (string, error) {
	//这里可以是数据库调用  看用户名是否对的上？

	c := MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpirationDelta)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	//表明期限

	signedToken, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err //错误返回空令牌
	}

	viper.Set("token", signedToken)
	return signedToken, nil
}
