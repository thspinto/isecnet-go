package ports

import (
	"context"

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

func (g GRPCServer) GetZones(ctx context.Context, in *pb.GetZoneRequest) (*pb.GetZoneResponse, error) {
	zones, err := g.app.Queries.Zones.Handle(ctx, in.All)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var response []*pb.Zone
	for _, r := range zones {
		response = append(response,
			&pb.Zone{
				Id:     r.Id,
				Name:   r.Name,
				Status: r.Status,
			})
	}
	return &pb.GetZoneResponse{Zones: response}, nil
}
