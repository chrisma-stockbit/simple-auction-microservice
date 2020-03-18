package endpoint

import (
	"context"
	"log"

	dtorequest "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/dto/request"
	dtoresponse "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/dto/response"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/entity"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/jinzhu/copier"
)

func BidderRegistrationEndpoint(s *service.BidderRegistrationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dtorequest.BidderRegistrationRequest)
		bidder := entity.Bidder{}
		copier.Copy(&bidder, &req)

		bidder, err = s.RegisterBidder(bidder)
		if err != nil {
			log.Fatalf("register bidder failed: %v", err)
		}

		resp := dtoresponse.BidderRegistrationResponse{}
		copier.Copy(&resp, &bidder)
		return resp, nil
	}
}
