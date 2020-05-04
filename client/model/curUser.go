package model

import (
	"net"
	"application/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}