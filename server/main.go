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

// 处理登录的请求
func serverProcessSignin(conn net.Conn,mes *message.Message) (err error) {

	// 先从mes中取出mes.Data
	var signMes message.SignInMes
	err = json.Unmarshal([]byte(mes.Data), &signMes)
	if err !=  nil {
		fmt.Println("json.Unmarshal err")
		return
	}

	// 定义一个返回数据的message
	var resMes message.Message
	resMes.Type = message.SignInResMesType
	var signInResMes message.SignInResMes
	
	// 先暂定一个测试账号 1292033353 wowowowo0
	if signMes.UserId == 1292033353 && signMes.UserPwd == "wowowowo0" {
		signInResMes.Code = 200
	} else {
		signInResMes.Code = 500
		signInResMes.Error = "账户不存在..."
	}

	// 序列化结构体
	data, err := json.Marshal(signInResMes)
	if err != nil {
		fmt.Println("json.Marshal() fail",err)
	}
	resMes.Data = string(data)

	// 序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal() fail",err)
	}

	// 数据已准备好，准备返回
	err = writepkg(conn, data)
	return
}

// 根据消息类型不同决定用哪个函数来处理
func ServerProcessMes(conn net.Conn, mes *message.Message) (err error) {
	
	switch mes.Type {
		case message.SignInMesType:
			// 处理登录
			err = serverProcessSignin(conn, mes)
		case message.SignUpMesType:
			// 处理注册
		default:
			fmt.Println("消息类型不存在........")
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
		// fmt.Println("mes = ", mes)

		err = ServerProcessMes(conn, &mes)
		if err != nil {
			fmt.Println("ServerProcessMes() err", err)
			return
		}
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