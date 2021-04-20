package alarm

import (
	"context"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/thspinto/isecnet-go/pkg/alarm/frame"
)

// Zone is the raw zone status collected
// from the alarm central
type Zone struct {
	Violated     bool
	Anulated     bool
	Open         bool
	LowBattery   bool // not implemented
	ShortCircuit bool // not implemented
	Tamper       bool // not implemented
}

// PGM indicates the status of given PGM
type PGM struct {
	Enabled bool
}

// Partition indicates the status of given partition
type Partition struct {
	Enabled bool
}

// Keyboard is the status of a keyboard
type Keyboard struct {
	ReceiverIssue bool
	Issue         bool // keyboard has some issue
	Tamper        bool
}

// Central is all the information relative to the central
// collected from the partial status command
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

// Siren is the status of the siren
type Siren struct {
	Enabled      bool
	WireCut      bool
	ShortCircuit bool
}

// Battery is the status for the given battery
type Battery struct {
	Low              bool
	AbsentOrInverted bool
	ShortCircuit     bool
}

// StatusResponse is all the information
// collected from the partial status command
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
func (c *Client) GetPartialStatus(ctx context.Context) (response *StatusResponse, err error) {
	request := frame.ISECNet{
		Command: frame.COMMAND,
		Data: frame.ISECNetMobile{
			Password: []byte(c.password),
			Command:  []byte{0x5a},
		},
	}

	r, err := c.command(request.Bytes())
	if err != nil {
		return
	}
	log.WithFields(log.Fields{
		"response": fmt.Sprintf("%x", r),
	}).Debug("Partial Status Response")
	if len(r) <= 3 {
		return nil, errors.New(frame.GetShortResponse(r).Description)
	}

	response = &StatusResponse{
		Zones:      parseZones(r),
		Central:    parseCentral(r),
		Date:       parseDate(r),
		Keyboards:  parseKeyboard(r),
		Partitions: parsePartitions(r),
	}

	return
}

func parsePartitions(b []byte) []Partition {
	partitions := make([]Partition, 2)

	for i := range partitions {
		partitions[i].Enabled = (b[22]>>i)&0x01 == 0x01
	}

	return partitions
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
