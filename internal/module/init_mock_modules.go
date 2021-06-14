package module

import (
	"github.com/Erwin011895/shorty-challenge/internal/mocks"
	"github.com/Erwin011895/shorty-challenge/internal/mocks/mockmodule"
)

type MockModule struct {
	ShortURLModule *mockmodule.MockShortURLModuleWrapper
}

func InitMockModule(mc *mocks.MockComponent) *Modules {
	return &Modules{
		ShortURLModule: mc.ShortURLModule,
	}
}
