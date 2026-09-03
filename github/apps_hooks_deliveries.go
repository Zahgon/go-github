package github

import (
	"context"
)

//meta:operation GET /app/hook/deliveries
func (s *AppsService) ListHookDeliveries(ctx context.Context, opts *ListCursorOptions) ([]*HookDelivery, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /app/hook/deliveries/{delivery_id}
func (s *AppsService) GetHookDelivery(ctx context.Context, deliveryID int64) (*HookDelivery, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /app/hook/deliveries/{delivery_id}/attempts
func (s *AppsService) RedeliverHookDelivery(ctx context.Context, deliveryID int64) (*HookDelivery, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
