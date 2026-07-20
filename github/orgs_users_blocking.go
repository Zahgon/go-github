package github

import (
	"context"
)

//meta:operation GET /orgs/{org}/blocks
func (s *OrganizationsService) ListBlockedUsers(ctx context.Context, org string, opts *ListOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/blocks/{username}
func (s *OrganizationsService) IsBlocked(ctx context.Context, org, user string) (bool, *Response, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

//meta:operation PUT /orgs/{org}/blocks/{username}
func (s *OrganizationsService) BlockUser(ctx context.Context, org, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/blocks/{username}
func (s *OrganizationsService) UnblockUser(ctx context.Context, org, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
