package github

import (
	"context"
)

type OrganizationCustomPropertyValues struct {
	Properties []*CustomPropertyValue `json:"properties,omitempty"`
}

//meta:operation GET /organizations/{org}/org-properties/values
func (s *OrganizationsService) GetOrganizationCustomPropertyValues(ctx context.Context, org string) ([]*CustomPropertyValue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /organizations/{org}/org-properties/values
func (s *OrganizationsService) CreateOrUpdateOrganizationCustomPropertyValues(ctx context.Context, org string, body OrganizationCustomPropertyValues) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
