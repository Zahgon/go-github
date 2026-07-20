package github

import (
	"context"
)

type LicensesService service

type RepositoryLicense struct {
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`

	SHA         *string  `json:"sha,omitempty"`
	Size        *int     `json:"size,omitempty"`
	URL         *string  `json:"url,omitempty"`
	HTMLURL     *string  `json:"html_url,omitempty"`
	GitURL      *string  `json:"git_url,omitempty"`
	DownloadURL *string  `json:"download_url,omitempty"`
	Type        *string  `json:"type,omitempty"`
	Content     *string  `json:"content,omitempty"`
	Encoding    *string  `json:"encoding,omitempty"`
	License     *License `json:"license,omitempty"`
}

func (l RepositoryLicense) String() string { _ = "STUB: not implemented"; return "" }

type License struct {
	Key  *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`

	SPDXID         *string   `json:"spdx_id,omitempty"`
	HTMLURL        *string   `json:"html_url,omitempty"`
	Featured       *bool     `json:"featured,omitempty"`
	Description    *string   `json:"description,omitempty"`
	Implementation *string   `json:"implementation,omitempty"`
	Permissions    *[]string `json:"permissions,omitempty"`
	Conditions     *[]string `json:"conditions,omitempty"`
	Limitations    *[]string `json:"limitations,omitempty"`
	Body           *string   `json:"body,omitempty"`
}

func (l License) String() string { _ = "STUB: not implemented"; return "" }

type ListLicensesOptions struct {
	Featured *bool `url:"featured,omitempty"`

	ListOptions
}

//meta:operation GET /licenses
func (s *LicensesService) List(ctx context.Context, opts *ListLicensesOptions) ([]*License, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /licenses/{license}
func (s *LicensesService) Get(ctx context.Context, licenseName string) (*License, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
