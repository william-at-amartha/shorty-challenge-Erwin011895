package component

import (
	"github.com/Erwin011895/shorty-challenge/internal/config"
)

type SharedComponent struct {
	Cache *Cache
	Config *config.Config
}

func InitSharedComponent() *SharedComponent {
	cache := InitializeCache()

	return &SharedComponent{
		Cache: cache,
		Config: config.Get(),
	}
}
