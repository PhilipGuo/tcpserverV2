package tcpserver

import (
	"fmt"
	"net"
	"strings"
)

// TCPHandler inerface
//
type TCPHandler interface {
	Handle(con net.Conn)
}

// TCPServer construct
//
func TCPServer(listener net.Listener, handler TCPHandler) {
	for {
		con, err := listener.Accept()
		if nil != err {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				// log
				fmt.Println("tempoary error")
				continue
			}
			if !strings.Contains(err.Error(), "use of closed network connection") {
				// log
				fmt.Println("use of closed con")
			}
			break
		}

		go handler.Handle(con)
	}
}
