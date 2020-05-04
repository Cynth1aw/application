package process

import (
	"encoding/json"
	"application/client/utils"
	"application/common/message"
)

type SmsProcess struct {

}

// 发送群聊消息
 func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//  创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	// 创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		return
	}

	mes.Data =  string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		return
	}

	tf := &utils.Transfer{
		Conn : CurUser.Conn,
	}
	tf.Writepkg(data)
	if err != nil {
		return
	}
	return
 }