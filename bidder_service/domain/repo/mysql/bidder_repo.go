package mysql

import (
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/entity"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/repo"
	"github.com/jinzhu/gorm"
)

type BidderRepository struct {
	repo.BidderRepository
	Conn func() (*gorm.DB, error)
}

func (r *BidderRepository) Save(b *entity.Bidder) (*entity.Bidder, error) {
	db, err := r.Conn()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Create(b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}
