package app

import (
	"github.com/spf13/viper"
	"github.com/thspinto/isecnet-go/internal/app/query"
	"github.com/thspinto/isecnet-go/pkg/alarm"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	Zones *query.ZonesHandler
}

func NewApplication() Application {
	c := alarm.NewClient(
		viper.GetString("alarm_host"),
		viper.GetString("alarm_port"),
		viper.GetString("alarm_password"),
	)

	q := Queries{
		Zones: query.NewZonesHandler(c),
	}

	return Application{
		Queries: q,
	}
}
