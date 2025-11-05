package db

import (
	"context"
	"fmt"

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

	return nil
}
