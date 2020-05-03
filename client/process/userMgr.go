package process

import (
	"fmt"
	"application/common/message"
)

// 客户端维护的在线用户map
var onlineUsers map[int]*message.User = make(map[int]*message.User)

// 在客户端显示当前在线的用户
func outputOnlineUser() {
	for id, user := range onlineUsers {
		fmt.Println("用户ID：", id, user)
	}
}

// 处理返回的NotifyUserStatusMes
func UpdataUserStatus(NotifyUserStatusMes *message.NotifyUserStatusMes) {
	
	user, ok := onlineUsers[NotifyUserStatusMes.UserId]

	if !ok {
		user = &message.User{
			UserId : NotifyUserStatusMes.UserId,
		}
	}
	
	user.UserStatus = NotifyUserStatusMes.Status

	onlineUsers[NotifyUserStatusMes.UserId] = user
}