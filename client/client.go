package main

import (
	"os"
	"fmt"
	"net"
	"bufio"
)

func main() {
	conn, err := net.Dial("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Println("连接失败.........")
		return
	}
	fmt.Printf("conn成功，suc=%v",conn)
	//发送聊天内容
	reader := bufio.NewReader(os.Stdin) //标准输入
	
	//从终端读取用户的输入
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失败.........")
	}
	//把line发送给服务器
	n, err := conn.Write([]byte(line))

	if err != nil {
		fmt.Println("conn.Write() err.........")
	}
	fmt.Printf("客户端发送了 %d 字节",n)
}