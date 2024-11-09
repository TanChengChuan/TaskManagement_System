package models

import (
	"TaskManagement_System/controllers"
	"TaskManagement_System/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

var dsn = "root:hsw20050529@tcp(127.0.0.1:3306)/taskmanagement?charset=utf8mb4&parseTime=True&loc=Local"
var db *gorm.DB

type Task struct {
	ID          uint      `gorm:"primary_key" json:"ID"`
	Title       string    ` json:"Title"`
	Description string    ` json:"Description"`
	Status      int       `json:"Status" ` //可以用数字 分别表达pending, in-progress, completed
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
	OwnerID     uint      `json:"OwnerID"`
}

type TaskBatch struct {
	Tasks []Task `json:"tasks"`
}

func TMTask() {
	//数据库
	//数据库后面加连接池优化。
	//用户名:密码啊@tcp(ip:端口)/数据库的名字
	//dsn := "root:hsw20050529@tcp(127.0.0.1:3306)/taskmanagement?charset=utf8mb4&parseTime=True&loc=Local"
	//引入数据库
	var err error
	e := gin.Default()
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		slog.Error("db connect error", err)
	}
	e.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	slog.Info("db connect success") //这里加一个switch - case
	fmt.Println("请选择你想要的进行的选项\n\t[1]创建新任务\n\t[2]删除任务\n\t[3]更新任务\n\t[4]获取单项任务\n\t[5]获取所有任务\n\t[6]导入任务\n\t[7]退出程序")
	//可加循环，然后退出。
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
		controllers.ClearScreen()
		return
	}
}
