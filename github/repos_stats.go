package github

import (
	"context"
)

type ContributorStats struct {
	Author *Contributor   `json:"author,omitempty"`
	Total  *int           `json:"total,omitempty"`
	Weeks  []*WeeklyStats `json:"weeks,omitempty"`
}

func (c ContributorStats) String() string { _ = "STUB: not implemented"; return "" }

type WeeklyStats struct {
	Week      *Timestamp `json:"w,omitempty"`
	Additions *int       `json:"a,omitempty"`
	Deletions *int       `json:"d,omitempty"`
	Commits   *int       `json:"c,omitempty"`
}

func (w WeeklyStats) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/stats/contributors
func (s *RepositoriesService) ListContributorsStats(ctx context.Context, owner, repo string) ([]*ContributorStats, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type WeeklyCommitActivity struct {
	Days  []int      `json:"days,omitempty"`
	Total *int       `json:"total,omitempty"`
	Week  *Timestamp `json:"week,omitempty"`
}

func (w WeeklyCommitActivity) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/stats/commit_activity
func (s *RepositoriesService) ListCommitActivity(ctx context.Context, owner, repo string) ([]*WeeklyCommitActivity, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/stats/code_frequency
func (s *RepositoriesService) ListCodeFrequency(ctx context.Context, owner, repo string) ([]*WeeklyStats, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type RepositoryParticipation struct {
	All   []int `json:"all,omitempty"`
	Owner []int `json:"owner,omitempty"`
}

func (r RepositoryParticipation) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/stats/participation
func (s *RepositoriesService) ListParticipation(ctx context.Context, owner, repo string) (*RepositoryParticipation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type PunchCard struct {
	Day     *int
	Hour    *int
	Commits *int
}

//meta:operation GET /repos/{owner}/{repo}/stats/punch_card
func (s *RepositoriesService) ListPunchCard(ctx context.Context, owner, repo string) ([]*PunchCard, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
