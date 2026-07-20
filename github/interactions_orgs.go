package github

import (
	"context"
)

//meta:operation GET /orgs/{org}/interaction-limits
func (s *InteractionsService) GetRestrictionsForOrg(ctx context.Context, organization string) (*InteractionRestriction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/interaction-limits
func (s *InteractionsService) UpdateRestrictionsForOrg(ctx context.Context, organization, limit string) (*InteractionRestriction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/interaction-limits
func (s *InteractionsService) RemoveRestrictionsFromOrg(ctx context.Context, organization string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
