package github

import (
	"context"
)

type TeamsService service

type Team struct {
	ID          *int64  `json:"id,omitempty"`
	NodeID      *string `json:"node_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
	Slug        *string `json:"slug,omitempty"`

	Permission *string `json:"permission,omitempty"`

	Privacy *string `json:"privacy,omitempty"`

	NotificationSetting *string `json:"notification_setting,omitempty"`

	MembersCount    *int          `json:"members_count,omitempty"`
	ReposCount      *int          `json:"repos_count,omitempty"`
	Organization    *Organization `json:"organization,omitempty"`
	HTMLURL         *string       `json:"html_url,omitempty"`
	MembersURL      *string       `json:"members_url,omitempty"`
	RepositoriesURL *string       `json:"repositories_url,omitempty"`
	Parent          *Team         `json:"parent,omitempty"`

	LDAPDN *string `json:"ldap_dn,omitempty"`

	Permissions map[string]bool `json:"permissions,omitempty"`

	Assignment *string `json:"assignment,omitempty"`

	Type *string `json:"type,omitempty"`

	AccessSource *string `json:"access_source,omitempty"`
}

func (t Team) String() string { _ = "STUB: not implemented"; return "" }

type Invitation struct {
	ID     *int64  `json:"id,omitempty"`
	NodeID *string `json:"node_id,omitempty"`
	Login  *string `json:"login,omitempty"`
	Email  *string `json:"email,omitempty"`

	Role              *string    `json:"role,omitempty"`
	CreatedAt         *Timestamp `json:"created_at,omitempty"`
	Inviter           *User      `json:"inviter,omitempty"`
	TeamCount         *int       `json:"team_count,omitempty"`
	InvitationTeamURL *string    `json:"invitation_team_url,omitempty"`
	FailedAt          *Timestamp `json:"failed_at,omitempty"`
	FailedReason      *string    `json:"failed_reason,omitempty"`
}

func (i Invitation) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /orgs/{org}/teams
func (s *TeamsService) ListTeams(ctx context.Context, org string, opts *ListOptions) ([]*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /organizations/{organization_id}/team/{team_id}
func (s *TeamsService) GetTeamByID(ctx context.Context, orgID, teamID int64) (*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}
func (s *TeamsService) GetTeamBySlug(ctx context.Context, org, slug string) (*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type CreateTeamRequest struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Maintainers []string `json:"maintainers,omitempty"`

	RepoNames []string `json:"repo_names,omitempty"`

	Privacy *string `json:"privacy,omitempty"`

	NotificationSetting *string `json:"notification_setting,omitempty"`

	Permission *string `json:"permission,omitempty"`

	ParentTeamID *int64 `json:"parent_team_id,omitempty"`

	ParentTeamSlug *string `json:"parent_team_slug,omitempty"`
}

func (r CreateTeamRequest) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation POST /orgs/{org}/teams
func (s *TeamsService) CreateTeam(ctx context.Context, org string, body CreateTeamRequest) (*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type UpdateTeamRequest struct {
	Name *string `json:"name"`

	Description *string `json:"description,omitempty"`

	Privacy *string `json:"privacy,omitempty"`

	NotificationSetting *string `json:"notification_setting,omitempty"`

	Permission *string `json:"permission,omitempty"`

	ParentTeamID *int64 `json:"parent_team_id,omitempty"`

	ParentTeamSlug *string `json:"parent_team_slug,omitempty"`

	RemoveParentTeam bool `json:"-"`
}

func (r UpdateTeamRequest) String() string { _ = "STUB: not implemented"; return "" }

func (r UpdateTeamRequest) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PATCH /organizations/{organization_id}/team/{team_id}
func (s *TeamsService) UpdateTeamByID(ctx context.Context, orgID, teamID int64, body UpdateTeamRequest) (*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/teams/{team_slug}
func (s *TeamsService) UpdateTeamBySlug(ctx context.Context, org, slug string, body UpdateTeamRequest) (*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /organizations/{organization_id}/team/{team_id}
func (s *TeamsService) DeleteTeamByID(ctx context.Context, orgID, teamID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/teams/{team_slug}
func (s *TeamsService) DeleteTeamBySlug(ctx context.Context, org, slug string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /organizations/{organization_id}/team/{team_id}/teams
func (s *TeamsService) ListChildTeamsByParentID(ctx context.Context, orgID, teamID int64, opts *ListOptions) ([]*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/teams
func (s *TeamsService) ListChildTeamsByParentSlug(ctx context.Context, org, slug string, opts *ListOptions) ([]*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /organizations/{organization_id}/team/{team_id}/repos
func (s *TeamsService) ListTeamReposByID(ctx context.Context, orgID, teamID int64, opts *ListOptions) ([]*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/repos
func (s *TeamsService) ListTeamReposBySlug(ctx context.Context, org, slug string, opts *ListOptions) ([]*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /organizations/{organization_id}/team/{team_id}/repos/{owner}/{repo}
func (s *TeamsService) IsTeamRepoByID(ctx context.Context, orgID, teamID int64, owner, repo string) (*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/repos/{owner}/{repo}
func (s *TeamsService) IsTeamRepoBySlug(ctx context.Context, org, slug, owner, repo string) (*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type TeamAddTeamRepoOptions struct {
	Permission string `json:"permission,omitempty"`
}

//meta:operation PUT /organizations/{organization_id}/team/{team_id}/repos/{owner}/{repo}
func (s *TeamsService) AddTeamRepoByID(ctx context.Context, orgID, teamID int64, owner, repo string, body *TeamAddTeamRepoOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/teams/{team_slug}/repos/{owner}/{repo}
func (s *TeamsService) AddTeamRepoBySlug(ctx context.Context, org, slug, owner, repo string, body *TeamAddTeamRepoOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /organizations/{organization_id}/team/{team_id}/repos/{owner}/{repo}
func (s *TeamsService) RemoveTeamRepoByID(ctx context.Context, orgID, teamID int64, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/teams/{team_slug}/repos/{owner}/{repo}
func (s *TeamsService) RemoveTeamRepoBySlug(ctx context.Context, org, slug, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /user/teams
func (s *TeamsService) ListUserTeams(ctx context.Context, opts *ListOptions) ([]*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /organizations/{organization_id}/team/{team_id}/projects
func (s *TeamsService) ListTeamProjectsByID(ctx context.Context, orgID, teamID int64) ([]*ProjectV2, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/projects
func (s *TeamsService) ListTeamProjectsBySlug(ctx context.Context, org, slug string) ([]*ProjectV2, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /organizations/{organization_id}/team/{team_id}/projects/{project_id}
func (s *TeamsService) ReviewTeamProjectsByID(ctx context.Context, orgID, teamID, projectID int64) (*ProjectV2, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/projects/{project_id}
func (s *TeamsService) ReviewTeamProjectsBySlug(ctx context.Context, org, slug string, projectID int64) (*ProjectV2, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type TeamProjectOptions struct {
	Permission *string `json:"permission,omitempty"`
}

//meta:operation PUT /organizations/{organization_id}/team/{team_id}/projects/{project_id}
func (s *TeamsService) AddTeamProjectByID(ctx context.Context, orgID, teamID, projectID int64, body *TeamProjectOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/teams/{team_slug}/projects/{project_id}
func (s *TeamsService) AddTeamProjectBySlug(ctx context.Context, org, slug string, projectID int64, body *TeamProjectOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /organizations/{organization_id}/team/{team_id}/projects/{project_id}
func (s *TeamsService) RemoveTeamProjectByID(ctx context.Context, orgID, teamID, projectID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/teams/{team_slug}/projects/{project_id}
func (s *TeamsService) RemoveTeamProjectBySlug(ctx context.Context, org, slug string, projectID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ListIDPGroupsOptions struct {
	Query string `url:"q,omitempty"`

	ListCursorOptions
}

type IDPGroupList struct {
	Groups []*IDPGroup `json:"groups"`
}

type IDPGroup struct {
	GroupID          *string `json:"group_id,omitempty"`
	GroupName        *string `json:"group_name,omitempty"`
	GroupDescription *string `json:"group_description,omitempty"`
}

//meta:operation GET /orgs/{org}/team-sync/groups
func (s *TeamsService) ListIDPGroupsInOrganization(ctx context.Context, org string, opts *ListIDPGroupsOptions) (*IDPGroupList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /organizations/{organization_id}/team/{team_id}/team-sync/group-mappings
func (s *TeamsService) ListIDPGroupsForTeamByID(ctx context.Context, orgID, teamID int64) (*IDPGroupList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/team-sync/group-mappings
func (s *TeamsService) ListIDPGroupsForTeamBySlug(ctx context.Context, org, slug string) (*IDPGroupList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /organizations/{organization_id}/team/{team_id}/team-sync/group-mappings
func (s *TeamsService) CreateOrUpdateIDPGroupConnectionsByID(ctx context.Context, orgID, teamID int64, body IDPGroupList) (*IDPGroupList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/teams/{team_slug}/team-sync/group-mappings
func (s *TeamsService) CreateOrUpdateIDPGroupConnectionsBySlug(ctx context.Context, org, slug string, body IDPGroupList) (*IDPGroupList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ExternalGroupMember struct {
	MemberID    *int64  `json:"member_id,omitempty"`
	MemberLogin *string `json:"member_login,omitempty"`
	MemberName  *string `json:"member_name,omitempty"`
	MemberEmail *string `json:"member_email,omitempty"`
}

type ExternalGroupTeam struct {
	TeamID   *int64  `json:"team_id,omitempty"`
	TeamName *string `json:"team_name,omitempty"`
}

type ExternalGroup struct {
	GroupID   *int64                 `json:"group_id,omitempty"`
	GroupName *string                `json:"group_name,omitempty"`
	UpdatedAt *Timestamp             `json:"updated_at,omitempty"`
	Teams     []*ExternalGroupTeam   `json:"teams,omitempty"`
	Members   []*ExternalGroupMember `json:"members,omitempty"`
}

type ExternalGroupList struct {
	Groups []*ExternalGroup `json:"groups"`
}

type UpdateConnectedExternalGroupRequest struct {
	GroupID int64 `json:"group_id"`
}

//meta:operation GET /orgs/{org}/external-group/{group_id}
func (s *TeamsService) GetExternalGroup(ctx context.Context, org string, groupID int64) (*ExternalGroup, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ListExternalGroupsOptions struct {
	DisplayName *string `url:"display_name,omitempty"`

	ListOptions
}

//meta:operation GET /orgs/{org}/external-groups
func (s *TeamsService) ListExternalGroups(ctx context.Context, org string, opts *ListExternalGroupsOptions) (*ExternalGroupList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/external-groups
func (s *TeamsService) ListExternalGroupsForTeamBySlug(ctx context.Context, org, slug string) (*ExternalGroupList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/teams/{team_slug}/external-groups
func (s *TeamsService) UpdateConnectedExternalGroup(ctx context.Context, org, slug string, body UpdateConnectedExternalGroupRequest) (*ExternalGroup, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/teams/{team_slug}/external-groups
func (s *TeamsService) RemoveConnectedExternalGroup(ctx context.Context, org, slug string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
