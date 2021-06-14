package component

import (
	"github.com/Erwin011895/shorty-challenge/internal/config"

	"github.com/Erwin011895/shorty-challenge/internal/mocks"
)

func InitMockSharedComponent(mc *mocks.MockComponent) *SharedComponent {
	cache := &Cache{
		KodingCache: mc.KodingCache,
	}

	return &SharedComponent{
		Cache: cache,
		Config: &config.Config{
			Environment: "unit-test",
			HttpPort: "8080",
		},
	}
}
