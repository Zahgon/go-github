package github

import (
	"context"
)

type CodespacesService service

type Codespace struct {
	ID                             *int64                        `json:"id,omitempty"`
	Name                           *string                       `json:"name,omitempty"`
	DisplayName                    *string                       `json:"display_name,omitempty"`
	EnvironmentID                  *string                       `json:"environment_id,omitempty"`
	Owner                          *User                         `json:"owner,omitempty"`
	BillableOwner                  *User                         `json:"billable_owner,omitempty"`
	Repository                     *Repository                   `json:"repository,omitempty"`
	Machine                        *CodespacesMachine            `json:"machine,omitempty"`
	DevcontainerPath               *string                       `json:"devcontainer_path,omitempty"`
	Prebuild                       *bool                         `json:"prebuild,omitempty"`
	CreatedAt                      *Timestamp                    `json:"created_at,omitempty"`
	UpdatedAt                      *Timestamp                    `json:"updated_at,omitempty"`
	LastUsedAt                     *Timestamp                    `json:"last_used_at,omitempty"`
	State                          *string                       `json:"state,omitempty"`
	URL                            *string                       `json:"url,omitempty"`
	GitStatus                      *CodespacesGitStatus          `json:"git_status,omitempty"`
	Location                       *string                       `json:"location,omitempty"`
	IdleTimeoutMinutes             *int                          `json:"idle_timeout_minutes,omitempty"`
	WebURL                         *string                       `json:"web_url,omitempty"`
	MachinesURL                    *string                       `json:"machines_url,omitempty"`
	StartURL                       *string                       `json:"start_url,omitempty"`
	StopURL                        *string                       `json:"stop_url,omitempty"`
	PullsURL                       *string                       `json:"pulls_url,omitempty"`
	RecentFolders                  []string                      `json:"recent_folders,omitempty"`
	RuntimeConstraints             *CodespacesRuntimeConstraints `json:"runtime_constraints,omitempty"`
	PendingOperation               *bool                         `json:"pending_operation,omitempty"`
	PendingOperationDisabledReason *string                       `json:"pending_operation_disabled_reason,omitempty"`
	IdleTimeoutNotice              *string                       `json:"idle_timeout_notice,omitempty"`
	RetentionPeriodMinutes         *int                          `json:"retention_period_minutes,omitempty"`
	RetentionExpiresAt             *Timestamp                    `json:"retention_expires_at,omitempty"`
	LastKnownStopNotice            *string                       `json:"last_known_stop_notice,omitempty"`
}

type CodespacesGitStatus struct {
	Ahead                 *int    `json:"ahead,omitempty"`
	Behind                *int    `json:"behind,omitempty"`
	HasUnpushedChanges    *bool   `json:"has_unpushed_changes,omitempty"`
	HasUncommittedChanges *bool   `json:"has_uncommitted_changes,omitempty"`
	Ref                   *string `json:"ref,omitempty"`
}

type CodespacesMachine struct {
	Name                 *string `json:"name,omitempty"`
	DisplayName          *string `json:"display_name,omitempty"`
	OperatingSystem      *string `json:"operating_system,omitempty"`
	StorageInBytes       *int64  `json:"storage_in_bytes,omitempty"`
	MemoryInBytes        *int64  `json:"memory_in_bytes,omitempty"`
	CPUs                 *int    `json:"cpus,omitempty"`
	PrebuildAvailability *string `json:"prebuild_availability,omitempty"`
}

type CodespacesRuntimeConstraints struct {
	AllowedPortPrivacySettings []string `json:"allowed_port_privacy_settings,omitempty"`
}

type ListCodespaces struct {
	TotalCount *int         `json:"total_count,omitempty"`
	Codespaces []*Codespace `json:"codespaces"`
}

//meta:operation GET /repos/{owner}/{repo}/codespaces
func (s *CodespacesService) ListInRepo(ctx context.Context, owner, repo string, opts *ListOptions) (*ListCodespaces, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ListCodespacesOptions struct {
	ListOptions
	RepositoryID int64 `url:"repository_id,omitempty"`
}

//meta:operation GET /user/codespaces
func (s *CodespacesService) List(ctx context.Context, opts *ListCodespacesOptions) (*ListCodespaces, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type CreateCodespaceOptions struct {
	Ref *string `json:"ref,omitempty"`

	Geo                        *string `json:"geo,omitempty"`
	ClientIP                   *string `json:"client_ip,omitempty"`
	Machine                    *string `json:"machine,omitempty"`
	DevcontainerPath           *string `json:"devcontainer_path,omitempty"`
	MultiRepoPermissionsOptOut *bool   `json:"multi_repo_permissions_opt_out,omitempty"`
	WorkingDirectory           *string `json:"working_directory,omitempty"`
	IdleTimeoutMinutes         *int    `json:"idle_timeout_minutes,omitempty"`
	DisplayName                *string `json:"display_name,omitempty"`

	RetentionPeriodMinutes *int    `json:"retention_period_minutes,omitempty"`
	Location               *string `json:"location,omitempty"`
}

type DevContainer struct {
	Path        string  `json:"path"`
	Name        *string `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}

type DevContainerConfigurations struct {
	Devcontainers []*DevContainer `json:"devcontainers"`
	TotalCount    int64           `json:"total_count"`
}

type CodespaceDefaults struct {
	Location         string  `json:"location"`
	DevcontainerPath *string `json:"devcontainer_path,omitempty"`
}

type CodespaceDefaultAttributes struct {
	BillableOwner *User              `json:"billable_owner"`
	Defaults      *CodespaceDefaults `json:"defaults"`
}

type CodespaceGetDefaultAttributesOptions struct {
	Ref *string `url:"ref,omitempty"`

	ClientIP *string `url:"client_ip,omitempty"`
}

type CodespacePullRequestOptions struct {
	PullRequestNumber int64 `json:"pull_request_number"`

	RepositoryID int64 `json:"repository_id"`
}

type CodespaceCreateForUserOptions struct {
	PullRequest *CodespacePullRequestOptions `json:"pull_request"`

	RepositoryID               int64   `json:"repository_id"`
	Ref                        *string `json:"ref,omitempty"`
	Geo                        *string `json:"geo,omitempty"`
	ClientIP                   *string `json:"client_ip,omitempty"`
	RetentionPeriodMinutes     *int    `json:"retention_period_minutes,omitempty"`
	Location                   *string `json:"location,omitempty"`
	Machine                    *string `json:"machine,omitempty"`
	DevcontainerPath           *string `json:"devcontainer_path,omitempty"`
	MultiRepoPermissionsOptOut *bool   `json:"multi_repo_permissions_opt_out,omitempty"`
	WorkingDirectory           *string `json:"working_directory,omitempty"`
	IdleTimeoutMinutes         *int    `json:"idle_timeout_minutes,omitempty"`
	DisplayName                *string `json:"display_name,omitempty"`
}

type UpdateCodespaceOptions struct {
	Machine *string `json:"machine,omitempty"`

	RecentFolders []string `json:"recent_folders,omitempty"`
}

type CodespaceExport struct {
	State       *string    `json:"state,omitempty"`
	CompletedAt *Timestamp `json:"completed_at,omitempty"`
	Branch      *string    `json:"branch,omitempty"`
	SHA         *string    `json:"sha,omitempty"`
	ID          *string    `json:"id,omitempty"`
	ExportURL   *string    `json:"export_url,omitempty"`
	HTMLURL     *string    `json:"html_url,omitempty"`
}

type PublishCodespaceOptions struct {
	Name *string `json:"name,omitempty"`

	Private *bool `json:"private,omitempty"`
}

type CodespacePermissions struct {
	Accepted bool `json:"accepted"`
}

//meta:operation POST /repos/{owner}/{repo}/codespaces
func (s *CodespacesService) CreateInRepo(ctx context.Context, owner, repo string, body *CreateCodespaceOptions) (*Codespace, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /user/codespaces/{codespace_name}/start
func (s *CodespacesService) Start(ctx context.Context, codespaceName string) (*Codespace, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /user/codespaces/{codespace_name}/stop
func (s *CodespacesService) Stop(ctx context.Context, codespaceName string) (*Codespace, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /user/codespaces/{codespace_name}
func (s *CodespacesService) Delete(ctx context.Context, codespaceName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/codespaces/devcontainers
func (s *CodespacesService) ListDevContainerConfigurations(ctx context.Context, owner, repo string, opts *ListOptions) (*DevContainerConfigurations, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/codespaces/new
func (s *CodespacesService) GetDefaultAttributes(ctx context.Context, owner, repo string, opts *CodespaceGetDefaultAttributesOptions) (*CodespaceDefaultAttributes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/codespaces/permissions_check
func (s *CodespacesService) CheckPermissions(ctx context.Context, owner, repo, ref, devcontainerPath string) (*CodespacePermissions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/pulls/{pull_number}/codespaces
func (s *CodespacesService) CreateFromPullRequest(ctx context.Context, owner, repo string, pullNumber int, body *CreateCodespaceOptions) (*Codespace, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /user/codespaces
func (s *CodespacesService) Create(ctx context.Context, body *CodespaceCreateForUserOptions) (*Codespace, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/codespaces/{codespace_name}
func (s *CodespacesService) Get(ctx context.Context, codespaceName string) (*Codespace, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /user/codespaces/{codespace_name}
func (s *CodespacesService) Update(ctx context.Context, codespaceName string, body *UpdateCodespaceOptions) (*Codespace, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /user/codespaces/{codespace_name}/exports
func (s *CodespacesService) ExportCodespace(ctx context.Context, codespaceName string) (*CodespaceExport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/codespaces/{codespace_name}/exports/{export_id}
func (s *CodespacesService) GetLatestCodespaceExport(ctx context.Context, codespaceName string) (*CodespaceExport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /user/codespaces/{codespace_name}/publish
func (s *CodespacesService) Publish(ctx context.Context, codespaceName string, body *PublishCodespaceOptions) (*Codespace, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
