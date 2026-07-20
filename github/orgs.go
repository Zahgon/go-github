package github

import (
	"context"
)

type OrganizationsService service

type Organization struct {
	Login                       *string    `json:"login,omitempty"`
	ID                          *int64     `json:"id,omitempty"`
	NodeID                      *string    `json:"node_id,omitempty"`
	AvatarURL                   *string    `json:"avatar_url,omitempty"`
	HTMLURL                     *string    `json:"html_url,omitempty"`
	Name                        *string    `json:"name,omitempty"`
	Company                     *string    `json:"company,omitempty"`
	Blog                        *string    `json:"blog,omitempty"`
	Location                    *string    `json:"location,omitempty"`
	Email                       *string    `json:"email,omitempty"`
	TwitterUsername             *string    `json:"twitter_username,omitempty"`
	Description                 *string    `json:"description,omitempty"`
	PublicRepos                 *int       `json:"public_repos,omitempty"`
	PublicGists                 *int       `json:"public_gists,omitempty"`
	Followers                   *int       `json:"followers,omitempty"`
	Following                   *int       `json:"following,omitempty"`
	CreatedAt                   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt                   *Timestamp `json:"updated_at,omitempty"`
	ArchivedAt                  *Timestamp `json:"archived_at,omitempty"`
	TotalPrivateRepos           *int64     `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos           *int64     `json:"owned_private_repos,omitempty"`
	PrivateGists                *int       `json:"private_gists,omitempty"`
	DiskUsage                   *int       `json:"disk_usage,omitempty"`
	Collaborators               *int       `json:"collaborators,omitempty"`
	BillingEmail                *string    `json:"billing_email,omitempty"`
	Type                        *string    `json:"type,omitempty"`
	Plan                        *Plan      `json:"plan,omitempty"`
	TwoFactorRequirementEnabled *bool      `json:"two_factor_requirement_enabled,omitempty"`
	IsVerified                  *bool      `json:"is_verified,omitempty"`
	HasOrganizationProjects     *bool      `json:"has_organization_projects,omitempty"`
	HasRepositoryProjects       *bool      `json:"has_repository_projects,omitempty"`

	DefaultRepoPermission *string `json:"default_repository_permission,omitempty"`

	DefaultRepoSettings *string `json:"default_repository_settings,omitempty"`

	MembersCanCreateRepos *bool `json:"members_can_create_repositories,omitempty"`

	MembersCanCreatePublicRepos   *bool `json:"members_can_create_public_repositories,omitempty"`
	MembersCanCreatePrivateRepos  *bool `json:"members_can_create_private_repositories,omitempty"`
	MembersCanCreateInternalRepos *bool `json:"members_can_create_internal_repositories,omitempty"`

	MembersCanForkPrivateRepos *bool `json:"members_can_fork_private_repositories,omitempty"`

	DeployKeysEnabledForRepositories *bool `json:"deploy_keys_enabled_for_repositories,omitempty"`

	MembersAllowedRepositoryCreationType *string `json:"members_allowed_repository_creation_type,omitempty"`

	MembersCanCreatePages *bool `json:"members_can_create_pages,omitempty"`

	MembersCanCreatePublicPages *bool `json:"members_can_create_public_pages,omitempty"`

	MembersCanCreatePrivatePages *bool `json:"members_can_create_private_pages,omitempty"`

	WebCommitSignoffRequired *bool `json:"web_commit_signoff_required,omitempty"`

	AdvancedSecurityEnabledForNewRepos *bool `json:"advanced_security_enabled_for_new_repositories,omitempty"`

	DependabotAlertsEnabledForNewRepos *bool `json:"dependabot_alerts_enabled_for_new_repositories,omitempty"`

	DependabotSecurityUpdatesEnabledForNewRepos *bool `json:"dependabot_security_updates_enabled_for_new_repositories,omitempty"`

	DependencyGraphEnabledForNewRepos *bool `json:"dependency_graph_enabled_for_new_repositories,omitempty"`

	SecretScanningEnabledForNewRepos *bool `json:"secret_scanning_enabled_for_new_repositories,omitempty"`

	SecretScanningPushProtectionEnabledForNewRepos *bool `json:"secret_scanning_push_protection_enabled_for_new_repositories,omitempty"`

	SecretScanningValidityChecksEnabled *bool `json:"secret_scanning_validity_checks_enabled,omitempty"`

	SecretScanningPushProtectionCustomLinkEnabled *bool `json:"secret_scanning_push_protection_custom_link_enabled,omitempty"`

	SecretScanningPushProtectionCustomLink *string `json:"secret_scanning_push_protection_custom_link,omitempty"`

	MembersCanDeleteRepositories *bool `json:"members_can_delete_repositories,omitempty"`

	MembersCanChangeRepoVisibility *bool `json:"members_can_change_repo_visibility,omitempty"`

	MembersCanInviteOutsideCollaborators *bool `json:"members_can_invite_outside_collaborators,omitempty"`

	MembersCanDeleteIssues *bool `json:"members_can_delete_issues,omitempty"`

	DisplayCommenterFullNameSettingEnabled *bool `json:"display_commenter_full_name_setting_enabled,omitempty"`

	ReadersCanCreateDiscussions *bool `json:"readers_can_create_discussions,omitempty"`

	MembersCanCreateTeams *bool `json:"members_can_create_teams,omitempty"`

	MembersCanViewDependencyInsights *bool `json:"members_can_view_dependency_insights,omitempty"`

	DefaultRepositoryBranch *string `json:"default_repository_branch,omitempty"`

	URL              *string `json:"url,omitempty"`
	EventsURL        *string `json:"events_url,omitempty"`
	HooksURL         *string `json:"hooks_url,omitempty"`
	IssuesURL        *string `json:"issues_url,omitempty"`
	MembersURL       *string `json:"members_url,omitempty"`
	PublicMembersURL *string `json:"public_members_url,omitempty"`
	ReposURL         *string `json:"repos_url,omitempty"`
}

type OrganizationInstallations struct {
	TotalCount    *int            `json:"total_count,omitempty"`
	Installations []*Installation `json:"installations,omitempty"`
}

func (o Organization) String() string { _ = "STUB: not implemented"; return "" }

type Plan struct {
	Name          *string `json:"name,omitempty"`
	Space         *int    `json:"space,omitempty"`
	Collaborators *int    `json:"collaborators,omitempty"`
	PrivateRepos  *int64  `json:"private_repos,omitempty"`
	FilledSeats   *int    `json:"filled_seats,omitempty"`
	Seats         *int    `json:"seats,omitempty"`
}

func (p Plan) String() string { _ = "STUB: not implemented"; return "" }

type OrganizationsListOptions struct {
	Since int64 `url:"since,omitempty"`

	PerPage int `url:"per_page,omitempty"`
}

//meta:operation GET /organizations
func (s *OrganizationsService) ListAll(ctx context.Context, opts *OrganizationsListOptions) ([]*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/orgs
//meta:operation GET /users/{username}/orgs
func (s *OrganizationsService) List(ctx context.Context, user string, opts *ListOptions) ([]*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}
func (s *OrganizationsService) Get(ctx context.Context, org string) (*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /organizations/{organization_id}
func (s *OrganizationsService) GetByID(ctx context.Context, id int64) (*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}
func (s *OrganizationsService) Edit(ctx context.Context, name string, body *Organization) (*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}
func (s *OrganizationsService) Delete(ctx context.Context, org string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/installations
func (s *OrganizationsService) ListInstallations(ctx context.Context, org string, opts *ListOptions) (*OrganizationInstallations, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
