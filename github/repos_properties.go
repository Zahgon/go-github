package github

import (
	"context"
)

//meta:operation GET /repos/{owner}/{repo}/properties/values
func (s *RepositoriesService) GetAllCustomPropertyValues(ctx context.Context, org, repo string) ([]*CustomPropertyValue, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/properties/values
func (s *RepositoriesService) CreateOrUpdateCustomProperties(ctx context.Context, org, repo string, customPropertyValues []*CustomPropertyValue) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
