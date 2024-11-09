package controllers

import (
	"TaskManagement_System/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func TMTask(e *gin.Engine) {
	//数据库
	//数据库后面加连接池优化。
	//用户名:密码啊@tcp(ip:端口)/数据库的名字
	//引入数据库
	///迁移数据库
	//// 将数据库实例存储到Gin的上下文中
	//每次启动要 数据库连接

	e.Use(func(c *gin.Context) {
		c.Set("db", models.DB)
		c.Next()
	})
	//这里加一个switch - case

	//可加循环，然后退出。-》加屁，错了
	models.DB.AutoMigrate(&models.Task{}) //每一次都迁移 ->可以优化为 创建，删除，更新，导入时 再迁移。 （后面再说）
	fmt.Println("请选择你想要的进行的选项\n\t[1]创建新任务\n\t[2]删除任务\n\t[3]更新任务\n\t[4]获取单项任务\n\t[5]获取所有任务\n\t[6]导入任务\n\t[7]退出程序")
	var i int
	fmt.Scanf("%d", &i)
	switch int(i) { //这玩意没啥用，本来就是靠URL实现的
	case 1:
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
		//应该是在main函数调用吗？->是的
	case 7:
		fmt.Println("期待与你的下一次相遇~")
		ClearScreen()
		return //os.Exit(0)

	default:
		fmt.Println("无效选项，请重新选择")
	}

}
