package scrape

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

var defaultBaseURL = "https://github.com/"

type Client struct {
	*http.Client

	baseURL *url.URL
}

func NewClient(transport http.RoundTripper) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) SaveCookies() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Client) LoadCookies(v []byte) error { _ = "STUB: not implemented"; return nil }

func (c *Client) get(urlStr string, a ...any) (*goquery.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Authenticate(username, password, otpseed string) error {
	_ = "STUB: not implemented"
	return nil
}
