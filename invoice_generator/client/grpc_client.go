package client

import (
	"github.com/kevalsabhani/toll-calculator/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InvoiceGeneratorGrpcClient struct {
}

func NewInvoiceGeneratorGrpcClient(grpcServerHost string) (pb.InvoiceServiceClient, error) {
	conn, err := grpc.Dial(grpcServerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewInvoiceServiceClient(conn)
	return client, nil
}
