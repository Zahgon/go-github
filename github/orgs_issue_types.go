package github

import (
	"context"
)

type CreateOrUpdateIssueTypesOptions struct {
	Name        string  `json:"name"`
	IsEnabled   bool    `json:"is_enabled"`
	IsPrivate   *bool   `json:"is_private,omitempty"`
	Description *string `json:"description,omitempty"`
	Color       *string `json:"color,omitempty"`
}

//meta:operation GET /orgs/{org}/issue-types
func (s *OrganizationsService) ListIssueTypes(ctx context.Context, org string) ([]*IssueType, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/issue-types
func (s *OrganizationsService) CreateIssueType(ctx context.Context, org string, body *CreateOrUpdateIssueTypesOptions) (*IssueType, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/issue-types/{issue_type_id}
func (s *OrganizationsService) UpdateIssueType(ctx context.Context, org string, issueTypeID int64, body *CreateOrUpdateIssueTypesOptions) (*IssueType, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/issue-types/{issue_type_id}
func (s *OrganizationsService) DeleteIssueType(ctx context.Context, org string, issueTypeID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
