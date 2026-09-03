package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/v90/github"
)

func encryptSecret(publicKeyB64, secret string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %v <create|delete> [flags]\n", os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		runCreate(os.Args[2:])
	case "delete":
		runDelete(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %q. Must be one of: create, delete\n", os.Args[1])
		os.Exit(1)
	}
}

func newFlagSet(name string) (*flag.FlagSet, *string) { _ = "STUB: not implemented"; return nil, nil }

func parseAndInit(fs *flag.FlagSet, enterprise *string, args []string) (context.Context, *github.Client, string) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, ""
}

func runCreate(args []string) { _ = "STUB: not implemented"; return }

func runDelete(args []string) { _ = "STUB: not implemented"; return }

func newClient(token, apiURL string) *github.Client { _ = "STUB: not implemented"; return nil }

func requireEnv(name string) string { _ = "STUB: not implemented"; return "" }

func requireFlag(name, val string) { _ = "STUB: not implemented"; return }

func requireIntFlag(name string, val int64) { _ = "STUB: not implemented"; return }
