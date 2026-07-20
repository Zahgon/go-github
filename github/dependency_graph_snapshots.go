package github

import (
	"context"
)

type DependencyGraphSnapshotResolvedDependency struct {
	PackageURL *string `json:"package_url,omitempty"`

	Metadata map[string]any `json:"metadata,omitempty"`

	Relationship *string `json:"relationship,omitempty"`

	Scope        *string  `json:"scope,omitempty"`
	Dependencies []string `json:"dependencies,omitempty"`
}

type DependencyGraphSnapshotJob struct {
	Correlator *string `json:"correlator,omitempty"`
	ID         *string `json:"id,omitempty"`
	HTMLURL    *string `json:"html_url,omitempty"`
}

type DependencyGraphSnapshotDetector struct {
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
	URL     *string `json:"url,omitempty"`
}

type DependencyGraphSnapshotManifestFile struct {
	SourceLocation *string `json:"source_location,omitempty"`
}

type DependencyGraphSnapshotManifest struct {
	Name     *string                                               `json:"name,omitempty"`
	File     *DependencyGraphSnapshotManifestFile                  `json:"file,omitempty"`
	Metadata map[string]any                                        `json:"metadata,omitempty"`
	Resolved map[string]*DependencyGraphSnapshotResolvedDependency `json:"resolved,omitempty"`
}

type DependencyGraphSnapshot struct {
	Version   int                                         `json:"version"`
	Sha       *string                                     `json:"sha,omitempty"`
	Ref       *string                                     `json:"ref,omitempty"`
	Job       *DependencyGraphSnapshotJob                 `json:"job,omitempty"`
	Detector  *DependencyGraphSnapshotDetector            `json:"detector,omitempty"`
	Scanned   *Timestamp                                  `json:"scanned,omitempty"`
	Metadata  map[string]any                              `json:"metadata,omitempty"`
	Manifests map[string]*DependencyGraphSnapshotManifest `json:"manifests,omitempty"`
}

type DependencyGraphSnapshotCreationData struct {
	ID        int64      `json:"id"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
	Message   *string    `json:"message,omitempty"`

	Result *string `json:"result,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/dependency-graph/snapshots
func (s *DependencyGraphService) CreateSnapshot(ctx context.Context, owner, repo string, body *DependencyGraphSnapshot) (*DependencyGraphSnapshotCreationData, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
