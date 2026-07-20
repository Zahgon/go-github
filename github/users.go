package github

import (
	"context"
)

type UsersService service

type User struct {
	Login                   *string    `json:"login,omitempty"`
	ID                      *int64     `json:"id,omitempty"`
	UserViewType            *string    `json:"user_view_type,omitempty"`
	NodeID                  *string    `json:"node_id,omitempty"`
	AvatarURL               *string    `json:"avatar_url,omitempty"`
	HTMLURL                 *string    `json:"html_url,omitempty"`
	GravatarID              *string    `json:"gravatar_id,omitempty"`
	Name                    *string    `json:"name,omitempty"`
	Company                 *string    `json:"company,omitempty"`
	Blog                    *string    `json:"blog,omitempty"`
	Location                *string    `json:"location,omitempty"`
	Email                   *string    `json:"email,omitempty"`
	NotificationEmail       *string    `json:"notification_email,omitempty"`
	Hireable                *bool      `json:"hireable,omitempty"`
	Bio                     *string    `json:"bio,omitempty"`
	TwitterUsername         *string    `json:"twitter_username,omitempty"`
	PublicRepos             *int       `json:"public_repos,omitempty"`
	PublicGists             *int       `json:"public_gists,omitempty"`
	Followers               *int       `json:"followers,omitempty"`
	Following               *int       `json:"following,omitempty"`
	CreatedAt               *Timestamp `json:"created_at,omitempty"`
	UpdatedAt               *Timestamp `json:"updated_at,omitempty"`
	SuspendedAt             *Timestamp `json:"suspended_at,omitempty"`
	Type                    *string    `json:"type,omitempty"`
	SiteAdmin               *bool      `json:"site_admin,omitempty"`
	TotalPrivateRepos       *int64     `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos       *int64     `json:"owned_private_repos,omitempty"`
	PrivateGists            *int       `json:"private_gists,omitempty"`
	DiskUsage               *int       `json:"disk_usage,omitempty"`
	Collaborators           *int       `json:"collaborators,omitempty"`
	TwoFactorAuthentication *bool      `json:"two_factor_authentication,omitempty"`
	Plan                    *Plan      `json:"plan,omitempty"`
	BusinessPlus            *bool      `json:"business_plus,omitempty"`
	LdapDn                  *string    `json:"ldap_dn,omitempty"`

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

	TextMatches []*TextMatch `json:"text_matches,omitempty"`

	Permissions *RepositoryPermissions `json:"permissions,omitempty"`

	RoleName *string `json:"role_name,omitempty"`

	Assignment *string `json:"assignment,omitempty"`

	InheritedFrom []*Team `json:"inherited_from,omitempty"`

	Role *string `json:"role,omitempty"`

	Inherited *bool `json:"inherited,omitempty"`
}

func (u User) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /user
//meta:operation GET /users/{username}
func (s *UsersService) Get(ctx context.Context, user string) (*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/{account_id}
func (s *UsersService) GetByID(ctx context.Context, id int64) (*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /user
func (s *UsersService) Edit(ctx context.Context, body *User) (*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type HovercardOptions struct {
	SubjectType string `url:"subject_type"`

	SubjectID string `url:"subject_id"`
}

type Hovercard struct {
	Contexts []*UserContext `json:"contexts,omitempty"`
}

type UserContext struct {
	Message *string `json:"message,omitempty"`
	Octicon *string `json:"octicon,omitempty"`
}

//meta:operation GET /users/{username}/hovercard
func (s *UsersService) GetHovercard(ctx context.Context, user string, opts *HovercardOptions) (*Hovercard, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type UserListOptions struct {
	Since   int64 `url:"since,omitempty"`
	PerPage int   `url:"per_page,omitempty"`
}

//meta:operation GET /users
func (s *UsersService) ListAll(ctx context.Context, opts *UserListOptions) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/repository_invitations
func (s *UsersService) ListInvitations(ctx context.Context, opts *ListOptions) ([]*RepositoryInvitation, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /user/repository_invitations/{invitation_id}
func (s *UsersService) AcceptInvitation(ctx context.Context, invitationID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /user/repository_invitations/{invitation_id}
func (s *UsersService) DeclineInvitation(ctx context.Context, invitationID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
