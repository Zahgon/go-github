package github

import (
	"context"
)

//meta:operation GET /orgs/{org}/actions/permissions
func (s *OrganizationsService) GetActionsPermissions(ctx context.Context, org string) (*ActionsPermissions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions
func (s *OrganizationsService) UpdateActionsPermissions(ctx context.Context, org string, actionsPermissions ActionsPermissions) (*ActionsPermissions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
