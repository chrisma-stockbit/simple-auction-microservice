package response

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/copier"

	dto "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/usecase/dto/response"
)

func EncodeBitterRegistrationResponse(_ context.Context, w http.ResponseWriter, bidder interface{}) error {
	response := &dto.BidderRegistrationResponse{}
	copier.Copy(response, bidder)

	return json.NewEncoder(w).Encode(response)
}
