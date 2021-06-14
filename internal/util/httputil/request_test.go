package httputil

import (
	"testing"
	"encoding/json"
	"bytes"

	"github.com/Erwin011895/shorty-challenge/internal/util/testutil"
	"github.com/Erwin011895/shorty-challenge/pkg/dto"
	"github.com/Erwin011895/shorty-challenge/pkg/constant"
)

func TestGetBodyPostShortenURL(t *testing.T) {
    // sub tests
    t.Run("success", func(t *testing.T){
    	wantedBody := &dto.BodyPostShortenURL{
    		URL: "https://example.com",
    		Shortcode: "exampl",
    	}

    	b, _ := json.Marshal(wantedBody)
    	req := testutil.CreateRequest("POST", "/shorten", bytes.NewReader(b))
		body, err := GetBodyPostShortenURL(req)

		if err != nil {
			t.Fatalf(`got %v, %v, want match for %v, %v`, body, err, wantedBody, nil)
		}
    })

    t.Run("invalid body", func(t *testing.T){
    	b := []byte{65}
    	req := testutil.CreateRequest("POST", "/shorten", bytes.NewReader(b))
		_, err := GetBodyPostShortenURL(req)

		wantedErr := constant.ErrBodyInvalid
		if err != wantedErr {
			t.Fatalf(`got %v, %v, want match for %v, %v`, "", err, "", wantedErr)
		}
    })
}
