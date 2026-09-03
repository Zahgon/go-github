package github

import (
	"context"
)

type SearchService service

type SearchOptions struct {
	Sort string `url:"sort,omitempty"`

	Order string `url:"order,omitempty"`

	TextMatch bool `url:"-"`

	AdvancedSearch *bool `url:"advanced_search,omitempty"`

	SearchType string `url:"search_type,omitempty"`

	ListOptions
}

type searchParameters struct {
	Query        string
	RepositoryID *int64
}

type RepositoriesSearchResult struct {
	Total             *int          `json:"total_count,omitempty"`
	IncompleteResults *bool         `json:"incomplete_results,omitempty"`
	Repositories      []*Repository `json:"items,omitempty"`
}

//meta:operation GET /search/repositories
func (s *SearchService) Repositories(ctx context.Context, query string, opts *SearchOptions) (*RepositoriesSearchResult, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type TopicsSearchResult struct {
	Total             *int           `json:"total_count,omitempty"`
	IncompleteResults *bool          `json:"incomplete_results,omitempty"`
	Topics            []*TopicResult `json:"items,omitempty"`
}

type TopicResult struct {
	Name             *string    `json:"name,omitempty"`
	DisplayName      *string    `json:"display_name,omitempty"`
	ShortDescription *string    `json:"short_description,omitempty"`
	Description      *string    `json:"description,omitempty"`
	CreatedBy        *string    `json:"created_by,omitempty"`
	CreatedAt        *Timestamp `json:"created_at,omitempty"`
	UpdatedAt        *string    `json:"updated_at,omitempty"`
	Featured         *bool      `json:"featured,omitempty"`
	Curated          *bool      `json:"curated,omitempty"`
	Score            *float64   `json:"score,omitempty"`
}

//meta:operation GET /search/topics
func (s *SearchService) Topics(ctx context.Context, query string, opts *SearchOptions) (*TopicsSearchResult, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type CommitsSearchResult struct {
	Total             *int            `json:"total_count,omitempty"`
	IncompleteResults *bool           `json:"incomplete_results,omitempty"`
	Commits           []*CommitResult `json:"items,omitempty"`
}

type CommitResult struct {
	SHA         *string   `json:"sha,omitempty"`
	Commit      *Commit   `json:"commit,omitempty"`
	Author      *User     `json:"author,omitempty"`
	Committer   *User     `json:"committer,omitempty"`
	Parents     []*Commit `json:"parents,omitempty"`
	HTMLURL     *string   `json:"html_url,omitempty"`
	URL         *string   `json:"url,omitempty"`
	CommentsURL *string   `json:"comments_url,omitempty"`

	Repository *Repository `json:"repository,omitempty"`
	Score      *float64    `json:"score,omitempty"`
}

//meta:operation GET /search/commits
func (s *SearchService) Commits(ctx context.Context, query string, opts *SearchOptions) (*CommitsSearchResult, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type IssuesSearchResult struct {
	Total             *int     `json:"total_count,omitempty"`
	IncompleteResults *bool    `json:"incomplete_results,omitempty"`
	SearchType        *string  `json:"search_type,omitempty"`
	Issues            []*Issue `json:"items,omitempty"`
}

//meta:operation GET /search/issues
func (s *SearchService) Issues(ctx context.Context, query string, opts *SearchOptions) (*IssuesSearchResult, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type UsersSearchResult struct {
	Total             *int    `json:"total_count,omitempty"`
	IncompleteResults *bool   `json:"incomplete_results,omitempty"`
	Users             []*User `json:"items,omitempty"`
}

//meta:operation GET /search/users
func (s *SearchService) Users(ctx context.Context, query string, opts *SearchOptions) (*UsersSearchResult, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type Match struct {
	Text    *string `json:"text,omitempty"`
	Indices []int   `json:"indices,omitempty"`
}

type TextMatch struct {
	ObjectURL  *string  `json:"object_url,omitempty"`
	ObjectType *string  `json:"object_type,omitempty"`
	Property   *string  `json:"property,omitempty"`
	Fragment   *string  `json:"fragment,omitempty"`
	Matches    []*Match `json:"matches,omitempty"`
}

func (tm TextMatch) String() string { _ = "STUB: not implemented"; return "" }

type CodeSearchResult struct {
	Total             *int          `json:"total_count,omitempty"`
	IncompleteResults *bool         `json:"incomplete_results,omitempty"`
	CodeResults       []*CodeResult `json:"items,omitempty"`
}

type CodeResult struct {
	Name        *string      `json:"name,omitempty"`
	Path        *string      `json:"path,omitempty"`
	SHA         *string      `json:"sha,omitempty"`
	HTMLURL     *string      `json:"html_url,omitempty"`
	Repository  *Repository  `json:"repository,omitempty"`
	TextMatches []*TextMatch `json:"text_matches,omitempty"`
}

func (c CodeResult) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /search/code
func (s *SearchService) Code(ctx context.Context, query string, opts *SearchOptions) (*CodeSearchResult, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type LabelsSearchResult struct {
	Total             *int           `json:"total_count,omitempty"`
	IncompleteResults *bool          `json:"incomplete_results,omitempty"`
	Labels            []*LabelResult `json:"items,omitempty"`
}

type LabelResult struct {
	ID          *int64   `json:"id,omitempty"`
	URL         *string  `json:"url,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Color       *string  `json:"color,omitempty"`
	Default     *bool    `json:"default,omitempty"`
	Description *string  `json:"description,omitempty"`
	Score       *float64 `json:"score,omitempty"`
}

func (l LabelResult) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /search/labels
func (s *SearchService) Labels(ctx context.Context, repoID int64, query string, opts *SearchOptions) (*LabelsSearchResult, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *SearchService) search(ctx context.Context, searchType string, parameters *searchParameters, opts *SearchOptions, result any) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
