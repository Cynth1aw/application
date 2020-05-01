package main

import (
	"os"
	"fmt"
	"application/client/process"
)
// 定义两个全局变量 ID 、密码
var (
	userId int
	userPwd string
)

func main() {
	// 1、接收用户的输入
	var option int
	// 2、判断是否继续显示菜单
	var loop bool = true

	for {
		fmt.Println("----------欢迎登陆聊天系统------------")
		fmt.Println("             1 用户登录")
		fmt.Println("             2 注册用户")
		fmt.Println("             3 退    出")
		fmt.Println("             请选择(1-3)")
		fmt.Println("--------------------------------------")
		fmt.Print("请选择您要操作的选项：")
		// 这里使用Scanf 要注意加\n
		fmt.Scanf("%d\n", &option)
		switch option {
			case 1:
				fmt.Print("请输入账户：")
				fmt.Scanln(&userId)
				fmt.Print("请输入密码：")
				fmt.Scanln(&userPwd)
				// signIn(userId, userPwd)
				up := &process.UserProcess{}
				up.SignIn(userId,userPwd)
				loop = false
			case 2:
				// signUp()
				loop = false
			case 3:
				os.Exit(0)
			default:
				fmt.Print("输入有误，请重新输入\n")
		}
		if !loop {
			break
		}
	}
}