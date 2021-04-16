package main

import (
	"fmt"
	"log"

	"github.com/panjf2000/gnet"
)

type Server struct {
	*gnet.EventServer
	count int
}

func (es *Server) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	if frame[len(frame)-1] != checksum(frame) {
		out = []byte{0x2, 0xe9, 0xe0}
		out = append(out, checksum(out))
		fmt.Println("Invalid packet")
		fmt.Printf("%b", checksum(frame))
	} else if frame[1] != 0xe9 {
		out = []byte{0x2, 0xe9, 0xe2}
		out = append(out, checksum(out))
		fmt.Println("Invalid command")
	} else if string(frame[3:7]) != string([]byte{0x31, 0x32, 0x33, 0x34}) {
		out = []byte{0x2, 0xe9, 0xe1}
		out = append(out, checksum(out))
		fmt.Println("Invalid password")
	} else if frame[7] == 0x5a {
		fmt.Printf("Partial Status: %v\n", es.count)
		zones := [][]byte{
			{
				// Zone 1: ok
				// Zone 2: open
				// Zone 3: anulated
				// Zone 4: open and anulated
				// Zone 5: violated
				// Zone 6: Anulated and violated
				0x2c,
				0xe9,
				0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x2c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00,
			},
			{
				// Zone 1: open
				// Zone 2: violated
				// Zone 3: ok
				// Zone 4: ok
				// Zone 5: ok
				// Zone 6: ok
				0x2c,
				0xe9,
				0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00,
			},
			{
				// Zone 1: ok
				// Zone 2: anulated
				// Zone 3: ok
				// Zone 4: ok
				// Zone 5: ok
				// Zone 6: ok
				0x2c,
				0xe9,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00,
			},
			{
				// Zone 1-8: ok
				// Zone 8-16: open
				0x2c,
				0xe9,
				0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00,
			},
			{
				// Zone 1-8: open
				// Zone 8-16: ok
				0x2c,
				0xe9,
				0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00,
			},
		}
		out = zones[es.count%len(zones)]
		out = append(out, checksum(out))
		es.count++
	}
	return
}

func checksum(b []byte) byte {
	sum := byte(0xff)
	for _, b := range b[:len(b)-1] {
		sum = sum ^ b
	}
	return sum
}

// Serve starts the server
func main() {
	s := new(Server)
	fmt.Println("Running server...")
	log.Fatal(gnet.Serve(s, "tcp://:9009"))
}
