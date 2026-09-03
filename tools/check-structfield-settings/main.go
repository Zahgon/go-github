package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/golangci/plugin-module-register/register"
	"github.com/google/go-github/v90/tools/structfield"
)

func init() {
	register.Plugin("structfield", structfield.New)
}

func main() {
	log.SetFlags(0)
	configPath := flag.String("config", "", "path to .golangci.yml (defaults to searching up from cwd)")
	packagesFlag := flag.String("packages", "./...", "comma-separated list of package patterns to analyze")
	includeTests := flag.Bool("tests", false, "include test files in analysis")
	fix := flag.Bool("fix", false, "remove obsolete exceptions and sort/dedupe lists in .golangci.yml")
	flag.Parse()

	resolvedConfig, repoRoot, err := resolveConfig(*configPath)
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	allowedNamesList, allowedTypesList, allowedNames, allowedTypes, err := readStructfieldSettings(resolvedConfig)
	if err != nil {
		log.Fatalf("parse config: %v", err)
	}
	if len(allowedNames) == 0 && len(allowedTypes) == 0 {
		log.Fatalf("no structfield settings found in %s", resolvedConfig)
	}

	duplicateNames := findDuplicates(allowedNamesList)
	duplicateTypes := findDuplicates(allowedTypesList)

	patterns := strings.Split(*packagesFlag, ",")
	for i, pattern := range patterns {
		patterns[i] = strings.TrimSpace(pattern)
	}

	usedNames, usedTypes, err := analyzeRepo(repoRoot, patterns, *includeTests, allowedNames, allowedTypes)
	if err != nil {
		log.Fatalf("analyze: %v", err)
	}

	obsoleteNames := diffKeys(allowedNames, usedNames)
	obsoleteTypes := diffKeys(allowedTypes, usedTypes)

	if len(obsoleteNames) == 0 && len(obsoleteTypes) == 0 && len(duplicateNames) == 0 && len(duplicateTypes) == 0 {
		return
	}

	if *fix {
		if err := removeObsoleteExceptions(resolvedConfig, obsoleteNames, obsoleteTypes); err != nil {
			log.Fatalf("fix: %v", err)
		}
		return
	}

	if len(obsoleteNames) > 0 {
		fmt.Println("Obsolete allowed-tag-names:")
		for _, name := range obsoleteNames {
			fmt.Printf("  - %v\n", name)
		}
	}
	if len(obsoleteTypes) > 0 {
		if len(obsoleteNames) > 0 {
			fmt.Println()
		}
		fmt.Println("Obsolete allowed-tag-types:")
		for _, name := range obsoleteTypes {
			fmt.Printf("  - %v\n", name)
		}
	}
	if len(duplicateNames) > 0 {
		if len(obsoleteNames) > 0 || len(obsoleteTypes) > 0 {
			fmt.Println()
		}
		fmt.Println("Duplicate allowed-tag-names:")
		for _, name := range sortedKeys(duplicateNames) {
			fmt.Printf("  - %v (%v)\n", name, duplicateNames[name])
		}
	}
	if len(duplicateTypes) > 0 {
		if len(obsoleteNames) > 0 || len(obsoleteTypes) > 0 || len(duplicateNames) > 0 {
			fmt.Println()
		}
		fmt.Println("Duplicate allowed-tag-types:")
		for _, name := range sortedKeys(duplicateTypes) {
			fmt.Printf("  - %v (%v)\n", name, duplicateTypes[name])
		}
	}
}

type golangciConfig struct {
	Linters struct {
		Settings struct {
			Custom struct {
				Structfield struct {
					Settings struct {
						structfield.Settings `yaml:",inline"`
					} `yaml:"settings"`
				} `yaml:"structfield"`
			} `yaml:"custom"`
		} `yaml:"settings"`
	} `yaml:"linters"`
}

func resolveConfig(configPath string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func readStructfieldSettings(configPath string) ([]string, []string, map[string]bool, map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

func analyzeRepo(repoRoot string, patterns []string, includeTests bool, allowedNames, allowedTypes map[string]bool) (map[string]bool, map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

var (
	nameMismatchRE  = regexp.MustCompile(`^change Go field name "([^"]+)" to ".*" for .* tag ".*" in struct "([^"]+)"$`)
	typeChangeRE    = regexp.MustCompile(`^change the "([^"]+)" field type to ".*" in the struct "([^"]+)"`)
	fieldInStructRE = regexp.MustCompile(`^the "([^"]+)" field in struct "([^"]+)" .*`)
)

func markUsedException(msg string, allowedNames, allowedTypes, usedNames, usedTypes map[string]bool) {
	_ = "STUB: not implemented"
	return
}

func diffKeys(all, used map[string]bool) []string { _ = "STUB: not implemented"; return nil }

func findDuplicates(values []string) map[string]int { _ = "STUB: not implemented"; return nil }

func sortedKeys(values map[string]int) []string { _ = "STUB: not implemented"; return nil }

func removeObsoleteExceptions(configPath string, obsoleteNames, obsoleteTypes []string) error {
	_ = "STUB: not implemented"
	return nil
}

type listItem struct {
	value string
	line  string
}

func appendSortedItems(lines []string, items []*listItem) []string {
	_ = "STUB: not implemented"
	return nil
}
