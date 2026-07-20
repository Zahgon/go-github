package github

import (
	"context"
)

type Tag struct {
	Tag          *string                `json:"tag,omitempty"`
	SHA          *string                `json:"sha,omitempty"`
	URL          *string                `json:"url,omitempty"`
	Message      *string                `json:"message,omitempty"`
	Tagger       *CommitAuthor          `json:"tagger,omitempty"`
	Object       *GitObject             `json:"object,omitempty"`
	Verification *SignatureVerification `json:"verification,omitempty"`
	NodeID       *string                `json:"node_id,omitempty"`
}

type CreateTag struct {
	Tag     string        `json:"tag,omitempty"`
	Message string        `json:"message,omitempty"`
	Object  string        `json:"object,omitempty"`
	Type    string        `json:"type,omitempty"`
	Tagger  *CommitAuthor `json:"tagger,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/git/tags/{tag_sha}
func (s *GitService) GetTag(ctx context.Context, owner, repo, sha string) (*Tag, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/git/tags
func (s *GitService) CreateTag(ctx context.Context, owner, repo string, body CreateTag) (*Tag, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
