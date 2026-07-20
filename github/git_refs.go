package github

import (
	"context"
)

type Reference struct {
	Ref    *string    `json:"ref"`
	URL    *string    `json:"url"`
	Object *GitObject `json:"object"`
	NodeID *string    `json:"node_id,omitempty"`
}

func (r Reference) String() string { _ = "STUB: not implemented"; return "" }

type GitObject struct {
	Type *string `json:"type"`
	SHA  *string `json:"sha"`
	URL  *string `json:"url"`
}

func (o GitObject) String() string { _ = "STUB: not implemented"; return "" }

type CreateRef struct {
	Ref string `json:"ref"`
	SHA string `json:"sha"`
}

type UpdateRef struct {
	SHA   string `json:"sha"`
	Force *bool  `json:"force,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/git/ref/{ref}
func (s *GitService) GetRef(ctx context.Context, owner, repo, ref string) (*Reference, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func refURLEscape(ref string) string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/git/matching-refs/{ref}
func (s *GitService) ListMatchingRefs(ctx context.Context, owner, repo, ref string) ([]*Reference, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/git/refs
func (s *GitService) CreateRef(ctx context.Context, owner, repo string, body CreateRef) (*Reference, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/git/refs/{ref}
func (s *GitService) UpdateRef(ctx context.Context, owner, repo, ref string, body UpdateRef) (*Reference, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/git/refs/{ref}
func (s *GitService) DeleteRef(ctx context.Context, owner, repo, ref string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
