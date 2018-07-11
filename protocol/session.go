package protocol

import (
	"net"
)

// Session struct
//
type Session struct {
	con net.Conn
}

// NewSession instance
//
func NewSession(con net.Conn) *Session {
	return &Session{
		con: con,
	}
}

func (s *Session) IOLoop() {

}
func (s *Session) String() string {
	return s.con.RemoteAddr().String()
}

func (s *Session) Read(bs []byte) (int, error) {

	return 0, nil
}
