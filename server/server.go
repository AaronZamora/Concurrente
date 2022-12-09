package server

import (
	"net/http"
)

type country struct {
	Nombre   string
	Language string
}

var countries []*country = []*country{}

func New(addr string) *http.Server {
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}
