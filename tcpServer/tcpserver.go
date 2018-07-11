package tcpserver

import (
	"net"
	protocol "tcpServerV2/protocol"
)

// Tcpserver object
//
type Tcpserver struct {
}

// NewTcpserver return instance of Tcpserver
func NewTcpserver() *Tcpserver {
	return &Tcpserver{}
}

// Handle connection
//
func (s *Tcpserver) Handle(con net.Conn) {
	if nil == s {
		return
	}

	var prot protocol.Protocol
	prot = protocol.NewProtocolV1()
	prot.IOLoop(con)
	// prot = &prot.IOLoop(con)
}
