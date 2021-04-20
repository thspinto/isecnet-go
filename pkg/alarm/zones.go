package alarm

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ZoneModel is a simple abstraction for Zone status
type ZoneModel struct {
	Id     int
	Name   string
	Status string
}

// ZoneDescription a configuration to add
// contextual metadata to Zones
type ZoneDescription struct {
	Id          int    `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

// GetZones returns Zone statuses in a simple abstraction
func (c *Client) GetZones(ctx context.Context, all bool) ([]ZoneModel, error) {
	status, err := c.GetPartialStatus(ctx)
	if err != nil {
		return nil, err
	}
	var zonesDesc []ZoneDescription
	err = viper.UnmarshalKey("zones", &zonesDesc)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Failed to decode zone descriptions")
	}

	return buildZoneModels(status.Zones, zonesDesc, all), nil
}

func buildZoneModels(zones []Zone, zonesDesc []ZoneDescription, all bool) []ZoneModel {
	zonesDescMap := zonesDescMap(zonesDesc)
	zonesModels := make([]ZoneModel, len(zones))
	for i, z := range zones {
		if zonesDescMap[i+1].Name == "" && !all {
			continue
		}
		zonesModels[i].Id = i + 1
		zonesModels[i].Name = zonesDescMap[i+1].Name

		switch {
		case z.Open:
			zonesModels[i].Status = "Open"
		case z.Anulated:
			zonesModels[i].Status = "Anulated"
		case z.LowBattery:
			zonesModels[i].Status = "LowBattery"
		case z.Violated:
			zonesModels[i].Status = "Violated"
		case z.ShortCircuit:
			zonesModels[i].Status = "ShortCircuit"
		case z.Tamper:
			zonesModels[i].Status = "Tamper"
		}
	}

	return zonesModels
}

func zonesDescMap(zonesDesc []ZoneDescription) map[int]ZoneDescription {
	if zonesDesc == nil {
		return map[int]ZoneDescription{}
	}
	zonesDescMap := make(map[int]ZoneDescription, len(zonesDesc))
	for _, d := range zonesDesc {
		zonesDescMap[d.Id] = d
	}
	return zonesDescMap
}
