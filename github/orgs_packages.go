package github

import (
	"context"
)

//meta:operation GET /orgs/{org}/packages
func (s *OrganizationsService) ListPackages(ctx context.Context, org string, opts *PackageListOptions) ([]*Package, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/packages/{package_type}/{package_name}
func (s *OrganizationsService) GetPackage(ctx context.Context, org, packageType, packageName string) (*Package, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/packages/{package_type}/{package_name}
func (s *OrganizationsService) DeletePackage(ctx context.Context, org, packageType, packageName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /orgs/{org}/packages/{package_type}/{package_name}/restore
func (s *OrganizationsService) RestorePackage(ctx context.Context, org, packageType, packageName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /orgs/{org}/packages/{package_type}/{package_name}/versions
func (s *OrganizationsService) PackageGetAllVersions(ctx context.Context, org, packageType, packageName string, opts *PackageListOptions) ([]*PackageVersion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/packages/{package_type}/{package_name}/versions/{package_version_id}
func (s *OrganizationsService) PackageGetVersion(ctx context.Context, org, packageType, packageName string, packageVersionID int64) (*PackageVersion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/packages/{package_type}/{package_name}/versions/{package_version_id}
func (s *OrganizationsService) PackageDeleteVersion(ctx context.Context, org, packageType, packageName string, packageVersionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /orgs/{org}/packages/{package_type}/{package_name}/versions/{package_version_id}/restore
func (s *OrganizationsService) PackageRestoreVersion(ctx context.Context, org, packageType, packageName string, packageVersionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
