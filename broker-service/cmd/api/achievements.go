package main

import (
	"net/http"
	"portfolio/internal/data"
)

func (app *App) achievements(w http.ResponseWriter, r *http.Request) {
	var payload jsonResponse
	payload.Error = false
	payload.Message = ""

	achievements, err := app.allAchievements()
	if err != nil {
		payload.Error = true
		payload.Message = err.Error()
		_ = app.writeJSON(w, http.StatusInternalServerError, payload)
		return
	}

	payload.Data = achievements

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *App) allAchievements() ([]*data.Achievement, error) {
	allDocs, err := app.models.Achievement.All()
	if err != nil {
		app.log.PrintError(err, nil)
		return nil, err
	}

	return allDocs, nil
}
