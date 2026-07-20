package github

import (
	"context"
)

//meta:operation GET /orgs/{org}/security-managers
func (s *OrganizationsService) ListSecurityManagerTeams(ctx context.Context, org string) ([]*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/security-managers/teams/{team_slug}
func (s *OrganizationsService) AddSecurityManagerTeam(ctx context.Context, org, team string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/security-managers/teams/{team_slug}
func (s *OrganizationsService) RemoveSecurityManagerTeam(ctx context.Context, org, team string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
