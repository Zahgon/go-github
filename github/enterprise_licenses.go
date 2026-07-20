package github

import (
	"context"
)

type EnterpriseConsumedLicenses struct {
	TotalSeatsConsumed  int                        `json:"total_seats_consumed"`
	TotalSeatsPurchased int                        `json:"total_seats_purchased"`
	Users               []*EnterpriseLicensedUsers `json:"users,omitempty"`
}

type EnterpriseLicensedUsers struct {
	GithubComLogin                  string   `json:"github_com_login"`
	GithubComName                   *string  `json:"github_com_name"`
	EnterpriseServerUserIDs         []string `json:"enterprise_server_user_ids,omitempty"`
	GithubComUser                   bool     `json:"github_com_user"`
	EnterpriseServerUser            *bool    `json:"enterprise_server_user"`
	VisualStudioSubscriptionUser    bool     `json:"visual_studio_subscription_user"`
	LicenseType                     string   `json:"license_type"`
	GithubComProfile                *string  `json:"github_com_profile"`
	GithubComMemberRoles            []string `json:"github_com_member_roles,omitempty"`
	GithubComEnterpriseRoles        []string `json:"github_com_enterprise_roles,omitempty"`
	GithubComVerifiedDomainEmails   []string `json:"github_com_verified_domain_emails,omitempty"`
	GithubComSamlNameID             *string  `json:"github_com_saml_name_id"`
	GithubComOrgsWithPendingInvites []string `json:"github_com_orgs_with_pending_invites,omitempty"`
	GithubComTwoFactorAuth          *bool    `json:"github_com_two_factor_auth"`
	EnterpriseServerEmails          []string `json:"enterprise_server_emails,omitempty"`
	VisualStudioLicenseStatus       *string  `json:"visual_studio_license_status"`
	VisualStudioSubscriptionEmail   *string  `json:"visual_studio_subscription_email"`
	TotalUserAccounts               int      `json:"total_user_accounts"`
}

type EnterpriseLicenseSyncStatus struct {
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	Properties  *ServerInstanceProperties `json:"properties,omitempty"`
}

type ServerInstanceProperties struct {
	ServerInstances *ServerInstances `json:"server_instances,omitempty"`
}

type ServerInstances struct {
	Type  string                `json:"type"`
	Items *ServiceInstanceItems `json:"items,omitempty"`
}

type ServiceInstanceItems struct {
	Type       string                `json:"type"`
	Properties *ServerItemProperties `json:"properties,omitempty"`
}

type ServerItemProperties struct {
	ServerID string           `json:"server_id"`
	Hostname string           `json:"hostname"`
	LastSync *LastLicenseSync `json:"last_sync,omitempty"`
}

type LastLicenseSync struct {
	Type       string                     `json:"type"`
	Properties *LastLicenseSyncProperties `json:"properties,omitempty"`
}

type LastLicenseSyncProperties struct {
	Date   *Timestamp `json:"date,omitempty"`
	Status string     `json:"status"`
	Error  string     `json:"error"`
}

//meta:operation GET /enterprises/{enterprise}/consumed-licenses
func (s *EnterpriseService) ListConsumedLicenses(ctx context.Context, enterprise string, opts *ListOptions) (*EnterpriseConsumedLicenses, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/license-sync-status
func (s *EnterpriseService) GetLicenseSyncStatus(ctx context.Context, enterprise string) (*EnterpriseLicenseSyncStatus, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
