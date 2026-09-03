package main

import (
	"fmt"
	"log"

	"github.com/google/go-github/v90/github"
)

func fetchTopics(topic string) (*github.TopicsSearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	var topic string
	fmt.Print("Enter GitHub topic: ")
	fmt.Scanf("%s", &topic)

	topics, err := fetchTopics(topic)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, topic := range topics.Topics {
		fmt.Println(*topic.Name)
	}
}
