package httputil

import (
	"context"
	"net/http"
	"encoding/json"

	internaldto "github.com/Erwin011895/shorty-challenge/internal/dto"
	pkgdto "github.com/Erwin011895/shorty-challenge/pkg/dto"
)

var redirectHTTPStatuses = map[int]bool{
	http.StatusMultipleChoices: true, // 300
    http.StatusMovedPermanently: true, // 301
    http.StatusFound: true, // 302
    http.StatusSeeOther: true, // 303
    http.StatusNotModified: true, // 304
    http.StatusUseProxy: true, // 305
}

func WriteResponseFromHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, resp *internaldto.ResponseHandler) {
	isRedirect := redirectHTTPStatuses[resp.StatusCode]

	if isRedirect {
		http.Redirect(w, r, resp.RedirectURL, resp.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(resp.StatusCode)
	err := json.NewEncoder(w).Encode(&resp.Data)
	if err != nil {
		json.NewEncoder(w).Encode(&pkgdto.ResponseError{
			Tag: "ErrorEncodeToJSON",
			Message: "failed to encode response",
		})
	}
}
