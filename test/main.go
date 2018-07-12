package main

import (
	"fmt"
	"net"
	"tcpServerV2/client"
	tcpserver "tcpServerV2/tcpserver"
	"time"
)

func main() {

	tcplistener, err := net.Listen("tcp", "127.0.0.1:8299")
	if nil != err {
		fmt.Println("listen err:", err.Error())
		return
	}

	server := tcpserver.NewTcpserver() // &tcpserver.Tcpserver{}
	var wrapper tcpserver.WaitGroupWrapper
	wrapper.Wrap(func() {
		tcpserver.TCPServer(tcplistener, server)
	})

	fmt.Println("server launched")

	for i := 0; i < 70000; i++ {
		fmt.Println("enter ", i)
		con, _ := net.Dial("tcp", "127.0.0.1:8299")
		cl := client.NewClient(con, i)
		cl.Connect()
		time.Sleep(time.Duration(1))
		fmt.Println("out ", i)
	}

	select {}
}
