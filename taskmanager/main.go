package main

import (
	"TaskManagement_System/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//先启动JWT认证，先让用户确定 注册还是登录
	var i int
	fmt.Printf("请输入你要选的选项\n\t[1]注册账号\n\t[2]登录账号\n\t（第一次使用请选择注册账号）\n")
	fmt.Scanf("%d", &i)
	switch int(i) {
	case 1:
		{
			e := gin.Default()
			routes.RegisterHandler(c * gin.Context) //有一个问题，它的context 上下文 数据在不同包是一样的吗?
		}

	case 2:
		{

		}

	default:
		{
		}

	}
}
