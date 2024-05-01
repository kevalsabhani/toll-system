package main

import (
	"flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kevalsabhani/toll-calculator/invoice_generator/grpchandlers"
	"github.com/kevalsabhani/toll-calculator/invoice_generator/handlers"
	"github.com/kevalsabhani/toll-calculator/invoice_generator/storage"
	"github.com/kevalsabhani/toll-calculator/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	httpPort = ":3000"
	grpcPort = ":50051"
	env      = "DEVELOPMENT"
)

var logger *zap.Logger

func main() {
	transportPtr := flag.String("t", "http", "select the transport[http | grcp]")
	flag.Parse()

	//zap logger
	logger = zap.Must(zap.NewProduction())
	if env == "DEVELOPMENT" {
		logger = zap.Must(zap.NewDevelopment())
	}

	store := storage.NewMemoryStore(logger)
	aggregator := handlers.NewDistanceAggregator(store, logger)
	generator := handlers.NewInvoiceGenerator(store, logger)

	// Create a signal channel to handle graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	if *transportPtr == "http" {
		// Http Transport
		go makeHttpTransport(aggregator, generator)
	} else {
		// Grpc Transport
		go makeGrpcTransport(store)
	}

	<-sigchan
	logger.Info("Closing the server...")
}

func makeHttpTransport(aggregator *handlers.DistanceAggregator, generator *handlers.InvoiceGenerator) {
	logger.Info("Http transport running on", zap.String("port", httpPort))
	http.HandleFunc("/aggregate_distance", aggregator.AggregateDistance)
	http.HandleFunc("/invoice", generator.GenerateInvoice)
	http.ListenAndServe(httpPort, nil)
}

func makeGrpcTransport(store storage.Store) {
	logger.Info("Grpc transport running on", zap.String("port", grpcPort))
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Fatal(err.Error())
	}
	grpcServer := grpc.NewServer()
	grpcInvoiceServiceServer := grpchandlers.NewGrpcInvoiceServiceServer(store)
	pb.RegisterInvoiceServiceServer(grpcServer, grpcInvoiceServiceServer)
	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal(err.Error())
	}
}
