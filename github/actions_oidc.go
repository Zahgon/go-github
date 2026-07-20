package github

import (
	"context"
)

type OIDCSubjectClaimCustomTemplate struct {
	UseDefault          *bool    `json:"use_default,omitempty"`
	IncludeClaimKeys    []string `json:"include_claim_keys,omitempty"`
	UseImmutableSubject *bool    `json:"use_immutable_subject,omitempty"`
	SubClaimPrefix      *string  `json:"sub_claim_prefix,omitempty"`
}

//meta:operation GET /orgs/{org}/actions/oidc/customization/sub
func (s *ActionsService) GetOrgOIDCSubjectClaimCustomTemplate(ctx context.Context, org string) (*OIDCSubjectClaimCustomTemplate, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/actions/oidc/customization/sub
func (s *ActionsService) GetRepoOIDCSubjectClaimCustomTemplate(ctx context.Context, owner, repo string) (*OIDCSubjectClaimCustomTemplate, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) getOIDCSubjectClaimCustomTemplate(ctx context.Context, url string) (*OIDCSubjectClaimCustomTemplate, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /orgs/{org}/actions/oidc/customization/sub
func (s *ActionsService) SetOrgOIDCSubjectClaimCustomTemplate(ctx context.Context, org string, body OIDCSubjectClaimCustomTemplate) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/actions/oidc/customization/sub
func (s *ActionsService) SetRepoOIDCSubjectClaimCustomTemplate(ctx context.Context, owner, repo string, body OIDCSubjectClaimCustomTemplate) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ActionsService) setOIDCSubjectClaimCustomTemplate(ctx context.Context, url string, body OIDCSubjectClaimCustomTemplate) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
