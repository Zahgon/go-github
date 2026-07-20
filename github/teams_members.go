package github

import (
	"context"
)

type TeamListTeamMembersOptions struct {
	Role string `url:"role,omitempty"`

	ListOptions
}

//meta:operation GET /organizations/{organization_id}/team/{team_id}/members
func (s *TeamsService) ListTeamMembersByID(ctx context.Context, orgID, teamID int64, opts *TeamListTeamMembersOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/members
func (s *TeamsService) ListTeamMembersBySlug(ctx context.Context, org, slug string, opts *TeamListTeamMembersOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /organizations/{organization_id}/team/{team_id}/memberships/{username}
func (s *TeamsService) GetTeamMembershipByID(ctx context.Context, orgID, teamID int64, user string) (*Membership, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/memberships/{username}
func (s *TeamsService) GetTeamMembershipBySlug(ctx context.Context, org, slug, user string) (*Membership, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type TeamAddTeamMembershipOptions struct {
	Role string `json:"role,omitempty"`
}

//meta:operation PUT /organizations/{organization_id}/team/{team_id}/memberships/{username}
func (s *TeamsService) AddTeamMembershipByID(ctx context.Context, orgID, teamID int64, user string, body *TeamAddTeamMembershipOptions) (*Membership, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/teams/{team_slug}/memberships/{username}
func (s *TeamsService) AddTeamMembershipBySlug(ctx context.Context, org, slug, user string, body *TeamAddTeamMembershipOptions) (*Membership, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /organizations/{organization_id}/team/{team_id}/memberships/{username}
func (s *TeamsService) RemoveTeamMembershipByID(ctx context.Context, orgID, teamID int64, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/teams/{team_slug}/memberships/{username}
func (s *TeamsService) RemoveTeamMembershipBySlug(ctx context.Context, org, slug, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /organizations/{organization_id}/team/{team_id}/invitations
func (s *TeamsService) ListPendingTeamInvitationsByID(ctx context.Context, orgID, teamID int64, opts *ListOptions) ([]*Invitation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/invitations
func (s *TeamsService) ListPendingTeamInvitationsBySlug(ctx context.Context, org, slug string, opts *ListOptions) ([]*Invitation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
