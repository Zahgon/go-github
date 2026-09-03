package github

import (
	"context"
)

//meta:operation PUT /users/{username}/site_admin
func (s *UsersService) PromoteSiteAdmin(ctx context.Context, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /users/{username}/site_admin
func (s *UsersService) DemoteSiteAdmin(ctx context.Context, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type UserSuspendOptions struct {
	Reason *string `json:"reason,omitempty"`
}

//meta:operation PUT /users/{username}/suspended
func (s *UsersService) Suspend(ctx context.Context, user string, body *UserSuspendOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /users/{username}/suspended
func (s *UsersService) Unsuspend(ctx context.Context, user string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
