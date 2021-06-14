package model

import (
	"time"
)

type ShortURL struct {
	URL string `json:"url"`
	StartDate time.Time `json:"startDate"`
	LastSeenDate time.Time `json:"lastSeenDate"`
	RedirectCount int `json:"redirectCount"`
}
