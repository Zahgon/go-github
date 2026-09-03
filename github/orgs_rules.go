package github

import (
	"context"
)

//meta:operation GET /orgs/{org}/rulesets
func (s *OrganizationsService) ListAllRepositoryRulesets(ctx context.Context, org string, opts *ListOptions) ([]*RepositoryRuleset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/rulesets
func (s *OrganizationsService) CreateRepositoryRuleset(ctx context.Context, org string, body RepositoryRuleset) (*RepositoryRuleset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/rulesets/{ruleset_id}
func (s *OrganizationsService) GetRepositoryRuleset(ctx context.Context, org string, rulesetID int64) (*RepositoryRuleset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/rulesets/{ruleset_id}
func (s *OrganizationsService) UpdateRepositoryRuleset(ctx context.Context, org string, rulesetID int64, body RepositoryRuleset) (*RepositoryRuleset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/rulesets/{ruleset_id}
func (s *OrganizationsService) DeleteRepositoryRuleset(ctx context.Context, org string, rulesetID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
