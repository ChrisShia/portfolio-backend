package main

import "net/http"

func (app *App) achievements(w http.ResponseWriter, r *http.Request) {
	var payload jsonResponse
	payload.Error = false
	payload.Message = ""
	payload.Data = achievements

	app.writeJSON(w, http.StatusOK, payload)
}

var achievements = []struct {
	Year    string `json:"year,omitempty"`
	Title   string `json:"title"`
	Content string `json:"content,omitempty"`
}{
	{"2022", "Software Developer Albourne", "Say a few words for the projects and technologies"},
	{"2021", "Physics Degree Msc", "Some words for Physics Msc"},
	{"2017", "Cern Internship", "Some words for Cern"},
	{"2013", "Physics Degree Bsc", "Thesis on Thermodynamic."},
	//{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	//{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
}
