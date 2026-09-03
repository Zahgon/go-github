package github

import (
	"context"
)

type TeamDiscussion struct {
	Author        *User      `json:"author,omitempty"`
	Body          *string    `json:"body,omitempty"`
	BodyHTML      *string    `json:"body_html,omitempty"`
	BodyVersion   *string    `json:"body_version,omitempty"`
	CommentsCount *int       `json:"comments_count,omitempty"`
	CommentsURL   *string    `json:"comments_url,omitempty"`
	CreatedAt     *Timestamp `json:"created_at,omitempty"`
	LastEditedAt  *Timestamp `json:"last_edited_at,omitempty"`
	HTMLURL       *string    `json:"html_url,omitempty"`
	NodeID        *string    `json:"node_id,omitempty"`
	Number        *int       `json:"number,omitempty"`
	Pinned        *bool      `json:"pinned,omitempty"`
	Private       *bool      `json:"private,omitempty"`
	TeamURL       *string    `json:"team_url,omitempty"`
	Title         *string    `json:"title,omitempty"`
	UpdatedAt     *Timestamp `json:"updated_at,omitempty"`
	URL           *string    `json:"url,omitempty"`
	Reactions     *Reactions `json:"reactions,omitempty"`
}

func (d TeamDiscussion) String() string { _ = "STUB: not implemented"; return "" }

type DiscussionListOptions struct {
	Direction string `url:"direction,omitempty"`

	ListOptions
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/discussions
func (s *TeamsService) ListDiscussionsByID(ctx context.Context, orgID, teamID int64, opts *DiscussionListOptions) ([]*TeamDiscussion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/discussions
func (s *TeamsService) ListDiscussionsBySlug(ctx context.Context, org, slug string, opts *DiscussionListOptions) ([]*TeamDiscussion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}
func (s *TeamsService) GetDiscussionByID(ctx context.Context, orgID, teamID int64, discussionNumber int) (*TeamDiscussion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}
func (s *TeamsService) GetDiscussionBySlug(ctx context.Context, org, slug string, discussionNumber int) (*TeamDiscussion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/teams/{team_slug}/discussions
func (s *TeamsService) CreateDiscussionByID(ctx context.Context, orgID, teamID int64, body TeamDiscussion) (*TeamDiscussion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/teams/{team_slug}/discussions
func (s *TeamsService) CreateDiscussionBySlug(ctx context.Context, org, slug string, body TeamDiscussion) (*TeamDiscussion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}
func (s *TeamsService) EditDiscussionByID(ctx context.Context, orgID, teamID int64, discussionNumber int, body TeamDiscussion) (*TeamDiscussion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}
func (s *TeamsService) EditDiscussionBySlug(ctx context.Context, org, slug string, discussionNumber int, body TeamDiscussion) (*TeamDiscussion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}
func (s *TeamsService) DeleteDiscussionByID(ctx context.Context, orgID, teamID int64, discussionNumber int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}
func (s *TeamsService) DeleteDiscussionBySlug(ctx context.Context, org, slug string, discussionNumber int) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
