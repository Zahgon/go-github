package github

import (
	"context"
)

type SecretScanningCustomPattern struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Pattern string `json:"pattern"`

	Slug string `json:"slug"`

	State string `json:"state"`

	PushProtectionEnabled bool `json:"push_protection_enabled"`

	StartDelimiter *string `json:"start_delimiter,omitempty"`

	EndDelimiter *string `json:"end_delimiter,omitempty"`

	MustMatch []string `json:"must_match,omitempty"`

	MustNotMatch []string `json:"must_not_match,omitempty"`

	CustomPatternVersion *string `json:"custom_pattern_version,omitempty"`

	CreatedAt *Timestamp `json:"created_at,omitempty"`

	UpdatedAt *Timestamp `json:"updated_at,omitempty"`
}

type SecretScanningCustomPatternRequest struct {
	Name string `json:"name"`

	Pattern string `json:"pattern"`

	StartDelimiter *string `json:"start_delimiter,omitempty"`

	EndDelimiter *string `json:"end_delimiter,omitempty"`

	MustMatch []string `json:"must_match,omitempty"`

	MustNotMatch []string `json:"must_not_match,omitempty"`
}

type SecretScanningCreateCustomPatternsRequest struct {
	Patterns []*SecretScanningCustomPatternRequest `json:"patterns"`
}

type SecretScanningCreateCustomPatternsResponse struct {
	CreatedPatterns []*SecretScanningCustomPattern `json:"created_patterns"`
}

type SecretScanningCustomPatternToDelete struct {
	PatternID int64 `json:"pattern_id"`

	CustomPatternVersion *string `json:"custom_pattern_version,omitempty"`
}

type SecretScanningDeleteCustomPatternsRequest struct {
	Patterns []*SecretScanningCustomPatternToDelete `json:"patterns"`

	PostDeleteAction *string `json:"post_delete_action,omitempty"`
}

type SecretScanningUpdateCustomPatternRequest struct {
	Pattern *string `json:"pattern,omitempty"`

	StartDelimiter *string `json:"start_delimiter,omitempty"`

	EndDelimiter *string `json:"end_delimiter,omitempty"`

	MustMatch []string `json:"must_match,omitempty"`

	MustNotMatch []string `json:"must_not_match,omitempty"`

	CustomPatternVersion *string `json:"custom_pattern_version"`
}

type SecretScanningCustomPatternListOptions struct {
	State string `url:"state,omitempty"`

	PushProtection string `url:"push_protection,omitempty"`

	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/secret-scanning/custom-patterns
func (s *SecretScanningService) ListCustomPatternsForRepo(ctx context.Context, owner, repo string, opts *SecretScanningCustomPatternListOptions) ([]*SecretScanningCustomPattern, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/secret-scanning/custom-patterns
func (s *SecretScanningService) CreateCustomPatternsForRepo(ctx context.Context, owner, repo string, body SecretScanningCreateCustomPatternsRequest) (*SecretScanningCreateCustomPatternsResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/secret-scanning/custom-patterns/{pattern_id}
func (s *SecretScanningService) UpdateCustomPatternForRepo(ctx context.Context, owner, repo string, patternID int64, body SecretScanningUpdateCustomPatternRequest) (*SecretScanningCustomPattern, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/secret-scanning/custom-patterns
func (s *SecretScanningService) DeleteCustomPatternsForRepo(ctx context.Context, owner, repo string, body SecretScanningDeleteCustomPatternsRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/secret-scanning/custom-patterns
func (s *SecretScanningService) ListCustomPatternsForOrg(ctx context.Context, org string, opts *SecretScanningCustomPatternListOptions) ([]*SecretScanningCustomPattern, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/secret-scanning/custom-patterns
func (s *SecretScanningService) CreateCustomPatternsForOrg(ctx context.Context, org string, body SecretScanningCreateCustomPatternsRequest) (*SecretScanningCreateCustomPatternsResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/secret-scanning/custom-patterns/{pattern_id}
func (s *SecretScanningService) UpdateCustomPatternForOrg(ctx context.Context, org string, patternID int64, body SecretScanningUpdateCustomPatternRequest) (*SecretScanningCustomPattern, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/secret-scanning/custom-patterns
func (s *SecretScanningService) DeleteCustomPatternsForOrg(ctx context.Context, org string, body SecretScanningDeleteCustomPatternsRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /enterprises/{enterprise}/secret-scanning/custom-patterns
func (s *SecretScanningService) ListCustomPatternsForEnterprise(ctx context.Context, enterprise string, opts *SecretScanningCustomPatternListOptions) ([]*SecretScanningCustomPattern, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/secret-scanning/custom-patterns
func (s *SecretScanningService) CreateCustomPatternsForEnterprise(ctx context.Context, enterprise string, body SecretScanningCreateCustomPatternsRequest) (*SecretScanningCreateCustomPatternsResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/secret-scanning/custom-patterns/{pattern_id}
func (s *SecretScanningService) UpdateCustomPatternForEnterprise(ctx context.Context, enterprise string, patternID int64, body SecretScanningUpdateCustomPatternRequest) (*SecretScanningCustomPattern, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/secret-scanning/custom-patterns
func (s *SecretScanningService) DeleteCustomPatternsForEnterprise(ctx context.Context, enterprise string, body SecretScanningDeleteCustomPatternsRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
