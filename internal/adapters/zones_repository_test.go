package adapters

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/thspinto/isecnet-go/internal/domain/zone"
	"github.com/thspinto/isecnet-go/pkg/alarm"
	"github.com/thspinto/isecnet-go/pkg/alarm/mock_alarm"
)

func TestZonesRepository_GetZones(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_alarm.NewMockAlarmClient(ctrl)

	expected := []zone.Zone{
		{
			Id:     "1",
			Name:   "one",
			Status: "open",
		},
	}

	m.EXPECT().
		GetZones(context.Background(), false).
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

	z, err := c.GetZones(context.Background(), false)
	assert.NoError(t, err)
	assert.Equal(t, expected, z)
}

func TestZonesRepository_GetZonesError(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_alarm.NewMockAlarmClient(ctrl)

	m.EXPECT().
		GetZones(context.Background(), false).
		Return(nil, errors.New("error"))

	c := ZonesRepository{
		client: m,
	}

	_, err := c.GetZones(context.Background(), false)
	assert.Error(t, err)
}
