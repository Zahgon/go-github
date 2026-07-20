package github

import (
	"context"
)

type ActionsCache struct {
	ID             *int64     `json:"id,omitempty" url:"-"`
	Ref            *string    `json:"ref,omitempty" url:"ref"`
	Key            *string    `json:"key,omitempty" url:"key"`
	Version        *string    `json:"version,omitempty" url:"-"`
	LastAccessedAt *Timestamp `json:"last_accessed_at,omitempty" url:"-"`
	CreatedAt      *Timestamp `json:"created_at,omitempty" url:"-"`
	SizeInBytes    *int64     `json:"size_in_bytes,omitempty" url:"-"`
}

type ActionsCacheList struct {
	TotalCount    int             `json:"total_count"`
	ActionsCaches []*ActionsCache `json:"actions_caches,omitempty"`
}

type ActionsCacheUsage struct {
	FullName                string `json:"full_name"`
	ActiveCachesSizeInBytes int64  `json:"active_caches_size_in_bytes"`
	ActiveCachesCount       int    `json:"active_caches_count"`
}

type ActionsCacheUsageList struct {
	TotalCount     int                  `json:"total_count"`
	RepoCacheUsage []*ActionsCacheUsage `json:"repository_cache_usages,omitempty"`
}

type TotalCacheUsage struct {
	TotalActiveCachesUsageSizeInBytes int64 `json:"total_active_caches_size_in_bytes"`
	TotalActiveCachesCount            int   `json:"total_active_caches_count"`
}

type ActionsCacheListOptions struct {
	ListOptions

	Ref *string `url:"ref,omitempty"`
	Key *string `url:"key,omitempty"`

	Sort *string `url:"sort,omitempty"`

	Direction *string `url:"direction,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/actions/caches
func (s *ActionsService) ListCaches(ctx context.Context, owner, repo string, opts *ActionsCacheListOptions) (*ActionsCacheList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/actions/caches
func (s *ActionsService) DeleteCachesByKey(ctx context.Context, owner, repo, key string, ref *string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/actions/caches/{cache_id}
func (s *ActionsService) DeleteCachesByID(ctx context.Context, owner, repo string, cacheID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/cache/usage
func (s *ActionsService) GetCacheUsageForRepo(ctx context.Context, owner, repo string) (*ActionsCacheUsage, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/cache/usage-by-repository
func (s *ActionsService) ListCacheUsageByRepoForOrg(ctx context.Context, org string, opts *ListOptions) (*ActionsCacheUsageList, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/cache/usage
func (s *ActionsService) GetTotalCacheUsageForOrg(ctx context.Context, org string) (*TotalCacheUsage, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/actions/cache/usage
func (s *ActionsService) GetTotalCacheUsageForEnterprise(ctx context.Context, enterprise string) (*TotalCacheUsage, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
