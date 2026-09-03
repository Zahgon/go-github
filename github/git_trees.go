package github

import (
	"context"
)

type Tree struct {
	SHA     *string      `json:"sha,omitempty"`
	Entries []*TreeEntry `json:"tree,omitempty"`

	Truncated *bool `json:"truncated,omitempty"`
}

func (t Tree) String() string { _ = "STUB: not implemented"; return "" }

type TreeEntry struct {
	SHA     *string `json:"sha,omitempty"`
	Path    *string `json:"path,omitempty"`
	Mode    *string `json:"mode,omitempty"`
	Type    *string `json:"type,omitempty"`
	Size    *int    `json:"size,omitempty"`
	Content *string `json:"content,omitempty"`
	URL     *string `json:"url,omitempty"`
}

func (t TreeEntry) String() string { _ = "STUB: not implemented"; return "" }

type treeEntryWithFileDelete struct {
	SHA     *string `json:"sha"`
	Path    *string `json:"path,omitempty"`
	Mode    *string `json:"mode,omitempty"`
	Type    *string `json:"type,omitempty"`
	Size    *int    `json:"size,omitempty"`
	Content *string `json:"content,omitempty"`
	URL     *string `json:"url,omitempty"`
}

func (t TreeEntry) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//meta:operation GET /repos/{owner}/{repo}/git/trees/{tree_sha}
func (s *GitService) GetTree(ctx context.Context, owner, repo, sha string, recursive bool) (*Tree, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type createTree struct {
	BaseTree string `json:"base_tree,omitempty"`
	Entries  []any  `json:"tree"`
}

//meta:operation POST /repos/{owner}/{repo}/git/trees
func (s *GitService) CreateTree(ctx context.Context, owner, repo, baseTree string, entries []*TreeEntry) (*Tree, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
