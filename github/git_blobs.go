package github

import (
	"context"
)

type Blob struct {
	Content  *string `json:"content,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	SHA      *string `json:"sha,omitempty"`
	Size     *int    `json:"size,omitempty"`
	URL      *string `json:"url,omitempty"`
	NodeID   *string `json:"node_id,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/git/blobs/{file_sha}
func (s *GitService) GetBlob(ctx context.Context, owner, repo, sha string) (*Blob, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/git/blobs/{file_sha}
func (s *GitService) GetBlobRaw(ctx context.Context, owner, repo, sha string) ([]byte, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/git/blobs
func (s *GitService) CreateBlob(ctx context.Context, owner, repo string, body Blob) (*Blob, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
