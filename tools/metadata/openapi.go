package main

import (
	"context"
	"regexp"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/google/go-github/v90/github"
)

const (
	descriptionsOwnerName = "github"
	descriptionsRepoName  = "rest-api-description"
	descriptionsPath      = "descriptions"
)

type openapiFile struct {
	description  *openapi3.T
	filename     string
	plan         string
	planIdx      int
	releaseMajor int
	releaseMinor int
}

func getOpsFromGithub(ctx context.Context, client *github.Client, gitRef string) ([]*operation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deprecateOperations(descs []*openapiFile, ops []*operation) { _ = "STUB: not implemented"; return }

func (o *openapiFile) loadDescription(ctx context.Context, client *github.Client, gitRef string) error {
	_ = "STUB: not implemented"
	return nil
}

var dirPatterns = []*regexp.Regexp{
	regexp.MustCompile(`^(?P<plan>api\.github\.com)(-(?P<major>\d+)\.(?P<minor>\d+))?$`),
	regexp.MustCompile(`^(?P<plan>ghec)(-(?P<major>\d+)\.(?P<minor>\d+))?$`),
	regexp.MustCompile(`^(?P<plan>ghes)(-(?P<major>\d+)\.(?P<minor>\d+))?$`),
}

func getDescriptions(ctx context.Context, client *github.Client, gitRef string) ([]*openapiFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
