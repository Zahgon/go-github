package github

import (
	"context"
)

type Membership struct {
	URL *string `json:"url,omitempty"`

	State *string `json:"state,omitempty"`

	Role *string `json:"role,omitempty"`

	OrganizationURL *string `json:"organization_url,omitempty"`

	Organization *Organization `json:"organization,omitempty"`

	User *User `json:"user,omitempty"`
}

func (m Membership) String() string { _ = "STUB: not implemented"; return "" }

type ListMembersOptions struct {
	PublicOnly bool `url:"-"`

	Filter string `url:"filter,omitempty"`

	Role string `url:"role,omitempty"`

	ListOptions
}

//meta:operation GET /orgs/{org}/members
//meta:operation GET /orgs/{org}/public_members
func (s *OrganizationsService) ListMembers(ctx context.Context, org string, opts *ListMembersOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/members/{username}
func (s *OrganizationsService) IsMember(ctx context.Context, org, user string) (bool, *Response, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

//meta:operation GET /orgs/{org}/public_members/{username}
func (s *OrganizationsService) IsPublicMember(ctx context.Context, org, user string) (bool, *Response, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

//meta:operation DELETE /orgs/{org}/members/{username}
func (s *OrganizationsService) RemoveMember(ctx context.Context, org, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/invitations/{invitation_id}
func (s *OrganizationsService) CancelInvite(ctx context.Context, org string, invitationID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/public_members/{username}
func (s *OrganizationsService) PublicizeMembership(ctx context.Context, org, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/public_members/{username}
func (s *OrganizationsService) ConcealMembership(ctx context.Context, org, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ListOrgMembershipsOptions struct {
	State string `url:"state,omitempty"`

	ListOptions
}

//meta:operation GET /user/memberships/orgs
func (s *OrganizationsService) ListOrgMemberships(ctx context.Context, opts *ListOrgMembershipsOptions) ([]*Membership, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/memberships/{username}
//meta:operation GET /user/memberships/orgs/{org}
func (s *OrganizationsService) GetOrgMembership(ctx context.Context, user, org string) (*Membership, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/memberships/{username}
//meta:operation PATCH /user/memberships/orgs/{org}
func (s *OrganizationsService) EditOrgMembership(ctx context.Context, user, org string, membership *Membership) (*Membership, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/memberships/{username}
func (s *OrganizationsService) RemoveOrgMembership(ctx context.Context, user, org string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/invitations
func (s *OrganizationsService) ListPendingOrgInvitations(ctx context.Context, org string, opts *ListOptions) ([]*Invitation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type CreateOrgInvitationOptions struct {
	InviteeID *int64 `json:"invitee_id,omitempty"`

	Email *string `json:"email,omitempty"`

	Role   *string `json:"role,omitempty"`
	TeamID []int64 `json:"team_ids,omitempty"`
}

//meta:operation POST /orgs/{org}/invitations
func (s *OrganizationsService) CreateOrgInvitation(ctx context.Context, org string, body *CreateOrgInvitationOptions) (*Invitation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/invitations/{invitation_id}/teams
func (s *OrganizationsService) ListOrgInvitationTeams(ctx context.Context, org, invitationID string, opts *ListOptions) ([]*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/failed_invitations
func (s *OrganizationsService) ListFailedOrgInvitations(ctx context.Context, org string, opts *ListOptions) ([]*Invitation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
