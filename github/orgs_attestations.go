package github

import (
	"context"
)

//meta:operation GET /orgs/{org}/attestations/{subject_digest}
func (s *OrganizationsService) ListAttestations(ctx context.Context, org, subjectDigest string, opts *ListOptions) (*AttestationsResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
