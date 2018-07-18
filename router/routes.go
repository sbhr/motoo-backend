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

func NewRoutes(h handler.Handler) []Route {
	return []Route{
		{
			handler: h.GetAll,
			path:    "/",
			methods: []string{
				http.MethodGet,
			},
		},
	}
}
