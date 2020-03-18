package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/dto/request"
)

func DecodeRegisterBidderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var domainReq request.BidderRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&domainReq); err != nil {
		return nil, err
	}
	return domainReq, nil
}
