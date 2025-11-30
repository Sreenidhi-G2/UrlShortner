package repository

import (
	"shorten-service/internal/db"
	"shorten-service/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type URLRepository struct{}

// GetOriginalURL returns the long/original URL for a given short key.
func (r *URLRepository) GetOriginalURL(shortKey string) (string, error) {

	// Directly check MongoDB (Redis removed)
	var result model.URL
	err := db.URLCollection.FindOne(db.Ctx, bson.M{"_id": shortKey}).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.OriginalURL, nil
}

// SaveShortURL stores a new shortKey → longURL mapping.
func (r *URLRepository) SaveShortURL(shortKey, longURL string) error {
	doc := model.URL{
		ID:          shortKey,
		OriginalURL: longURL,
		ShortCode:   shortKey,
		CreatedAt:   time.Now(),
	}

	_, err := db.URLCollection.InsertOne(db.Ctx, doc)
	return err
}

// GetShortKeyByLongURL returns the short code for the given long URL.
func (r *URLRepository) GetShortKeyByLongURL(longURL string) (string, error) {

	// Directly query MongoDB (Redis removed)
	var result model.URL
	err := db.URLCollection.FindOne(db.Ctx, bson.M{"original_url": longURL}).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.ShortCode, nil
}
