package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	httpadapter "github.com/chrisma-stockbit/simple-auction-microservice/bidder-service/adapter/http/controller"
)

type Server struct {
	port int
	r    *mux.Router
}

func NewServer(c *httpadapter.Controller, port int) *Server {
	s := &Server{port: port}
	s.registerRoutes(c)
	return s
}

func (s *Server) registerRoutes(c *httpadapter.Controller) {
	s.r = mux.NewRouter()
	for _, route := range c.Routes {
		s.r.PathPrefix("/api" + route.Path).Handler(route.Handler).Methods(route.Method)
	}
}

func (s *Server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.r)
}
