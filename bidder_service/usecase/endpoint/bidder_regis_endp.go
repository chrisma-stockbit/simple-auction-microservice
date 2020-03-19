package endpoint

import (
	"context"

	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/entity"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/repo"
	dtorequest "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/usecase/dto/request"
	"github.com/go-kit/kit/endpoint"
	"github.com/jinzhu/copier"
)

func BidderRegistrationEndpoint(r repo.BidderRepository) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dtorequest.BidderRegistrationRequest)
		bidder := entity.NewActiveBidder()
		copier.Copy(bidder, &req)

		bidder, err = r.Save(bidder)
		if err != nil {
			return nil, err
		}
		return bidder, nil
	}
}
