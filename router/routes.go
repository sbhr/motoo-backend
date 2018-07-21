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
			handler: h.GetAllConversations,
			path:    "/convo/",
			methods: []string{
				http.MethodGet,
			},
		},
		{
			handler: h.GetConversation,
			path:    "/convo/{id}",
			methods: []string{
				http.MethodGet,
			},
		},
	}
}
