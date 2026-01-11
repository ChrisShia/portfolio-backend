package main

import (
	"net/http"
	"portfolio/internal/data"
)

func (app *App) codingSkills(w http.ResponseWriter, r *http.Request) {
	var payload jsonResponse
	payload.Error = false
	payload.Message = ""

	skills, err := app.allCodingSkills()
	if err != nil {
		payload.Error = true
		payload.Message = err.Error()
		_ = app.writeJSON(w, http.StatusInternalServerError, payload)
		return
	}

	payload.Data = skills

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *App) allCodingSkills() ([]*data.CodingSkill, error) {
	allDocs, err := app.models.CodingSkill.All()
	if err != nil {
		app.log.PrintError(err, nil)
		return nil, err
	}

	return allDocs, nil
}
