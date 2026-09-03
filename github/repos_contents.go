package github

import (
	"context"
	"errors"
	"io"
	"net/url"
)

var ErrContentsDirectory = errors.New("contents not available for directory")

var ErrContentsSubmodule = errors.New("contents not available for submodule")

var ErrContentsNoDownloadURL = errors.New("contents download url is empty")

type RepositoryContent struct {
	Type *string `json:"type,omitempty"`

	Target   *string `json:"target,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	Size     *int    `json:"size,omitempty"`
	Name     *string `json:"name,omitempty"`
	Path     *string `json:"path,omitempty"`

	Content         *string `json:"content,omitempty"`
	SHA             *string `json:"sha,omitempty"`
	URL             *string `json:"url,omitempty"`
	GitURL          *string `json:"git_url,omitempty"`
	HTMLURL         *string `json:"html_url,omitempty"`
	DownloadURL     *string `json:"download_url,omitempty"`
	SubmoduleGitURL *string `json:"submodule_git_url,omitempty"`
}

type RepositoryContentResponse struct {
	Content *RepositoryContent `json:"content,omitempty"`
	Commit  `json:"commit"`
}

type RepositoryContentFileOptions struct {
	Message   *string       `json:"message,omitempty"`
	Content   []byte        `json:"content"`
	SHA       *string       `json:"sha,omitempty"`
	Branch    *string       `json:"branch,omitempty"`
	Author    *CommitAuthor `json:"author,omitempty"`
	Committer *CommitAuthor `json:"committer,omitempty"`
}

type RepositoryContentGetOptions struct {
	Ref string `url:"ref,omitempty"`
}

func (r RepositoryContent) String() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContent) GetContent() (string, error) { _ = "STUB: not implemented"; return "", nil }

//meta:operation GET /repos/{owner}/{repo}/readme
func (s *RepositoriesService) GetReadme(ctx context.Context, owner, repo string, opts *RepositoryContentGetOptions) (*RepositoryContent, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/contents/{path}
func (s *RepositoriesService) DownloadContents(ctx context.Context, owner, repo, filepath string, opts *RepositoryContentGetOptions) (io.ReadCloser, *Response, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/contents/{path}
func (s *RepositoriesService) DownloadContentsWithMeta(ctx context.Context, owner, repo, filepath string, opts *RepositoryContentGetOptions) (io.ReadCloser, *RepositoryContent, *Response, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/contents/{path}
func (s *RepositoriesService) GetContents(ctx context.Context, owner, repo, path string, opts *RepositoryContentGetOptions) (fileContent *RepositoryContent, directoryContent []*RepositoryContent, resp *Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/contents/{path}
func (s *RepositoriesService) CreateFile(ctx context.Context, owner, repo, path string, body *RepositoryContentFileOptions) (*RepositoryContentResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/contents/{path}
func (s *RepositoriesService) UpdateFile(ctx context.Context, owner, repo, path string, body *RepositoryContentFileOptions) (*RepositoryContentResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/contents/{path}
func (s *RepositoriesService) DeleteFile(ctx context.Context, owner, repo, path string, opts *RepositoryContentFileOptions) (*RepositoryContentResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ArchiveFormat string

const (
	Tarball ArchiveFormat = "tarball"

	Zipball ArchiveFormat = "zipball"
)

//meta:operation GET /repos/{owner}/{repo}/tarball/{ref}
//meta:operation GET /repos/{owner}/{repo}/zipball/{ref}
func (s *RepositoriesService) GetArchiveLink(ctx context.Context, owner, repo string, archiveformat ArchiveFormat, opts *RepositoryContentGetOptions, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *RepositoriesService) getArchiveLinkWithoutRateLimit(ctx context.Context, u string, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *RepositoriesService) getArchiveLinkWithRateLimit(ctx context.Context, u string, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
