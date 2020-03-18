package gokit

import (
	"net/http"

	httpadapter "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/adapter/http"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/endpoint"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/service"
	httptransport "github.com/go-kit/kit/transport/http"
)

type HttpRoute struct {
	Path    string
	Method  string
	Handler *httptransport.Server
}

func HttpRoutes() []HttpRoute {
	var httpRoute []HttpRoute
	httpRoute = append(httpRoute, bidderRegistrationRoute())
	return httpRoute
}

func bidderRegistrationRoute() HttpRoute {
	service := new(service.BidderRegistrationService)
	handler := httptransport.NewServer(
		endpoint.BidderRegistrationEndpoint(service),
		httpadapter.DecodeRegisterBidderRequest,
		httpadapter.EncodeBitterRegistrationResponse,
	)
	return HttpRoute{"/bidder", http.MethodPost, handler}
}
