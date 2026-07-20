package github

import "context"

type RateLimitService service

type Rate struct {
	Limit int `json:"limit"`

	Remaining int `json:"remaining"`

	Used int `json:"used"`

	Reset Timestamp `json:"reset"`

	Resource string `json:"resource,omitempty"`
}

func (r Rate) String() string { _ = "STUB: not implemented"; return "" }

type RateLimits struct {
	Core *Rate `json:"core"`

	Search *Rate `json:"search"`

	GraphQL *Rate `json:"graphql"`

	IntegrationManifest *Rate `json:"integration_manifest"`

	SourceImport              *Rate `json:"source_import"`
	CodeScanningUpload        *Rate `json:"code_scanning_upload"`
	ActionsRunnerRegistration *Rate `json:"actions_runner_registration"`
	SCIM                      *Rate `json:"scim"`
	DependencySnapshots       *Rate `json:"dependency_snapshots"`
	CodeSearch                *Rate `json:"code_search"`
	AuditLog                  *Rate `json:"audit_log"`
	DependencySBOM            *Rate `json:"dependency_sbom"`
}

func (r RateLimits) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /rate_limit
func (s *RateLimitService) Get(ctx context.Context) (*RateLimits, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
