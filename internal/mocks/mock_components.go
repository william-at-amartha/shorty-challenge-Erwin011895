package mocks

import (
	"testing"
	
	"github.com/golang/mock/gomock"

	"github.com/Erwin011895/shorty-challenge/internal/mocks/mockcache"
	"github.com/Erwin011895/shorty-challenge/internal/mocks/mockmodule"
)

type MockComponent struct {
	Controller *gomock.Controller
	KodingCache *mockcache.MockCache
	ShortURLModule *mockmodule.MockShortURLModuleWrapper
}

func InitMockComponent(t *testing.T) *MockComponent {
	mockCtrl := gomock.NewController(t)

	return &MockComponent{
		Controller: mockCtrl,
		KodingCache: mockcache.NewMockCache(mockCtrl),
		ShortURLModule: mockmodule.NewMockShortURLModuleWrapper(mockCtrl),
	}
}
