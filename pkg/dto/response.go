package dto

import (
	"time"
)

type ResponsePing struct {
	AppName         string `json:"app_name"`
	Environment     string `json:"environment"`
	Message         string `json:"message"`
	ServerTimestamp int64  `json:"server_timestamp"`
}

type ResponseError struct {
	Tag string `json:"tag,omitempty"`
	Message string `json:"message"`
}

type ResponsePostShortenURL struct {
	Shortcode string `json:"shortcode"`
}

type ResponseGetStats struct {
	StartDate time.Time `json:"startDate"`
	LastSeenDate time.Time `json:"lastSeenDate"`
	RedirectCount int `json:"redirectCount"`
}
