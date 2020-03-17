package gokit

import (
	"context"

	domendpoint "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/endpoint"
	encdec "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/grpc"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/service"
	pb "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/infra/grpc"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	registerBidderGrpcHandler grpctransport.Handler
}

func NewGrpcServer() pb.BidderServiceServer {
	bidderRegistrationService := service.BidderRegistrationService{}

	return &grpcServer{
		registerBidderGrpcHandler: grpctransport.NewServer(
			domendpoint.BidderRegistrationEndpoint(bidderRegistrationService),
			encdec.DecodeGrpcRegisterBidderRequest,
			encdec.EncodeGrpcBitterRegistrationResponse,
		),
	}
}

func (s *grpcServer) RegisterBidder(ctx context.Context, req *pb.BidderRegistrationRequest) (*pb.BidderRegistrationResponse, error) {
	_, resp, err := s.registerBidderGrpcHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.BidderRegistrationResponse), nil
}
