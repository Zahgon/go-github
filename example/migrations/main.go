package main

import (
	"fmt"
	"log"

	"github.com/google/go-github/v89/github"
)

func fetchAllUserMigrations() ([]*github.UserMigration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	migrations, err := fetchAllUserMigrations()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for i, m := range migrations {
		fmt.Printf("%v. %v", i+1, m.GetID())
	}
}
