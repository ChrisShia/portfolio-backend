package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *Application) routes() http.Handler {
	//mux := http.NewServeMux()
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Post("/projects", app.projects)

	mux.Get("/achievements", app.achievements)

	return mux
}

func (app *Application) projects(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) achievements(w http.ResponseWriter, r *http.Request) {

	var payload jsonResponse
	payload.Error = false
	payload.Message = ""
	payload.Data = achievements

	app.writeJSON(w, http.StatusOK, payload)
}

var achievements = []struct {
	Title       string `json:"title"`
	Institution string `json:"institution"`
	Year        string `json:"year,omitempty"`
	Content     string `json:"content,omitempty"`
}{
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes.Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Thesis on Thermodynamic of Black Holes."},
	{"Physics Undergraduate Degree", "University of Cyprus", "2013", "Something silly for last."},
}
