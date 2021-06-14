package handler

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io"
	"encoding/json"
	"bytes"

	"github.com/golang/mock/gomock"

	"github.com/Erwin011895/shorty-challenge/internal/mocks"
	"github.com/Erwin011895/shorty-challenge/internal/component"
	"github.com/Erwin011895/shorty-challenge/internal/module"
	"github.com/Erwin011895/shorty-challenge/pkg/dto"
	"github.com/Erwin011895/shorty-challenge/pkg/constant"
)

func TestPostShortenURL(t *testing.T) {
	mc := mocks.InitMockComponent(t)
	sc := component.InitMockSharedComponent(mc)
	mm := module.InitMockModule(mc)
	hc := NewHandler(sc, mm)

	ts := httptest.NewServer(http.HandlerFunc(hc.PostShortenURL))
	defer ts.Close()

    t.Run("success", func(t *testing.T){
		reqBody := &dto.BodyPostShortenURL{
    		URL: "https://example.com",
    		Shortcode: "exampl",
    	}
    	b, _ := json.Marshal(reqBody)

    	// expected module calls
    	mc.ShortURLModule.EXPECT().ShortenURL(gomock.Any(), reqBody).Return(reqBody.Shortcode, nil)

    	// call mock API
		res, err := http.Post(ts.URL, "", bytes.NewReader(b))
		if err != nil {
			t.Fatalf("[TestPostShortenURL] error on API call. %+v", err)
		}

		if res.StatusCode != http.StatusCreated {
			t.Fatalf("[TestPostShortenURL] got status code %v, want %v.", res.StatusCode, http.StatusCreated)
		}

		_, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			t.Fatalf("[TestPostShortenURL] error on read response body. %+v", err)
		}
    })

    t.Run("error shortcode already in use", func(t *testing.T){
		reqBody := &dto.BodyPostShortenURL{
    		URL: "https://example.com",
    		Shortcode: "exampl",
    	}
    	b, _ := json.Marshal(reqBody)

    	// expected module calls
    	mc.ShortURLModule.EXPECT().ShortenURL(gomock.Any(), reqBody).Return(reqBody.Shortcode, constant.ErrShortcodeAlreadyInUse)

    	// call mock API
		res, err := http.Post(ts.URL, "", bytes.NewReader(b))
		if err != nil {
			t.Fatalf("[TestPostShortenURL] error on API call. %+v", err)
		}

		if res.StatusCode != http.StatusConflict {
			t.Fatalf("[TestPostShortenURL] got status code %v, want %v.", res.StatusCode, http.StatusCreated)
		}
    })
}
