package github

import (
	"context"
)

//meta:operation GET /repos/{owner}/{repo}/attestations/{subject_digest}
func (s *RepositoriesService) ListAttestations(ctx context.Context, owner, repo, subjectDigest string, opts *ListOptions) (*AttestationsResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
