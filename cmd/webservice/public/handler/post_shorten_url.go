package handler

import (
	"net/http"
	"context"

	internaldto "github.com/Erwin011895/shorty-challenge/internal/dto"
	pkgdto "github.com/Erwin011895/shorty-challenge/pkg/dto"
	"github.com/Erwin011895/shorty-challenge/internal/util/httputil"
	"github.com/Erwin011895/shorty-challenge/pkg/constant"
)

func (h *HandlerComponent) PostShortenURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	bodyPostShortenURL, err := httputil.GetBodyPostShortenURL(r)
	if err != nil {
		httputil.WriteResponseFromHandler(ctx, w, r, &internaldto.ResponseHandler{
			StatusCode: http.StatusBadRequest,
			Data: pkgdto.ResponseError{
				Message: err.Error(),
			},
		})
		return
	}

	resp := h.postShortenURL(ctx, bodyPostShortenURL)
	httputil.WriteResponseFromHandler(ctx, w, r, resp)
}

func (h *HandlerComponent) postShortenURL(ctx context.Context, param *pkgdto.BodyPostShortenURL) *internaldto.ResponseHandler {
	resp := &internaldto.ResponseHandler{}

	shortcode, err := h.shortURLModule.ShortenURL(ctx, param)
	if err != nil {
		return h.postShortenURLError(err)
	}

	// successfully shorten the url
	resp.StatusCode = http.StatusCreated // 201
	resp.Data = pkgdto.ResponsePostShortenURL{
		Shortcode: shortcode,
	}
	return resp
}

func (h *HandlerComponent) postShortenURLError(err error) *internaldto.ResponseHandler {
	errorResponses := map[error]*internaldto.ResponseHandler{
		constant.ErrMissingBodyURL: &internaldto.ResponseHandler{
			StatusCode: http.StatusBadRequest, // 400
			Data: pkgdto.ResponseError{
				Tag: constant.ErrMissingBodyURLTag,
				Message: constant.ErrMissingBodyURL.Error(),
			},
		},
		constant.ErrShortcodeAlreadyInUse: &internaldto.ResponseHandler{
			StatusCode: http.StatusConflict, // 409
			Data: pkgdto.ResponseError{
				Tag: constant.ErrShortcodeAlreadyInUseTag,
				Message: err.Error(),
			},
		},
		constant.ErrShortcodeNotAlphaNumeric: &internaldto.ResponseHandler{
			StatusCode: http.StatusUnprocessableEntity, // 422
			Data: pkgdto.ResponseError{
				Tag: constant.ErrShortcodeNotAlphaNumericTag,
				Message: err.Error(),
			},
		},
	}

	return errorResponses[err]
}

