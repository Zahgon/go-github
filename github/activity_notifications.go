package github

import (
	"context"
	"time"
)

type Notification struct {
	ID         *string              `json:"id,omitempty"`
	Repository *Repository          `json:"repository,omitempty"`
	Subject    *NotificationSubject `json:"subject,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Unread     *bool      `json:"unread,omitempty"`
	UpdatedAt  *Timestamp `json:"updated_at,omitempty"`
	LastReadAt *Timestamp `json:"last_read_at,omitempty"`
	URL        *string    `json:"url,omitempty"`
}

type NotificationSubject struct {
	Title            *string `json:"title,omitempty"`
	URL              *string `json:"url,omitempty"`
	LatestCommentURL *string `json:"latest_comment_url,omitempty"`
	Type             *string `json:"type,omitempty"`
}

type NotificationListOptions struct {
	All           bool      `url:"all,omitempty"`
	Participating bool      `url:"participating,omitempty"`
	Since         time.Time `url:"since,omitempty"`
	Before        time.Time `url:"before,omitempty"`

	ListOptions
}

//meta:operation GET /notifications
func (s *ActivityService) ListNotifications(ctx context.Context, opts *NotificationListOptions) ([]*Notification, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/notifications
func (s *ActivityService) ListRepositoryNotifications(ctx context.Context, owner, repo string, opts *NotificationListOptions) ([]*Notification, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type markReadOptions struct {
	LastReadAt Timestamp `json:"last_read_at,omitzero"`
}

//meta:operation PUT /notifications
func (s *ActivityService) MarkNotificationsRead(ctx context.Context, lastRead Timestamp) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/notifications
func (s *ActivityService) MarkRepositoryNotificationsRead(ctx context.Context, owner, repo string, lastRead Timestamp) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /notifications/threads/{thread_id}
func (s *ActivityService) GetThread(ctx context.Context, id string) (*Notification, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /notifications/threads/{thread_id}
func (s *ActivityService) MarkThreadRead(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /notifications/threads/{thread_id}
func (s *ActivityService) MarkThreadDone(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /notifications/threads/{thread_id}/subscription
func (s *ActivityService) GetThreadSubscription(ctx context.Context, id string) (*Subscription, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /notifications/threads/{thread_id}/subscription
func (s *ActivityService) SetThreadSubscription(ctx context.Context, id string, body *Subscription) (*Subscription, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /notifications/threads/{thread_id}/subscription
func (s *ActivityService) DeleteThreadSubscription(ctx context.Context, id string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
