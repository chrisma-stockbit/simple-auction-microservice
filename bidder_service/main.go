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
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", 10000))
	if err != nil {
		log.Fatalf("failed to run grpc server: %v", err)
	}
	grpcServer := grpc.NewServer()
	//We pass proto method - gokit endpoint integrator as the grpc server implementation
	pb.RegisterBidderServiceServer(grpcServer, gokit.NewGrpcServer())
	log.Println("serving grpc server..")
	grpcServer.Serve(ln)
}
