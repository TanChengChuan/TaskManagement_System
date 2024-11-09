package main

import "C"
import (
	"TaskManagement_System/controllers"
	"fmt"
	"time"
)

func main() {
	//先让用户确定 注册还是登录(登录后给予JWT认证)
	//加一个for循环
	var i int
	fmt.Printf("请输入你要选的选项\n\t[1]注册账号\n\t[2]登录账号\n\t（第一次使用请选择注册账号）\n")
	fmt.Scanf("%d", &i)
	switch int(i) {
	case 1: //注册
		{

			controllers.TMuser(i)
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
			controllers.TMuser(i)
			//在task内也可以加一个循环，然后选择所选的
			controllers.TMTask()
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
