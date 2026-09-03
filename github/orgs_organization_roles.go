package github

import (
	"context"
)

type OrganizationCustomRoles struct {
	TotalCount      *int             `json:"total_count,omitempty"`
	CustomRepoRoles []*CustomOrgRole `json:"roles,omitempty"`
}

type CustomOrgRole struct {
	ID          *int64        `json:"id,omitempty"`
	Name        *string       `json:"name,omitempty"`
	Description *string       `json:"description,omitempty"`
	Permissions []string      `json:"permissions,omitempty"`
	Org         *Organization `json:"organization,omitempty"`
	CreatedAt   *Timestamp    `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp    `json:"updated_at,omitempty"`
	Source      *string       `json:"source,omitempty"`
	BaseRole    *string       `json:"base_role,omitempty"`
}

type CreateCustomOrgRoleRequest struct {
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Permissions []string `json:"permissions"`
	BaseRole    *string  `json:"base_role,omitempty"`
}

type UpdateCustomOrgRoleRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	BaseRole    *string  `json:"base_role,omitempty"`
}

type OrganizationFineGrainedPermission struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

//meta:operation GET /orgs/{org}/organization-roles
func (s *OrganizationsService) ListRoles(ctx context.Context, org string) (*OrganizationCustomRoles, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/organization-roles/{role_id}
func (s *OrganizationsService) GetOrgRole(ctx context.Context, org string, roleID int64) (*CustomOrgRole, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/organization-roles
func (s *OrganizationsService) CreateCustomOrgRole(ctx context.Context, org string, body CreateCustomOrgRoleRequest) (*CustomOrgRole, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/organization-roles/{role_id}
func (s *OrganizationsService) UpdateCustomOrgRole(ctx context.Context, org string, roleID int64, body UpdateCustomOrgRoleRequest) (*CustomOrgRole, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/organization-roles/{role_id}
func (s *OrganizationsService) DeleteCustomOrgRole(ctx context.Context, org string, roleID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/organization-roles/teams/{team_slug}/{role_id}
func (s *OrganizationsService) AssignOrgRoleToTeam(ctx context.Context, org, teamSlug string, roleID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/organization-roles/teams/{team_slug}/{role_id}
func (s *OrganizationsService) RemoveOrgRoleFromTeam(ctx context.Context, org, teamSlug string, roleID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/organization-roles/users/{username}/{role_id}
func (s *OrganizationsService) AssignOrgRoleToUser(ctx context.Context, org, username string, roleID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/organization-roles/users/{username}/{role_id}
func (s *OrganizationsService) RemoveOrgRoleFromUser(ctx context.Context, org, username string, roleID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/organization-roles/{role_id}/teams
func (s *OrganizationsService) ListTeamsAssignedToOrgRole(ctx context.Context, org string, roleID int64, opts *ListOptions) ([]*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/organization-roles/{role_id}/users
func (s *OrganizationsService) ListUsersAssignedToOrgRole(ctx context.Context, org string, roleID int64, opts *ListOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/organization-fine-grained-permissions
func (s *OrganizationsService) ListFineGrainedPermissions(ctx context.Context, org string) ([]*OrganizationFineGrainedPermission, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
