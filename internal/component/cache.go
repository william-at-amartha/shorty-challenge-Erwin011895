package component

import (
	"github.com/koding/cache"
)

// InMemory component
type Cache struct {
	KodingCache cache.Cache
}

// InitializeInMemory to initialize cache
func InitializeCache() *Cache {
	return &Cache{
		KodingCache: cache.NewMemory(),
	}
}
