package github

import (
	"context"
)

type Key struct {
	ID        *int64     `json:"id,omitempty"`
	Key       *string    `json:"key,omitempty"`
	URL       *string    `json:"url,omitempty"`
	Title     *string    `json:"title,omitempty"`
	ReadOnly  *bool      `json:"read_only,omitempty"`
	Verified  *bool      `json:"verified,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
	AddedBy   *string    `json:"added_by,omitempty"`
	LastUsed  *Timestamp `json:"last_used,omitempty"`
}

func (k Key) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /user/keys
//meta:operation GET /users/{username}/keys
func (s *UsersService) ListKeys(ctx context.Context, user string, opts *ListOptions) ([]*Key, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/keys/{key_id}
func (s *UsersService) GetKey(ctx context.Context, id int64) (*Key, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /user/keys
func (s *UsersService) CreateKey(ctx context.Context, body *Key) (*Key, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /user/keys/{key_id}
func (s *UsersService) DeleteKey(ctx context.Context, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
