package github

import (
	"context"
)

type AdminService service

type TeamLDAPMapping struct {
	ID          *int64  `json:"id,omitempty"`
	LDAPDN      *string `json:"ldap_dn,omitempty"`
	URL         *string `json:"url,omitempty"`
	Name        *string `json:"name,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	Description *string `json:"description,omitempty"`
	Privacy     *string `json:"privacy,omitempty"`
	Permission  *string `json:"permission,omitempty"`

	MembersURL      *string `json:"members_url,omitempty"`
	RepositoriesURL *string `json:"repositories_url,omitempty"`
}

func (m TeamLDAPMapping) String() string { _ = "STUB: not implemented"; return "" }

type UserLDAPMapping struct {
	ID         *int64  `json:"id,omitempty"`
	LDAPDN     *string `json:"ldap_dn,omitempty"`
	Login      *string `json:"login,omitempty"`
	AvatarURL  *string `json:"avatar_url,omitempty"`
	GravatarID *string `json:"gravatar_id,omitempty"`
	Type       *string `json:"type,omitempty"`
	SiteAdmin  *bool   `json:"site_admin,omitempty"`

	URL               *string `json:"url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`
}

func (m UserLDAPMapping) String() string { _ = "STUB: not implemented"; return "" }

type Enterprise struct {
	ID          *int       `json:"id,omitempty"`
	Slug        *string    `json:"slug,omitempty"`
	Name        *string    `json:"name,omitempty"`
	NodeID      *string    `json:"node_id,omitempty"`
	AvatarURL   *string    `json:"avatar_url,omitempty"`
	Description *string    `json:"description,omitempty"`
	WebsiteURL  *string    `json:"website_url,omitempty"`
	HTMLURL     *string    `json:"html_url,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
}

func (m Enterprise) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation PATCH /admin/ldap/users/{username}/mapping
func (s *AdminService) UpdateUserLDAPMapping(ctx context.Context, user string, body *UserLDAPMapping) (*UserLDAPMapping, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /admin/ldap/teams/{team_id}/mapping
func (s *AdminService) UpdateTeamLDAPMapping(ctx context.Context, team int64, body *TeamLDAPMapping) (*TeamLDAPMapping, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
