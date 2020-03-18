package grpc

import (
	"context"

	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/dto/request"
	pb "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/infra/grpc"
	"github.com/jinzhu/copier"
)

//DecodeGrpcRegisterBidderRequest converts grpc request to user-domain BidderRegistrationRequest
func DecodeGrpcRegisterBidderRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.BidderRegistrationRequest)
	apiReq := request.BidderRegistrationRequest{}
	copier.Copy(&apiReq, &req)
	return apiReq, nil
}
