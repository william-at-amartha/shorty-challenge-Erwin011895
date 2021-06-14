package handler

import (
	"net/http"
	"context"

	"github.com/gorilla/mux"

	internaldto "github.com/Erwin011895/shorty-challenge/internal/dto"
	pkgdto "github.com/Erwin011895/shorty-challenge/pkg/dto"
	"github.com/Erwin011895/shorty-challenge/internal/util/httputil"
	"github.com/Erwin011895/shorty-challenge/pkg/constant"
)

func (h *HandlerComponent) GetStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortcode := vars["shortcode"]

	ctx := r.Context()
	resp := h.getStats(ctx, shortcode)
	httputil.WriteResponseFromHandler(ctx, w, r, resp)
}

func (h *HandlerComponent) getStats(ctx context.Context, shortcode string) *internaldto.ResponseHandler {
	resp := &internaldto.ResponseHandler{}

	stats, err := h.shortURLModule.GetStats(ctx, shortcode)
	if err != nil {
		return h.redirectFromShortURLError(err)
	}

	// successfully shorten the url
	resp.StatusCode = http.StatusOK // 200
	resp.Data = stats
	return resp
}

func (h *HandlerComponent) getStatsError(err error) *internaldto.ResponseHandler {
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

