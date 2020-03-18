package http

import (
	"context"
	"encoding/json"
	"net/http"
)

func EncodeBitterRegistrationResponse(_ context.Context, w http.ResponseWriter, domainResp interface{}) error {
	return json.NewEncoder(w).Encode(domainResp)
}
