package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	Client        *mongo.Client
	URLCollection *mongo.Collection
	Ctx           = context.TODO()
)

func Connect(dbURL string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbURL).SetServerAPIOptions(serverAPI)

	var err error

	Client, err = mongo.Connect(opts)
	if err != nil {
		return err
	}

	if err = Client.Ping(Ctx, readpref.Primary()); err != nil {
		return err
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	database := Client.Database("Shorten_service") // <<-- choose your database name
	URLCollection = database.Collection("small")   // <<-- choose your collection name

	if err = createIndexes(); err != nil {
		return fmt.Errorf("failed to create index %w", err)
	}

	return nil
}

func createIndexes() error {
	// Index on original_url for looking up if a long URL already exists
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "original_url", Value: 1}},
		Options: options.Index().SetUnique(false), // Not unique - same URL can't be shortened twice anyway due to our logic
	}

	_, err := URLCollection.Indexes().CreateOne(Ctx, indexModel)
	if err != nil {
		return err
	}

	fmt.Println("Successfully created index on original_url")
	return nil
}
