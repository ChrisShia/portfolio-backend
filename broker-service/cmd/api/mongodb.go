package main

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (app *App) connectMongoDB() {
	uri := app.config.mongo.URI

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		app.log.PrintFatal(
			err, map[string]string{"message": "Failed to connect to MongoDB"})
	}

	app.log.PrintInfo("MongoDB", map[string]string{
		"status": "Connected!",
		"uri":    uri,
	})

	app.mongo.client = client
}

func (app *App) disconnectMongoDB() {
	if err := app.mongo.client.Disconnect(context.TODO()); err != nil {
		app.log.PrintFatal(err, nil)
	}
}
