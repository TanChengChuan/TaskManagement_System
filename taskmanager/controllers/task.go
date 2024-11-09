package controllers

import (
	"TaskManagement_System/models"
	"TaskManagement_System/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
)

func TMTask() {
	//数据库
	//数据库后面加连接池优化。
	//用户名:密码啊@tcp(ip:端口)/数据库的名字
	//引入数据库
	///迁移数据库
	//// 将数据库实例存储到Gin的上下文中
	//每次启动要 数据库连接
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
	slog.Info("db connect success") //这里加一个switch - case

	//可加循环，然后退出。
	for {
		models.DB.AutoMigrate(&models.Task{}) //每一次都迁移
		fmt.Println("请选择你想要的进行的选项\n\t[1]创建新任务\n\t[2]删除任务\n\t[3]更新任务\n\t[4]获取单项任务\n\t[5]获取所有任务\n\t[6]导入任务\n\t[7]退出程序")
		var i int
		fmt.Scanf("%d", &i)
		switch int(i) {
		case 1:
			e.POST("/task", routes.CreateTask)
		case 2:
			e.DELETE("/task", routes.DeleteTask)
		case 3:
			e.PUT("/task/:id", routes.UpdateTask)
		case 4:
			e.GET("/task:id", routes.GetsingleTask)
		case 5:
			e.GET("/task", routes.GetallTasks)
		case 6:
			e.POST("/task/import", routes.ImportTask)
		case 7:
			fmt.Println("期待与你的下一次相遇~")
			ClearScreen()
			return //os.Exit(0)

		default:
			fmt.Println("无效选项，请重新选择")
		}
	}
}
