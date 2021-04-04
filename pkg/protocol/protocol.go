package protocol

const (
	FRAME_DELIMITER = 0x21
	COMMAND         = 0xE9
)

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
	bytes := []byte{byte(size), COMMAND}
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
