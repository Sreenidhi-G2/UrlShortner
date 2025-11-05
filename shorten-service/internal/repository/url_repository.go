package repository

import (
	"shorten-service/internal/db"
	"shorten-service/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type URLRepository struct{}

func (r *URLRepository) GetOriginalURL(shortKey string) (string, error) {
	// 1) Check Redis first
	longURL, err := db.RedisClient.Get(db.Ctx, shortKey).Result()
	if err == nil {
		return longURL, nil // cache hit
	}

	// 2) If not found, lookup in Mongo
	var result model.URL
	err = db.URLCollection.FindOne(db.Ctx, bson.M{"_id": shortKey}).Decode(&result)
	if err != nil {
		return "", err
	}

	// 3) Save in Redis (so next lookup is fast)
	db.RedisClient.Set(db.Ctx, shortKey, result.OriginalURL, 24*time.Hour)

	return result.OriginalURL, nil
}
