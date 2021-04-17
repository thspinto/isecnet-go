package app

import (
	"github.com/spf13/viper"
	"github.com/thspinto/isecnet-go/internal/app/handlers"
	"github.com/thspinto/isecnet-go/pkg/client"
)

type Application struct {
	handlers Handlers
}

type Handlers struct {
	GetZones *handlers.GetZonesHandler
}

func NewApplication() Application {
	c := client.NewClient(
		viper.GetString("alarm_host"),
		viper.GetString("alarm_port"),
		viper.GetString("alarm_password"),
	)

	h := Handlers{
		GetZones: handlers.NewGetZonesHandler(c),
	}

	return Application{
		handlers: h,
	}
}
