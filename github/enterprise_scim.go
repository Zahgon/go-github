package github

import (
	"context"
)

const SCIMSchemasURINamespacesGroups = "urn:ietf:params:scim:schemas:core:2.0:Group"

const SCIMSchemasURINamespacesUser = "urn:ietf:params:scim:schemas:core:2.0:User"

const SCIMSchemasURINamespacesListResponse = "urn:ietf:params:scim:api:messages:2.0:ListResponse"

const SCIMSchemasURINamespacesPatchOp = "urn:ietf:params:scim:api:messages:2.0:PatchOp"

type SCIMEnterpriseGroupAttributes struct {
	DisplayName *string                           `json:"displayName,omitempty"`
	Members     []*SCIMEnterpriseDisplayReference `json:"members,omitempty"`
	ExternalID  *string                           `json:"externalId,omitempty"`
	Schemas     []string                          `json:"schemas,omitempty"`

	ID   *string             `json:"id,omitempty"`
	Meta *SCIMEnterpriseMeta `json:"meta,omitempty"`
}

type SCIMEnterpriseDisplayReference struct {
	Value   string  `json:"value"`
	Ref     *string `json:"$ref,omitempty"`
	Display *string `json:"display,omitempty"`
}

type SCIMEnterpriseMeta struct {
	ResourceType string     `json:"resourceType"`
	Created      *Timestamp `json:"created,omitempty"`
	LastModified *Timestamp `json:"lastModified,omitempty"`
	Location     *string    `json:"location,omitempty"`
}

type SCIMEnterpriseGroups struct {
	Schemas      []string                         `json:"schemas,omitempty"`
	TotalResults *int                             `json:"totalResults,omitempty"`
	Resources    []*SCIMEnterpriseGroupAttributes `json:"Resources,omitempty"`
	StartIndex   *int                             `json:"startIndex,omitempty"`
	ItemsPerPage *int                             `json:"itemsPerPage,omitempty"`
}

type ListProvisionedSCIMGroupsEnterpriseOptions struct {
	Filter *string `url:"filter,omitempty"`

	ExcludedAttributes *string `url:"excludedAttributes,omitempty"`

	StartIndex *int `url:"startIndex,omitempty"`

	Count *int `url:"count,omitempty"`
}

type GetProvisionedSCIMGroupEnterpriseOptions struct {
	ExcludedAttributes *string `url:"excludedAttributes,omitempty"`
}

type SCIMEnterpriseUserAttributes struct {
	DisplayName string                     `json:"displayName"`
	Name        *SCIMEnterpriseUserName    `json:"name,omitempty"`
	UserName    string                     `json:"userName"`
	Emails      []*SCIMEnterpriseUserEmail `json:"emails"`
	Roles       []*SCIMEnterpriseUserRole  `json:"roles,omitempty"`
	ExternalID  string                     `json:"externalId"`
	Active      bool                       `json:"active"`
	Schemas     []string                   `json:"schemas"`

	ID     *string                           `json:"id,omitempty"`
	Groups []*SCIMEnterpriseDisplayReference `json:"groups,omitempty"`
	Meta   *SCIMEnterpriseMeta               `json:"meta,omitempty"`
}

type SCIMEnterpriseUserName struct {
	GivenName  string  `json:"givenName"`
	FamilyName string  `json:"familyName"`
	Formatted  *string `json:"formatted,omitempty"`
	MiddleName *string `json:"middleName,omitempty"`
}

type SCIMEnterpriseUserEmail struct {
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
}

type SCIMEnterpriseUserRole struct {
	Value   string  `json:"value"`
	Display *string `json:"display,omitempty"`
	Type    *string `json:"type,omitempty"`
	Primary *bool   `json:"primary,omitempty"`
}

type SCIMEnterpriseUsers struct {
	Schemas      []string                        `json:"schemas,omitempty"`
	TotalResults *int                            `json:"totalResults,omitempty"`
	ItemsPerPage *int                            `json:"itemsPerPage,omitempty"`
	StartIndex   *int                            `json:"startIndex,omitempty"`
	Resources    []*SCIMEnterpriseUserAttributes `json:"Resources,omitempty"`
}

type ListProvisionedSCIMUsersEnterpriseOptions struct {
	Filter *string `url:"filter,omitempty"`

	StartIndex *int `url:"startIndex,omitempty"`

	Count *int `url:"count,omitempty"`
}

type SCIMEnterpriseAttribute struct {
	Schemas    []string                            `json:"schemas"`
	Operations []*SCIMEnterpriseAttributeOperation `json:"Operations"`
}

type SCIMEnterpriseAttributeOperation struct {
	Op    string  `json:"op"`
	Path  *string `json:"path,omitempty"`
	Value any     `json:"value,omitempty"`
}

//meta:operation GET /scim/v2/enterprises/{enterprise}/Groups
func (s *EnterpriseService) ListProvisionedSCIMGroups(ctx context.Context, enterprise string, opts *ListProvisionedSCIMGroupsEnterpriseOptions) (*SCIMEnterpriseGroups, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /scim/v2/enterprises/{enterprise}/Users
func (s *EnterpriseService) ListProvisionedSCIMUsers(ctx context.Context, enterprise string, opts *ListProvisionedSCIMUsersEnterpriseOptions) (*SCIMEnterpriseUsers, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}
func (s *EnterpriseService) SetProvisionedSCIMGroup(ctx context.Context, enterprise, scimGroupID string, body SCIMEnterpriseGroupAttributes) (*SCIMEnterpriseGroupAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /scim/v2/enterprises/{enterprise}/Users/{scim_user_id}
func (s *EnterpriseService) SetProvisionedSCIMUser(ctx context.Context, enterprise, scimUserID string, body SCIMEnterpriseUserAttributes) (*SCIMEnterpriseUserAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}
func (s *EnterpriseService) UpdateSCIMGroupAttribute(ctx context.Context, enterprise, scimGroupID string, body SCIMEnterpriseAttribute) (*SCIMEnterpriseGroupAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /scim/v2/enterprises/{enterprise}/Users/{scim_user_id}
func (s *EnterpriseService) UpdateSCIMUserAttribute(ctx context.Context, enterprise, scimUserID string, body SCIMEnterpriseAttribute) (*SCIMEnterpriseUserAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /scim/v2/enterprises/{enterprise}/Groups
func (s *EnterpriseService) ProvisionSCIMGroup(ctx context.Context, enterprise string, body SCIMEnterpriseGroupAttributes) (*SCIMEnterpriseGroupAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /scim/v2/enterprises/{enterprise}/Users
func (s *EnterpriseService) ProvisionSCIMUser(ctx context.Context, enterprise string, body SCIMEnterpriseUserAttributes) (*SCIMEnterpriseUserAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}
func (s *EnterpriseService) GetProvisionedSCIMGroup(ctx context.Context, enterprise, scimGroupID string, opts *GetProvisionedSCIMGroupEnterpriseOptions) (*SCIMEnterpriseGroupAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /scim/v2/enterprises/{enterprise}/Users/{scim_user_id}
func (s *EnterpriseService) GetProvisionedSCIMUser(ctx context.Context, enterprise, scimUserID string) (*SCIMEnterpriseUserAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}
func (s *EnterpriseService) DeleteSCIMGroup(ctx context.Context, enterprise, scimGroupID string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /scim/v2/enterprises/{enterprise}/Users/{scim_user_id}
func (s *EnterpriseService) DeleteSCIMUser(ctx context.Context, enterprise, scimUserID string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
