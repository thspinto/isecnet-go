package adapters

import (
	"strconv"

	"github.com/thspinto/isecnet-go/internal/app/query"
	"github.com/thspinto/isecnet-go/internal/common/errors"
	"github.com/thspinto/isecnet-go/pkg/alarm"
)

type ZonesRepository struct {
	client alarm.AlarmClient
}

func NewZonesRepository(c *alarm.Client) ZonesRepository {
	if c == nil {
		panic("missing alarm client")
	}

	return ZonesRepository{c}
}

func (r *ZonesRepository) GetZones() (zones []query.Zone, err error) {
	response, err := r.client.GetZones()
	if err != nil {
		return nil, errors.NewAlarmCentralError("Failed to get partial status", "alarm-zone-status-get-error")
	}

	zones = make([]query.Zone, len(response))
	for i, r := range response {
		zones[i].Id = strconv.Itoa(r.Id)
		zones[i].Name = r.Name
		zones[i].Status = r.Status
	}

	return
}
