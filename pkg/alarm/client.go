package alarm

import (
	"fmt"
	"io"
	"net"

	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination mock_alarm/mock_alarm.go github.com/thspinto/isecnet-go/pkg/alarm AlarmClient

// AlarmClient is the client interface
type AlarmClient interface {
	GetPartialStatus() (response *StatusResponse, err error)
	GetZones() ([]ZoneModel, error)
}

// Client is the client for
// communicating with the alarm central
type Client struct {
	host     string
	port     string
	conn     net.Conn
	password string
}

// NewClient returns client for communicating with the server through a tcp connection
func NewClient(host, port, password string) (client AlarmClient) {
	log.WithFields(log.Fields{
		"address": host + ":" + port,
	}).Info("Connecting...")

	client = &Client{
		host:     host,
		port:     port,
		password: password,
	}

	return
}

func (c *Client) connect() (err error) {
	conn, err := net.Dial("tcp", c.host+":"+c.port)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Failed to establish TCP connection")
		return
	}
	c.conn = conn
	return
}

// command dispatches the command and return the response
func (c *Client) command(b []byte) (response []byte, err error) {
	if c.conn == nil {
		if err = c.connect(); err != nil {
			return
		}
	}

	_, err = c.conn.Write(b)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Failed writing to stream")
		return
	}

	sizeBuf := make([]byte, 1)
	_, err = io.ReadFull(c.conn, sizeBuf)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Failed reading response size")
		return
	}

	size := int64(sizeBuf[0])
	log.WithFields(log.Fields{
		"size": fmt.Sprintf("%x", size),
	}).Debug("Response Size")

	response = make([]byte, size+1)
	_, err = io.ReadFull(c.conn, response)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Failed reading response")
		return
	}

	return
}