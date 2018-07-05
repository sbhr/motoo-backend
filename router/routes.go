package router

import (
	"net/http"

	"github.com/sbhr/motoo-backend/handler"
)

type Route struct {
	handler http.HandlerFunc
	path    string
	methods []string
}

var (
	routes = []Route{
		{
			handler: handler.Get,
			path:    "/",
			methods: []string{
				http.MethodGet,
			},
		},
	}
)
