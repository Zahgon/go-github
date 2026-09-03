package github

import (
	"context"
	"io"
	"net/http"
	"os"
)

type RepositoryRelease struct {
	TagName         string          `json:"tag_name"`
	TargetCommitish string          `json:"target_commitish"`
	Name            *string         `json:"name"`
	Body            *string         `json:"body,omitempty"`
	Draft           bool            `json:"draft"`
	Prerelease      bool            `json:"prerelease"`
	Immutable       *bool           `json:"immutable,omitempty"`
	ID              int64           `json:"id"`
	CreatedAt       Timestamp       `json:"created_at"`
	PublishedAt     *Timestamp      `json:"published_at"`
	UpdatedAt       *Timestamp      `json:"updated_at,omitempty"`
	URL             string          `json:"url"`
	HTMLURL         string          `json:"html_url"`
	AssetsURL       string          `json:"assets_url"`
	Assets          []*ReleaseAsset `json:"assets"`
	UploadURL       string          `json:"upload_url"`
	ZipballURL      *string         `json:"zipball_url"`
	TarballURL      *string         `json:"tarball_url"`
	Author          *User           `json:"author"`
	NodeID          string          `json:"node_id"`
	BodyHTML        *string         `json:"body_html,omitempty"`
	BodyText        *string         `json:"body_text,omitempty"`
	MentionsCount   *int            `json:"mentions_count,omitempty"`
	DiscussionURL   *string         `json:"discussion_url,omitempty"`
	Reactions       *Reactions      `json:"reactions,omitempty"`
}

func (r RepositoryRelease) String() string { _ = "STUB: not implemented"; return "" }

type RepositoryReleaseNotes struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

type GenerateNotesRequest struct {
	TagName               string  `json:"tag_name"`
	PreviousTagName       *string `json:"previous_tag_name,omitempty"`
	TargetCommitish       *string `json:"target_commitish,omitempty"`
	ConfigurationFilePath *string `json:"configuration_file_path,omitempty"`
}

type ReleaseAsset struct {
	ID                 *int64     `json:"id,omitempty"`
	URL                *string    `json:"url,omitempty"`
	Name               *string    `json:"name,omitempty"`
	Label              *string    `json:"label,omitempty"`
	State              *string    `json:"state,omitempty"`
	ContentType        *string    `json:"content_type,omitempty"`
	Size               *int       `json:"size,omitempty"`
	DownloadCount      *int       `json:"download_count,omitempty"`
	CreatedAt          *Timestamp `json:"created_at,omitempty"`
	UpdatedAt          *Timestamp `json:"updated_at,omitempty"`
	BrowserDownloadURL *string    `json:"browser_download_url,omitempty"`
	Uploader           *User      `json:"uploader,omitempty"`
	NodeID             *string    `json:"node_id,omitempty"`
	Digest             *string    `json:"digest,omitempty"`
}

type UpdateReleaseAssetRequest struct {
	Name  *string `json:"name,omitempty"`
	Label *string `json:"label,omitempty"`
	State *string `json:"state,omitempty"`
}

func (r ReleaseAsset) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/releases
func (s *RepositoriesService) ListReleases(ctx context.Context, owner, repo string, opts *ListOptions) ([]*RepositoryRelease, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/releases/{release_id}
func (s *RepositoriesService) GetRelease(ctx context.Context, owner, repo string, id int64) (*RepositoryRelease, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/releases/latest
func (s *RepositoriesService) GetLatestRelease(ctx context.Context, owner, repo string) (*RepositoryRelease, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/releases/tags/{tag}
func (s *RepositoriesService) GetReleaseByTag(ctx context.Context, owner, repo, tag string) (*RepositoryRelease, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/releases/generate-notes
func (s *RepositoriesService) GenerateReleaseNotes(ctx context.Context, owner, repo string, body GenerateNotesRequest) (*RepositoryReleaseNotes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *RepositoriesService) getSingleRelease(ctx context.Context, url string) (*RepositoryRelease, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type CreateReleaseRequest struct {
	TagName         string  `json:"tag_name"`
	TargetCommitish *string `json:"target_commitish,omitempty"`
	Name            *string `json:"name,omitempty"`
	Body            *string `json:"body,omitempty"`
	Draft           *bool   `json:"draft,omitempty"`
	Prerelease      *bool   `json:"prerelease,omitempty"`

	MakeLatest             *string `json:"make_latest,omitempty"`
	DiscussionCategoryName *string `json:"discussion_category_name,omitempty"`
	GenerateReleaseNotes   *bool   `json:"generate_release_notes,omitempty"`
}

type UpdateReleaseRequest struct {
	TagName         *string `json:"tag_name,omitempty"`
	TargetCommitish *string `json:"target_commitish,omitempty"`
	Name            *string `json:"name,omitempty"`
	Body            *string `json:"body,omitempty"`
	Draft           *bool   `json:"draft,omitempty"`
	Prerelease      *bool   `json:"prerelease,omitempty"`

	MakeLatest             *string `json:"make_latest,omitempty"`
	DiscussionCategoryName *string `json:"discussion_category_name,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/releases
func (s *RepositoriesService) CreateRelease(ctx context.Context, owner, repo string, body CreateReleaseRequest) (*RepositoryRelease, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/releases/{release_id}
func (s *RepositoriesService) UpdateRelease(ctx context.Context, owner, repo string, id int64, body UpdateReleaseRequest) (*RepositoryRelease, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/releases/{release_id}
func (s *RepositoriesService) DeleteRelease(ctx context.Context, owner, repo string, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/releases/{release_id}/assets
func (s *RepositoriesService) ListReleaseAssets(ctx context.Context, owner, repo string, id int64, opts *ListOptions) ([]*ReleaseAsset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/releases/assets/{asset_id}
func (s *RepositoriesService) GetReleaseAsset(ctx context.Context, owner, repo string, id int64) (*ReleaseAsset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/releases/assets/{asset_id}
func (s *RepositoriesService) DownloadReleaseAsset(ctx context.Context, owner, repo string, id int64, followRedirectsClient *http.Client) (rc io.ReadCloser, redirectURL string, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), "", nil
}

func (s *RepositoriesService) downloadReleaseAssetFromURL(ctx context.Context, followRedirectsClient *http.Client, url string) (rc io.ReadCloser, err error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

//meta:operation PATCH /repos/{owner}/{repo}/releases/assets/{asset_id}
func (s *RepositoriesService) UpdateReleaseAsset(ctx context.Context, owner, repo string, id int64, body UpdateReleaseAssetRequest) (*ReleaseAsset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/releases/assets/{asset_id}
func (s *RepositoriesService) DeleteReleaseAsset(ctx context.Context, owner, repo string, id int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/releases/{release_id}/assets
func (s *RepositoriesService) UploadReleaseAsset(ctx context.Context, owner, repo string, id int64, opts *UploadOptions, file *os.File) (*ReleaseAsset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/releases/{release_id}/assets
func (s *RepositoriesService) UploadReleaseAssetFromRelease(
	ctx context.Context,
	release *RepositoryRelease,
	opts *UploadOptions,
	reader io.Reader,
	size int64,
) (*ReleaseAsset, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
