package github

import (
	"context"
)

//meta:operation GET /users/{username}/attestations/{subject_digest}
func (s *UsersService) ListAttestations(ctx context.Context, user, subjectDigest string, opts *ListOptions) (*AttestationsResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
