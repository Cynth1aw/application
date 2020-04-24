package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
// 定义一个全局的pool
var pool *redis.Pool

// 初始化redis连接池
func init() {
	pool = &redis.Pool{
		MaxIdle: 8, //最大空闲链接数
		MaxActive: 0, // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个ip的redis
			return redis.Dial("tcp", "0.0.0.0:6379")
		},
	}
}

func main () {
	// ******************
	// Redis
	// 执行下面命令获取redis包 需要到GOPATH路劲下执行  我的这里时cd到GoProgram执行go get github.com/garyburd/redigo/redis  
	// go get github.com/garyburd/redigo/redis  
	// ******************
	
	// 从pool中取出redis链接
	conn := pool.Get()

	defer conn.Close()

	// 字符串类型
	_, err := conn.Do("set", "name", "lauren7ce")
	if err != nil {
		fmt.Println("redis set fail")
		return
	}

	//取出来的适合顺便转成字符串
	res, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("redis get fail")
		return
	}
	fmt.Println(res)

	//哈希类型
	_, err = conn.Do("HSet", "user01", "name", "cynth1aw")
	if err != nil {
		fmt.Println("redis set fail")
		return
	}
	_, err = conn.Do("HSet", "user01", "age", 20)
	if err != nil {
		fmt.Println("redis hset fail")
		return
	}

	//取出来的适合顺便转成字符串
	res, err = redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("redis hget fail")
		return
	}
	age, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("redis hget fail")
		return
	}
	fmt.Println(res,age)

	// 哈希批量操作
	_, err = conn.Do("HMSet", "user01", "nname", "Cynth1aw", "age", 21)
	if err != nil {
		fmt.Println("redis hset fail")
		return
	}
	// 同时取出多个值,返回的是一个集合 遍历一下即可取出值
	mres, err := redis.Strings(conn.Do("HMGet", "user01", "nname","age"))
	if err != nil {
		fmt.Println("redis hget fail")
		return
	}
	fmt.Println(mres)
	for _, v := range mres {
		fmt.Println(v)
	}
	
	//设置有效时间为10s
	_, err = conn.Do("expire", "name", 10)
}