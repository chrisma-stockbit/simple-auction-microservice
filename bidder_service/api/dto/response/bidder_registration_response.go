package response

import (
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/dto/request"
)

type BidderRegistrationResponse struct {
	request.BidderRegistrationRequest
	Balance int64  `json:"balance,omitempty"`
	Guid    string `json:"guid"`
}
