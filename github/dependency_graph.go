package github

import (
	"context"
)

type DependencyGraphService service

type SBOM struct {
	SBOM *SBOMInfo `json:"sbom,omitempty"`
}

type CreationInfo struct {
	Created  *Timestamp `json:"created,omitempty"`
	Creators []string   `json:"creators,omitempty"`
}

type RepoDependencies struct {
	SPDXID *string `json:"SPDXID,omitempty"`

	Name             *string               `json:"name,omitempty"`
	VersionInfo      *string               `json:"versionInfo,omitempty"`
	DownloadLocation *string               `json:"downloadLocation,omitempty"`
	FilesAnalyzed    *bool                 `json:"filesAnalyzed,omitempty"`
	LicenseConcluded *string               `json:"licenseConcluded,omitempty"`
	LicenseDeclared  *string               `json:"licenseDeclared,omitempty"`
	ExternalRefs     []*PackageExternalRef `json:"externalRefs,omitempty"`
}

type PackageExternalRef struct {
	ReferenceCategory string `json:"referenceCategory"`

	ReferenceType string `json:"referenceType"`

	ReferenceLocator string `json:"referenceLocator"`
}

type SBOMRelationship struct {
	SPDXElementID string `json:"spdxElementId"`

	RelatedSPDXElement string `json:"relatedSpdxElement"`

	RelationshipType string `json:"relationshipType"`
}

type SBOMInfo struct {
	SPDXID       *string       `json:"SPDXID,omitempty"`
	SPDXVersion  *string       `json:"spdxVersion,omitempty"`
	CreationInfo *CreationInfo `json:"creationInfo,omitempty"`

	Name              *string  `json:"name,omitempty"`
	DataLicense       *string  `json:"dataLicense,omitempty"`
	DocumentDescribes []string `json:"documentDescribes,omitempty"`
	DocumentNamespace *string  `json:"documentNamespace,omitempty"`

	Packages []*RepoDependencies `json:"packages,omitempty"`

	Relationships []*SBOMRelationship `json:"relationships,omitempty"`
}

func (s SBOM) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /repos/{owner}/{repo}/dependency-graph/sbom
func (s *DependencyGraphService) GetSBOM(ctx context.Context, owner, repo string) (*SBOM, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
