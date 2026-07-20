package scrape

import (
	"net/http"

	"github.com/google/go-github/v89/github"
)

func (c *Client) AppRestrictionsEnabled(org string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Client) ListOAuthApps(org string) ([]*OAuthApp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func intFromLastPathSegment(s string) int { _ = "STUB: not implemented"; return 0 }

type OAuthAppReviewState int

const (
	OAuthAppRequested OAuthAppReviewState = iota + 1

	OAuthAppApproved

	OAuthAppDenied
)

type OAuthApp struct {
	ID          int
	Name        string
	Description string
	State       OAuthAppReviewState
	RequestedBy string
}

type AppManifest struct {
	Name *string `json:"name,omitempty"`

	URL *string `json:"url,omitempty"`

	CallbackURLs []string `json:"callback_urls,omitempty"`

	HookAttributes map[string]string `json:"hook_attributes,omitempty"`

	RedirectURL *string `json:"redirect_url,omitempty"`

	Description *string `json:"description,omitempty"`

	Public *bool `json:"public,omitempty"`

	DefaultEvents []string `json:"default_events,omitempty"`

	DefaultPermissions *github.InstallationPermissions `json:"default_permissions,omitempty"`
}

func (c *Client) CreateApp(m *AppManifest, orgName string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
