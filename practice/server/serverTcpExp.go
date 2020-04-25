package main

import (
	"fmt"
	"net"
)

//net.Conn类型的参数
func process(conn net.Conn) {
	defer conn.Close()
	//循环接收内容
	for {
		buf := make([]byte, 1024)
		fmt.Printf("服务器在等待 IP = %v 发送信息\n", conn.RemoteAddr().String()) //获取链接的IP：conn.RemoteAddr().String()
		// conn.Read(buf)
		// 1 等待客户端通过conn发送信息
		// 2 如果客户端没有write[f发送]，那么协程就阻塞在这里
		n, err := conn.Read(buf) //从conn中读取内容
		
		if err == nil {
			fmt.Println("服务器端的Read err 退出")
			//这里要退出，否则会发生死循环
			return
		}
		//显示客户端发送的内容 注意：这里不能用Println()
		fmt.Print(string(buf[:n]))
	}
}

func main()  {
	fmt.Println("服务器开始监听了.........")

	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听失败........")
		fmt.Println("发送邮件通知Admin")
	}
	defer listen.Close()

	// fmt.Printf("suc = %v\n", listen)//查看类型

	for {
		fmt.Println("等待连接中........")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept()........")
			fmt.Println("发送邮件通知Admi2n")
		}
		fmt.Printf("suc = %v IP = %v\n", conn, conn.RemoteAddr().String()) //获取链接的IP：conn.RemoteAddr().String()
		//起一个协程为客户端服务
		go process(conn)
	}
}