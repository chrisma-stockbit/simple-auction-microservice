package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"

	controller "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/adapter/http/controller"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/entity"
	infrahttp "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/infra/http"
	dtoreq "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/usecase/dto/request"
	dtoresp "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/usecase/dto/response"
)

func Test(t *testing.T) {
	bidder := entity.NewActiveBidder()
	regisReq := dtoreq.BidderRegistrationRequest{
		FullName:      "Test Name",
		AccountNumber: "123",
		BankName:      "BTEST",
		PhoneNumber:   "08121234567",
	}
	copier.Copy(bidder, &regisReq)

	mockRepo := new(MockBidderRepository)
	mockRepo.On("Save", bidder).Return(bidder, nil)
	httpController := controller.Build(mockRepo)
	httpserver := infrahttp.NewServer(&httpController, 8080)

	rr := httptest.NewRecorder()
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(regisReq)
	req, _ := http.NewRequest("POST", "/api/bidder", reqBodyBytes)
	httpserver.Router.ServeHTTP(rr, req)

	expected := dtoresp.BidderRegistrationResponse{regisReq, MockGuid}
	expected.Guid = MockGuid
	actual := dtoresp.BidderRegistrationResponse{}
	json.NewDecoder(rr.Body).Decode(&actual)
	assert.Equal(t, expected, actual)
}
