package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/google/go-github/v90/github"
)

var helpVars = kong.Vars{
	"update_openapi_help": `
Update openapi_operations.yaml from OpenAPI descriptions in github.com/github/rest-api-description at the given git ref.
`,

	"update_go_help": `
Update go source code to be consistent with openapi_operations.yaml.
 - Adds and updates "// GitHub API docs:" comments for service methods.
 - Updates "//meta:operation" comments to use canonical operation names.
 - Updates formatting of "//meta:operation" comments to make sure there isn't a space between the "//" and the "meta".
 - Formats modified files with the equivalent of "go fmt".
`,

	"format_help": `Format white space in openapi_operations.yaml and sort its operations.`,
	"unused_help": `List operations in openapi_operations.yaml that aren't used by any service methods.`,

	"working_dir_help": `Working directory. Should be the root of the go-github repository.`,
	"openapi_ref_help": `Git ref to pull OpenAPI descriptions from.`,

	"openapi_validate_help": `
Instead of updating, make sure that the operations in openapi_operations.yaml's "openapi_operations" field are
consistent with the SHA listed in "openapi_commit". This is run in CI as a convenience so that reviewers can trust
changes to openapi_operations.yaml.
`,

	"unused_deprecated_help": `Only list deprecated operations in openapi_operations.yaml that aren't used by any service methods.`,
	"output_json_help":       `Output JSON.`,
}

type rootCmd struct {
	UpdateOpenAPI updateOpenAPICmd `kong:"cmd,name=update-openapi,help=${update_openapi_help}"`
	UpdateGo      updateGoCmd      `kong:"cmd,help=${update_go_help}"`
	Format        formatCmd        `kong:"cmd,help=${format_help}"`
	Unused        unusedCmd        `kong:"cmd,help=${unused_help}"`

	WorkingDir string `kong:"short=C,default=.,help=${working_dir_help}"`

	GithubURL string `kong:"hidden,default='https://api.github.com'"`
	UploadURL string `kong:"hidden,default='https://uploads.github.com'"`
}

func (c *rootCmd) opsFile() (string, *operationsFile, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func githubClient(apiURL, uploadURL string) (*github.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type updateOpenAPICmd struct {
	Ref            string `kong:"default=main,help=${openapi_ref_help}"`
	ValidateGithub bool   `kong:"name=validate,help=${openapi_validate_help}"`
}

func (c *updateOpenAPICmd) Run(root *rootCmd) error { _ = "STUB: not implemented"; return nil }

type formatCmd struct{}

func (c *formatCmd) Run(root *rootCmd) error { _ = "STUB: not implemented"; return nil }

type updateGoCmd struct{}

func (c *updateGoCmd) Run(root *rootCmd) error { _ = "STUB: not implemented"; return nil }

type unusedCmd struct {
	Deprecated bool `kong:"help=${unused_deprecated_help}"`
	JSON       bool `kong:"help=${output_json_help}"`
}

func (c *unusedCmd) Run(root *rootCmd, k *kong.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	err := run(os.Args[1:], nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, opts []kong.Option) error { _ = "STUB: not implemented"; return nil }
