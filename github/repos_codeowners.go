package github

import (
	"context"
)

type GetCodeownersErrorsOptions struct {
	Ref string `url:"ref,omitempty"`
}

type CodeownersErrors struct {
	Errors []*CodeownersError `json:"errors"`
}

type CodeownersError struct {
	Line       int     `json:"line"`
	Column     int     `json:"column"`
	Kind       string  `json:"kind"`
	Source     string  `json:"source"`
	Suggestion *string `json:"suggestion,omitempty"`
	Message    string  `json:"message"`
	Path       string  `json:"path"`
}

//meta:operation GET /repos/{owner}/{repo}/codeowners/errors
func (s *RepositoriesService) GetCodeownersErrors(ctx context.Context, owner, repo string, opts *GetCodeownersErrorsOptions) (*CodeownersErrors, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
