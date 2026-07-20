package github

import (
	"context"
)

type GPGKey struct {
	ID                *int64      `json:"id,omitempty"`
	PrimaryKeyID      *int64      `json:"primary_key_id,omitempty"`
	KeyID             *string     `json:"key_id,omitempty"`
	RawKey            *string     `json:"raw_key,omitempty"`
	PublicKey         *string     `json:"public_key,omitempty"`
	Emails            []*GPGEmail `json:"emails,omitempty"`
	Subkeys           []*GPGKey   `json:"subkeys,omitempty"`
	CanSign           *bool       `json:"can_sign,omitempty"`
	CanEncryptComms   *bool       `json:"can_encrypt_comms,omitempty"`
	CanEncryptStorage *bool       `json:"can_encrypt_storage,omitempty"`
	CanCertify        *bool       `json:"can_certify,omitempty"`
	CreatedAt         *Timestamp  `json:"created_at,omitempty"`
	ExpiresAt         *Timestamp  `json:"expires_at,omitempty"`
}

func (k GPGKey) String() string { _ = "STUB: not implemented"; return "" }

type GPGEmail struct {
	Email    *string `json:"email,omitempty"`
	Verified *bool   `json:"verified,omitempty"`
}

//meta:operation GET /user/gpg_keys
//meta:operation GET /users/{username}/gpg_keys
func (s *UsersService) ListGPGKeys(ctx context.Context, user string, opts *ListOptions) ([]*GPGKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/gpg_keys/{gpg_key_id}
func (s *UsersService) GetGPGKey(ctx context.Context, id int64) (*GPGKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /user/gpg_keys
func (s *UsersService) CreateGPGKey(ctx context.Context, armoredPublicKey string) (*GPGKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /user/gpg_keys/{gpg_key_id}
func (s *UsersService) DeleteGPGKey(ctx context.Context, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
