package github

import (
	"context"
)

type OrganizationCustomRepoRoles struct {
	TotalCount      *int               `json:"total_count,omitempty"`
	CustomRepoRoles []*CustomRepoRoles `json:"custom_roles,omitempty"`
}

type CustomRepoRoles struct {
	ID          *int64        `json:"id,omitempty"`
	Name        *string       `json:"name,omitempty"`
	Description *string       `json:"description,omitempty"`
	BaseRole    *string       `json:"base_role,omitempty"`
	Permissions []string      `json:"permissions,omitempty"`
	Org         *Organization `json:"organization,omitempty"`
	CreatedAt   *Timestamp    `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp    `json:"updated_at,omitempty"`
}

type CreateCustomRepoRoleRequest struct {
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	BaseRole    string   `json:"base_role"`
	Permissions []string `json:"permissions"`
}

type UpdateCustomRepoRoleRequest struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	BaseRole    *string  `json:"base_role,omitempty"`
	Permissions []string `json:"permissions,omitzero"`
}

type RepoFineGrainedPermission struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

//meta:operation GET /orgs/{org}/custom-repository-roles
func (s *OrganizationsService) ListCustomRepoRoles(ctx context.Context, org string) (*OrganizationCustomRepoRoles, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/custom-repository-roles/{role_id}
func (s *OrganizationsService) GetCustomRepoRole(ctx context.Context, org string, roleID int64) (*CustomRepoRoles, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/custom-repository-roles
func (s *OrganizationsService) CreateCustomRepoRole(ctx context.Context, org string, body CreateCustomRepoRoleRequest) (*CustomRepoRoles, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/custom-repository-roles/{role_id}
func (s *OrganizationsService) UpdateCustomRepoRole(ctx context.Context, org string, roleID int64, body UpdateCustomRepoRoleRequest) (*CustomRepoRoles, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/custom-repository-roles/{role_id}
func (s *OrganizationsService) DeleteCustomRepoRole(ctx context.Context, org string, roleID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/repository-fine-grained-permissions
func (s *OrganizationsService) ListRepositoryFineGrainedPermissions(ctx context.Context, org string) ([]*RepoFineGrainedPermission, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
