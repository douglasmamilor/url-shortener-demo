package model

import "time"

// URL model
type URL struct {
	ID          string    `db:"id" json:"_"`
	OriginalURL string    `db:"original_url" json:"original_url"`
	ShortCode   string    `db:"short_code" json:"short_code"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

// ShortenURLRequest represents URL shortening request payloads
type ShortenURLRequest struct {
	URL string `json:"url"`
}
