package protocol

import (
	"net"
)

// Protocol interface
//
type Protocol interface {
	IOLoop(con net.Conn)
}
