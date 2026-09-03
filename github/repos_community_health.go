package github

import (
	"context"
)

type Metric struct {
	Name    *string `json:"name"`
	Key     *string `json:"key"`
	SPDXID  *string `json:"spdx_id"`
	URL     *string `json:"url"`
	HTMLURL *string `json:"html_url"`
	NodeID  *string `json:"node_id"`
}

type CommunityHealthFiles struct {
	CodeOfConduct       *Metric `json:"code_of_conduct"`
	CodeOfConductFile   *Metric `json:"code_of_conduct_file"`
	Contributing        *Metric `json:"contributing"`
	IssueTemplate       *Metric `json:"issue_template"`
	PullRequestTemplate *Metric `json:"pull_request_template"`
	License             *Metric `json:"license"`
	Readme              *Metric `json:"readme"`
}

type CommunityHealthMetrics struct {
	HealthPercentage      *int                  `json:"health_percentage"`
	Description           *string               `json:"description"`
	Documentation         *string               `json:"documentation"`
	Files                 *CommunityHealthFiles `json:"files"`
	UpdatedAt             *Timestamp            `json:"updated_at"`
	ContentReportsEnabled *bool                 `json:"content_reports_enabled"`
}

//meta:operation GET /repos/{owner}/{repo}/community/profile
func (s *RepositoriesService) GetCommunityHealthMetrics(ctx context.Context, owner, repo string) (*CommunityHealthMetrics, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
