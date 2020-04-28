package main

import (
	"net"
	"fmt"
	_"application/common/message"
)

func processController(conn net.Conn) {
	fmt.Println("等待发送消息")
	defer conn.Close()
	
	// 创建总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.Process2()
	if err != nil {
		return
	}
}

func main() {
	listen, err := net.Listen("tcp","0.0.0.0:8888")
	defer listen.Close()

	if err !=  nil {
		fmt.Println("net.Listen() err")
		return
	}

	for {
		conn, err := listen.Accept()

		if err !=  nil {
			fmt.Println("listen.Accept() err")
		}
		go processController(conn)
	}
}