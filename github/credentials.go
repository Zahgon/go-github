package github

import (
	"context"
)

type CredentialsService service

type revokeCredentialsRequest struct {
	Credentials []string `json:"credentials"`
}

//meta:operation POST /credentials/revoke
func (s *CredentialsService) Revoke(ctx context.Context, credentials []string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
