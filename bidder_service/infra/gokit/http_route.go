package gokit

import (
	"net/http"

	httpadapter "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/adapter/http"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/api/endpoint"
	mysqlrepo "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/repo/mysql"
	"github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/domain/service"
	mysqldb "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/infra/mysql"
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
	bidderRepo := &mysqlrepo.BidderRepository{}
	bidderRepo.Conn = mysqldb.GetDbConnection
	service.BidderRepo = bidderRepo
	handler := httptransport.NewServer(
		endpoint.BidderRegistrationEndpoint(service),
		httpadapter.DecodeRegisterBidderRequest,
		httpadapter.EncodeBitterRegistrationResponse,
	)
	return HttpRoute{"/bidder", http.MethodPost, handler}
}
