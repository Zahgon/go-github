package github

import (
	"context"
)

type ConfigApplyOptions struct {
	RunID *string `json:"run_id,omitempty"`
}

type ConfigApplyStatus struct {
	Running    *bool                    `json:"running,omitempty"`
	Successful *bool                    `json:"successful,omitempty"`
	Nodes      []*ConfigApplyStatusNode `json:"nodes"`
}

type ConfigApplyStatusNode struct {
	Hostname   *string `json:"hostname,omitempty"`
	Running    *bool   `json:"running,omitempty"`
	Successful *bool   `json:"successful,omitempty"`
	RunID      *string `json:"run_id,omitempty"`
}

type ConfigApplyEventsOptions struct {
	LastRequestID *string `url:"last_request_id,omitempty"`
}

type ConfigApplyEvents struct {
	Nodes []*ConfigApplyEventsNode `json:"nodes"`
}

type ConfigApplyEventsNode struct {
	Node          *string                       `json:"node,omitempty"`
	LastRequestID *string                       `json:"last_request_id,omitempty"`
	Events        []*ConfigApplyEventsNodeEvent `json:"events"`
}

type ConfigApplyEventsNodeEvent struct {
	Timestamp    *Timestamp `json:"timestamp,omitempty"`
	SeverityText *string    `json:"severity_text,omitempty"`
	Body         *string    `json:"body,omitempty"`
	EventName    *string    `json:"event_name,omitempty"`
	Topology     *string    `json:"topology,omitempty"`
	Hostname     *string    `json:"hostname,omitempty"`
	ConfigRunID  *string    `json:"config_run_id,omitempty"`
	TraceID      *string    `json:"trace_id,omitempty"`
	SpanID       *string    `json:"span_id,omitempty"`
	SpanParentID *int64     `json:"span_parent_id,omitempty"`
	SpanDepth    *int       `json:"span_depth,omitempty"`
}

type InitialConfigOptions struct {
	License  string `json:"license"`
	Password string `json:"password"`
}

type LicenseStatus struct {
	AdvancedSecurityEnabled      *bool      `json:"advancedSecurityEnabled,omitempty"`
	AdvancedSecuritySeats        *int       `json:"advancedSecuritySeats,omitempty"`
	ClusterSupport               *bool      `json:"clusterSupport,omitempty"`
	Company                      *string    `json:"company,omitempty"`
	CroquetSupport               *bool      `json:"croquetSupport,omitempty"`
	CustomTerms                  *bool      `json:"customTerms,omitempty"`
	Evaluation                   *bool      `json:"evaluation,omitempty"`
	ExpireAt                     *Timestamp `json:"expireAt,omitempty"`
	InsightsEnabled              *bool      `json:"insightsEnabled,omitempty"`
	InsightsExpireAt             *Timestamp `json:"insightsExpireAt,omitempty"`
	LearningLabEvaluationExpires *Timestamp `json:"learningLabEvaluationExpires,omitempty"`
	LearningLabSeats             *int       `json:"learningLabSeats,omitempty"`
	Perpetual                    *bool      `json:"perpetual,omitempty"`
	ReferenceNumber              *string    `json:"referenceNumber,omitempty"`
	Seats                        *int       `json:"seats,omitempty"`
	SSHAllowed                   *bool      `json:"sshAllowed,omitempty"`

	SupportKey       *bool `json:"supportKey,omitempty"`
	UnlimitedSeating *bool `json:"unlimitedSeating,omitempty"`
}

type UploadLicenseOptions struct {
	License string `url:"license"`
}

type LicenseCheck struct {
	Status *string `json:"status,omitempty"`
}

type ConfigSettings struct {
	PrivateMode           *bool                          `json:"private_mode,omitempty"`
	PublicPages           *bool                          `json:"public_pages,omitempty"`
	SubdomainIsolation    *bool                          `json:"subdomain_isolation,omitempty"`
	SignupEnabled         *bool                          `json:"signup_enabled,omitempty"`
	GithubHostname        *string                        `json:"github_hostname,omitempty"`
	IdenticonsHost        *string                        `json:"identicons_host,omitempty"`
	HTTPProxy             *string                        `json:"http_proxy,omitempty"`
	AuthMode              *string                        `json:"auth_mode,omitempty"`
	ExpireSessions        *bool                          `json:"expire_sessions,omitempty"`
	AdminPassword         *string                        `json:"admin_password,omitempty"`
	ConfigurationID       *int64                         `json:"configuration_id,omitempty"`
	ConfigurationRunCount *int                           `json:"configuration_run_count,omitempty"`
	Avatar                *ConfigSettingsAvatar          `json:"avatar,omitempty"`
	Customer              *ConfigSettingsCustomer        `json:"customer,omitempty"`
	License               *ConfigSettingsLicenseSettings `json:"license,omitempty"`
	GithubSSL             *ConfigSettingsGithubSSL       `json:"github_ssl,omitempty"`
	LDAP                  *ConfigSettingsLDAP            `json:"ldap,omitempty"`
	CAS                   *ConfigSettingsCAS             `json:"cas,omitempty"`
	SAML                  *ConfigSettingsSAML            `json:"saml,omitempty"`
	GithubOAuth           *ConfigSettingsGithubOAuth     `json:"github_oauth,omitempty"`
	SMTP                  *ConfigSettingsSMTP            `json:"smtp,omitempty"`
	NTP                   *ConfigSettingsNTP             `json:"ntp,omitempty"`
	Timezone              *string                        `json:"timezone,omitempty"`
	SNMP                  *ConfigSettingsSNMP            `json:"snmp,omitempty"`
	Syslog                *ConfigSettingsSyslog          `json:"syslog,omitempty"`
	Assets                *string                        `json:"assets,omitempty"`
	Pages                 *ConfigSettingsPagesSettings   `json:"pages,omitempty"`
	Collectd              *ConfigSettingsCollectd        `json:"collectd,omitempty"`
	Mapping               *ConfigSettingsMapping         `json:"mapping,omitempty"`
	LoadBalancer          *string                        `json:"load_balancer,omitempty"`
}

type ConfigSettingsAvatar struct {
	Enabled *bool   `json:"enabled,omitempty"`
	URI     *string `json:"uri,omitempty"`
}

type ConfigSettingsCustomer struct {
	Name          *string `json:"name,omitempty"`
	Email         *string `json:"email,omitempty"`
	UUID          *string `json:"uuid,omitempty"`
	Secret        *string `json:"secret,omitempty"`
	PublicKeyData *string `json:"public_key_data,omitempty"`
}

type ConfigSettingsLicenseSettings struct {
	Seats            *int       `json:"seats,omitempty"`
	Evaluation       *bool      `json:"evaluation,omitempty"`
	Perpetual        *bool      `json:"perpetual,omitempty"`
	UnlimitedSeating *bool      `json:"unlimited_seating,omitempty"`
	SupportKey       *string    `json:"support_key,omitempty"`
	SSHAllowed       *bool      `json:"ssh_allowed,omitempty"`
	ClusterSupport   *bool      `json:"cluster_support,omitempty"`
	ExpireAt         *Timestamp `json:"expire_at,omitempty"`
}

type ConfigSettingsGithubSSL struct {
	Enabled *bool   `json:"enabled,omitempty"`
	Cert    *string `json:"cert,omitempty"`
	Key     *string `json:"key,omitempty"`
}

type ConfigSettingsLDAP struct {
	Host                    *string                           `json:"host,omitempty"`
	Port                    *int                              `json:"port,omitempty"`
	Base                    []string                          `json:"base,omitempty"`
	UID                     *string                           `json:"uid,omitempty"`
	BindDN                  *string                           `json:"bind_dn,omitempty"`
	Password                *string                           `json:"password,omitempty"`
	Method                  *string                           `json:"method,omitempty"`
	SearchStrategy          *string                           `json:"search_strategy,omitempty"`
	UserGroups              []string                          `json:"user_groups,omitempty"`
	AdminGroup              *string                           `json:"admin_group,omitempty"`
	VirtualAttributeEnabled *bool                             `json:"virtual_attribute_enabled,omitempty"`
	RecursiveGroupSearch    *bool                             `json:"recursive_group_search,omitempty"`
	PosixSupport            *bool                             `json:"posix_support,omitempty"`
	UserSyncEmails          *bool                             `json:"user_sync_emails,omitempty"`
	UserSyncKeys            *bool                             `json:"user_sync_keys,omitempty"`
	UserSyncInterval        *int                              `json:"user_sync_interval,omitempty"`
	TeamSyncInterval        *int                              `json:"team_sync_interval,omitempty"`
	SyncEnabled             *bool                             `json:"sync_enabled,omitempty"`
	Reconciliation          *ConfigSettingsLDAPReconciliation `json:"reconciliation,omitempty"`
	Profile                 *ConfigSettingsLDAPProfile        `json:"profile,omitempty"`
}

type ConfigSettingsLDAPReconciliation struct {
	User *string `json:"user,omitempty"`
	Org  *string `json:"org,omitempty"`
}

type ConfigSettingsLDAPProfile struct {
	UID  *string `json:"uid,omitempty"`
	Name *string `json:"name,omitempty"`
	Mail *string `json:"mail,omitempty"`
	Key  *string `json:"key,omitempty"`
}

type ConfigSettingsCAS struct {
	URL *string `json:"url,omitempty"`
}

type ConfigSettingsSAML struct {
	SSOURL             *string `json:"sso_url,omitempty"`
	Certificate        *string `json:"certificate,omitempty"`
	CertificatePath    *string `json:"certificate_path,omitempty"`
	Issuer             *string `json:"issuer,omitempty"`
	IDPInitiatedSSO    *bool   `json:"idp_initiated_sso,omitempty"`
	DisableAdminDemote *bool   `json:"disable_admin_demote,omitempty"`
}

type ConfigSettingsGithubOAuth struct {
	ClientID         *string `json:"client_id,omitempty"`
	ClientSecret     *string `json:"client_secret,omitempty"`
	OrganizationName *string `json:"organization_name,omitempty"`
	OrganizationTeam *string `json:"organization_team,omitempty"`
}

type ConfigSettingsSMTP struct {
	Enabled                 *bool   `json:"enabled,omitempty"`
	Address                 *string `json:"address,omitempty"`
	Authentication          *string `json:"authentication,omitempty"`
	Port                    *string `json:"port,omitempty"`
	Domain                  *string `json:"domain,omitempty"`
	Username                *string `json:"username,omitempty"`
	UserName                *string `json:"user_name,omitempty"`
	EnableStarttlsAuto      *bool   `json:"enable_starttls_auto,omitempty"`
	Password                *string `json:"password,omitempty"`
	DiscardToNoreplyAddress *bool   `json:"discard-to-noreply-address,omitempty"`
	SupportAddress          *string `json:"support_address,omitempty"`
	SupportAddressType      *string `json:"support_address_type,omitempty"`
	NoreplyAddress          *string `json:"noreply_address,omitempty"`
}

type ConfigSettingsNTP struct {
	PrimaryServer   *string `json:"primary_server,omitempty"`
	SecondaryServer *string `json:"secondary_server,omitempty"`
}

type ConfigSettingsSNMP struct {
	Enabled   *bool   `json:"enabled,omitempty"`
	Community *string `json:"community,omitempty"`
}

type ConfigSettingsSyslog struct {
	Enabled      *bool   `json:"enabled,omitempty"`
	Server       *string `json:"server,omitempty"`
	ProtocolName *string `json:"protocol_name,omitempty"`
}

type ConfigSettingsPagesSettings struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type ConfigSettingsCollectd struct {
	Enabled    *bool   `json:"enabled,omitempty"`
	Server     *string `json:"server,omitempty"`
	Port       *int    `json:"port,omitempty"`
	Encryption *string `json:"encryption,omitempty"`
	Username   *string `json:"username,omitempty"`
	Password   *string `json:"password,omitempty"`
}

type ConfigSettingsMapping struct {
	Enabled    *bool   `json:"enabled,omitempty"`
	Tileserver *string `json:"tileserver,omitempty"`
	Basemap    *string `json:"basemap,omitempty"`
	Token      *string `json:"token,omitempty"`
}

type NodeMetadataStatus struct {
	Topology *string        `json:"topology,omitempty"`
	Nodes    []*NodeDetails `json:"nodes"`
}

type NodeDetails struct {
	Hostname     *string  `json:"hostname,omitempty"`
	UUID         *string  `json:"uuid,omitempty"`
	ClusterRoles []string `json:"cluster_roles,omitempty"`
}

//meta:operation GET /manage/v1/config/apply/events
func (s *EnterpriseService) ConfigApplyEvents(ctx context.Context, opts *ConfigApplyEventsOptions) (*ConfigApplyEvents, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /manage/v1/config/init
func (s *EnterpriseService) InitialConfig(ctx context.Context, license, password string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /manage/v1/config/license
func (s *EnterpriseService) License(ctx context.Context) (*LicenseStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /manage/v1/config/license
func (s *EnterpriseService) UploadLicense(ctx context.Context, license string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /manage/v1/config/license/check
func (s *EnterpriseService) LicenseStatus(ctx context.Context) ([]*LicenseCheck, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /manage/v1/config/nodes
func (s *EnterpriseService) NodeMetadata(ctx context.Context, opts *NodeQueryOptions) (*NodeMetadataStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /manage/v1/config/settings
func (s *EnterpriseService) Settings(ctx context.Context) (*ConfigSettings, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /manage/v1/config/settings
func (s *EnterpriseService) UpdateSettings(ctx context.Context, body *ConfigSettings) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /manage/v1/config/apply
func (s *EnterpriseService) ConfigApply(ctx context.Context, body *ConfigApplyOptions) (*ConfigApplyOptions, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /manage/v1/config/apply
func (s *EnterpriseService) ConfigApplyStatus(ctx context.Context, opts *ConfigApplyOptions) (*ConfigApplyStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
