package github

import (
	"context"
)

type EmojisService service

//meta:operation GET /emojis
func (s *EmojisService) List(ctx context.Context) (map[string]string, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Client) ListEmojis(ctx context.Context) (map[string]string, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
