package response

import (
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/dto/request"
)

type BidderRegistrationResponse struct {
	request.BidderRegistrationRequest
	Guid string `json:"guid"`
}
