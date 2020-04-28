package main

import (
	"io"
	"net"
	"fmt"
	"application/server/utils"
	"application/server/process"
	"application/common/message"
)
type Processor struct {
	Conn net.Conn
}

// 根据消息类型不同决定用哪个函数来处理
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	
	switch mes.Type {
		case message.SignInMesType:
			// 处理登录
			up := &processcon.UserProcess{
				Conn: this.Conn,
			}
			err = up.ServerProcessSignin(mes)
		case message.SignUpMesType:
			// 处理注册
		default:
			fmt.Println("消息类型不存在........")
	} 
	return
}

func (this *Processor) Process2() (err error) {
	for {
		// 读取客户端发送的数据
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.Readpkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器自动断开... ")
			} else {
				fmt.Println("readpkg() err ")
			}
			return err
		}
		// fmt.Println("mes = ", mes)

		err = this.ServerProcessMes( &mes)
		if err != nil {
			fmt.Println("ServerProcessMes() err", err)
			return err
		}
	}
}