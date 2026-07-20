package github

import (
	"context"
)

//meta:operation GET /enterprises/{enterprise}/network-configurations
func (s *EnterpriseService) ListEnterpriseNetworkConfigurations(ctx context.Context, enterprise string, opts *ListOptions) (*NetworkConfigurations, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/network-configurations
func (s *EnterpriseService) CreateEnterpriseNetworkConfiguration(ctx context.Context, enterprise string, body NetworkConfigurationRequest) (*NetworkConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/network-configurations/{network_configuration_id}
func (s *EnterpriseService) GetEnterpriseNetworkConfiguration(ctx context.Context, enterprise, networkID string) (*NetworkConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/network-configurations/{network_configuration_id}
func (s *EnterpriseService) UpdateEnterpriseNetworkConfiguration(ctx context.Context, enterprise, networkID string, body NetworkConfigurationRequest) (*NetworkConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/network-configurations/{network_configuration_id}
func (s *EnterpriseService) DeleteEnterpriseNetworkConfiguration(ctx context.Context, enterprise, networkID string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /enterprises/{enterprise}/network-settings/{network_settings_id}
func (s *EnterpriseService) GetEnterpriseNetworkSettingsResource(ctx context.Context, enterprise, networkID string) (*NetworkSettingsResource, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
