package github

import (
	"context"
)

type InstallableOrganization struct {
	ID                        int64   `json:"id"`
	Login                     string  `json:"login"`
	AccessibleRepositoriesURL *string `json:"accessible_repositories_url,omitempty"`
}

type AccessibleRepository struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

type InstallAppRequest struct {
	ClientID string `json:"client_id"`

	RepositorySelection string `json:"repository_selection"`

	Repositories []string `json:"repositories,omitempty"`
}

//meta:operation GET /enterprises/{enterprise}/apps/installable_organizations
func (s *EnterpriseService) ListAppInstallableOrganizations(ctx context.Context, enterprise string, opts *ListOptions) ([]*InstallableOrganization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/apps/installable_organizations/{org}/accessible_repositories
func (s *EnterpriseService) ListAppAccessibleOrganizationRepositories(ctx context.Context, enterprise, org string, opts *ListOptions) ([]*AccessibleRepository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/apps/organizations/{org}/installations
func (s *EnterpriseService) ListAppInstallations(ctx context.Context, enterprise, org string, opts *ListOptions) ([]*Installation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/apps/organizations/{org}/installations
func (s *EnterpriseService) InstallApp(ctx context.Context, enterprise, org string, body InstallAppRequest) (*Installation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/apps/organizations/{org}/installations/{installation_id}
func (s *EnterpriseService) UninstallApp(ctx context.Context, enterprise, org string, installationID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type AppInstallationRepositoriesRequest struct {
	Repositories []string `json:"repositories"`
}

type UpdateAppInstallationRepositoriesRequest struct {
	RepositorySelection *string `json:"repository_selection,omitempty"`

	Repositories []string `json:"repositories,omitempty"`
}

//meta:operation GET /enterprises/{enterprise}/apps/organizations/{org}/installations/{installation_id}/repositories
func (s *EnterpriseService) ListRepositoriesForOrgAppInstallation(ctx context.Context, enterprise, org string, installationID int64, opts *ListOptions) ([]*AccessibleRepository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/apps/organizations/{org}/installations/{installation_id}/repositories
func (s *EnterpriseService) UpdateAppInstallationRepositories(ctx context.Context, enterprise, org string, installationID int64, body UpdateAppInstallationRepositoriesRequest) (*Installation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/apps/organizations/{org}/installations/{installation_id}/repositories/add
func (s *EnterpriseService) AddRepositoriesToAppInstallation(ctx context.Context, enterprise, org string, installationID int64, body AppInstallationRepositoriesRequest) ([]*AccessibleRepository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/apps/organizations/{org}/installations/{installation_id}/repositories/remove
func (s *EnterpriseService) RemoveRepositoriesFromAppInstallation(ctx context.Context, enterprise, org string, installationID int64, body AppInstallationRepositoriesRequest) ([]*AccessibleRepository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
