package github

import (
	"context"
)

type CopilotCloudAgentConfiguration struct {
	MCPConfiguration                      any                            `json:"mcp_configuration"`
	EnabledTools                          *CopilotCloudAgentEnabledTools `json:"enabled_tools"`
	RequireActionsWorkflowApproval        bool                           `json:"require_actions_workflow_approval"`
	IsFirewallEnabled                     bool                           `json:"is_firewall_enabled"`
	IsFirewallRecommendedAllowlistEnabled bool                           `json:"is_firewall_recommended_allowlist_enabled"`
	CustomAllowlist                       []string                       `json:"custom_allowlist"`
}

type CopilotCloudAgentEnabledTools struct {
	Codeql                        bool `json:"codeql"`
	CopilotCodeReview             bool `json:"copilot_code_review"`
	SecretScanning                bool `json:"secret_scanning"`
	DependencyVulnerabilityChecks bool `json:"dependency_vulnerability_checks"`
}

//meta:operation GET /repos/{owner}/{repo}/copilot/cloud-agent/configuration
func (s *CopilotService) GetCloudAgentConfiguration(ctx context.Context, owner, repo string) (*CopilotCloudAgentConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
