package github

import (
	"context"
)

type EnterpriseCustomPropertySchema struct {
	Properties []*CustomProperty `json:"properties,omitempty"`
}

type EnterpriseCustomPropertiesValues struct {
	OrganizationID *int64 `json:"organization_id,omitempty"`

	OrganizationLogin *string `json:"organization_login,omitempty"`

	Properties []*CustomPropertyValue `json:"properties,omitempty"`
}

type EnterpriseCustomPropertyValuesRequest struct {
	OrganizationLogin []string `json:"organization_login"`

	Properties []*CustomPropertyValue `json:"properties"`
}

//meta:operation GET /enterprises/{enterprise}/org-properties/schema
func (s *EnterpriseService) GetOrganizationCustomPropertySchema(ctx context.Context, enterprise string) (*EnterpriseCustomPropertySchema, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/org-properties/schema
func (s *EnterpriseService) CreateOrUpdateOrganizationCustomPropertySchema(ctx context.Context, enterprise string, body EnterpriseCustomPropertySchema) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /enterprises/{enterprise}/org-properties/schema/{custom_property_name}
func (s *EnterpriseService) GetOrganizationCustomProperty(ctx context.Context, enterprise, customPropertyName string) (*CustomProperty, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/org-properties/schema/{custom_property_name}
func (s *EnterpriseService) CreateOrUpdateOrganizationCustomProperty(ctx context.Context, enterprise, customPropertyName string, body CustomProperty) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/org-properties/schema/{custom_property_name}
func (s *EnterpriseService) DeleteOrganizationCustomProperty(ctx context.Context, enterprise, customPropertyName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /enterprises/{enterprise}/org-properties/values
func (s *EnterpriseService) ListOrganizationCustomPropertyValues(ctx context.Context, enterprise string, opts *ListOptions) ([]*EnterpriseCustomPropertiesValues, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/org-properties/values
func (s *EnterpriseService) CreateOrUpdateOrganizationCustomPropertyValues(ctx context.Context, enterprise string, body EnterpriseCustomPropertyValuesRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
