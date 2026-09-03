package github

import (
	"context"
)

type RepositoryMergeRequest struct {
	Base          string  `json:"base"`
	Head          string  `json:"head"`
	CommitMessage *string `json:"commit_message,omitempty"`
}

type RepoMergeUpstreamRequest struct {
	Branch string `json:"branch"`
}

type RepoMergeUpstreamResult struct {
	Message    *string `json:"message,omitempty"`
	MergeType  *string `json:"merge_type,omitempty"`
	BaseBranch *string `json:"base_branch,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/merges
func (s *RepositoriesService) Merge(ctx context.Context, owner, repo string, body RepositoryMergeRequest) (*RepositoryCommit, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/merge-upstream
func (s *RepositoriesService) MergeUpstream(ctx context.Context, owner, repo string, body RepoMergeUpstreamRequest) (*RepoMergeUpstreamResult, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
