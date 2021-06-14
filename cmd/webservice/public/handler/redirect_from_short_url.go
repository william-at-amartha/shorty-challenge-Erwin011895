package handler

import (
	"net/http"
	"context"

	"github.com/gorilla/mux"

	internaldto "github.com/Erwin011895/shorty-challenge/internal/dto"
	pkgdto "github.com/Erwin011895/shorty-challenge/pkg/dto"
	"github.com/Erwin011895/shorty-challenge/pkg/constant"
	"github.com/Erwin011895/shorty-challenge/internal/util/httputil"
)

func (h *HandlerComponent) RedirectFromShortURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortcode := vars["shortcode"]

	ctx := r.Context()
	resp := h.redirectFromShortURL(ctx, shortcode)
	httputil.WriteResponseFromHandler(ctx, w, r, resp)
}

func (h *HandlerComponent) redirectFromShortURL(ctx context.Context, shortcode string) *internaldto.ResponseHandler {
	resp := &internaldto.ResponseHandler{}

	url, err := h.shortURLModule.GetURL(ctx, shortcode)
	if err != nil {
		return h.redirectFromShortURLError(err)
	}

	// successfully get the url
	resp.StatusCode = http.StatusFound // 302
	resp.RedirectURL = url
	return resp
}

func (h *HandlerComponent) redirectFromShortURLError(err error) *internaldto.ResponseHandler {
	errorResponses := map[error]*internaldto.ResponseHandler{
		constant.ErrShortcodeNotFound: &internaldto.ResponseHandler{
			StatusCode: http.StatusNotFound, // 404
			Data: pkgdto.ResponseError{
				Tag: constant.ErrShortcodeNotFoundTag,
				Message: constant.ErrShortcodeNotFound.Error(),
			},
		},
	}

	return errorResponses[err]
}
