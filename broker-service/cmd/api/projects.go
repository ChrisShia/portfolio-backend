package main

import (
	"net/http"
	"portfolio/internal/data"
)

func (app *App) projects(w http.ResponseWriter, r *http.Request) {
	var payload jsonResponse
	payload.Error = false
	payload.Message = ""
	allProjects, err := app.allProjects()
	if err != nil {
		payload.Error = true
		payload.Message = err.Error()
		_ = app.writeJSON(w, http.StatusInternalServerError, payload)
	}

	payload.Data = allProjects

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *App) allProjects() ([]*data.Project, error) {
	allDocs, err := app.models.Project.All()
	if err != nil {
		app.log.PrintError(err, nil)
		return nil, err
	}

	return allDocs, nil
}

var projects = []struct {
	Title            string   `json:"title"`
	GithubUrl        string   `json:"github_url"`
	Url              string   `json:"url"`
	Technologies     []string `json:"technologies"`
	ShortDescription string   `json:"short_description"`
	LongDescription  string   `json:"long_description"`
	Icon             string   `json:"icon"`
}{
	{"Hub", "https://www.github.com/ChrisShia/hub",
		"",
		[]string{golangIconSvg, redisIconSvg, dockerIconSvg, postgresqlIconSvg},
		"Golang REST API for a collection of services, meant to show-case skill set in REST API microservice architecture. Containerized tools and microservices such as authentication with PostgreSQL database and Redis for implementing a scalable rate limiter.", "", ""},
	{"Math Depot", "https://github.com/ChrisShia/math-depot",
		"",
		[]string{golangIconSvg},
		"Centralization of mathematical functions, tools and structures for linear algebra.", "", ""},
	{"Hanoi", "https://github.com/ChrisShia/hanoi",
		"",
		[]string{golangIconSvg},
		"Recursive approach to the solution of the tower of hanoi problem.", "", ""},
	{"Rate Limiter", "https://github.com/ChrisShia/ratelimiter",
		"",
		[]string{golangIconSvg, redisIconSvg},
		"A rate limiter interface. Provides a redis implementation.", "", ""},
	{"Reading", "https://github.com/ChrisShia/goread",
		"",
		[]string{golangIconSvg},
		"Centralization of functions for reading text input. Utilization of concurrent fan-in-fan-out pattern for parallel execution, using go native tools, i.e. routines and channels.", "", ""},
	{"Sort", "https://github.com/ChrisShia/gosort",
		"",
		[]string{golangIconSvg},
		"Implementation and experimenting with established sorting algorithms.", "", ""},
	{"Sort", "https://github.com/ChrisShia/gosort",
		"",
		[]string{golangIconSvg},
		"Implementation and experimenting with established sorting algorithms.", "", ""},
	{"Json Logger", "https://github.com/ChrisShia/jsonlog",
		"",
		[]string{golangIconSvg},
		"A small but handy logger in json format.", "", ""},
	{"Advent Of Code 2024", "https://github.com/ChrisShia/advent-of-code-2024",
		"adventofcode.com/2024",
		[]string{golangIconSvg},
		"Personal solutions to the exercises found in the 2024 advent of code calendar. String manipulation, graph data structures, mathematical functions.", "", ""},
}
