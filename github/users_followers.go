package github

import (
	"context"
)

//meta:operation GET /user/followers
//meta:operation GET /users/{username}/followers
func (s *UsersService) ListFollowers(ctx context.Context, user string, opts *ListOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/following
//meta:operation GET /users/{username}/following
func (s *UsersService) ListFollowing(ctx context.Context, user string, opts *ListOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/following/{username}
//meta:operation GET /users/{username}/following/{target_user}
func (s *UsersService) IsFollowing(ctx context.Context, user, target string) (bool, *Response, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

//meta:operation PUT /user/following/{username}
func (s *UsersService) Follow(ctx context.Context, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /user/following/{username}
func (s *UsersService) Unfollow(ctx context.Context, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
