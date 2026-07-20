package github

import (
	"context"
	"encoding/json"
)

type GetAuditLogOptions struct {
	Phrase  *string `url:"phrase,omitempty"`
	Include *string `url:"include,omitempty"`
	Order   *string `url:"order,omitempty"`

	ListCursorOptions
}

type ActorLocation struct {
	CountryCode *string `json:"country_code,omitempty"`
}

type AuditEntry struct {
	Action                   *string        `json:"action,omitempty"`
	Actor                    *string        `json:"actor,omitempty"`
	ActorID                  *int64         `json:"actor_id,omitempty"`
	ActorLocation            *ActorLocation `json:"actor_location,omitempty"`
	Business                 *string        `json:"business,omitempty"`
	BusinessID               *int64         `json:"business_id,omitempty"`
	CreatedAt                *Timestamp     `json:"created_at,omitempty"`
	DocumentID               *string        `json:"_document_id,omitempty"`
	ExternalIdentityNameID   *string        `json:"external_identity_nameid,omitempty"`
	ExternalIdentityUsername *string        `json:"external_identity_username,omitempty"`
	HashedToken              *string        `json:"hashed_token,omitempty"`
	Org                      *string        `json:"org,omitempty"`
	OrgID                    *int64         `json:"org_id,omitempty"`
	Timestamp                *Timestamp     `json:"@timestamp,omitempty"`
	TokenID                  *int64         `json:"token_id,omitempty"`
	TokenScopes              *string        `json:"token_scopes,omitempty"`
	User                     *string        `json:"user,omitempty"`
	UserID                   *int64         `json:"user_id,omitempty"`

	Data map[string]any `json:"data,omitempty"`

	AdditionalFields map[string]any `json:"-"`
}

func (a *AuditEntry) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func unmarshalStringOrArray(raw json.RawMessage) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalIntOrArray(raw json.RawMessage) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a AuditEntry) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//meta:operation GET /orgs/{org}/audit-log
func (s *OrganizationsService) GetAuditLog(ctx context.Context, org string, opts *GetAuditLogOptions) ([]*AuditEntry, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
