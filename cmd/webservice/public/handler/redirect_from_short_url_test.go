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
)

func TestRedirectFromShortURL(t *testing.T) {
	mc := mocks.InitMockComponent(t)
	sc := component.InitMockSharedComponent(mc)
	mm := module.InitMockModule(mc)
	hc := NewHandler(sc, mm)

	ts := httptest.NewServer(http.HandlerFunc(hc.RedirectFromShortURL))
	defer ts.Close()

	router := mux.NewRouter()
    router.HandleFunc("/{shortcode}", hc.RedirectFromShortURL)

    t.Run("success", func(t *testing.T){
    	shortcode := "example"
    	url := "https://example.com"

    	// expected module calls
    	mc.ShortURLModule.EXPECT().GetURL(gomock.Any(), shortcode).Return(url, nil)

    	// call mock API
    	path := fmt.Sprintf("/%s", shortcode)
    	req, err := http.NewRequest("GET", path, nil)
        if err != nil {
            t.Fatal(err)
        }

        rr := httptest.NewRecorder()
	
        router.ServeHTTP(rr, req)

		if rr.Code != http.StatusFound {
            t.Errorf("handler should redirect on route %s: got %v want %v",
                shortcode, rr.Code, http.StatusFound)
        }
    })

    t.Run("error not found", func(t *testing.T){
    	shortcode := "example"
    	url := "https://example.com"

    	// expected module calls
    	mc.ShortURLModule.EXPECT().GetURL(gomock.Any(), shortcode).Return(url, constant.ErrShortcodeNotFound)

    	// call mock API
    	path := fmt.Sprintf("/%s", shortcode)
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
