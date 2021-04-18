package adapters

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/thspinto/isecnet-go/internal/app/query"
	"github.com/thspinto/isecnet-go/pkg/alarm"
	"github.com/thspinto/isecnet-go/pkg/alarm/mock_alarm"
)

func TestZonesRepository_GetZones(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_alarm.NewMockAlarmClient(ctrl)

	expected := []query.Zone{
		{
			Id:     "1",
			Name:   "one",
			Status: "open",
		},
	}

	m.EXPECT().
		GetZones().
		Return([]alarm.ZoneModel{
			{
				Name:   "one",
				Id:     1,
				Status: "open",
			},
		}, nil)

	c := ZonesRepository{
		client: m,
	}

	z, err := c.GetZones()
	assert.NoError(t, err)
	assert.Equal(t, expected, z)
}

func TestZonesRepository_GetZonesError(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_alarm.NewMockAlarmClient(ctrl)

	m.EXPECT().
		GetZones().
		Return(nil, errors.New("error"))

	c := ZonesRepository{
		client: m,
	}

	_, err := c.GetZones()
	assert.Error(t, err)
}
