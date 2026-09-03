package github

import (
	"context"
)

type Subscription struct {
	Subscribed *bool      `json:"subscribed,omitempty"`
	Ignored    *bool      `json:"ignored,omitempty"`
	Reason     *string    `json:"reason,omitempty"`
	CreatedAt  *Timestamp `json:"created_at,omitempty"`
	URL        *string    `json:"url,omitempty"`

	RepositoryURL *string `json:"repository_url,omitempty"`

	ThreadURL *string `json:"thread_url,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/subscribers
func (s *ActivityService) ListWatchers(ctx context.Context, owner, repo string, opts *ListOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/subscriptions
//meta:operation GET /users/{username}/subscriptions
func (s *ActivityService) ListWatched(ctx context.Context, user string, opts *ListOptions) ([]*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/subscription
func (s *ActivityService) GetRepositorySubscription(ctx context.Context, owner, repo string) (*Subscription, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/subscription
func (s *ActivityService) SetRepositorySubscription(ctx context.Context, owner, repo string, body *Subscription) (*Subscription, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/subscription
func (s *ActivityService) DeleteRepositorySubscription(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
