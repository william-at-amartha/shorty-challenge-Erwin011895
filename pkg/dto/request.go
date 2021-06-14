package dto

type BodyPostShortenURL struct {
	URL string `json:"url"`
	Shortcode string `json:"shortcode"`
}
