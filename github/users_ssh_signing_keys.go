package github

import (
	"context"
)

type SSHSigningKey struct {
	ID        *int64     `json:"id,omitempty"`
	Key       *string    `json:"key,omitempty"`
	Title     *string    `json:"title,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
}

func (k SSHSigningKey) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /user/ssh_signing_keys
//meta:operation GET /users/{username}/ssh_signing_keys
func (s *UsersService) ListSSHSigningKeys(ctx context.Context, user string, opts *ListOptions) ([]*SSHSigningKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/ssh_signing_keys/{ssh_signing_key_id}
func (s *UsersService) GetSSHSigningKey(ctx context.Context, id int64) (*SSHSigningKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /user/ssh_signing_keys
func (s *UsersService) CreateSSHSigningKey(ctx context.Context, body *Key) (*SSHSigningKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /user/ssh_signing_keys/{ssh_signing_key_id}
func (s *UsersService) DeleteSSHSigningKey(ctx context.Context, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
