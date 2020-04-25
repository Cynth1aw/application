package message

const (
	SignInMesType = "SginMes"
	SignInResMesType = "SignInResMes"
)

type Message struct {
	Type string `json:"type"`
	Data string	`json:"data"`
}

// 消息类型
type SignInMes struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userpwd"`
	UserName string `json:"userName"`
}

// 500	没有注册
// 200	登录成功
type SignInResMes struct {
	Code int `json:"code"`
	Error string `json:"error"`//没有错误返回nil
}