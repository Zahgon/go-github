package github

import (
	"context"
)

type ListCollaboratorsOptions struct {
	Affiliation string `url:"affiliation,omitempty"`

	Permission string `url:"permission,omitempty"`

	ListOptions
}

type CollaboratorInvitation struct {
	ID          *int64      `json:"id,omitempty"`
	Repo        *Repository `json:"repository,omitempty"`
	Invitee     *User       `json:"invitee,omitempty"`
	Inviter     *User       `json:"inviter,omitempty"`
	Permissions *string     `json:"permissions,omitempty"`
	CreatedAt   *Timestamp  `json:"created_at,omitempty"`
	URL         *string     `json:"url,omitempty"`
	HTMLURL     *string     `json:"html_url,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/collaborators
func (s *RepositoriesService) ListCollaborators(ctx context.Context, owner, repo string, opts *ListCollaboratorsOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/collaborators/{username}
func (s *RepositoriesService) IsCollaborator(ctx context.Context, owner, repo, user string) (bool, *Response, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

type RepositoryPermissionLevel struct {
	Permission *string `json:"permission,omitempty"`

	User *User `json:"user,omitempty"`

	RoleName *string `json:"role_name,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/collaborators/{username}/permission
func (s *RepositoriesService) GetPermissionLevel(ctx context.Context, owner, repo, user string) (*RepositoryPermissionLevel, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type RepositoryAddCollaboratorOptions struct {
	Permission string `json:"permission,omitempty"`
}

//meta:operation PUT /repos/{owner}/{repo}/collaborators/{username}
func (s *RepositoriesService) AddCollaborator(ctx context.Context, owner, repo, user string, body *RepositoryAddCollaboratorOptions) (*CollaboratorInvitation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/collaborators/{username}
func (s *RepositoriesService) RemoveCollaborator(ctx context.Context, owner, repo, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
