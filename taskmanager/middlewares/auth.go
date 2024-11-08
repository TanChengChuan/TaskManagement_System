package middlewares

import (
	"TaskManagement_System/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

// - **认证中间件：** 验证请求头中的JWT token，确保用户已认证。

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is empty"})
			c.Abort()
			//当在请求处理过程中遇到无法恢复的错误时，可以使用c.Abort()来中断处理并返回错误响应
			return
		}

		//拆分 bearer 和 token  （正确的人 拿着正确的钥匙）
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is invalid"})
			c.Abort()
			return
		}

		//解析 token
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			//匹配成功
			return utils.JwtSecret, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"}) //签名错误
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Error parsing token"}) //解析报错
			}
			//后面加一个日志记录
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { //记得和jwt.go文件合在一起.
			//储存用户信息 在上下文中？ -》这样表明用户通过认证后，可以进行操作.
			username := claims["username"].(string)
			c.Set("username", username)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		} //解析出来 token 非法
		c.Next()
	}
}
