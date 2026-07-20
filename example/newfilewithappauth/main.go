package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v89/github"
)

func main() {
	const gitHost = "https://git.api.com"

	privatePem, err := os.ReadFile("path/to/pem")
	if err != nil {
		log.Fatalf("failed to read pem: %v", err)
	}

	itr, err := ghinstallation.NewAppsTransport(http.DefaultTransport, 10, privatePem)
	if err != nil {
		log.Fatalf("failed to create app transport: %v", err)
	}
	itr.BaseURL = gitHost

	client, err := github.NewClient(github.WithHTTPClient(&http.Client{
		Transport: itr,
		Timeout:   time.Second * 30,
	}), github.WithEnterpriseURLs(gitHost, gitHost))
	if err != nil {
		log.Fatalf("failed to create git client for app: %v", err)
	}

	installations, _, err := client.Apps.ListInstallations(context.Background(), &github.ListOptions{})
	if err != nil {
		log.Fatalf("failed to list installations: %v", err)
	}

	var installID int64
	for _, val := range installations {
		installID = val.GetID()
	}

	token, _, err := client.Apps.CreateInstallationToken(
		context.Background(),
		installID,
		&github.InstallationTokenOptions{})
	if err != nil {
		log.Fatalf("failed to create installation token: %v", err)
	}

	apiClient, err := github.NewClient(github.WithAuthToken(token.GetToken()), github.WithEnterpriseURLs(gitHost, gitHost))
	if err != nil {
		log.Fatalf("failed to create new git client with token: %v", err)
	}

	_, resp, err := apiClient.Repositories.CreateFile(
		context.Background(),
		"repoOwner",
		"sample-repo",
		"example/foo.txt",
		&github.RepositoryContentFileOptions{
			Content: []byte("foo"),
			Message: github.Ptr("sample commit"),
			SHA:     nil,
		})
	if err != nil {
		log.Fatalf("failed to create new file: %v", err)
	}

	log.Printf("file written status code: %v", resp.StatusCode)
}
