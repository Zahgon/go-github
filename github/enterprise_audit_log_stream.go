package github

import (
	"context"
)

type AuditLogStream struct {
	ID            int64      `json:"id"`
	StreamType    string     `json:"stream_type"`
	StreamDetails string     `json:"stream_details"`
	Enabled       bool       `json:"enabled"`
	CreatedAt     Timestamp  `json:"created_at"`
	UpdatedAt     Timestamp  `json:"updated_at"`
	PausedAt      *Timestamp `json:"paused_at,omitempty"`
}

type AuditLogStreamConfig struct {
	Enabled        bool                       `json:"enabled"`
	StreamType     string                     `json:"stream_type"`
	VendorSpecific AuditLogStreamVendorConfig `json:"vendor_specific"`
}

type AuditLogStreamVendorConfig interface {
	isAuditLogStreamVendorConfig()
}

type AuditLogStreamKey struct {
	KeyID string `json:"key_id"`
	Key   string `json:"key"`
}

type AzureBlobConfig struct {
	KeyID           string `json:"key_id"`
	EncryptedSASURL string `json:"encrypted_sas_url"`
	Container       string `json:"container"`
}

type AzureHubConfig struct {
	Name                string `json:"name"`
	EncryptedConnstring string `json:"encrypted_connstring"`
	KeyID               string `json:"key_id"`
}

type AmazonS3OIDCConfig struct {
	Bucket             string `json:"bucket"`
	Region             string `json:"region"`
	KeyID              string `json:"key_id"`
	AuthenticationType string `json:"authentication_type"`
	ArnRole            string `json:"arn_role"`
}

type AmazonS3AccessKeysConfig struct {
	Bucket               string `json:"bucket"`
	Region               string `json:"region"`
	KeyID                string `json:"key_id"`
	AuthenticationType   string `json:"authentication_type"`
	EncryptedSecretKey   string `json:"encrypted_secret_key"`
	EncryptedAccessKeyID string `json:"encrypted_access_key_id"`
}

type SplunkConfig struct {
	Domain         string `json:"domain"`
	Port           uint16 `json:"port"`
	KeyID          string `json:"key_id"`
	EncryptedToken string `json:"encrypted_token"`
	SSLVerify      bool   `json:"ssl_verify"`
}

type HecConfig struct {
	Domain         string `json:"domain"`
	Port           uint16 `json:"port"`
	KeyID          string `json:"key_id"`
	EncryptedToken string `json:"encrypted_token"`
	Path           string `json:"path"`
	SSLVerify      bool   `json:"ssl_verify"`
}

type GoogleCloudConfig struct {
	Bucket                   string `json:"bucket"`
	KeyID                    string `json:"key_id"`
	EncryptedJSONCredentials string `json:"encrypted_json_credentials"`
}

type DatadogConfig struct {
	EncryptedToken string `json:"encrypted_token"`
	Site           string `json:"site"`
	KeyID          string `json:"key_id"`
}

func (*AzureBlobConfig) isAuditLogStreamVendorConfig()          { _ = "STUB: not implemented"; return }
func (*AzureHubConfig) isAuditLogStreamVendorConfig()           { _ = "STUB: not implemented"; return }
func (*AmazonS3OIDCConfig) isAuditLogStreamVendorConfig()       { _ = "STUB: not implemented"; return }
func (*AmazonS3AccessKeysConfig) isAuditLogStreamVendorConfig() { _ = "STUB: not implemented"; return }
func (*SplunkConfig) isAuditLogStreamVendorConfig()             { _ = "STUB: not implemented"; return }
func (*HecConfig) isAuditLogStreamVendorConfig()                { _ = "STUB: not implemented"; return }
func (*GoogleCloudConfig) isAuditLogStreamVendorConfig()        { _ = "STUB: not implemented"; return }
func (*DatadogConfig) isAuditLogStreamVendorConfig()            { _ = "STUB: not implemented"; return }

func NewAzureBlobStreamConfig(enabled bool, cfg *AzureBlobConfig) *AuditLogStreamConfig {
	_ = "STUB: not implemented"
	return nil
}

func NewAzureHubStreamConfig(enabled bool, cfg *AzureHubConfig) *AuditLogStreamConfig {
	_ = "STUB: not implemented"
	return nil
}

func NewAmazonS3OIDCStreamConfig(enabled bool, cfg *AmazonS3OIDCConfig) *AuditLogStreamConfig {
	_ = "STUB: not implemented"
	return nil
}

func NewAmazonS3AccessKeysStreamConfig(enabled bool, cfg *AmazonS3AccessKeysConfig) *AuditLogStreamConfig {
	_ = "STUB: not implemented"
	return nil
}

func NewSplunkStreamConfig(enabled bool, cfg *SplunkConfig) *AuditLogStreamConfig {
	_ = "STUB: not implemented"
	return nil
}

func NewHecStreamConfig(enabled bool, cfg *HecConfig) *AuditLogStreamConfig {
	_ = "STUB: not implemented"
	return nil
}

func NewGoogleCloudStreamConfig(enabled bool, cfg *GoogleCloudConfig) *AuditLogStreamConfig {
	_ = "STUB: not implemented"
	return nil
}

func NewDatadogStreamConfig(enabled bool, cfg *DatadogConfig) *AuditLogStreamConfig {
	_ = "STUB: not implemented"
	return nil
}

//meta:operation GET /enterprises/{enterprise}/audit-log/stream-key
func (s *EnterpriseService) GetAuditLogStreamKey(ctx context.Context, enterprise string) (*AuditLogStreamKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/audit-log/streams
func (s *EnterpriseService) ListAuditLogStreams(ctx context.Context, enterprise string) ([]*AuditLogStream, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/audit-log/streams/{stream_id}
func (s *EnterpriseService) GetAuditLogStream(ctx context.Context, enterprise string, streamID int64) (*AuditLogStream, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/audit-log/streams
func (s *EnterpriseService) CreateAuditLogStream(ctx context.Context, enterprise string, body AuditLogStreamConfig) (*AuditLogStream, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /enterprises/{enterprise}/audit-log/streams/{stream_id}
func (s *EnterpriseService) UpdateAuditLogStream(ctx context.Context, enterprise string, streamID int64, body AuditLogStreamConfig) (*AuditLogStream, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/audit-log/streams/{stream_id}
func (s *EnterpriseService) DeleteAuditLogStream(ctx context.Context, enterprise string, streamID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
