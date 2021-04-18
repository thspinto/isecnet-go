package alarm

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_command(t *testing.T) {
	client, server := net.Pipe()
	// This is an example of a get partial status command
	expectedByServer := []byte{0x08, 0xe9, 0x21, 0x31, 0x32, 0x33, 0x34, 0x5a, 0x21, 0x40}
	// This is an example of a invalid password response
	expectedByClient := []byte{0x02, 0xe9, 0xe1, 0xf5}
	go func() {
		in := make([]byte, 10)
		_, err := server.Read(in)
		assert.NoError(t, err)
		assert.Equal(t, expectedByServer, in)
		_, err = server.Write(expectedByClient)
		assert.NoError(t, err)
	}()

	c := Client{
		conn: client,
	}
	r, err := c.command(expectedByServer)
	assert.Nil(t, err)
	assert.Equal(t, expectedByClient[1:], r)
}
