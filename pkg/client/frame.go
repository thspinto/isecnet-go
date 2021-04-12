package client

const (
	FRAME_DELIMITER = 0x21
	COMMAND         = 0xE9
)

// ShortResponseFrame is the representation of a short response from the server
type ShortResponseFrame struct {
	ack         bool
	description string
}

// ISECNetMobileFrame is the data frame in the ISECNet frame
type ISECNetFrame struct {
	command byte
	data    ISECNetMobileFrame
}

// ISECNetMobileFrame is the data frame in the ISECNet frame
type ISECNetMobileFrame struct {
	password []byte
	command  []byte
	content  []byte
}

func (frame *ISECNetMobileFrame) bytes() []byte {
	bytes := []byte{FRAME_DELIMITER}
	bytes = append(bytes, frame.password...)
	bytes = append(bytes, frame.command...)
	bytes = append(bytes, frame.content...)
	bytes = append(bytes, FRAME_DELIMITER)
	return bytes
}

func (frame *ISECNetFrame) bytes() []byte {
	data := frame.data.bytes()
	size := len(data) + 1
	bytes := []byte{byte(size), frame.command}
	bytes = append(bytes, data...)
	bytes = append(bytes, frame.checksum())
	return bytes
}

func (frame *ISECNetFrame) checksum() byte {
	size := len(frame.data.bytes()) + 1
	sum := byte(size) ^ frame.command // xor ISECNet frame data
	for _, b := range frame.data.bytes() {
		sum = sum ^ b
	}
	sum = sum ^ 0xFF // two's complement

	return sum
}

// Gets the translation for a short response
func GetShortResponse(b []byte) ShortResponseFrame {
	frames := map[byte]ShortResponseFrame{
		0xfe: {
			ack:         true,
			description: "Command successfully received",
		},
		0xe0: {
			ack:         false,
			description: "Invalid packet format",
		},
		0xe1: {
			ack:         false,
			description: "Invalid password",
		},
		0xe2: {
			ack:         false,
			description: "Invalid command",
		},
		0xe3: {
			ack:         false,
			description: "Central not partitioned",
		},
		0xe4: {
			ack:         false,
			description: "Open zones",
		},
		0xe5: {
			ack:         false,
			description: "Deprecated command",
		},
		0xe6: {
			ack:         false,
			description: "User does not have permission to bypass",
		},
		0xe7: {
			ack:         false,
			description: "User does not have permission to deactivate",
		},
	}

	return frames[b[1]]
}
