package github

import (
	"context"
)

type CodespacesOrgAccessControlRequest struct {
	Visibility string `json:"visibility"`

	SelectedUsernames []string `json:"selected_usernames,omitzero"`
}

//meta:operation GET /orgs/{org}/codespaces
func (s *CodespacesService) ListInOrg(ctx context.Context, org string, opts *ListOptions) (*ListCodespaces, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/codespaces/access
func (s *CodespacesService) SetOrgAccessControl(ctx context.Context, org string, body CodespacesOrgAccessControlRequest) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /orgs/{org}/codespaces/access/selected_users
func (s *CodespacesService) AddUsersToOrgAccess(ctx context.Context, org string, usernames []string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/codespaces/access/selected_users
func (s *CodespacesService) RemoveUsersFromOrgAccess(ctx context.Context, org string, usernames []string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/members/{username}/codespaces
func (s *CodespacesService) ListUserCodespacesInOrg(ctx context.Context, org, username string, opts *ListOptions) (*ListCodespaces, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/members/{username}/codespaces/{codespace_name}
func (s *CodespacesService) DeleteUserCodespaceInOrg(ctx context.Context, org, username, codespaceName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /orgs/{org}/members/{username}/codespaces/{codespace_name}/stop
func (s *CodespacesService) StopUserCodespaceInOrg(ctx context.Context, org, username, codespaceName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
