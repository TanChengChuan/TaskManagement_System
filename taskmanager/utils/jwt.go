package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var JwtSecret = []byte("#HswACG10100100010") // 替换为你的密钥
var jwtExpirationDelta = time.Hour * 720     // token 有效期，这里设置为 30 天

// 创建 JWT token 返回给用户登录后的。
func generateToken(username string) (string, error) {
	//这里可以是数据库调用  看用户名是否对的上？

	expirationTime := time.Now().Add(jwtExpirationDelta) //表明期限
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(), // .Unix将时间期限转换为时间戳，这样如果是跨时区，会更加方便。
	}) //对 ，声明验证  token中的username是否正确 和 期限是否到期

	signedToken, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err //错误返回空令牌
	}
	return signedToken, nil
}
