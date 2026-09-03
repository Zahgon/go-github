package github

import (
	"context"
)

type EnterpriseSecurityAnalysisSettings struct {
	AdvancedSecurityEnabledForNewRepositories             *bool   `json:"advanced_security_enabled_for_new_repositories,omitempty"`
	SecretScanningEnabledForNewRepositories               *bool   `json:"secret_scanning_enabled_for_new_repositories,omitempty"`
	SecretScanningPushProtectionEnabledForNewRepositories *bool   `json:"secret_scanning_push_protection_enabled_for_new_repositories,omitempty"`
	SecretScanningPushProtectionCustomLink                *string `json:"secret_scanning_push_protection_custom_link,omitempty"`
	SecretScanningValidityChecksEnabled                   *bool   `json:"secret_scanning_validity_checks_enabled,omitempty"`
}

//meta:operation GET /enterprises/{enterprise}/code_security_and_analysis
func (s *EnterpriseService) GetCodeSecurityAndAnalysis(ctx context.Context, enterprise string) (*EnterpriseSecurityAnalysisSettings, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/code_security_and_analysis
func (s *EnterpriseService) UpdateCodeSecurityAndAnalysis(ctx context.Context, enterprise string, body *EnterpriseSecurityAnalysisSettings) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /enterprises/{enterprise}/{security_product}/{enablement}
func (s *EnterpriseService) EnableDisableSecurityFeature(ctx context.Context, enterprise, securityProduct, enablement string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
