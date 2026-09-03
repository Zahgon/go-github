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

	if err := addUserSecret(ctx, client, secretName, secretValue, *owner, *repo); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added secret %q to the authenticated user\n", secretName)
}

func getSecretName() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getSecretValue(secretName string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func addUserSecret(ctx context.Context, client *github.Client, secretName, secretValue, owner, repo string) error {
	_ = "STUB: not implemented"
	return nil
}

func encryptSecretWithPublicKey(publicKey *github.PublicKey, secretName, secretValue string) (*github.EncryptedSecret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
