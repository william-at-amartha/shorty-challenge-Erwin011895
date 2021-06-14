package httputil

import (
	"net/http"
	"encoding/json"

	"github.com/Erwin011895/shorty-challenge/pkg/constant"
	"github.com/Erwin011895/shorty-challenge/pkg/dto"
)

func GetBodyPostShortenURL(r *http.Request) (bodyPostShortenURL *dto.BodyPostShortenURL, err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&bodyPostShortenURL)
	if err != nil {
		return nil, constant.ErrBodyInvalid
	}
	return
}
