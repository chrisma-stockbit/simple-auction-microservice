package test

import (
	"testing"

	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/entity"
	dto "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/usecase/dto/request"
	endp "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/usecase/endpoint"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var MockGuid = "012345678"

type MockBidderRepository struct {
	mock.Mock
}

func (m *MockBidderRepository) Save(b *entity.Bidder) (*entity.Bidder, error) {
	args := m.Called(b)
	bidder := args.Get(0).(*entity.Bidder)
	bidder.Guid = MockGuid
	return bidder, args.Error(1)
}

func TestBidderRegistrationEndpoint(t *testing.T) {
	req := dto.BidderRegistrationRequest{
		FullName:      "Test",
		AccountNumber: "0123456789",
		BankName:      "BTEST",
		PhoneNumber:   "08121234567",
	}
	bidder := entity.NewActiveBidder()
	copier.Copy(bidder, &req)

	mockRepo := new(MockBidderRepository)
	mockRepo.On("Save", bidder).Return(bidder, nil)

	resp, _ := endp.BidderRegistrationEndpoint(mockRepo)(nil, req)
	mockRepo.AssertExpectations(t)

	expected := entity.NewActiveBidder()
	copier.Copy(expected, req)
	expected.Guid = MockGuid
	assert.Equal(t, expected, resp.(*entity.Bidder))
}
