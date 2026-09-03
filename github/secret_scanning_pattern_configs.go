package github

import (
	"context"
)

type SecretScanningPatternConfigs struct {
	PatternConfigVersion     *string                          `json:"pattern_config_version,omitempty"`
	ProviderPatternOverrides []*SecretScanningPatternOverride `json:"provider_pattern_overrides,omitempty"`
	CustomPatternOverrides   []*SecretScanningPatternOverride `json:"custom_pattern_overrides,omitempty"`
}

type SecretScanningPatternOverride struct {
	TokenType            *string `json:"token_type,omitempty"`
	CustomPatternVersion *string `json:"custom_pattern_version,omitempty"`
	Slug                 *string `json:"slug,omitempty"`
	DisplayName          *string `json:"display_name,omitempty"`
	AlertTotal           *int    `json:"alert_total,omitempty"`
	AlertTotalPercentage *int    `json:"alert_total_percentage,omitempty"`
	FalsePositives       *int    `json:"false_positives,omitempty"`
	FalsePositiveRate    *int    `json:"false_positive_rate,omitempty"`
	Bypassrate           *int    `json:"bypass_rate,omitempty"`
	DefaultSetting       *string `json:"default_setting,omitempty"`
	EnterpriseSetting    *string `json:"enterprise_setting,omitempty"`
	Setting              *string `json:"setting,omitempty"`
}

type SecretScanningPatternConfigsUpdate struct {
	PatternConfigVersion *string `json:"pattern_config_version,omitempty"`
}

type SecretScanningPatternConfigsUpdateOptions struct {
	PatternConfigVersion *string `json:"pattern_config_version,omitempty"`

	ProviderPatternSettings []*SecretScanningProviderPatternSetting `json:"provider_pattern_settings,omitempty"`

	CustomPatternSettings []*SecretScanningCustomPatternSetting `json:"custom_pattern_settings,omitempty"`
}

type SecretScanningProviderPatternSetting struct {
	TokenType string `json:"token_type"`

	PushProtectionSetting string `json:"push_protection_setting"`
}

type SecretScanningCustomPatternSetting struct {
	TokenType string `json:"token_type"`

	CustomPatternVersion *string `json:"custom_pattern_version,omitempty"`

	PushProtectionSetting string `json:"push_protection_setting"`
}

//meta:operation GET /enterprises/{enterprise}/secret-scanning/pattern-configurations
func (s *SecretScanningService) ListPatternConfigsForEnterprise(ctx context.Context, enterprise string) (*SecretScanningPatternConfigs, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/secret-scanning/pattern-configurations
func (s *SecretScanningService) ListPatternConfigsForOrg(ctx context.Context, org string) (*SecretScanningPatternConfigs, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/secret-scanning/pattern-configurations
func (s *SecretScanningService) UpdatePatternConfigsForEnterprise(ctx context.Context, enterprise string, body *SecretScanningPatternConfigsUpdateOptions) (*SecretScanningPatternConfigsUpdate, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/secret-scanning/pattern-configurations
func (s *SecretScanningService) UpdatePatternConfigsForOrg(ctx context.Context, org string, body *SecretScanningPatternConfigsUpdateOptions) (*SecretScanningPatternConfigsUpdate, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
