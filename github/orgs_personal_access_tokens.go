package github

import (
	"context"
)

type PersonalAccessToken struct {
	ID *int64 `json:"id"`

	Owner *User `json:"owner"`

	RepositorySelection *string `json:"repository_selection"`

	RepositoriesURL *string `json:"repositories_url"`

	Permissions *PersonalAccessTokenPermissions `json:"permissions"`

	AccessGrantedAt *Timestamp `json:"access_granted_at"`

	TokenExpired *bool `json:"token_expired"`

	TokenExpiresAt *Timestamp `json:"token_expires_at"`

	TokenID *int64 `json:"token_id"`

	TokenName *string `json:"token_name"`

	TokenLastUsedAt *Timestamp `json:"token_last_used_at"`
}

type ListFineGrainedPATOptions struct {
	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	Owner []string `url:"-"`

	Repository string `url:"repository,omitempty"`

	Permission string `url:"permission,omitempty"`

	LastUsedBefore string `url:"last_used_before,omitempty"`

	LastUsedAfter string `url:"last_used_after,omitempty"`

	TokenID []int64 `url:"-"`

	ListOptions
}

//meta:operation GET /orgs/{org}/personal-access-tokens
func (s *OrganizationsService) ListFineGrainedPersonalAccessTokens(ctx context.Context, org string, opts *ListFineGrainedPATOptions) ([]*PersonalAccessToken, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type FineGrainedPersonalAccessTokenRequest struct {
	ID int64 `json:"id"`

	Reason string `json:"reason"`

	Owner User `json:"owner"`

	RepositorySelection string `json:"repository_selection"`

	RepositoriesURL string `json:"repositories_url"`

	Permissions PersonalAccessTokenPermissions `json:"permissions"`

	CreatedAt *Timestamp `json:"created_at"`

	TokenExpired bool `json:"token_expired"`

	TokenExpiresAt *Timestamp `json:"token_expires_at"`

	TokenID int64 `json:"token_id"`

	TokenName string `json:"token_name"`

	TokenLastUsedAt *Timestamp `json:"token_last_used_at"`
}

//meta:operation GET /orgs/{org}/personal-access-token-requests
func (s *OrganizationsService) ListFineGrainedPersonalAccessTokenRequests(ctx context.Context, org string, opts *ListFineGrainedPATOptions) ([]*FineGrainedPersonalAccessTokenRequest, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ReviewPersonalAccessTokenRequestOptions struct {
	Action string  `json:"action"`
	Reason *string `json:"reason,omitempty"`
}

//meta:operation POST /orgs/{org}/personal-access-token-requests/{pat_request_id}
func (s *OrganizationsService) ReviewPersonalAccessTokenRequest(ctx context.Context, org string, requestID int64, opts ReviewPersonalAccessTokenRequestOptions) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addListFineGrainedPATOptions(s string, opts *ListFineGrainedPATOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
