package main

import (
	"net/http"
)

func main() {

	app := App{}

	//client := github.NewClient(nil)

	//go updateRepos(client)

	http.ListenAndServe(":8080", app.routes())
}
