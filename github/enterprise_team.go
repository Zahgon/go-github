package github

import (
	"context"
)

type EnterpriseTeam struct {
	ID                        int64     `json:"id"`
	URL                       string    `json:"url"`
	MemberURL                 string    `json:"member_url"`
	Name                      string    `json:"name"`
	Description               *string   `json:"description,omitempty"`
	HTMLURL                   string    `json:"html_url"`
	Slug                      string    `json:"slug"`
	CreatedAt                 Timestamp `json:"created_at"`
	UpdatedAt                 Timestamp `json:"updated_at"`
	GroupID                   string    `json:"group_id"`
	OrganizationSelectionType *string   `json:"organization_selection_type,omitempty"`
}

type EnterpriseTeamCreateOrUpdateRequest struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	OrganizationSelectionType *string `json:"organization_selection_type,omitempty"`

	GroupID *string `json:"group_id,omitempty"`
}

//meta:operation GET /enterprises/{enterprise}/teams
func (s *EnterpriseService) ListTeams(ctx context.Context, enterprise string, opts *ListOptions) ([]*EnterpriseTeam, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/teams
func (s *EnterpriseService) CreateTeam(ctx context.Context, enterprise string, body EnterpriseTeamCreateOrUpdateRequest) (*EnterpriseTeam, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/teams/{team_slug}
func (s *EnterpriseService) GetTeam(ctx context.Context, enterprise, teamSlug string) (*EnterpriseTeam, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/teams/{team_slug}
func (s *EnterpriseService) UpdateTeam(ctx context.Context, enterprise, teamSlug string, body EnterpriseTeamCreateOrUpdateRequest) (*EnterpriseTeam, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/teams/{team_slug}
func (s *EnterpriseService) DeleteTeam(ctx context.Context, enterprise, teamSlug string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /enterprises/{enterprise}/teams/{enterprise-team}/memberships
func (s *EnterpriseService) ListTeamMembers(ctx context.Context, enterprise, enterpriseTeam string, opts *ListOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/teams/{enterprise-team}/memberships/add
func (s *EnterpriseService) BulkAddTeamMembers(ctx context.Context, enterprise, enterpriseTeam string, username []string) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/teams/{enterprise-team}/memberships/remove
func (s *EnterpriseService) BulkRemoveTeamMembers(ctx context.Context, enterprise, enterpriseTeam string, username []string) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/teams/{enterprise-team}/memberships/{username}
func (s *EnterpriseService) GetTeamMembership(ctx context.Context, enterprise, enterpriseTeam, username string) (*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/teams/{enterprise-team}/memberships/{username}
func (s *EnterpriseService) AddTeamMember(ctx context.Context, enterprise, enterpriseTeam, username string) (*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/teams/{enterprise-team}/memberships/{username}
func (s *EnterpriseService) RemoveTeamMember(ctx context.Context, enterprise, enterpriseTeam, username string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /enterprises/{enterprise}/teams/{enterprise-team}/organizations
func (s *EnterpriseService) ListAssignments(ctx context.Context, enterprise, enterpriseTeam string, opts *ListOptions) ([]*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/teams/{enterprise-team}/organizations/add
func (s *EnterpriseService) AddMultipleAssignments(ctx context.Context, enterprise, enterpriseTeam string, organizationSlugs []string) ([]*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/teams/{enterprise-team}/organizations/remove
func (s *EnterpriseService) RemoveMultipleAssignments(ctx context.Context, enterprise, enterpriseTeam string, organizationSlugs []string) ([]*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/teams/{enterprise-team}/organizations/{org}
func (s *EnterpriseService) GetAssignment(ctx context.Context, enterprise, enterpriseTeam, org string) (*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/teams/{enterprise-team}/organizations/{org}
func (s *EnterpriseService) AddAssignment(ctx context.Context, enterprise, enterpriseTeam, org string) (*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/teams/{enterprise-team}/organizations/{org}
func (s *EnterpriseService) RemoveAssignment(ctx context.Context, enterprise, enterpriseTeam, org string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
