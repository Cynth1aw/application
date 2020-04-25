package main

import (
	"net"
	"fmt"
)

func process(conn net.Conn) {
	fmt.Println("等待发送消息")
	defer conn.Close()
	for {
		buf := make([]byte, 8096)
		n, err := conn.Read(buf[:4])
		if n != 4 || err !=  nil {
			fmt.Println("conn.Read() err")
			return
		}
		fmt.Println(buf[:4])
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
		go process(conn)
	}
}