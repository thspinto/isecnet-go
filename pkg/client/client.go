package client

import (
	"io"
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/thspinto/isecnet-go/pkg/handlers"
)

type Client struct {
	host     string
	port     string
	conn     net.Conn
	password string
}

// NewClient returns client for communicating with the server through a tcp connection
func NewClient(host, port, password string) Client {
	log.WithFields(log.Fields{
		"address": host + ":" + port,
	}).Info("Connecting...")

	conn, err := net.Dial("tcp", host+":"+port)
	handlers.CheckError("tcp connection failed", err)

	return Client{
		host:     host,
		port:     port,
		conn:     conn,
		password: password,
	}
}

// command dispatches the command and return the response
func (c *Client) command(b []byte) []byte {
	_, err := c.conn.Write(b)
	handlers.CheckError("failed writing to stream", err)

	sizeBuf := make([]byte, 1)
	_, err = io.ReadFull(c.conn, sizeBuf)
	handlers.CheckError("failed reading response size", err)
	size := int64(sizeBuf[0])

	buf := make([]byte, size)
	_, err = io.ReadFull(c.conn, buf)
	handlers.CheckError("failed reading response", err)

	return buf
}
