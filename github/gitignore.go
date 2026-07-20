package github

import (
	"context"
)

type GitignoresService service

type Gitignore struct {
	Name   *string `json:"name,omitempty"`
	Source *string `json:"source,omitempty"`
}

func (g Gitignore) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /gitignore/templates
func (s *GitignoresService) List(ctx context.Context) ([]string, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /gitignore/templates/{name}
func (s *GitignoresService) Get(ctx context.Context, name string) (*Gitignore, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
