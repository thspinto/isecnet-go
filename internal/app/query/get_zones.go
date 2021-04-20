package query

import (
	"context"

	"github.com/thspinto/isecnet-go/internal/domain/zone"
)

type GetZonesHandler struct {
	readModel ZonesReadModel
}

func NewGetZonesHandler(readModel ZonesReadModel) *GetZonesHandler {
	if readModel == nil {
		panic("nil zones repository")
	}
	return &GetZonesHandler{
		readModel: readModel,
	}
}

type ZonesReadModel interface {
	GetZones(ctx context.Context, all bool) (zones []zone.Zone, err error)
}

func (h GetZonesHandler) Handle(ctx context.Context, all bool) ([]zone.Zone, error) {
	return h.readModel.GetZones(ctx, all)
}
