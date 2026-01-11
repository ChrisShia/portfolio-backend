package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"portfolio/internal/data"

	"github.com/ChrisShia/jsonlog"
)

func main() {
	cfg := config{}
	cfg.flags()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	app := App{config: cfg, log: logger}

	app.connectMongoDB()
	defer app.disconnectMongoDB()

	app.models = data.New(app.mongo.client)

	err := http.ListenAndServe(fmt.Sprintf(":%d", app.config.port), app.Routes())
	if err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			return
		default:
			app.log.PrintFatal(err, nil)
		}
	}
}
