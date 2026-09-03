package github

import (
	"context"
)

type CreateUserRequest struct {
	Login     string  `json:"login"`
	Email     *string `json:"email,omitempty"`
	Suspended *bool   `json:"suspended,omitempty"`
}

//meta:operation POST /admin/users
func (s *AdminService) CreateUser(ctx context.Context, body CreateUserRequest) (*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /admin/users/{username}
func (s *AdminService) DeleteUser(ctx context.Context, username string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CreateUserImpersonationRequest struct {
	Scopes []string `json:"scopes"`
}

type OAuthAPP struct {
	URL      *string `json:"url,omitempty"`
	Name     *string `json:"name,omitempty"`
	ClientID *string `json:"client_id,omitempty"`
}

func (s OAuthAPP) String() string { _ = "STUB: not implemented"; return "" }

type UserAuthorization struct {
	ID             *int64     `json:"id,omitempty"`
	URL            *string    `json:"url,omitempty"`
	Scopes         []string   `json:"scopes,omitempty"`
	Token          *string    `json:"token,omitempty"`
	TokenLastEight *string    `json:"token_last_eight,omitempty"`
	HashedToken    *string    `json:"hashed_token,omitempty"`
	App            *OAuthAPP  `json:"app,omitempty"`
	Note           *string    `json:"note,omitempty"`
	NoteURL        *string    `json:"note_url,omitempty"`
	UpdatedAt      *Timestamp `json:"updated_at,omitempty"`
	CreatedAt      *Timestamp `json:"created_at,omitempty"`
	Fingerprint    *string    `json:"fingerprint,omitempty"`
}

//meta:operation POST /admin/users/{username}/authorizations
func (s *AdminService) CreateUserImpersonation(ctx context.Context, username string, body CreateUserImpersonationRequest) (*UserAuthorization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /admin/users/{username}/authorizations
func (s *AdminService) DeleteUserImpersonation(ctx context.Context, username string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
