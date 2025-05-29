package model

import (
	"errors"
	"time"

	"github.com/samber/mo"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrIdentifierExists = errors.New("identifier already exists")
)

type Shortening struct {
	Identifier  string    `json:"identifier"` // The short version for the URL
	OriginalURL string    `json:"original_url"`
	Visits      int64     `json:"visits"` // The number of times the shortened URL has been visited
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// This struct is designed to capture client input (e.g., from an HTTP request)
// and validate it before creating a Shortening struct,
// which is the actual data model stored in MongoDB
type ShortenInput struct {
	RawURL     string
	Identifier mo.Option[string]
	CreatedBy  string
}
