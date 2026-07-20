package github

import (
	"context"
	"regexp"
)

type ComputeService string

const (
	ComputeServiceNone       ComputeService = "none"
	ComputeServiceActions    ComputeService = "actions"
	ComputeServiceCodespaces ComputeService = "codespaces"
)

type NetworkConfigurations struct {
	TotalCount            *int64                  `json:"total_count,omitempty"`
	NetworkConfigurations []*NetworkConfiguration `json:"network_configurations,omitempty"`
}

type NetworkConfiguration struct {
	ID                 *string         `json:"id,omitempty"`
	Name               *string         `json:"name,omitempty"`
	ComputeService     *ComputeService `json:"compute_service,omitempty"`
	NetworkSettingsIDs []string        `json:"network_settings_ids,omitempty"`
	CreatedOn          *Timestamp      `json:"created_on"`
}

type NetworkSettingsResource struct {
	ID                     *string `json:"id,omitempty"`
	NetworkConfigurationID *string `json:"network_configuration_id,omitempty"`
	Name                   *string `json:"name,omitempty"`
	SubnetID               *string `json:"subnet_id,omitempty"`
	Region                 *string `json:"region,omitempty"`
}

func validateComputeService(compute *ComputeService) error { _ = "STUB: not implemented"; return nil }

var validNetworkNameRE = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)

func validateNetworkName(name string) error { _ = "STUB: not implemented"; return nil }

func validateNetworkSettingsID(settingsID []string) error { _ = "STUB: not implemented"; return nil }

func validateNetworkConfigurationRequest(req NetworkConfigurationRequest) error {
	_ = "STUB: not implemented"
	return nil
}

type NetworkConfigurationRequest struct {
	Name               *string         `json:"name,omitempty"`
	ComputeService     *ComputeService `json:"compute_service,omitempty"`
	NetworkSettingsIDs []string        `json:"network_settings_ids,omitempty"`
}

//meta:operation GET /orgs/{org}/settings/network-configurations
func (s *OrganizationsService) ListNetworkConfigurations(ctx context.Context, org string, opts *ListOptions) (*NetworkConfigurations, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/settings/network-configurations
func (s *OrganizationsService) CreateNetworkConfiguration(ctx context.Context, org string, body NetworkConfigurationRequest) (*NetworkConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/settings/network-configurations/{network_configuration_id}
func (s *OrganizationsService) GetNetworkConfiguration(ctx context.Context, org, networkID string) (*NetworkConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/settings/network-configurations/{network_configuration_id}
func (s *OrganizationsService) UpdateNetworkConfiguration(ctx context.Context, org, networkID string, body NetworkConfigurationRequest) (*NetworkConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/settings/network-configurations/{network_configuration_id}
func (s *OrganizationsService) DeleteNetworkConfigurations(ctx context.Context, org, networkID string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/settings/network-settings/{network_settings_id}
func (s *OrganizationsService) GetNetworkConfigurationResource(ctx context.Context, org, networkID string) (*NetworkSettingsResource, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
