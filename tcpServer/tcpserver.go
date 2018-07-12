package tcpserver

import (
	"net"
	"tcp"
	protocol "tcpServerV2/protocol"
)

// Tcpserver object
//
type Tcpserver struct {
	buk *tcp.TCPConnBucket
}

// NewTcpserver return instance of Tcpserver
func NewTcpserver() *Tcpserver {
	return &Tcpserver{
		buk: tcp.NewTCPConnBucket(),
	}
}

// Handle connection
//
func (s *Tcpserver) Handle(con net.Conn) {
	if nil == s {
		return
	}

	s.buk.Put(con.RemoteAddr())
	var prot protocol.Protocol
	prot = protocol.NewProtocolV1()
	prot.IOLoop(con)
	// prot = &prot.IOLoop(con)
}
