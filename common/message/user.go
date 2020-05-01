package message

type User struct {
	// 这里不加tag 序列化和反序列化会失败
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserNam string `json:"userName"`
}