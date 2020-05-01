package message

const (
	SignInMesType = "SginMes"
	SignInResMesType = "SignInResMes"
	SignUpMesType = "SignUpMes"
	SignUpResMesType = "SignUpResMes"
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