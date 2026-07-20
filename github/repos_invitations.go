package github

import (
	"context"
)

type RepositoryInvitation struct {
	ID      *int64      `json:"id,omitempty"`
	Repo    *Repository `json:"repository,omitempty"`
	Invitee *User       `json:"invitee,omitempty"`
	Inviter *User       `json:"inviter,omitempty"`

	Permissions *string    `json:"permissions,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	URL         *string    `json:"url,omitempty"`
	HTMLURL     *string    `json:"html_url,omitempty"`
	Expired     *bool      `json:"expired,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/invitations
func (s *RepositoriesService) ListInvitations(ctx context.Context, owner, repo string, opts *ListOptions) ([]*RepositoryInvitation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/invitations/{invitation_id}
func (s *RepositoriesService) DeleteInvitation(ctx context.Context, owner, repo string, invitationID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/invitations/{invitation_id}
func (s *RepositoriesService) UpdateInvitation(ctx context.Context, owner, repo string, invitationID int64, permissions string) (*RepositoryInvitation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
