package main

import (
	"net"
	"fmt"
	_"errors"
	"encoding/binary"
	"encoding/json"
	"application/common/message"
)
// 读
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

// 写
func writepkg(conn net.Conn, data []byte) (err error) {

	// 先发一个长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write() err")
		return
	}
	
	// 发送Data数据本身
	n, err = conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.Write() err")
		return
	}
	return
}