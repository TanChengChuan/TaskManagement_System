package models

import (
	"TaskManagement_System/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
)

// 加入数据库
type User struct {
	Username string `json:"username" bind:"required"`
	Password string `json:"password" bind:"required"`
}

func TMuser() {
	var erru error
	db, erru = gorm.Open(mysql.Open(dsn))
	if erru != nil {
		slog.Error("db connect error", erru)
	}

	e := gin.Default()
	e.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}) //这里不知道与task的 数据库连接 是否冲突。
	//要加一个经过认证 才能访问任务管理相关的接口,其实直接在 调用处加入一个if语句，然后接收JWTtoken 并且与密钥进行比对即可，

	e.POST("/POST", routes.RegisterHandler)
	e.POST("/POST", routes.LoginHandler)

}
