package github

import (
	"context"
)

type CustomDeploymentProtectionRuleApp struct {
	ID             *int64  `json:"id,omitempty"`
	Slug           *string `json:"slug,omitempty"`
	IntegrationURL *string `json:"integration_url,omitempty"`
	NodeID         *string `json:"node_id,omitempty"`
}

type CustomDeploymentProtectionRule struct {
	ID      *int64                             `json:"id,omitempty"`
	NodeID  *string                            `json:"node_id,omitempty"`
	Enabled *bool                              `json:"enabled,omitempty"`
	App     *CustomDeploymentProtectionRuleApp `json:"app,omitempty"`
}

type ListDeploymentProtectionRuleResponse struct {
	TotalCount      *int                              `json:"total_count,omitempty"`
	ProtectionRules []*CustomDeploymentProtectionRule `json:"custom_deployment_protection_rules,omitempty"`
}

type ListCustomDeploymentRuleIntegrationsResponse struct {
	TotalCount            *int                                 `json:"total_count,omitempty"`
	AvailableIntegrations []*CustomDeploymentProtectionRuleApp `json:"available_custom_deployment_protection_rule_integrations,omitempty"`
}

type CustomDeploymentProtectionRuleRequest struct {
	IntegrationID *int64 `json:"integration_id,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules
func (s *RepositoriesService) GetAllDeploymentProtectionRules(ctx context.Context, owner, repo, environment string) (*ListDeploymentProtectionRuleResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules
func (s *RepositoriesService) CreateCustomDeploymentProtectionRule(ctx context.Context, owner, repo, environment string, body *CustomDeploymentProtectionRuleRequest) (*CustomDeploymentProtectionRule, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules/apps
func (s *RepositoriesService) ListCustomDeploymentRuleIntegrations(ctx context.Context, owner, repo, environment string, opts *ListOptions) (*ListCustomDeploymentRuleIntegrationsResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules/{protection_rule_id}
func (s *RepositoriesService) GetCustomDeploymentProtectionRule(ctx context.Context, owner, repo, environment string, protectionRuleID int64) (*CustomDeploymentProtectionRule, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/environments/{environment_name}/deployment_protection_rules/{protection_rule_id}
func (s *RepositoriesService) DisableCustomDeploymentProtectionRule(ctx context.Context, owner, repo, environment string, protectionRuleID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
