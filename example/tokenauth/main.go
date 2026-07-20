package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v89/github"
	"golang.org/x/term"
)

func main() {
	fmt.Print("GitHub Token: ")
	token, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()

	ctx := context.Background()
	client, err := github.NewClient(github.WithAuthToken(string(token)))
	if err != nil {
		log.Fatalf("Error creating GitHub client: %v", err)
	}

	user, resp, err := client.Users.Get(ctx, "")
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}

	log.Printf("Rate: %#v\n", resp.Rate)

	if !resp.TokenExpiration.IsZero() {
		log.Printf("Token Expiration: %v\n", resp.TokenExpiration)
	}

	fmt.Printf("\n%v\n", github.Stringify(user))
}
