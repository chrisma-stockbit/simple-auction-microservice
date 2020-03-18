package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	gokit "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/infra/gokit"
	pb "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/infra/grpc"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	bidderServiceServerImpl := gokit.NewGrpcServer()
	grpcPort := 10000
	go runGrpcServer(grpcServer, bidderServiceServerImpl, grpcPort)

	httpRoutes := gokit.HttpRoutes()
	httpPort := 8080
	go runHttpServer(httpRoutes, httpPort)

	for {
	}
}

func registerHttpRoutes(routes []gokit.HttpRoute) *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.PathPrefix("/api").Handler(http.StripPrefix("/api", route.Handler)).Methods(route.Method)
	}
	return r
}

func runHttpServer(routes []gokit.HttpRoute, port int) {
	r := registerHttpRoutes(routes)
	log.Printf("running http server on port %d...", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
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
