package client

import (
	"errors"
	"fmt"
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

type Partition struct {
	Enabled bool
}

type Keyboard struct {
	ReceiverIssue bool
	Issue         bool // keyboard has some issue
	Tamper        bool
}
type Central struct {
	Model                     string
	Firmware                  string
	Battery                   Battery
	Activated                 bool
	Alerting                  bool
	IssueWarn                 bool // central has some issue
	ExternalPowerFault        bool // central loss external power source
	ExternalAuxOverload       bool // external auxiliary exit overload
	PhoneLineCut              bool
	EventCommunicationFailure bool
	Siren                     Siren
}

type Siren struct {
	Enabled      bool
	WireCut      bool
	ShortCircuit bool
}

type Battery struct {
	Low              bool
	AbsentOrInverted bool
	ShortCircuit     bool
}

type StatusResponse struct {
	Date             time.Time
	Zones            []Zone
	Keyboards        []Keyboard
	Central          Central
	PGMs             []PGM
	Partitions       []Partition
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
		Zones:     parseZones(response),
		Central:   parseCentral(response),
		Date:      parseDate(response),
		Keyboards: parseKeyboard(response),
	}

	return &status, nil
}

func parseKeyboard(b []byte) []Keyboard {
	keyboards := make([]Keyboard, 4)

	for i := range keyboards {
		keyboards[i].Issue = (b[30]>>i)&0x01 == 0x01
		keyboards[i].ReceiverIssue = (b[30]>>i+4)&0x01 == 0x01
		keyboards[i].Tamper = (b[32]>>(i+4))&0x01 == 0x01
	}

	return keyboards
}

func parseZones(b []byte) []Zone {
	zones := make([]Zone, 48)

	for i := 0; i < int(len(zones)/8); i++ {
		for j := 0; j < 8; j++ {
			zone := (i * 8) + j
			zones[zone].Open = (b[i+1]>>j)&0x01 == 0x01
			zones[zone].Violated = (b[i+7]>>j)&0x01 == 0x01
			zones[zone].Anulated = (b[i+13]>>j)&0x01 == 0x01
			if zone <= 39 {
				// for battery status there are only 40 zones
				zones[zone].LowBattery = (b[i+39]>>j)&0x01 == 0x01
			}
		}
		if i <= 8 {
			// tamper and short circuit are only for zones 1 to 8 and 11 to 18
			zones[i].Tamper = (b[34]>>i)&0x01 == 0x01
			zones[i].ShortCircuit = (b[36]>>i)&0x01 == 0x01
			zones[i+10].Tamper = (b[35]>>i)&0x01 == 0x01
			zones[i+10].ShortCircuit = (b[37]>>i)&0x01 == 0x01
		}
	}
	return zones
}

func parseDate(b []byte) time.Time {
	return time.Date(
		2000+int((b[28]>>4*10)+b[28]&0x0f),
		time.Month((b[27]>>4*10)+b[27]&0x0f),
		int((b[26]>>4*10)+b[26]&0x0f),
		int((b[24]>>4*10)+b[24]&0x0f),
		int((b[25]>>4*10)+b[25]&0x0f),
		0,
		0,
		time.Local,
	)
}

func parseCentral(b []byte) Central {
	c := Central{}

	if b[19] == 0x1e {
		c.Model = "AMT2018 E/EG"
	}
	c.Firmware = fmt.Sprintf("%v.%v", b[20]>>4, b[20]&0x0f)
	c.Activated = (b[23]>>3&0x01 == 0x01)
	c.Alerting = (b[23]>>2&0x01 == 0x01)
	c.IssueWarn = (b[23]>>0&0x01 == 0x01)

	c.ExternalPowerFault = (b[29]>>0&0x01 == 0x01)
	c.Battery = Battery{
		Low:              (b[29]>>1&0x01 == 0x01),
		AbsentOrInverted: (b[29]>>2&0x01 == 0x01),
		ShortCircuit:     (b[29]>>3&0x01 == 0x01),
	}
	c.ExternalAuxOverload = (b[29]>>4&0x01 == 0x01)
	c.Siren = Siren{
		Enabled:      (b[23]>>1 == 0x01) || (b[38]>>2&0x01 == 0x01),
		WireCut:      (b[33]>>0&0x01 == 0x01),
		ShortCircuit: (b[33]>>1&0x01 == 0x01),
	}
	c.PhoneLineCut = (b[33]>>2&0x01 == 0x01)
	c.EventCommunicationFailure = (b[33]>>3&0x01 == 0x01)
	return c
}
