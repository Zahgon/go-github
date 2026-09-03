package github

import (
	"context"
)

//meta:operation POST /enterprises/{enterprise}/rulesets
func (s *EnterpriseService) CreateRepositoryRuleset(ctx context.Context, enterprise string, body RepositoryRuleset) (*RepositoryRuleset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/rulesets/{ruleset_id}
func (s *EnterpriseService) GetRepositoryRuleset(ctx context.Context, enterprise string, rulesetID int64) (*RepositoryRuleset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/rulesets/{ruleset_id}
func (s *EnterpriseService) UpdateRepositoryRuleset(ctx context.Context, enterprise string, rulesetID int64, body RepositoryRuleset) (*RepositoryRuleset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/rulesets/{ruleset_id}
func (s *EnterpriseService) DeleteRepositoryRuleset(ctx context.Context, enterprise string, rulesetID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
