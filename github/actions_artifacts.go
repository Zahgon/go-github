package github

import (
	"context"
	"net/url"
)

type ArtifactWorkflowRun struct {
	ID               *int64  `json:"id,omitempty"`
	RepositoryID     *int64  `json:"repository_id,omitempty"`
	HeadRepositoryID *int64  `json:"head_repository_id,omitempty"`
	HeadBranch       *string `json:"head_branch,omitempty"`
	HeadSHA          *string `json:"head_sha,omitempty"`
}

type Artifact struct {
	ID                 *int64     `json:"id,omitempty"`
	NodeID             *string    `json:"node_id,omitempty"`
	Name               *string    `json:"name,omitempty"`
	SizeInBytes        *int64     `json:"size_in_bytes,omitempty"`
	URL                *string    `json:"url,omitempty"`
	ArchiveDownloadURL *string    `json:"archive_download_url,omitempty"`
	Expired            *bool      `json:"expired,omitempty"`
	CreatedAt          *Timestamp `json:"created_at,omitempty"`
	UpdatedAt          *Timestamp `json:"updated_at,omitempty"`
	ExpiresAt          *Timestamp `json:"expires_at,omitempty"`

	Digest      *string              `json:"digest,omitempty"`
	WorkflowRun *ArtifactWorkflowRun `json:"workflow_run,omitempty"`
}

type ArtifactList struct {
	TotalCount *int64      `json:"total_count,omitempty"`
	Artifacts  []*Artifact `json:"artifacts,omitempty"`
}

type ListArtifactsOptions struct {
	Name *string `url:"name,omitempty"`

	ListOptions
}

type ArtifactPeriod struct {
	Days               *int `json:"days,omitempty"`
	MaximumAllowedDays *int `json:"maximum_allowed_days,omitempty"`
}

func (a ArtifactPeriod) String() string { _ = "STUB: not implemented"; return "" }

type ArtifactPeriodOpt struct {
	Days *int `json:"days,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/actions/artifacts
func (s *ActionsService) ListArtifacts(ctx context.Context, owner, repo string, opts *ListArtifactsOptions) (*ArtifactList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/runs/{run_id}/artifacts
func (s *ActionsService) ListWorkflowRunArtifacts(ctx context.Context, owner, repo string, runID int64, opts *ListOptions) (*ArtifactList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/artifacts/{artifact_id}
func (s *ActionsService) GetArtifact(ctx context.Context, owner, repo string, artifactID int64) (*Artifact, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/artifacts/{artifact_id}/{archive_format}
func (s *ActionsService) DownloadArtifact(ctx context.Context, owner, repo string, artifactID int64, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) downloadArtifactWithoutRateLimit(ctx context.Context, u string, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) downloadArtifactWithRateLimit(ctx context.Context, u string, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/actions/artifacts/{artifact_id}
func (s *ActionsService) DeleteArtifact(ctx context.Context, owner, repo string, artifactID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
