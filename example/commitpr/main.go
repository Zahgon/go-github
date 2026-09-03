package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/google/go-github/v90/github"
)

var (
	sourceOwner   = flag.String("source-owner", "", "Name of the owner (user or org) of the repo to create the commit in.")
	sourceRepo    = flag.String("source-repo", "", "Name of repo to create the commit in.")
	commitMessage = flag.String("commit-message", "", "Content of the commit message.")
	commitBranch  = flag.String("commit-branch", "", "Name of branch to create the commit in. If it does not already exists, it will be created using the `base-branch` parameter")
	repoBranch    = flag.String("repo-branch", "", "Name of the repository where the changes in the pull request were made. This field is required for cross-repository pull requests if both repositories are owned by the same organization")
	baseBranch    = flag.String("base-branch", "master", "Name of branch to create the `commit-branch` from.")
	prRepoOwner   = flag.String("merge-repo-owner", "", "Name of the owner (user or org) of the repo to create the PR against. If not specified, the value of the `-source-owner` flag will be used.")
	prRepo        = flag.String("merge-repo", "", "Name of repo to create the PR against. If not specified, the value of the `-source-repo` flag will be used.")
	prBranch      = flag.String("merge-branch", "master", "Name of branch to create the PR against (the one you want to merge your branch in via the PR).")
	prSubject     = flag.String("pr-title", "", "Title of the pull request. If not specified, no pull request will be created.")
	prDescription = flag.String("pr-text", "", "Text to put in the description of the pull request.")
	sourceFiles   = flag.String("files", "", `Comma-separated list of files to commit and their location.
The local file is separated by its target location by a semi-colon.
If the file should be in the same location with the same name, you can just put the file name and omit the repetition.
Example: README.md,main.go:github/examples/commitpr/main.go`)
	authorName  = flag.String("author-name", "", "Name of the author of the commit.")
	authorEmail = flag.String("author-email", "", "Email of the author of the commit.")
	privateKey  = flag.String("private-key", "", "Path to the private key to use to sign the commit.")
)

var (
	client *github.Client
	ctx    = context.Background()
)

func getRef() (ref *github.Reference, err error) { _ = "STUB: not implemented"; return nil, nil }

func branchRef(name string) string { _ = "STUB: not implemented"; return "" }

func getTree(ref *github.Reference) (tree *github.Tree, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFileContent(fileArg string) (targetName string, b []byte, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func pushCommit(ref *github.Reference, tree *github.Tree) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func createPR() (err error) { _ = "STUB: not implemented"; return nil }

func main() {
	flag.Parse()
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	if *sourceOwner == "" || *sourceRepo == "" || *commitBranch == "" || *sourceFiles == "" || *authorName == "" || *authorEmail == "" {
		log.Fatal("You need to specify a non-empty value for the flags `-source-owner`, `-source-repo`, `-commit-branch`, `-files`, `-author-name` and `-author-email`")
	}
	c, err := github.NewClient(github.WithAuthToken(token))
	if err != nil {
		log.Fatal(err)
	}
	client = c

	ref, err := getRef()
	if err != nil {
		log.Fatalf("Unable to get/create the commit reference: %v", err)
	}
	if ref == nil {
		log.Fatal("No error where returned but the reference is nil")
	}

	tree, err := getTree(ref)
	if err != nil {
		log.Fatalf("Unable to create the tree based on the provided files: %v", err)
	}

	if err := pushCommit(ref, tree); err != nil {
		log.Fatalf("Unable to create the commit: %v", err)
	}

	if err := createPR(); err != nil {
		log.Fatalf("Error while creating the pull request: %v", err)
	}
}
