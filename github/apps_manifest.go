package github

import (
	"context"
)

type AppConfig struct {
	ID            *int64     `json:"id,omitempty"`
	Slug          *string    `json:"slug,omitempty"`
	NodeID        *string    `json:"node_id,omitempty"`
	Owner         *User      `json:"owner,omitempty"`
	Name          *string    `json:"name,omitempty"`
	Description   *string    `json:"description,omitempty"`
	ExternalURL   *string    `json:"external_url,omitempty"`
	HTMLURL       *string    `json:"html_url,omitempty"`
	CreatedAt     *Timestamp `json:"created_at,omitempty"`
	UpdatedAt     *Timestamp `json:"updated_at,omitempty"`
	ClientID      *string    `json:"client_id,omitempty"`
	ClientSecret  *string    `json:"client_secret,omitempty"`
	WebhookSecret *string    `json:"webhook_secret,omitempty"`
	PEM           *string    `json:"pem,omitempty"`
}

//meta:operation POST /app-manifests/{code}/conversions
func (s *AppsService) CompleteAppManifest(ctx context.Context, code string) (*AppConfig, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
