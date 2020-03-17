package service

import (
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/entity"
)

type BidderRegistrationService struct{}

func (b *BidderRegistrationService) RegisterBidder(bidder entity.Bidder) (entity.Bidder, error) {
	//TODO insert to mysql db
	bidder.Guid = "1234567"
	return bidder, nil
}
