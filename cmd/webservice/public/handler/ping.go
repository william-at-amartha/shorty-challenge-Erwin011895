package handler

import (
	"net/http"
	"time"

	internaldto "github.com/Erwin011895/shorty-challenge/internal/dto"
	pkgdto "github.com/Erwin011895/shorty-challenge/pkg/dto"
	"github.com/Erwin011895/shorty-challenge/internal/util/httputil"
)

func (h *HandlerComponent) Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	httputil.WriteResponseFromHandler(ctx, w, r, &internaldto.ResponseHandler{
		StatusCode: http.StatusOK,
		Data: pkgdto.ResponsePing{
			Message:         "pong",
			ServerTimestamp: time.Now().Unix(),
			AppName:         "shorty-challenge",
			Environment:     h.config.Environment,
		},
	})
}
