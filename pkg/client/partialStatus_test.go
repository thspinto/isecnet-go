package client

import (
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Client_GetPartialStatusError(t *testing.T) {
	client, server := net.Pipe()
	// This is an example of a get partial status command
	expectedByServer := []byte{0x08, 0xe9, 0x21, 0x31, 0x32, 0x33, 0x34, 0x5a, 0x21, 0x40}
	// This is an example of a invalid password response
	expectedByClient := []byte{0x02, 0xe9, 0xe1}
	go func() {
		in := make([]byte, 10)
		_, err := server.Read(in)
		assert.NoError(t, err)
		assert.Equal(t, expectedByServer, in)
		_, err = server.Write(expectedByClient)
		assert.NoError(t, err)
	}()

	c := Client{
		conn:     client,
		password: "1234",
	}

	_, err := c.GetPartialStatus()
	assert.Error(t, err, "Invalid password")
}

func Test_Client_parceZones(t *testing.T) {
	// Zone 1: violated false, anulated false, open false, shortCircuit false, lowBattery false, tamper false
	// Zone 10 violated true, anulated true, open true, shortCircuit true, lowBattery true, tamper true
	data := []byte{
		0xe9,
		0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x04,
		0x00, 0x00,
		0x40}

	response := parseZones(data)

	assert.False(t, response[0].Anulated)
	assert.False(t, response[0].Open)
	assert.False(t, response[0].Violated)
	assert.False(t, response[0].Tamper)
	assert.False(t, response[0].ShortCircuit)
	assert.False(t, response[0].LowBattery)

	assert.True(t, response[9].Anulated)
	assert.True(t, response[9].Open)
	assert.True(t, response[9].Violated)

	assert.True(t, response[10].LowBattery)
	assert.True(t, response[10].Tamper)
	assert.True(t, response[10].ShortCircuit)
}

func Test_Client_parceCentral(t *testing.T) {
	data := []byte{
		0xe9,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x31,
		0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00,
		0x00, 0x00, 0x0f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
		0x40}

	response := parseCentral(data)

	assert.True(t, response.Activated)
	assert.True(t, response.ExternalPowerFault)
	assert.True(t, response.Battery.Low)
	assert.True(t, response.Battery.AbsentOrInverted)
	assert.True(t, response.Battery.ShortCircuit)
	assert.True(t, response.ExternalAuxOverload)
	assert.False(t, response.Siren.Enabled)
	assert.True(t, response.Siren.WireCut)
	assert.True(t, response.Siren.ShortCircuit)
	assert.True(t, response.PhoneLineCut)
	assert.True(t, response.EventCommunicationFailure)
}

func Test_Client_parceCentralZoneAlerting1(t *testing.T) {
	data := []byte{
		0xe9,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x31,
		0x00, 0x00, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
		0x40}

	response := parseCentral(data)

	assert.True(t, response.Alerting)
	assert.False(t, response.Activated)
	assert.False(t, response.ExternalPowerFault)
	assert.False(t, response.Battery.Low)
	assert.False(t, response.Battery.AbsentOrInverted)
	assert.False(t, response.Battery.ShortCircuit)
	assert.False(t, response.ExternalAuxOverload)
	assert.False(t, response.Siren.Enabled)
	assert.False(t, response.Siren.WireCut)
	assert.False(t, response.Siren.ShortCircuit)
	assert.False(t, response.PhoneLineCut)
	assert.False(t, response.EventCommunicationFailure)
}

func Test_Client_parceCentralZoneAlerting2(t *testing.T) {
	// Zone 1: violated false, anulated false, open false, shortCircuit false, lowBattery false, tamper false
	// Zone 10 violated true, anulated true, open true, shortCircuit true, lowBattery true, tamper true
	data := []byte{
		0xe9,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x31,
		0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
		0x40}

	response := parseCentral(data)

	assert.True(t, response.Alerting)
	assert.False(t, response.Activated)
	assert.False(t, response.ExternalPowerFault)
	assert.False(t, response.Battery.Low)
	assert.False(t, response.Battery.AbsentOrInverted)
	assert.False(t, response.Battery.ShortCircuit)
	assert.False(t, response.ExternalAuxOverload)
	assert.False(t, response.Siren.Enabled)
	assert.False(t, response.Siren.WireCut)
	assert.False(t, response.Siren.ShortCircuit)
	assert.False(t, response.PhoneLineCut)
	assert.False(t, response.EventCommunicationFailure)
}

func Test_Client_parceCentralSirenOn(t *testing.T) {
	data := []byte{
		0xe9,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x31,
		0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
		0x40}

	response := parseCentral(data)

	assert.True(t, response.Siren.Enabled)
	assert.False(t, response.Alerting)
	assert.False(t, response.Activated)
	assert.False(t, response.ExternalPowerFault)
	assert.False(t, response.Battery.Low)
	assert.False(t, response.Battery.AbsentOrInverted)
	assert.False(t, response.Battery.ShortCircuit)
	assert.False(t, response.ExternalAuxOverload)
	assert.False(t, response.Siren.WireCut)
	assert.False(t, response.Siren.ShortCircuit)
	assert.False(t, response.PhoneLineCut)
	assert.False(t, response.EventCommunicationFailure)
}

func Test_Client_parceCentralIssueOnCentral(t *testing.T) {
	data := []byte{
		0xe9,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x31,
		0x00, 0x00, 0x11, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
		0x40}

	response := parseCentral(data)

	assert.True(t, response.IssueWarn)
	assert.False(t, response.Siren.Enabled)
	assert.False(t, response.Alerting)
	assert.False(t, response.Activated)
	assert.False(t, response.ExternalPowerFault)
	assert.False(t, response.Battery.Low)
	assert.False(t, response.Battery.AbsentOrInverted)
	assert.False(t, response.Battery.ShortCircuit)
	assert.False(t, response.ExternalAuxOverload)
	assert.False(t, response.Siren.WireCut)
	assert.False(t, response.Siren.ShortCircuit)
	assert.False(t, response.PhoneLineCut)
	assert.False(t, response.EventCommunicationFailure)
}

func Test_Client_parceCentralModelVersion(t *testing.T) {
	data := []byte{
		0xe9,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x31,
		0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
		0x40}

	response := parseCentral(data)

	assert.Equal(t, "AMT2018 E/EG", response.Model)
	assert.Equal(t, "3.1", response.Firmware)
}

func Test_Client_parceCentralDate(t *testing.T) {
	data := []byte{
		0xe9,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x31,
		0x00, 0x00, 0x04, 0x12, 0x12, 0x12, 0x12, 0x12, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
		0x40}

	response := parseDate(data)

	assert.Equal(t, time.Date(2012, time.December, 12, 12, 12, 0, 0, time.Local), response)
}
