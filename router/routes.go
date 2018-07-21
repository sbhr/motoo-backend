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
			handler: h.GetConversations,
			path:    "/convo/",
			methods: []string{
				http.MethodGet,
			},
		},
		{
			handler: h.GetConversationByID,
			path:    "/convo/{id}",
			methods: []string{
				http.MethodGet,
			},
		},
		{
			handler: h.PostConversation,
			path:    "/convo/",
			methods: []string{
				http.MethodPost,
			},
		},
		{
			handler: h.UpdateConversation,
			path:    "/convo/{id}",
			methods: []string{
				http.MethodPut,
			},
		},
		{
			handler: h.DeleteConversation,
			path:    "/convo/{id}",
			methods: []string{
				http.MethodDelete,
			},
		},
		{
			handler: h.PostUser,
			path:    "/user/",
			methods: []string{
				http.MethodPost,
			},
		},
		{
			handler: h.PostPlaylog,
			path:    "/playlog/",
			methods: []string{
				http.MethodPost,
			},
		},
	}
}
