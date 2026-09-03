package github

import (
	"context"
	"encoding/json"
)

type HookDelivery struct {
	ID             *int64     `json:"id,omitempty"`
	GUID           *string    `json:"guid,omitempty"`
	DeliveredAt    *Timestamp `json:"delivered_at,omitempty"`
	Redelivery     *bool      `json:"redelivery,omitempty"`
	Duration       *float64   `json:"duration,omitempty"`
	Status         *string    `json:"status,omitempty"`
	StatusCode     *int       `json:"status_code,omitempty"`
	Event          *string    `json:"event,omitempty"`
	Action         *string    `json:"action,omitempty"`
	InstallationID *int64     `json:"installation_id,omitempty"`
	RepositoryID   *int64     `json:"repository_id,omitempty"`

	Request *HookRequest `json:"request,omitempty"`

	Response *HookResponse `json:"response,omitempty"`
}

func (d HookDelivery) String() string { _ = "STUB: not implemented"; return "" }

func getHeader(headers map[string]string, key string) string { _ = "STUB: not implemented"; return "" }

type HookRequest struct {
	Headers    map[string]string `json:"headers,omitempty"`
	RawPayload *json.RawMessage  `json:"payload,omitempty"`
}

func (r *HookRequest) GetHeader(key string) string { _ = "STUB: not implemented"; return "" }

func (r HookRequest) String() string { _ = "STUB: not implemented"; return "" }

type HookResponse struct {
	Headers    map[string]string `json:"headers,omitempty"`
	RawPayload *json.RawMessage  `json:"payload,omitempty"`
}

func (r *HookResponse) GetHeader(key string) string { _ = "STUB: not implemented"; return "" }

func (r HookResponse) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/hooks/{hook_id}/deliveries
func (s *RepositoriesService) ListHookDeliveries(ctx context.Context, owner, repo string, id int64, opts *ListCursorOptions) ([]*HookDelivery, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/hooks/{hook_id}/deliveries/{delivery_id}
func (s *RepositoriesService) GetHookDelivery(ctx context.Context, owner, repo string, hookID, deliveryID int64) (*HookDelivery, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/hooks/{hook_id}/deliveries/{delivery_id}/attempts
func (s *RepositoriesService) RedeliverHookDelivery(ctx context.Context, owner, repo string, hookID, deliveryID int64) (*HookDelivery, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (d *HookDelivery) ParseRequestPayload() (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
