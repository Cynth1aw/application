package process

import (
	"net"
	"fmt"
	"encoding/binary"
	"encoding/json"
	"application/server/utils"
	"application/common/message"
)

type UserProcess struct {
	// 字段 暂时不需要
}


func (this *UserProcess) SignIn(userId int, userPwd string) (err error) {
	
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
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.Readpkg()
	if err != nil {
		fmt.Println("client readpkg() err", err)
		return
	}
	var signInResMes message.SignInResMes
	err = json.Unmarshal([]byte(mes.Data), &signInResMes)
	if signInResMes.Code == 200 {
		// fmt.Println("signin success")
		
		// 启动一个协程 和服务器保持通讯，
		// 接收服务器的推送消息并显示在客户端
		go PercessServerMes(conn)

		// 登陆成功显示菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(signInResMes.Error)
	}
	return 
}