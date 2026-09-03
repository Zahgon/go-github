package github

import (
	"context"
)

type ReactionsService service

type Reaction struct {
	ID     *int64  `json:"id,omitempty"`
	User   *User   `json:"user,omitempty"`
	NodeID *string `json:"node_id,omitempty"`

	Content   *string    `json:"content,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
}

type Reactions struct {
	TotalCount *int    `json:"total_count,omitempty"`
	PlusOne    *int    `json:"+1,omitempty"`
	MinusOne   *int    `json:"-1,omitempty"`
	Laugh      *int    `json:"laugh,omitempty"`
	Confused   *int    `json:"confused,omitempty"`
	Heart      *int    `json:"heart,omitempty"`
	Hooray     *int    `json:"hooray,omitempty"`
	Rocket     *int    `json:"rocket,omitempty"`
	Eyes       *int    `json:"eyes,omitempty"`
	URL        *string `json:"url,omitempty"`
}

func (r Reaction) String() string { _ = "STUB: not implemented"; return "" }

type ListReactionOptions struct {
	Content string `url:"content,omitempty"`

	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/comments/{comment_id}/reactions
func (s *ReactionsService) ListCommentReactions(ctx context.Context, owner, repo string, id int64, opts *ListReactionOptions) ([]*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/comments/{comment_id}/reactions
func (s *ReactionsService) CreateCommentReaction(ctx context.Context, owner, repo string, id int64, content string) (*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/comments/{comment_id}/reactions/{reaction_id}
func (s *ReactionsService) DeleteCommentReaction(ctx context.Context, owner, repo string, commentID, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/comments/{comment_id}/reactions/{reaction_id}
func (s *ReactionsService) DeleteCommentReactionByID(ctx context.Context, repoID, commentID, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/issues/{issue_number}/reactions
func (s *ReactionsService) ListIssueReactions(ctx context.Context, owner, repo string, number int, opts *ListReactionOptions) ([]*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/issues/{issue_number}/reactions
func (s *ReactionsService) CreateIssueReaction(ctx context.Context, owner, repo string, number int, content string) (*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/issues/{issue_number}/reactions/{reaction_id}
func (s *ReactionsService) DeleteIssueReaction(ctx context.Context, owner, repo string, issueNumber int, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/issues/{issue_number}/reactions/{reaction_id}
func (s *ReactionsService) DeleteIssueReactionByID(ctx context.Context, repoID, issueNumber int, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/issues/comments/{comment_id}/reactions
func (s *ReactionsService) ListIssueCommentReactions(ctx context.Context, owner, repo string, id int64, opts *ListReactionOptions) ([]*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/issues/comments/{comment_id}/reactions
func (s *ReactionsService) CreateIssueCommentReaction(ctx context.Context, owner, repo string, id int64, content string) (*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/issues/comments/{comment_id}/reactions/{reaction_id}
func (s *ReactionsService) DeleteIssueCommentReaction(ctx context.Context, owner, repo string, commentID, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/issues/comments/{comment_id}/reactions/{reaction_id}
func (s *ReactionsService) DeleteIssueCommentReactionByID(ctx context.Context, repoID, commentID, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions
func (s *ReactionsService) ListPullRequestCommentReactions(ctx context.Context, owner, repo string, id int64, opts *ListReactionOptions) ([]*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions
func (s *ReactionsService) CreatePullRequestCommentReaction(ctx context.Context, owner, repo string, id int64, content string) (*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions/{reaction_id}
func (s *ReactionsService) DeletePullRequestCommentReaction(ctx context.Context, owner, repo string, commentID, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions/{reaction_id}
func (s *ReactionsService) DeletePullRequestCommentReactionByID(ctx context.Context, repoID, commentID, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /teams/{team_id}/discussions/{discussion_number}/reactions
func (s *ReactionsService) ListTeamDiscussionReactions(ctx context.Context, teamID int64, discussionNumber int, opts *ListReactionOptions) ([]*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /teams/{team_id}/discussions/{discussion_number}/reactions
func (s *ReactionsService) CreateTeamDiscussionReaction(ctx context.Context, teamID int64, discussionNumber int, content string) (*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/reactions/{reaction_id}
func (s *ReactionsService) DeleteTeamDiscussionReaction(ctx context.Context, org, teamSlug string, discussionNumber int, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/reactions
func (s *ReactionsService) DeleteTeamDiscussionReactionByOrgIDAndTeamID(ctx context.Context, orgID, teamID, discussionNumber int, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}/reactions
func (s *ReactionsService) ListTeamDiscussionCommentReactions(ctx context.Context, teamID int64, discussionNumber, commentNumber int, opts *ListReactionOptions) ([]*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}/reactions
func (s *ReactionsService) CreateTeamDiscussionCommentReaction(ctx context.Context, teamID int64, discussionNumber, commentNumber int, content string) (*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}/reactions/{reaction_id}
func (s *ReactionsService) DeleteTeamDiscussionCommentReaction(ctx context.Context, org, teamSlug string, discussionNumber, commentNumber int, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}/reactions
func (s *ReactionsService) DeleteTeamDiscussionCommentReactionByOrgIDAndTeamID(ctx context.Context, orgID, teamID, discussionNumber, commentNumber int, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ReactionsService) deleteReaction(ctx context.Context, url string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/releases/{release_id}/reactions
func (s *ReactionsService) CreateReleaseReaction(ctx context.Context, owner, repo string, releaseID int64, content string) (*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/releases/{release_id}/reactions
func (s *ReactionsService) ListReleaseReactions(ctx context.Context, owner, repo string, releaseID int64, opts *ListReactionOptions) ([]*Reaction, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/releases/{release_id}/reactions/{reaction_id}
func (s *ReactionsService) DeleteReleaseReaction(ctx context.Context, owner, repo string, releaseID, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/releases/{release_id}/reactions/{reaction_id}
func (s *ReactionsService) DeleteReleaseReactionByID(ctx context.Context, repoID, releaseID, reactionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
