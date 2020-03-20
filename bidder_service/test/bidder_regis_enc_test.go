package test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/copier"

	responseencoder "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/adapter/http/codec/response"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/entity"
	dto "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/usecase/dto/response"
	"github.com/stretchr/testify/assert"
)

func TestEncodeBitterRegistrationResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	bidder := entity.Bidder{
		FullName:      "TestName",
		AccountNumber: "123",
		BankName:      "BTEST",
		PhoneNumber:   "08121234567",
		Guid:          "0102030405",
	}

	expected := dto.BidderRegistrationResponse{}
	copier.Copy(&expected, bidder)

	responseencoder.EncodeBitterRegistrationResponse(nil, rr, bidder)
	actual := dto.BidderRegistrationResponse{}
	json.NewDecoder(rr.Body).Decode(&actual)
	assert.Equal(t, expected, actual)
}
