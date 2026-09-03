package github

import (
	"context"
)

type Scope string

const (
	ScopeNone           Scope = "(no scope)"
	ScopeUser           Scope = "user"
	ScopeUserEmail      Scope = "user:email"
	ScopeUserFollow     Scope = "user:follow"
	ScopePublicRepo     Scope = "public_repo"
	ScopeRepo           Scope = "repo"
	ScopeRepoDeployment Scope = "repo_deployment"
	ScopeRepoStatus     Scope = "repo:status"
	ScopeDeleteRepo     Scope = "delete_repo"
	ScopeNotifications  Scope = "notifications"
	ScopeGist           Scope = "gist"
	ScopeReadRepoHook   Scope = "read:repo_hook"
	ScopeWriteRepoHook  Scope = "write:repo_hook"
	ScopeAdminRepoHook  Scope = "admin:repo_hook"
	ScopeAdminOrgHook   Scope = "admin:org_hook"
	ScopeReadOrg        Scope = "read:org"
	ScopeWriteOrg       Scope = "write:org"
	ScopeAdminOrg       Scope = "admin:org"
	ScopeReadPublicKey  Scope = "read:public_key"
	ScopeWritePublicKey Scope = "write:public_key"
	ScopeAdminPublicKey Scope = "admin:public_key"
	ScopeReadGPGKey     Scope = "read:gpg_key"
	ScopeWriteGPGKey    Scope = "write:gpg_key"
	ScopeAdminGPGKey    Scope = "admin:gpg_key"
	ScopeSecurityEvents Scope = "security_events"
)

type AuthorizationsService service

type Authorization struct {
	ID             *int64            `json:"id,omitempty"`
	URL            *string           `json:"url,omitempty"`
	Scopes         []Scope           `json:"scopes,omitempty"`
	Token          *string           `json:"token,omitempty"`
	TokenLastEight *string           `json:"token_last_eight,omitempty"`
	HashedToken    *string           `json:"hashed_token,omitempty"`
	App            *AuthorizationApp `json:"app,omitempty"`
	Note           *string           `json:"note,omitempty"`
	NoteURL        *string           `json:"note_url,omitempty"`
	UpdatedAt      *Timestamp        `json:"updated_at,omitempty"`
	CreatedAt      *Timestamp        `json:"created_at,omitempty"`
	Fingerprint    *string           `json:"fingerprint,omitempty"`

	User *User `json:"user,omitempty"`
}

func (a Authorization) String() string { _ = "STUB: not implemented"; return "" }

type AuthorizationApp struct {
	URL      *string `json:"url,omitempty"`
	Name     *string `json:"name,omitempty"`
	ClientID *string `json:"client_id,omitempty"`
}

func (a AuthorizationApp) String() string { _ = "STUB: not implemented"; return "" }

type Grant struct {
	ID        *int64            `json:"id,omitempty"`
	URL       *string           `json:"url,omitempty"`
	App       *AuthorizationApp `json:"app,omitempty"`
	CreatedAt *Timestamp        `json:"created_at,omitempty"`
	UpdatedAt *Timestamp        `json:"updated_at,omitempty"`
	Scopes    []string          `json:"scopes,omitempty"`
}

func (g Grant) String() string { _ = "STUB: not implemented"; return "" }

type AuthorizationRequest struct {
	Scopes       []Scope `json:"scopes,omitempty"`
	Note         *string `json:"note,omitempty"`
	NoteURL      *string `json:"note_url,omitempty"`
	ClientID     *string `json:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty"`
	Fingerprint  *string `json:"fingerprint,omitempty"`
}

func (a AuthorizationRequest) String() string { _ = "STUB: not implemented"; return "" }

type AuthorizationUpdateRequest struct {
	Scopes       []string `json:"scopes,omitempty"`
	AddScopes    []string `json:"add_scopes,omitempty"`
	RemoveScopes []string `json:"remove_scopes,omitempty"`
	Note         *string  `json:"note,omitempty"`
	NoteURL      *string  `json:"note_url,omitempty"`
	Fingerprint  *string  `json:"fingerprint,omitempty"`
}

func (a AuthorizationUpdateRequest) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation POST /applications/{client_id}/token
func (s *AuthorizationsService) Check(ctx context.Context, clientID, accessToken string) (*Authorization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /applications/{client_id}/token
func (s *AuthorizationsService) Reset(ctx context.Context, clientID, accessToken string) (*Authorization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /applications/{client_id}/token
func (s *AuthorizationsService) Revoke(ctx context.Context, clientID, accessToken string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /applications/{client_id}/grant
func (s *AuthorizationsService) DeleteGrant(ctx context.Context, clientID, accessToken string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /admin/users/{username}/authorizations
func (s *AuthorizationsService) CreateImpersonation(ctx context.Context, username string, body *AuthorizationRequest) (*Authorization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /admin/users/{username}/authorizations
func (s *AuthorizationsService) DeleteImpersonation(ctx context.Context, username string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
