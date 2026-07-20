package github

import (
	"context"
)

type ActionsCreateOrgVariableRequest struct {
	Name                  string  `json:"name"`
	Value                 string  `json:"value"`
	Visibility            string  `json:"visibility"`
	SelectedRepositoryIDs []int64 `json:"selected_repository_ids,omitzero"`
}

type ActionsUpdateOrgVariableRequest struct {
	Name                  *string `json:"name,omitempty"`
	Value                 *string `json:"value,omitempty"`
	Visibility            *string `json:"visibility,omitempty"`
	SelectedRepositoryIDs []int64 `json:"selected_repository_ids,omitzero"`
}

type ActionsCreateVariableRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ActionsUpdateVariableRequest struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type ActionsVariable struct {
	Name      string     `json:"name"`
	Value     string     `json:"value"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
	UpdatedAt *Timestamp `json:"updated_at,omitempty"`

	Visibility              *string `json:"visibility,omitempty"`
	SelectedRepositoriesURL *string `json:"selected_repositories_url,omitempty"`
}

type ActionsVariables struct {
	TotalCount int                `json:"total_count"`
	Variables  []*ActionsVariable `json:"variables"`
}

//meta:operation GET /repos/{owner}/{repo}/actions/variables
func (s *ActionsService) ListRepoVariables(ctx context.Context, owner, repo string, opts *ListOptions) (*ActionsVariables, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/organization-variables
func (s *ActionsService) ListRepoOrgVariables(ctx context.Context, owner, repo string, opts *ListOptions) (*ActionsVariables, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/variables
func (s *ActionsService) ListOrgVariables(ctx context.Context, org string, opts *ListOptions) (*ActionsVariables, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/environments/{environment_name}/variables
func (s *ActionsService) ListEnvVariables(ctx context.Context, owner, repo, env string, opts *ListOptions) (*ActionsVariables, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/variables/{name}
func (s *ActionsService) GetRepoVariable(ctx context.Context, owner, repo, name string) (*ActionsVariable, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/variables/{name}
func (s *ActionsService) GetOrgVariable(ctx context.Context, org, name string) (*ActionsVariable, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/environments/{environment_name}/variables/{name}
func (s *ActionsService) GetEnvVariable(ctx context.Context, owner, repo, env, variableName string) (*ActionsVariable, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/actions/variables
func (s *ActionsService) CreateRepoVariable(ctx context.Context, owner, repo string, body ActionsCreateVariableRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /orgs/{org}/actions/variables
func (s *ActionsService) CreateOrgVariable(ctx context.Context, org string, body ActionsCreateOrgVariableRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/environments/{environment_name}/variables
func (s *ActionsService) CreateEnvVariable(ctx context.Context, owner, repo, env string, body ActionsCreateVariableRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/actions/variables/{name}
func (s *ActionsService) UpdateRepoVariable(ctx context.Context, owner, repo, name string, body ActionsUpdateVariableRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PATCH /orgs/{org}/actions/variables/{name}
func (s *ActionsService) UpdateOrgVariable(ctx context.Context, org, name string, body ActionsUpdateOrgVariableRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/environments/{environment_name}/variables/{name}
func (s *ActionsService) UpdateEnvVariable(ctx context.Context, owner, repo, env, name string, body ActionsUpdateVariableRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/actions/variables/{name}
func (s *ActionsService) DeleteRepoVariable(ctx context.Context, owner, repo, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/variables/{name}
func (s *ActionsService) DeleteOrgVariable(ctx context.Context, org, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/environments/{environment_name}/variables/{name}
func (s *ActionsService) DeleteEnvVariable(ctx context.Context, owner, repo, env, variableName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/actions/variables/{name}/repositories
func (s *ActionsService) ListSelectedReposForOrgVariable(ctx context.Context, org, name string, opts *ListOptions) (*SelectedReposList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/variables/{name}/repositories
func (s *ActionsService) SetSelectedReposForOrgVariable(ctx context.Context, org, name string, ids []int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/actions/variables/{name}/repositories/{repository_id}
func (s *ActionsService) AddSelectedRepoToOrgVariable(ctx context.Context, org, name string, repo *Repository) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/variables/{name}/repositories/{repository_id}
func (s *ActionsService) RemoveSelectedRepoFromOrgVariable(ctx context.Context, org, name string, repo *Repository) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
