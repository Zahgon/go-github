package github

import (
	"context"
)

//meta:operation GET /enterprises/{enterprise}/properties/schema
func (s *EnterpriseService) GetAllCustomProperties(ctx context.Context, enterprise string) ([]*CustomProperty, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/properties/schema
func (s *EnterpriseService) CreateOrUpdateCustomProperties(ctx context.Context, enterprise string, properties []*CustomProperty) ([]*CustomProperty, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/properties/schema/{custom_property_name}
func (s *EnterpriseService) GetCustomProperty(ctx context.Context, enterprise, customPropertyName string) (*CustomProperty, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/properties/schema/{custom_property_name}
func (s *EnterpriseService) CreateOrUpdateCustomProperty(ctx context.Context, enterprise, customPropertyName string, body *CustomProperty) (*CustomProperty, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/properties/schema/{custom_property_name}
func (s *EnterpriseService) RemoveCustomProperty(ctx context.Context, enterprise, customPropertyName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
