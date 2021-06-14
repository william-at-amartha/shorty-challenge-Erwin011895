package router

import (
	"net/http"

	"github.com/Erwin011895/shorty-challenge/cmd/webservice/public/handler"
	"github.com/Erwin011895/shorty-challenge/internal/component"
	"github.com/gorilla/mux"
)

func Init(h *handler.HandlerComponent, sharedComponent *component.SharedComponent) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/ping", h.Ping).Methods("GET")
	router.HandleFunc("/shorten", h.PostShortenURL).Methods("POST")
	router.HandleFunc("/{shortcode}", h.RedirectFromShortURL).Methods("GET")
	router.HandleFunc("/{shortcode}/stats", h.GetStats).Methods("GET")
	
	return router
}
