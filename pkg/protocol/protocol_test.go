package protocol

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

	assert.Equal(t, byte(0x40), frame.checksum(), "checksum failed")
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
