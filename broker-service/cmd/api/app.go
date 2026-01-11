package main

import (
	"portfolio/internal/data"

	"github.com/ChrisShia/jsonlog"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type config struct {
	port  int
	mongo struct {
		URI string
	}
}

type App struct {
	config config
	log    *jsonlog.Logger
	mongo  struct {
		client *mongo.Client
	}
	models data.Models
}
