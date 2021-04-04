package protocol

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	data := ISECNetMobileFrame{
		password: []byte{'1', '2', '3', '4'},
		command:  []byte{0x5a},
	}
	frame := ISECNetFrame{
		command: COMMAND,
		data:    data,
	}

	if frame.checksum() != 0x40 {
		t.Errorf("checksum failed. Expected %x, got %x", 0x40, frame.checksum())
	}
}

func TestBytes(t *testing.T) {
	data := ISECNetMobileFrame{
		password: []byte{'1', '2', '3', '4'},
		command:  []byte{0x5a},
	}
	frame := ISECNetFrame{
		command: COMMAND,
		data:    data,
	}

	bytes := frame.bytes()
	expected := []byte{0x08, 0xe9, 0x21, 0x31, 0x32, 0x33, 0x34, 0x5a, 0x21, 0x40}
	for i := range bytes {
		if bytes[i] != expected[i] {
			t.Errorf("serialization failed. Expected %x, got %x", expected, bytes)
			break
		}
	}
}
