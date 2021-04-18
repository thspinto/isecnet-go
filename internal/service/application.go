package service

import (
	"context"

	"github.com/spf13/viper"
	"github.com/thspinto/isecnet-go/internal/adapters"
	"github.com/thspinto/isecnet-go/internal/app"
	"github.com/thspinto/isecnet-go/internal/app/query"
	"github.com/thspinto/isecnet-go/pkg/alarm"
)

func NewApplication(ctx context.Context) app.Application {
	c := alarm.NewClient(
		viper.GetString("alarm_host"),
		viper.GetString("alarm_port"),
		viper.GetString("alarm_password"),
	)

	r := adapters.NewZonesRepository(c)

	q := app.Queries{
		Zones: query.NewGetZonesHandler(r),
	}

	return app.Application{
		Queries: q,
	}
}
