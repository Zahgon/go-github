package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

var (
	sinceTag = flag.String("tag", "", "List all changes since this tag (e.g. 'v76.0.0')")

	descriptionRE = regexp.MustCompile(`^\* (.*?\((#[^\)]+)\))`)
	releaseTagRE  = regexp.MustCompile(`[^a-zA-Z0-9.\-_]+`)
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	priorRelease := *sinceTag
	if priorRelease == "" {
		priorRelease = getPriorRelease()
		log.Printf("Prior release: %v", priorRelease)
	}

	newChanges := newChangesSinceRelease(priorRelease)

	releaseNotes := genReleaseNotes(newChanges)
	fmt.Printf("%v%v", releaseNotes, "\n")

	log.Print("Done.")
}

func genReleaseNotes(text string) string { _ = "STUB: not implemented"; return "" }

func splitIntoPRs(text string) []string { _ = "STUB: not implemented"; return nil }

func splitBreakingLines(allLines []string) (breaking, nonBreaking []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func genRefLines(breaking, nonBreaking []string) (ref, refNon []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func runCommand(cmdArgs []string) string { _ = "STUB: not implemented"; return "" }

//nolint:gosec

func newChangesSinceRelease(priorRelease string) string { _ = "STUB: not implemented"; return "" }

func getPriorRelease() string { _ = "STUB: not implemented"; return "" }

const releaseNotesFmt = `
This release contains the following breaking API changes:

%v

...and the following additional changes:

%v

&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&

This release contains the following breaking API changes:

%v

...and the following additional changes:

%v
`
