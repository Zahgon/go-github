package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v90/github"
	"github.com/sigstore/sigstore-go/pkg/bundle"
	"github.com/sigstore/sigstore-go/pkg/root"
	"github.com/sigstore/sigstore-go/pkg/verify"
)

var (
	owner = flag.String("owner", "cli", "GitHub organization or user to scope attestation lookup by")

	artifactDigest = flag.String("artifact-digest", "", "The digest of the artifact")

	artifactDigestAlgorithm = flag.String("artifact-digest-algorithm", "sha256", "The algorithm used to compute the digest of the artifact")

	expectedIssuer = flag.String("expected-issuer", "https://token.actions.githubusercontent.com", "Issuer of the OIDC token")

	expectedSAN = flag.String("expected-san", "https://github.com/cli/cli/.github/workflows/deployment.yml@refs/heads/trunk", "The expected Subject Alternative Name (SAN) of the certificate used to sign the attestation")

	trustedRootJSONPath = flag.String("trusted-root-json-path", "verifyartifact/trusted-root-public-good.json", "Path to the trusted root JSON file")
)

func usage() { _ = "STUB: not implemented"; return }

func main() {
	flag.Parse()
	if *artifactDigest == "" {
		fmt.Fprintln(os.Stderr, "artifact-digest is required.")
		usage()
		os.Exit(1)
	}

	token := os.Getenv("GITHUB_AUTH_TOKEN")

	if token == "" {
		log.Fatal("Unauthorized: No token present. Please set the GITHUB_AUTH_TOKEN environment variable to a valid token with `attestations:read` permission.")
	}

	ctx := context.Background()
	client, err := github.NewClient(github.WithAuthToken(token))
	if err != nil {
		log.Fatalf("Error creating GitHub client: %v", err)
	}

	attestations, _, err := client.Organizations.ListAttestations(ctx, *owner, fmt.Sprintf("%v:%v", *artifactDigestAlgorithm, *artifactDigest), nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(attestations.Attestations) == 0 {
		log.Fatal("No attestations found.")
	}

	sev, err := getSignedEntityVerifier()
	if err != nil {
		log.Fatal(err)
	}

	pb, err := getPolicyBuilder()
	if err != nil {
		log.Fatal(err)
	}

	var b *bundle.Bundle
	for _, attestation := range attestations.Attestations {
		if err := json.Unmarshal(attestation.Bundle, &b); err != nil {
			log.Fatal(err)
		}

		err := runVerification(sev, pb, b)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getTrustedMaterial() (root.TrustedMaterialCollection, error) {
	_ = "STUB: not implemented"
	return *new(root.TrustedMaterialCollection), nil
}

func getIdentityPolicies() ([]verify.PolicyOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSignedEntityVerifier() (*verify.Verifier, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPolicyBuilder() (*verify.PolicyBuilder, error) { _ = "STUB: not implemented"; return nil, nil }

func runVerification(sev *verify.Verifier, pb *verify.PolicyBuilder, b *bundle.Bundle) error {
	_ = "STUB: not implemented"
	return nil
}
