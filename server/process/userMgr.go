package processcon

import (
	"fmt"
)

var (
	userMgr *UserMgr
)

// 维护在线用户
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成userMgr初始化
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 用户登录时添加到onlineUsers
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// 用户退出时删除onlineUsers
func (this *UserMgr) DelOnlineUsers(userId int) {
	delete(this.onlineUsers, userId)
}

// 查询返回当前所有在线的用户
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

// 根据ID返回
func (this *UserMgr) GetOnlineUsersById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err =  fmt.Errorf("该用户不在线...") 
	}
	return
}