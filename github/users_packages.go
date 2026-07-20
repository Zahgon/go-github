package github

import (
	"context"
)

//meta:operation GET /user/packages
//meta:operation GET /users/{username}/packages
func (s *UsersService) ListPackages(ctx context.Context, user string, opts *PackageListOptions) ([]*Package, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/packages/{package_type}/{package_name}
//meta:operation GET /users/{username}/packages/{package_type}/{package_name}
func (s *UsersService) GetPackage(ctx context.Context, user, packageType, packageName string) (*Package, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /user/packages/{package_type}/{package_name}
//meta:operation DELETE /users/{username}/packages/{package_type}/{package_name}
func (s *UsersService) DeletePackage(ctx context.Context, user, packageType, packageName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /user/packages/{package_type}/{package_name}/restore
//meta:operation POST /users/{username}/packages/{package_type}/{package_name}/restore
func (s *UsersService) RestorePackage(ctx context.Context, user, packageType, packageName string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ListPackageVersionsOptions struct {
	State string `url:"state,omitempty"`

	ListOptions
}

//meta:operation GET /user/packages/{package_type}/{package_name}/versions
func (s *UsersService) ListPackageVersions(ctx context.Context, packageType, packageName string, opts *ListPackageVersionsOptions) ([]*PackageVersion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/packages/{package_type}/{package_name}/versions
func (s *UsersService) ListUserPackageVersions(ctx context.Context, user, packageType, packageName string) ([]*PackageVersion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /user/packages/{package_type}/{package_name}/versions/{package_version_id}
//meta:operation GET /users/{username}/packages/{package_type}/{package_name}/versions/{package_version_id}
func (s *UsersService) PackageGetVersion(ctx context.Context, user, packageType, packageName string, packageVersionID int64) (*PackageVersion, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /user/packages/{package_type}/{package_name}/versions/{package_version_id}
//meta:operation DELETE /users/{username}/packages/{package_type}/{package_name}/versions/{package_version_id}
func (s *UsersService) PackageDeleteVersion(ctx context.Context, user, packageType, packageName string, packageVersionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation POST /user/packages/{package_type}/{package_name}/versions/{package_version_id}/restore
//meta:operation POST /users/{username}/packages/{package_type}/{package_name}/versions/{package_version_id}/restore
func (s *UsersService) PackageRestoreVersion(ctx context.Context, user, packageType, packageName string, packageVersionID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
