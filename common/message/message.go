package message

const (
	SignInMesType = "SginMes"
	SignInResMesType = "SignInResMes"
	SignUpMesType = "SignUpMes"
	SignUpResMesType = "SignUpResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)

// 用户在线状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`
	Data string	`json:"data"`
}

// 登录消息类型
type SignInMes struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userpwd"`
	UserName string `json:"userName"`
}

// 500	没有注册
// 200	登录成功
type SignInResMes struct {
	Code int `json:"code"`
	UsersId []int // 保存用户id的切片
	Error string `json:"error"`//没有错误返回nil
}

// 注册类型
type SignUpMes struct {
	User User `json:"user"`
}

// 注册响应类型
type SignUpResMes struct {
	Code int `json:"code"`
	Error string `json:"error"`//没有错误返回nil
}

// 配合服务器推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}