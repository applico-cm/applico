package main

import (
	"github.com/applico-cm/applico/core"
	"github.com/gin-gonic/gin"
)

type Server struct {
	service core.CoreService
}

func NewServer() *Server {
	s := &Server{
		service: core.NewCoreService(),
	}
	return s
}

func (s Server) ListUsers(ctx *gin.Context) {
	customerID := ctx.Param("customerid")
	_, _ = s.service.FindUsersByCustomerID(customerID)
}
