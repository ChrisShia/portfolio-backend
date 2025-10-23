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
	IconSvg string `json:"icon"`
	Content string `json:"content,omitempty"`
}{
	{"2022", "Java Software Developer", softwareIconSvg, "Say a few words for the projects and technologies"},
	{"2021", "Physics Degree Msc", emcIconSvg, "Some words for Physics Msc"},
	{"2017", "Cern Internship", emcIconSvg, "Some words for Cern"},
	{"2013", "Physics Degree Bsc", emcIconSvg, "Thesis on Thermodynamic."},
	//{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	//{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
}
