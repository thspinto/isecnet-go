package app

import (
	"github.com/thspinto/isecnet-go/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	Zones *query.GetZonesHandler
}
