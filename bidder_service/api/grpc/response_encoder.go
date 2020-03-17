package grpc

import (
	"context"

	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/dto/response"
	pb "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/infra/grpc"
	"github.com/jinzhu/copier"
)

func EncodeGrpcBitterRegistrationResponse(_ context.Context, domainResp interface{}) (interface{}, error) {
	resp := domainResp.(response.BidderRegistrationResponse)
	grpcResp := pb.BidderRegistrationResponse{}
	grpcRespData := pb.BidderRegistrationRequest{}

	copier.Copy(&grpcRespData, &resp)
	grpcResp.Guid = resp.Guid
	grpcResp.Data = &grpcRespData
	return &grpcResp, nil
}
