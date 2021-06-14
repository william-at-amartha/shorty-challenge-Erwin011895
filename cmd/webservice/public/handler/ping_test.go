package handler

import (
	"testing"
	"net/http"
	"net/http/httptest"
	// "io"

	"github.com/Erwin011895/shorty-challenge/internal/mocks"
	"github.com/Erwin011895/shorty-challenge/internal/component"
	"github.com/Erwin011895/shorty-challenge/internal/module"
)

func TestPing(t *testing.T) {
    t.Run("success", func(t *testing.T){
    	mc := mocks.InitMockComponent(t)
    	sc := component.InitMockSharedComponent(mc)
    	mm := module.InitMockModule(mc)
    	hc := NewHandler(sc, mm)

    	ts := httptest.NewServer(http.HandlerFunc(hc.Ping))
		defer ts.Close()

		// res, err := http.Get(ts.URL)
		_, err := http.Get(ts.URL)
		if err != nil {
			t.Fatalf("[Ping] error on API call. %+v", err)
		}

		// byteBodyResp, err = io.ReadAll(res.Body)
		// res.Body.Close()
		// if err != nil {
		// 	t.Fatalf("[Ping] error on read response body. %+v", err)
		// }
    })
}
