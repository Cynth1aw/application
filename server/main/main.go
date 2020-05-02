package main

import (
	"net"
	"fmt"
	"time"
	"application/server/model"
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

// 完成UserDao初始化任务
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	// 当服务器启动时就初始化连接池
	initPool("0.0.0.0:6379", 16, 0, 300 * time.Second)

	initUserDao()
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