package processcon

import (
	"net"
	"fmt"
	"encoding/json"
	"application/server/model"
	"application/server/utils"
	"application/common/message"
	
)

// UserProcessor
type UserProcess struct {
	Conn net.Conn
	// 增加一个字段，区分每个Conn的用户
	UserId int
}

// 通知所有在线的用户的方法
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		// 过滤自己
		if id == this.UserId {
			continue
		}
		// 开始通知 -----------这里.............
		up.NotifyMeOnlineUser(this.UserId)
	} 
}

func (this *UserProcess) NotifyMeOnlineUser(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline
	data, err := json.Marshal(notifyUserStatusMes)
	if err !=  nil {
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err !=  nil {
		return
	}

	// 数据已准备好，准备返回
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.Writepkg(data)
	return
}

// 处理注册的请求
func (this *UserProcess) ServerProcessSignup(mes *message.Message) (err error) {
	// 先从mes中取出mes.Data
	var signUpMes message.SignUpMes
	err = json.Unmarshal([]byte(mes.Data), &signUpMes)
	if err !=  nil {
		return
	}

	// 定义一个返回数据的message
	var resMes message.Message
	resMes.Type = message.SignUpResMesType
	var signUpResMes message.SignUpResMes

	err = model.MyUserDao.SignUp(&signUpMes.User)
	if err != nil {
		if err == model.ERROR_USER_XISTS {
			signUpResMes.Code = 505
			signUpResMes.Error = model.ERROR_USER_XISTS.Error()
		} else {
			signUpResMes.Code = 505
			signUpResMes.Error = "注册发生错误..."
		}
	} else {
		signUpResMes.Code = 200
	}

	// 序列化结构体
	data, err := json.Marshal(signUpResMes)
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
	
	user, err := model.MyUserDao.SignIn(signMes.UserId, signMes.UserPwd)
	if err != nil {
		signInResMes.Code = 500
		signInResMes.Error = "Account does not exist..."
	} else {
		fmt.Println(user)
		signInResMes.Code = 200
		// 登陆成功把用户存到在线列表的结构体
		this.UserId = signMes.UserId
		userMgr.AddOnlineUser(this)

		this.NotifyOthersOnlineUser(signMes.UserId)

		for id, _ := range userMgr.onlineUsers {
			signInResMes.UsersId = append(signInResMes.UsersId, id)
		}
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