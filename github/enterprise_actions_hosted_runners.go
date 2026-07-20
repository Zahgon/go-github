package github

import (
	"context"
)

//meta:operation GET /enterprises/{enterprise}/actions/hosted-runners
func (s *EnterpriseService) ListHostedRunners(ctx context.Context, enterprise string, opts *ListOptions) (*HostedRunners, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/actions/hosted-runners
func (s *EnterpriseService) CreateHostedRunner(ctx context.Context, enterprise string, body CreateHostedRunnerRequest) (*HostedRunner, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/hosted-runners/images/github-owned
func (s *EnterpriseService) GetHostedRunnerGitHubOwnedImages(ctx context.Context, enterprise string) (*HostedRunnerImages, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/hosted-runners/images/partner
func (s *EnterpriseService) GetHostedRunnerPartnerImages(ctx context.Context, enterprise string) (*HostedRunnerImages, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/hosted-runners/limits
func (s *EnterpriseService) GetHostedRunnerLimits(ctx context.Context, enterprise string) (*HostedRunnerPublicIPLimits, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/hosted-runners/machine-sizes
func (s *EnterpriseService) GetHostedRunnerMachineSpecs(ctx context.Context, enterprise string) (*HostedRunnerMachineSpecs, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/hosted-runners/platforms
func (s *EnterpriseService) GetHostedRunnerPlatforms(ctx context.Context, enterprise string) (*HostedRunnerPlatforms, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/hosted-runners/{hosted_runner_id}
func (s *EnterpriseService) GetHostedRunner(ctx context.Context, enterprise string, runnerID int64) (*HostedRunner, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /enterprises/{enterprise}/actions/hosted-runners/{hosted_runner_id}
func (s *EnterpriseService) UpdateHostedRunner(ctx context.Context, enterprise string, runnerID int64, body UpdateHostedRunnerRequest) (*HostedRunner, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/actions/hosted-runners/{hosted_runner_id}
func (s *EnterpriseService) DeleteHostedRunner(ctx context.Context, enterprise string, runnerID int64) (*HostedRunner, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/hosted-runners/images/custom
func (s *EnterpriseService) ListHostedRunnerCustomImages(ctx context.Context, enterprise string) (*HostedRunnerCustomImages, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/hosted-runners/images/custom/{image_definition_id}
func (s *EnterpriseService) GetHostedRunnerCustomImage(ctx context.Context, enterprise string, imageDefinitionID int64) (*HostedRunnerCustomImage, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/actions/hosted-runners/images/custom/{image_definition_id}
func (s *EnterpriseService) DeleteHostedRunnerCustomImage(ctx context.Context, enterprise string, imageDefinitionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/hosted-runners/images/custom/{image_definition_id}/versions
func (s *EnterpriseService) ListHostedRunnerCustomImageVersions(ctx context.Context, enterprise string, imageDefinitionID int64) (*HostedRunnerCustomImageVersions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/hosted-runners/images/custom/{image_definition_id}/versions/{version}
func (s *EnterpriseService) GetHostedRunnerCustomImageVersion(ctx context.Context, enterprise string, imageDefinitionID int64, version string) (*HostedRunnerCustomImageVersion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/actions/hosted-runners/images/custom/{image_definition_id}/versions/{version}
func (s *EnterpriseService) DeleteHostedRunnerCustomImageVersion(ctx context.Context, enterprise string, imageDefinitionID int64, version string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
