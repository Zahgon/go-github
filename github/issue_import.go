package github

import (
	"context"
)

type IssueImportService service

type IssueImportRequest struct {
	IssueImport IssueImport `json:"issue"`
	Comments    []*Comment  `json:"comments,omitempty"`
}

type IssueImport struct {
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
	ClosedAt  *Timestamp `json:"closed_at,omitempty"`
	UpdatedAt *Timestamp `json:"updated_at,omitempty"`
	Assignee  *string    `json:"assignee,omitempty"`
	Milestone *int       `json:"milestone,omitempty"`
	Closed    *bool      `json:"closed,omitempty"`
	Labels    []string   `json:"labels,omitempty"`
}

type Comment struct {
	CreatedAt *Timestamp `json:"created_at,omitempty"`
	Body      string     `json:"body"`
}

type IssueImportResponse struct {
	ID               *int                `json:"id,omitempty"`
	Status           *string             `json:"status,omitempty"`
	URL              *string             `json:"url,omitempty"`
	ImportIssuesURL  *string             `json:"import_issues_url,omitempty"`
	RepositoryURL    *string             `json:"repository_url,omitempty"`
	CreatedAt        *Timestamp          `json:"created_at,omitempty"`
	UpdatedAt        *Timestamp          `json:"updated_at,omitempty"`
	Message          *string             `json:"message,omitempty"`
	DocumentationURL *string             `json:"documentation_url,omitempty"`
	Errors           []*IssueImportError `json:"errors,omitempty"`
}

type IssueImportError struct {
	Location *string `json:"location,omitempty"`
	Resource *string `json:"resource,omitempty"`
	Field    *string `json:"field,omitempty"`
	Value    *string `json:"value,omitempty"`
	Code     *string `json:"code,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/import/issues
func (s *IssueImportService) Create(ctx context.Context, owner, repo string, body *IssueImportRequest) (*IssueImportResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/import/issues/{issue_number}
func (s *IssueImportService) CheckStatus(ctx context.Context, owner, repo string, issueID int64) (*IssueImportResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/import/issues
func (s *IssueImportService) CheckStatusSince(ctx context.Context, owner, repo string, since Timestamp) ([]*IssueImportResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
