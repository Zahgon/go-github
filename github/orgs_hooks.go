package github

import (
	"context"
)

//meta:operation GET /orgs/{org}/hooks
func (s *OrganizationsService) ListHooks(ctx context.Context, org string, opts *ListOptions) ([]*Hook, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/hooks/{hook_id}
func (s *OrganizationsService) GetHook(ctx context.Context, org string, id int64) (*Hook, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/hooks
func (s *OrganizationsService) CreateHook(ctx context.Context, org string, hook *Hook) (*Hook, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/hooks/{hook_id}
func (s *OrganizationsService) EditHook(ctx context.Context, org string, id int64, body *Hook) (*Hook, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/hooks/{hook_id}/pings
func (s *OrganizationsService) PingHook(ctx context.Context, org string, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/hooks/{hook_id}
func (s *OrganizationsService) DeleteHook(ctx context.Context, org string, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
