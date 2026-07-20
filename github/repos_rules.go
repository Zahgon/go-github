package github

import (
	"context"
	"iter"
)

//meta:operation GET /repos/{owner}/{repo}/rules/branches/{branch}
func (s *RepositoriesService) ListRulesForBranch(ctx context.Context, owner, repo, branch string, opts *ListOptions) (*BranchRules, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *RepositoriesService) ListRulesForBranchIter(ctx context.Context, owner, repo, branch string, opts *ListOptions) iter.Seq2[any, error] {
	_ = "STUB: not implemented"
	return nil
}

type RepositoryListRulesetsOptions struct {
	IncludesParents *bool `url:"includes_parents,omitempty"`
	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/rulesets
func (s *RepositoriesService) GetAllRulesets(ctx context.Context, owner, repo string, opts *RepositoryListRulesetsOptions) ([]*RepositoryRuleset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/rulesets
func (s *RepositoriesService) CreateRuleset(ctx context.Context, owner, repo string, body RepositoryRuleset) (*RepositoryRuleset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/rulesets/{ruleset_id}
func (s *RepositoriesService) GetRuleset(ctx context.Context, owner, repo string, rulesetID int64, includesParents bool) (*RepositoryRuleset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/rulesets/{ruleset_id}
func (s *RepositoriesService) UpdateRuleset(ctx context.Context, owner, repo string, rulesetID int64, body RepositoryRuleset) (*RepositoryRuleset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/rulesets/{ruleset_id}
func (s *RepositoriesService) DeleteRuleset(ctx context.Context, owner, repo string, rulesetID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
