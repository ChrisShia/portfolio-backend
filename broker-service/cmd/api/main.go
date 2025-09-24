package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/v74/github"
)

func main() {

	app := Application{}

	//client := github.NewClient(nil)

	//go updateRepos(client)

	http.ListenAndServe(":8080", app.routes())
}

func updateRepos(client *github.Client) []*github.Repository {

	//TODO: Create timer... use time.
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fetchReposFromGithub(client)
		}
	}
}

func fetchReposFromGithub(client *github.Client) {
	opt := &github.RepositoryListByUserOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByUser(context.Background(), "ChrisShia", opt)
	if err != nil {
		fmt.Println(err)
	}

	for _, repo := range repos {
		if *repo.Fork {
			continue
		}
		fmt.Println(repo.GetHTMLURL())
	}
}
