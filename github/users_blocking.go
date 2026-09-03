package github

import (
	"context"
)

//meta:operation GET /user/blocks
func (s *UsersService) ListBlockedUsers(ctx context.Context, opts *ListOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/blocks/{username}
func (s *UsersService) IsBlocked(ctx context.Context, user string) (bool, *Response, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

//meta:operation PUT /user/blocks/{username}
func (s *UsersService) BlockUser(ctx context.Context, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /user/blocks/{username}
func (s *UsersService) UnblockUser(ctx context.Context, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
