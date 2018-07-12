package client

import (
	"errors"
	"fmt"
	"net"
	"time"
)

// Client object
//
type Client struct {
	con net.Conn
	cid int // client index
}

// NewClient return new client instance
//
func NewClient(con net.Conn, id int) *Client {
	return &Client{
		con: con,
		cid: id,
	}
}

// Connect ..
//
func (c *Client) Connect() (bool, error) {
	if nil == c {
		return false, errors.New("client is nil")
	}
	fmt.Println("connected:", c.con.LocalAddr())
	go c.heartBeat()
	go c.recv()

	return true, nil
}

// heartBeat
//
func (c *Client) heartBeat() {
	if nil == c {
		return
	}
	for {
		//fmt.Println("before apply")
		c.con.Write([]byte("apply"))
		//fmt.Println("after apply")
		time.Sleep(time.Second * 1)
		fmt.Println("after 1 sec ", time.Now().String())
	}
}

func (c *Client) recv() {
	if nil == c {
		return
	}

	for {
		p := make([]byte, 512)
		cnt, err := c.con.Read(p)
		if nil != err {
			fmt.Println(err.Error())
			break
		}

		fmt.Println(string(p[:cnt]), c.cid)
	}
}
