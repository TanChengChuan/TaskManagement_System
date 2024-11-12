package tests

import (
	"TaskManagement_System/controllers"
	"TaskManagement_System/models"
	"TaskManagement_System/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"testing"
	"time"
)

///编写单元测试和集成测试，覆盖核心功能。

//auth -》 user

func Test_auth(t *testing.T) {
	viper.SetConfigName("config")
	viper.AddConfigPath("E:\\Goland\\TaskManagement_System\\taskmanager\\config")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	var err error
	e := gin.Default()
	models.DB, err = gorm.Open(mysql.Open(models.DSN))
	if err != nil {
		fmt.Println("db connect error")
	}
	e.Use(func(c *gin.Context) {
		c.Set("db", models.DB)
		c.Next()
	})
	slog.Info("db connect success")
	//---
	var i int
	fmt.Printf("请输入你要选的选项\n\t[1]注册账号\n\t[2]登录账号\n\t（第一次使用请选择注册账号）\n")
	fmt.Scanf("%d", &i) //错误判断。
	routes.SetupRouter(e)
	port := viper.GetString("server.port")
	log.Printf("Starting server on port %s", port)
	//if err := e.Run(":" + port); err != nil {
	//	log.Fatalf("Failed to start server: %s", err)
	//}
	switch int(i) {
	case 1: //注册
		{

			controllers.TMuser(e)
			fmt.Println("注册完成后，请重新登录账号.")
			fmt.Println("重新启动中...")
			time.Sleep(time.Second)
			controllers.ClearScreen()
			//有一个问题，它的context 上下文 数据在不同包是一样的吗? -》 在一个应用程序是的，所以可以一样
		}

	case 2: //登录
		{
			//补充： 首先根据用户 寻找其OwnerID 然后对应任务， 这样用户进行操作的时候就会比较方便，
			//像是：创建关联，查找关联，删除关联，更新关联。
			fmt.Println("请输入账号密码")
			controllers.TMuser(e)
			//在task内也可以加一个循环，然后选择所选的
			controllers.TMTask(e)
			//这里应该只有报错 或者 说是自行选择退出 才可以出来了。所以加一个break,直接退出。
			break
		}

	default:
		{
			fmt.Println("输入错误，请重新输入")
			return
		}

	}
}
