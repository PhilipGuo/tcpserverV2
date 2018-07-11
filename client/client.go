package client

import (
	"errors"
	"net"
)

// Client object
//
type Client struct {
	con net.Conn
}

// NewClient return new client instance
//
func NewClient(con net.Conn) *Client {
	return &Client{
		con: con,
	}
}

func (c *Client) Connect() (bool, error) {
	if nil == c {
		return false, errors.New("client is nil")
	}

	return true, nil
}
