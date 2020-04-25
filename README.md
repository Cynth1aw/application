# application

海量用户通讯系统 - 练习版

 #Redis
	执行下面命令获取redis包 需要到GOPATH路劲下执行  我的这里时cd到GoProgram执行go get github.com/garyburd/redigo/redis  
	go get github.com/garyburd/redigo/redis   
    	import "github.com/garyburd/redigo/redis"

#编译client
	PS D:\GoProgram> go build -o client.exe  application/client

#编译服务器
	PS D:\GoProgram> go build -o server.exe application/server
