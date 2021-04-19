package ports

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/thspinto/isecnet-go/api/pb"
	"github.com/thspinto/isecnet-go/internal/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	app app.Application
	pb.UnimplementedZoneServiceServer
}

func NewGrpcServer(application app.Application) GRPCServer {
	return GRPCServer{app: application}
}

func (g GRPCServer) GetZones(ctx context.Context, in *empty.Empty) (*pb.ZoneResponse, error) {
	zones, err := g.app.Queries.Zones.Handle(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := make([]*pb.Zone, len(zones))

	for i, r := range zones {
		response[i].Id = r.Id
		response[i].Name = r.Name
		response[i].Status = r.Status
	}
	return &pb.ZoneResponse{Zones: response}, nil
}
