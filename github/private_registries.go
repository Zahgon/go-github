package github

import (
	"context"
)

type PrivateRegistriesService service

type PrivateRegistryType string

const (
	PrivateRegistryTypeMavenRepository    PrivateRegistryType = "maven_repository"
	PrivateRegistryTypeNugetFeed          PrivateRegistryType = "nuget_feed"
	PrivateRegistryTypeGoProxyServer      PrivateRegistryType = "goproxy_server"
	PrivateRegistryTypeNpmRegistry        PrivateRegistryType = "npm_registry"
	PrivateRegistryTypeRubygemsServer     PrivateRegistryType = "rubygems_server"
	PrivateRegistryTypeCargoRegistry      PrivateRegistryType = "cargo_registry"
	PrivateRegistryTypeComposerRepository PrivateRegistryType = "composer_repository"
	PrivateRegistryTypeDockerRegistry     PrivateRegistryType = "docker_registry"
	PrivateRegistryTypeGitSource          PrivateRegistryType = "git_source"
	PrivateRegistryTypeHelmRegistry       PrivateRegistryType = "helm_registry"
	PrivateRegistryTypeHexOrganization    PrivateRegistryType = "hex_organization"
	PrivateRegistryTypeHexRepository      PrivateRegistryType = "hex_repository"
	PrivateRegistryTypePubRepository      PrivateRegistryType = "pub_repository"
	PrivateRegistryTypePythonIndex        PrivateRegistryType = "python_index"
	PrivateRegistryTypeTerraformRegistry  PrivateRegistryType = "terraform_registry"
)

type PrivateRegistryVisibility string

const (
	PrivateRegistryVisibilityPrivate  PrivateRegistryVisibility = "private"
	PrivateRegistryVisibilityAll      PrivateRegistryVisibility = "all"
	PrivateRegistryVisibilitySelected PrivateRegistryVisibility = "selected"
)

type PrivateRegistryAuthType string

const (
	PrivateRegistryAuthTypeToken            PrivateRegistryAuthType = "token"
	PrivateRegistryAuthTypeUsernamePassword PrivateRegistryAuthType = "username_password"
	PrivateRegistryAuthTypeOIDCAzure        PrivateRegistryAuthType = "oidc_azure"
	PrivateRegistryAuthTypeOIDCAWS          PrivateRegistryAuthType = "oidc_aws"
	PrivateRegistryAuthTypeOIDCJFrog        PrivateRegistryAuthType = "oidc_jfrog"
)

type PrivateRegistry struct {
	Name *string `json:"name,omitempty"`

	RegistryType *PrivateRegistryType `json:"registry_type,omitempty"`

	AuthType *PrivateRegistryAuthType `json:"auth_type,omitempty"`

	URL *string `json:"url,omitempty"`

	Username *string `json:"username,omitempty"`

	ReplacesBase *bool `json:"replaces_base,omitempty"`

	Visibility *PrivateRegistryVisibility `json:"visibility,omitempty"`

	SelectedRepositoryIDs []int64 `json:"selected_repository_ids,omitempty"`

	TenantID *string `json:"tenant_id,omitempty"`

	ClientID *string `json:"client_id,omitempty"`

	AWSRegion *string `json:"aws_region,omitempty"`

	AccountID *string `json:"account_id,omitempty"`

	RoleName *string `json:"role_name,omitempty"`

	Domain *string `json:"domain,omitempty"`

	DomainOwner *string `json:"domain_owner,omitempty"`

	JFrogOIDCProviderName *string `json:"jfrog_oidc_provider_name,omitempty"`

	Audience *string `json:"audience,omitempty"`

	IdentityMappingName *string `json:"identity_mapping_name,omitempty"`

	CreatedAt *Timestamp `json:"created_at,omitempty"`

	UpdatedAt *Timestamp `json:"updated_at,omitempty"`
}

type PrivateRegistries struct {
	TotalCount *int `json:"total_count,omitempty"`

	Configurations []*PrivateRegistry `json:"configurations,omitempty"`
}

type CreateOrganizationPrivateRegistry struct {
	RegistryType PrivateRegistryType `json:"registry_type"`

	URL string `json:"url"`

	Username *string `json:"username,omitempty"`

	ReplacesBase *bool `json:"replaces_base,omitempty"`

	EncryptedValue *string `json:"encrypted_value,omitempty"`

	KeyID *string `json:"key_id,omitempty"`

	Visibility PrivateRegistryVisibility `json:"visibility"`

	SelectedRepositoryIDs []int64 `json:"selected_repository_ids,omitempty"`

	AuthType *string `json:"auth_type,omitempty"`

	TenantID *string `json:"tenant_id,omitempty"`

	ClientID *string `json:"client_id,omitempty"`

	AWSRegion *string `json:"aws_region,omitempty"`

	AccountID *string `json:"account_id,omitempty"`

	RoleName *string `json:"role_name,omitempty"`

	Domain *string `json:"domain,omitempty"`

	DomainOwner *string `json:"domain_owner,omitempty"`

	JFrogOIDCProviderName *string `json:"jfrog_oidc_provider_name,omitempty"`

	Audience *string `json:"audience,omitempty"`

	IdentityMappingName *string `json:"identity_mapping_name,omitempty"`
}

type UpdateOrganizationPrivateRegistry struct {
	RegistryType *PrivateRegistryType `json:"registry_type,omitempty"`

	URL *string `json:"url,omitempty"`

	Username *string `json:"username,omitempty"`

	ReplacesBase *bool `json:"replaces_base,omitempty"`

	EncryptedValue *string `json:"encrypted_value,omitempty"`

	KeyID *string `json:"key_id,omitempty"`

	Visibility *PrivateRegistryVisibility `json:"visibility,omitempty"`

	SelectedRepositoryIDs []int64 `json:"selected_repository_ids,omitempty"`

	AuthType *string `json:"auth_type,omitempty"`

	TenantID *string `json:"tenant_id,omitempty"`

	ClientID *string `json:"client_id,omitempty"`

	AWSRegion *string `json:"aws_region,omitempty"`

	AccountID *string `json:"account_id,omitempty"`

	RoleName *string `json:"role_name,omitempty"`

	Domain *string `json:"domain,omitempty"`

	DomainOwner *string `json:"domain_owner,omitempty"`

	JFrogOIDCProviderName *string `json:"jfrog_oidc_provider_name,omitempty"`

	Audience *string `json:"audience,omitempty"`

	IdentityMappingName *string `json:"identity_mapping_name,omitempty"`
}

//meta:operation GET /orgs/{org}/private-registries
func (s *PrivateRegistriesService) ListOrganizationPrivateRegistries(ctx context.Context, org string, opts *ListOptions) (*PrivateRegistries, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/private-registries
func (s *PrivateRegistriesService) CreateOrganizationPrivateRegistry(ctx context.Context, org string, body CreateOrganizationPrivateRegistry) (*PrivateRegistry, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/private-registries/public-key
func (s *PrivateRegistriesService) GetOrganizationPrivateRegistriesPublicKey(ctx context.Context, org string) (*PublicKey, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/private-registries/{secret_name}
func (s *PrivateRegistriesService) GetOrganizationPrivateRegistry(ctx context.Context, org, secretName string) (*PrivateRegistry, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/private-registries/{secret_name}
func (s *PrivateRegistriesService) UpdateOrganizationPrivateRegistry(ctx context.Context, org, secretName string, body UpdateOrganizationPrivateRegistry) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /orgs/{org}/private-registries/{secret_name}
func (s *PrivateRegistriesService) DeleteOrganizationPrivateRegistry(ctx context.Context, org, secretName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
