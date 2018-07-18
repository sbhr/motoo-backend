package router

import (
	"github.com/gorilla/mux"

	"github.com/sbhr/motoo-backend/handler"
)

func New(h handler.Handler) *mux.Router {
	router := mux.NewRouter()

	for _, r := range NewRoutes(h) {
		router.Methods(r.methods...).Path(r.path).Handler(r.handler)
	}

	return router
}
