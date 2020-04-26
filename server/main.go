package main

import (
	"io"
	"net"
	"fmt"
	_"errors"
	"encoding/binary"
	"encoding/json"
	"application/common/message"
)

func readpkg(conn net.Conn) (mes message.Message,err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据......")
	// 这里是双向的，只有两边都不断开才会阻塞在这里
	n, err := conn.Read(buf[:4])
	if n != 4 || err !=  nil {
		fmt.Println("conn.Read() err")
		// 返回自定义错误
		// err = errors.New("read pkg header err")
		return
	}
	// fmt.Println(buf[:4])
	// 根据buf[:4]转成uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])

	// 根据pkgLen读取消息内容
	n,err = conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read() err ")
		// 返回自定义错误
		// err = errors.New("read pkg body err")
		return
	}

	// 把pkgLen反序列化成message类型
	err = json.Unmarshal(buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal() err ")
		// 返回自定义错误
		// err = errors.New("read pkg header err")
		return
	}
	return 
}

func process(conn net.Conn) {
	fmt.Println("等待发送消息")
	defer conn.Close()
	
	for {
		// 读取客户端发送的数据
		mes, err := readpkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器自动断开... ")
			} else {
				fmt.Println("readpkg() err ")
			}
			
			return
		}
		fmt.Println("mes = ", mes)
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