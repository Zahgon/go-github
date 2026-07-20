package github

import (
	"context"
)

//meta:operation GET /user/codespaces/secrets
func (s *CodespacesService) ListUserSecrets(ctx context.Context, opts *ListOptions) (*Secrets, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/codespaces/secrets
func (s *CodespacesService) ListOrgSecrets(ctx context.Context, org string, opts *ListOptions) (*Secrets, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/codespaces/secrets
func (s *CodespacesService) ListRepoSecrets(ctx context.Context, owner, repo string, opts *ListOptions) (*Secrets, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *CodespacesService) listSecrets(ctx context.Context, url string) (*Secrets, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/codespaces/secrets/public-key
func (s *CodespacesService) GetUserPublicKey(ctx context.Context) (*PublicKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/codespaces/secrets/public-key
func (s *CodespacesService) GetOrgPublicKey(ctx context.Context, org string) (*PublicKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/codespaces/secrets/public-key
func (s *CodespacesService) GetRepoPublicKey(ctx context.Context, owner, repo string) (*PublicKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *CodespacesService) getPublicKey(ctx context.Context, url string) (*PublicKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/codespaces/secrets/{secret_name}
func (s *CodespacesService) GetUserSecret(ctx context.Context, name string) (*Secret, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/codespaces/secrets/{secret_name}
func (s *CodespacesService) GetOrgSecret(ctx context.Context, org, name string) (*Secret, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/codespaces/secrets/{secret_name}
func (s *CodespacesService) GetRepoSecret(ctx context.Context, owner, repo, name string) (*Secret, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *CodespacesService) getSecret(ctx context.Context, url string) (*Secret, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /user/codespaces/secrets/{secret_name}
func (s *CodespacesService) CreateOrUpdateUserSecret(ctx context.Context, body *EncryptedSecret) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/codespaces/secrets/{secret_name}
func (s *CodespacesService) CreateOrUpdateOrgSecret(ctx context.Context, org string, body *EncryptedSecret) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/codespaces/secrets/{secret_name}
func (s *CodespacesService) CreateOrUpdateRepoSecret(ctx context.Context, owner, repo string, body *EncryptedSecret) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CodespacesService) createOrUpdateSecret(ctx context.Context, url string, body *EncryptedSecret) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /user/codespaces/secrets/{secret_name}
func (s *CodespacesService) DeleteUserSecret(ctx context.Context, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/codespaces/secrets/{secret_name}
func (s *CodespacesService) DeleteOrgSecret(ctx context.Context, org, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/codespaces/secrets/{secret_name}
func (s *CodespacesService) DeleteRepoSecret(ctx context.Context, owner, repo, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CodespacesService) deleteSecret(ctx context.Context, url string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /user/codespaces/secrets/{secret_name}/repositories
func (s *CodespacesService) ListSelectedReposForUserSecret(ctx context.Context, name string, opts *ListOptions) (*SelectedReposList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/codespaces/secrets/{secret_name}/repositories
func (s *CodespacesService) ListSelectedReposForOrgSecret(ctx context.Context, org, name string, opts *ListOptions) (*SelectedReposList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *CodespacesService) listSelectedReposForSecret(ctx context.Context, url string) (*SelectedReposList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /user/codespaces/secrets/{secret_name}/repositories
func (s *CodespacesService) SetSelectedReposForUserSecret(ctx context.Context, name string, ids SelectedRepoIDs) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/codespaces/secrets/{secret_name}/repositories
func (s *CodespacesService) SetSelectedReposForOrgSecret(ctx context.Context, org, name string, ids SelectedRepoIDs) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CodespacesService) setSelectedRepoForSecret(ctx context.Context, url string, ids SelectedRepoIDs) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /user/codespaces/secrets/{secret_name}/repositories/{repository_id}
func (s *CodespacesService) AddSelectedRepoToUserSecret(ctx context.Context, name string, repo *Repository) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/codespaces/secrets/{secret_name}/repositories/{repository_id}
func (s *CodespacesService) AddSelectedRepoToOrgSecret(ctx context.Context, org, name string, repo *Repository) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CodespacesService) addSelectedRepoToSecret(ctx context.Context, url string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /user/codespaces/secrets/{secret_name}/repositories/{repository_id}
func (s *CodespacesService) RemoveSelectedRepoFromUserSecret(ctx context.Context, name string, repo *Repository) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/codespaces/secrets/{secret_name}/repositories/{repository_id}
func (s *CodespacesService) RemoveSelectedRepoFromOrgSecret(ctx context.Context, org, name string, repo *Repository) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CodespacesService) removeSelectedRepoFromSecret(ctx context.Context, url string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
