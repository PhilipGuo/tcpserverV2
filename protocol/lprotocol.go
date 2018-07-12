package protocol

import (
	"fmt"
	"net"
)

// ProtocolV1 is v1 of protocol
//
type ProtocolV1 struct {
}

// NewProtocolV1 return New ProtocolV1 instance
//
func NewProtocolV1() *ProtocolV1 {
	return &ProtocolV1{}
}

// IOLoop io loop
//
func (p *ProtocolV1) IOLoop(con net.Conn) {
	session := NewSession(con)
	fmt.Println("client connected:", session.String())
	for {
		p := make([]byte, 512)
		cnt, err := session.Read(p)
		if nil != err {
			fmt.Println("read err")
			continue
		}

		apply := string(p[:cnt])
		if apply == "apply" {
			HeartBeatReponse(session)
		}
	}
}

//HeartBeatReponse reply
//
func HeartBeatReponse(session *Session) {
	if nil == session {
		return
	}
	session.con.Write([]byte("reply"))
}
