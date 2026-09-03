package github

type InteractionsService service

type InteractionRestriction struct {
	Limit *string `json:"limit,omitempty"`

	Origin *string `json:"origin,omitempty"`

	ExpiresAt *Timestamp `json:"expires_at,omitempty"`
}
