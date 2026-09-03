package github

import (
	"context"
)

type PublicKey struct {
	KeyID *string `json:"key_id"`
	Key   *string `json:"key"`
}

func (p *PublicKey) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s *ActionsService) getPublicKey(ctx context.Context, url string) (*PublicKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/secrets/public-key
func (s *ActionsService) GetRepoPublicKey(ctx context.Context, owner, repo string) (*PublicKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/secrets/public-key
func (s *ActionsService) GetOrgPublicKey(ctx context.Context, org string) (*PublicKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/environments/{environment_name}/secrets/public-key
func (s *ActionsService) GetEnvPublicKey(ctx context.Context, owner, repo, env string) (*PublicKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type Secret struct {
	Name                    string    `json:"name"`
	CreatedAt               Timestamp `json:"created_at"`
	UpdatedAt               Timestamp `json:"updated_at"`
	Visibility              string    `json:"visibility,omitempty"`
	SelectedRepositoriesURL string    `json:"selected_repositories_url,omitempty"`
}

type Secrets struct {
	TotalCount int       `json:"total_count"`
	Secrets    []*Secret `json:"secrets"`
}

//meta:operation GET /repos/{owner}/{repo}/actions/secrets
func (s *ActionsService) ListRepoSecrets(ctx context.Context, owner, repo string, opts *ListOptions) (*Secrets, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/organization-secrets
func (s *ActionsService) ListRepoOrgSecrets(ctx context.Context, owner, repo string, opts *ListOptions) (*Secrets, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/secrets
func (s *ActionsService) ListOrgSecrets(ctx context.Context, org string, opts *ListOptions) (*Secrets, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/environments/{environment_name}/secrets
func (s *ActionsService) ListEnvSecrets(ctx context.Context, owner, repo, env string, opts *ListOptions) (*Secrets, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/secrets/{secret_name}
func (s *ActionsService) GetRepoSecret(ctx context.Context, owner, repo, name string) (*Secret, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/secrets/{secret_name}
func (s *ActionsService) GetOrgSecret(ctx context.Context, org, name string) (*Secret, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/environments/{environment_name}/secrets/{secret_name}
func (s *ActionsService) GetEnvSecret(ctx context.Context, owner, repo, env, secretName string) (*Secret, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type SelectedRepoIDs []int64

type EncryptedSecret struct {
	Name                  string          `json:"-"`
	KeyID                 string          `json:"key_id"`
	EncryptedValue        string          `json:"encrypted_value"`
	Visibility            string          `json:"visibility,omitempty"`
	SelectedRepositoryIDs SelectedRepoIDs `json:"selected_repository_ids,omitempty"`
}

type SecretOrgRequest struct {
	KeyID                 string  `json:"key_id"`
	EncryptedValue        string  `json:"encrypted_value"`
	Visibility            string  `json:"visibility"`
	SelectedRepositoryIDs []int64 `json:"selected_repository_ids,omitzero"`
}

type SecretRequest struct {
	KeyID          string `json:"key_id"`
	EncryptedValue string `json:"encrypted_value"`
}

//meta:operation PUT /repos/{owner}/{repo}/actions/secrets/{secret_name}
func (s *ActionsService) CreateOrUpdateRepoSecret(ctx context.Context, owner, repo, name string, body SecretRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/actions/secrets/{secret_name}
func (s *ActionsService) CreateOrUpdateOrgSecret(ctx context.Context, org, name string, body SecretOrgRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/environments/{environment_name}/secrets/{secret_name}
func (s *ActionsService) CreateOrUpdateEnvSecret(ctx context.Context, owner, repo, env, name string, body SecretRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/actions/secrets/{secret_name}
func (s *ActionsService) DeleteRepoSecret(ctx context.Context, owner, repo, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/secrets/{secret_name}
func (s *ActionsService) DeleteOrgSecret(ctx context.Context, org, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/environments/{environment_name}/secrets/{secret_name}
func (s *ActionsService) DeleteEnvSecret(ctx context.Context, owner, repo, env, secretName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type SelectedReposList struct {
	TotalCount   *int          `json:"total_count,omitempty"`
	Repositories []*Repository `json:"repositories,omitempty"`
}

//meta:operation GET /orgs/{org}/actions/secrets/{secret_name}/repositories
func (s *ActionsService) ListSelectedReposForOrgSecret(ctx context.Context, org, name string, opts *ListOptions) (*SelectedReposList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/secrets/{secret_name}/repositories
func (s *ActionsService) SetSelectedReposForOrgSecret(ctx context.Context, org, name string, ids []int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/actions/secrets/{secret_name}/repositories/{repository_id}
func (s *ActionsService) AddSelectedRepoToOrgSecret(ctx context.Context, org, name string, repoID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/secrets/{secret_name}/repositories/{repository_id}
func (s *ActionsService) RemoveSelectedRepoFromOrgSecret(ctx context.Context, org, name string, repoID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
