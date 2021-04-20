package main

import (
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	go main()
}

func readResponse(conn net.Conn) ([]byte, error) {
	sizeBuf := make([]byte, 1)
	_, err := io.ReadFull(conn, sizeBuf)
	if err != nil {
		return nil, err
	}
	size := int64(sizeBuf[0])
	buf := make([]byte, size)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return nil, err
	}

	return append(sizeBuf, buf...), err
}

func Test_InvalidPasswordResponse(t *testing.T) {
	// Example of get partial status request
	getPartialStatus := []byte{0x08, 0xe9, 0x21, 0x31, 0x32, 0x33, 0x35, 0x5a, 0x21, 0x41}
	// This is an example of a invalid password response
	expectedResponse := []byte{0x02, 0xe9, 0xe1}

	conn, err := net.Dial("tcp", "localhost:9009")
	assert.NoError(t, err)
	_, err = conn.Write(getPartialStatus)
	assert.NoError(t, err)
	response, err := readResponse(conn)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func Test_InvalidPacketResponse(t *testing.T) {
	// Example of get partial status request
	getPartialStatus := []byte{0x08, 0xe9, 0x21, 0x31, 0x32, 0x33, 0x35, 0x5a, 0x21, 0x40}
	// This is an example of a invalid password response
	invalidPasswordResponse := []byte{0x02, 0xe9, 0xe0}

	conn, err := net.Dial("tcp", "localhost:9009")
	assert.NoError(t, err)
	_, err = conn.Write(getPartialStatus)
	assert.NoError(t, err)
	response, err := readResponse(conn)
	assert.NoError(t, err)
	assert.Equal(t, invalidPasswordResponse, response)
}

func Test_InvalidCommandResponse(t *testing.T) {
	// Example of get partial status request
	getPartialStatus := []byte{0x08, 0xe1, 0x21, 0x31, 0x32, 0x33, 0x35, 0x5a, 0x21, 0x49}
	// This is an example of a invalid password response
	expectedResponse := []byte{0x02, 0xe9, 0xe2}

	conn, err := net.Dial("tcp", "localhost:9009")
	assert.NoError(t, err)
	_, err = conn.Write(getPartialStatus)
	assert.NoError(t, err)
	response, err := readResponse(conn)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, response)
}

func Test_PartialStatusResponse(t *testing.T) {
	// Example of get partial status request
	getPartialStatus := []byte{0x08, 0xe9, 0x21, 0x31, 0x32, 0x33, 0x34, 0x5a, 0x21, 0x40}

	conn, err := net.Dial("tcp", "localhost:9009")
	assert.NoError(t, err)
	_, err = conn.Write(getPartialStatus)
	assert.NoError(t, err)
	response, err := readResponse(conn)
	assert.NoError(t, err)
	assert.True(t, len(response) > 4)
}
