package github

import (
	"context"
)

//meta:operation GET /enterprises/{enterprise}/audit-log
func (s *EnterpriseService) GetAuditLog(ctx context.Context, enterprise string, opts *GetAuditLogOptions) ([]*AuditEntry, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
