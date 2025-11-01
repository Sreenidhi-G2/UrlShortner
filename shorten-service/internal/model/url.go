package model

import (
	"time"
)

type URL struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	OriginalURL string    `bson:"original_url" json:"original_url"`
	ShortCode   string    `bson:"short_code" json:"short_code"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	ExpiresAt   time.Time `bson:"expires_at,omitempty" json:"expires_at,omitempty"`
}
