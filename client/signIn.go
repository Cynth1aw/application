package main

import (
	"net"
	"fmt"
	_"time"
	"encoding/binary"
	"encoding/json"
	"application/common/message"
)
func signIn(userId int, userPwd string) (err error) {
	
	// 链接服务器
	conn, err := net.Dial("tcp","0.0.0.0:8888")
	defer conn.Close()
	if err !=  nil {
		fmt.Println("net.Dial() err")
		return
	}
	var mes message.Message
	//类型
	mes.Type = message.SignInMesType

	// 创建一个signIn结构体
	var signInMes message.SignInMes
	signInMes.UserId = userId
	signInMes.UserPwd = userPwd

	data, err := json.Marshal(signInMes)
	if err != nil {
		fmt.Println("json.Marshal() err")
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal() err")
		return
	}

	// 先把data的长度发过去  把长度转成一个切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write() err")
		return
	}

	// 发送真实消息
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write() err")
		return
	}

	// 休眠20s
	// time.Sleep(20 * time.Second)
	mes, err = readpkg(conn)
	if err != nil {
		fmt.Println("client readpkg() err", err)
		return
	}
	var signInResMes message.SignInResMes
	err = json.Unmarshal([]byte(mes.Data), &signInResMes)
	if signInResMes.Code == 200 {
		fmt.Println("signin success")
	} else {
		fmt.Println(signInResMes.Error)
	}
	return 
}