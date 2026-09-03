package github

import (
	"context"
)

type SSHKeyStatus struct {
	Hostname *string `json:"hostname,omitempty"`
	UUID     *string `json:"uuid,omitempty"`
	Message  *string `json:"message,omitempty"`
	Modified *bool   `json:"modified,omitempty"`
}

type SSHKeyOptions struct {
	Key string `json:"key"`
}

type ClusterSSHKey struct {
	Key         *string `json:"key,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty"`
}

//meta:operation DELETE /manage/v1/access/ssh
func (s *EnterpriseService) DeleteSSHKey(ctx context.Context, key string) ([]*SSHKeyStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /manage/v1/access/ssh
func (s *EnterpriseService) GetSSHKey(ctx context.Context) ([]*ClusterSSHKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /manage/v1/access/ssh
func (s *EnterpriseService) CreateSSHKey(ctx context.Context, key string) ([]*SSHKeyStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
