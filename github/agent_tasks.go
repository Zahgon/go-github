package github

import (
	"context"
	"time"
)

type AgentTasksService service

type AgentTask struct {
	ID          string            `json:"id"`
	URL         *string           `json:"url,omitempty"`
	HTMLURL     *string           `json:"html_url,omitempty"`
	Name        *string           `json:"name,omitempty"`
	Creator     *AgentTaskCreator `json:"creator,omitempty"`
	CreatorType *string           `json:"creator_type,omitempty"`

	UserCollaborators []*User              `json:"user_collaborators,omitempty"`
	Owner             *AgentTaskOwner      `json:"owner,omitempty"`
	Repository        *AgentTaskRepository `json:"repository,omitempty"`
	State             string               `json:"state"`
	SessionCount      *int                 `json:"session_count,omitempty"`
	Artifacts         []*AgentTaskArtifact `json:"artifacts,omitempty"`
	ArchivedAt        *Timestamp           `json:"archived_at,omitempty"`
	CreatedAt         Timestamp            `json:"created_at"`
	UpdatedAt         *Timestamp           `json:"updated_at,omitempty"`
	Sessions          []*AgentTaskSession  `json:"sessions,omitempty"`
}

type AgentTaskCreator struct {
	ID *int64 `json:"id,omitempty"`
}

type AgentTaskOwner struct {
	ID *int64 `json:"id,omitempty"`
}

type AgentTaskRepository struct {
	ID *int64 `json:"id,omitempty"`
}

type AgentTaskArtifact struct {
	Provider string                 `json:"provider"`
	Type     string                 `json:"type"`
	Data     *AgentTaskArtifactData `json:"data,omitempty"`
}

type AgentTaskArtifactData struct {
	ID       *int64  `json:"id,omitempty"`
	GlobalID *string `json:"global_id,omitempty"`
	HeadRef  *string `json:"head_ref,omitempty"`
	BaseRef  *string `json:"base_ref,omitempty"`
}

type AgentTaskSession struct {
	ID          string                 `json:"id"`
	Name        *string                `json:"name,omitempty"`
	User        *User                  `json:"user,omitempty"`
	Owner       *AgentTaskOwner        `json:"owner,omitempty"`
	Repository  *AgentTaskRepository   `json:"repository,omitempty"`
	TaskID      *string                `json:"task_id,omitempty"`
	State       string                 `json:"state"`
	CreatedAt   Timestamp              `json:"created_at"`
	UpdatedAt   *Timestamp             `json:"updated_at,omitempty"`
	CompletedAt *Timestamp             `json:"completed_at,omitempty"`
	Prompt      *string                `json:"prompt,omitempty"`
	HeadRef     *string                `json:"head_ref,omitempty"`
	BaseRef     *string                `json:"base_ref,omitempty"`
	Model       *string                `json:"model,omitempty"`
	Error       *AgentTaskSessionError `json:"error,omitempty"`
}

type AgentTaskSessionError struct {
	Message *string `json:"message,omitempty"`
}

type AgentTaskList struct {
	Tasks              []*AgentTask `json:"tasks"`
	TotalActiveCount   *int         `json:"total_active_count,omitempty"`
	TotalArchivedCount *int         `json:"total_archived_count,omitempty"`
}

type AgentTaskListOptions struct {
	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	State string `url:"state,omitempty"`

	IsArchived bool `url:"is_archived,omitempty"`

	Since *time.Time `url:"since,omitempty"`

	ListOptions
}

type AgentTaskListByRepoOptions struct {
	AgentTaskListOptions

	CreatorID []int64 `url:"creator_id,omitempty"`
}

type CreateAgentTaskRequest struct {
	Prompt string `json:"prompt"`

	Model *string `json:"model,omitempty"`

	CreatePullRequest *bool `json:"create_pull_request,omitempty"`

	HeadRef *string `json:"head_ref,omitempty"`

	BaseRef *string `json:"base_ref,omitempty"`
}

//meta:operation GET /agents/repos/{owner}/{repo}/tasks
func (s *AgentTasksService) ListByRepo(ctx context.Context, owner, repo string, opts *AgentTaskListByRepoOptions) (*AgentTaskList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /agents/repos/{owner}/{repo}/tasks
func (s *AgentTasksService) Create(ctx context.Context, owner, repo string, body CreateAgentTaskRequest) (*AgentTask, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /agents/repos/{owner}/{repo}/tasks/{task_id}
func (s *AgentTasksService) GetByRepoAndID(ctx context.Context, owner, repo, taskID string) (*AgentTask, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /agents/tasks
func (s *AgentTasksService) List(ctx context.Context, opts *AgentTaskListOptions) (*AgentTaskList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /agents/tasks/{task_id}
func (s *AgentTasksService) Get(ctx context.Context, taskID string) (*AgentTask, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
