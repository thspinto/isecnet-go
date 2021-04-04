package network

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewTCPClient(t *testing.T) {
	ln, err := net.Listen("tcp", ":12345")
	assert.NoError(t, err)
	NewTCPClient("localhost", "12345")
	_, err = ln.Accept()
	assert.NoError(t, err)
}

func Test_Command(t *testing.T) {
	client, server := net.Pipe()
	// This is an example of a get partial status command
	expectedByServer := []byte{0x08, 0xe9, 0x21, 0x31, 0x32, 0x33, 0x34, 0x5a, 0x21, 0x4b}
	// This is an example of a invalid password response
	expectedByClient := []byte{0x02, 0xe2, 0xe1}
	go func() {
		in := make([]byte, 10)
		_, err := server.Read(in)
		assert.NoError(t, err)
		assert.Equal(t, expectedByServer, in)
		_, err = server.Write(expectedByClient)
		assert.NoError(t, err)
	}()

	c := TCPClient{
		conn: client,
	}
	assert.Equal(t, expectedByClient[1:], c.Command(expectedByServer))
}
