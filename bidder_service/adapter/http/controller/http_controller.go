package controller

import (
	"net/http"

	requestdec "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/adapter/http/codec/request"
	responseenc "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/adapter/http/codec/response"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/repo"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/usecase/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Controller struct {
	Routes []httpRoute
}

type httpRoute struct {
	Path    string
	Method  string
	Handler *httptransport.Server
}

func Build(r repo.BidderRepository) Controller {
	var routes []httpRoute
	routes = append(routes, bidderRegistrationRoute(r))
	return Controller{routes}
}

func bidderRegistrationRoute(r repo.BidderRepository) httpRoute {
	handler := httptransport.NewServer(
		endpoint.BidderRegistrationEndpoint(r),
		requestdec.DecodeRegisterBidderRequest,
		responseenc.EncodeBitterRegistrationResponse,
	)
	return httpRoute{"/bidder", http.MethodPost, handler}
}
