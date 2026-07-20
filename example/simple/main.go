package main

import (
	"fmt"
	"log"

	"github.com/google/go-github/v89/github"
)

func fetchOrganizations(username string) ([]*github.Organization, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	var username string
	fmt.Print("Enter GitHub username: ")
	fmt.Scanf("%s", &username)

	organizations, err := fetchOrganizations(username)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for i, organization := range organizations {
		fmt.Printf("%v. %v\n", i+1, organization.GetLogin())
	}
}
