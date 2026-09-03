package github

import (
	"context"
)

//meta:operation GET /repos/{owner}/{repo}/dependabot/secrets/public-key
func (s *DependabotService) GetRepoPublicKey(ctx context.Context, owner, repo string) (*PublicKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/dependabot/secrets/public-key
func (s *DependabotService) GetOrgPublicKey(ctx context.Context, org string) (*PublicKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/dependabot/secrets
func (s *DependabotService) ListRepoSecrets(ctx context.Context, owner, repo string, opts *ListOptions) (*Secrets, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/dependabot/secrets
func (s *DependabotService) ListOrgSecrets(ctx context.Context, org string, opts *ListOptions) (*Secrets, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/dependabot/secrets/{secret_name}
func (s *DependabotService) GetRepoSecret(ctx context.Context, owner, repo, name string) (*Secret, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/dependabot/secrets/{secret_name}
func (s *DependabotService) GetOrgSecret(ctx context.Context, org, name string) (*Secret, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/dependabot/secrets/{secret_name}
func (s *DependabotService) CreateOrUpdateRepoSecret(ctx context.Context, owner, repo, name string, body SecretRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/dependabot/secrets/{secret_name}
func (s *DependabotService) CreateOrUpdateOrgSecret(ctx context.Context, org, name string, body SecretOrgRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/dependabot/secrets/{secret_name}
func (s *DependabotService) DeleteRepoSecret(ctx context.Context, owner, repo, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/dependabot/secrets/{secret_name}
func (s *DependabotService) DeleteOrgSecret(ctx context.Context, org, name string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/dependabot/secrets/{secret_name}/repositories
func (s *DependabotService) ListSelectedReposForOrgSecret(ctx context.Context, org, name string, opts *ListOptions) (*SelectedReposList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/dependabot/secrets/{secret_name}/repositories
func (s *DependabotService) SetSelectedReposForOrgSecret(ctx context.Context, org, name string, ids []int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /orgs/{org}/dependabot/secrets/{secret_name}/repositories/{repository_id}
func (s *DependabotService) AddSelectedRepoToOrgSecret(ctx context.Context, org, name string, repoID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/dependabot/secrets/{secret_name}/repositories/{repository_id}
func (s *DependabotService) RemoveSelectedRepoFromOrgSecret(ctx context.Context, org, name string, repoID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
