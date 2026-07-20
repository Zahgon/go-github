package github

import (
	"context"
)

type TrafficReferrer struct {
	Referrer *string `json:"referrer,omitempty"`
	Count    *int    `json:"count,omitempty"`
	Uniques  *int    `json:"uniques,omitempty"`
}

type TrafficPath struct {
	Path    *string `json:"path,omitempty"`
	Title   *string `json:"title,omitempty"`
	Count   *int    `json:"count,omitempty"`
	Uniques *int    `json:"uniques,omitempty"`
}

type TrafficData struct {
	Timestamp *Timestamp `json:"timestamp,omitempty"`
	Count     *int       `json:"count,omitempty"`
	Uniques   *int       `json:"uniques,omitempty"`
}

type TrafficViews struct {
	Views   []*TrafficData `json:"views,omitempty"`
	Count   *int           `json:"count,omitempty"`
	Uniques *int           `json:"uniques,omitempty"`
}

type TrafficClones struct {
	Clones  []*TrafficData `json:"clones,omitempty"`
	Count   *int           `json:"count,omitempty"`
	Uniques *int           `json:"uniques,omitempty"`
}

type TrafficBreakdownOptions struct {
	Per string `url:"per,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/traffic/popular/referrers
func (s *RepositoriesService) ListTrafficReferrers(ctx context.Context, owner, repo string) ([]*TrafficReferrer, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/traffic/popular/paths
func (s *RepositoriesService) ListTrafficPaths(ctx context.Context, owner, repo string) ([]*TrafficPath, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/traffic/views
func (s *RepositoriesService) ListTrafficViews(ctx context.Context, owner, repo string, opts *TrafficBreakdownOptions) (*TrafficViews, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/traffic/clones
func (s *RepositoriesService) ListTrafficClones(ctx context.Context, owner, repo string, opts *TrafficBreakdownOptions) (*TrafficClones, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
