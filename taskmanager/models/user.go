package models

import (
	"TaskManagement_System/middlewares"
	"TaskManagement_System/routes"
	"TaskManagement_System/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
)

// 加入数据库
type User struct {
	Username string `json:"username" bind:"required"`
	Password string `json:"password" bind:"required"`
}

func TMuser(i int) {
	//用viper 读取zap日志
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return

	}
	// 将配置解析到结构体中
	var zapconfig utils.ZapConfig
	if err := viper.Unmarshal(&zapconfig); err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
		return
	}
	logger := utils.InitZap(&zapconfig)
	defer logger.Sync() // 确保在程序退出前刷新缓冲区
	//---
	var erru error
	db, erru = gorm.Open(mysql.Open(dsn))
	if erru != nil {
		slog.Error("db connect error", erru)
	}

	e := gin.Default()
	e.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}, middlewares.GinLogger(logger), middlewares.GinRecovery(logger, true)) //打印堆栈信息
	// 这里不知道与task的 数据库连接 是否冲突。
	//要加一个经过认证 才能访问任务管理相关的接口,其实直接在 调用处加入一个if语句，然后接收JWTtoken 并且与密钥进行比对即可，
	switch i {
	case 1:
		e.POST("/POST", routes.RegisterHandler)
	case 2:
		e.POST("/POST", routes.LoginHandler)
	}
}
