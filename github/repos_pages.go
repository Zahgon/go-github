package github

import (
	"context"
)

type Pages struct {
	URL              *string                `json:"url,omitempty"`
	Status           *string                `json:"status,omitempty"`
	CNAME            *string                `json:"cname,omitempty"`
	Custom404        *bool                  `json:"custom_404,omitempty"`
	HTMLURL          *string                `json:"html_url,omitempty"`
	BuildType        *string                `json:"build_type,omitempty"`
	Source           *PagesSource           `json:"source,omitempty"`
	Public           *bool                  `json:"public,omitempty"`
	HTTPSCertificate *PagesHTTPSCertificate `json:"https_certificate,omitempty"`
	HTTPSEnforced    *bool                  `json:"https_enforced,omitempty"`
}

type PagesSource struct {
	Branch *string `json:"branch,omitempty"`
	Path   *string `json:"path,omitempty"`
}

type PagesError struct {
	Message *string `json:"message,omitempty"`
}

type PagesBuild struct {
	URL       *string     `json:"url,omitempty"`
	Status    *string     `json:"status,omitempty"`
	Error     *PagesError `json:"error,omitempty"`
	Pusher    *User       `json:"pusher,omitempty"`
	Commit    *string     `json:"commit,omitempty"`
	Duration  *int        `json:"duration,omitempty"`
	CreatedAt *Timestamp  `json:"created_at,omitempty"`
	UpdatedAt *Timestamp  `json:"updated_at,omitempty"`
}

type PagesDomain struct {
	Host                          *string `json:"host,omitempty"`
	URI                           *string `json:"uri,omitempty"`
	Nameservers                   *string `json:"nameservers,omitempty"`
	DNSResolves                   *bool   `json:"dns_resolves,omitempty"`
	IsProxied                     *bool   `json:"is_proxied,omitempty"`
	IsCloudflareIP                *bool   `json:"is_cloudflare_ip,omitempty"`
	IsFastlyIP                    *bool   `json:"is_fastly_ip,omitempty"`
	IsOldIPAddress                *bool   `json:"is_old_ip_address,omitempty"`
	IsARecord                     *bool   `json:"is_a_record,omitempty"`
	HasCNAMERecord                *bool   `json:"has_cname_record,omitempty"`
	HasMXRecordsPresent           *bool   `json:"has_mx_records_present,omitempty"`
	IsValidDomain                 *bool   `json:"is_valid_domain,omitempty"`
	IsApexDomain                  *bool   `json:"is_apex_domain,omitempty"`
	ShouldBeARecord               *bool   `json:"should_be_a_record,omitempty"`
	IsCNAMEToGithubUserDomain     *bool   `json:"is_cname_to_github_user_domain,omitempty"`
	IsCNAMEToPagesDotGithubDotCom *bool   `json:"is_cname_to_pages_dot_github_dot_com,omitempty"`
	IsCNAMEToFastly               *bool   `json:"is_cname_to_fastly,omitempty"`
	IsPointedToGithubPagesIP      *bool   `json:"is_pointed_to_github_pages_ip,omitempty"`
	IsNonGithubPagesIPPresent     *bool   `json:"is_non_github_pages_ip_present,omitempty"`
	IsPagesDomain                 *bool   `json:"is_pages_domain,omitempty"`
	IsServedByPages               *bool   `json:"is_served_by_pages,omitempty"`
	IsValid                       *bool   `json:"is_valid,omitempty"`
	Reason                        *string `json:"reason,omitempty"`
	RespondsToHTTPS               *bool   `json:"responds_to_https,omitempty"`
	EnforcesHTTPS                 *bool   `json:"enforces_https,omitempty"`
	HTTPSError                    *string `json:"https_error,omitempty"`
	IsHTTPSEligible               *bool   `json:"is_https_eligible,omitempty"`
	CAAError                      *string `json:"caa_error,omitempty"`
}

type PagesHealthCheckResponse struct {
	Domain    *PagesDomain `json:"domain,omitempty"`
	AltDomain *PagesDomain `json:"alt_domain,omitempty"`
}

type PagesHTTPSCertificate struct {
	State       *string  `json:"state,omitempty"`
	Description *string  `json:"description,omitempty"`
	Domains     []string `json:"domains,omitempty"`

	ExpiresAt *string `json:"expires_at,omitempty"`
}

type createPagesRequest struct {
	BuildType *string      `json:"build_type,omitempty"`
	Source    *PagesSource `json:"source,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/pages
func (s *RepositoriesService) EnablePages(ctx context.Context, owner, repo string, pages *Pages) (*Pages, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type PagesUpdate struct {
	CNAME *string `json:"cname"`

	BuildType *string `json:"build_type,omitempty"`

	Source *PagesSource `json:"source,omitempty"`

	Public *bool `json:"public,omitempty"`

	HTTPSEnforced *bool `json:"https_enforced,omitempty"`
}

//meta:operation PUT /repos/{owner}/{repo}/pages
func (s *RepositoriesService) UpdatePages(ctx context.Context, owner, repo string, body *PagesUpdate) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PagesUpdateWithoutCNAME struct {
	BuildType     *string      `json:"build_type,omitempty"`
	Source        *PagesSource `json:"source,omitempty"`
	Public        *bool        `json:"public,omitempty"`
	HTTPSEnforced *bool        `json:"https_enforced,omitempty"`
}

//meta:operation PUT /repos/{owner}/{repo}/pages
func (s *RepositoriesService) UpdatePagesGHES(ctx context.Context, owner, repo string, body *PagesUpdateWithoutCNAME) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/pages
func (s *RepositoriesService) DisablePages(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pages
func (s *RepositoriesService) GetPagesInfo(ctx context.Context, owner, repo string) (*Pages, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pages/builds
func (s *RepositoriesService) ListPagesBuilds(ctx context.Context, owner, repo string, opts *ListOptions) ([]*PagesBuild, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pages/builds/latest
func (s *RepositoriesService) GetLatestPagesBuild(ctx context.Context, owner, repo string) (*PagesBuild, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pages/builds/{build_id}
func (s *RepositoriesService) GetPageBuild(ctx context.Context, owner, repo string, id int64) (*PagesBuild, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/pages/builds
func (s *RepositoriesService) RequestPageBuild(ctx context.Context, owner, repo string) (*PagesBuild, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/pages/health
func (s *RepositoriesService) GetPageHealthCheck(ctx context.Context, owner, repo string) (*PagesHealthCheckResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
