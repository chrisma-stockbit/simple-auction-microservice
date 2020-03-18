package main

import (
	"fmt"
	"log"
	"net"

	gokit "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/infra/gokit"
	pb "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/infra/grpc"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	bidderServiceServerImpl := gokit.NewGrpcServer()
	grpcPort := 10000
	setupGrpcServer(grpcServer, bidderServiceServerImpl, grpcPort)

}

func runGrpcServer(grpcServer *grpc.Server, bidderServiceServerImpl pb.BidderServiceServer, port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to run grpc server: %v", err)
	}

	//We pass proto method - gokit endpoint integrator as the grpc server implementation
	pb.RegisterBidderServiceServer(grpcServer, bidderServiceServerImpl)
	log.Printf("serving grpc server on port %d..", port)
	grpcServer.Serve(ln)
}
