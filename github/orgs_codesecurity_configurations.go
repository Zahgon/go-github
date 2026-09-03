package github

import (
	"context"
)

type DependencyGraphAutosubmitActionOptions struct {
	LabeledRunners *bool `json:"labeled_runners,omitempty"`
}

type CodeScanningOptions struct {
	AllowAdvanced *bool `json:"allow_advanced,omitempty"`
}

type CodeScanningDefaultSetupOptions struct {
	RunnerType  string  `json:"runner_type"`
	RunnerLabel *string `json:"runner_label,omitempty"`
}

type RepositoryAttachment struct {
	Status     *string     `json:"status"`
	Repository *Repository `json:"repository"`
}

type SecretScanningDelegatedBypassOptions struct {
	Reviewers []*BypassReviewer `json:"reviewers,omitzero"`
}

type BypassReviewer struct {
	ReviewerID              int64  `json:"reviewer_id"`
	ReviewerType            string `json:"reviewer_type"`
	SecurityConfigurationID *int64 `json:"security_configuration_id,omitempty"`
}

type CodeSecurityConfiguration struct {
	ID                                     *int64                                  `json:"id,omitempty"`
	TargetType                             *string                                 `json:"target_type,omitempty"`
	Name                                   string                                  `json:"name"`
	Description                            string                                  `json:"description"`
	AdvancedSecurity                       *string                                 `json:"advanced_security,omitempty"`
	DependencyGraph                        *string                                 `json:"dependency_graph,omitempty"`
	DependencyGraphAutosubmitAction        *string                                 `json:"dependency_graph_autosubmit_action,omitempty"`
	DependencyGraphAutosubmitActionOptions *DependencyGraphAutosubmitActionOptions `json:"dependency_graph_autosubmit_action_options,omitempty"`
	DependabotAlerts                       *string                                 `json:"dependabot_alerts,omitempty"`
	DependabotDelegatedAlertDismissal      *string                                 `json:"dependabot_delegated_alert_dismissal,omitempty"`
	DependabotSecurityUpdates              *string                                 `json:"dependabot_security_updates,omitempty"`
	CodeScanningDefaultSetup               *string                                 `json:"code_scanning_default_setup,omitempty"`
	CodeScanningDefaultSetupOptions        *CodeScanningDefaultSetupOptions        `json:"code_scanning_default_setup_options,omitempty"`
	CodeScanningDelegatedAlertDismissal    *string                                 `json:"code_scanning_delegated_alert_dismissal,omitempty"`
	CodeScanningOptions                    *CodeScanningOptions                    `json:"code_scanning_options,omitempty"`
	CodeSecurity                           *string                                 `json:"code_security,omitempty"`
	SecretScanning                         *string                                 `json:"secret_scanning,omitempty"`
	SecretScanningPushProtection           *string                                 `json:"secret_scanning_push_protection,omitempty"`
	SecretScanningDelegatedBypass          *string                                 `json:"secret_scanning_delegated_bypass,omitempty"`
	SecretScanningDelegatedBypassOptions   *SecretScanningDelegatedBypassOptions   `json:"secret_scanning_delegated_bypass_options,omitempty"`
	SecretScanningValidityChecks           *string                                 `json:"secret_scanning_validity_checks,omitempty"`
	SecretScanningNonProviderPatterns      *string                                 `json:"secret_scanning_non_provider_patterns,omitempty"`
	SecretScanningGenericSecrets           *string                                 `json:"secret_scanning_generic_secrets,omitempty"`
	SecretScanningDelegatedAlertDismissal  *string                                 `json:"secret_scanning_delegated_alert_dismissal,omitempty"`
	SecretScanningExtendedMetadata         *string                                 `json:"secret_scanning_extended_metadata,omitempty"`
	SecretProtection                       *string                                 `json:"secret_protection,omitempty"`
	PrivateVulnerabilityReporting          *string                                 `json:"private_vulnerability_reporting,omitempty"`
	Enforcement                            *string                                 `json:"enforcement,omitempty"`
	URL                                    *string                                 `json:"url,omitempty"`
	HTMLURL                                *string                                 `json:"html_url,omitempty"`
	CreatedAt                              *Timestamp                              `json:"created_at,omitempty"`
	UpdatedAt                              *Timestamp                              `json:"updated_at,omitempty"`
}

type CodeSecurityConfigurationWithDefaultForNewRepos struct {
	Configuration      *CodeSecurityConfiguration `json:"configuration"`
	DefaultForNewRepos *string                    `json:"default_for_new_repos,omitempty"`
}

type RepositoryCodeSecurityConfiguration struct {
	State         *string                    `json:"state,omitempty"`
	Configuration *CodeSecurityConfiguration `json:"configuration,omitempty"`
}

type ListOrgCodeSecurityConfigurationOptions struct {
	Before string `url:"before,omitempty"`

	After string `url:"after,omitempty"`

	PerPage int `url:"per_page,omitempty"`

	TargetType string `url:"target_type,omitempty"`
}

type ListCodeSecurityConfigurationRepositoriesOptions struct {
	Before string `url:"before,omitempty"`

	After string `url:"after,omitempty"`

	PerPage int `url:"per_page,omitempty"`

	Status string `url:"status,omitempty"`
}

//meta:operation GET /orgs/{org}/code-security/configurations
func (s *OrganizationsService) ListCodeSecurityConfigurations(ctx context.Context, org string, opts *ListOrgCodeSecurityConfigurationOptions) ([]*CodeSecurityConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/code-security/configurations
func (s *OrganizationsService) CreateCodeSecurityConfiguration(ctx context.Context, org string, body CodeSecurityConfiguration) (*CodeSecurityConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/code-security/configurations/defaults
func (s *OrganizationsService) ListDefaultCodeSecurityConfigurations(ctx context.Context, org string) ([]*CodeSecurityConfigurationWithDefaultForNewRepos, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/code-security/configurations/detach
func (s *OrganizationsService) DetachCodeSecurityConfigurationsFromRepositories(ctx context.Context, org string, repoIDs []int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/code-security/configurations/{configuration_id}
func (s *OrganizationsService) GetCodeSecurityConfiguration(ctx context.Context, org string, configurationID int64) (*CodeSecurityConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/code-security/configurations/{configuration_id}
func (s *OrganizationsService) UpdateCodeSecurityConfiguration(ctx context.Context, org string, configurationID int64, body CodeSecurityConfiguration) (*CodeSecurityConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/code-security/configurations/{configuration_id}
func (s *OrganizationsService) DeleteCodeSecurityConfiguration(ctx context.Context, org string, configurationID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /orgs/{org}/code-security/configurations/{configuration_id}/attach
func (s *OrganizationsService) AttachCodeSecurityConfigurationToRepositories(ctx context.Context, org string, configurationID int64, scope string, repoIDs []int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/code-security/configurations/{configuration_id}/defaults
func (s *OrganizationsService) SetDefaultCodeSecurityConfiguration(ctx context.Context, org string, configurationID int64, newReposParam string) (*CodeSecurityConfigurationWithDefaultForNewRepos, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/code-security/configurations/{configuration_id}/repositories
func (s *OrganizationsService) ListCodeSecurityConfigurationRepositories(ctx context.Context, org string, configurationID int64, opts *ListCodeSecurityConfigurationRepositoriesOptions) ([]*RepositoryAttachment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/code-security-configuration
func (s *OrganizationsService) GetCodeSecurityConfigurationForRepository(ctx context.Context, org, repo string) (*RepositoryCodeSecurityConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
