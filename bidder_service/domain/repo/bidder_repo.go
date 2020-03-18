package repo

import (
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/entity"
)

type BidderRepository interface {
	Save(entity.Bidder) (entity.Bidder, error)
}
