package github

import (
	"encoding/json"
)

type Event struct {
	Type       *string          `json:"type,omitempty"`
	Public     *bool            `json:"public,omitempty"`
	RawPayload *json.RawMessage `json:"payload,omitempty"`
	Repo       *Repository      `json:"repo,omitempty"`
	Actor      *User            `json:"actor,omitempty"`
	Org        *Organization    `json:"org,omitempty"`
	CreatedAt  *Timestamp       `json:"created_at,omitempty"`
	ID         *string          `json:"id,omitempty"`
}

func (e Event) String() string { _ = "STUB: not implemented"; return "" }

func (e *Event) ParsePayload() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (e *Event) Payload() (payload any) { _ = "STUB: not implemented"; return *new(any) }
