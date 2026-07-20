package github

import (
	"context"
)

type HostedRunnerPublicIP struct {
	Enabled bool   `json:"enabled"`
	Prefix  string `json:"prefix"`
	Length  int    `json:"length"`
}

type HostedRunnerMachineSpec struct {
	ID        string `json:"id"`
	CPUCores  int    `json:"cpu_cores"`
	MemoryGB  int    `json:"memory_gb"`
	StorageGB int    `json:"storage_gb"`
}

type HostedRunner struct {
	ID                 *int64                   `json:"id,omitempty"`
	Name               *string                  `json:"name,omitempty"`
	RunnerGroupID      *int64                   `json:"runner_group_id,omitempty"`
	Platform           *string                  `json:"platform,omitempty"`
	ImageDetails       *HostedRunnerImageDetail `json:"image_details,omitempty"`
	MachineSizeDetails *HostedRunnerMachineSpec `json:"machine_size_details,omitempty"`
	Status             *string                  `json:"status,omitempty"`
	MaximumRunners     *int64                   `json:"maximum_runners,omitempty"`
	PublicIPEnabled    *bool                    `json:"public_ip_enabled,omitempty"`
	PublicIPs          []*HostedRunnerPublicIP  `json:"public_ips,omitempty"`
	LastActiveOn       *Timestamp               `json:"last_active_on,omitempty"`
}

type HostedRunnerImageDetail struct {
	ID          *string `json:"id"`
	SizeGB      *int64  `json:"size_gb"`
	DisplayName *string `json:"display_name"`
	Source      *string `json:"source"`
	Version     *string `json:"version"`
}

type HostedRunners struct {
	TotalCount int             `json:"total_count"`
	Runners    []*HostedRunner `json:"runners"`
}

//meta:operation GET /orgs/{org}/actions/hosted-runners
func (s *ActionsService) ListHostedRunners(ctx context.Context, org string, opts *ListOptions) (*HostedRunners, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type HostedRunnerImage struct {
	ID string `json:"id"`

	Source string `json:"source"`

	Version *string `json:"version,omitempty"`
}

type CreateHostedRunnerRequest struct {
	Name           string            `json:"name"`
	Image          HostedRunnerImage `json:"image"`
	Size           string            `json:"size"`
	RunnerGroupID  int64             `json:"runner_group_id"`
	MaximumRunners *int64            `json:"maximum_runners,omitempty"`
	EnableStaticIP *bool             `json:"enable_static_ip,omitempty"`
	ImageGen       *bool             `json:"image_gen,omitempty"`
}

type UpdateHostedRunnerRequest struct {
	Name           *string `json:"name,omitempty"`
	RunnerGroupID  *int64  `json:"runner_group_id,omitempty"`
	MaximumRunners *int64  `json:"maximum_runners,omitempty"`
	EnableStaticIP *bool   `json:"enable_static_ip,omitempty"`
	Size           *string `json:"size,omitempty"`
	ImageID        *string `json:"image_id,omitempty"`
	ImageVersion   *string `json:"image_version,omitempty"`
}

func validateCreateHostedRunnerRequest(request *CreateHostedRunnerRequest) error {
	_ = "STUB: not implemented"
	return nil
}

//meta:operation POST /orgs/{org}/actions/hosted-runners
func (s *ActionsService) CreateHostedRunner(ctx context.Context, org string, body CreateHostedRunnerRequest) (*HostedRunner, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type HostedRunnerCustomImage struct {
	ID                int64  `json:"id"`
	Platform          string `json:"platform"`
	Name              string `json:"name"`
	Source            string `json:"source"`
	VersionsCount     int    `json:"versions_count"`
	TotalVersionsSize int    `json:"total_versions_size"`
	LatestVersion     string `json:"latest_version"`
	State             string `json:"state"`
}

type HostedRunnerCustomImages struct {
	TotalCount int                        `json:"total_count"`
	Images     []*HostedRunnerCustomImage `json:"images"`
}

type HostedRunnerCustomImageVersion struct {
	Version      string    `json:"version"`
	SizeGB       int       `json:"size_gb"`
	State        string    `json:"state"`
	StateDetails string    `json:"state_details"`
	CreatedOn    Timestamp `json:"created_on"`
}

type HostedRunnerCustomImageVersions struct {
	TotalCount    int                               `json:"total_count"`
	ImageVersions []*HostedRunnerCustomImageVersion `json:"image_versions"`
}

type HostedRunnerImageSpecs struct {
	ID          string `json:"id"`
	Platform    string `json:"platform"`
	SizeGB      int    `json:"size_gb"`
	DisplayName string `json:"display_name"`
	Source      string `json:"source"`
}

type HostedRunnerImages struct {
	TotalCount int                       `json:"total_count"`
	Images     []*HostedRunnerImageSpecs `json:"images"`
}

//meta:operation GET /orgs/{org}/actions/hosted-runners/images/github-owned
func (s *ActionsService) GetHostedRunnerGitHubOwnedImages(ctx context.Context, org string) (*HostedRunnerImages, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/hosted-runners/images/partner
func (s *ActionsService) GetHostedRunnerPartnerImages(ctx context.Context, org string) (*HostedRunnerImages, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type HostedRunnerPublicIPLimits struct {
	PublicIPs *PublicIPUsage `json:"public_ips"`
}

type PublicIPUsage struct {
	Maximum      int64 `json:"maximum"`
	CurrentUsage int64 `json:"current_usage"`
}

//meta:operation GET /orgs/{org}/actions/hosted-runners/limits
func (s *ActionsService) GetHostedRunnerLimits(ctx context.Context, org string) (*HostedRunnerPublicIPLimits, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type HostedRunnerMachineSpecs struct {
	TotalCount   int                        `json:"total_count"`
	MachineSpecs []*HostedRunnerMachineSpec `json:"machine_specs"`
}

//meta:operation GET /orgs/{org}/actions/hosted-runners/machine-sizes
func (s *ActionsService) GetHostedRunnerMachineSpecs(ctx context.Context, org string) (*HostedRunnerMachineSpecs, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type HostedRunnerPlatforms struct {
	TotalCount int      `json:"total_count"`
	Platforms  []string `json:"platforms"`
}

//meta:operation GET /orgs/{org}/actions/hosted-runners/platforms
func (s *ActionsService) GetHostedRunnerPlatforms(ctx context.Context, org string) (*HostedRunnerPlatforms, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/hosted-runners/{hosted_runner_id}
func (s *ActionsService) GetHostedRunner(ctx context.Context, org string, runnerID int64) (*HostedRunner, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/actions/hosted-runners/{hosted_runner_id}
func (s *ActionsService) UpdateHostedRunner(ctx context.Context, org string, runnerID int64, body UpdateHostedRunnerRequest) (*HostedRunner, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/hosted-runners/{hosted_runner_id}
func (s *ActionsService) DeleteHostedRunner(ctx context.Context, org string, runnerID int64) (*HostedRunner, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/hosted-runners/images/custom
func (s *ActionsService) ListHostedRunnerCustomImages(ctx context.Context, org string) (*HostedRunnerCustomImages, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/hosted-runners/images/custom/{image_definition_id}
func (s *ActionsService) GetHostedRunnerCustomImage(ctx context.Context, org string, imageDefinitionID int64) (*HostedRunnerCustomImage, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/hosted-runners/images/custom/{image_definition_id}
func (s *ActionsService) DeleteHostedRunnerCustomImage(ctx context.Context, org string, imageDefinitionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/actions/hosted-runners/images/custom/{image_definition_id}/versions
func (s *ActionsService) ListHostedRunnerCustomImageVersions(ctx context.Context, org string, imageDefinitionID int64) (*HostedRunnerCustomImageVersions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/hosted-runners/images/custom/{image_definition_id}/versions/{version}
func (s *ActionsService) GetHostedRunnerCustomImageVersion(ctx context.Context, org string, imageDefinitionID int64, version string) (*HostedRunnerCustomImageVersion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/hosted-runners/images/custom/{image_definition_id}/versions/{version}
func (s *ActionsService) DeleteHostedRunnerCustomImageVersion(ctx context.Context, org string, imageDefinitionID int64, version string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
