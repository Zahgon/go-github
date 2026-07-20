package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v89/github"
)

func main() {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	repo := os.Getenv("GITHUB_REPOSITORY_NAME")
	owner := os.Getenv("GITHUB_REPOSITORY_OWNER")

	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}

	ctx := context.Background()
	client, err := github.NewClient(github.WithAuthToken(token))
	if err != nil {
		log.Fatal(err)
	}

	expectedPageSize := 2

	opts := &github.EnvironmentListOptions{ListOptions: github.ListOptions{PerPage: expectedPageSize}}
	envResponse, _, err := client.Repositories.ListEnvironments(ctx, owner, repo, opts)
	if err != nil {
		log.Fatal(err)
	}

	if len(envResponse.Environments) != expectedPageSize {
		log.Fatal("Unexpected number of environments returned")
	}

	fmt.Printf("%v environments returned\n", len(envResponse.Environments))
}
