package github

import (
	"context"
)

type AppsService service

type App struct {
	ID                 *int64                   `json:"id,omitempty"`
	Slug               *string                  `json:"slug,omitempty"`
	ClientID           *string                  `json:"client_id,omitempty"`
	NodeID             *string                  `json:"node_id,omitempty"`
	Owner              *User                    `json:"owner,omitempty"`
	Name               *string                  `json:"name,omitempty"`
	Description        *string                  `json:"description,omitempty"`
	ExternalURL        *string                  `json:"external_url,omitempty"`
	HTMLURL            *string                  `json:"html_url,omitempty"`
	CreatedAt          *Timestamp               `json:"created_at,omitempty"`
	UpdatedAt          *Timestamp               `json:"updated_at,omitempty"`
	Permissions        *InstallationPermissions `json:"permissions,omitempty"`
	Events             []string                 `json:"events,omitempty"`
	InstallationsCount *int                     `json:"installations_count,omitempty"`
}

type InstallationToken struct {
	Token        *string                  `json:"token,omitempty"`
	ExpiresAt    *Timestamp               `json:"expires_at,omitempty"`
	Permissions  *InstallationPermissions `json:"permissions,omitempty"`
	Repositories []*Repository            `json:"repositories,omitempty"`
}

type InstallationTokenOptions struct {
	RepositoryIDs []int64 `json:"repository_ids,omitempty"`

	Repositories []string `json:"repositories,omitempty"`

	Permissions *InstallationPermissions `json:"permissions,omitempty"`
}

type InstallationTokenListRepoOptions struct {
	RepositoryIDs []int64 `json:"repository_ids"`

	Repositories []string `json:"repositories,omitempty"`

	Permissions *InstallationPermissions `json:"permissions,omitempty"`
}

type InstallationPermissions struct {
	Actions                                 *string `json:"actions,omitempty"`
	ActionsVariables                        *string `json:"actions_variables,omitempty"`
	Administration                          *string `json:"administration,omitempty"`
	Attestations                            *string `json:"attestations,omitempty"`
	Blocking                                *string `json:"blocking,omitempty"`
	Checks                                  *string `json:"checks,omitempty"`
	Codespaces                              *string `json:"codespaces,omitempty"`
	CodespacesLifecycleAdmin                *string `json:"codespaces_lifecycle_admin,omitempty"`
	CodespacesMetadata                      *string `json:"codespaces_metadata,omitempty"`
	CodespacesSecrets                       *string `json:"codespaces_secrets,omitempty"`
	CodespacesUserSecrets                   *string `json:"codespaces_user_secrets,omitempty"`
	Contents                                *string `json:"contents,omitempty"`
	ContentReferences                       *string `json:"content_references,omitempty"`
	CopilotMessages                         *string `json:"copilot_messages,omitempty"`
	DependabotSecrets                       *string `json:"dependabot_secrets,omitempty"`
	Deployments                             *string `json:"deployments,omitempty"`
	Discussions                             *string `json:"discussions,omitempty"`
	Emails                                  *string `json:"emails,omitempty"`
	EnterpriseAIControls                    *string `json:"enterprise_ai_controls,omitempty"`
	EnterpriseCopilotMetrics                *string `json:"enterprise_copilot_metrics,omitempty"`
	EnterpriseCredentials                   *string `json:"enterprise_credentials,omitempty"`
	EnterpriseCustomEnterpriseRoles         *string `json:"enterprise_custom_enterprise_roles,omitempty"`
	EnterpriseCustomOrgRoles                *string `json:"enterprise_custom_org_roles,omitempty"`
	EnterpriseCustomProperties              *string `json:"enterprise_custom_properties,omitempty"`
	EnterpriseCustomPropertiesForOrgs       *string `json:"enterprise_custom_properties_for_organizations,omitempty"`
	EnterpriseOrganizations                 *string `json:"enterprise_organizations,omitempty"`
	EnterpriseOrganizationInstallations     *string `json:"enterprise_organization_installations,omitempty"`
	EnterpriseOrgInstallationRepos          *string `json:"enterprise_organization_installation_repositories,omitempty"`
	EnterprisePeople                        *string `json:"enterprise_people,omitempty"`
	EnterpriseSSO                           *string `json:"enterprise_sso,omitempty"`
	EnterpriseTeams                         *string `json:"enterprise_teams,omitempty"`
	Environments                            *string `json:"environments,omitempty"`
	Followers                               *string `json:"followers,omitempty"`
	Gists                                   *string `json:"gists,omitempty"`
	GitSigningSSHPublicKeys                 *string `json:"git_signing_ssh_public_keys,omitempty"`
	GPGKeys                                 *string `json:"gpg_keys,omitempty"`
	InteractionLimits                       *string `json:"interaction_limits,omitempty"`
	Issues                                  *string `json:"issues,omitempty"`
	Keys                                    *string `json:"keys,omitempty"`
	Metadata                                *string `json:"metadata,omitempty"`
	Members                                 *string `json:"members,omitempty"`
	MergeQueues                             *string `json:"merge_queues,omitempty"`
	OrganizationActionsVariables            *string `json:"organization_actions_variables,omitempty"`
	OrganizationAdministration              *string `json:"organization_administration,omitempty"`
	OrganizationAnnouncementBanners         *string `json:"organization_announcement_banners,omitempty"`
	OrganizationAPIInsights                 *string `json:"organization_api_insights,omitempty"`
	OrganizationCodespaces                  *string `json:"organization_codespaces,omitempty"`
	OrganizationCodespacesSecrets           *string `json:"organization_codespaces_secrets,omitempty"`
	OrganizationCodespacesSettings          *string `json:"organization_codespaces_settings,omitempty"`
	OrganizationCopilotMetrics              *string `json:"organization_copilot_metrics,omitempty"`
	OrganizationCopilotSeatManagement       *string `json:"organization_copilot_seat_management,omitempty"`
	OrganizationCustomProperties            *string `json:"organization_custom_properties,omitempty"`
	OrganizationCustomRoles                 *string `json:"organization_custom_roles,omitempty"`
	OrganizationCustomOrgRoles              *string `json:"organization_custom_org_roles,omitempty"`
	OrganizationDependabotSecrets           *string `json:"organization_dependabot_secrets,omitempty"`
	OrganizationEvents                      *string `json:"organization_events,omitempty"`
	OrganizationHooks                       *string `json:"organization_hooks,omitempty"`
	OrganizationKnowledgeBases              *string `json:"organization_knowledge_bases,omitempty"`
	OrganizationPackages                    *string `json:"organization_packages,omitempty"`
	OrganizationPersonalAccessTokens        *string `json:"organization_personal_access_tokens,omitempty"`
	OrganizationPersonalAccessTokenRequests *string `json:"organization_personal_access_token_requests,omitempty"`
	OrganizationPlan                        *string `json:"organization_plan,omitempty"`
	OrganizationPreReceiveHooks             *string `json:"organization_pre_receive_hooks,omitempty"`
	OrganizationProjects                    *string `json:"organization_projects,omitempty"`
	OrganizationSecrets                     *string `json:"organization_secrets,omitempty"`
	OrganizationSelfHostedRunners           *string `json:"organization_self_hosted_runners,omitempty"`
	OrganizationUserBlocking                *string `json:"organization_user_blocking,omitempty"`
	Packages                                *string `json:"packages,omitempty"`
	Pages                                   *string `json:"pages,omitempty"`
	Plan                                    *string `json:"plan,omitempty"`
	Profile                                 *string `json:"profile,omitempty"`
	PullRequests                            *string `json:"pull_requests,omitempty"`
	RepositoryAdvisories                    *string `json:"repository_advisories,omitempty"`
	RepositoryCustomProperties              *string `json:"repository_custom_properties,omitempty"`
	RepositoryHooks                         *string `json:"repository_hooks,omitempty"`
	RepositoryProjects                      *string `json:"repository_projects,omitempty"`
	RepositoryPreReceiveHooks               *string `json:"repository_pre_receive_hooks,omitempty"`
	Secrets                                 *string `json:"secrets,omitempty"`
	SecretScanningAlerts                    *string `json:"secret_scanning_alerts,omitempty"`
	SecurityEvents                          *string `json:"security_events,omitempty"`
	SingleFile                              *string `json:"single_file,omitempty"`
	Starring                                *string `json:"starring,omitempty"`
	Statuses                                *string `json:"statuses,omitempty"`
	TeamDiscussions                         *string `json:"team_discussions,omitempty"`
	UserEvents                              *string `json:"user_events,omitempty"`
	VulnerabilityAlerts                     *string `json:"vulnerability_alerts,omitempty"`
	Watching                                *string `json:"watching,omitempty"`
	Workflows                               *string `json:"workflows,omitempty"`
}

type InstallationRequest struct {
	ID        *int64     `json:"id,omitempty"`
	NodeID    *string    `json:"node_id,omitempty"`
	Account   *User      `json:"account,omitempty"`
	Requester *User      `json:"requester,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
}

type Installation struct {
	ID                     *int64                   `json:"id,omitempty"`
	NodeID                 *string                  `json:"node_id,omitempty"`
	ClientID               *string                  `json:"client_id,omitempty"`
	AppID                  *int64                   `json:"app_id,omitempty"`
	AppSlug                *string                  `json:"app_slug,omitempty"`
	TargetID               *int64                   `json:"target_id,omitempty"`
	Account                *User                    `json:"account,omitempty"`
	AccessTokensURL        *string                  `json:"access_tokens_url,omitempty"`
	RepositoriesURL        *string                  `json:"repositories_url,omitempty"`
	HTMLURL                *string                  `json:"html_url,omitempty"`
	TargetType             *string                  `json:"target_type,omitempty"`
	SingleFileName         *string                  `json:"single_file_name,omitempty"`
	RepositorySelection    *string                  `json:"repository_selection,omitempty"`
	Events                 []string                 `json:"events,omitempty"`
	SingleFilePaths        []string                 `json:"single_file_paths,omitempty"`
	Permissions            *InstallationPermissions `json:"permissions,omitempty"`
	CreatedAt              *Timestamp               `json:"created_at,omitempty"`
	UpdatedAt              *Timestamp               `json:"updated_at,omitempty"`
	HasMultipleSingleFiles *bool                    `json:"has_multiple_single_files,omitempty"`
	SuspendedBy            *User                    `json:"suspended_by,omitempty"`
	SuspendedAt            *Timestamp               `json:"suspended_at,omitempty"`
}

type Attachment struct {
	ID    *int64  `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Body  *string `json:"body,omitempty"`
}

type ContentReference struct {
	ID        *int64  `json:"id,omitempty"`
	NodeID    *string `json:"node_id,omitempty"`
	Reference *string `json:"reference,omitempty"`
}

func (i Installation) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /app
//meta:operation GET /apps/{app_slug}
func (s *AppsService) Get(ctx context.Context, appSlug string) (*App, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /app/installation-requests
func (s *AppsService) ListInstallationRequests(ctx context.Context, opts *ListOptions) ([]*InstallationRequest, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /app/installations
func (s *AppsService) ListInstallations(ctx context.Context, opts *ListOptions) ([]*Installation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /app/installations/{installation_id}
func (s *AppsService) GetInstallation(ctx context.Context, id int64) (*Installation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/installations
func (s *AppsService) ListUserInstallations(ctx context.Context, opts *ListOptions) ([]*Installation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /app/installations/{installation_id}/suspended
func (s *AppsService) SuspendInstallation(ctx context.Context, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /app/installations/{installation_id}/suspended
func (s *AppsService) UnsuspendInstallation(ctx context.Context, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /app/installations/{installation_id}
func (s *AppsService) DeleteInstallation(ctx context.Context, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /app/installations/{installation_id}/access_tokens
func (s *AppsService) CreateInstallationToken(ctx context.Context, id int64, body *InstallationTokenOptions) (*InstallationToken, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /app/installations/{installation_id}/access_tokens
func (s *AppsService) CreateInstallationTokenListRepos(ctx context.Context, id int64, body *InstallationTokenListRepoOptions) (*InstallationToken, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/content_references/{content_reference_id}/attachments
func (s *AppsService) CreateAttachment(ctx context.Context, contentReferenceID int64, title, body string) (*Attachment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/installation
func (s *AppsService) GetOrganizationInstallation(ctx context.Context, org string) (*Installation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/installation
func (s *AppsService) GetEnterpriseInstallation(ctx context.Context, enterprise string) (*Installation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/installation
func (s *AppsService) GetRepositoryInstallation(ctx context.Context, owner, repo string) (*Installation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repositories/{repository_id}/installation
func (s *AppsService) GetRepositoryInstallationByID(ctx context.Context, id int64) (*Installation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/installation
func (s *AppsService) GetUserInstallation(ctx context.Context, user string) (*Installation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *AppsService) getInstallation(ctx context.Context, url string) (*Installation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
