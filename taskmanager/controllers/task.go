package controllers

import (
	"TaskManagement_System/models"
	"github.com/gin-gonic/gin"
)

// 数据库后面加连接池优化。
// 用户名:密码啊@tcp(ip:端口)/数据库的名字
// 引入数据库
// /迁移数据库
func TMTask(e *gin.Engine) {
	e.Use(func(c *gin.Context) {
		c.Set("db", models.DB)
		c.Next()
	})
	//这里加一个switch - case

	models.DB.AutoMigrate(&models.Task{}) //每一次都迁移 ->可以优化为 创建，删除，更新，导入时 再迁移。 （后面再说）

}
