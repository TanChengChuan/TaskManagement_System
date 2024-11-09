package controllers

import (
	"TaskManagement_System/middlewares"
	"TaskManagement_System/models"
	"TaskManagement_System/routes"
	"TaskManagement_System/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	"os"
	"os/exec"
	"sync"
	"time"
)

// 使用WaitGroup来等待所有goroutine完成
var Wg sync.WaitGroup

func ClearScreen() { //清个屏，稍微好看一点
	time.Sleep(200 * time.Millisecond)
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func TMuser(i int) {
	//每次都进行数据库的连接》？

	//用viper 读取zap日志
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return

	}
	// 将配置解析到结构体中
	var userconfig models.User
	var zapconfig utils.ZapConfig
	if err := viper.Unmarshal(&zapconfig); err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
		return
	}
	logger := utils.InitZap(&zapconfig)
	defer logger.Sync() // 确保在程序退出前刷新缓冲区
	//---
	var erru error
	models.DB, erru = gorm.Open(mysql.Open(models.DSN))
	models.DB.AutoMigrate(&models.User{}) //每一次都迁移 , 这里不仅是用户的账号和密码，也包括拥有任务的信息。
	if erru != nil {
		slog.Error("db connect error", erru)
	}

	e := gin.Default()
	e.Use(func(c *gin.Context) {
		c.Set("db", models.DB)
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
