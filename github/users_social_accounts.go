package github

import (
	"context"
)

type SocialAccount struct {
	Provider *string `json:"provider,omitempty"`
	URL      *string `json:"url,omitempty"`
}

type socialAccountsRequest struct {
	AccountURLs []string `json:"account_urls"`
}

//meta:operation GET /user/social_accounts
func (s *UsersService) ListSocialAccounts(ctx context.Context, opts *ListOptions) ([]*SocialAccount, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /user/social_accounts
func (s *UsersService) AddSocialAccounts(ctx context.Context, accountURLs []string) ([]*SocialAccount, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /user/social_accounts
func (s *UsersService) DeleteSocialAccounts(ctx context.Context, accountURLs []string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /users/{username}/social_accounts
func (s *UsersService) ListUserSocialAccounts(ctx context.Context, username string, opts *ListOptions) ([]*SocialAccount, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
