package github

import (
	"context"
)

type CodesOfConductService service

type CodeOfConduct struct {
	Name *string `json:"name,omitempty"`
	Key  *string `json:"key,omitempty"`
	URL  *string `json:"url,omitempty"`
	Body *string `json:"body,omitempty"`
}

func (c *CodeOfConduct) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /codes_of_conduct
func (s *CodesOfConductService) List(ctx context.Context) ([]*CodeOfConduct, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Client) ListCodesOfConduct(ctx context.Context) ([]*CodeOfConduct, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /codes_of_conduct/{key}
func (s *CodesOfConductService) Get(ctx context.Context, key string) (*CodeOfConduct, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Client) GetCodeOfConduct(ctx context.Context, key string) (*CodeOfConduct, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
