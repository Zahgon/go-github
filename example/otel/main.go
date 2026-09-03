package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/otel/v90"
	"github.com/google/go-github/v90/github"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {

	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("failed to initialize stdouttrace exporter: %v", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	t := otel.NewTransport(
		http.DefaultTransport,
		otel.WithTracerProvider(tp),
	)

	client, err := github.NewClient(github.WithTransport(t))
	if err != nil {
		log.Fatalf("Error creating GitHub client: %v", err)
	}

	limits, resp, err := client.RateLimit.Get(context.Background())
	if err != nil {
		log.Printf("Error fetching rate limits: %v", err)
	} else {
		fmt.Printf("Core Rate Limit: %v/%v (Resets at %v)\n",
			limits.GetCore().Remaining,
			limits.GetCore().Limit,
			limits.GetCore().Reset)
	}

	if resp != nil {
		fmt.Printf("Response Status: %v\n", resp.Status)
	}
}
