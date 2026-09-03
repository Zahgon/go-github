package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofri/go-github-pagination/githubpagination"
	"github.com/gofri/go-github-ratelimit/v2/github_ratelimit"
	"github.com/gofri/go-github-ratelimit/v2/github_ratelimit/github_primary_ratelimit"
	"github.com/gofri/go-github-ratelimit/v2/github_ratelimit/github_secondary_ratelimit"
	"github.com/google/go-github/v90/github"
)

func main() {
	var username string
	fmt.Print("Enter GitHub username: ")
	fmt.Scanf("%s", &username)

	rateLimiter := github_ratelimit.New(nil,
		github_primary_ratelimit.WithLimitDetectedCallback(func(ctx *github_primary_ratelimit.CallbackContext) {
			fmt.Printf("Primary rate limit detected: category %v, reset time: %v\n", ctx.Category, ctx.ResetTime)
		}),
		github_secondary_ratelimit.WithLimitDetectedCallback(func(ctx *github_secondary_ratelimit.CallbackContext) {
			fmt.Printf("Secondary rate limit detected: reset time: %v, total sleep time: %v\n", ctx.ResetTime, ctx.TotalSleepTime)
		}),
	)

	paginator := githubpagination.NewClient(rateLimiter,
		githubpagination.WithPerPage(100),
	)
	client, err := github.NewClient(github.WithHTTPClient(paginator))
	if err != nil {
		log.Fatalf("Error creating GitHub client: %v", err)
	}

	repos, _, err := client.Repositories.ListByUser(context.Background(), username, nil)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for i, repo := range repos {
		fmt.Printf("%v. %v\n", i+1, repo.GetName())
	}
}
