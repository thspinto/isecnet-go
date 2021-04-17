package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checksum(t *testing.T) {
	data := ISECNetMobileFrame{
		password: []byte{'1', '2', '3', '4'},
		command:  []byte{0x5a},
	}
	frame := ISECNetFrame{
		command: COMMAND,
		data:    data,
	}
	bytes := frame.bytes()

	assert.Equal(t, byte(0x40), bytes[len(bytes)-1], "checksum failed")
}

func Test_bytes(t *testing.T) {
	data := ISECNetMobileFrame{
		password: []byte{'1', '2', '3', '4'},
		command:  []byte{0x5a},
	}
	frame := ISECNetFrame{
		command: COMMAND,
		data:    data,
	}

	expected := []byte{0x08, 0xe9, 0x21, 0x31, 0x32, 0x33, 0x34, 0x5a, 0x21, 0x40}
	assert.Equal(t, expected, frame.bytes(), "serialization failed")
}

func Test_ShortResponseFrameGet(t *testing.T) {
	r := GetShortResponse([]byte{0xe9, 0xe1})
	expected := ShortResponseFrame{ack: false, description: "Invalid password"}
	assert.Equal(t, expected, r, "wrong short response translation")
}
