package alarm

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ZoneModel struct {
	Id     int
	Name   string
	Status string
}

type ZoneDescription struct {
	Id          int    `yaml:"id"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

func (c *Client) GetZones() ([]ZoneModel, error) {
	status, err := c.GetPartialStatus()
	if err != nil {
		return nil, err
	}
	var zonesDesc []ZoneDescription
	err = viper.UnmarshalKey("zones", &zonesDesc)
	log.WithFields(log.Fields{
		"error": err,
	}).Error("Failed to decode zone descriptions")

	return buildZoneModels(status.Zones, zonesDesc), nil
}

func buildZoneModels(zones []Zone, zonesDesc []ZoneDescription) []ZoneModel {
	zonesDescMap := zonesDescMap(zonesDesc)
	zonesModels := make([]ZoneModel, len(zones))
	for i, z := range zones {
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
