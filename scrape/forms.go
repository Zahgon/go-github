package scrape

import (
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

type htmlForm struct {
	Action string

	Method string

	Values url.Values
}

func parseForms(node *html.Node) (forms []*htmlForm) { _ = "STUB: not implemented"; return nil }

func fetchAndSubmitForm(client *http.Client, urlStr string, setValues func(url.Values)) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
