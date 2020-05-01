package processcon

import (
	"net"
	"fmt"
	"encoding/json"
	"application/common/message"
	"application/server/utils"
)

// UserProcessor
type UserProcess struct {
	Conn net.Conn
}

// 处理登录的请求
func (this *UserProcess) ServerProcessSignin(mes *message.Message) (err error) {

	// 先从mes中取出mes.Data
	var signMes message.SignInMes
	err = json.Unmarshal([]byte(mes.Data), &signMes)
	if err !=  nil {
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
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.Writepkg(data)
	return
}