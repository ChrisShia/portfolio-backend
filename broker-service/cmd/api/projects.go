package main

import (
	"net/http"
)

func (app *App) projects(w http.ResponseWriter, r *http.Request) {
	var payload jsonResponse
	payload.Error = false
	payload.Message = ""
	payload.Data = projects

	app.writeJSON(w, http.StatusOK, payload)
}

var projects = []struct {
	Title            string   `json:"title"`
	GithubUrl        string   `json:"github_url"`
	Url              string   `json:"url"`
	Technologies     []string `json:"technologies"`
	ShortDescription string   `json:"short_description"`
}{
	{"First Project", "https://www.github.com/ChrisShia",
		"",
		[]string{javaIconSvg, golangIconSvg, reactIconSvg},
		"First Project is about going the extra mile with REACT..and what a mile this is...I prefer go..."},
	{"Another Project", "https://www.github.com/ChrisShia",
		"",
		[]string{javaIconSvg, redisIconSvg, dockerIconSvg, golangIconSvg, reactIconSvg},
		"Project 2 is about going the extra mile with REACT..and what a mile this is...I prefer go..."},
	{"Some other Project", "https://www.github.com/ChrisShia",
		"",
		[]string{javaIconSvg, reactIconSvg, golangIconSvg},
		"Project 1 is about going the extra mile with REACT..and what a mile this is...I prefer go..."},
	{"Project blah", "https://www.github.com/ChrisShia",
		"",
		[]string{javaIconSvg, redisIconSvg, golangIconSvg, reactIconSvg},
		"Project 1 is about going the extra mile with REACT..and what a mile this is...I prefer go..."},
}
