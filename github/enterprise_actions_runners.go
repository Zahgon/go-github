package github

import (
	"context"
)

//meta:operation GET /enterprises/{enterprise}/actions/runners/downloads
func (s *EnterpriseService) ListRunnerApplicationDownloads(ctx context.Context, enterprise string) ([]*RunnerApplicationDownload, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/actions/runners/generate-jitconfig
func (s *EnterpriseService) CreateJITConfig(ctx context.Context, enterprise string, body CreateJITConfigRequest) (*JITRunnerConfig, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/actions/runners/registration-token
func (s *EnterpriseService) CreateRegistrationToken(ctx context.Context, enterprise string) (*RegistrationToken, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/runners
func (s *EnterpriseService) ListRunners(ctx context.Context, enterprise string, opts *ListRunnersOptions) (*Runners, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/runners/{runner_id}
func (s *EnterpriseService) GetRunner(ctx context.Context, enterprise string, runnerID int64) (*Runner, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/actions/runners/{runner_id}
func (s *EnterpriseService) RemoveRunner(ctx context.Context, enterprise string, runnerID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
