package github

import (
	"context"
)

type CodeScanningService service

type Rule struct {
	ID                    *string  `json:"id,omitempty"`
	Severity              *string  `json:"severity,omitempty"`
	Description           *string  `json:"description,omitempty"`
	Name                  *string  `json:"name,omitempty"`
	SecuritySeverityLevel *string  `json:"security_severity_level,omitempty"`
	FullDescription       *string  `json:"full_description,omitempty"`
	Tags                  []string `json:"tags,omitempty"`
	Help                  *string  `json:"help,omitempty"`
}

type Location struct {
	Path        *string `json:"path,omitempty"`
	StartLine   *int    `json:"start_line,omitempty"`
	EndLine     *int    `json:"end_line,omitempty"`
	StartColumn *int    `json:"start_column,omitempty"`
	EndColumn   *int    `json:"end_column,omitempty"`
}

type Message struct {
	Text *string `json:"text,omitempty"`
}

type MostRecentInstance struct {
	Ref             *string   `json:"ref,omitempty"`
	AnalysisKey     *string   `json:"analysis_key,omitempty"`
	Category        *string   `json:"category,omitempty"`
	Environment     *string   `json:"environment,omitempty"`
	State           *string   `json:"state,omitempty"`
	CommitSHA       *string   `json:"commit_sha,omitempty"`
	Message         *Message  `json:"message,omitempty"`
	Location        *Location `json:"location,omitempty"`
	HTMLURL         *string   `json:"html_url,omitempty"`
	Classifications []string  `json:"classifications,omitempty"`
}

type Tool struct {
	Name    *string `json:"name,omitempty"`
	GUID    *string `json:"guid,omitempty"`
	Version *string `json:"version,omitempty"`
}

type Alert struct {
	Number             *int                  `json:"number,omitempty"`
	Repository         *Repository           `json:"repository,omitempty"`
	RuleID             *string               `json:"rule_id,omitempty"`
	RuleSeverity       *string               `json:"rule_severity,omitempty"`
	RuleDescription    *string               `json:"rule_description,omitempty"`
	Rule               *Rule                 `json:"rule,omitempty"`
	Tool               *Tool                 `json:"tool,omitempty"`
	CreatedAt          *Timestamp            `json:"created_at,omitempty"`
	UpdatedAt          *Timestamp            `json:"updated_at,omitempty"`
	FixedAt            *Timestamp            `json:"fixed_at,omitempty"`
	State              *string               `json:"state,omitempty"`
	ClosedBy           *User                 `json:"closed_by,omitempty"`
	ClosedAt           *Timestamp            `json:"closed_at,omitempty"`
	URL                *string               `json:"url,omitempty"`
	HTMLURL            *string               `json:"html_url,omitempty"`
	MostRecentInstance *MostRecentInstance   `json:"most_recent_instance,omitempty"`
	Instances          []*MostRecentInstance `json:"instances,omitempty"`
	DismissedBy        *User                 `json:"dismissed_by,omitempty"`
	DismissedAt        *Timestamp            `json:"dismissed_at,omitempty"`
	DismissedReason    *string               `json:"dismissed_reason,omitempty"`
	DismissedComment   *string               `json:"dismissed_comment,omitempty"`
	InstancesURL       *string               `json:"instances_url,omitempty"`
}

func (a *Alert) ID() int64 { _ = "STUB: not implemented"; return 0 }

type AlertInstancesListOptions struct {
	Ref string `url:"ref,omitempty"`

	ListOptions
}

type AlertListOptions struct {
	State string `url:"state,omitempty"`

	Ref string `url:"ref,omitempty"`

	Severity string `url:"severity,omitempty"`

	ToolName string `url:"tool_name,omitempty"`

	ToolGUID string `url:"tool_guid,omitempty"`

	Direction string `url:"direction,omitempty"`

	Sort string `url:"sort,omitempty"`

	ListCursorOptions

	ListOptions
}

type AnalysesListOptions struct {
	SarifID *string `url:"sarif_id,omitempty"`

	Ref *string `url:"ref,omitempty"`

	ListOptions
}

type CodeQLDatabase struct {
	ID          *int64     `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Language    *string    `json:"language,omitempty"`
	Uploader    *User      `json:"uploader,omitempty"`
	ContentType *string    `json:"content_type,omitempty"`
	Size        *int64     `json:"size,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
	URL         *string    `json:"url,omitempty"`
}

type ScanningAnalysis struct {
	ID           *int64     `json:"id,omitempty"`
	Ref          *string    `json:"ref,omitempty"`
	CommitSHA    *string    `json:"commit_sha,omitempty"`
	AnalysisKey  *string    `json:"analysis_key,omitempty"`
	Environment  *string    `json:"environment,omitempty"`
	Error        *string    `json:"error,omitempty"`
	Category     *string    `json:"category,omitempty"`
	CreatedAt    *Timestamp `json:"created_at,omitempty"`
	ResultsCount *int       `json:"results_count,omitempty"`
	RulesCount   *int       `json:"rules_count,omitempty"`
	URL          *string    `json:"url,omitempty"`
	SarifID      *string    `json:"sarif_id,omitempty"`
	Tool         *Tool      `json:"tool,omitempty"`
	Deletable    *bool      `json:"deletable,omitempty"`
	Warning      *string    `json:"warning,omitempty"`
}

type SarifAnalysis struct {
	CommitSHA   string     `json:"commit_sha"`
	Ref         string     `json:"ref"`
	Sarif       string     `json:"sarif"`
	CheckoutURI *string    `json:"checkout_uri,omitempty"`
	StartedAt   *Timestamp `json:"started_at,omitempty"`
	ToolName    *string    `json:"tool_name,omitempty"`
	Validate    *bool      `json:"validate,omitempty"`
}

type CodeScanningAlertState struct {
	State string `json:"state"`

	DismissedReason *string `json:"dismissed_reason,omitempty"`

	DismissedComment *string `json:"dismissed_comment,omitempty"`
}

type SarifID struct {
	ID  *string `json:"id,omitempty"`
	URL *string `json:"url,omitempty"`
}

//meta:operation GET /orgs/{org}/code-scanning/alerts
func (s *CodeScanningService) ListAlertsForOrg(ctx context.Context, org string, opts *AlertListOptions) ([]*Alert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/code-scanning/alerts
func (s *CodeScanningService) ListAlertsForRepo(ctx context.Context, owner, repo string, opts *AlertListOptions) ([]*Alert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/code-scanning/alerts/{alert_number}
func (s *CodeScanningService) GetAlert(ctx context.Context, owner, repo string, id int64) (*Alert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/code-scanning/alerts/{alert_number}
func (s *CodeScanningService) UpdateAlert(ctx context.Context, owner, repo string, id int64, body *CodeScanningAlertState) (*Alert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/code-scanning/alerts/{alert_number}/instances
func (s *CodeScanningService) ListAlertInstances(ctx context.Context, owner, repo string, id int64, opts *AlertInstancesListOptions) ([]*MostRecentInstance, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/code-scanning/sarifs
func (s *CodeScanningService) UploadSarif(ctx context.Context, owner, repo string, body SarifAnalysis) (*SarifID, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type SARIFUpload struct {
	ProcessingStatus *string `json:"processing_status,omitempty"`

	AnalysesURL *string `json:"analyses_url,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/code-scanning/sarifs/{sarif_id}
func (s *CodeScanningService) GetSARIF(ctx context.Context, owner, repo, sarifID string) (*SARIFUpload, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/code-scanning/analyses
func (s *CodeScanningService) ListAnalysesForRepo(ctx context.Context, owner, repo string, opts *AnalysesListOptions) ([]*ScanningAnalysis, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/code-scanning/analyses/{analysis_id}
func (s *CodeScanningService) GetAnalysis(ctx context.Context, owner, repo string, id int64) (*ScanningAnalysis, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type DeleteAnalysis struct {
	NextAnalysisURL *string `json:"next_analysis_url,omitempty"`

	ConfirmDeleteURL *string `json:"confirm_delete_url,omitempty"`
}

//meta:operation DELETE /repos/{owner}/{repo}/code-scanning/analyses/{analysis_id}
func (s *CodeScanningService) DeleteAnalysis(ctx context.Context, owner, repo string, id int64) (*DeleteAnalysis, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/code-scanning/codeql/databases
func (s *CodeScanningService) ListCodeQLDatabases(ctx context.Context, owner, repo string) ([]*CodeQLDatabase, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/code-scanning/codeql/databases/{language}
func (s *CodeScanningService) GetCodeQLDatabase(ctx context.Context, owner, repo, language string) (*CodeQLDatabase, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type DefaultSetupConfiguration struct {
	State      *string    `json:"state,omitempty"`
	Languages  []string   `json:"languages,omitempty"`
	QuerySuite *string    `json:"query_suite,omitempty"`
	UpdatedAt  *Timestamp `json:"updated_at,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/code-scanning/default-setup
func (s *CodeScanningService) GetDefaultSetupConfiguration(ctx context.Context, owner, repo string) (*DefaultSetupConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type UpdateDefaultSetupConfigurationOptions struct {
	State      string   `json:"state"`
	QuerySuite *string  `json:"query_suite,omitempty"`
	Languages  []string `json:"languages,omitempty"`
}

type UpdateDefaultSetupConfigurationResponse struct {
	RunID  *int64  `json:"run_id,omitempty"`
	RunURL *string `json:"run_url,omitempty"`
}

//meta:operation PATCH /repos/{owner}/{repo}/code-scanning/default-setup
func (s *CodeScanningService) UpdateDefaultSetupConfiguration(ctx context.Context, owner, repo string, body *UpdateDefaultSetupConfigurationOptions) (*UpdateDefaultSetupConfigurationResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
