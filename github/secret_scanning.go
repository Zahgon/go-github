package github

import (
	"context"
)

type SecretScanningService service

type SecretScanningAlert struct {
	Number                                     *int                                `json:"number,omitempty"`
	CreatedAt                                  *Timestamp                          `json:"created_at,omitempty"`
	URL                                        *string                             `json:"url,omitempty"`
	HTMLURL                                    *string                             `json:"html_url,omitempty"`
	LocationsURL                               *string                             `json:"locations_url,omitempty"`
	FirstLocationDetected                      *SecretScanningAlertLocationDetails `json:"first_location_detected,omitempty"`
	HasMoreLocations                           *bool                               `json:"has_more_locations,omitempty"`
	State                                      *string                             `json:"state,omitempty"`
	Resolution                                 *string                             `json:"resolution,omitempty"`
	ResolvedAt                                 *Timestamp                          `json:"resolved_at,omitempty"`
	ResolvedBy                                 *User                               `json:"resolved_by,omitempty"`
	SecretType                                 *string                             `json:"secret_type,omitempty"`
	SecretTypeDisplayName                      *string                             `json:"secret_type_display_name,omitempty"`
	Secret                                     *string                             `json:"secret,omitempty"`
	Repository                                 *Repository                         `json:"repository,omitempty"`
	UpdatedAt                                  *Timestamp                          `json:"updated_at,omitempty"`
	IsBase64Encoded                            *bool                               `json:"is_base64_encoded,omitempty"`
	MultiRepo                                  *bool                               `json:"multi_repo,omitempty"`
	PubliclyLeaked                             *bool                               `json:"publicly_leaked,omitempty"`
	PushProtectionBypassed                     *bool                               `json:"push_protection_bypassed,omitempty"`
	PushProtectionBypassedBy                   *User                               `json:"push_protection_bypassed_by,omitempty"`
	PushProtectionBypassedAt                   *Timestamp                          `json:"push_protection_bypassed_at,omitempty"`
	ResolutionComment                          *string                             `json:"resolution_comment,omitempty"`
	PushProtectionBypassRequestComment         *string                             `json:"push_protection_bypass_request_comment,omitempty"`
	PushProtectionBypassRequestHTMLURL         *string                             `json:"push_protection_bypass_request_html_url,omitempty"`
	PushProtectionBypassRequestReviewer        *User                               `json:"push_protection_bypass_request_reviewer,omitempty"`
	PushProtectionBypassRequestReviewerComment *string                             `json:"push_protection_bypass_request_reviewer_comment,omitempty"`
	Validity                                   *string                             `json:"validity,omitempty"`
}

type SecretScanningAlertLocation struct {
	Type    *string                             `json:"type,omitempty"`
	Details *SecretScanningAlertLocationDetails `json:"details,omitempty"`
}

type SecretScanningAlertLocationDetails struct {
	Path                  *string `json:"path,omitempty"`
	Startline             *int    `json:"start_line,omitempty"`
	EndLine               *int    `json:"end_line,omitempty"`
	StartColumn           *int    `json:"start_column,omitempty"`
	EndColumn             *int    `json:"end_column,omitempty"`
	BlobSHA               *string `json:"blob_sha,omitempty"`
	BlobURL               *string `json:"blob_url,omitempty"`
	CommitSHA             *string `json:"commit_sha,omitempty"`
	CommitURL             *string `json:"commit_url,omitempty"`
	PullRequestCommentURL *string `json:"pull_request_comment_url,omitempty"`
}

type SecretScanningAlertListOptions struct {
	State string `url:"state,omitempty"`

	SecretType string `url:"secret_type,omitempty"`

	Resolution string `url:"resolution,omitempty"`

	Validity string `url:"validity,omitempty"`

	IsPubliclyLeaked bool `url:"is_publicly_leaked,omitempty"`

	IsMultiRepo bool `url:"is_multi_repo,omitempty"`

	Direction string `url:"direction,omitempty"`

	Sort string `url:"sort,omitempty"`

	ListCursorOptions

	ListOptions
}

type SecretScanningAlertUpdateOptions struct {
	State string `json:"state"`

	Resolution *string `json:"resolution,omitempty"`

	ResolutionComment *string `json:"resolution_comment,omitempty"`
}

type PushProtectionBypassRequest struct {
	Reason string `json:"reason"`

	PlaceholderID string `json:"placeholder_id"`
}

type PushProtectionBypass struct {
	Reason string `json:"reason"`

	ExpireAt *Timestamp `json:"expire_at"`

	TokenType string `json:"token_type"`
}

type SecretsScan struct {
	Type        string     `json:"type"`
	Status      string     `json:"status"`
	CompletedAt *Timestamp `json:"completed_at,omitempty"`
	StartedAt   *Timestamp `json:"started_at,omitempty"`
}

type CustomPatternBackfillScan struct {
	SecretsScan
	PatternSlug  *string `json:"pattern_slug,omitempty"`
	PatternScope *string `json:"pattern_scope,omitempty"`
}

type SecretScanningScanHistory struct {
	IncrementalScans []*SecretsScan `json:"incremental_scans,omitempty"`

	BackfillScans []*SecretsScan `json:"backfill_scans,omitempty"`

	PatternUpdateScans []*SecretsScan `json:"pattern_update_scans,omitempty"`

	CustomPatternBackfillScans []*CustomPatternBackfillScan `json:"custom_pattern_backfill_scans,omitempty"`
}

//meta:operation GET /enterprises/{enterprise}/secret-scanning/alerts
func (s *SecretScanningService) ListAlertsForEnterprise(ctx context.Context, enterprise string, opts *SecretScanningAlertListOptions) ([]*SecretScanningAlert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/secret-scanning/alerts
func (s *SecretScanningService) ListAlertsForOrg(ctx context.Context, org string, opts *SecretScanningAlertListOptions) ([]*SecretScanningAlert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/secret-scanning/alerts
func (s *SecretScanningService) ListAlertsForRepo(ctx context.Context, owner, repo string, opts *SecretScanningAlertListOptions) ([]*SecretScanningAlert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/secret-scanning/alerts/{alert_number}
func (s *SecretScanningService) GetAlert(ctx context.Context, owner, repo string, number int64) (*SecretScanningAlert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/secret-scanning/alerts/{alert_number}
func (s *SecretScanningService) UpdateAlert(ctx context.Context, owner, repo string, number int64, body *SecretScanningAlertUpdateOptions) (*SecretScanningAlert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/secret-scanning/alerts/{alert_number}/locations
func (s *SecretScanningService) ListLocationsForAlert(ctx context.Context, owner, repo string, number int64, opts *ListOptions) ([]*SecretScanningAlertLocation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/secret-scanning/push-protection-bypasses
func (s *SecretScanningService) CreatePushProtectionBypass(ctx context.Context, owner, repo string, body PushProtectionBypassRequest) (*PushProtectionBypass, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/secret-scanning/scan-history
func (s *SecretScanningService) GetScanHistory(ctx context.Context, owner, repo string) (*SecretScanningScanHistory, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
