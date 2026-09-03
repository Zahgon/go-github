package github

import (
	"context"
)

type Dependency struct {
	Package      *VulnerabilityPackage `json:"package,omitempty"`
	ManifestPath *string               `json:"manifest_path,omitempty"`
	Scope        *string               `json:"scope,omitempty"`
}

type AdvisoryCVSS struct {
	Score        *float64 `json:"score,omitempty"`
	VectorString *string  `json:"vector_string,omitempty"`
}

type AdvisoryCVSSSeverities struct {
	CVSSV3 *AdvisoryCVSS `json:"cvss_v3,omitempty"`
	CVSSV4 *AdvisoryCVSS `json:"cvss_v4,omitempty"`
}

type AdvisoryCWEs struct {
	CWEID *string `json:"cwe_id,omitempty"`
	Name  *string `json:"name,omitempty"`
}

type AdvisoryEPSS struct {
	Percentage float64 `json:"percentage"`
	Percentile float64 `json:"percentile"`
}

type DependabotSecurityAdvisory struct {
	GHSAID          *string                  `json:"ghsa_id,omitempty"`
	CVEID           *string                  `json:"cve_id,omitempty"`
	Summary         *string                  `json:"summary,omitempty"`
	Description     *string                  `json:"description,omitempty"`
	Vulnerabilities []*AdvisoryVulnerability `json:"vulnerabilities,omitempty"`
	Severity        *string                  `json:"severity,omitempty"`
	Classification  *string                  `json:"classification,omitempty"`
	CVSS            *AdvisoryCVSS            `json:"cvss,omitempty"`
	CVSSSeverities  *AdvisoryCVSSSeverities  `json:"cvss_severities,omitempty"`
	CWEs            []*AdvisoryCWEs          `json:"cwes,omitempty"`
	EPSS            *AdvisoryEPSS            `json:"epss,omitempty"`
	Identifiers     []*AdvisoryIdentifier    `json:"identifiers,omitempty"`
	References      []*AdvisoryReference     `json:"references,omitempty"`
	PublishedAt     *Timestamp               `json:"published_at,omitempty"`
	UpdatedAt       *Timestamp               `json:"updated_at,omitempty"`
	WithdrawnAt     *Timestamp               `json:"withdrawn_at,omitempty"`
}

type DependabotAlert struct {
	Number                *int                        `json:"number,omitempty"`
	State                 *string                     `json:"state,omitempty"`
	Dependency            *Dependency                 `json:"dependency,omitempty"`
	SecurityAdvisory      *DependabotSecurityAdvisory `json:"security_advisory,omitempty"`
	SecurityVulnerability *AdvisoryVulnerability      `json:"security_vulnerability,omitempty"`
	URL                   *string                     `json:"url,omitempty"`
	HTMLURL               *string                     `json:"html_url,omitempty"`
	CreatedAt             *Timestamp                  `json:"created_at,omitempty"`
	UpdatedAt             *Timestamp                  `json:"updated_at,omitempty"`
	DismissedAt           *Timestamp                  `json:"dismissed_at,omitempty"`
	DismissedBy           *User                       `json:"dismissed_by,omitempty"`
	DismissedReason       *string                     `json:"dismissed_reason,omitempty"`
	DismissedComment      *string                     `json:"dismissed_comment,omitempty"`
	FixedAt               *Timestamp                  `json:"fixed_at,omitempty"`
	AutoDismissedAt       *Timestamp                  `json:"auto_dismissed_at,omitempty"`

	Repository *Repository `json:"repository,omitempty"`
}

type DependabotAlertState struct {
	State string `json:"state"`

	DismissedReason *string `json:"dismissed_reason,omitempty"`

	DismissedComment *string `json:"dismissed_comment,omitempty"`
}

type ListAlertsOptions struct {
	State     *string `url:"state,omitempty"`
	Severity  *string `url:"severity,omitempty"`
	Ecosystem *string `url:"ecosystem,omitempty"`
	Package   *string `url:"package,omitempty"`
	Scope     *string `url:"scope,omitempty"`
	Sort      *string `url:"sort,omitempty"`
	Direction *string `url:"direction,omitempty"`

	ListOptions
	ListCursorOptions
}

func (s *DependabotService) listAlerts(ctx context.Context, url string, opts *ListAlertsOptions) ([]*DependabotAlert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/dependabot/alerts
func (s *DependabotService) ListRepoAlerts(ctx context.Context, owner, repo string, opts *ListAlertsOptions) ([]*DependabotAlert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/dependabot/alerts
func (s *DependabotService) ListOrgAlerts(ctx context.Context, org string, opts *ListAlertsOptions) ([]*DependabotAlert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/dependabot/alerts/{alert_number}
func (s *DependabotService) GetRepoAlert(ctx context.Context, owner, repo string, number int) (*DependabotAlert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/dependabot/alerts/{alert_number}
func (s *DependabotService) UpdateAlert(ctx context.Context, owner, repo string, number int, body *DependabotAlertState) (*DependabotAlert, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
