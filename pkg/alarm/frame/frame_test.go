package frame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checksum(t *testing.T) {
	data := ISECNetMobile{
		Password: []byte{'1', '2', '3', '4'},
		Command:  []byte{0x5a},
	}
	frame := ISECNet{
		Command: COMMAND,
		Data:    data,
	}
	bytes := frame.Bytes()

	assert.Equal(t, byte(0x40), bytes[len(bytes)-1], "checksum failed")
}

func Test_bytes(t *testing.T) {
	data := ISECNetMobile{
		Password: []byte{'1', '2', '3', '4'},
		Command:  []byte{0x5a},
	}
	frame := ISECNet{
		Command: COMMAND,
		Data:    data,
	}

	expected := []byte{0x08, 0xe9, 0x21, 0x31, 0x32, 0x33, 0x34, 0x5a, 0x21, 0x40}
	assert.Equal(t, expected, frame.Bytes(), "serialization failed")
}

func Test_ShortResponseGet(t *testing.T) {
	r := GetShortResponse([]byte{0xe9, 0xe1})
	expected := ShortResponse{Ack: false, Description: "Invalid password"}
	assert.Equal(t, expected, r, "wrong short response translation")
}
