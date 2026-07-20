package github

import (
	"context"
)

type SecurityAdvisoriesService service

type SecurityAdvisorySubmission struct {
	Accepted *bool `json:"accepted,omitempty"`
}

type RepoAdvisoryCredit struct {
	Login *string `json:"login,omitempty"`
	Type  *string `json:"type,omitempty"`
}

type RepoAdvisoryCreditDetailed struct {
	User  *User   `json:"user,omitempty"`
	Type  *string `json:"type,omitempty"`
	State *string `json:"state,omitempty"`
}

type ListRepositorySecurityAdvisoriesOptions struct {
	ListCursorOptions

	Direction string `url:"direction,omitempty"`

	Sort string `url:"sort,omitempty"`

	State string `url:"state,omitempty"`
}

type ListGlobalSecurityAdvisoriesOptions struct {
	ListCursorOptions

	GHSAID *string `url:"ghsa_id,omitempty"`

	Type *string `url:"type,omitempty"`

	CVEID *string `url:"cve_id,omitempty"`

	Ecosystem *string `url:"ecosystem,omitempty"`

	Severity *string `url:"severity,omitempty"`

	CWEs []string `url:"cwes,omitempty"`

	IsWithdrawn *bool `url:"is_withdrawn,omitempty"`

	Affects *string `url:"affects,omitempty"`

	Published *string `url:"published,omitempty"`

	Updated *string `url:"updated,omitempty"`

	Modified *string `url:"modified,omitempty"`
}

type GlobalSecurityAdvisory struct {
	SecurityAdvisory
	ID                    *int64                         `json:"id,omitempty"`
	RepositoryAdvisoryURL *string                        `json:"repository_advisory_url,omitempty"`
	Type                  *string                        `json:"type,omitempty"`
	SourceCodeLocation    *string                        `json:"source_code_location,omitempty"`
	References            []string                       `json:"references,omitempty"`
	Vulnerabilities       []*GlobalSecurityVulnerability `json:"vulnerabilities,omitempty"`
	GithubReviewedAt      *Timestamp                     `json:"github_reviewed_at,omitempty"`
	NVDPublishedAt        *Timestamp                     `json:"nvd_published_at,omitempty"`
	Credits               []*Credit                      `json:"credits,omitempty"`
}

type GlobalSecurityVulnerability struct {
	Package                *VulnerabilityPackage `json:"package,omitempty"`
	FirstPatchedVersion    *string               `json:"first_patched_version,omitempty"`
	VulnerableVersionRange *string               `json:"vulnerable_version_range,omitempty"`
	VulnerableFunctions    []string              `json:"vulnerable_functions,omitempty"`
}

type Credit struct {
	User *User   `json:"user,omitempty"`
	Type *string `json:"type,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/security-advisories/{ghsa_id}/cve
func (s *SecurityAdvisoriesService) RequestCVE(ctx context.Context, owner, repo, ghsaID string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/security-advisories/{ghsa_id}/forks
func (s *SecurityAdvisoriesService) CreateTemporaryPrivateFork(ctx context.Context, owner, repo, ghsaID string) (*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/security-advisories
func (s *SecurityAdvisoriesService) ListRepositorySecurityAdvisoriesForOrg(ctx context.Context, org string, opts *ListRepositorySecurityAdvisoriesOptions) ([]*SecurityAdvisory, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/security-advisories
func (s *SecurityAdvisoriesService) ListRepositorySecurityAdvisories(ctx context.Context, owner, repo string, opts *ListRepositorySecurityAdvisoriesOptions) ([]*SecurityAdvisory, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /advisories
func (s *SecurityAdvisoriesService) ListGlobalSecurityAdvisories(ctx context.Context, opts *ListGlobalSecurityAdvisoriesOptions) ([]*GlobalSecurityAdvisory, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /advisories/{ghsa_id}
func (s *SecurityAdvisoriesService) GetGlobalSecurityAdvisories(ctx context.Context, ghsaID string) (*GlobalSecurityAdvisory, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
