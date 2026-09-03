package github

import (
	"context"
)

type ListEnterpriseCodeSecurityConfigurationOptions struct {
	Before string `url:"before,omitempty"`

	After string `url:"after,omitempty"`

	PerPage int `url:"per_page,omitempty"`
}

//meta:operation GET /enterprises/{enterprise}/code-security/configurations
func (s *EnterpriseService) ListCodeSecurityConfigurations(ctx context.Context, enterprise string, opts *ListEnterpriseCodeSecurityConfigurationOptions) ([]*CodeSecurityConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/code-security/configurations
func (s *EnterpriseService) CreateCodeSecurityConfiguration(ctx context.Context, enterprise string, body CodeSecurityConfiguration) (*CodeSecurityConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/code-security/configurations/defaults
func (s *EnterpriseService) ListDefaultCodeSecurityConfigurations(ctx context.Context, enterprise string) ([]*CodeSecurityConfigurationWithDefaultForNewRepos, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/code-security/configurations/{configuration_id}
func (s *EnterpriseService) GetCodeSecurityConfiguration(ctx context.Context, enterprise string, configurationID int64) (*CodeSecurityConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/code-security/configurations/{configuration_id}
func (s *EnterpriseService) UpdateCodeSecurityConfiguration(ctx context.Context, enterprise string, configurationID int64, body CodeSecurityConfiguration) (*CodeSecurityConfiguration, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/code-security/configurations/{configuration_id}
func (s *EnterpriseService) DeleteCodeSecurityConfiguration(ctx context.Context, enterprise string, configurationID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /enterprises/{enterprise}/code-security/configurations/{configuration_id}/attach
func (s *EnterpriseService) AttachCodeSecurityConfigurationToRepositories(ctx context.Context, enterprise string, configurationID int64, scope string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/code-security/configurations/{configuration_id}/defaults
func (s *EnterpriseService) SetDefaultCodeSecurityConfiguration(ctx context.Context, enterprise string, configurationID int64, defaultForNewRepos string) (*CodeSecurityConfigurationWithDefaultForNewRepos, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/code-security/configurations/{configuration_id}/repositories
func (s *EnterpriseService) ListCodeSecurityConfigurationRepositories(ctx context.Context, enterprise string, configurationID int64, opts *ListCodeSecurityConfigurationRepositoriesOptions) ([]*RepositoryAttachment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
