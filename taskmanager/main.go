package main

import (
	"TaskManagement_System/controllers"
	"TaskManagement_System/models"
	"TaskManagement_System/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"log/slog"
)

func main() {
	//先让用户确定 注册还是登录(登录后给予JWT认证)
	//加一个for循环
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	var err error
	e := gin.Default()
	models.DB, err = gorm.Open(mysql.Open(models.DSN))
	if err != nil {
		slog.Error("db connect error", err)
	}
	e.Use(func(c *gin.Context) {
		c.Set("db", models.DB)
		c.Next()
	})
	slog.Info("db connect success")
	routes.SetupRouter(e)
	controllers.TMTask(e)
	controllers.TMuser(e)
	port := viper.GetString("server.port")
	log.Printf("Starting server on port %s", port)
	if err := e.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
