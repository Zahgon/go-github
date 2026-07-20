package github

import "context"

type ActivityService service

type FeedLink struct {
	HRef *string `json:"href,omitempty"`
	Type *string `json:"type,omitempty"`
}

type Feeds struct {
	TimelineURL                 *string    `json:"timeline_url,omitempty"`
	UserURL                     *string    `json:"user_url,omitempty"`
	CurrentUserPublicURL        *string    `json:"current_user_public_url,omitempty"`
	CurrentUserURL              *string    `json:"current_user_url,omitempty"`
	CurrentUserActorURL         *string    `json:"current_user_actor_url,omitempty"`
	CurrentUserOrganizationURL  *string    `json:"current_user_organization_url,omitempty"`
	CurrentUserOrganizationURLs []string   `json:"current_user_organization_urls,omitempty"`
	Links                       *FeedLinks `json:"_links,omitempty"`
}

type FeedLinks struct {
	Timeline                 *FeedLink   `json:"timeline,omitempty"`
	User                     *FeedLink   `json:"user,omitempty"`
	CurrentUserPublic        *FeedLink   `json:"current_user_public,omitempty"`
	CurrentUser              *FeedLink   `json:"current_user,omitempty"`
	CurrentUserActor         *FeedLink   `json:"current_user_actor,omitempty"`
	CurrentUserOrganization  *FeedLink   `json:"current_user_organization,omitempty"`
	CurrentUserOrganizations []*FeedLink `json:"current_user_organizations,omitempty"`
}

//meta:operation GET /feeds
func (s *ActivityService) ListFeeds(ctx context.Context) (*Feeds, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
