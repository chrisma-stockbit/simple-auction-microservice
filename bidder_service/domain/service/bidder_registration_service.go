package service

import (
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/entity"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/repo/mysql"
)

type BidderRegistrationService struct {
	BidderRepo *mysql.BidderRepository
}

//RegisterBidder saves new bidder that active.
func (b *BidderRegistrationService) RegisterBidder(bidder *entity.Bidder) (*entity.Bidder, error) {
	bidder.IsActive = true
	bidder, err := b.BidderRepo.Save(bidder)
	if err != nil {
		return nil, err
	}
	return bidder, nil
}
