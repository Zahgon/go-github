package github

import (
	"context"
	"encoding/json"
)

type SCIMService service

type SCIMUserAttributes struct {
	UserName    string           `json:"userName"`
	Name        SCIMUserName     `json:"name"`
	DisplayName *string          `json:"displayName,omitempty"`
	Emails      []*SCIMUserEmail `json:"emails"`
	Schemas     []string         `json:"schemas,omitempty"`
	ExternalID  *string          `json:"externalId,omitempty"`
	Groups      []string         `json:"groups,omitempty"`
	Roles       []*SCIMUserRole  `json:"roles,omitempty"`
	Active      *bool            `json:"active,omitempty"`

	ID   *string   `json:"id,omitempty"`
	Meta *SCIMMeta `json:"meta,omitempty"`
}

type SCIMUserName struct {
	GivenName  string  `json:"givenName"`
	FamilyName string  `json:"familyName"`
	Formatted  *string `json:"formatted,omitempty"`
}

type SCIMUserEmail struct {
	Value   string  `json:"value"`
	Primary *bool   `json:"primary,omitempty"`
	Type    *string `json:"type,omitempty"`
}

type SCIMUserRole struct {
	Value   string  `json:"value"`
	Display *string `json:"display,omitempty"`
	Type    *string `json:"type,omitempty"`
	Primary *bool   `json:"primary,omitempty"`
}

type SCIMMeta struct {
	ResourceType *string    `json:"resourceType,omitempty"`
	Created      *Timestamp `json:"created,omitempty"`
	LastModified *Timestamp `json:"lastModified,omitempty"`
	Location     *string    `json:"location,omitempty"`
}

type SCIMProvisionedIdentities struct {
	Schemas      []string              `json:"schemas,omitempty"`
	TotalResults *int                  `json:"totalResults,omitempty"`
	ItemsPerPage *int                  `json:"itemsPerPage,omitempty"`
	StartIndex   *int                  `json:"startIndex,omitempty"`
	Resources    []*SCIMUserAttributes `json:"Resources,omitempty"`
}

type ListSCIMProvisionedIdentitiesOptions struct {
	StartIndex *int `url:"startIndex,omitempty"`
	Count      *int `url:"count,omitempty"`

	Filter *string `url:"filter,omitempty"`
}

//meta:operation GET /scim/v2/organizations/{org}/Users
func (s *SCIMService) ListSCIMProvisionedIdentities(ctx context.Context, org string, opts *ListSCIMProvisionedIdentitiesOptions) (*SCIMProvisionedIdentities, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /scim/v2/organizations/{org}/Users
func (s *SCIMService) ProvisionAndInviteSCIMUser(ctx context.Context, org string, body *SCIMUserAttributes) (*SCIMUserAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /scim/v2/organizations/{org}/Users/{scim_user_id}
func (s *SCIMService) GetSCIMProvisioningInfoForUser(ctx context.Context, org, scimUserID string) (*SCIMUserAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type UpdateProvisionedOrgMembershipRequest struct {
	UserName    string           `json:"userName"`
	Name        SCIMUserName     `json:"name"`
	DisplayName *string          `json:"displayName,omitempty"`
	Emails      []*SCIMUserEmail `json:"emails"`
	Schemas     []string         `json:"schemas,omitempty"`
	ExternalID  *string          `json:"externalId,omitempty"`
	Groups      []string         `json:"groups,omitempty"`
	Active      *bool            `json:"active,omitempty"`
}

//meta:operation PUT /scim/v2/organizations/{org}/Users/{scim_user_id}
func (s *SCIMService) UpdateProvisionedOrgMembership(ctx context.Context, org, scimUserID string, body UpdateProvisionedOrgMembershipRequest) (*SCIMUserAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type UpdateAttributeForSCIMUserRequest struct {
	Schemas    []string                                `json:"schemas,omitempty"`
	Operations []*UpdateAttributeForSCIMUserOperations `json:"Operations"`
}

type UpdateAttributeForSCIMUserOperations struct {
	Op    string          `json:"op"`
	Path  *string         `json:"path,omitempty"`
	Value json.RawMessage `json:"value,omitempty"`
}

//meta:operation PATCH /scim/v2/organizations/{org}/Users/{scim_user_id}
func (s *SCIMService) UpdateAttributeForSCIMUser(ctx context.Context, org, scimUserID string, body UpdateAttributeForSCIMUserRequest) (*SCIMUserAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /scim/v2/organizations/{org}/Users/{scim_user_id}
func (s *SCIMService) DeleteSCIMUserFromOrg(ctx context.Context, org, scimUserID string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
