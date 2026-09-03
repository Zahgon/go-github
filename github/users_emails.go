package github

import "context"

type UserEmail struct {
	Email      *string `json:"email,omitempty"`
	Primary    *bool   `json:"primary,omitempty"`
	Verified   *bool   `json:"verified,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
}

//meta:operation GET /user/emails
func (s *UsersService) ListEmails(ctx context.Context, opts *ListOptions) ([]*UserEmail, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /user/emails
func (s *UsersService) AddEmails(ctx context.Context, body []string) ([]*UserEmail, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /user/emails
func (s *UsersService) DeleteEmails(ctx context.Context, emails []string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PATCH /user/email/visibility
func (s *UsersService) SetEmailVisibility(ctx context.Context, visibility string) ([]*UserEmail, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
