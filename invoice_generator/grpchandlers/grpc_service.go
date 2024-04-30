package grpchandlers

import (
	"context"

	"github.com/kevalsabhani/toll-calculator/invoice_generator/storage"
	"github.com/kevalsabhani/toll-calculator/invoice_generator/types"
	"github.com/kevalsabhani/toll-calculator/pb"
)

type GrpcInvoiceServiceServer struct {
	pb.UnimplementedInvoiceServiceServer
	store storage.Store
}

func NewGrpcInvoiceServiceServer(store storage.Store) *GrpcInvoiceServiceServer {
	return &GrpcInvoiceServiceServer{
		store: store,
	}
}

func (gis *GrpcInvoiceServiceServer) AggregateDistance(
	ctx context.Context,
	in *pb.AggregateDistanceRequest,
) (*pb.AggregateDistanceResponse, error) {
	distanceData := &types.Distance{
		Value:     in.GetValue(),
		OBUId:     int(in.GetOBUId()),
		Timestamp: in.GetTimestamp(),
	}
	gis.store.Insert(distanceData)
	return &pb.AggregateDistanceResponse{Message: "Distance added successfully"}, nil
}

func (gis *GrpcInvoiceServiceServer) GenerateInvoice(
	ctx context.Context,
	in *pb.GenerateInvoiceRequest,
) (*pb.GenerateInvoiceResponse, error) {
	obuId := in.GetOBUId()
	totalDistance, err := gis.store.Get(int(obuId))
	if err != nil {
		return nil, err
	}
	return &pb.GenerateInvoiceResponse{
		OBUId:         obuId,
		TotalDistance: totalDistance,
		Amount:        totalDistance * 10.0,
	}, nil
}
