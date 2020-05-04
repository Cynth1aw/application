package process

import (
	"os"
	"fmt"
	"net"
	"encoding/json"
	"application/server/utils"
	"application/common/message"
)

// 显示登陆成功后的界面
func ShowMenu() {
	fmt.Println("----------xxxxxxxxxLanded successfully------------")
	fmt.Println("             1 Online user list")
	fmt.Println("             2 发送消息")
	fmt.Println("             3 信息列表")
	fmt.Println("             4 退出登录")
	fmt.Println("             请选择(1-4)")
	fmt.Println("--------------------------------------")
	fmt.Print("请选择您要操作的选项：")

	// 1、接收用户的输入
	var option int
	var content string
	// 定义在外部，之后不用在重复定义
	smsProcess := &SmsProcess{}
	// 这里使用Scanf 要注意加\n
	fmt.Scanf("%d\n", &option)
	switch option {
		case 1:
			// fmt.Println("在线列表")
			outputOnlineUser()
		case 2:
			// fmt.Println("发送消息")
			fmt.Println("请输入：")
			fmt.Scanln(&content)
			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("信息列表")
		case 4:
			// 这里可以发送一条消息，通知服务器
			os.Exit(0)
		default:
			fmt.Print("The input is incorrect, please re-enter\n")
	}
}

// 和服务器端保持通讯
func PercessServerMes(conn net.Conn) {
	// 创建一个transfer 不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn:conn,
	}
	for {
		mes, err := tf.Readpkg()
		if err != nil {
			fmt.Println("tf.Readpkg() err = ", err)
			return
		}
		// fmt.Println(mes)
		switch mes.Type {
			case message.NotifyUserStatusMesType:
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)

				UpdataUserStatus(&notifyUserStatusMes)
			case message.SmsMesType:
				// 处理服务器转发消息返回来的内容
				outputGroupMes(&mes)
			default :
				fmt.Println("error...")
		}
	}
}