package client

import (
	"errors"
	"time"
)

type Zone struct {
	Violated     bool
	Anulated     bool
	Open         bool
	LowBattery   bool // not implemented
	ShortCircuit bool // not implemented
	Tamper       bool // not implemented
}

type PGM struct {
	Enabled bool
}

type Keyboard struct {
	IssueWarn bool // keyboard has some issue
	Tamper    bool
}
type Central struct {
	Model                     string
	Firmware                  string
	Siren                     Siren
	Battery                   Battery
	Activated                 bool
	Alerting                  bool
	IssueWarn                 bool // central has some issue
	ExternalPowerFault        bool // central loss external power source
	ExternalAuxOverload       bool // external auxiliary exit overload
	PhoneLineCut              bool
	EventCommunicationFailure bool
}

type Siren struct {
	Enabled      bool
	WireCut      bool
	ShortCircuit bool
}

type Battery struct {
	Low           bool
	ShortCircuit  bool
	Level         int
	BypassEnabled bool
	BypassBlink   bool
}

type StatusResponse struct {
	Time             time.Time
	Zones            []Zone
	Keyboards        []Keyboard
	Central          Central
	PGM              PGM
	PartitionEnabled bool
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

	status := StatusResponse{
		Zones: parseZones(response),
	}

	return &status, nil
}

func parseZones(b []byte) []Zone {
	zones := make([]Zone, 48)

	for i := 0; i < int(len(zones)/8); i++ {
		for j := 0; j < 8; j++ {
			sensor := (i * 8) + j
			zones[sensor].Open = (b[i+1]>>j)&0x01 == 0x01
			zones[sensor].Violated = (b[i+9]>>j)&0x01 == 0x01
			zones[sensor].Anulated = (b[i+17]>>j)&0x01 == 0x01
		}
	}
	return zones
}
