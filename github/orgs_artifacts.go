package github

import (
	"context"
)

type DeploymentRuntimeRisk string

const (
	DeploymentRuntimeRiskCriticalResource DeploymentRuntimeRisk = "critical-resource"
	DeploymentRuntimeRiskInternetExposed  DeploymentRuntimeRisk = "internet-exposed"
	DeploymentRuntimeRiskLateralMovement  DeploymentRuntimeRisk = "lateral-movement"
	DeploymentRuntimeRiskSensitiveData    DeploymentRuntimeRisk = "sensitive-data"
)

type ArtifactDeploymentRecord struct {
	ID                  *int64                  `json:"id,omitempty"`
	Digest              *string                 `json:"digest,omitempty"`
	LogicalEnvironment  *string                 `json:"logical_environment,omitempty"`
	PhysicalEnvironment *string                 `json:"physical_environment,omitempty"`
	Cluster             *string                 `json:"cluster,omitempty"`
	DeploymentName      *string                 `json:"deployment_name,omitempty"`
	Tags                map[string]string       `json:"tags,omitempty"`
	RuntimeRisks        []DeploymentRuntimeRisk `json:"runtime_risks,omitempty"`
	AttestationID       *int64                  `json:"attestation_id,omitempty"`
	CreatedAt           *Timestamp              `json:"created_at,omitempty"`
	UpdatedAt           *Timestamp              `json:"updated_at,omitempty"`
}

type CreateArtifactDeploymentRequest struct {
	Name                string                  `json:"name"`
	Digest              string                  `json:"digest"`
	Version             *string                 `json:"version,omitempty"`
	Status              string                  `json:"status"`
	LogicalEnvironment  string                  `json:"logical_environment"`
	PhysicalEnvironment *string                 `json:"physical_environment,omitempty"`
	Cluster             *string                 `json:"cluster,omitempty"`
	DeploymentName      string                  `json:"deployment_name"`
	Tags                map[string]string       `json:"tags,omitempty"`
	RuntimeRisks        []DeploymentRuntimeRisk `json:"runtime_risks,omitempty"`
	GithubRepository    *string                 `json:"github_repository,omitempty"`
}

type ArtifactDeploymentResponse struct {
	TotalCount        *int                        `json:"total_count,omitempty"`
	DeploymentRecords []*ArtifactDeploymentRecord `json:"deployment_records,omitempty"`
}

type ClusterArtifactDeployment struct {
	Name             string                  `json:"name"`
	Digest           string                  `json:"digest"`
	Version          *string                 `json:"version,omitempty"`
	Status           string                  `json:"status"`
	DeploymentName   string                  `json:"deployment_name"`
	Tags             map[string]string       `json:"tags,omitempty"`
	RuntimeRisks     []DeploymentRuntimeRisk `json:"runtime_risks,omitempty"`
	GithubRepository *string                 `json:"github_repository,omitempty"`
}

type ClusterDeploymentRecordsRequest struct {
	LogicalEnvironment  string                       `json:"logical_environment"`
	PhysicalEnvironment *string                      `json:"physical_environment,omitempty"`
	Deployments         []*ClusterArtifactDeployment `json:"deployments"`
}

type ArtifactStorageRecord struct {
	ID          *int64     `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Digest      *string    `json:"digest,omitempty"`
	ArtifactURL *string    `json:"artifact_url,omitempty"`
	RegistryURL *string    `json:"registry_url,omitempty"`
	Repository  *string    `json:"repository,omitempty"`
	Status      *string    `json:"status,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
}

type CreateArtifactStorageRequest struct {
	Name             string  `json:"name"`
	Digest           string  `json:"digest"`
	Version          *string `json:"version,omitempty"`
	ArtifactURL      *string `json:"artifact_url,omitempty"`
	Path             *string `json:"path,omitempty"`
	RegistryURL      string  `json:"registry_url"`
	Repository       *string `json:"repository,omitempty"`
	Status           *string `json:"status,omitempty"`
	GithubRepository *string `json:"github_repository,omitempty"`
}

type ArtifactStorageResponse struct {
	TotalCount     *int                     `json:"total_count,omitempty"`
	StorageRecords []*ArtifactStorageRecord `json:"storage_records,omitempty"`
}

//meta:operation POST /orgs/{org}/artifacts/metadata/deployment-record
func (s *OrganizationsService) CreateArtifactDeploymentRecord(ctx context.Context, org string, body CreateArtifactDeploymentRequest) (*ArtifactDeploymentResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/artifacts/metadata/deployment-record/cluster/{cluster}
func (s *OrganizationsService) SetClusterDeploymentRecords(ctx context.Context, org, cluster string, body ClusterDeploymentRecordsRequest) (*ArtifactDeploymentResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/artifacts/metadata/storage-record
func (s *OrganizationsService) CreateArtifactStorageRecord(ctx context.Context, org string, body CreateArtifactStorageRequest) (*ArtifactStorageResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/artifacts/{subject_digest}/metadata/deployment-records
func (s *OrganizationsService) ListArtifactDeploymentRecords(ctx context.Context, org, subjectDigest string) (*ArtifactDeploymentResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/artifacts/{subject_digest}/metadata/storage-records
func (s *OrganizationsService) ListArtifactStorageRecords(ctx context.Context, org, subjectDigest string) (*ArtifactStorageResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
