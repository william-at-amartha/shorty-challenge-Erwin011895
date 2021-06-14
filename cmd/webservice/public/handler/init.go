package handler

import (
	"net/http"
	
	"github.com/Erwin011895/shorty-challenge/internal/component"
	"github.com/Erwin011895/shorty-challenge/internal/config"
	"github.com/Erwin011895/shorty-challenge/internal/module"
)

type HandlerComponent struct {
	cache *component.Cache
	config *config.Config
	shortURLModule module.ShortURLModuleWrapper
}

type Handler interface {
	Ping(w http.ResponseWriter, r *http.Request)
	PostShortenURL(w http.ResponseWriter, r *http.Request)
	RedirectFromShortURL(w http.ResponseWriter, r *http.Request)
	GetStats(w http.ResponseWriter, r *http.Request)
}

func NewHandler(sharedComponent *component.SharedComponent, modules *module.Modules) *HandlerComponent {
	return &HandlerComponent{
		config: sharedComponent.Config,
		cache : sharedComponent.Cache,
		shortURLModule: modules.ShortURLModule,
	}
}
