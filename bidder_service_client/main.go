package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service-client/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/jinzhu/copier"
	grpc "google.golang.org/grpc"
)

func main() {
	serverHost := "localhost"
	serverPort := 10000
	cc, err := grpc.Dial(fmt.Sprintf("%s:%d", serverHost, serverPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to dial: %v", err)
	}
	client := grpctransport.NewClient(cc, "bidder_service.BidderService", "RegisterBidder",
		encodeGrpcRegisterBidderRequest, decodeGrpcSumResponse, pb.BidderRegistrationResponse{})
	domainReq := BidderRegistrationRequest{FullName: "Chrisma Manalu",
		AccountNumber: "11223344556677", BankName: "BCA", PhoneNumber: "08121122334455"}
	resp, err := client.Endpoint()(context.Background(), domainReq)
	if err != nil {
		log.Fatalf("fail to invoke grpc service: %v", err)
	}
	fmt.Printf("response from server: %v\n", resp)
}

type BidderRegistrationRequest struct {
	FullName      string
	AccountNumber string
	BankName      string
	PhoneNumber   string
}

type BidderRegistrationResponse struct {
	BidderRegistrationRequest
	Guid string
}

func encodeGrpcRegisterBidderRequest(_ context.Context, request interface{}) (interface{}, error) {
	domainReq := request.(BidderRegistrationRequest)
	grpcReq := &pb.BidderRegistrationRequest{}
	copier.Copy(grpcReq, domainReq)
	return grpcReq, nil
}

func decodeGrpcSumResponse(_ context.Context, reply interface{}) (interface{}, error) {
	grpcReply := reply.(*pb.BidderRegistrationResponse)
	domainResponse := BidderRegistrationResponse{}
	copier.Copy(&domainResponse, grpcReply)
	copier.Copy(&domainResponse, grpcReply.GetData())
	return domainResponse, nil
}
