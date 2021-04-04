package client

import (
	"errors"
	"time"
)

type Zone struct {
	violated     bool
	anulated     bool
	open         bool
	lowBattery   bool // not implemented
	shortCircuit bool // not implemented
	tamper       bool // not implemented
}

type PGM struct {
	enabled bool
}

type Keyboard struct {
	issueWarn bool // keyboard has some issue
	tamper    bool
}
type Central struct {
	model                     string
	firmware                  string
	siren                     Siren
	battery                   Battery
	activated                 bool
	alerting                  bool
	issueWarn                 bool // central has some issue
	externalPowerFault        bool // central loss external power source
	externalAuxOverload       bool // external auxiliary exit overload
	phoneLineCut              bool
	eventCommunicationFailure bool
}

type Siren struct {
	enabled      bool
	wireCut      bool
	shortCircuit bool
}

type Battery struct {
	low           bool
	shortCircuit  bool
	level         int
	bypassEnabled bool
	bypassBlink   bool
}

type StatusResponse struct {
	time             time.Time
	zones            []Zone
	keyboards        []Keyboard
	Central          Central
	pgm              PGM
	partitionEnabled bool
}

// GetPartialStatus get the partial status from the Central
func (c *Client) GetPartialStatus() (*StatusResponse, error) {
	request := ISECNetFrame{
		command: COMMAND,
		data: ISECNetMobileFrame{
			password: []byte(c.password),
			command:  []byte{0x5a},
		},
	}

	response := c.command(request.bytes())

	if len(response) <= 2 {
		return nil, errors.New(GetShortResponse(response).description)
	}

	return nil, nil
}

func parseZones(b []byte) []Zone {
	zones := make([]Zone, 48)

	for i := 0; i < int(len(zones)/8); i++ {
		for j := 0; j < 8; j++ {
			sensor := (i * 8) + j
			zones[sensor].open = (b[i+1]>>j)&0x01 == 0x01
			zones[sensor].violated = (b[i+9]>>j)&0x01 == 0x01
			zones[sensor].anulated = (b[i+17]>>j)&0x01 == 0x01
		}
	}
	return zones
}
