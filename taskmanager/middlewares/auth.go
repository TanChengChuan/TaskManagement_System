package middlewares

import (
	"TaskManagement_System/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"net/http"
)

// - **认证中间件：** 验证请求头中的JWT token，确保用户已认证。

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := viper.GetString("token")
		tokenString, err := jwt.ParseWithClaims(token, &utils.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
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
		//2024.1111.这里改变了一下逻辑，测试一下是不是没储存到用户信息-》不是 -》 换成配置读取 成功！！！！
		if tokenString != nil {
			claims := tokenString.Claims.(*utils.MyClaims) //&& tokenString.Valid { //记得和jwt.go文件合在一起.
			c.Set("username", claims.Username)
			//储存用户信息 在上下文中？ -》这样表明用户通过认证后，可以进行操作.
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
			//} //解析出来 token 非法
		}
		c.Next()

	}
}
