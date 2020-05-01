package model

import (
	"encoding/json"
	"application/common/message"
	"github.com/garyburd/redigo/redis"
)

// 服务器启动时就初始化一个全局的UserDao
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	return &UserDao{
		pool: pool,
	}
}

func (this *UserDao) getUserById(conn redis.Conn,id int) (user *User,err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		// err == redis.ErrNil表示查不到对应的数据
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	// 把res反序列化
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		return
	}
	return
}

// 校验参数是否正确
func (this *UserDao) SignIn(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	if userPwd != user.UserPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

// 校验参数是否正确
func (this *UserDao) SignUp(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_XISTS
		return
	}

	// 序列化写库完成注册
	data, err := json.Marshal(user)
	if err == nil {
		return
	}
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err == nil {
		return
	}
	return
}