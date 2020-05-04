package processcon

import (
	"net"
	"encoding/json"
	"application/server/utils"
	"application/common/message"
)

type SmeProcess struct {

}

// 转发消息
func (this * SmeProcess) SendGroupMes(mes *message.Message) {
	
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		return
	}
	for id, up := range userMgr.onlineUsers {
		// 过滤自己
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmeProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn : conn,
	}
	err := tf.Writepkg(data)
	if err != nil {
		return
	}
	
}