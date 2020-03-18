package http

import (
	"context"
	"encoding/json"
	"net/http"
)

func encodeBitterRegistrationResponse(_ context.Context, w http.ResponseWriter, domainResp interface{}) error {
	return json.NewEncoder(w).Encode(domainResp)
}
