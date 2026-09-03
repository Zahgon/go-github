package github

import (
	"context"
)

type MarkdownService service

type MarkdownOptions struct {
	Mode string

	Context string
}

type markdownRenderRequest struct {
	Text    *string `json:"text,omitempty"`
	Mode    *string `json:"mode,omitempty"`
	Context *string `json:"context,omitempty"`
}

//meta:operation POST /markdown
func (s *MarkdownService) Render(ctx context.Context, text string, opts *MarkdownOptions) (string, *Response, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
