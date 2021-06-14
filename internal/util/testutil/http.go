package testutil

import (
	"context"
	"net/http"
	"io"
)

func CreateContext() context.Context {
	req, _ := http.NewRequest("GET", "/", nil)
    return req.Context()
}

func CreateRequest(method, url string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, url, body)
    return req
}
