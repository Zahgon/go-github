package github

import (
	"context"
	"encoding/json"
)

type Deployment struct {
	URL           *string         `json:"url,omitempty"`
	ID            *int64          `json:"id,omitempty"`
	SHA           *string         `json:"sha,omitempty"`
	Ref           *string         `json:"ref,omitempty"`
	Task          *string         `json:"task,omitempty"`
	Payload       json.RawMessage `json:"payload,omitempty"`
	Environment   *string         `json:"environment,omitempty"`
	Description   *string         `json:"description,omitempty"`
	Creator       *User           `json:"creator,omitempty"`
	CreatedAt     *Timestamp      `json:"created_at,omitempty"`
	UpdatedAt     *Timestamp      `json:"updated_at,omitempty"`
	StatusesURL   *string         `json:"statuses_url,omitempty"`
	RepositoryURL *string         `json:"repository_url,omitempty"`
	NodeID        *string         `json:"node_id,omitempty"`
}

type DeploymentRequest struct {
	Ref                   string   `json:"ref"`
	Task                  *string  `json:"task,omitempty"`
	AutoMerge             *bool    `json:"auto_merge,omitempty"`
	RequiredContexts      []string `json:"required_contexts,omitzero"`
	Payload               any      `json:"payload,omitempty"`
	Environment           *string  `json:"environment,omitempty"`
	Description           *string  `json:"description,omitempty"`
	TransientEnvironment  *bool    `json:"transient_environment,omitempty"`
	ProductionEnvironment *bool    `json:"production_environment,omitempty"`
}

type DeploymentsListOptions struct {
	SHA string `url:"sha,omitempty"`

	Ref string `url:"ref,omitempty"`

	Task string `url:"task,omitempty"`

	Environment string `url:"environment,omitempty"`

	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/deployments
func (s *RepositoriesService) ListDeployments(ctx context.Context, owner, repo string, opts *DeploymentsListOptions) ([]*Deployment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/deployments/{deployment_id}
func (s *RepositoriesService) GetDeployment(ctx context.Context, owner, repo string, deploymentID int64) (*Deployment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/deployments
func (s *RepositoriesService) CreateDeployment(ctx context.Context, owner, repo string, body DeploymentRequest) (*Deployment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/deployments/{deployment_id}
func (s *RepositoriesService) DeleteDeployment(ctx context.Context, owner, repo string, deploymentID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type DeploymentStatus struct {
	ID *int64 `json:"id,omitempty"`

	State          *string    `json:"state,omitempty"`
	Creator        *User      `json:"creator,omitempty"`
	Description    *string    `json:"description,omitempty"`
	Environment    *string    `json:"environment,omitempty"`
	NodeID         *string    `json:"node_id,omitempty"`
	CreatedAt      *Timestamp `json:"created_at,omitempty"`
	UpdatedAt      *Timestamp `json:"updated_at,omitempty"`
	TargetURL      *string    `json:"target_url,omitempty"`
	DeploymentURL  *string    `json:"deployment_url,omitempty"`
	RepositoryURL  *string    `json:"repository_url,omitempty"`
	EnvironmentURL *string    `json:"environment_url,omitempty"`
	LogURL         *string    `json:"log_url,omitempty"`
	URL            *string    `json:"url,omitempty"`
}

type DeploymentStatusRequest struct {
	State string `json:"state"`

	TargetURL      *string `json:"target_url,omitempty"`
	LogURL         *string `json:"log_url,omitempty"`
	Description    *string `json:"description,omitempty"`
	Environment    *string `json:"environment,omitempty"`
	EnvironmentURL *string `json:"environment_url,omitempty"`
	AutoInactive   *bool   `json:"auto_inactive,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/deployments/{deployment_id}/statuses
func (s *RepositoriesService) ListDeploymentStatuses(ctx context.Context, owner, repo string, deployment int64, opts *ListOptions) ([]*DeploymentStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/deployments/{deployment_id}/statuses/{status_id}
func (s *RepositoriesService) GetDeploymentStatus(ctx context.Context, owner, repo string, deploymentID, deploymentStatusID int64) (*DeploymentStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/deployments/{deployment_id}/statuses
func (s *RepositoriesService) CreateDeploymentStatus(ctx context.Context, owner, repo string, deploymentID int64, body DeploymentStatusRequest) (*DeploymentStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
