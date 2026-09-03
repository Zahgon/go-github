package github

import (
	"context"
	"io"
)

type SignatureVerification struct {
	Verified  *bool   `json:"verified,omitempty"`
	Reason    *string `json:"reason,omitempty"`
	Signature *string `json:"signature,omitempty"`
	Payload   *string `json:"payload,omitempty"`
}

type MessageSigner interface {
	Sign(w io.Writer, r io.Reader) error
}

type MessageSignerFunc func(w io.Writer, r io.Reader) error

func (f MessageSignerFunc) Sign(w io.Writer, r io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

type Commit struct {
	SHA          *string                `json:"sha,omitempty"`
	Author       *CommitAuthor          `json:"author,omitempty"`
	Committer    *CommitAuthor          `json:"committer,omitempty"`
	Message      *string                `json:"message,omitempty"`
	Tree         *Tree                  `json:"tree,omitempty"`
	Parents      []*Commit              `json:"parents,omitempty"`
	HTMLURL      *string                `json:"html_url,omitempty"`
	URL          *string                `json:"url,omitempty"`
	Verification *SignatureVerification `json:"verification,omitempty"`
	NodeID       *string                `json:"node_id,omitempty"`

	CommentCount *int `json:"comment_count,omitempty"`
}

func (c Commit) String() string { _ = "STUB: not implemented"; return "" }

type CommitAuthor struct {
	Date  *Timestamp `json:"date,omitempty"`
	Name  *string    `json:"name,omitempty"`
	Email *string    `json:"email,omitempty"`

	Login *string `json:"username,omitempty"`
}

func (c CommitAuthor) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/git/commits/{commit_sha}
func (s *GitService) GetCommit(ctx context.Context, owner, repo, sha string) (*Commit, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type createCommit struct {
	Author    *CommitAuthor `json:"author,omitempty"`
	Committer *CommitAuthor `json:"committer,omitempty"`
	Message   *string       `json:"message,omitempty"`
	Tree      *string       `json:"tree,omitempty"`
	Parents   []string      `json:"parents,omitempty"`
	Signature *string       `json:"signature,omitempty"`
}

type CreateCommitOptions struct {
	Signer MessageSigner
}

//meta:operation POST /repos/{owner}/{repo}/git/commits
func (s *GitService) CreateCommit(ctx context.Context, owner, repo string, commit Commit, opts *CreateCommitOptions) (*Commit, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func createSignature(signer MessageSigner, commit *createCommit) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func createSignatureMessage(commit *createCommit) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
