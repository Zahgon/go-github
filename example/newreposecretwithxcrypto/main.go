package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v90/github"
)

var (
	repo  = flag.String("repo", "", "The repo that the secret should be added to, ex. go-github")
	owner = flag.String("owner", "", "The owner of there repo this should be added to, ex. google")
)

func main() {
	flag.Parse()

	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("please provide a GitHub API token via env variable GITHUB_AUTH_TOKEN")
	}

	if *repo == "" {
		log.Fatal("please provide required flag --repo to specify GitHub repository ")
	}

	if *owner == "" {
		log.Fatal("please provide required flag --owner to specify GitHub user/org owner")
	}

	secretName, err := getSecretName()
	if err != nil {
		log.Fatal(err)
	}

	secretValue, err := getSecretValue(secretName)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client, err := github.NewClient(github.WithAuthToken(token))
	if err != nil {
		log.Fatal(err)
	}

	if err := addRepoSecret(ctx, client, *owner, *repo, secretName, secretValue); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added secret %q to the repo %v/%v\n", secretName, *owner, *repo)
}

func getSecretName() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getSecretValue(secretName string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func addRepoSecret(ctx context.Context, client *github.Client, owner, repo, secretName, secretValue string) error {
	_ = "STUB: not implemented"
	return nil
}
