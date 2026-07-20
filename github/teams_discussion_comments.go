package github

import (
	"context"
)

type DiscussionComment struct {
	Author        *User      `json:"author,omitempty"`
	Body          *string    `json:"body,omitempty"`
	BodyHTML      *string    `json:"body_html,omitempty"`
	BodyVersion   *string    `json:"body_version,omitempty"`
	CreatedAt     *Timestamp `json:"created_at,omitempty"`
	LastEditedAt  *Timestamp `json:"last_edited_at,omitempty"`
	DiscussionURL *string    `json:"discussion_url,omitempty"`
	HTMLURL       *string    `json:"html_url,omitempty"`
	NodeID        *string    `json:"node_id,omitempty"`
	Number        *int       `json:"number,omitempty"`
	UpdatedAt     *Timestamp `json:"updated_at,omitempty"`
	URL           *string    `json:"url,omitempty"`
	Reactions     *Reactions `json:"reactions,omitempty"`
}

func (c DiscussionComment) String() string { _ = "STUB: not implemented"; return "" }

type DiscussionCommentListOptions struct {
	Direction string `url:"direction,omitempty"`
	ListOptions
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments
func (s *TeamsService) ListCommentsByID(ctx context.Context, orgID, teamID int64, discussionNumber int, opts *DiscussionCommentListOptions) ([]*DiscussionComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments
func (s *TeamsService) ListCommentsBySlug(ctx context.Context, org, slug string, discussionNumber int, opts *DiscussionCommentListOptions) ([]*DiscussionComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}
func (s *TeamsService) GetCommentByID(ctx context.Context, orgID, teamID int64, discussionNumber, commentNumber int) (*DiscussionComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}
func (s *TeamsService) GetCommentBySlug(ctx context.Context, org, slug string, discussionNumber, commentNumber int) (*DiscussionComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments
func (s *TeamsService) CreateCommentByID(ctx context.Context, orgID, teamID int64, discussionNumber int, body DiscussionComment) (*DiscussionComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments
func (s *TeamsService) CreateCommentBySlug(ctx context.Context, org, slug string, discussionNumber int, body DiscussionComment) (*DiscussionComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}
func (s *TeamsService) EditCommentByID(ctx context.Context, orgID, teamID int64, discussionNumber, commentNumber int, body DiscussionComment) (*DiscussionComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}
func (s *TeamsService) EditCommentBySlug(ctx context.Context, org, slug string, discussionNumber, commentNumber int, body DiscussionComment) (*DiscussionComment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}
func (s *TeamsService) DeleteCommentByID(ctx context.Context, orgID, teamID int64, discussionNumber, commentNumber int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}
func (s *TeamsService) DeleteCommentBySlug(ctx context.Context, org, slug string, discussionNumber, commentNumber int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
