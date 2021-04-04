package client

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Client_GetPartialStatusError(t *testing.T) {
	client, server := net.Pipe()
	// This is an example of a get partial status command
	expectedByServer := []byte{0x08, 0xe9, 0x21, 0x31, 0x32, 0x33, 0x34, 0x5a, 0x21, 0x40}
	// This is an example of a invalid password response
	expectedByClient := []byte{0x02, 0xe9, 0xe1}
	go func() {
		in := make([]byte, 10)
		_, err := server.Read(in)
		assert.NoError(t, err)
		assert.Equal(t, expectedByServer, in)
		_, err = server.Write(expectedByClient)
		assert.NoError(t, err)
	}()

	c := Client{
		conn:     client,
		password: "1234",
	}

	_, err := c.GetPartialStatus()
	assert.Error(t, err, "Invalid password")
}

func Test_Client_parceZones(t *testing.T) {
	// Zone 1: violated false, anulated false, open false, shortCircuit false, lowBattery false, tamper false
	// Zone 10 violated true, anulated true, open true, shortCircuit true, lowBattery true, tamper true
	data := []byte{
		0xe9,
		0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00,
		0x40}

	response := parseZones(data)

	assert.False(t, response[0].anulated)
	assert.False(t, response[0].open)
	assert.False(t, response[0].violated)

	assert.True(t, response[9].anulated)
	assert.True(t, response[9].open)
	assert.True(t, response[9].violated)
}
