package github

import (
	"context"
)

type MetaService service

type APIMeta struct {
	Hooks []string `json:"hooks,omitempty"`

	Git []string `json:"git,omitempty"`

	VerifiablePasswordAuthentication *bool `json:"verifiable_password_authentication,omitempty"`

	Packages []string `json:"packages,omitempty"`

	Pages []string `json:"pages,omitempty"`

	Importer []string `json:"importer,omitempty"`

	GithubEnterpriseImporter []string `json:"github_enterprise_importer,omitempty"`

	Actions []string `json:"actions,omitempty"`

	ActionsMacos []string `json:"actions_macos,omitempty"`

	Codespaces []string `json:"codespaces,omitempty"`

	Copilot []string `json:"copilot,omitempty"`

	Dependabot []string `json:"dependabot,omitempty"`

	SSHKeyFingerprints map[string]string `json:"ssh_key_fingerprints,omitempty"`

	SSHKeys []string `json:"ssh_keys,omitempty"`

	Web []string `json:"web,omitempty"`

	API []string `json:"api,omitempty"`

	Domains *APIMetaDomains `json:"domains,omitempty"`
}

type APIMetaDomains struct {
	Website              []string                     `json:"website,omitempty"`
	Codespaces           []string                     `json:"codespaces,omitempty"`
	Copilot              []string                     `json:"copilot,omitempty"`
	Packages             []string                     `json:"packages,omitempty"`
	Actions              []string                     `json:"actions,omitempty"`
	ActionsInbound       *ActionsInboundDomains       `json:"actions_inbound,omitempty"`
	ArtifactAttestations *APIMetaArtifactAttestations `json:"artifact_attestations,omitempty"`
}

type ActionsInboundDomains struct {
	FullDomains     []string `json:"full_domains,omitempty"`
	WildcardDomains []string `json:"wildcard_domains,omitempty"`
}

type APIMetaArtifactAttestations struct {
	TrustDomain string   `json:"trust_domain,omitempty"`
	Services    []string `json:"services,omitempty"`
}

//meta:operation GET /meta
func (s *MetaService) Get(ctx context.Context) (*APIMeta, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Client) APIMeta(ctx context.Context) (*APIMeta, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /octocat
func (s *MetaService) Octocat(ctx context.Context, message string) (string, *Response, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (c *Client) Octocat(ctx context.Context, message string) (string, *Response, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

//meta:operation GET /zen
func (s *MetaService) Zen(ctx context.Context) (string, *Response, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (c *Client) Zen(ctx context.Context) (string, *Response, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
