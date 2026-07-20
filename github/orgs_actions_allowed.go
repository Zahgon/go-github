package github

import (
	"context"
)

//meta:operation GET /orgs/{org}/actions/permissions/selected-actions
func (s *OrganizationsService) GetActionsAllowed(ctx context.Context, org string) (*ActionsAllowed, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/permissions/selected-actions
func (s *OrganizationsService) UpdateActionsAllowed(ctx context.Context, org string, actionsAllowed ActionsAllowed) (*ActionsAllowed, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
