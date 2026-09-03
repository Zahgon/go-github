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

type InclusionSource string

const (
	InclusionSourceEnterprise   InclusionSource = "enterprise"
	InclusionSourceOrganization InclusionSource = "organization"
)

type OIDCCustomPropertyClaim struct {
	CustomPropertyName string `json:"custom_property_name"`
}

type OIDCCustomPropertyClaimResponse struct {
	OIDCCustomPropertyClaim
	InclusionSource InclusionSource `json:"inclusion_source"`
}

//meta:operation GET /enterprises/{enterprise}/actions/oidc/customization/properties/repo
func (s *ActionsService) ListEnterpriseOIDCCustomPropertyClaims(ctx context.Context, enterprise string) ([]*OIDCCustomPropertyClaimResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/actions/oidc/customization/properties/repo
func (s *ActionsService) ListOrgOIDCCustomPropertyClaims(ctx context.Context, org string) ([]*OIDCCustomPropertyClaimResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) listOIDCCustomPropertyClaims(ctx context.Context, url string) ([]*OIDCCustomPropertyClaimResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /enterprises/{enterprise}/actions/oidc/customization/properties/repo
func (s *ActionsService) SetEnterpriseOIDCCustomPropertyClaim(ctx context.Context, enterprise string, body OIDCCustomPropertyClaim) (*OIDCCustomPropertyClaim, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/actions/oidc/customization/properties/repo
func (s *ActionsService) SetOrgOIDCCustomPropertyClaim(ctx context.Context, org string, body OIDCCustomPropertyClaim) (*OIDCCustomPropertyClaim, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ActionsService) setOIDCCustomPropertyClaim(ctx context.Context, url string, body OIDCCustomPropertyClaim) (*OIDCCustomPropertyClaim, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /enterprises/{enterprise}/actions/oidc/customization/properties/repo/{custom_property_name}
func (s *ActionsService) DeleteEnterpriseOIDCCustomPropertyClaim(ctx context.Context, enterprise, customProperty string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/actions/oidc/customization/properties/repo/{custom_property_name}
func (s *ActionsService) DeleteOrgOIDCCustomPropertyClaim(ctx context.Context, enterprise, customProperty string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ActionsService) deleteOIDCCustomPropertyClaim(ctx context.Context, url string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
