package github

import (
	"context"
)

type createOrgRequest struct {
	Login *string `json:"login,omitempty"`
	Admin *string `json:"admin,omitempty"`
}

//meta:operation POST /admin/organizations
func (s *AdminService) CreateOrg(ctx context.Context, org *Organization, admin string) (*Organization, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type renameOrgRequest struct {
	Login *string `json:"login,omitempty"`
}

type RenameOrgResponse struct {
	Message *string `json:"message,omitempty"`
	URL     *string `json:"url,omitempty"`
}

//meta:operation PATCH /admin/organizations/{org}
func (s *AdminService) RenameOrg(ctx context.Context, org *Organization, newName string) (*RenameOrgResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /admin/organizations/{org}
func (s *AdminService) RenameOrgByName(ctx context.Context, org, newName string) (*RenameOrgResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
