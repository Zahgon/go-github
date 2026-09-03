package github

import (
	"context"
)

type CodespacesMachines struct {
	TotalCount int64                `json:"total_count"`
	Machines   []*CodespacesMachine `json:"machines"`
}

type ListRepoMachineTypesOptions struct {
	Ref *string `url:"ref,omitempty"`

	Location *string `url:"location,omitempty"`

	ClientIP *string `url:"client_ip,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/codespaces/machines
func (s *CodespacesService) ListRepositoryMachineTypes(ctx context.Context, owner, repo string, opts *ListRepoMachineTypesOptions) (*CodespacesMachines, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/codespaces/{codespace_name}/machines
func (s *CodespacesService) ListCodespaceMachineTypes(ctx context.Context, codespaceName string) (*CodespacesMachines, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
