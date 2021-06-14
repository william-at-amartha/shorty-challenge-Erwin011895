package module

import (
	"github.com/Erwin011895/shorty-challenge/internal/component"
)

type Modules struct {
	ShortURLModule ShortURLModuleWrapper
}

func InitModules(sharedComponent *component.SharedComponent) *Modules {
	// init modules
	shortURLModule := NewShortURLModule(&ShortURLModuleParams{
		Cache: sharedComponent.Cache,
	})

	return &Modules{
		ShortURLModule: shortURLModule,
	}
}
