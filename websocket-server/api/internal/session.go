package internal

import (
	"net"
)

type Session struct {
	Uid    string
	Conn   *net.Conn
	events uint8
}

func NewSession(uid string, conn *net.Conn) (session *Session) {
	session = &Session{
		Uid:  uid,
		Conn: conn,
	}
	return
}
