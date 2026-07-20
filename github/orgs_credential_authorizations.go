package github

import (
	"context"
)

type CredentialAuthorization struct {
	Login *string `json:"login,omitempty"`

	CredentialID *int64 `json:"credential_id,omitempty"`

	CredentialType *string `json:"credential_type,omitempty"`

	TokenLastEight *string `json:"token_last_eight,omitempty"`

	CredentialAuthorizedAt *Timestamp `json:"credential_authorized_at,omitempty"`

	CredentialAccessedAt *Timestamp `json:"credential_accessed_at,omitempty"`

	Scopes []string `json:"scopes,omitempty"`

	Fingerprint *string `json:"fingerprint,omitempty"`

	AuthorizedCredentialID *int64 `json:"authorized_credential_id,omitempty"`

	AuthorizedCredentialTitle *string `json:"authorized_credential_title,omitempty"`

	AuthorizedCredentialNote *string `json:"authorized_credential_note,omitempty"`

	AuthorizedCredentialExpiresAt *Timestamp `json:"authorized_credential_expires_at,omitempty"`
}

type CredentialAuthorizationsListOptions struct {
	ListOptions

	Login string `url:"login,omitempty"`
}

//meta:operation GET /orgs/{org}/credential-authorizations
func (s *OrganizationsService) ListCredentialAuthorizations(ctx context.Context, org string, opts *CredentialAuthorizationsListOptions) ([]*CredentialAuthorization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/credential-authorizations/{credential_id}
func (s *OrganizationsService) RemoveCredentialAuthorization(ctx context.Context, org string, credentialID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
