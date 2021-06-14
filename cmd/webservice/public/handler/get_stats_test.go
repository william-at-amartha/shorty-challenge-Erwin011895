package handler

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/Erwin011895/shorty-challenge/internal/mocks"
	"github.com/Erwin011895/shorty-challenge/internal/component"
	"github.com/Erwin011895/shorty-challenge/internal/module"
	"github.com/Erwin011895/shorty-challenge/pkg/constant"
	"github.com/Erwin011895/shorty-challenge/pkg/dto"
)

func TestGetStats(t *testing.T) {
	mc := mocks.InitMockComponent(t)
	sc := component.InitMockSharedComponent(mc)
	mm := module.InitMockModule(mc)
	hc := NewHandler(sc, mm)

	ts := httptest.NewServer(http.HandlerFunc(hc.GetStats))
	defer ts.Close()

	router := mux.NewRouter()
    router.HandleFunc("/{shortcode}/stats", hc.GetStats)

    t.Run("success", func(t *testing.T){
    	shortcode := "example"

    	// expected module calls
    	mc.ShortURLModule.EXPECT().GetStats(gomock.Any(), shortcode).Return(dto.ResponseGetStats{}, nil)

    	// call mock API
    	path := fmt.Sprintf("/%s/stats", shortcode)
    	req, err := http.NewRequest("GET", path, nil)
        if err != nil {
            t.Fatal(err)
        }

        rr := httptest.NewRecorder()
	
        router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
            t.Errorf("handler should give status OK on route %s: got %v want %v",
                shortcode, rr.Code, http.StatusOK)
        }
    })

    t.Run("error not found", func(t *testing.T){
    	shortcode := "example"

    	// expected module calls
    	mc.ShortURLModule.EXPECT().GetStats(gomock.Any(), shortcode).Return(dto.ResponseGetStats{}, constant.ErrShortcodeNotFound)

    	// call mock API
    	path := fmt.Sprintf("/%s/stats", shortcode)
    	req, err := http.NewRequest("GET", path, nil)
        if err != nil {
            t.Fatal(err)
        }

        rr := httptest.NewRecorder()
	
        router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
            t.Errorf("handler should give not found response on route %s: got %v want %v",
                shortcode, rr.Code, http.StatusNotFound)
        }
    })
}
