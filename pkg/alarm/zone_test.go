package alarm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildZoneModelsAll(t *testing.T) {
	zoneDesc := []ZoneDescription{
		{
			Id:   1,
			Name: "one",
		},
		{
			Id:   2,
			Name: "two",
		},
		{
			Id:   4,
			Name: "four",
		},
	}

	zones := []Zone{
		{
			Open: true,
		},
		{
			Anulated: true,
		},
		{
			Violated: true,
		},
		{
			Tamper: true,
		},
		{
			LowBattery: true,
		},
		{
			ShortCircuit: true,
		},
	}

	models := buildZoneModels(zones, zoneDesc, true)
	assert.Equal(t, len(models), 6)

	assert.Equal(t,
		ZoneModel{
			Id:     1,
			Name:   "one",
			Status: "Open",
		},
		models[0])

	assert.Equal(t,
		ZoneModel{
			Id:     2,
			Name:   "two",
			Status: "Anulated",
		},
		models[1])

	assert.Equal(t,
		ZoneModel{
			Id:     3,
			Status: "Violated",
		},
		models[2])

	assert.Equal(t,
		ZoneModel{
			Id:     4,
			Name:   "four",
			Status: "Tamper",
		},
		models[3])

	assert.Equal(t,
		ZoneModel{
			Id:     5,
			Status: "LowBattery",
		},
		models[4])

	assert.Equal(t,
		ZoneModel{
			Id:     6,
			Status: "ShortCircuit",
		},
		models[5])
}

func Test_buildZoneNameModels(t *testing.T) {
	zoneDesc := []ZoneDescription{
		{
			Id:   1,
			Name: "one",
		},
		{
			Id:   2,
			Name: "two",
		},
		{
			Id:   4,
			Name: "four",
		},
	}

	zones := []Zone{
		{
			Open: true,
		},
		{
			Anulated: true,
		},
		{
			Violated: true,
		},
		{
			Tamper: true,
		},
		{
			LowBattery: true,
		},
		{
			ShortCircuit: true,
		},
	}

	models := buildZoneModels(zones, zoneDesc, false)
	assert.Equal(t, len(models), 3)

	assert.Equal(t,
		ZoneModel{
			Id:     1,
			Name:   "one",
			Status: "Open",
		},
		models[0])

	assert.Equal(t,
		ZoneModel{
			Id:     2,
			Name:   "two",
			Status: "Anulated",
		},
		models[1])

	assert.Equal(t,
		ZoneModel{
			Id:     4,
			Name:   "four",
			Status: "Tamper",
		},
		models[2])
}
