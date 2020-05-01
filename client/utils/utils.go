package utils

import (
	"fmt"
	"net"
	"encoding/binary"
	"encoding/json"
	"application/common/message"
)

// 传输的结构体
type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}	

// 读
func (this *Transfer) Readpkg() (mes message.Message,err error) {

	fmt.Println("读取客户端发送的数据......")
	// 这里是双向的，只有两边都不断开才会阻塞在这里
	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err !=  nil {
		// 返回自定义错误
		// err = errors.New("read pkg header err")
		return
	}
	// fmt.Println(buf[:4])
	// 根据buf[:4]转成uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	// 根据pkgLen读取消息内容
	n,err = this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// 返回自定义错误
		// err = errors.New("read pkg body err")
		return
	}

	// 把pkgLen反序列化成message类型
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal() err ")
		// 返回自定义错误
		// err = errors.New("read pkg header err")
		return
	}
	return 
}

// 写
func (this *Transfer) Writepkg(data []byte) (err error) {

	// 先发一个长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if err != nil || n != 4 {
		fmt.Println("conn.Write() err")
		return
	}
	
	// 发送Data数据本身
	n, err = this.Conn.Write(data)
	if err != nil || n != int(pkgLen) {
		fmt.Println("conn.Write() err")
		return
	}
	return
}
