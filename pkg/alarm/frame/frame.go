package frame

const (
	FRAME_DELIMITER = 0x21
	COMMAND         = 0xE9
)

// ShortResponse is the representation of a short response from the server
type ShortResponse struct {
	Ack         bool
	Description string
}

// ISECNet is the data frame in the ISECNet frame
type ISECNet struct {
	Command byte
	Data    ISECNetMobile
}

// ISECNetMobile is the data frame in the ISECNet frame
type ISECNetMobile struct {
	Password []byte
	Command  []byte
	Content  []byte
}

func (frame *ISECNetMobile) bytes() []byte {
	bytes := []byte{FRAME_DELIMITER}
	bytes = append(bytes, frame.Password...)
	bytes = append(bytes, frame.Command...)
	bytes = append(bytes, frame.Content...)
	bytes = append(bytes, FRAME_DELIMITER)
	return bytes
}

// Bytes returns the frame representation in bytes
func (frame *ISECNet) Bytes() []byte {
	data := frame.Data.bytes()
	size := len(data) + 1
	bytes := []byte{byte(size), frame.Command}
	bytes = append(bytes, data...)
	bytes = append(bytes, Checksum(bytes))
	return bytes
}

// Checksum calculates the integrity byte
// The integrity byte is the last byte of the
// message and is an xor of all bytes of the frame.
func Checksum(b []byte) byte {
	sum := byte(0xFF) // two's complement
	for _, b := range b {
		sum = sum ^ b
	}
	return sum
}

// Gets the translation for a short response
func GetShortResponse(b []byte) ShortResponse {
	frames := map[byte]ShortResponse{
		0xfe: {
			Ack:         true,
			Description: "Command successfully received",
		},
		0xe0: {
			Ack:         false,
			Description: "Invalid packet format",
		},
		0xe1: {
			Ack:         false,
			Description: "Invalid password",
		},
		0xe2: {
			Ack:         false,
			Description: "Invalid command",
		},
		0xe3: {
			Ack:         false,
			Description: "Central not partitioned",
		},
		0xe4: {
			Ack:         false,
			Description: "Open zones",
		},
		0xe5: {
			Ack:         false,
			Description: "Deprecated command",
		},
		0xe6: {
			Ack:         false,
			Description: "User does not have permission to bypass",
		},
		0xe7: {
			Ack:         false,
			Description: "User does not have permission to deactivate",
		},
	}

	return frames[b[1]]
}
