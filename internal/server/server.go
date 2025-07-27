package server

import (
	"github.com/StEvseeva/cleany/internal/service"
)

type Server struct {
	service service.Service
}

func NewServer(service service.Service) *Server {
	return &Server{
		service: service,
	}
}
