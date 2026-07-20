package github

import (
	"context"
)

type ListOutsideCollaboratorsOptions struct {
	Filter string `url:"filter,omitempty"`

	ListOptions
}

//meta:operation GET /orgs/{org}/outside_collaborators
func (s *OrganizationsService) ListOutsideCollaborators(ctx context.Context, org string, opts *ListOutsideCollaboratorsOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/outside_collaborators/{username}
func (s *OrganizationsService) RemoveOutsideCollaborator(ctx context.Context, org, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/outside_collaborators/{username}
func (s *OrganizationsService) ConvertMemberToOutsideCollaborator(ctx context.Context, org, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
