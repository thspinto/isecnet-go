package handlers

import "github.com/thspinto/isecnet-go/pkg/client"

type GetZonesHandler struct {
	client *client.Client
}

func NewGetZonesHandler(c *client.Client) *GetZonesHandler {
	return &GetZonesHandler{
		client: c,
	}
}
