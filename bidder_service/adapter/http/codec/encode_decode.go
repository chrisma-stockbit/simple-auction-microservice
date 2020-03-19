package http

import (
	"context"
	"net/http"
)

// DecodeRequestFunc decodes HTTP request to domain layer request
type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)

// EncodeResponseFunc encodes response from domain layer to HTTP response
type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error
