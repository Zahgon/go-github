package github

import (
	"encoding/json"
)

type Package struct {
	ID          *int64      `json:"id,omitempty"`
	Name        *string     `json:"name,omitempty"`
	PackageType *string     `json:"package_type,omitempty"`
	HTMLURL     *string     `json:"html_url,omitempty"`
	Visibility  *string     `json:"visibility,omitempty"`
	Owner       *User       `json:"owner,omitempty"`
	Repository  *Repository `json:"repository,omitempty"`
	CreatedAt   *Timestamp  `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp  `json:"updated_at,omitempty"`

	Namespace      *string          `json:"namespace,omitempty"`
	Description    *string          `json:"description,omitempty"`
	Ecosystem      *string          `json:"ecosystem,omitempty"`
	PackageVersion *PackageVersion  `json:"package_version,omitempty"`
	Registry       *PackageRegistry `json:"registry,omitempty"`

	URL          *string `json:"url,omitempty"`
	VersionCount *int64  `json:"version_count,omitempty"`
}

func (p Package) String() string { _ = "STUB: not implemented"; return "" }

type PackageVersion struct {
	ID             *int64          `json:"id,omitempty"`
	Name           *string         `json:"name,omitempty"`
	URL            *string         `json:"url,omitempty"`
	PackageHTMLURL *string         `json:"package_html_url,omitempty"`
	License        *string         `json:"license,omitempty"`
	Description    *string         `json:"description,omitempty"`
	CreatedAt      *Timestamp      `json:"created_at,omitempty"`
	UpdatedAt      *Timestamp      `json:"updated_at,omitempty"`
	Metadata       json.RawMessage `json:"metadata,omitempty"`

	Version             *string                        `json:"version,omitempty"`
	Summary             *string                        `json:"summary,omitempty"`
	Body                json.RawMessage                `json:"body,omitempty"`
	BodyHTML            *string                        `json:"body_html,omitempty"`
	Release             *PackageRelease                `json:"release,omitempty"`
	Manifest            *string                        `json:"manifest,omitempty"`
	HTMLURL             *string                        `json:"html_url,omitempty"`
	TagName             *string                        `json:"tag_name,omitempty"`
	TargetCommitish     *string                        `json:"target_commitish,omitempty"`
	TargetOID           *string                        `json:"target_oid,omitempty"`
	Draft               *bool                          `json:"draft,omitempty"`
	Prerelease          *bool                          `json:"prerelease,omitempty"`
	ContainerMetadata   *PackageEventContainerMetadata `json:"container_metadata,omitempty"`
	DockerMetadata      []any                          `json:"docker_metadata,omitempty"`
	NPMMetadata         *PackageNPMMetadata            `json:"npm_metadata,omitempty"`
	NugetMetadata       []*PackageNugetMetadata        `json:"nuget_metadata,omitempty"`
	RubyMetadata        map[string]any                 `json:"ruby_metadata,omitempty"`
	PackageFiles        []*PackageFile                 `json:"package_files,omitempty"`
	PackageURL          *string                        `json:"package_url,omitempty"`
	Author              *User                          `json:"author,omitempty"`
	SourceURL           *string                        `json:"source_url,omitempty"`
	InstallationCommand *string                        `json:"installation_command,omitempty"`

	DeletedAt *Timestamp `json:"deleted_at,omitempty"`
}

func (pv *PackageVersion) GetBody() (body string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (pv *PackageVersion) GetBodyAsPackageVersionBody() (body *PackageVersionBody, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (pv *PackageVersion) GetMetadata() (metadata *PackageMetadata, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (pv *PackageVersion) GetRawMetadata() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (pv PackageVersion) String() string { _ = "STUB: not implemented"; return "" }

type PackageRelease struct {
	URL             *string    `json:"url,omitempty"`
	HTMLURL         *string    `json:"html_url,omitempty"`
	ID              *int64     `json:"id,omitempty"`
	TagName         *string    `json:"tag_name,omitempty"`
	TargetCommitish *string    `json:"target_commitish,omitempty"`
	Name            *string    `json:"name,omitempty"`
	Draft           *bool      `json:"draft,omitempty"`
	Author          *User      `json:"author,omitempty"`
	Prerelease      *bool      `json:"prerelease,omitempty"`
	CreatedAt       *Timestamp `json:"created_at,omitempty"`
	PublishedAt     *Timestamp `json:"published_at,omitempty"`
}

func (r PackageRelease) String() string { _ = "STUB: not implemented"; return "" }

type PackageFile struct {
	DownloadURL *string    `json:"download_url,omitempty"`
	ID          *int64     `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	SHA256      *string    `json:"sha256,omitempty"`
	SHA1        *string    `json:"sha1,omitempty"`
	MD5         *string    `json:"md5,omitempty"`
	ContentType *string    `json:"content_type,omitempty"`
	State       *string    `json:"state,omitempty"`
	Author      *User      `json:"author,omitempty"`
	Size        *int64     `json:"size,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
}

func (pf PackageFile) String() string { _ = "STUB: not implemented"; return "" }

type PackageRegistry struct {
	AboutURL *string `json:"about_url,omitempty"`
	Name     *string `json:"name,omitempty"`
	Type     *string `json:"type,omitempty"`
	URL      *string `json:"url,omitempty"`
	Vendor   *string `json:"vendor,omitempty"`
}

func (r PackageRegistry) String() string { _ = "STUB: not implemented"; return "" }

type PackageListOptions struct {
	Visibility *string `url:"visibility,omitempty"`

	PackageType *string `url:"package_type,omitempty"`

	State *string `url:"state,omitempty"`

	ListOptions
}

type PackageMetadata struct {
	PackageType *string                   `json:"package_type,omitempty"`
	Container   *PackageContainerMetadata `json:"container,omitempty"`
}

func (r PackageMetadata) String() string { _ = "STUB: not implemented"; return "" }

type PackageContainerMetadata struct {
	Tags []string `json:"tags,omitempty"`
}

func (r PackageContainerMetadata) String() string { _ = "STUB: not implemented"; return "" }

type PackageVersionBody struct {
	Repo *Repository             `json:"repository,omitempty"`
	Info *PackageVersionBodyInfo `json:"info,omitempty"`
}

func (b PackageVersionBody) String() string { _ = "STUB: not implemented"; return "" }

type PackageVersionBodyInfo struct {
	Type       *string `json:"type,omitempty"`
	OID        *string `json:"oid,omitempty"`
	Mode       *int64  `json:"mode,omitempty"`
	Name       *string `json:"name,omitempty"`
	Path       *string `json:"path,omitempty"`
	Size       *int64  `json:"size,omitempty"`
	Collection *bool   `json:"collection,omitempty"`
}

func (bi PackageVersionBodyInfo) String() string { _ = "STUB: not implemented"; return "" }

type PackageEventContainerMetadata struct {
	Labels   map[string]any                    `json:"labels,omitempty"`
	Manifest map[string]any                    `json:"manifest,omitempty"`
	Tag      *PackageEventContainerMetadataTag `json:"tag,omitempty"`
}

func (m PackageEventContainerMetadata) String() string { _ = "STUB: not implemented"; return "" }

type PackageEventContainerMetadataTag struct {
	Name   *string `json:"name,omitempty"`
	Digest *string `json:"digest,omitempty"`
}

func (mt PackageEventContainerMetadataTag) String() string { _ = "STUB: not implemented"; return "" }

type PackageNugetMetadata struct {
	ID    json.RawMessage `json:"id,omitempty"`
	Name  *string         `json:"name,omitempty"`
	Value json.RawMessage `json:"value,omitempty"`
}

func (nm PackageNugetMetadata) String() string { _ = "STUB: not implemented"; return "" }

type PackageNPMMetadata struct {
	Name                 *string           `json:"name,omitempty"`
	Version              *string           `json:"version,omitempty"`
	NPMUser              *string           `json:"npm_user,omitempty"`
	Author               map[string]string `json:"author,omitempty"`
	Bugs                 map[string]string `json:"bugs,omitempty"`
	Dependencies         map[string]string `json:"dependencies,omitempty"`
	DevDependencies      map[string]string `json:"dev_dependencies,omitempty"`
	PeerDependencies     map[string]string `json:"peer_dependencies,omitempty"`
	OptionalDependencies map[string]string `json:"optional_dependencies,omitempty"`
	Description          *string           `json:"description,omitempty"`
	Dist                 map[string]string `json:"dist,omitempty"`
	GitHead              *string           `json:"git_head,omitempty"`
	Homepage             *string           `json:"homepage,omitempty"`
	License              *string           `json:"license,omitempty"`
	Main                 *string           `json:"main,omitempty"`
	Repository           map[string]string `json:"repository,omitempty"`
	Scripts              map[string]any    `json:"scripts,omitempty"`
	ID                   *string           `json:"id,omitempty"`
	NodeVersion          *string           `json:"node_version,omitempty"`
	NPMVersion           *string           `json:"npm_version,omitempty"`
	HasShrinkwrap        *bool             `json:"has_shrinkwrap,omitempty"`
	Maintainers          []any             `json:"maintainers,omitempty"`
	Contributors         []any             `json:"contributors,omitempty"`
	Engines              map[string]string `json:"engines,omitempty"`
	Keywords             []string          `json:"keywords,omitempty"`
	Files                []string          `json:"files,omitempty"`
	Bin                  map[string]any    `json:"bin,omitempty"`
	Man                  map[string]any    `json:"man,omitempty"`
	Directories          map[string]string `json:"directories,omitempty"`
	OS                   []string          `json:"os,omitempty"`
	CPU                  []string          `json:"cpu,omitempty"`
	Readme               *string           `json:"readme,omitempty"`
	InstallationCommand  *string           `json:"installation_command,omitempty"`
	ReleaseID            *int64            `json:"release_id,omitempty"`
	CommitOID            *string           `json:"commit_oid,omitempty"`
	PublishedViaActions  *bool             `json:"published_via_actions,omitempty"`
	DeletedByID          *int64            `json:"deleted_by_id,omitempty"`
}

func (nm PackageNPMMetadata) String() string { _ = "STUB: not implemented"; return "" }
