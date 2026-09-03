package github

import (
	"context"
)

type NodeQueryOptions struct {
	UUID *string `url:"uuid,omitempty"`

	ClusterRoles *string `url:"cluster_roles,omitempty"`
}

type ClusterStatus struct {
	Status *string              `json:"status,omitempty"`
	Nodes  []*ClusterStatusNode `json:"nodes"`
}

type ClusterStatusNode struct {
	Hostname *string                         `json:"hostname,omitempty"`
	Status   *string                         `json:"status,omitempty"`
	Services []*ClusterStatusNodeServiceItem `json:"services"`
}

type ClusterStatusNodeServiceItem struct {
	Status  *string `json:"status,omitempty"`
	Name    *string `json:"name,omitempty"`
	Details *string `json:"details,omitempty"`
}

type SystemRequirements struct {
	Status *string                   `json:"status,omitempty"`
	Nodes  []*SystemRequirementsNode `json:"nodes"`
}

type SystemRequirementsNode struct {
	Hostname    *string                             `json:"hostname,omitempty"`
	Status      *string                             `json:"status,omitempty"`
	RolesStatus []*SystemRequirementsNodeRoleStatus `json:"roles_status"`
}

type SystemRequirementsNodeRoleStatus struct {
	Status *string `json:"status,omitempty"`
	Role   *string `json:"role,omitempty"`
}

type NodeReleaseVersion struct {
	Hostname *string         `json:"hostname,omitempty"`
	Version  *ReleaseVersion `json:"version"`
}

type ReleaseVersion struct {
	Version   *string `json:"version,omitempty"`
	Platform  *string `json:"platform,omitempty"`
	BuildID   *string `json:"build_id,omitempty"`
	BuildDate *string `json:"build_date,omitempty"`
}

//meta:operation GET /manage/v1/checks/system-requirements
func (s *EnterpriseService) CheckSystemRequirements(ctx context.Context) (*SystemRequirements, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /manage/v1/cluster/status
func (s *EnterpriseService) ClusterStatus(ctx context.Context) (*ClusterStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /manage/v1/replication/status
func (s *EnterpriseService) ReplicationStatus(ctx context.Context, opts *NodeQueryOptions) (*ClusterStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /manage/v1/version
func (s *EnterpriseService) GetNodeReleaseVersions(ctx context.Context, opts *NodeQueryOptions) ([]*NodeReleaseVersion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
