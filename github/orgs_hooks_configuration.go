package github

import (
	"context"
)

//meta:operation GET /orgs/{org}/hooks/{hook_id}/config
func (s *OrganizationsService) GetHookConfiguration(ctx context.Context, org string, id int64) (*HookConfig, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/hooks/{hook_id}/config
func (s *OrganizationsService) UpdateHookConfiguration(ctx context.Context, org string, id int64, body HookConfig) (*HookConfig, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
