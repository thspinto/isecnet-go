package network

import (
	"io"
	"net"

	"github.com/thspinto/isecnet-go/pkg/handlers"
)

type TCPClient struct {
	host string
	port string
	conn net.Conn
}

// NewTCPClient returns client for communicating with the server through a tcp connection
func NewTCPClient(host, port string) TCPClient {
	conn, err := net.Dial("tcp", host+":"+port)
	handlers.CheckError("tcp connection failed", err)

	return TCPClient{
		host: host,
		port: port,
		conn: conn,
	}
}

// Command dispatches the command and return the response
func (c *TCPClient) Command(b []byte) []byte {
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
