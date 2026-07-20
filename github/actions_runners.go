package github

import (
	"context"
)

type RunnerApplicationDownload struct {
	OS                *string `json:"os,omitempty"`
	Architecture      *string `json:"architecture,omitempty"`
	DownloadURL       *string `json:"download_url,omitempty"`
	Filename          *string `json:"filename,omitempty"`
	TempDownloadToken *string `json:"temp_download_token,omitempty"`
	SHA256Checksum    *string `json:"sha256_checksum,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/actions/runners/downloads
func (s *ActionsService) ListRunnerApplicationDownloads(ctx context.Context, owner, repo string) ([]*RunnerApplicationDownload, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type CreateJITConfigRequest struct {
	Name          string  `json:"name"`
	RunnerGroupID int64   `json:"runner_group_id"`
	WorkFolder    *string `json:"work_folder,omitempty"`

	Labels []string `json:"labels"`
}

type JITRunnerConfig struct {
	Runner           *Runner `json:"runner,omitempty"`
	EncodedJITConfig *string `json:"encoded_jit_config,omitempty"`
}

//meta:operation POST /orgs/{org}/actions/runners/generate-jitconfig
func (s *ActionsService) CreateOrgJITConfig(ctx context.Context, org string, body CreateJITConfigRequest) (*JITRunnerConfig, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/actions/runners/generate-jitconfig
func (s *ActionsService) CreateRepoJITConfig(ctx context.Context, owner, repo string, body CreateJITConfigRequest) (*JITRunnerConfig, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type RegistrationToken struct {
	Token     *string    `json:"token,omitempty"`
	ExpiresAt *Timestamp `json:"expires_at,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/actions/runners/registration-token
func (s *ActionsService) CreateRegistrationToken(ctx context.Context, owner, repo string) (*RegistrationToken, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type Runner struct {
	ID            *int64          `json:"id,omitempty"`
	RunnerGroupID *int64          `json:"runner_group_id,omitempty"`
	Name          *string         `json:"name,omitempty"`
	OS            *string         `json:"os,omitempty"`
	Status        *string         `json:"status,omitempty"`
	Busy          *bool           `json:"busy,omitempty"`
	Labels        []*RunnerLabels `json:"labels,omitempty"`
	Ephemeral     *bool           `json:"ephemeral,omitempty"`
	Version       *string         `json:"version,omitempty"`
}

type RunnerLabels struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

type Runners struct {
	TotalCount int       `json:"total_count"`
	Runners    []*Runner `json:"runners"`
}

type ListRunnersOptions struct {
	Name *string `url:"name,omitempty"`
	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/actions/runners
func (s *ActionsService) ListRunners(ctx context.Context, owner, repo string, opts *ListRunnersOptions) (*Runners, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/runners/{runner_id}
func (s *ActionsService) GetRunner(ctx context.Context, owner, repo string, runnerID int64) (*Runner, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type RemoveToken struct {
	Token     *string    `json:"token,omitempty"`
	ExpiresAt *Timestamp `json:"expires_at,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/actions/runners/remove-token
func (s *ActionsService) CreateRemoveToken(ctx context.Context, owner, repo string) (*RemoveToken, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/actions/runners/{runner_id}
func (s *ActionsService) RemoveRunner(ctx context.Context, owner, repo string, runnerID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/actions/runners/downloads
func (s *ActionsService) ListOrganizationRunnerApplicationDownloads(ctx context.Context, org string) ([]*RunnerApplicationDownload, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/actions/runners/registration-token
func (s *ActionsService) CreateOrganizationRegistrationToken(ctx context.Context, org string) (*RegistrationToken, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/runners
func (s *ActionsService) ListOrganizationRunners(ctx context.Context, org string, opts *ListRunnersOptions) (*Runners, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/runners/{runner_id}
func (s *ActionsService) GetOrganizationRunner(ctx context.Context, org string, runnerID int64) (*Runner, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/actions/runners/remove-token
func (s *ActionsService) CreateOrganizationRemoveToken(ctx context.Context, org string) (*RemoveToken, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/runners/{runner_id}
func (s *ActionsService) RemoveOrganizationRunner(ctx context.Context, org string, runnerID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
