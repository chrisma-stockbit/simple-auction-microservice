package response

import (
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/usecase/dto/request"
)

type BidderRegistrationResponse struct {
	request.BidderRegistrationRequest
	Guid string `json:"guid"`
}
