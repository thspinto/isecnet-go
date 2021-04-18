package query

import (
	"github.com/thspinto/isecnet-go/pkg/alarm"
)

type ZonesHandler struct {
	client *alarm.Client
}

type Zone struct {
	Id     string
	Name   string
	Status string
}

func NewZonesHandler(c *alarm.Client) *ZonesHandler {
	if c == nil {
		panic("nil zones repository")
	}
	return &ZonesHandler{
		client: c,
	}
}
