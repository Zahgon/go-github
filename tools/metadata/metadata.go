package main

import (
	"context"
	"go/ast"
	"regexp"
	"sync"

	"github.com/google/go-github/v89/github"
)

type operation struct {
	Name             string   `yaml:"name,omitempty" json:"name,omitempty"`
	DocumentationURL string   `yaml:"documentation_url,omitempty" json:"documentation_url,omitempty"`
	OpenAPIFiles     []string `yaml:"openapi_files,omitempty" json:"openapi_files,omitempty"`
	Deprecated       bool     `yaml:"deprecated,omitempty" json:"deprecated,omitempty"`
}

func (o *operation) equal(other *operation) bool { _ = "STUB: not implemented"; return false }

func (o *operation) clone() *operation { _ = "STUB: not implemented"; return nil }

func operationsEqual(a, b []*operation) bool { _ = "STUB: not implemented"; return false }

func sortOperations(ops []*operation) { _ = "STUB: not implemented"; return }

func normalizeOpPath(opPath string) string { _ = "STUB: not implemented"; return "" }

func normalizedOpName(name string) string { _ = "STUB: not implemented"; return "" }

var opNameRe = regexp.MustCompile(`(?i)(\S+)(?:\s+(\S.*))?`)

func parseOpName(id string) (verb, url string) { _ = "STUB: not implemented"; return "", "" }

type operationsFile struct {
	ManualOps   []*operation `yaml:"operations,omitempty"`
	OverrideOps []*operation `yaml:"operation_overrides,omitempty"`
	GitCommit   string       `yaml:"openapi_commit,omitempty"`
	OpenapiOps  []*operation `yaml:"openapi_operations,omitempty"`

	mu          sync.Mutex
	resolvedOps map[string]*operation
}

func (m *operationsFile) resolve() { _ = "STUB: not implemented"; return }

func (m *operationsFile) saveFile(filename string) (errOut error) {
	_ = "STUB: not implemented"
	return nil
}

func (m *operationsFile) updateFromGithub(ctx context.Context, client *github.Client, ref string) error {
	_ = "STUB: not implemented"
	return nil
}

func loadOperationsFile(filename string) (*operationsFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addOperation(ops []*operation, filename, opName, docURL string, deprecated bool) []*operation {
	_ = "STUB: not implemented"
	return nil
}

func unusedOps(opsFile *operationsFile, dir string) ([]*operation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updateDocsVisitor(opsFile *operationsFile) nodeVisitor {
	_ = "STUB: not implemented"
	return *new(nodeVisitor)
}

func updateDocs(opsFile *operationsFile, dir string) error { _ = "STUB: not implemented"; return nil }

type nodeVisitor func(serviceMethod string, fn *ast.FuncDecl, cmap ast.CommentMap) error

func visitServiceMethods(dir string, writeFiles bool, visit nodeVisitor) error {
	_ = "STUB: not implemented"
	return nil
}

func visitFileMethods(updateFile bool, filename string, visit nodeVisitor) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	metaOpRe        = regexp.MustCompile(`(?i)\s*//\s*meta:operation\s+(\S.+)`)
	undocRE         = regexp.MustCompile(`(?i)\s*//\s*Note:\s+\S.+ uses the undocumented GitHub API endpoint`)
	docLineRE       = regexp.MustCompile(`(?i)\s*//\s*GitHub\s+API\s+docs:`)
	deprecatedRE    = regexp.MustCompile(`(?i)\s*//\s*Deprecated: This endpoint has been deprecated by GitHub\.`)
	anyDeprecatedRE = regexp.MustCompile(`(?i)\s*//\s*Deprecated:`)
)

func methodOps(opsFile *operationsFile, cmap ast.CommentMap, fn *ast.FuncDecl) ([]*operation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const metadataDocsAPIVersion = "2022-11-28"

func normalizeDocURL(docURL string) string { _ = "STUB: not implemented"; return "" }

func nodeServiceMethod(fn *ast.FuncDecl) string { _ = "STUB: not implemented"; return "" }

var skipServiceMethod = map[string]bool{
	"CopilotService.DownloadCopilotMetrics":      true,
	"CopilotService.DownloadDailyMetrics":        true,
	"CopilotService.DownloadPeriodicMetrics":     true,
	"CopilotService.DownloadUserDailyMetrics":    true,
	"CopilotService.DownloadUserPeriodicMetrics": true,
}
