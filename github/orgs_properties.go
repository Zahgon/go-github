package github

import (
	"context"
)

type PropertyValueType string

const (
	PropertyValueTypeString       PropertyValueType = "string"
	PropertyValueTypeSingleSelect PropertyValueType = "single_select"
	PropertyValueTypeMultiSelect  PropertyValueType = "multi_select"
	PropertyValueTypeTrueFalse    PropertyValueType = "true_false"
	PropertyValueTypeURL          PropertyValueType = "url"
)

type CustomProperty struct {
	PropertyName *string `json:"property_name,omitempty"`

	URL *string `json:"url,omitempty"`

	SourceType *string `json:"source_type,omitempty"`

	ValueType PropertyValueType `json:"value_type"`

	Required *bool `json:"required,omitempty"`

	DefaultValue any `json:"default_value,omitempty"`

	Description *string `json:"description,omitempty"`

	AllowedValues []string `json:"allowed_values,omitempty"`

	ValuesEditableBy *string `json:"values_editable_by,omitempty"`
}

func (cp CustomProperty) DefaultValueString() (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (cp CustomProperty) DefaultValueStrings() ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (cp CustomProperty) DefaultValueBool() (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

type RepoCustomPropertyValue struct {
	RepositoryID       int64                  `json:"repository_id"`
	RepositoryName     string                 `json:"repository_name"`
	RepositoryFullName string                 `json:"repository_full_name"`
	Properties         []*CustomPropertyValue `json:"properties"`
}

type CustomPropertyValue struct {
	PropertyName string `json:"property_name"`
	Value        any    `json:"value"`
}

type ListCustomPropertyValuesOptions struct {
	RepositoryQuery string `url:"repository_query,omitempty"`
	ListOptions
}

func (cpv *CustomPropertyValue) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

//meta:operation GET /orgs/{org}/properties/schema
func (s *OrganizationsService) GetAllCustomProperties(ctx context.Context, org string) ([]*CustomProperty, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/properties/schema
func (s *OrganizationsService) CreateOrUpdateCustomProperties(ctx context.Context, org string, properties []*CustomProperty) ([]*CustomProperty, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/properties/schema/{custom_property_name}
func (s *OrganizationsService) GetCustomProperty(ctx context.Context, org, name string) (*CustomProperty, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/properties/schema/{custom_property_name}
func (s *OrganizationsService) CreateOrUpdateCustomProperty(ctx context.Context, org, customPropertyName string, body *CustomProperty) (*CustomProperty, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/properties/schema/{custom_property_name}
func (s *OrganizationsService) RemoveCustomProperty(ctx context.Context, org, customPropertyName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/properties/values
func (s *OrganizationsService) ListCustomPropertyValues(ctx context.Context, org string, opts *ListCustomPropertyValuesOptions) ([]*RepoCustomPropertyValue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/properties/values
func (s *OrganizationsService) CreateOrUpdateRepoCustomPropertyValues(ctx context.Context, org string, repoNames []string, properties []*CustomPropertyValue) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
