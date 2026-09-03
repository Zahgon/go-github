package github

import (
	"context"
)

//meta:operation GET /orgs/{org}/hooks/{hook_id}/deliveries
func (s *OrganizationsService) ListHookDeliveries(ctx context.Context, org string, id int64, opts *ListCursorOptions) ([]*HookDelivery, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/hooks/{hook_id}/deliveries/{delivery_id}
func (s *OrganizationsService) GetHookDelivery(ctx context.Context, owner string, hookID, deliveryID int64) (*HookDelivery, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/hooks/{hook_id}/deliveries/{delivery_id}/attempts
func (s *OrganizationsService) RedeliverHookDelivery(ctx context.Context, owner string, hookID, deliveryID int64) (*HookDelivery, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
