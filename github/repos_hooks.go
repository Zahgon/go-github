package github

import (
	"context"
	"net/http"
)

type WebHookPayload = PushEvent

type WebHookCommit = HeadCommit

type WebHookAuthor = CommitAuthor

type Hook struct {
	CreatedAt    *Timestamp     `json:"created_at,omitempty"`
	UpdatedAt    *Timestamp     `json:"updated_at,omitempty"`
	URL          *string        `json:"url,omitempty"`
	ID           *int64         `json:"id,omitempty"`
	Type         *string        `json:"type,omitempty"`
	Name         *string        `json:"name,omitempty"`
	TestURL      *string        `json:"test_url,omitempty"`
	PingURL      *string        `json:"ping_url,omitempty"`
	LastResponse map[string]any `json:"last_response,omitempty"`

	Config *HookConfig `json:"config,omitempty"`
	Events []string    `json:"events,omitempty"`
	Active *bool       `json:"active,omitempty"`
}

func (h Hook) String() string { _ = "STUB: not implemented"; return "" }

type createHookRequest struct {
	Name   string      `json:"name"`
	Config *HookConfig `json:"config,omitempty"`
	Events []string    `json:"events,omitempty"`
	Active *bool       `json:"active,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/hooks
func (s *RepositoriesService) CreateHook(ctx context.Context, owner, repo string, hook *Hook) (*Hook, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/hooks
func (s *RepositoriesService) ListHooks(ctx context.Context, owner, repo string, opts *ListOptions) ([]*Hook, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/hooks/{hook_id}
func (s *RepositoriesService) GetHook(ctx context.Context, owner, repo string, id int64) (*Hook, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/hooks/{hook_id}
func (s *RepositoriesService) EditHook(ctx context.Context, owner, repo string, id int64, body *Hook) (*Hook, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/hooks/{hook_id}
func (s *RepositoriesService) DeleteHook(ctx context.Context, owner, repo string, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/hooks/{hook_id}/pings
func (s *RepositoriesService) PingHook(ctx context.Context, owner, repo string, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/hooks/{hook_id}/tests
func (s *RepositoriesService) TestHook(ctx context.Context, owner, repo string, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /hub
func (s *RepositoriesService) Subscribe(ctx context.Context, owner, repo, event, callback string, secret []byte) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /hub
func (s *RepositoriesService) Unsubscribe(ctx context.Context, owner, repo, event, callback string, secret []byte) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *RepositoriesService) createWebSubRequest(ctx context.Context, hubMode, owner, repo, event, callback string, secret []byte) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
