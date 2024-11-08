package models

import (
	"TaskManagement_System/routes"
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
	ID          uint
	Title       string
	Description string
	Status      int //可以用数字 分别表达pending, in-progress, completed
	CreatedAt   time.Time
	UpdatedAt   time.Time
	OwnerID     uint
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
	slog.Info("db connect success")
	e.POST("/task", routes.CreateTask)
	e.DELETE("/task", routes.DeleteTask)
	e.PUT("/task/:id", routes.UpdateTask)
	e.GET("/task:id", routes.GetsingleTask)
	e.GET("/task", routes.GetallTasks)
	e.POST("/task/import", routes.ImportTask)

}
