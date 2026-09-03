package github

import (
	"context"
)

type MaintenanceOperationStatus struct {
	Hostname *string `json:"hostname,omitempty"`
	UUID     *string `json:"uuid,omitempty"`
	Message  *string `json:"message,omitempty"`
}

type MaintenanceStatus struct {
	Hostname               *string                  `json:"hostname,omitempty"`
	UUID                   *string                  `json:"uuid,omitempty"`
	Status                 *string                  `json:"status,omitempty"`
	ScheduledTime          *Timestamp               `json:"scheduled_time,omitempty"`
	ConnectionServices     []*ConnectionServiceItem `json:"connection_services,omitempty"`
	CanUnsetMaintenance    *bool                    `json:"can_unset_maintenance,omitempty"`
	IPExceptionList        []string                 `json:"ip_exception_list,omitempty"`
	MaintenanceModeMessage *string                  `json:"maintenance_mode_message,omitempty"`
}

type ConnectionServiceItem struct {
	Name   *string `json:"name,omitempty"`
	Number *int    `json:"number,omitempty"`
}

type MaintenanceOptions struct {
	Enabled                bool     `json:"enabled"`
	UUID                   *string  `json:"uuid,omitempty"`
	When                   *string  `json:"when,omitempty"`
	IPExceptionList        []string `json:"ip_exception_list,omitempty"`
	MaintenanceModeMessage *string  `json:"maintenance_mode_message,omitempty"`
}

//meta:operation GET /manage/v1/maintenance
func (s *EnterpriseService) GetMaintenanceStatus(ctx context.Context, opts *NodeQueryOptions) ([]*MaintenanceStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /manage/v1/maintenance
func (s *EnterpriseService) CreateMaintenance(ctx context.Context, enable bool, body *MaintenanceOptions) ([]*MaintenanceOperationStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
