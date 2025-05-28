package model

import (
	"errors"
	"time"
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
	UpdateddAt  time.Time `json:"updated_at"`
}
