package request

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/usecase/dto/request"
)

func DecodeRegisterBidderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var domainReq request.BidderRegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&domainReq)
	if err != nil {
		return nil, err
	}
	return domainReq, nil
}
