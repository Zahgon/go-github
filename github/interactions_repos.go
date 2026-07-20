package github

import (
	"context"
)

//meta:operation GET /repos/{owner}/{repo}/interaction-limits
func (s *InteractionsService) GetRestrictionsForRepo(ctx context.Context, owner, repo string) (*InteractionRestriction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/interaction-limits
func (s *InteractionsService) UpdateRestrictionsForRepo(ctx context.Context, owner, repo, limit string) (*InteractionRestriction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/interaction-limits
func (s *InteractionsService) RemoveRestrictionsFromRepo(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
