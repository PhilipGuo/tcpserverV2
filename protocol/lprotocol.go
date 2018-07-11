package protocol

import (
	"fmt"
	"net"
)

// ProtocolV1 is v1 of protocol
//
type ProtocolV1 struct {
}

func NewProtocolV1() *ProtocolV1 {
	return &ProtocolV1{}
}
func (p *ProtocolV1) IOLoop(con net.Conn) {
	session := NewSession(con)
	fmt.Println("client connected:", session.String())
	for {
		//reader := bufio.NewReader(session)
	}
}
