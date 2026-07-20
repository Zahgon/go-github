package github

import (
	"encoding/json"
	"time"
)

func (a *AbuseRateLimitError) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (a *AbuseRateLimitError) GetRetryAfter() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (a *AcceptedAssignment) GetAssignment() *ClassroomAssignment {
	_ = "STUB: not implemented"
	return nil
}

func (a *AcceptedAssignment) GetCommitCount() int { _ = "STUB: not implemented"; return 0 }

func (a *AcceptedAssignment) GetGrade() string { _ = "STUB: not implemented"; return "" }

func (a *AcceptedAssignment) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AcceptedAssignment) GetPassing() bool { _ = "STUB: not implemented"; return false }

func (a *AcceptedAssignment) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (a *AcceptedAssignment) GetStudents() []*ClassroomUser { _ = "STUB: not implemented"; return nil }

func (a *AcceptedAssignment) GetSubmitted() bool { _ = "STUB: not implemented"; return false }

func (a *AcceptedError) GetRaw() []byte { _ = "STUB: not implemented"; return nil }

func (a *AccessibleRepository) GetFullName() string { _ = "STUB: not implemented"; return "" }

func (a *AccessibleRepository) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AccessibleRepository) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsAllowed) GetGithubOwnedAllowed() bool { _ = "STUB: not implemented"; return false }

func (a *ActionsAllowed) GetPatternsAllowed() []string { _ = "STUB: not implemented"; return nil }

func (a *ActionsAllowed) GetVerifiedAllowed() bool { _ = "STUB: not implemented"; return false }

func (a *ActionsCache) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *ActionsCache) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *ActionsCache) GetKey() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsCache) GetLastAccessedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (a *ActionsCache) GetRef() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsCache) GetSizeInBytes() int64 { _ = "STUB: not implemented"; return 0 }

func (a *ActionsCache) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsCacheList) GetActionsCaches() []*ActionsCache {
	_ = "STUB: not implemented"
	return nil
}

func (a *ActionsCacheList) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (a *ActionsCacheListOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsCacheListOptions) GetKey() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsCacheListOptions) GetRef() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsCacheListOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsCacheUsage) GetActiveCachesCount() int { _ = "STUB: not implemented"; return 0 }

func (a *ActionsCacheUsage) GetActiveCachesSizeInBytes() int64 { _ = "STUB: not implemented"; return 0 }

func (a *ActionsCacheUsage) GetFullName() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsCacheUsageList) GetRepoCacheUsage() []*ActionsCacheUsage {
	_ = "STUB: not implemented"
	return nil
}

func (a *ActionsCacheUsageList) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (a *ActionsCreateOrgVariableRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsCreateOrgVariableRequest) GetSelectedRepositoryIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (a *ActionsCreateOrgVariableRequest) GetValue() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsCreateOrgVariableRequest) GetVisibility() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *ActionsCreateVariableRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsCreateVariableRequest) GetValue() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsEnabledOnEnterpriseRepos) GetOrganizations() []*Organization {
	_ = "STUB: not implemented"
	return nil
}

func (a *ActionsEnabledOnEnterpriseRepos) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (a *ActionsEnabledOnOrgRepos) GetRepositories() []*Repository {
	_ = "STUB: not implemented"
	return nil
}

func (a *ActionsEnabledOnOrgRepos) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (a *ActionsInboundDomains) GetFullDomains() []string { _ = "STUB: not implemented"; return nil }

func (a *ActionsInboundDomains) GetWildcardDomains() []string {
	_ = "STUB: not implemented"
	return nil
}

func (a *ActionsPermissions) GetAllowedActions() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsPermissions) GetEnabledRepositories() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsPermissions) GetSelectedActionsURL() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsPermissions) GetSHAPinningRequired() bool { _ = "STUB: not implemented"; return false }

func (a *ActionsPermissionsEnterprise) GetAllowedActions() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *ActionsPermissionsEnterprise) GetEnabledOrganizations() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *ActionsPermissionsEnterprise) GetSelectedActionsURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *ActionsPermissionsRepository) GetAllowedActions() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *ActionsPermissionsRepository) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (a *ActionsPermissionsRepository) GetSelectedActionsURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *ActionsPermissionsRepository) GetSHAPinningRequired() bool {
	_ = "STUB: not implemented"
	return false
}

func (a *ActionsUpdateOrgVariableRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsUpdateOrgVariableRequest) GetSelectedRepositoryIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (a *ActionsUpdateOrgVariableRequest) GetValue() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsUpdateOrgVariableRequest) GetVisibility() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *ActionsUpdateVariableRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsUpdateVariableRequest) GetValue() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsVariable) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (a *ActionsVariable) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsVariable) GetSelectedRepositoriesURL() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsVariable) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (a *ActionsVariable) GetValue() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsVariable) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (a *ActionsVariables) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (a *ActionsVariables) GetVariables() []*ActionsVariable { _ = "STUB: not implemented"; return nil }

func (a *ActiveCommitters) GetMaximumAdvancedSecurityCommitters() int {
	_ = "STUB: not implemented"
	return 0
}

func (a *ActiveCommitters) GetPurchasedAdvancedSecurityCommitters() int {
	_ = "STUB: not implemented"
	return 0
}

func (a *ActiveCommitters) GetRepositories() []*RepositoryActiveCommitters {
	_ = "STUB: not implemented"
	return nil
}

func (a *ActiveCommitters) GetTotalAdvancedSecurityCommitters() int {
	_ = "STUB: not implemented"
	return 0
}

func (a *ActiveCommitters) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (a *ActiveCommittersListOptions) GetAdvancedSecurityProduct() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *ActivityListStarredOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (a *ActivityListStarredOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (a *ActorLocation) GetCountryCode() string { _ = "STUB: not implemented"; return "" }

func (a *AddProjectItemOptions) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AddProjectItemOptions) GetType() *ProjectV2ItemContentType {
	_ = "STUB: not implemented"
	return nil
}

func (a *AddProjectV2FieldRequest) GetDataType() string { _ = "STUB: not implemented"; return "" }

func (a *AddProjectV2FieldRequest) GetIssueFieldID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AddProjectV2FieldRequest) GetIterationConfiguration() *ProjectV2FieldIterationConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (a *AddProjectV2FieldRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *AddProjectV2FieldRequest) GetSingleSelectOptions() []*ProjectV2FieldSingleSelectOption {
	_ = "STUB: not implemented"
	return nil
}

func (a *AddResourcesToCostCenterResponse) GetMessage() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *AddResourcesToCostCenterResponse) GetReassignedResources() []*ReassignedResource {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdminEnforcedChanges) GetFrom() bool { _ = "STUB: not implemented"; return false }

func (a *AdminEnforcement) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (a *AdminEnforcement) GetURL() string { _ = "STUB: not implemented"; return "" }

func (a *AdminStats) GetComments() *CommentStats { _ = "STUB: not implemented"; return nil }

func (a *AdminStats) GetGists() *GistStats { _ = "STUB: not implemented"; return nil }

func (a *AdminStats) GetHooks() *HookStats { _ = "STUB: not implemented"; return nil }

func (a *AdminStats) GetIssues() *IssueStats { _ = "STUB: not implemented"; return nil }

func (a *AdminStats) GetMilestones() *MilestoneStats { _ = "STUB: not implemented"; return nil }

func (a *AdminStats) GetOrgs() *OrgStats { _ = "STUB: not implemented"; return nil }

func (a *AdminStats) GetPages() *PageStats { _ = "STUB: not implemented"; return nil }

func (a *AdminStats) GetPulls() *PullStats { _ = "STUB: not implemented"; return nil }

func (a *AdminStats) GetRepos() *RepoStats { _ = "STUB: not implemented"; return nil }

func (a *AdminStats) GetUsers() *UserStats { _ = "STUB: not implemented"; return nil }

func (a *AdvancedSecurity) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (a *AdvancedSecurityCommittersBreakdown) GetLastPushedDate() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *AdvancedSecurityCommittersBreakdown) GetLastPushedEmail() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *AdvancedSecurityCommittersBreakdown) GetUserLogin() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *AdvisoryCVSS) GetScore() float64 { _ = "STUB: not implemented"; return 0 }

func (a *AdvisoryCVSS) GetVectorString() string { _ = "STUB: not implemented"; return "" }

func (a *AdvisoryCWEs) GetCWEID() string { _ = "STUB: not implemented"; return "" }

func (a *AdvisoryCWEs) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *AdvisoryEPSS) GetPercentage() float64 { _ = "STUB: not implemented"; return 0 }

func (a *AdvisoryEPSS) GetPercentile() float64 { _ = "STUB: not implemented"; return 0 }

func (a *AdvisoryIdentifier) GetType() string { _ = "STUB: not implemented"; return "" }

func (a *AdvisoryIdentifier) GetValue() string { _ = "STUB: not implemented"; return "" }

func (a *AdvisoryReference) GetURL() string { _ = "STUB: not implemented"; return "" }

func (a *AdvisoryVulnerability) GetFirstPatchedVersion() *FirstPatchedVersion {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdvisoryVulnerability) GetPackage() *VulnerabilityPackage {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdvisoryVulnerability) GetPatchedVersions() string { _ = "STUB: not implemented"; return "" }

func (a *AdvisoryVulnerability) GetSeverity() string { _ = "STUB: not implemented"; return "" }

func (a *AdvisoryVulnerability) GetVulnerableFunctions() []string {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdvisoryVulnerability) GetVulnerableVersionRange() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *Alert) GetClosedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *Alert) GetClosedBy() *User { _ = "STUB: not implemented"; return nil }

func (a *Alert) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *Alert) GetDismissedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *Alert) GetDismissedBy() *User { _ = "STUB: not implemented"; return nil }

func (a *Alert) GetDismissedComment() string { _ = "STUB: not implemented"; return "" }

func (a *Alert) GetDismissedReason() string { _ = "STUB: not implemented"; return "" }

func (a *Alert) GetFixedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *Alert) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (a *Alert) GetInstances() []*MostRecentInstance { _ = "STUB: not implemented"; return nil }

func (a *Alert) GetInstancesURL() string { _ = "STUB: not implemented"; return "" }

func (a *Alert) GetMostRecentInstance() *MostRecentInstance { _ = "STUB: not implemented"; return nil }

func (a *Alert) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (a *Alert) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (a *Alert) GetRule() *Rule { _ = "STUB: not implemented"; return nil }

func (a *Alert) GetRuleDescription() string { _ = "STUB: not implemented"; return "" }

func (a *Alert) GetRuleID() string { _ = "STUB: not implemented"; return "" }

func (a *Alert) GetRuleSeverity() string { _ = "STUB: not implemented"; return "" }

func (a *Alert) GetState() string { _ = "STUB: not implemented"; return "" }

func (a *Alert) GetTool() *Tool { _ = "STUB: not implemented"; return nil }

func (a *Alert) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *Alert) GetURL() string { _ = "STUB: not implemented"; return "" }

func (a *AlertInstancesListOptions) GetRef() string { _ = "STUB: not implemented"; return "" }

func (a *AlertListOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (a *AlertListOptions) GetRef() string { _ = "STUB: not implemented"; return "" }

func (a *AlertListOptions) GetSeverity() string { _ = "STUB: not implemented"; return "" }

func (a *AlertListOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (a *AlertListOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (a *AlertListOptions) GetToolGUID() string { _ = "STUB: not implemented"; return "" }

func (a *AlertListOptions) GetToolName() string { _ = "STUB: not implemented"; return "" }

func (a *AllowDeletions) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (a *AllowDeletionsEnforcementLevelChanges) GetFrom() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *AllowForcePushes) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (a *AllowForkSyncing) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (a *AmazonS3AccessKeysConfig) GetAuthenticationType() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *AmazonS3AccessKeysConfig) GetBucket() string { _ = "STUB: not implemented"; return "" }

func (a *AmazonS3AccessKeysConfig) GetEncryptedAccessKeyID() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *AmazonS3AccessKeysConfig) GetEncryptedSecretKey() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *AmazonS3AccessKeysConfig) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (a *AmazonS3AccessKeysConfig) GetRegion() string { _ = "STUB: not implemented"; return "" }

func (a *AmazonS3OIDCConfig) GetArnRole() string { _ = "STUB: not implemented"; return "" }

func (a *AmazonS3OIDCConfig) GetAuthenticationType() string { _ = "STUB: not implemented"; return "" }

func (a *AmazonS3OIDCConfig) GetBucket() string { _ = "STUB: not implemented"; return "" }

func (a *AmazonS3OIDCConfig) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (a *AmazonS3OIDCConfig) GetRegion() string { _ = "STUB: not implemented"; return "" }

func (a *AnalysesListOptions) GetRef() string { _ = "STUB: not implemented"; return "" }

func (a *AnalysesListOptions) GetSarifID() string { _ = "STUB: not implemented"; return "" }

func (a *APIMeta) GetActions() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetActionsMacos() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetAPI() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetCodespaces() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetCopilot() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetDependabot() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetDomains() *APIMetaDomains { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetGit() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetGithubEnterpriseImporter() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetHooks() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetImporter() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetPackages() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetPages() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetSSHKeyFingerprints() map[string]string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetSSHKeys() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMeta) GetVerifiablePasswordAuthentication() bool {
	_ = "STUB: not implemented"
	return false
}

func (a *APIMeta) GetWeb() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMetaArtifactAttestations) GetServices() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMetaArtifactAttestations) GetTrustDomain() string { _ = "STUB: not implemented"; return "" }

func (a *APIMetaDomains) GetActions() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMetaDomains) GetActionsInbound() *ActionsInboundDomains {
	_ = "STUB: not implemented"
	return nil
}

func (a *APIMetaDomains) GetArtifactAttestations() *APIMetaArtifactAttestations {
	_ = "STUB: not implemented"
	return nil
}

func (a *APIMetaDomains) GetCodespaces() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMetaDomains) GetCopilot() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMetaDomains) GetPackages() []string { _ = "STUB: not implemented"; return nil }

func (a *APIMetaDomains) GetWebsite() []string { _ = "STUB: not implemented"; return nil }

func (a *App) GetClientID() string { _ = "STUB: not implemented"; return "" }

func (a *App) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *App) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (a *App) GetEvents() []string { _ = "STUB: not implemented"; return nil }

func (a *App) GetExternalURL() string { _ = "STUB: not implemented"; return "" }

func (a *App) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (a *App) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *App) GetInstallationsCount() int { _ = "STUB: not implemented"; return 0 }

func (a *App) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *App) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (a *App) GetOwner() *User { _ = "STUB: not implemented"; return nil }

func (a *App) GetPermissions() *InstallationPermissions { _ = "STUB: not implemented"; return nil }

func (a *App) GetSlug() string { _ = "STUB: not implemented"; return "" }

func (a *App) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *AppConfig) GetClientID() string { _ = "STUB: not implemented"; return "" }

func (a *AppConfig) GetClientSecret() string { _ = "STUB: not implemented"; return "" }

func (a *AppConfig) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *AppConfig) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (a *AppConfig) GetExternalURL() string { _ = "STUB: not implemented"; return "" }

func (a *AppConfig) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (a *AppConfig) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AppConfig) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *AppConfig) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (a *AppConfig) GetOwner() *User { _ = "STUB: not implemented"; return nil }

func (a *AppConfig) GetPEM() string { _ = "STUB: not implemented"; return "" }

func (a *AppConfig) GetSlug() string { _ = "STUB: not implemented"; return "" }

func (a *AppConfig) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *AppConfig) GetWebhookSecret() string { _ = "STUB: not implemented"; return "" }

func (a *AppInstallationRepositoriesRequest) GetRepositories() []string {
	_ = "STUB: not implemented"
	return nil
}

func (a *ArchivedAt) GetFrom() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *ArchivedAt) GetTo() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *Artifact) GetArchiveDownloadURL() string { _ = "STUB: not implemented"; return "" }

func (a *Artifact) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *Artifact) GetDigest() string { _ = "STUB: not implemented"; return "" }

func (a *Artifact) GetExpired() bool { _ = "STUB: not implemented"; return false }

func (a *Artifact) GetExpiresAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *Artifact) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *Artifact) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *Artifact) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (a *Artifact) GetSizeInBytes() int64 { _ = "STUB: not implemented"; return 0 }

func (a *Artifact) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *Artifact) GetURL() string { _ = "STUB: not implemented"; return "" }

func (a *Artifact) GetWorkflowRun() *ArtifactWorkflowRun { _ = "STUB: not implemented"; return nil }

func (a *ArtifactDeploymentRecord) GetAttestationID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *ArtifactDeploymentRecord) GetCluster() string { _ = "STUB: not implemented"; return "" }

func (a *ArtifactDeploymentRecord) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (a *ArtifactDeploymentRecord) GetDeploymentName() string { _ = "STUB: not implemented"; return "" }

func (a *ArtifactDeploymentRecord) GetDigest() string { _ = "STUB: not implemented"; return "" }

func (a *ArtifactDeploymentRecord) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *ArtifactDeploymentRecord) GetLogicalEnvironment() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *ArtifactDeploymentRecord) GetPhysicalEnvironment() string {
	_ = "STUB: not implemented"
	return ""
}

func (a *ArtifactDeploymentRecord) GetRuntimeRisks() []DeploymentRuntimeRisk {
	_ = "STUB: not implemented"
	return nil
}

func (a *ArtifactDeploymentRecord) GetTags() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (a *ArtifactDeploymentRecord) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (a *ArtifactDeploymentResponse) GetDeploymentRecords() []*ArtifactDeploymentRecord {
	_ = "STUB: not implemented"
	return nil
}

func (a *ArtifactDeploymentResponse) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (a *ArtifactList) GetArtifacts() []*Artifact { _ = "STUB: not implemented"; return nil }

func (a *ArtifactList) GetTotalCount() int64 { _ = "STUB: not implemented"; return 0 }

func (a *ArtifactPeriod) GetDays() int { _ = "STUB: not implemented"; return 0 }

func (a *ArtifactPeriod) GetMaximumAllowedDays() int { _ = "STUB: not implemented"; return 0 }

func (a *ArtifactPeriodOpt) GetDays() int { _ = "STUB: not implemented"; return 0 }

func (a *ArtifactStorageRecord) GetArtifactURL() string { _ = "STUB: not implemented"; return "" }

func (a *ArtifactStorageRecord) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (a *ArtifactStorageRecord) GetDigest() string { _ = "STUB: not implemented"; return "" }

func (a *ArtifactStorageRecord) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *ArtifactStorageRecord) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *ArtifactStorageRecord) GetRegistryURL() string { _ = "STUB: not implemented"; return "" }

func (a *ArtifactStorageRecord) GetRepository() string { _ = "STUB: not implemented"; return "" }

func (a *ArtifactStorageRecord) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (a *ArtifactStorageRecord) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (a *ArtifactStorageResponse) GetStorageRecords() []*ArtifactStorageRecord {
	_ = "STUB: not implemented"
	return nil
}

func (a *ArtifactStorageResponse) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (a *ArtifactWorkflowRun) GetHeadBranch() string { _ = "STUB: not implemented"; return "" }

func (a *ArtifactWorkflowRun) GetHeadRepositoryID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *ArtifactWorkflowRun) GetHeadSHA() string { _ = "STUB: not implemented"; return "" }

func (a *ArtifactWorkflowRun) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *ArtifactWorkflowRun) GetRepositoryID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AssignmentGrade) GetAssignmentName() string { _ = "STUB: not implemented"; return "" }

func (a *AssignmentGrade) GetAssignmentURL() string { _ = "STUB: not implemented"; return "" }

func (a *AssignmentGrade) GetGithubUsername() string { _ = "STUB: not implemented"; return "" }

func (a *AssignmentGrade) GetGroupName() string { _ = "STUB: not implemented"; return "" }

func (a *AssignmentGrade) GetPointsAvailable() int { _ = "STUB: not implemented"; return 0 }

func (a *AssignmentGrade) GetPointsAwarded() int { _ = "STUB: not implemented"; return 0 }

func (a *AssignmentGrade) GetRosterIdentifier() string { _ = "STUB: not implemented"; return "" }

func (a *AssignmentGrade) GetStarterCodeURL() string { _ = "STUB: not implemented"; return "" }

func (a *AssignmentGrade) GetStudentRepositoryName() string { _ = "STUB: not implemented"; return "" }

func (a *AssignmentGrade) GetStudentRepositoryURL() string { _ = "STUB: not implemented"; return "" }

func (a *AssignmentGrade) GetSubmissionTimestamp() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (a *Attachment) GetBody() string { _ = "STUB: not implemented"; return "" }

func (a *Attachment) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *Attachment) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (a *Attestation) GetBundle() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (a *Attestation) GetRepositoryID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AttestationsResponse) GetAttestations() []*Attestation {
	_ = "STUB: not implemented"
	return nil
}

func (a *AuditEntry) GetAction() string { _ = "STUB: not implemented"; return "" }

func (a *AuditEntry) GetActor() string { _ = "STUB: not implemented"; return "" }

func (a *AuditEntry) GetActorID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AuditEntry) GetActorLocation() *ActorLocation { _ = "STUB: not implemented"; return nil }

func (a *AuditEntry) GetAdditionalFields() map[string]any { _ = "STUB: not implemented"; return nil }

func (a *AuditEntry) GetBusiness() string { _ = "STUB: not implemented"; return "" }

func (a *AuditEntry) GetBusinessID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AuditEntry) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *AuditEntry) GetData() map[string]any { _ = "STUB: not implemented"; return nil }

func (a *AuditEntry) GetDocumentID() string { _ = "STUB: not implemented"; return "" }

func (a *AuditEntry) GetExternalIdentityNameID() string { _ = "STUB: not implemented"; return "" }

func (a *AuditEntry) GetExternalIdentityUsername() string { _ = "STUB: not implemented"; return "" }

func (a *AuditEntry) GetHashedToken() string { _ = "STUB: not implemented"; return "" }

func (a *AuditEntry) GetOrg() string { _ = "STUB: not implemented"; return "" }

func (a *AuditEntry) GetOrgID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AuditEntry) GetTimestamp() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *AuditEntry) GetTokenID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AuditEntry) GetTokenScopes() string { _ = "STUB: not implemented"; return "" }

func (a *AuditEntry) GetUser() string { _ = "STUB: not implemented"; return "" }

func (a *AuditEntry) GetUserID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AuditLogStream) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (a *AuditLogStream) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (a *AuditLogStream) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AuditLogStream) GetPausedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *AuditLogStream) GetStreamDetails() string { _ = "STUB: not implemented"; return "" }

func (a *AuditLogStream) GetStreamType() string { _ = "STUB: not implemented"; return "" }

func (a *AuditLogStream) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (a *AuditLogStreamConfig) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (a *AuditLogStreamConfig) GetStreamType() string { _ = "STUB: not implemented"; return "" }

func (a *AuditLogStreamConfig) GetVendorSpecific() AuditLogStreamVendorConfig {
	_ = "STUB: not implemented"
	return *new(AuditLogStreamVendorConfig)
}

func (a *AuditLogStreamKey) GetKey() string { _ = "STUB: not implemented"; return "" }

func (a *AuditLogStreamKey) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (a *Authorization) GetApp() *AuthorizationApp { _ = "STUB: not implemented"; return nil }

func (a *Authorization) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *Authorization) GetFingerprint() string { _ = "STUB: not implemented"; return "" }

func (a *Authorization) GetHashedToken() string { _ = "STUB: not implemented"; return "" }

func (a *Authorization) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *Authorization) GetNote() string { _ = "STUB: not implemented"; return "" }

func (a *Authorization) GetNoteURL() string { _ = "STUB: not implemented"; return "" }

func (a *Authorization) GetScopes() []Scope { _ = "STUB: not implemented"; return nil }

func (a *Authorization) GetToken() string { _ = "STUB: not implemented"; return "" }

func (a *Authorization) GetTokenLastEight() string { _ = "STUB: not implemented"; return "" }

func (a *Authorization) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (a *Authorization) GetURL() string { _ = "STUB: not implemented"; return "" }

func (a *Authorization) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (a *AuthorizationApp) GetClientID() string { _ = "STUB: not implemented"; return "" }

func (a *AuthorizationApp) GetName() string { _ = "STUB: not implemented"; return "" }

func (a *AuthorizationApp) GetURL() string { _ = "STUB: not implemented"; return "" }

func (a *AuthorizationRequest) GetClientID() string { _ = "STUB: not implemented"; return "" }

func (a *AuthorizationRequest) GetClientSecret() string { _ = "STUB: not implemented"; return "" }

func (a *AuthorizationRequest) GetFingerprint() string { _ = "STUB: not implemented"; return "" }

func (a *AuthorizationRequest) GetNote() string { _ = "STUB: not implemented"; return "" }

func (a *AuthorizationRequest) GetNoteURL() string { _ = "STUB: not implemented"; return "" }

func (a *AuthorizationRequest) GetScopes() []Scope { _ = "STUB: not implemented"; return nil }

func (a *AuthorizationUpdateRequest) GetAddScopes() []string { _ = "STUB: not implemented"; return nil }

func (a *AuthorizationUpdateRequest) GetFingerprint() string { _ = "STUB: not implemented"; return "" }

func (a *AuthorizationUpdateRequest) GetNote() string { _ = "STUB: not implemented"; return "" }

func (a *AuthorizationUpdateRequest) GetNoteURL() string { _ = "STUB: not implemented"; return "" }

func (a *AuthorizationUpdateRequest) GetRemoveScopes() []string {
	_ = "STUB: not implemented"
	return nil
}

func (a *AuthorizationUpdateRequest) GetScopes() []string { _ = "STUB: not implemented"; return nil }

func (a *AuthorizedActorNames) GetFrom() []string { _ = "STUB: not implemented"; return nil }

func (a *AuthorizedActorsOnly) GetFrom() bool { _ = "STUB: not implemented"; return false }

func (a *AuthorizedDismissalActorsOnlyChanges) GetFrom() bool {
	_ = "STUB: not implemented"
	return false
}

func (a *Autolink) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *Autolink) GetIsAlphanumeric() bool { _ = "STUB: not implemented"; return false }

func (a *Autolink) GetKeyPrefix() string { _ = "STUB: not implemented"; return "" }

func (a *Autolink) GetURLTemplate() string { _ = "STUB: not implemented"; return "" }

func (a *AutolinkOptions) GetIsAlphanumeric() bool { _ = "STUB: not implemented"; return false }

func (a *AutolinkOptions) GetKeyPrefix() string { _ = "STUB: not implemented"; return "" }

func (a *AutolinkOptions) GetURLTemplate() string { _ = "STUB: not implemented"; return "" }

func (a *AutomatedSecurityFixes) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (a *AutomatedSecurityFixes) GetPaused() bool { _ = "STUB: not implemented"; return false }

func (a *AutoTriggerCheck) GetAppID() int64 { _ = "STUB: not implemented"; return 0 }

func (a *AutoTriggerCheck) GetSetting() bool { _ = "STUB: not implemented"; return false }

func (a *AzureBlobConfig) GetContainer() string { _ = "STUB: not implemented"; return "" }

func (a *AzureBlobConfig) GetEncryptedSASURL() string { _ = "STUB: not implemented"; return "" }

func (a *AzureBlobConfig) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (a *AzureHubConfig) GetEncryptedConnstring() string { _ = "STUB: not implemented"; return "" }

func (a *AzureHubConfig) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (a *AzureHubConfig) GetName() string { _ = "STUB: not implemented"; return "" }

func (b *BasicAuthTransport) GetOTP() string { _ = "STUB: not implemented"; return "" }

func (b *BasicAuthTransport) GetPassword() string { _ = "STUB: not implemented"; return "" }

func (b *BasicAuthTransport) GetUsername() string { _ = "STUB: not implemented"; return "" }

func (b *BillingCostCenter) GetID() string { _ = "STUB: not implemented"; return "" }

func (b *BillingCostCenter) GetName() string { _ = "STUB: not implemented"; return "" }

func (b *Blob) GetContent() string { _ = "STUB: not implemented"; return "" }

func (b *Blob) GetEncoding() string { _ = "STUB: not implemented"; return "" }

func (b *Blob) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (b *Blob) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (b *Blob) GetSize() int { _ = "STUB: not implemented"; return 0 }

func (b *Blob) GetURL() string { _ = "STUB: not implemented"; return "" }

func (b *BlockCreations) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (b *Branch) GetCommit() *RepositoryCommit { _ = "STUB: not implemented"; return nil }

func (b *Branch) GetName() string { _ = "STUB: not implemented"; return "" }

func (b *Branch) GetProtected() bool { _ = "STUB: not implemented"; return false }

func (b *Branch) GetProtection() *Protection { _ = "STUB: not implemented"; return nil }

func (b *Branch) GetProtectionURL() string { _ = "STUB: not implemented"; return "" }

func (b *BranchCommit) GetCommit() *Commit { _ = "STUB: not implemented"; return nil }

func (b *BranchCommit) GetName() string { _ = "STUB: not implemented"; return "" }

func (b *BranchCommit) GetProtected() bool { _ = "STUB: not implemented"; return false }

func (b *BranchListOptions) GetProtected() bool { _ = "STUB: not implemented"; return false }

func (b *BranchPolicy) GetCustomBranchPolicies() bool { _ = "STUB: not implemented"; return false }

func (b *BranchPolicy) GetProtectedBranches() bool { _ = "STUB: not implemented"; return false }

func (b *BranchProtectionConfigurationEvent) GetAction() string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BranchProtectionConfigurationEvent) GetEnterprise() *Enterprise {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchProtectionConfigurationEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchProtectionConfigurationEvent) GetOrg() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchProtectionConfigurationEvent) GetRepo() *Repository {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchProtectionConfigurationEvent) GetSender() *User {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchProtectionRule) GetAdminEnforced() bool { _ = "STUB: not implemented"; return false }

func (b *BranchProtectionRule) GetAllowDeletionsEnforcementLevel() string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BranchProtectionRule) GetAllowForcePushesEnforcementLevel() string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BranchProtectionRule) GetAuthorizedActorNames() []string {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchProtectionRule) GetAuthorizedActorsOnly() bool {
	_ = "STUB: not implemented"
	return false
}

func (b *BranchProtectionRule) GetAuthorizedDismissalActorsOnly() bool {
	_ = "STUB: not implemented"
	return false
}

func (b *BranchProtectionRule) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (b *BranchProtectionRule) GetDismissStaleReviewsOnPush() bool {
	_ = "STUB: not implemented"
	return false
}

func (b *BranchProtectionRule) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (b *BranchProtectionRule) GetIgnoreApprovalsFromContributors() bool {
	_ = "STUB: not implemented"
	return false
}

func (b *BranchProtectionRule) GetLinearHistoryRequirementEnforcementLevel() string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BranchProtectionRule) GetMergeQueueEnforcementLevel() string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BranchProtectionRule) GetName() string { _ = "STUB: not implemented"; return "" }

func (b *BranchProtectionRule) GetPullRequestReviewsEnforcementLevel() string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BranchProtectionRule) GetRepositoryID() int64 { _ = "STUB: not implemented"; return 0 }

func (b *BranchProtectionRule) GetRequireCodeOwnerReview() bool {
	_ = "STUB: not implemented"
	return false
}

func (b *BranchProtectionRule) GetRequiredApprovingReviewCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (b *BranchProtectionRule) GetRequiredConversationResolutionLevel() string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BranchProtectionRule) GetRequiredDeploymentsEnforcementLevel() string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BranchProtectionRule) GetRequiredStatusChecks() []string {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchProtectionRule) GetRequiredStatusChecksEnforcementLevel() string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BranchProtectionRule) GetRequireLastPushApproval() bool {
	_ = "STUB: not implemented"
	return false
}

func (b *BranchProtectionRule) GetSignatureRequirementEnforcementLevel() string {
	_ = "STUB: not implemented"
	return ""
}

func (b *BranchProtectionRule) GetStrictRequiredStatusChecksPolicy() bool {
	_ = "STUB: not implemented"
	return false
}

func (b *BranchProtectionRule) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (b *BranchProtectionRuleEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (b *BranchProtectionRuleEvent) GetChanges() *ProtectionChanges {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchProtectionRuleEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchProtectionRuleEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (b *BranchProtectionRuleEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (b *BranchProtectionRuleEvent) GetRule() *BranchProtectionRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchProtectionRuleEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (b *BranchRestrictions) GetApps() []*App { _ = "STUB: not implemented"; return nil }

func (b *BranchRestrictions) GetTeams() []*Team { _ = "STUB: not implemented"; return nil }

func (b *BranchRestrictions) GetUsers() []*User { _ = "STUB: not implemented"; return nil }

func (b *BranchRestrictionsRequest) GetApps() []string { _ = "STUB: not implemented"; return nil }

func (b *BranchRestrictionsRequest) GetTeams() []string { _ = "STUB: not implemented"; return nil }

func (b *BranchRestrictionsRequest) GetUsers() []string { _ = "STUB: not implemented"; return nil }

func (b *BranchRuleMetadata) GetRulesetID() int64 { _ = "STUB: not implemented"; return 0 }

func (b *BranchRuleMetadata) GetRulesetSource() string { _ = "STUB: not implemented"; return "" }

func (b *BranchRuleMetadata) GetRulesetSourceType() RulesetSourceType {
	_ = "STUB: not implemented"
	return *new(RulesetSourceType)
}

func (b *BranchRules) GetBranchNamePattern() []*PatternBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetCodeScanning() []*CodeScanningBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetCommitAuthorEmailPattern() []*PatternBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetCommitMessagePattern() []*PatternBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetCommitterEmailPattern() []*PatternBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetCopilotCodeReview() []*CopilotCodeReviewBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetCreation() []*BranchRuleMetadata { _ = "STUB: not implemented"; return nil }

func (b *BranchRules) GetDeletion() []*BranchRuleMetadata { _ = "STUB: not implemented"; return nil }

func (b *BranchRules) GetFileExtensionRestriction() []*FileExtensionRestrictionBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetFilePathRestriction() []*FilePathRestrictionBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetMaxFilePathLength() []*MaxFilePathLengthBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetMaxFileSize() []*MaxFileSizeBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetMergeQueue() []*MergeQueueBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetNonFastForward() []*BranchRuleMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetPullRequest() []*PullRequestBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetRequiredDeployments() []*RequiredDeploymentsBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetRequiredLinearHistory() []*BranchRuleMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetRequiredSignatures() []*BranchRuleMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetRequiredStatusChecks() []*RequiredStatusChecksBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetTagNamePattern() []*PatternBranchRule {
	_ = "STUB: not implemented"
	return nil
}

func (b *BranchRules) GetUpdate() []*UpdateBranchRule { _ = "STUB: not implemented"; return nil }

func (b *BranchRules) GetWorkflows() []*WorkflowsBranchRule { _ = "STUB: not implemented"; return nil }

func (b *BypassActor) GetActorID() int64 { _ = "STUB: not implemented"; return 0 }

func (b *BypassActor) GetActorType() *BypassActorType { _ = "STUB: not implemented"; return nil }

func (b *BypassActor) GetBypassMode() *BypassMode { _ = "STUB: not implemented"; return nil }

func (b *BypassPullRequestAllowances) GetApps() []*App { _ = "STUB: not implemented"; return nil }

func (b *BypassPullRequestAllowances) GetTeams() []*Team { _ = "STUB: not implemented"; return nil }

func (b *BypassPullRequestAllowances) GetUsers() []*User { _ = "STUB: not implemented"; return nil }

func (b *BypassPullRequestAllowancesRequest) GetApps() []string {
	_ = "STUB: not implemented"
	return nil
}

func (b *BypassPullRequestAllowancesRequest) GetTeams() []string {
	_ = "STUB: not implemented"
	return nil
}

func (b *BypassPullRequestAllowancesRequest) GetUsers() []string {
	_ = "STUB: not implemented"
	return nil
}

func (b *BypassReviewer) GetReviewerID() int64 { _ = "STUB: not implemented"; return 0 }

func (b *BypassReviewer) GetReviewerType() string { _ = "STUB: not implemented"; return "" }

func (b *BypassReviewer) GetSecurityConfigurationID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CheckRun) GetApp() *App { _ = "STUB: not implemented"; return nil }

func (c *CheckRun) GetCheckSuite() *CheckSuite { _ = "STUB: not implemented"; return nil }

func (c *CheckRun) GetCompletedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (c *CheckRun) GetConclusion() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRun) GetDetailsURL() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRun) GetExternalID() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRun) GetHeadSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRun) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRun) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CheckRun) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRun) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRun) GetOutput() *CheckRunOutput { _ = "STUB: not implemented"; return nil }

func (c *CheckRun) GetPullRequests() []*PullRequest { _ = "STUB: not implemented"; return nil }

func (c *CheckRun) GetStartedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (c *CheckRun) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRun) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunAction) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunAction) GetIdentifier() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunAction) GetLabel() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunAnnotation) GetAnnotationLevel() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunAnnotation) GetEndColumn() int { _ = "STUB: not implemented"; return 0 }

func (c *CheckRunAnnotation) GetEndLine() int { _ = "STUB: not implemented"; return 0 }

func (c *CheckRunAnnotation) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunAnnotation) GetPath() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunAnnotation) GetRawDetails() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunAnnotation) GetStartColumn() int { _ = "STUB: not implemented"; return 0 }

func (c *CheckRunAnnotation) GetStartLine() int { _ = "STUB: not implemented"; return 0 }

func (c *CheckRunAnnotation) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunEvent) GetCheckRun() *CheckRun { _ = "STUB: not implemented"; return nil }

func (c *CheckRunEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (c *CheckRunEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (c *CheckRunEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (c *CheckRunEvent) GetRequestedAction() *RequestedAction {
	_ = "STUB: not implemented"
	return nil
}

func (c *CheckRunEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (c *CheckRunImage) GetAlt() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunImage) GetCaption() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunImage) GetImageURL() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunOutput) GetAnnotations() []*CheckRunAnnotation {
	_ = "STUB: not implemented"
	return nil
}

func (c *CheckRunOutput) GetAnnotationsCount() int { _ = "STUB: not implemented"; return 0 }

func (c *CheckRunOutput) GetAnnotationsURL() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunOutput) GetImages() []*CheckRunImage { _ = "STUB: not implemented"; return nil }

func (c *CheckRunOutput) GetSummary() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunOutput) GetText() string { _ = "STUB: not implemented"; return "" }

func (c *CheckRunOutput) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (c *CheckSuite) GetAfterSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CheckSuite) GetApp() *App { _ = "STUB: not implemented"; return nil }

func (c *CheckSuite) GetBeforeSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CheckSuite) GetConclusion() string { _ = "STUB: not implemented"; return "" }

func (c *CheckSuite) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (c *CheckSuite) GetHeadBranch() string { _ = "STUB: not implemented"; return "" }

func (c *CheckSuite) GetHeadCommit() *Commit { _ = "STUB: not implemented"; return nil }

func (c *CheckSuite) GetHeadSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CheckSuite) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CheckSuite) GetLatestCheckRunsCount() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CheckSuite) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (c *CheckSuite) GetPullRequests() []*PullRequest { _ = "STUB: not implemented"; return nil }

func (c *CheckSuite) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (c *CheckSuite) GetRerequestable() bool { _ = "STUB: not implemented"; return false }

func (c *CheckSuite) GetRunsRerequestable() bool { _ = "STUB: not implemented"; return false }

func (c *CheckSuite) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (c *CheckSuite) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (c *CheckSuite) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *CheckSuiteEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (c *CheckSuiteEvent) GetCheckSuite() *CheckSuite { _ = "STUB: not implemented"; return nil }

func (c *CheckSuiteEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (c *CheckSuiteEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (c *CheckSuiteEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (c *CheckSuiteEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (c *CheckSuitePreferenceOptions) GetAutoTriggerChecks() []*AutoTriggerCheck {
	_ = "STUB: not implemented"
	return nil
}

func (c *CheckSuitePreferenceResults) GetPreferences() *PreferenceList {
	_ = "STUB: not implemented"
	return nil
}

func (c *CheckSuitePreferenceResults) GetRepository() *Repository {
	_ = "STUB: not implemented"
	return nil
}

func (c *Classroom) GetArchived() bool { _ = "STUB: not implemented"; return false }

func (c *Classroom) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *Classroom) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *Classroom) GetOrganization() *Organization { _ = "STUB: not implemented"; return nil }

func (c *Classroom) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *ClassroomAssignment) GetAccepted() int { _ = "STUB: not implemented"; return 0 }

func (c *ClassroomAssignment) GetClassroom() *Classroom { _ = "STUB: not implemented"; return nil }

func (c *ClassroomAssignment) GetDeadline() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *ClassroomAssignment) GetEditor() string { _ = "STUB: not implemented"; return "" }

func (c *ClassroomAssignment) GetFeedbackPullRequestsEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ClassroomAssignment) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *ClassroomAssignment) GetInvitationsEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *ClassroomAssignment) GetInviteLink() string { _ = "STUB: not implemented"; return "" }

func (c *ClassroomAssignment) GetLanguage() string { _ = "STUB: not implemented"; return "" }

func (c *ClassroomAssignment) GetMaxMembers() int { _ = "STUB: not implemented"; return 0 }

func (c *ClassroomAssignment) GetMaxTeams() int { _ = "STUB: not implemented"; return 0 }

func (c *ClassroomAssignment) GetPassing() int { _ = "STUB: not implemented"; return 0 }

func (c *ClassroomAssignment) GetPublicRepo() bool { _ = "STUB: not implemented"; return false }

func (c *ClassroomAssignment) GetSlug() string { _ = "STUB: not implemented"; return "" }

func (c *ClassroomAssignment) GetStarterCodeRepository() *Repository {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClassroomAssignment) GetStudentsAreRepoAdmins() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ClassroomAssignment) GetSubmitted() int { _ = "STUB: not implemented"; return 0 }

func (c *ClassroomAssignment) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (c *ClassroomAssignment) GetType() string { _ = "STUB: not implemented"; return "" }

func (c *ClassroomUser) GetAvatarURL() string { _ = "STUB: not implemented"; return "" }

func (c *ClassroomUser) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (c *ClassroomUser) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *ClassroomUser) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (c *ClusterArtifactDeployment) GetDeploymentName() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *ClusterArtifactDeployment) GetDigest() string { _ = "STUB: not implemented"; return "" }

func (c *ClusterArtifactDeployment) GetGithubRepository() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *ClusterArtifactDeployment) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *ClusterArtifactDeployment) GetRuntimeRisks() []DeploymentRuntimeRisk {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterArtifactDeployment) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (c *ClusterArtifactDeployment) GetTags() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterArtifactDeployment) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (c *ClusterDeploymentRecordsRequest) GetDeployments() []*ClusterArtifactDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterDeploymentRecordsRequest) GetLogicalEnvironment() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *ClusterDeploymentRecordsRequest) GetPhysicalEnvironment() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *ClusterSSHKey) GetFingerprint() string { _ = "STUB: not implemented"; return "" }

func (c *ClusterSSHKey) GetKey() string { _ = "STUB: not implemented"; return "" }

func (c *ClusterStatus) GetNodes() []*ClusterStatusNode { _ = "STUB: not implemented"; return nil }

func (c *ClusterStatus) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (c *ClusterStatusNode) GetHostname() string { _ = "STUB: not implemented"; return "" }

func (c *ClusterStatusNode) GetServices() []*ClusterStatusNodeServiceItem {
	_ = "STUB: not implemented"
	return nil
}

func (c *ClusterStatusNode) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (c *ClusterStatusNodeServiceItem) GetDetails() string { _ = "STUB: not implemented"; return "" }

func (c *ClusterStatusNodeServiceItem) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *ClusterStatusNodeServiceItem) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (c *CodeOfConduct) GetBody() string { _ = "STUB: not implemented"; return "" }

func (c *CodeOfConduct) GetKey() string { _ = "STUB: not implemented"; return "" }

func (c *CodeOfConduct) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CodeOfConduct) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *CodeownersError) GetColumn() int { _ = "STUB: not implemented"; return 0 }

func (c *CodeownersError) GetKind() string { _ = "STUB: not implemented"; return "" }

func (c *CodeownersError) GetLine() int { _ = "STUB: not implemented"; return 0 }

func (c *CodeownersError) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (c *CodeownersError) GetPath() string { _ = "STUB: not implemented"; return "" }

func (c *CodeownersError) GetSource() string { _ = "STUB: not implemented"; return "" }

func (c *CodeownersError) GetSuggestion() string { _ = "STUB: not implemented"; return "" }

func (c *CodeownersErrors) GetErrors() []*CodeownersError { _ = "STUB: not implemented"; return nil }

func (c *CodeQLDatabase) GetContentType() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQLDatabase) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CodeQLDatabase) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CodeQLDatabase) GetLanguage() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQLDatabase) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQLDatabase) GetSize() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CodeQLDatabase) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CodeQLDatabase) GetUploader() *User { _ = "STUB: not implemented"; return nil }

func (c *CodeQLDatabase) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualityFinding) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CodeQualityFinding) GetLocation() CodeQualityFindingLocation {
	_ = "STUB: not implemented"
	return *new(CodeQualityFindingLocation)
}

func (c *CodeQualityFinding) GetMessage() CodeQualityFindingMessage {
	_ = "STUB: not implemented"
	return *new(CodeQualityFindingMessage)
}

func (c *CodeQualityFinding) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (c *CodeQualityFinding) GetRule() CodeQualityFindingRule {
	_ = "STUB: not implemented"
	return *new(CodeQualityFindingRule)
}

func (c *CodeQualityFinding) GetState() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualityFinding) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualityFindingLocation) GetEndColumn() int { _ = "STUB: not implemented"; return 0 }

func (c *CodeQualityFindingLocation) GetEndLine() int { _ = "STUB: not implemented"; return 0 }

func (c *CodeQualityFindingLocation) GetPath() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualityFindingLocation) GetStartColumn() int { _ = "STUB: not implemented"; return 0 }

func (c *CodeQualityFindingLocation) GetStartLine() int { _ = "STUB: not implemented"; return 0 }

func (c *CodeQualityFindingMessage) GetMarkdown() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualityFindingMessage) GetText() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualityFindingRule) GetCategory() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualityFindingRule) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualityFindingRule) GetHelp() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualityFindingRule) GetID() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualityFindingRule) GetSeverity() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualityFindingRule) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualitySetupConfiguration) GetLanguages() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodeQualitySetupConfiguration) GetRunnerLabel() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeQualitySetupConfiguration) GetRunnerType() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeQualitySetupConfiguration) GetSchedule() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualitySetupConfiguration) GetState() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualitySetupConfiguration) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CodeQualityUpdateSetupRequest) GetLanguages() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodeQualityUpdateSetupRequest) GetRunnerLabel() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeQualityUpdateSetupRequest) GetRunnerType() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeQualityUpdateSetupRequest) GetState() string { _ = "STUB: not implemented"; return "" }

func (c *CodeQualityUpdateSetupResponse) GetRunID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CodeQualityUpdateSetupResponse) GetRunURL() string { _ = "STUB: not implemented"; return "" }

func (c *CodeResult) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (c *CodeResult) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CodeResult) GetPath() string { _ = "STUB: not implemented"; return "" }

func (c *CodeResult) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (c *CodeResult) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CodeResult) GetTextMatches() []*TextMatch { _ = "STUB: not implemented"; return nil }

func (c *CodeScanningAlertEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (c *CodeScanningAlertEvent) GetAlert() *Alert { _ = "STUB: not implemented"; return nil }

func (c *CodeScanningAlertEvent) GetCommitOID() string { _ = "STUB: not implemented"; return "" }

func (c *CodeScanningAlertEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodeScanningAlertEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (c *CodeScanningAlertEvent) GetRef() string { _ = "STUB: not implemented"; return "" }

func (c *CodeScanningAlertEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (c *CodeScanningAlertEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (c *CodeScanningAlertState) GetDismissedComment() string { _ = "STUB: not implemented"; return "" }

func (c *CodeScanningAlertState) GetDismissedReason() string { _ = "STUB: not implemented"; return "" }

func (c *CodeScanningAlertState) GetState() string { _ = "STUB: not implemented"; return "" }

func (c *CodeScanningBranchRule) GetParameters() CodeScanningRuleParameters {
	_ = "STUB: not implemented"
	return *new(CodeScanningRuleParameters)
}

func (c *CodeScanningDefaultSetupOptions) GetRunnerLabel() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeScanningDefaultSetupOptions) GetRunnerType() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeScanningOptions) GetAllowAdvanced() bool { _ = "STUB: not implemented"; return false }

func (c *CodeScanningRuleParameters) GetCodeScanningTools() []*RuleCodeScanningTool {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodeSearchResult) GetCodeResults() []*CodeResult { _ = "STUB: not implemented"; return nil }

func (c *CodeSearchResult) GetIncompleteResults() bool { _ = "STUB: not implemented"; return false }

func (c *CodeSearchResult) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (c *CodeSecurity) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (c *CodeSecurityConfiguration) GetAdvancedSecurity() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetCodeScanningDefaultSetup() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetCodeScanningDefaultSetupOptions() *CodeScanningDefaultSetupOptions {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodeSecurityConfiguration) GetCodeScanningDelegatedAlertDismissal() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetCodeScanningOptions() *CodeScanningOptions {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodeSecurityConfiguration) GetCodeSecurity() string { _ = "STUB: not implemented"; return "" }

func (c *CodeSecurityConfiguration) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CodeSecurityConfiguration) GetDependabotAlerts() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetDependabotDelegatedAlertDismissal() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetDependabotSecurityUpdates() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetDependencyGraph() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetDependencyGraphAutosubmitAction() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetDependencyGraphAutosubmitActionOptions() *DependencyGraphAutosubmitActionOptions {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodeSecurityConfiguration) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (c *CodeSecurityConfiguration) GetEnforcement() string { _ = "STUB: not implemented"; return "" }

func (c *CodeSecurityConfiguration) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (c *CodeSecurityConfiguration) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CodeSecurityConfiguration) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CodeSecurityConfiguration) GetPrivateVulnerabilityReporting() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetSecretProtection() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetSecretScanning() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetSecretScanningDelegatedAlertDismissal() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetSecretScanningDelegatedBypass() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetSecretScanningDelegatedBypassOptions() *SecretScanningDelegatedBypassOptions {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodeSecurityConfiguration) GetSecretScanningExtendedMetadata() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetSecretScanningGenericSecrets() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetSecretScanningNonProviderPatterns() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetSecretScanningPushProtection() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetSecretScanningValidityChecks() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodeSecurityConfiguration) GetTargetType() string { _ = "STUB: not implemented"; return "" }

func (c *CodeSecurityConfiguration) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CodeSecurityConfiguration) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *CodeSecurityConfigurationWithDefaultForNewRepos) GetConfiguration() *CodeSecurityConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodeSecurityConfigurationWithDefaultForNewRepos) GetDefaultForNewRepos() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Codespace) GetBillableOwner() *User { _ = "STUB: not implemented"; return nil }

func (c *Codespace) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (c *Codespace) GetDevcontainerPath() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetDisplayName() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetEnvironmentID() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetGitStatus() *CodespacesGitStatus { _ = "STUB: not implemented"; return nil }

func (c *Codespace) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *Codespace) GetIdleTimeoutMinutes() int { _ = "STUB: not implemented"; return 0 }

func (c *Codespace) GetIdleTimeoutNotice() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetLastKnownStopNotice() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetLastUsedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (c *Codespace) GetLocation() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetMachine() *CodespacesMachine { _ = "STUB: not implemented"; return nil }

func (c *Codespace) GetMachinesURL() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetOwner() *User { _ = "STUB: not implemented"; return nil }

func (c *Codespace) GetPendingOperation() bool { _ = "STUB: not implemented"; return false }

func (c *Codespace) GetPendingOperationDisabledReason() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Codespace) GetPrebuild() bool { _ = "STUB: not implemented"; return false }

func (c *Codespace) GetPullsURL() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetRecentFolders() []string { _ = "STUB: not implemented"; return nil }

func (c *Codespace) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (c *Codespace) GetRetentionExpiresAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *Codespace) GetRetentionPeriodMinutes() int { _ = "STUB: not implemented"; return 0 }

func (c *Codespace) GetRuntimeConstraints() *CodespacesRuntimeConstraints {
	_ = "STUB: not implemented"
	return nil
}

func (c *Codespace) GetStartURL() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetState() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetStopURL() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (c *Codespace) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *Codespace) GetWebURL() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceCreateForUserOptions) GetClientIP() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceCreateForUserOptions) GetDevcontainerPath() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodespaceCreateForUserOptions) GetDisplayName() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodespaceCreateForUserOptions) GetGeo() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceCreateForUserOptions) GetIdleTimeoutMinutes() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CodespaceCreateForUserOptions) GetLocation() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceCreateForUserOptions) GetMachine() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceCreateForUserOptions) GetMultiRepoPermissionsOptOut() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CodespaceCreateForUserOptions) GetPullRequest() *CodespacePullRequestOptions {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodespaceCreateForUserOptions) GetRef() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceCreateForUserOptions) GetRepositoryID() int64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *CodespaceCreateForUserOptions) GetRetentionPeriodMinutes() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CodespaceCreateForUserOptions) GetWorkingDirectory() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodespaceDefaultAttributes) GetBillableOwner() *User {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodespaceDefaultAttributes) GetDefaults() *CodespaceDefaults {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodespaceDefaults) GetDevcontainerPath() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceDefaults) GetLocation() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceExport) GetBranch() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceExport) GetCompletedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CodespaceExport) GetExportURL() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceExport) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceExport) GetID() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceExport) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceExport) GetState() string { _ = "STUB: not implemented"; return "" }

func (c *CodespaceGetDefaultAttributesOptions) GetClientIP() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodespaceGetDefaultAttributesOptions) GetRef() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodespacePermissions) GetAccepted() bool { _ = "STUB: not implemented"; return false }

func (c *CodespacePullRequestOptions) GetPullRequestNumber() int64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *CodespacePullRequestOptions) GetRepositoryID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CodespacesGitStatus) GetAhead() int { _ = "STUB: not implemented"; return 0 }

func (c *CodespacesGitStatus) GetBehind() int { _ = "STUB: not implemented"; return 0 }

func (c *CodespacesGitStatus) GetHasUncommittedChanges() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CodespacesGitStatus) GetHasUnpushedChanges() bool { _ = "STUB: not implemented"; return false }

func (c *CodespacesGitStatus) GetRef() string { _ = "STUB: not implemented"; return "" }

func (c *CodespacesMachine) GetCPUs() int { _ = "STUB: not implemented"; return 0 }

func (c *CodespacesMachine) GetDisplayName() string { _ = "STUB: not implemented"; return "" }

func (c *CodespacesMachine) GetMemoryInBytes() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CodespacesMachine) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CodespacesMachine) GetOperatingSystem() string { _ = "STUB: not implemented"; return "" }

func (c *CodespacesMachine) GetPrebuildAvailability() string { _ = "STUB: not implemented"; return "" }

func (c *CodespacesMachine) GetStorageInBytes() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CodespacesMachines) GetMachines() []*CodespacesMachine {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodespacesMachines) GetTotalCount() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CodespacesOrgAccessControlRequest) GetSelectedUsernames() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CodespacesOrgAccessControlRequest) GetVisibility() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CodespacesRuntimeConstraints) GetAllowedPortPrivacySettings() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CollaboratorInvitation) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CollaboratorInvitation) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (c *CollaboratorInvitation) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CollaboratorInvitation) GetInvitee() *User { _ = "STUB: not implemented"; return nil }

func (c *CollaboratorInvitation) GetInviter() *User { _ = "STUB: not implemented"; return nil }

func (c *CollaboratorInvitation) GetPermissions() string { _ = "STUB: not implemented"; return "" }

func (c *CollaboratorInvitation) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (c *CollaboratorInvitation) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *CombinedStatus) GetCommitURL() string { _ = "STUB: not implemented"; return "" }

func (c *CombinedStatus) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CombinedStatus) GetRepositoryURL() string { _ = "STUB: not implemented"; return "" }

func (c *CombinedStatus) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CombinedStatus) GetState() string { _ = "STUB: not implemented"; return "" }

func (c *CombinedStatus) GetStatuses() []*RepoStatus { _ = "STUB: not implemented"; return nil }

func (c *CombinedStatus) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (c *Comment) GetBody() string { _ = "STUB: not implemented"; return "" }

func (c *Comment) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (c *CommentDiscussion) GetAuthorAssociation() string { _ = "STUB: not implemented"; return "" }

func (c *CommentDiscussion) GetBody() string { _ = "STUB: not implemented"; return "" }

func (c *CommentDiscussion) GetChildCommentCount() int { _ = "STUB: not implemented"; return 0 }

func (c *CommentDiscussion) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CommentDiscussion) GetDiscussionID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CommentDiscussion) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommentDiscussion) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CommentDiscussion) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (c *CommentDiscussion) GetParentID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CommentDiscussion) GetReactions() *Reactions { _ = "STUB: not implemented"; return nil }

func (c *CommentDiscussion) GetRepositoryURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommentDiscussion) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CommentDiscussion) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (c *CommentStats) GetTotalCommitComments() int { _ = "STUB: not implemented"; return 0 }

func (c *CommentStats) GetTotalGistComments() int { _ = "STUB: not implemented"; return 0 }

func (c *CommentStats) GetTotalIssueComments() int { _ = "STUB: not implemented"; return 0 }

func (c *CommentStats) GetTotalPullRequestComments() int { _ = "STUB: not implemented"; return 0 }

func (c *Commit) GetAuthor() *CommitAuthor { _ = "STUB: not implemented"; return nil }

func (c *Commit) GetCommentCount() int { _ = "STUB: not implemented"; return 0 }

func (c *Commit) GetCommitter() *CommitAuthor { _ = "STUB: not implemented"; return nil }

func (c *Commit) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (c *Commit) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (c *Commit) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (c *Commit) GetParents() []*Commit { _ = "STUB: not implemented"; return nil }

func (c *Commit) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (c *Commit) GetTree() *Tree { _ = "STUB: not implemented"; return nil }

func (c *Commit) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *Commit) GetVerification() *SignatureVerification { _ = "STUB: not implemented"; return nil }

func (c *CommitAuthor) GetDate() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (c *CommitAuthor) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (c *CommitAuthor) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (c *CommitAuthor) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CommitCommentEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (c *CommitCommentEvent) GetComment() *RepositoryComment { _ = "STUB: not implemented"; return nil }

func (c *CommitCommentEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (c *CommitCommentEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (c *CommitCommentEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (c *CommitCommentEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (c *CommitFile) GetAdditions() int { _ = "STUB: not implemented"; return 0 }

func (c *CommitFile) GetBlobURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommitFile) GetChanges() int { _ = "STUB: not implemented"; return 0 }

func (c *CommitFile) GetContentsURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommitFile) GetDeletions() int { _ = "STUB: not implemented"; return 0 }

func (c *CommitFile) GetFilename() string { _ = "STUB: not implemented"; return "" }

func (c *CommitFile) GetPatch() string { _ = "STUB: not implemented"; return "" }

func (c *CommitFile) GetPreviousFilename() string { _ = "STUB: not implemented"; return "" }

func (c *CommitFile) GetRawURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommitFile) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CommitFile) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (c *CommitResult) GetAuthor() *User { _ = "STUB: not implemented"; return nil }

func (c *CommitResult) GetCommentsURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommitResult) GetCommit() *Commit { _ = "STUB: not implemented"; return nil }

func (c *CommitResult) GetCommitter() *User { _ = "STUB: not implemented"; return nil }

func (c *CommitResult) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommitResult) GetParents() []*Commit { _ = "STUB: not implemented"; return nil }

func (c *CommitResult) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (c *CommitResult) GetScore() float64 { _ = "STUB: not implemented"; return 0 }

func (c *CommitResult) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CommitResult) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommitsComparison) GetAheadBy() int { _ = "STUB: not implemented"; return 0 }

func (c *CommitsComparison) GetBaseCommit() *RepositoryCommit {
	_ = "STUB: not implemented"
	return nil
}

func (c *CommitsComparison) GetBehindBy() int { _ = "STUB: not implemented"; return 0 }

func (c *CommitsComparison) GetCommits() []*RepositoryCommit { _ = "STUB: not implemented"; return nil }

func (c *CommitsComparison) GetDiffURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommitsComparison) GetFiles() []*CommitFile { _ = "STUB: not implemented"; return nil }

func (c *CommitsComparison) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommitsComparison) GetMergeBaseCommit() *RepositoryCommit {
	_ = "STUB: not implemented"
	return nil
}

func (c *CommitsComparison) GetPatchURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommitsComparison) GetPermalinkURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommitsComparison) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (c *CommitsComparison) GetTotalCommits() int { _ = "STUB: not implemented"; return 0 }

func (c *CommitsComparison) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *CommitsListOptions) GetAuthor() string { _ = "STUB: not implemented"; return "" }

func (c *CommitsListOptions) GetPath() string { _ = "STUB: not implemented"; return "" }

func (c *CommitsListOptions) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CommitsListOptions) GetSince() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (c *CommitsListOptions) GetUntil() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (c *CommitsSearchResult) GetCommits() []*CommitResult { _ = "STUB: not implemented"; return nil }

func (c *CommitsSearchResult) GetIncompleteResults() bool { _ = "STUB: not implemented"; return false }

func (c *CommitsSearchResult) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (c *CommitStats) GetAdditions() int { _ = "STUB: not implemented"; return 0 }

func (c *CommitStats) GetDeletions() int { _ = "STUB: not implemented"; return 0 }

func (c *CommitStats) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (c *CommunityHealthFiles) GetCodeOfConduct() *Metric { _ = "STUB: not implemented"; return nil }

func (c *CommunityHealthFiles) GetCodeOfConductFile() *Metric {
	_ = "STUB: not implemented"
	return nil
}

func (c *CommunityHealthFiles) GetContributing() *Metric { _ = "STUB: not implemented"; return nil }

func (c *CommunityHealthFiles) GetIssueTemplate() *Metric { _ = "STUB: not implemented"; return nil }

func (c *CommunityHealthFiles) GetLicense() *Metric { _ = "STUB: not implemented"; return nil }

func (c *CommunityHealthFiles) GetPullRequestTemplate() *Metric {
	_ = "STUB: not implemented"
	return nil
}

func (c *CommunityHealthFiles) GetReadme() *Metric { _ = "STUB: not implemented"; return nil }

func (c *CommunityHealthMetrics) GetContentReportsEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CommunityHealthMetrics) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (c *CommunityHealthMetrics) GetDocumentation() string { _ = "STUB: not implemented"; return "" }

func (c *CommunityHealthMetrics) GetFiles() *CommunityHealthFiles {
	_ = "STUB: not implemented"
	return nil
}

func (c *CommunityHealthMetrics) GetHealthPercentage() int { _ = "STUB: not implemented"; return 0 }

func (c *CommunityHealthMetrics) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *ConfigApplyEvents) GetNodes() []*ConfigApplyEventsNode {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigApplyEventsNode) GetEvents() []*ConfigApplyEventsNodeEvent {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigApplyEventsNode) GetLastRequestID() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyEventsNode) GetNode() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyEventsNodeEvent) GetBody() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyEventsNodeEvent) GetConfigRunID() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyEventsNodeEvent) GetEventName() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyEventsNodeEvent) GetHostname() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyEventsNodeEvent) GetSeverityText() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyEventsNodeEvent) GetSpanDepth() int { _ = "STUB: not implemented"; return 0 }

func (c *ConfigApplyEventsNodeEvent) GetSpanID() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyEventsNodeEvent) GetSpanParentID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *ConfigApplyEventsNodeEvent) GetTimestamp() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *ConfigApplyEventsNodeEvent) GetTopology() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyEventsNodeEvent) GetTraceID() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyEventsOptions) GetLastRequestID() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyOptions) GetRunID() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyStatus) GetNodes() []*ConfigApplyStatusNode {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigApplyStatus) GetRunning() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigApplyStatus) GetSuccessful() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigApplyStatusNode) GetHostname() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyStatusNode) GetRunID() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigApplyStatusNode) GetRunning() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigApplyStatusNode) GetSuccessful() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettings) GetAdminPassword() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettings) GetAssets() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettings) GetAuthMode() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettings) GetAvatar() *ConfigSettingsAvatar { _ = "STUB: not implemented"; return nil }

func (c *ConfigSettings) GetCAS() *ConfigSettingsCAS { _ = "STUB: not implemented"; return nil }

func (c *ConfigSettings) GetCollectd() *ConfigSettingsCollectd {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigSettings) GetConfigurationID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *ConfigSettings) GetConfigurationRunCount() int { _ = "STUB: not implemented"; return 0 }

func (c *ConfigSettings) GetCustomer() *ConfigSettingsCustomer {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigSettings) GetExpireSessions() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettings) GetGithubHostname() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettings) GetGithubOAuth() *ConfigSettingsGithubOAuth {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigSettings) GetGithubSSL() *ConfigSettingsGithubSSL {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigSettings) GetHTTPProxy() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettings) GetIdenticonsHost() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettings) GetLDAP() *ConfigSettingsLDAP { _ = "STUB: not implemented"; return nil }

func (c *ConfigSettings) GetLicense() *ConfigSettingsLicenseSettings {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigSettings) GetLoadBalancer() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettings) GetMapping() *ConfigSettingsMapping { _ = "STUB: not implemented"; return nil }

func (c *ConfigSettings) GetNTP() *ConfigSettingsNTP { _ = "STUB: not implemented"; return nil }

func (c *ConfigSettings) GetPages() *ConfigSettingsPagesSettings {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigSettings) GetPrivateMode() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettings) GetPublicPages() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettings) GetSAML() *ConfigSettingsSAML { _ = "STUB: not implemented"; return nil }

func (c *ConfigSettings) GetSignupEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettings) GetSMTP() *ConfigSettingsSMTP { _ = "STUB: not implemented"; return nil }

func (c *ConfigSettings) GetSNMP() *ConfigSettingsSNMP { _ = "STUB: not implemented"; return nil }

func (c *ConfigSettings) GetSubdomainIsolation() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettings) GetSyslog() *ConfigSettingsSyslog { _ = "STUB: not implemented"; return nil }

func (c *ConfigSettings) GetTimezone() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsAvatar) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsAvatar) GetURI() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsCAS) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsCollectd) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsCollectd) GetEncryption() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsCollectd) GetPassword() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsCollectd) GetPort() int { _ = "STUB: not implemented"; return 0 }

func (c *ConfigSettingsCollectd) GetServer() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsCollectd) GetUsername() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsCustomer) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsCustomer) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsCustomer) GetPublicKeyData() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsCustomer) GetSecret() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsCustomer) GetUUID() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsGithubOAuth) GetClientID() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsGithubOAuth) GetClientSecret() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsGithubOAuth) GetOrganizationName() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *ConfigSettingsGithubOAuth) GetOrganizationTeam() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *ConfigSettingsGithubSSL) GetCert() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsGithubSSL) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsGithubSSL) GetKey() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAP) GetAdminGroup() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAP) GetBase() []string { _ = "STUB: not implemented"; return nil }

func (c *ConfigSettingsLDAP) GetBindDN() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAP) GetHost() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAP) GetMethod() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAP) GetPassword() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAP) GetPort() int { _ = "STUB: not implemented"; return 0 }

func (c *ConfigSettingsLDAP) GetPosixSupport() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsLDAP) GetProfile() *ConfigSettingsLDAPProfile {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigSettingsLDAP) GetReconciliation() *ConfigSettingsLDAPReconciliation {
	_ = "STUB: not implemented"
	return nil
}

func (c *ConfigSettingsLDAP) GetRecursiveGroupSearch() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ConfigSettingsLDAP) GetSearchStrategy() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAP) GetSyncEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsLDAP) GetTeamSyncInterval() int { _ = "STUB: not implemented"; return 0 }

func (c *ConfigSettingsLDAP) GetUID() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAP) GetUserGroups() []string { _ = "STUB: not implemented"; return nil }

func (c *ConfigSettingsLDAP) GetUserSyncEmails() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsLDAP) GetUserSyncInterval() int { _ = "STUB: not implemented"; return 0 }

func (c *ConfigSettingsLDAP) GetUserSyncKeys() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsLDAP) GetVirtualAttributeEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ConfigSettingsLDAPProfile) GetKey() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAPProfile) GetMail() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAPProfile) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAPProfile) GetUID() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAPReconciliation) GetOrg() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLDAPReconciliation) GetUser() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsLicenseSettings) GetClusterSupport() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ConfigSettingsLicenseSettings) GetEvaluation() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ConfigSettingsLicenseSettings) GetExpireAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *ConfigSettingsLicenseSettings) GetPerpetual() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ConfigSettingsLicenseSettings) GetSeats() int { _ = "STUB: not implemented"; return 0 }

func (c *ConfigSettingsLicenseSettings) GetSSHAllowed() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ConfigSettingsLicenseSettings) GetSupportKey() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *ConfigSettingsLicenseSettings) GetUnlimitedSeating() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ConfigSettingsMapping) GetBasemap() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsMapping) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsMapping) GetTileserver() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsMapping) GetToken() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsNTP) GetPrimaryServer() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsNTP) GetSecondaryServer() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsPagesSettings) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsSAML) GetCertificate() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSAML) GetCertificatePath() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSAML) GetDisableAdminDemote() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsSAML) GetIDPInitiatedSSO() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsSAML) GetIssuer() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSAML) GetSSOURL() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSMTP) GetAddress() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSMTP) GetAuthentication() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSMTP) GetDiscardToNoreplyAddress() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *ConfigSettingsSMTP) GetDomain() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSMTP) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsSMTP) GetEnableStarttlsAuto() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsSMTP) GetNoreplyAddress() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSMTP) GetPassword() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSMTP) GetPort() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSMTP) GetSupportAddress() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSMTP) GetSupportAddressType() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSMTP) GetUsername() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSMTP) GetUserName() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSNMP) GetCommunity() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSNMP) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsSyslog) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *ConfigSettingsSyslog) GetProtocolName() string { _ = "STUB: not implemented"; return "" }

func (c *ConfigSettingsSyslog) GetServer() string { _ = "STUB: not implemented"; return "" }

func (c *ConnectionServiceItem) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *ConnectionServiceItem) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (c *ContentReference) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *ContentReference) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (c *ContentReference) GetReference() string { _ = "STUB: not implemented"; return "" }

func (c *ContentReferenceEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (c *ContentReferenceEvent) GetContentReference() *ContentReference {
	_ = "STUB: not implemented"
	return nil
}

func (c *ContentReferenceEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (c *ContentReferenceEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (c *ContentReferenceEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (c *Contributor) GetAvatarURL() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetContributions() int { _ = "STUB: not implemented"; return 0 }

func (c *Contributor) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetEventsURL() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetFollowersURL() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetFollowingURL() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetGistsURL() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetGravatarID() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *Contributor) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetOrganizationsURL() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetReceivedEventsURL() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetReposURL() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetSiteAdmin() bool { _ = "STUB: not implemented"; return false }

func (c *Contributor) GetStarredURL() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetSubscriptionsURL() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetType() string { _ = "STUB: not implemented"; return "" }

func (c *Contributor) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *ContributorApprovalPermissions) GetApprovalPolicy() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *ContributorStats) GetAuthor() *Contributor { _ = "STUB: not implemented"; return nil }

func (c *ContributorStats) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (c *ContributorStats) GetWeeks() []*WeeklyStats { _ = "STUB: not implemented"; return nil }

func (c *CopilotCloudAgentConfiguration) GetCustomAllowlist() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotCloudAgentConfiguration) GetEnabledTools() *CopilotCloudAgentEnabledTools {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotCloudAgentConfiguration) GetIsFirewallEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotCloudAgentConfiguration) GetIsFirewallRecommendedAllowlistEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotCloudAgentConfiguration) GetMCPConfiguration() any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (c *CopilotCloudAgentConfiguration) GetRequireActionsWorkflowApproval() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotCloudAgentEnabledTools) GetCodeql() bool { _ = "STUB: not implemented"; return false }

func (c *CopilotCloudAgentEnabledTools) GetCopilotCodeReview() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotCloudAgentEnabledTools) GetDependencyVulnerabilityChecks() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotCloudAgentEnabledTools) GetSecretScanning() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotCodeReviewBranchRule) GetParameters() CopilotCodeReviewRuleParameters {
	_ = "STUB: not implemented"
	return *new(CopilotCodeReviewRuleParameters)
}

func (c *CopilotCodeReviewRuleParameters) GetReviewDraftPullRequests() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotCodeReviewRuleParameters) GetReviewOnPush() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotDailyMetrics) GetCodeAcceptanceActivityCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotDailyMetrics) GetCodeGenerationActivityCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotDailyMetrics) GetDailyActiveCLIUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDailyMetrics) GetDailyActiveCopilotCloudAgentUsers() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotDailyMetrics) GetDailyActiveUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDailyMetrics) GetDay() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotDailyMetrics) GetEnterpriseID() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotDailyMetrics) GetLOCAddedSum() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDailyMetrics) GetLOCDeletedSum() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDailyMetrics) GetLOCSuggestedToAddSum() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDailyMetrics) GetLOCSuggestedToDeleteSum() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDailyMetrics) GetMonthlyActiveAgentUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDailyMetrics) GetMonthlyActiveChatUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDailyMetrics) GetMonthlyActiveCopilotCloudAgentUsers() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotDailyMetrics) GetMonthlyActiveUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDailyMetrics) GetOrganizationID() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotDailyMetrics) GetPullRequests() *CopilotMetricsPullRequests {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotDailyMetrics) GetTotalsByCLI() *CopilotMetricsCLI {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotDailyMetrics) GetTotalsByFeature() []*CopilotMetricsFeature {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotDailyMetrics) GetTotalsByIDE() []*CopilotMetricsIDE {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotDailyMetrics) GetTotalsByLanguageFeature() []*CopilotMetricsLanguageFeature {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotDailyMetrics) GetTotalsByLanguageModel() []*CopilotMetricsLanguageModel {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotDailyMetrics) GetTotalsByModelFeature() []*CopilotMetricsModelFeature {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotDailyMetrics) GetUserInitiatedInteractionCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotDailyMetrics) GetWeeklyActiveCopilotCloudAgentUsers() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotDailyMetrics) GetWeeklyActiveUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDailyMetricsReport) GetDownloadLinks() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotDailyMetricsReport) GetReportDay() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotDotcomChat) GetModels() []*CopilotDotcomChatModel {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotDotcomChat) GetTotalEngagedUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDotcomChatModel) GetCustomModelTrainingDate() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CopilotDotcomChatModel) GetIsCustomModel() bool { _ = "STUB: not implemented"; return false }

func (c *CopilotDotcomChatModel) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotDotcomChatModel) GetTotalChats() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDotcomChatModel) GetTotalEngagedUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDotcomPullRequests) GetRepositories() []*CopilotDotcomPullRequestsRepository {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotDotcomPullRequests) GetTotalEngagedUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotDotcomPullRequestsModel) GetCustomModelTrainingDate() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CopilotDotcomPullRequestsModel) GetIsCustomModel() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotDotcomPullRequestsModel) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotDotcomPullRequestsModel) GetTotalEngagedUsers() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotDotcomPullRequestsModel) GetTotalPRSummariesCreated() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotDotcomPullRequestsRepository) GetModels() []*CopilotDotcomPullRequestsModel {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotDotcomPullRequestsRepository) GetName() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CopilotDotcomPullRequestsRepository) GetTotalEngagedUsers() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotIDEChat) GetEditors() []*CopilotIDEChatEditor {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotIDEChat) GetTotalEngagedUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotIDEChatEditor) GetModels() []*CopilotIDEChatModel {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotIDEChatEditor) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotIDEChatEditor) GetTotalEngagedUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotIDEChatModel) GetCustomModelTrainingDate() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CopilotIDEChatModel) GetIsCustomModel() bool { _ = "STUB: not implemented"; return false }

func (c *CopilotIDEChatModel) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotIDEChatModel) GetTotalChatCopyEvents() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotIDEChatModel) GetTotalChatInsertionEvents() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotIDEChatModel) GetTotalChats() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotIDEChatModel) GetTotalEngagedUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotIDECodeCompletions) GetEditors() []*CopilotIDECodeCompletionsEditor {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotIDECodeCompletions) GetLanguages() []*CopilotIDECodeCompletionsLanguage {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotIDECodeCompletions) GetTotalEngagedUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotIDECodeCompletionsEditor) GetModels() []*CopilotIDECodeCompletionsModel {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotIDECodeCompletionsEditor) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotIDECodeCompletionsEditor) GetTotalEngagedUsers() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotIDECodeCompletionsLanguage) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotIDECodeCompletionsLanguage) GetTotalEngagedUsers() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotIDECodeCompletionsModel) GetCustomModelTrainingDate() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CopilotIDECodeCompletionsModel) GetIsCustomModel() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotIDECodeCompletionsModel) GetLanguages() []*CopilotIDECodeCompletionsModelLanguage {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotIDECodeCompletionsModel) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotIDECodeCompletionsModel) GetTotalEngagedUsers() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotIDECodeCompletionsModelLanguage) GetName() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CopilotIDECodeCompletionsModelLanguage) GetTotalCodeAcceptances() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotIDECodeCompletionsModelLanguage) GetTotalCodeLinesAccepted() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotIDECodeCompletionsModelLanguage) GetTotalCodeLinesSuggested() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotIDECodeCompletionsModelLanguage) GetTotalCodeSuggestions() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotIDECodeCompletionsModelLanguage) GetTotalEngagedUsers() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetrics) GetCopilotDotcomChat() *CopilotDotcomChat {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotMetrics) GetCopilotDotcomPullRequests() *CopilotDotcomPullRequests {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotMetrics) GetCopilotIDEChat() *CopilotIDEChat { _ = "STUB: not implemented"; return nil }

func (c *CopilotMetrics) GetCopilotIDECodeCompletions() *CopilotIDECodeCompletions {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotMetrics) GetDate() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotMetrics) GetTotalActiveUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetrics) GetTotalEngagedUsers() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsChatPanel) GetChatPanelAgentMode() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsChatPanel) GetChatPanelAskMode() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsChatPanel) GetChatPanelCustomMode() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsChatPanel) GetChatPanelEditMode() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsChatPanel) GetChatPanelUnknownMode() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsCLI) GetLastKnownCLIVersion() *CopilotMetricsCLIVersion {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotMetricsCLI) GetPromptCount() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsCLI) GetRequestCount() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsCLI) GetSessionCount() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsCLI) GetTokenUsage() *CopilotMetricsCLITokenUsage {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotMetricsCLITokenUsage) GetAvgTokensPerRequest() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsCLITokenUsage) GetOutputTokensSum() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsCLITokenUsage) GetPromptTokensSum() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsCLIVersion) GetCLIVersion() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotMetricsCLIVersion) GetSampledAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CopilotMetricsCodeActivity) GetCodeAcceptanceActivityCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsCodeActivity) GetCodeGenerationActivityCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsCodeActivity) GetLOCAddedSum() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsCodeActivity) GetLOCDeletedSum() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsCodeActivity) GetLOCSuggestedToAddSum() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsCodeActivity) GetLOCSuggestedToDeleteSum() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsFeature) GetFeature() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotMetricsFeature) GetUserInitiatedInteractionCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsIDE) GetIDE() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotMetricsIDE) GetUserInitiatedInteractionCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsLanguageFeature) GetFeature() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotMetricsLanguageFeature) GetLanguage() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotMetricsLanguageModel) GetLanguage() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotMetricsLanguageModel) GetModel() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotMetricsListOptions) GetSince() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (c *CopilotMetricsListOptions) GetUntil() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (c *CopilotMetricsModelFeature) GetFeature() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotMetricsModelFeature) GetModel() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotMetricsModelFeature) GetUserInitiatedInteractionCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsPullRequests) GetMedianMinutesToMerge() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsPullRequests) GetMedianMinutesToMergeCopilotAuthored() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsPullRequests) GetMedianMinutesToMergeCopilotReviewed() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsPullRequests) GetTotalAppliedSuggestions() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsPullRequests) GetTotalCopilotAppliedSuggestions() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsPullRequests) GetTotalCopilotSuggestions() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsPullRequests) GetTotalCreated() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsPullRequests) GetTotalCreatedByCopilot() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsPullRequests) GetTotalMerged() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsPullRequests) GetTotalMergedCreatedByCopilot() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsPullRequests) GetTotalMergedReviewedByCopilot() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsPullRequests) GetTotalReviewed() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsPullRequests) GetTotalReviewedByCopilot() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotMetricsPullRequests) GetTotalSuggestions() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotMetricsReport) GetDownloadLinks() []string { _ = "STUB: not implemented"; return nil }

func (c *CopilotMetricsReport) GetReportEndDay() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotMetricsReport) GetReportStartDay() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotMetricsReportOptions) GetDay() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotOrganizationDetails) GetCopilotChat() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotOrganizationDetails) GetPublicCodeSuggestions() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CopilotOrganizationDetails) GetSeatBreakdown() *CopilotSeatBreakdown {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotOrganizationDetails) GetSeatManagementSetting() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CopilotPeriodicMetrics) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CopilotPeriodicMetrics) GetDayTotals() []*CopilotDailyMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotPeriodicMetrics) GetEnterpriseID() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotPeriodicMetrics) GetOrganizationID() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotPeriodicMetrics) GetReportEndDay() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotPeriodicMetrics) GetReportStartDay() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotSeatBreakdown) GetActiveThisCycle() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotSeatBreakdown) GetAddedThisCycle() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotSeatBreakdown) GetInactiveThisCycle() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotSeatBreakdown) GetPendingCancellation() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotSeatBreakdown) GetPendingInvitation() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotSeatBreakdown) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotSeatDetails) GetAssignee() any { _ = "STUB: not implemented"; return *new(any) }

func (c *CopilotSeatDetails) GetAssigningTeam() *Team { _ = "STUB: not implemented"; return nil }

func (c *CopilotSeatDetails) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CopilotSeatDetails) GetLastActivityAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CopilotSeatDetails) GetLastActivityEditor() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotSeatDetails) GetPendingCancellationDate() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CopilotSeatDetails) GetPlanType() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotSeatDetails) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CopilotUserDailyMetrics) GetCodeAcceptanceActivityCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotUserDailyMetrics) GetCodeGenerationActivityCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotUserDailyMetrics) GetDay() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotUserDailyMetrics) GetEnterpriseID() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotUserDailyMetrics) GetLOCAddedSum() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotUserDailyMetrics) GetLOCDeletedSum() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotUserDailyMetrics) GetLOCSuggestedToAddSum() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotUserDailyMetrics) GetLOCSuggestedToDeleteSum() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotUserDailyMetrics) GetOrganizationID() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotUserDailyMetrics) GetTotalsByCLI() *CopilotMetricsCLI {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserDailyMetrics) GetTotalsByFeature() []*CopilotMetricsFeature {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserDailyMetrics) GetTotalsByIDE() []*CopilotUserMetricsIDE {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserDailyMetrics) GetTotalsByLanguageFeature() []*CopilotMetricsLanguageFeature {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserDailyMetrics) GetTotalsByLanguageModel() []*CopilotMetricsLanguageModel {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserDailyMetrics) GetTotalsByModelFeature() []*CopilotMetricsModelFeature {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserDailyMetrics) GetUsedAgent() bool { _ = "STUB: not implemented"; return false }

func (c *CopilotUserDailyMetrics) GetUsedChat() bool { _ = "STUB: not implemented"; return false }

func (c *CopilotUserDailyMetrics) GetUsedCLI() bool { _ = "STUB: not implemented"; return false }

func (c *CopilotUserDailyMetrics) GetUsedCopilotCodeReviewActive() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotUserDailyMetrics) GetUsedCopilotCodeReviewPassive() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotUserDailyMetrics) GetUsedCopilotCodingAgent() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotUserDailyMetrics) GetUserID() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotUserDailyMetrics) GetUserInitiatedInteractionCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotUserDailyMetrics) GetUserLogin() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotUserMetricsIDE) GetIDE() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotUserMetricsIDE) GetLastKnownIDEVersion() *CopilotUserMetricsIDEVersion {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserMetricsIDE) GetLastKnownPluginVersion() *CopilotUserMetricsPluginVersion {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserMetricsIDE) GetUserInitiatedInteractionCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotUserMetricsIDEVersion) GetIDEVersion() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotUserMetricsIDEVersion) GetSampledAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CopilotUserMetricsPluginVersion) GetPlugin() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotUserMetricsPluginVersion) GetPluginVersion() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CopilotUserMetricsPluginVersion) GetSampledAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CopilotUserPeriodicMetrics) GetCodeAcceptanceActivityCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotUserPeriodicMetrics) GetCodeGenerationActivityCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotUserPeriodicMetrics) GetDay() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotUserPeriodicMetrics) GetEnterpriseID() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotUserPeriodicMetrics) GetLOCAddedSum() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotUserPeriodicMetrics) GetLOCDeletedSum() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotUserPeriodicMetrics) GetLOCSuggestedToAddSum() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotUserPeriodicMetrics) GetLOCSuggestedToDeleteSum() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotUserPeriodicMetrics) GetOrganizationID() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CopilotUserPeriodicMetrics) GetReportEndDay() string { _ = "STUB: not implemented"; return "" }

func (c *CopilotUserPeriodicMetrics) GetReportStartDay() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CopilotUserPeriodicMetrics) GetTotalsByCLI() *CopilotMetricsCLI {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserPeriodicMetrics) GetTotalsByFeature() []*CopilotMetricsFeature {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserPeriodicMetrics) GetTotalsByIDE() []*CopilotUserMetricsIDE {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserPeriodicMetrics) GetTotalsByLanguageFeature() []*CopilotMetricsLanguageFeature {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserPeriodicMetrics) GetTotalsByLanguageModel() []*CopilotMetricsLanguageModel {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserPeriodicMetrics) GetTotalsByModelFeature() []*CopilotMetricsModelFeature {
	_ = "STUB: not implemented"
	return nil
}

func (c *CopilotUserPeriodicMetrics) GetUsedAgent() bool { _ = "STUB: not implemented"; return false }

func (c *CopilotUserPeriodicMetrics) GetUsedChat() bool { _ = "STUB: not implemented"; return false }

func (c *CopilotUserPeriodicMetrics) GetUsedCLI() bool { _ = "STUB: not implemented"; return false }

func (c *CopilotUserPeriodicMetrics) GetUsedCopilotCodeReviewActive() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotUserPeriodicMetrics) GetUsedCopilotCodeReviewPassive() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotUserPeriodicMetrics) GetUsedCopilotCodingAgent() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CopilotUserPeriodicMetrics) GetUserID() int { _ = "STUB: not implemented"; return 0 }

func (c *CopilotUserPeriodicMetrics) GetUserInitiatedInteractionCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CopilotUserPeriodicMetrics) GetUserLogin() string { _ = "STUB: not implemented"; return "" }

func (c *CostCenter) GetAzureSubscription() string { _ = "STUB: not implemented"; return "" }

func (c *CostCenter) GetID() string { _ = "STUB: not implemented"; return "" }

func (c *CostCenter) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CostCenter) GetResources() []*CostCenterResource { _ = "STUB: not implemented"; return nil }

func (c *CostCenter) GetState() string { _ = "STUB: not implemented"; return "" }

func (c *CostCenterRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CostCenterResource) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CostCenterResource) GetType() string { _ = "STUB: not implemented"; return "" }

func (c *CostCenterResourceRequest) GetOrganizations() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CostCenterResourceRequest) GetRepositories() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CostCenterResourceRequest) GetUsers() []string { _ = "STUB: not implemented"; return nil }

func (c *CostCenters) GetCostCenters() []*CostCenter { _ = "STUB: not implemented"; return nil }

func (c *CreateArtifactDeploymentRequest) GetCluster() string { _ = "STUB: not implemented"; return "" }

func (c *CreateArtifactDeploymentRequest) GetDeploymentName() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateArtifactDeploymentRequest) GetDigest() string { _ = "STUB: not implemented"; return "" }

func (c *CreateArtifactDeploymentRequest) GetGithubRepository() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateArtifactDeploymentRequest) GetLogicalEnvironment() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateArtifactDeploymentRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateArtifactDeploymentRequest) GetPhysicalEnvironment() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateArtifactDeploymentRequest) GetRuntimeRisks() []DeploymentRuntimeRisk {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateArtifactDeploymentRequest) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (c *CreateArtifactDeploymentRequest) GetTags() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateArtifactDeploymentRequest) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (c *CreateArtifactStorageRequest) GetArtifactURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateArtifactStorageRequest) GetDigest() string { _ = "STUB: not implemented"; return "" }

func (c *CreateArtifactStorageRequest) GetGithubRepository() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateArtifactStorageRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateArtifactStorageRequest) GetPath() string { _ = "STUB: not implemented"; return "" }

func (c *CreateArtifactStorageRequest) GetRegistryURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateArtifactStorageRequest) GetRepository() string { _ = "STUB: not implemented"; return "" }

func (c *CreateArtifactStorageRequest) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (c *CreateArtifactStorageRequest) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCheckRunOptions) GetActions() []*CheckRunAction {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateCheckRunOptions) GetCompletedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CreateCheckRunOptions) GetConclusion() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCheckRunOptions) GetDetailsURL() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCheckRunOptions) GetExternalID() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCheckRunOptions) GetHeadSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCheckRunOptions) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCheckRunOptions) GetOutput() *CheckRunOutput { _ = "STUB: not implemented"; return nil }

func (c *CreateCheckRunOptions) GetStartedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CreateCheckRunOptions) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCheckSuiteOptions) GetHeadBranch() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCheckSuiteOptions) GetHeadSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCodespaceOptions) GetClientIP() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCodespaceOptions) GetDevcontainerPath() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCodespaceOptions) GetDisplayName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCodespaceOptions) GetGeo() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCodespaceOptions) GetIdleTimeoutMinutes() int { _ = "STUB: not implemented"; return 0 }

func (c *CreateCodespaceOptions) GetLocation() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCodespaceOptions) GetMachine() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCodespaceOptions) GetMultiRepoPermissionsOptOut() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreateCodespaceOptions) GetRef() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCodespaceOptions) GetRetentionPeriodMinutes() int {
	_ = "STUB: not implemented"
	return 0
}

func (c *CreateCodespaceOptions) GetWorkingDirectory() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCommitOptions) GetSigner() MessageSigner {
	_ = "STUB: not implemented"
	return *new(MessageSigner)
}

func (c *CreateCustomOrgRoleRequest) GetBaseRole() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCustomOrgRoleRequest) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCustomOrgRoleRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateCustomOrgRoleRequest) GetPermissions() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateDeploymentBranchPolicyRequest) GetName() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateDeploymentBranchPolicyRequest) GetType() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateEnterpriseRunnerGroupRequest) GetAllowsPublicRepositories() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreateEnterpriseRunnerGroupRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateEnterpriseRunnerGroupRequest) GetNetworkConfigurationID() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateEnterpriseRunnerGroupRequest) GetRestrictedToWorkflows() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreateEnterpriseRunnerGroupRequest) GetRunners() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateEnterpriseRunnerGroupRequest) GetSelectedOrganizationIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateEnterpriseRunnerGroupRequest) GetSelectedWorkflows() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateEnterpriseRunnerGroupRequest) GetVisibility() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateEvent) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (c *CreateEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (c *CreateEvent) GetMasterBranch() string { _ = "STUB: not implemented"; return "" }

func (c *CreateEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (c *CreateEvent) GetPusherType() string { _ = "STUB: not implemented"; return "" }

func (c *CreateEvent) GetRef() string { _ = "STUB: not implemented"; return "" }

func (c *CreateEvent) GetRefType() string { _ = "STUB: not implemented"; return "" }

func (c *CreateEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (c *CreateEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (c *CreateGistCommentRequest) GetBody() string { _ = "STUB: not implemented"; return "" }

func (c *CreateGistFile) GetContent() string { _ = "STUB: not implemented"; return "" }

func (c *CreateGistRequest) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (c *CreateGistRequest) GetPublic() bool { _ = "STUB: not implemented"; return false }

func (c *CreateHostedRunnerRequest) GetEnableStaticIP() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreateHostedRunnerRequest) GetImage() HostedRunnerImage {
	_ = "STUB: not implemented"
	return *new(HostedRunnerImage)
}

func (c *CreateHostedRunnerRequest) GetImageGen() bool { _ = "STUB: not implemented"; return false }

func (c *CreateHostedRunnerRequest) GetMaximumRunners() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CreateHostedRunnerRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateHostedRunnerRequest) GetRunnerGroupID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CreateHostedRunnerRequest) GetSize() string { _ = "STUB: not implemented"; return "" }

func (c *CreateJITConfigRequest) GetLabels() []string { _ = "STUB: not implemented"; return nil }

func (c *CreateJITConfigRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateJITConfigRequest) GetRunnerGroupID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CreateJITConfigRequest) GetWorkFolder() string { _ = "STUB: not implemented"; return "" }

func (c *CreateOrganizationPrivateRegistry) GetAccountID() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetAudience() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetAuthType() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetAWSRegion() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetClientID() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetDomain() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetDomainOwner() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetEncryptedValue() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetIdentityMappingName() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetJFrogOIDCProviderName() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (c *CreateOrganizationPrivateRegistry) GetRegistryType() PrivateRegistryType {
	_ = "STUB: not implemented"
	return *new(PrivateRegistryType)
}

func (c *CreateOrganizationPrivateRegistry) GetReplacesBase() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreateOrganizationPrivateRegistry) GetRoleName() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetSelectedRepositoryIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateOrganizationPrivateRegistry) GetTenantID() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *CreateOrganizationPrivateRegistry) GetUsername() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrganizationPrivateRegistry) GetVisibility() PrivateRegistryVisibility {
	_ = "STUB: not implemented"
	return *new(PrivateRegistryVisibility)
}

func (c *CreateOrgInvitationOptions) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (c *CreateOrgInvitationOptions) GetInviteeID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CreateOrgInvitationOptions) GetRole() string { _ = "STUB: not implemented"; return "" }

func (c *CreateOrgInvitationOptions) GetTeamID() []int64 { _ = "STUB: not implemented"; return nil }

func (c *CreateOrUpdateCustomRepoRoleOptions) GetBaseRole() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrUpdateCustomRepoRoleOptions) GetDescription() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrUpdateCustomRepoRoleOptions) GetName() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrUpdateCustomRepoRoleOptions) GetPermissions() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateOrUpdateIssueTypesOptions) GetColor() string { _ = "STUB: not implemented"; return "" }

func (c *CreateOrUpdateIssueTypesOptions) GetDescription() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateOrUpdateIssueTypesOptions) GetIsEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreateOrUpdateIssueTypesOptions) GetIsPrivate() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreateOrUpdateIssueTypesOptions) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateProjectV2DraftItemRequest) GetBody() string { _ = "STUB: not implemented"; return "" }

func (c *CreateProjectV2DraftItemRequest) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (c *CreateProjectV2ViewRequest) GetFilter() string { _ = "STUB: not implemented"; return "" }

func (c *CreateProjectV2ViewRequest) GetLayout() string { _ = "STUB: not implemented"; return "" }

func (c *CreateProjectV2ViewRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateProjectV2ViewRequest) GetVisibleFields() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateProtectedChanges) GetFrom() bool { _ = "STUB: not implemented"; return false }

func (c *CreateRef) GetRef() string { _ = "STUB: not implemented"; return "" }

func (c *CreateRef) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (c *CreateReleaseRequest) GetBody() string { _ = "STUB: not implemented"; return "" }

func (c *CreateReleaseRequest) GetDiscussionCategoryName() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateReleaseRequest) GetDraft() bool { _ = "STUB: not implemented"; return false }

func (c *CreateReleaseRequest) GetGenerateReleaseNotes() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreateReleaseRequest) GetMakeLatest() string { _ = "STUB: not implemented"; return "" }

func (c *CreateReleaseRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateReleaseRequest) GetPrerelease() bool { _ = "STUB: not implemented"; return false }

func (c *CreateReleaseRequest) GetTagName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateReleaseRequest) GetTargetCommitish() string { _ = "STUB: not implemented"; return "" }

func (c *CreateRunnerGroupRequest) GetAllowsPublicRepositories() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreateRunnerGroupRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CreateRunnerGroupRequest) GetNetworkConfigurationID() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CreateRunnerGroupRequest) GetRestrictedToWorkflows() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreateRunnerGroupRequest) GetRunners() []int64 { _ = "STUB: not implemented"; return nil }

func (c *CreateRunnerGroupRequest) GetSelectedRepositoryIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateRunnerGroupRequest) GetSelectedWorkflows() []string {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateRunnerGroupRequest) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (c *CreateTag) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (c *CreateTag) GetObject() string { _ = "STUB: not implemented"; return "" }

func (c *CreateTag) GetTag() string { _ = "STUB: not implemented"; return "" }

func (c *CreateTag) GetTagger() *CommitAuthor { _ = "STUB: not implemented"; return nil }

func (c *CreateTag) GetType() string { _ = "STUB: not implemented"; return "" }

func (c *CreateUpdateEnvironment) GetCanAdminsBypass() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreateUpdateEnvironment) GetDeploymentBranchPolicy() *BranchPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateUpdateEnvironment) GetPreventSelfReview() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreateUpdateEnvironment) GetReviewers() []*EnvReviewers {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateUpdateEnvironment) GetWaitTimer() int { _ = "STUB: not implemented"; return 0 }

func (c *CreateUserRequest) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (c *CreateUserRequest) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (c *CreateUserRequest) GetSuspended() bool { _ = "STUB: not implemented"; return false }

func (c *CreateWorkflowDispatchEventRequest) GetInputs() map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (c *CreateWorkflowDispatchEventRequest) GetRef() string { _ = "STUB: not implemented"; return "" }

func (c *CreateWorkflowDispatchEventRequest) GetReturnRunDetails() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CreationInfo) GetCreated() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (c *CreationInfo) GetCreators() []string { _ = "STUB: not implemented"; return nil }

func (c *CredentialAuthorization) GetAuthorizedCredentialExpiresAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CredentialAuthorization) GetAuthorizedCredentialID() int64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *CredentialAuthorization) GetAuthorizedCredentialNote() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CredentialAuthorization) GetAuthorizedCredentialTitle() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CredentialAuthorization) GetCredentialAccessedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CredentialAuthorization) GetCredentialAuthorizedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CredentialAuthorization) GetCredentialID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CredentialAuthorization) GetCredentialType() string { _ = "STUB: not implemented"; return "" }

func (c *CredentialAuthorization) GetFingerprint() string { _ = "STUB: not implemented"; return "" }

func (c *CredentialAuthorization) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (c *CredentialAuthorization) GetScopes() []string { _ = "STUB: not implemented"; return nil }

func (c *CredentialAuthorization) GetTokenLastEight() string { _ = "STUB: not implemented"; return "" }

func (c *CredentialAuthorizationsListOptions) GetLogin() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *Credit) GetType() string { _ = "STUB: not implemented"; return "" }

func (c *Credit) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (c *CustomDeploymentProtectionRule) GetApp() *CustomDeploymentProtectionRuleApp {
	_ = "STUB: not implemented"
	return nil
}

func (c *CustomDeploymentProtectionRule) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (c *CustomDeploymentProtectionRule) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CustomDeploymentProtectionRule) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (c *CustomDeploymentProtectionRuleApp) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CustomDeploymentProtectionRuleApp) GetIntegrationURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CustomDeploymentProtectionRuleApp) GetNodeID() string {
	_ = "STUB: not implemented"
	return ""
}

func (c *CustomDeploymentProtectionRuleApp) GetSlug() string { _ = "STUB: not implemented"; return "" }

func (c *CustomDeploymentProtectionRuleRequest) GetIntegrationID() int64 {
	_ = "STUB: not implemented"
	return 0
}

func (c *CustomOrgRole) GetBaseRole() string { _ = "STUB: not implemented"; return "" }

func (c *CustomOrgRole) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (c *CustomOrgRole) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (c *CustomOrgRole) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CustomOrgRole) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CustomOrgRole) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (c *CustomOrgRole) GetPermissions() []string { _ = "STUB: not implemented"; return nil }

func (c *CustomOrgRole) GetSource() string { _ = "STUB: not implemented"; return "" }

func (c *CustomOrgRole) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (c *CustomPatternBackfillScan) GetPatternScope() string { _ = "STUB: not implemented"; return "" }

func (c *CustomPatternBackfillScan) GetPatternSlug() string { _ = "STUB: not implemented"; return "" }

func (c *CustomProperty) GetAllowedValues() []string { _ = "STUB: not implemented"; return nil }

func (c *CustomProperty) GetDefaultValue() any { _ = "STUB: not implemented"; return *new(any) }

func (c *CustomProperty) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (c *CustomProperty) GetPropertyName() string { _ = "STUB: not implemented"; return "" }

func (c *CustomProperty) GetRequired() bool { _ = "STUB: not implemented"; return false }

func (c *CustomProperty) GetSourceType() string { _ = "STUB: not implemented"; return "" }

func (c *CustomProperty) GetURL() string { _ = "STUB: not implemented"; return "" }

func (c *CustomProperty) GetValuesEditableBy() string { _ = "STUB: not implemented"; return "" }

func (c *CustomProperty) GetValueType() PropertyValueType {
	_ = "STUB: not implemented"
	return *new(PropertyValueType)
}

func (c *CustomPropertyEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (c *CustomPropertyEvent) GetDefinition() *CustomProperty {
	_ = "STUB: not implemented"
	return nil
}

func (c *CustomPropertyEvent) GetEnterprise() *Enterprise { _ = "STUB: not implemented"; return nil }

func (c *CustomPropertyEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (c *CustomPropertyEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (c *CustomPropertyEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (c *CustomPropertyValue) GetPropertyName() string { _ = "STUB: not implemented"; return "" }

func (c *CustomPropertyValue) GetValue() any { _ = "STUB: not implemented"; return *new(any) }

func (c *CustomPropertyValuesEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (c *CustomPropertyValuesEvent) GetEnterprise() *Enterprise {
	_ = "STUB: not implemented"
	return nil
}

func (c *CustomPropertyValuesEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (c *CustomPropertyValuesEvent) GetNewPropertyValues() []*CustomPropertyValue {
	_ = "STUB: not implemented"
	return nil
}

func (c *CustomPropertyValuesEvent) GetOldPropertyValues() []*CustomPropertyValue {
	_ = "STUB: not implemented"
	return nil
}

func (c *CustomPropertyValuesEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (c *CustomPropertyValuesEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (c *CustomPropertyValuesEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (c *CustomRepoRoles) GetBaseRole() string { _ = "STUB: not implemented"; return "" }

func (c *CustomRepoRoles) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (c *CustomRepoRoles) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (c *CustomRepoRoles) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (c *CustomRepoRoles) GetName() string { _ = "STUB: not implemented"; return "" }

func (c *CustomRepoRoles) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (c *CustomRepoRoles) GetPermissions() []string { _ = "STUB: not implemented"; return nil }

func (c *CustomRepoRoles) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DatadogConfig) GetEncryptedToken() string { _ = "STUB: not implemented"; return "" }

func (d *DatadogConfig) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (d *DatadogConfig) GetSite() string { _ = "STUB: not implemented"; return "" }

func (d *DefaultSetupConfiguration) GetLanguages() []string { _ = "STUB: not implemented"; return nil }

func (d *DefaultSetupConfiguration) GetQuerySuite() string { _ = "STUB: not implemented"; return "" }

func (d *DefaultSetupConfiguration) GetState() string { _ = "STUB: not implemented"; return "" }

func (d *DefaultSetupConfiguration) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DefaultWorkflowPermissionEnterprise) GetCanApprovePullRequestReviews() bool {
	_ = "STUB: not implemented"
	return false
}

func (d *DefaultWorkflowPermissionEnterprise) GetDefaultWorkflowPermissions() string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DefaultWorkflowPermissionOrganization) GetCanApprovePullRequestReviews() bool {
	_ = "STUB: not implemented"
	return false
}

func (d *DefaultWorkflowPermissionOrganization) GetDefaultWorkflowPermissions() string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DefaultWorkflowPermissionRepository) GetCanApprovePullRequestReviews() bool {
	_ = "STUB: not implemented"
	return false
}

func (d *DefaultWorkflowPermissionRepository) GetDefaultWorkflowPermissions() string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DeleteAnalysis) GetConfirmDeleteURL() string { _ = "STUB: not implemented"; return "" }

func (d *DeleteAnalysis) GetNextAnalysisURL() string { _ = "STUB: not implemented"; return "" }

func (d *DeleteCostCenterResponse) GetCostCenterState() string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DeleteCostCenterResponse) GetID() string { _ = "STUB: not implemented"; return "" }

func (d *DeleteCostCenterResponse) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (d *DeleteCostCenterResponse) GetName() string { _ = "STUB: not implemented"; return "" }

func (d *DeleteEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (d *DeleteEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (d *DeleteEvent) GetPusherType() string { _ = "STUB: not implemented"; return "" }

func (d *DeleteEvent) GetRef() string { _ = "STUB: not implemented"; return "" }

func (d *DeleteEvent) GetRefType() string { _ = "STUB: not implemented"; return "" }

func (d *DeleteEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (d *DeleteEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (d *DependabotAlert) GetAutoDismissedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DependabotAlert) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DependabotAlert) GetDependency() *Dependency { _ = "STUB: not implemented"; return nil }

func (d *DependabotAlert) GetDismissedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DependabotAlert) GetDismissedBy() *User { _ = "STUB: not implemented"; return nil }

func (d *DependabotAlert) GetDismissedComment() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotAlert) GetDismissedReason() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotAlert) GetFixedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (d *DependabotAlert) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotAlert) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (d *DependabotAlert) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (d *DependabotAlert) GetSecurityAdvisory() *DependabotSecurityAdvisory {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependabotAlert) GetSecurityVulnerability() *AdvisoryVulnerability {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependabotAlert) GetState() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotAlert) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DependabotAlert) GetURL() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotAlertEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotAlertEvent) GetAlert() *DependabotAlert { _ = "STUB: not implemented"; return nil }

func (d *DependabotAlertEvent) GetEnterprise() *Enterprise { _ = "STUB: not implemented"; return nil }

func (d *DependabotAlertEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependabotAlertEvent) GetOrganization() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependabotAlertEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (d *DependabotAlertEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (d *DependabotAlertState) GetDismissedComment() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotAlertState) GetDismissedReason() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotAlertState) GetState() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotSecurityAdvisory) GetCVEID() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotSecurityAdvisory) GetCVSS() *AdvisoryCVSS { _ = "STUB: not implemented"; return nil }

func (d *DependabotSecurityAdvisory) GetCWEs() []*AdvisoryCWEs {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependabotSecurityAdvisory) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotSecurityAdvisory) GetEPSS() *AdvisoryEPSS { _ = "STUB: not implemented"; return nil }

func (d *DependabotSecurityAdvisory) GetGHSAID() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotSecurityAdvisory) GetIdentifiers() []*AdvisoryIdentifier {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependabotSecurityAdvisory) GetPublishedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DependabotSecurityAdvisory) GetReferences() []*AdvisoryReference {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependabotSecurityAdvisory) GetSeverity() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotSecurityAdvisory) GetSummary() string { _ = "STUB: not implemented"; return "" }

func (d *DependabotSecurityAdvisory) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DependabotSecurityAdvisory) GetVulnerabilities() []*AdvisoryVulnerability {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependabotSecurityAdvisory) GetWithdrawnAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DependabotSecurityUpdates) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (d *Dependency) GetManifestPath() string { _ = "STUB: not implemented"; return "" }

func (d *Dependency) GetPackage() *VulnerabilityPackage { _ = "STUB: not implemented"; return nil }

func (d *Dependency) GetScope() string { _ = "STUB: not implemented"; return "" }

func (d *DependencyGraphAutosubmitActionOptions) GetLabeledRunners() bool {
	_ = "STUB: not implemented"
	return false
}

func (d *DependencyGraphSnapshot) GetDetector() *DependencyGraphSnapshotDetector {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependencyGraphSnapshot) GetJob() *DependencyGraphSnapshotJob {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependencyGraphSnapshot) GetMetadata() map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependencyGraphSnapshot) GetRef() string { _ = "STUB: not implemented"; return "" }

func (d *DependencyGraphSnapshot) GetScanned() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DependencyGraphSnapshot) GetSha() string { _ = "STUB: not implemented"; return "" }

func (d *DependencyGraphSnapshot) GetVersion() int { _ = "STUB: not implemented"; return 0 }

func (d *DependencyGraphSnapshotCreationData) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DependencyGraphSnapshotCreationData) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (d *DependencyGraphSnapshotCreationData) GetMessage() string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DependencyGraphSnapshotCreationData) GetResult() string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DependencyGraphSnapshotDetector) GetName() string { _ = "STUB: not implemented"; return "" }

func (d *DependencyGraphSnapshotDetector) GetURL() string { _ = "STUB: not implemented"; return "" }

func (d *DependencyGraphSnapshotDetector) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (d *DependencyGraphSnapshotJob) GetCorrelator() string { _ = "STUB: not implemented"; return "" }

func (d *DependencyGraphSnapshotJob) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (d *DependencyGraphSnapshotJob) GetID() string { _ = "STUB: not implemented"; return "" }

func (d *DependencyGraphSnapshotManifest) GetFile() *DependencyGraphSnapshotManifestFile {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependencyGraphSnapshotManifest) GetMetadata() map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependencyGraphSnapshotManifest) GetName() string { _ = "STUB: not implemented"; return "" }

func (d *DependencyGraphSnapshotManifestFile) GetSourceLocation() string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DependencyGraphSnapshotResolvedDependency) GetDependencies() []string {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependencyGraphSnapshotResolvedDependency) GetMetadata() map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (d *DependencyGraphSnapshotResolvedDependency) GetPackageURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DependencyGraphSnapshotResolvedDependency) GetRelationship() string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DependencyGraphSnapshotResolvedDependency) GetScope() string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DeployKeyEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (d *DeployKeyEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (d *DeployKeyEvent) GetKey() *Key { _ = "STUB: not implemented"; return nil }

func (d *DeployKeyEvent) GetOrganization() *Organization { _ = "STUB: not implemented"; return nil }

func (d *DeployKeyEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (d *DeployKeyEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (d *Deployment) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (d *Deployment) GetCreator() *User { _ = "STUB: not implemented"; return nil }

func (d *Deployment) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (d *Deployment) GetEnvironment() string { _ = "STUB: not implemented"; return "" }

func (d *Deployment) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (d *Deployment) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (d *Deployment) GetPayload() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (d *Deployment) GetRef() string { _ = "STUB: not implemented"; return "" }

func (d *Deployment) GetRepositoryURL() string { _ = "STUB: not implemented"; return "" }

func (d *Deployment) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (d *Deployment) GetStatusesURL() string { _ = "STUB: not implemented"; return "" }

func (d *Deployment) GetTask() string { _ = "STUB: not implemented"; return "" }

func (d *Deployment) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (d *Deployment) GetURL() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentBranchPolicy) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (d *DeploymentBranchPolicy) GetName() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentBranchPolicy) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentBranchPolicy) GetType() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentBranchPolicyResponse) GetBranchPolicies() []*DeploymentBranchPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentBranchPolicyResponse) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (d *DeploymentEvent) GetDeployment() *Deployment { _ = "STUB: not implemented"; return nil }

func (d *DeploymentEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (d *DeploymentEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (d *DeploymentEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (d *DeploymentEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (d *DeploymentEvent) GetWorkflow() *Workflow { _ = "STUB: not implemented"; return nil }

func (d *DeploymentEvent) GetWorkflowRun() *WorkflowRun { _ = "STUB: not implemented"; return nil }

func (d *DeploymentProtectionRuleEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentProtectionRuleEvent) GetDeployment() *Deployment {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentProtectionRuleEvent) GetDeploymentCallbackURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DeploymentProtectionRuleEvent) GetEnvironment() string {
	_ = "STUB: not implemented"
	return ""
}

func (d *DeploymentProtectionRuleEvent) GetEvent() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentProtectionRuleEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentProtectionRuleEvent) GetOrganization() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentProtectionRuleEvent) GetPullRequests() []*PullRequest {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentProtectionRuleEvent) GetRepo() *Repository {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentProtectionRuleEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (d *DeploymentRequest) GetAutoMerge() bool { _ = "STUB: not implemented"; return false }

func (d *DeploymentRequest) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentRequest) GetEnvironment() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentRequest) GetPayload() any { _ = "STUB: not implemented"; return *new(any) }

func (d *DeploymentRequest) GetProductionEnvironment() bool {
	_ = "STUB: not implemented"
	return false
}

func (d *DeploymentRequest) GetRef() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentRequest) GetRequiredContexts() []string { _ = "STUB: not implemented"; return nil }

func (d *DeploymentRequest) GetTask() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentRequest) GetTransientEnvironment() bool { _ = "STUB: not implemented"; return false }

func (d *DeploymentReviewEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentReviewEvent) GetApprover() *User { _ = "STUB: not implemented"; return nil }

func (d *DeploymentReviewEvent) GetComment() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentReviewEvent) GetEnterprise() *Enterprise { _ = "STUB: not implemented"; return nil }

func (d *DeploymentReviewEvent) GetEnvironment() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentReviewEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentReviewEvent) GetOrganization() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentReviewEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (d *DeploymentReviewEvent) GetRequester() *User { _ = "STUB: not implemented"; return nil }

func (d *DeploymentReviewEvent) GetReviewers() []*RequiredReviewer {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentReviewEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (d *DeploymentReviewEvent) GetSince() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentReviewEvent) GetWorkflowJobRun() *WorkflowJobRun {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentReviewEvent) GetWorkflowJobRuns() []*WorkflowJobRun {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentReviewEvent) GetWorkflowRun() *WorkflowRun {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentsListOptions) GetEnvironment() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentsListOptions) GetRef() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentsListOptions) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentsListOptions) GetTask() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatus) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DeploymentStatus) GetCreator() *User { _ = "STUB: not implemented"; return nil }

func (d *DeploymentStatus) GetDeploymentURL() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatus) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatus) GetEnvironment() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatus) GetEnvironmentURL() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatus) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (d *DeploymentStatus) GetLogURL() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatus) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatus) GetRepositoryURL() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatus) GetState() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatus) GetTargetURL() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatus) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DeploymentStatus) GetURL() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatusEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatusEvent) GetDeployment() *Deployment { _ = "STUB: not implemented"; return nil }

func (d *DeploymentStatusEvent) GetDeploymentStatus() *DeploymentStatus {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentStatusEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (d *DeploymentStatusEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (d *DeploymentStatusEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (d *DeploymentStatusEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (d *DeploymentStatusRequest) GetAutoInactive() bool { _ = "STUB: not implemented"; return false }

func (d *DeploymentStatusRequest) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatusRequest) GetEnvironment() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatusRequest) GetEnvironmentURL() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatusRequest) GetLogURL() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatusRequest) GetState() string { _ = "STUB: not implemented"; return "" }

func (d *DeploymentStatusRequest) GetTargetURL() string { _ = "STUB: not implemented"; return "" }

func (d *DevContainer) GetDisplayName() string { _ = "STUB: not implemented"; return "" }

func (d *DevContainer) GetName() string { _ = "STUB: not implemented"; return "" }

func (d *DevContainer) GetPath() string { _ = "STUB: not implemented"; return "" }

func (d *DevContainerConfigurations) GetDevcontainers() []*DevContainer {
	_ = "STUB: not implemented"
	return nil
}

func (d *DevContainerConfigurations) GetTotalCount() int64 { _ = "STUB: not implemented"; return 0 }

func (d *Discussion) GetActiveLockReason() string { _ = "STUB: not implemented"; return "" }

func (d *Discussion) GetAnswerChosenAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *Discussion) GetAnswerChosenBy() string { _ = "STUB: not implemented"; return "" }

func (d *Discussion) GetAnswerHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (d *Discussion) GetAuthorAssociation() string { _ = "STUB: not implemented"; return "" }

func (d *Discussion) GetBody() string { _ = "STUB: not implemented"; return "" }

func (d *Discussion) GetComments() int { _ = "STUB: not implemented"; return 0 }

func (d *Discussion) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (d *Discussion) GetDiscussionCategory() *DiscussionCategory {
	_ = "STUB: not implemented"
	return nil
}

func (d *Discussion) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (d *Discussion) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (d *Discussion) GetLocked() bool { _ = "STUB: not implemented"; return false }

func (d *Discussion) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (d *Discussion) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (d *Discussion) GetRepositoryURL() string { _ = "STUB: not implemented"; return "" }

func (d *Discussion) GetState() string { _ = "STUB: not implemented"; return "" }

func (d *Discussion) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (d *Discussion) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (d *Discussion) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (d *DiscussionCategory) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DiscussionCategory) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionCategory) GetEmoji() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionCategory) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (d *DiscussionCategory) GetIsAnswerable() bool { _ = "STUB: not implemented"; return false }

func (d *DiscussionCategory) GetName() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionCategory) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionCategory) GetRepositoryID() int64 { _ = "STUB: not implemented"; return 0 }

func (d *DiscussionCategory) GetSlug() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionCategory) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DiscussionComment) GetAuthor() *User { _ = "STUB: not implemented"; return nil }

func (d *DiscussionComment) GetBody() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionComment) GetBodyHTML() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionComment) GetBodyVersion() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionComment) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DiscussionComment) GetDiscussionURL() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionComment) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionComment) GetLastEditedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DiscussionComment) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionComment) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (d *DiscussionComment) GetReactions() *Reactions { _ = "STUB: not implemented"; return nil }

func (d *DiscussionComment) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (d *DiscussionComment) GetURL() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionCommentEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionCommentEvent) GetComment() *CommentDiscussion {
	_ = "STUB: not implemented"
	return nil
}

func (d *DiscussionCommentEvent) GetDiscussion() *Discussion { _ = "STUB: not implemented"; return nil }

func (d *DiscussionCommentEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (d *DiscussionCommentEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (d *DiscussionCommentEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (d *DiscussionCommentEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (d *DiscussionCommentListOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (d *DiscussionEvent) GetDiscussion() *Discussion { _ = "STUB: not implemented"; return nil }

func (d *DiscussionEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (d *DiscussionEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (d *DiscussionEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (d *DiscussionEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (d *DiscussionListOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (d *DismissalRestrictions) GetApps() []*App { _ = "STUB: not implemented"; return nil }

func (d *DismissalRestrictions) GetTeams() []*Team { _ = "STUB: not implemented"; return nil }

func (d *DismissalRestrictions) GetUsers() []*User { _ = "STUB: not implemented"; return nil }

func (d *DismissalRestrictionsRequest) GetApps() []string { _ = "STUB: not implemented"; return nil }

func (d *DismissalRestrictionsRequest) GetTeams() []string { _ = "STUB: not implemented"; return nil }

func (d *DismissalRestrictionsRequest) GetUsers() []string { _ = "STUB: not implemented"; return nil }

func (d *DismissedReview) GetDismissalCommitID() string { _ = "STUB: not implemented"; return "" }

func (d *DismissedReview) GetDismissalMessage() string { _ = "STUB: not implemented"; return "" }

func (d *DismissedReview) GetReviewID() int64 { _ = "STUB: not implemented"; return 0 }

func (d *DismissedReview) GetState() string { _ = "STUB: not implemented"; return "" }

func (d *DismissStaleReviewsOnPushChanges) GetFrom() bool { _ = "STUB: not implemented"; return false }

func (d *DispatchRequestOptions) GetClientPayload() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (d *DispatchRequestOptions) GetEventType() string { _ = "STUB: not implemented"; return "" }

func (d *DraftReviewComment) GetBody() string { _ = "STUB: not implemented"; return "" }

func (d *DraftReviewComment) GetLine() int { _ = "STUB: not implemented"; return 0 }

func (d *DraftReviewComment) GetPath() string { _ = "STUB: not implemented"; return "" }

func (d *DraftReviewComment) GetPosition() int { _ = "STUB: not implemented"; return 0 }

func (d *DraftReviewComment) GetSide() string { _ = "STUB: not implemented"; return "" }

func (d *DraftReviewComment) GetStartLine() int { _ = "STUB: not implemented"; return 0 }

func (d *DraftReviewComment) GetStartSide() string { _ = "STUB: not implemented"; return "" }

func (e *EditBase) GetRef() *EditRef { _ = "STUB: not implemented"; return nil }

func (e *EditBase) GetSHA() *EditSHA { _ = "STUB: not implemented"; return nil }

func (e *EditBody) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (e *EditChange) GetBase() *EditBase { _ = "STUB: not implemented"; return nil }

func (e *EditChange) GetBody() *EditBody { _ = "STUB: not implemented"; return nil }

func (e *EditChange) GetDefaultBranch() *EditDefaultBranch { _ = "STUB: not implemented"; return nil }

func (e *EditChange) GetOwner() *EditOwner { _ = "STUB: not implemented"; return nil }

func (e *EditChange) GetRepo() *EditRepo { _ = "STUB: not implemented"; return nil }

func (e *EditChange) GetTitle() *EditTitle { _ = "STUB: not implemented"; return nil }

func (e *EditChange) GetTopics() *EditTopics { _ = "STUB: not implemented"; return nil }

func (e *EditDefaultBranch) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (e *EditOwner) GetOwnerInfo() *OwnerInfo { _ = "STUB: not implemented"; return nil }

func (e *EditRef) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (e *EditRepo) GetName() *RepoName { _ = "STUB: not implemented"; return nil }

func (e *EditSHA) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (e *EditTitle) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (e *EditTopics) GetFrom() []string { _ = "STUB: not implemented"; return nil }

func (e *EncryptedSecret) GetEncryptedValue() string { _ = "STUB: not implemented"; return "" }

func (e *EncryptedSecret) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (e *EncryptedSecret) GetName() string { _ = "STUB: not implemented"; return "" }

func (e *EncryptedSecret) GetSelectedRepositoryIDs() SelectedRepoIDs {
	_ = "STUB: not implemented"
	return *new(SelectedRepoIDs)
}

func (e *EncryptedSecret) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (e *Enterprise) GetAvatarURL() string { _ = "STUB: not implemented"; return "" }

func (e *Enterprise) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (e *Enterprise) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (e *Enterprise) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (e *Enterprise) GetID() int { _ = "STUB: not implemented"; return 0 }

func (e *Enterprise) GetName() string { _ = "STUB: not implemented"; return "" }

func (e *Enterprise) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (e *Enterprise) GetSlug() string { _ = "STUB: not implemented"; return "" }

func (e *Enterprise) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (e *Enterprise) GetWebsiteURL() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseAggregatedUsageItem) GetDiscountAmount() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *EnterpriseAggregatedUsageItem) GetDiscountQuantity() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *EnterpriseAggregatedUsageItem) GetGrossAmount() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *EnterpriseAggregatedUsageItem) GetGrossQuantity() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *EnterpriseAggregatedUsageItem) GetModel() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseAggregatedUsageItem) GetNetAmount() float64 { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseAggregatedUsageItem) GetNetQuantity() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *EnterpriseAggregatedUsageItem) GetPricePerUnit() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *EnterpriseAggregatedUsageItem) GetProduct() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseAggregatedUsageItem) GetSKU() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseAggregatedUsageItem) GetUnitType() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseAggregatedUsageReport) GetCostCenter() *BillingCostCenter {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseAggregatedUsageReport) GetEnterprise() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseAggregatedUsageReport) GetModel() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseAggregatedUsageReport) GetOrganization() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseAggregatedUsageReport) GetProduct() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseAggregatedUsageReport) GetTimePeriod() EnterpriseUsageTimePeriod {
	_ = "STUB: not implemented"
	return *new(EnterpriseUsageTimePeriod)
}

func (e *EnterpriseAggregatedUsageReport) GetUsageItems() []*EnterpriseAggregatedUsageItem {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseAggregatedUsageReport) GetUser() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseBudget) GetBudgetAlerting() *EnterpriseBudgetAlerting {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseBudget) GetBudgetAmount() int { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseBudget) GetBudgetEntityName() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseBudget) GetBudgetProductSKU() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseBudget) GetBudgetScope() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseBudget) GetBudgetType() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseBudget) GetID() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseBudget) GetPreventFurtherUsage() bool { _ = "STUB: not implemented"; return false }

func (e *EnterpriseBudgetAlerting) GetAlertRecipients() []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseBudgetAlerting) GetWillAlert() bool { _ = "STUB: not implemented"; return false }

func (e *EnterpriseConsumedLicenses) GetTotalSeatsConsumed() int {
	_ = "STUB: not implemented"
	return 0
}

func (e *EnterpriseConsumedLicenses) GetTotalSeatsPurchased() int {
	_ = "STUB: not implemented"
	return 0
}

func (e *EnterpriseConsumedLicenses) GetUsers() []*EnterpriseLicensedUsers {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseCreateBudget) GetBudgetAlerting() *EnterpriseBudgetAlerting {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseCreateBudget) GetBudgetAmount() int { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseCreateBudget) GetBudgetEntityName() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseCreateBudget) GetBudgetProductSKU() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseCreateBudget) GetBudgetScope() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseCreateBudget) GetBudgetType() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseCreateBudget) GetPreventFurtherUsage() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *EnterpriseCreateOrUpdateBudgetResponse) GetBudget() *EnterpriseBudget {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseCreateOrUpdateBudgetResponse) GetMessage() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseCustomPropertiesValues) GetOrganizationID() int64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *EnterpriseCustomPropertiesValues) GetOrganizationLogin() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseCustomPropertiesValues) GetProperties() []*CustomPropertyValue {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseCustomPropertySchema) GetProperties() []*CustomProperty {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseCustomPropertyValuesRequest) GetOrganizationLogin() []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseCustomPropertyValuesRequest) GetProperties() []*CustomPropertyValue {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseDeleteBudgetResponse) GetID() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseDeleteBudgetResponse) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseLicensedUsers) GetEnterpriseServerEmails() []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseLicensedUsers) GetEnterpriseServerUser() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *EnterpriseLicensedUsers) GetEnterpriseServerUserIDs() []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseLicensedUsers) GetGithubComEnterpriseRoles() []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseLicensedUsers) GetGithubComLogin() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseLicensedUsers) GetGithubComMemberRoles() []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseLicensedUsers) GetGithubComName() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseLicensedUsers) GetGithubComOrgsWithPendingInvites() []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseLicensedUsers) GetGithubComProfile() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseLicensedUsers) GetGithubComSamlNameID() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseLicensedUsers) GetGithubComTwoFactorAuth() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *EnterpriseLicensedUsers) GetGithubComUser() bool { _ = "STUB: not implemented"; return false }

func (e *EnterpriseLicensedUsers) GetGithubComVerifiedDomainEmails() []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseLicensedUsers) GetLicenseType() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseLicensedUsers) GetTotalUserAccounts() int { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseLicensedUsers) GetVisualStudioLicenseStatus() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseLicensedUsers) GetVisualStudioSubscriptionEmail() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseLicensedUsers) GetVisualStudioSubscriptionUser() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *EnterpriseLicenseSyncStatus) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseLicenseSyncStatus) GetProperties() *ServerInstanceProperties {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseLicenseSyncStatus) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseListBudgets) GetBudgets() []*EnterpriseBudget {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseListBudgets) GetHasNextPage() bool { _ = "STUB: not implemented"; return false }

func (e *EnterpriseListBudgets) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (e *EnterprisePremiumRequestUsageReportOptions) GetModel() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterprisePremiumRequestUsageReportOptions) GetOrganization() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterprisePremiumRequestUsageReportOptions) GetProduct() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterprisePremiumRequestUsageReportOptions) GetUser() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseRunnerGroup) GetAllowsPublicRepositories() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *EnterpriseRunnerGroup) GetDefault() bool { _ = "STUB: not implemented"; return false }

func (e *EnterpriseRunnerGroup) GetHostedRunnersURL() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseRunnerGroup) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseRunnerGroup) GetInherited() bool { _ = "STUB: not implemented"; return false }

func (e *EnterpriseRunnerGroup) GetName() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseRunnerGroup) GetNetworkConfigurationID() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseRunnerGroup) GetRestrictedToWorkflows() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *EnterpriseRunnerGroup) GetRunnersURL() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseRunnerGroup) GetSelectedOrganizationsURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseRunnerGroup) GetSelectedWorkflows() []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseRunnerGroup) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseRunnerGroup) GetWorkflowRestrictionsReadOnly() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *EnterpriseRunnerGroups) GetRunnerGroups() []*EnterpriseRunnerGroup {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseRunnerGroups) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseSecurityAnalysisSettings) GetAdvancedSecurityEnabledForNewRepositories() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *EnterpriseSecurityAnalysisSettings) GetSecretScanningEnabledForNewRepositories() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *EnterpriseSecurityAnalysisSettings) GetSecretScanningPushProtectionCustomLink() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseSecurityAnalysisSettings) GetSecretScanningPushProtectionEnabledForNewRepositories() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *EnterpriseSecurityAnalysisSettings) GetSecretScanningValidityChecksEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *EnterpriseTeam) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (e *EnterpriseTeam) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseTeam) GetGroupID() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseTeam) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseTeam) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseTeam) GetMemberURL() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseTeam) GetName() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseTeam) GetOrganizationSelectionType() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseTeam) GetSlug() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseTeam) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (e *EnterpriseTeam) GetURL() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseTeamCreateOrUpdateRequest) GetDescription() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseTeamCreateOrUpdateRequest) GetGroupID() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseTeamCreateOrUpdateRequest) GetName() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseTeamCreateOrUpdateRequest) GetOrganizationSelectionType() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseUpdateBudget) GetBudgetAlerting() *EnterpriseBudgetAlerting {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseUpdateBudget) GetBudgetAmount() int { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUpdateBudget) GetBudgetEntityName() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUpdateBudget) GetBudgetProductSKU() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUpdateBudget) GetBudgetScope() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUpdateBudget) GetBudgetType() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUpdateBudget) GetPreventFurtherUsage() bool {
	_ = "STUB: not implemented"
	return false
}

func (e *EnterpriseUsageItem) GetDate() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageItem) GetDiscountAmount() float64 { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageItem) GetGrossAmount() float64 { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageItem) GetNetAmount() float64 { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageItem) GetOrganizationName() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageItem) GetPricePerUnit() float64 { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageItem) GetProduct() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageItem) GetQuantity() float64 { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageItem) GetRepositoryName() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageItem) GetSKU() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageItem) GetUnitType() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageReport) GetUsageItems() []*EnterpriseUsageItem {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseUsageReportOptions) GetCostCenterID() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseUsageReportOptions) GetDay() int { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageReportOptions) GetMonth() int { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageReportOptions) GetYear() int { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageSummaryItem) GetDiscountAmount() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *EnterpriseUsageSummaryItem) GetDiscountQuantity() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *EnterpriseUsageSummaryItem) GetGrossAmount() float64 { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageSummaryItem) GetGrossQuantity() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (e *EnterpriseUsageSummaryItem) GetNetAmount() float64 { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageSummaryItem) GetNetQuantity() float64 { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageSummaryItem) GetPricePerUnit() float64 { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageSummaryItem) GetProduct() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageSummaryItem) GetSKU() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageSummaryItem) GetUnitType() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageSummaryOptions) GetOrganization() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseUsageSummaryOptions) GetProduct() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageSummaryOptions) GetRepository() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseUsageSummaryOptions) GetSKU() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageSummaryReport) GetCostCenter() *BillingCostCenter {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseUsageSummaryReport) GetEnterprise() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageSummaryReport) GetOrganization() string {
	_ = "STUB: not implemented"
	return ""
}

func (e *EnterpriseUsageSummaryReport) GetProduct() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageSummaryReport) GetRepository() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageSummaryReport) GetSKU() string { _ = "STUB: not implemented"; return "" }

func (e *EnterpriseUsageSummaryReport) GetTimePeriod() EnterpriseUsageTimePeriod {
	_ = "STUB: not implemented"
	return *new(EnterpriseUsageTimePeriod)
}

func (e *EnterpriseUsageSummaryReport) GetUsageItems() []*EnterpriseUsageSummaryItem {
	_ = "STUB: not implemented"
	return nil
}

func (e *EnterpriseUsageTimePeriod) GetDay() int { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageTimePeriod) GetMonth() int { _ = "STUB: not implemented"; return 0 }

func (e *EnterpriseUsageTimePeriod) GetYear() int { _ = "STUB: not implemented"; return 0 }

func (e *Environment) GetCanAdminsBypass() bool { _ = "STUB: not implemented"; return false }

func (e *Environment) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (e *Environment) GetDeploymentBranchPolicy() *BranchPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (e *Environment) GetEnvironmentName() string { _ = "STUB: not implemented"; return "" }

func (e *Environment) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (e *Environment) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (e *Environment) GetName() string { _ = "STUB: not implemented"; return "" }

func (e *Environment) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (e *Environment) GetOwner() string { _ = "STUB: not implemented"; return "" }

func (e *Environment) GetProtectionRules() []*ProtectionRule { _ = "STUB: not implemented"; return nil }

func (e *Environment) GetRepo() string { _ = "STUB: not implemented"; return "" }

func (e *Environment) GetReviewers() []*EnvReviewers { _ = "STUB: not implemented"; return nil }

func (e *Environment) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (e *Environment) GetURL() string { _ = "STUB: not implemented"; return "" }

func (e *Environment) GetWaitTimer() int { _ = "STUB: not implemented"; return 0 }

func (e *EnvResponse) GetEnvironments() []*Environment { _ = "STUB: not implemented"; return nil }

func (e *EnvResponse) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (e *EnvReviewers) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (e *EnvReviewers) GetType() string { _ = "STUB: not implemented"; return "" }

func (e *Error) GetCode() string { _ = "STUB: not implemented"; return "" }

func (e *Error) GetField() string { _ = "STUB: not implemented"; return "" }

func (e *Error) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (e *Error) GetResource() string { _ = "STUB: not implemented"; return "" }

func (e *ErrorBlock) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (e *ErrorBlock) GetReason() string { _ = "STUB: not implemented"; return "" }

func (e *ErrorResponse) GetBlock() *ErrorBlock { _ = "STUB: not implemented"; return nil }

func (e *ErrorResponse) GetDocumentationURL() string { _ = "STUB: not implemented"; return "" }

func (e *ErrorResponse) GetErrors() []Error { _ = "STUB: not implemented"; return nil }

func (e *ErrorResponse) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (e *Event) GetActor() *User { _ = "STUB: not implemented"; return nil }

func (e *Event) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (e *Event) GetID() string { _ = "STUB: not implemented"; return "" }

func (e *Event) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (e *Event) GetPublic() bool { _ = "STUB: not implemented"; return false }

func (e *Event) GetRawPayload() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (e *Event) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (e *Event) GetType() string { _ = "STUB: not implemented"; return "" }

func (e *ExternalGroup) GetGroupID() int64 { _ = "STUB: not implemented"; return 0 }

func (e *ExternalGroup) GetGroupName() string { _ = "STUB: not implemented"; return "" }

func (e *ExternalGroup) GetMembers() []*ExternalGroupMember { _ = "STUB: not implemented"; return nil }

func (e *ExternalGroup) GetTeams() []*ExternalGroupTeam { _ = "STUB: not implemented"; return nil }

func (e *ExternalGroup) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (e *ExternalGroupList) GetGroups() []*ExternalGroup { _ = "STUB: not implemented"; return nil }

func (e *ExternalGroupMember) GetMemberEmail() string { _ = "STUB: not implemented"; return "" }

func (e *ExternalGroupMember) GetMemberID() int64 { _ = "STUB: not implemented"; return 0 }

func (e *ExternalGroupMember) GetMemberLogin() string { _ = "STUB: not implemented"; return "" }

func (e *ExternalGroupMember) GetMemberName() string { _ = "STUB: not implemented"; return "" }

func (e *ExternalGroupTeam) GetTeamID() int64 { _ = "STUB: not implemented"; return 0 }

func (e *ExternalGroupTeam) GetTeamName() string { _ = "STUB: not implemented"; return "" }

func (f *FeedLink) GetHRef() string { _ = "STUB: not implemented"; return "" }

func (f *FeedLink) GetType() string { _ = "STUB: not implemented"; return "" }

func (f *FeedLinks) GetCurrentUser() *FeedLink { _ = "STUB: not implemented"; return nil }

func (f *FeedLinks) GetCurrentUserActor() *FeedLink { _ = "STUB: not implemented"; return nil }

func (f *FeedLinks) GetCurrentUserOrganization() *FeedLink { _ = "STUB: not implemented"; return nil }

func (f *FeedLinks) GetCurrentUserOrganizations() []*FeedLink {
	_ = "STUB: not implemented"
	return nil
}

func (f *FeedLinks) GetCurrentUserPublic() *FeedLink { _ = "STUB: not implemented"; return nil }

func (f *FeedLinks) GetTimeline() *FeedLink { _ = "STUB: not implemented"; return nil }

func (f *FeedLinks) GetUser() *FeedLink { _ = "STUB: not implemented"; return nil }

func (f *Feeds) GetCurrentUserActorURL() string { _ = "STUB: not implemented"; return "" }

func (f *Feeds) GetCurrentUserOrganizationURL() string { _ = "STUB: not implemented"; return "" }

func (f *Feeds) GetCurrentUserOrganizationURLs() []string { _ = "STUB: not implemented"; return nil }

func (f *Feeds) GetCurrentUserPublicURL() string { _ = "STUB: not implemented"; return "" }

func (f *Feeds) GetCurrentUserURL() string { _ = "STUB: not implemented"; return "" }

func (f *Feeds) GetLinks() *FeedLinks { _ = "STUB: not implemented"; return nil }

func (f *Feeds) GetTimelineURL() string { _ = "STUB: not implemented"; return "" }

func (f *Feeds) GetUserURL() string { _ = "STUB: not implemented"; return "" }

func (f *FieldValue) GetFieldName() string { _ = "STUB: not implemented"; return "" }

func (f *FieldValue) GetFieldNodeID() string { _ = "STUB: not implemented"; return "" }

func (f *FieldValue) GetFieldType() string { _ = "STUB: not implemented"; return "" }

func (f *FieldValue) GetFrom() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (f *FieldValue) GetProjectNumber() int64 { _ = "STUB: not implemented"; return 0 }

func (f *FieldValue) GetTo() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (f *FileExtensionRestrictionBranchRule) GetParameters() FileExtensionRestrictionRuleParameters {
	_ = "STUB: not implemented"
	return *new(FileExtensionRestrictionRuleParameters)
}

func (f *FileExtensionRestrictionRuleParameters) GetRestrictedFileExtensions() []string {
	_ = "STUB: not implemented"
	return nil
}

func (f *FilePathRestrictionBranchRule) GetParameters() FilePathRestrictionRuleParameters {
	_ = "STUB: not implemented"
	return *new(FilePathRestrictionRuleParameters)
}

func (f *FilePathRestrictionRuleParameters) GetRestrictedFilePaths() []string {
	_ = "STUB: not implemented"
	return nil
}

func (f *FineGrainedPersonalAccessTokenRequest) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (f *FineGrainedPersonalAccessTokenRequest) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (f *FineGrainedPersonalAccessTokenRequest) GetOwner() User {
	_ = "STUB: not implemented"
	return *new(User)
}

func (f *FineGrainedPersonalAccessTokenRequest) GetPermissions() PersonalAccessTokenPermissions {
	_ = "STUB: not implemented"
	return *new(PersonalAccessTokenPermissions)
}

func (f *FineGrainedPersonalAccessTokenRequest) GetReason() string {
	_ = "STUB: not implemented"
	return ""
}

func (f *FineGrainedPersonalAccessTokenRequest) GetRepositoriesURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (f *FineGrainedPersonalAccessTokenRequest) GetRepositorySelection() string {
	_ = "STUB: not implemented"
	return ""
}

func (f *FineGrainedPersonalAccessTokenRequest) GetTokenExpired() bool {
	_ = "STUB: not implemented"
	return false
}

func (f *FineGrainedPersonalAccessTokenRequest) GetTokenExpiresAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (f *FineGrainedPersonalAccessTokenRequest) GetTokenID() int64 {
	_ = "STUB: not implemented"
	return 0
}

func (f *FineGrainedPersonalAccessTokenRequest) GetTokenLastUsedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (f *FineGrainedPersonalAccessTokenRequest) GetTokenName() string {
	_ = "STUB: not implemented"
	return ""
}

func (f *FirstPatchedVersion) GetIdentifier() string { _ = "STUB: not implemented"; return "" }

func (f *ForkEvent) GetForkee() *Repository { _ = "STUB: not implemented"; return nil }

func (f *ForkEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (f *ForkEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (f *ForkEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (g *GenerateNotesRequest) GetConfigurationFilePath() string {
	_ = "STUB: not implemented"
	return ""
}

func (g *GenerateNotesRequest) GetPreviousTagName() string { _ = "STUB: not implemented"; return "" }

func (g *GenerateNotesRequest) GetTagName() string { _ = "STUB: not implemented"; return "" }

func (g *GenerateNotesRequest) GetTargetCommitish() string { _ = "STUB: not implemented"; return "" }

func (g *GetAuditLogOptions) GetInclude() string { _ = "STUB: not implemented"; return "" }

func (g *GetAuditLogOptions) GetOrder() string { _ = "STUB: not implemented"; return "" }

func (g *GetAuditLogOptions) GetPhrase() string { _ = "STUB: not implemented"; return "" }

func (g *GetCodeownersErrorsOptions) GetRef() string { _ = "STUB: not implemented"; return "" }

func (g *GetProjectItemOptions) GetFields() []int64 { _ = "STUB: not implemented"; return nil }

func (g *GetProvisionedSCIMGroupEnterpriseOptions) GetExcludedAttributes() string {
	_ = "STUB: not implemented"
	return ""
}

func (g *Gist) GetComments() int { _ = "STUB: not implemented"; return 0 }

func (g *Gist) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (g *Gist) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (g *Gist) GetFiles() map[GistFilename]GistFile { _ = "STUB: not implemented"; return nil }

func (g *Gist) GetGitPullURL() string { _ = "STUB: not implemented"; return "" }

func (g *Gist) GetGitPushURL() string { _ = "STUB: not implemented"; return "" }

func (g *Gist) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (g *Gist) GetID() string { _ = "STUB: not implemented"; return "" }

func (g *Gist) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (g *Gist) GetOwner() *User { _ = "STUB: not implemented"; return nil }

func (g *Gist) GetPublic() bool { _ = "STUB: not implemented"; return false }

func (g *Gist) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (g *GistComment) GetBody() string { _ = "STUB: not implemented"; return "" }

func (g *GistComment) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (g *GistComment) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (g *GistComment) GetURL() string { _ = "STUB: not implemented"; return "" }

func (g *GistComment) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (g *GistCommit) GetChangeStatus() *CommitStats { _ = "STUB: not implemented"; return nil }

func (g *GistCommit) GetCommittedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (g *GistCommit) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (g *GistCommit) GetURL() string { _ = "STUB: not implemented"; return "" }

func (g *GistCommit) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (g *GistCommit) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (g *GistFile) GetContent() string { _ = "STUB: not implemented"; return "" }

func (g *GistFile) GetFilename() string { _ = "STUB: not implemented"; return "" }

func (g *GistFile) GetLanguage() string { _ = "STUB: not implemented"; return "" }

func (g *GistFile) GetRawURL() string { _ = "STUB: not implemented"; return "" }

func (g *GistFile) GetSize() int { _ = "STUB: not implemented"; return 0 }

func (g *GistFile) GetType() string { _ = "STUB: not implemented"; return "" }

func (g *GistFork) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (g *GistFork) GetID() string { _ = "STUB: not implemented"; return "" }

func (g *GistFork) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (g *GistFork) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (g *GistFork) GetURL() string { _ = "STUB: not implemented"; return "" }

func (g *GistFork) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (g *GistListOptions) GetSince() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (g *GistStats) GetPrivateGists() int { _ = "STUB: not implemented"; return 0 }

func (g *GistStats) GetPublicGists() int { _ = "STUB: not implemented"; return 0 }

func (g *GistStats) GetTotalGists() int { _ = "STUB: not implemented"; return 0 }

func (g *GitHubAppAuthorizationEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (g *GitHubAppAuthorizationEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (g *GitHubAppAuthorizationEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (g *Gitignore) GetName() string { _ = "STUB: not implemented"; return "" }

func (g *Gitignore) GetSource() string { _ = "STUB: not implemented"; return "" }

func (g *GitObject) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (g *GitObject) GetType() string { _ = "STUB: not implemented"; return "" }

func (g *GitObject) GetURL() string { _ = "STUB: not implemented"; return "" }

func (g *GlobalSecurityAdvisory) GetCredits() []*Credit { _ = "STUB: not implemented"; return nil }

func (g *GlobalSecurityAdvisory) GetGithubReviewedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (g *GlobalSecurityAdvisory) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (g *GlobalSecurityAdvisory) GetNVDPublishedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (g *GlobalSecurityAdvisory) GetReferences() []string { _ = "STUB: not implemented"; return nil }

func (g *GlobalSecurityAdvisory) GetRepositoryAdvisoryURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (g *GlobalSecurityAdvisory) GetSourceCodeLocation() string {
	_ = "STUB: not implemented"
	return ""
}

func (g *GlobalSecurityAdvisory) GetType() string { _ = "STUB: not implemented"; return "" }

func (g *GlobalSecurityAdvisory) GetVulnerabilities() []*GlobalSecurityVulnerability {
	_ = "STUB: not implemented"
	return nil
}

func (g *GlobalSecurityVulnerability) GetFirstPatchedVersion() string {
	_ = "STUB: not implemented"
	return ""
}

func (g *GlobalSecurityVulnerability) GetPackage() *VulnerabilityPackage {
	_ = "STUB: not implemented"
	return nil
}

func (g *GlobalSecurityVulnerability) GetVulnerableFunctions() []string {
	_ = "STUB: not implemented"
	return nil
}

func (g *GlobalSecurityVulnerability) GetVulnerableVersionRange() string {
	_ = "STUB: not implemented"
	return ""
}

func (g *GollumEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (g *GollumEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (g *GollumEvent) GetPages() []*Page { _ = "STUB: not implemented"; return nil }

func (g *GollumEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (g *GollumEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (g *GoogleCloudConfig) GetBucket() string { _ = "STUB: not implemented"; return "" }

func (g *GoogleCloudConfig) GetEncryptedJSONCredentials() string {
	_ = "STUB: not implemented"
	return ""
}

func (g *GoogleCloudConfig) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (g *GPGEmail) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (g *GPGEmail) GetVerified() bool { _ = "STUB: not implemented"; return false }

func (g *GPGKey) GetCanCertify() bool { _ = "STUB: not implemented"; return false }

func (g *GPGKey) GetCanEncryptComms() bool { _ = "STUB: not implemented"; return false }

func (g *GPGKey) GetCanEncryptStorage() bool { _ = "STUB: not implemented"; return false }

func (g *GPGKey) GetCanSign() bool { _ = "STUB: not implemented"; return false }

func (g *GPGKey) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (g *GPGKey) GetEmails() []*GPGEmail { _ = "STUB: not implemented"; return nil }

func (g *GPGKey) GetExpiresAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (g *GPGKey) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (g *GPGKey) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (g *GPGKey) GetPrimaryKeyID() int64 { _ = "STUB: not implemented"; return 0 }

func (g *GPGKey) GetPublicKey() string { _ = "STUB: not implemented"; return "" }

func (g *GPGKey) GetRawKey() string { _ = "STUB: not implemented"; return "" }

func (g *GPGKey) GetSubkeys() []*GPGKey { _ = "STUB: not implemented"; return nil }

func (g *Grant) GetApp() *AuthorizationApp { _ = "STUB: not implemented"; return nil }

func (g *Grant) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (g *Grant) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (g *Grant) GetScopes() []string { _ = "STUB: not implemented"; return nil }

func (g *Grant) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (g *Grant) GetURL() string { _ = "STUB: not implemented"; return "" }

func (h *HeadCommit) GetAdded() []string { _ = "STUB: not implemented"; return nil }

func (h *HeadCommit) GetAuthor() *CommitAuthor { _ = "STUB: not implemented"; return nil }

func (h *HeadCommit) GetCommitter() *CommitAuthor { _ = "STUB: not implemented"; return nil }

func (h *HeadCommit) GetDistinct() bool { _ = "STUB: not implemented"; return false }

func (h *HeadCommit) GetID() string { _ = "STUB: not implemented"; return "" }

func (h *HeadCommit) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (h *HeadCommit) GetModified() []string { _ = "STUB: not implemented"; return nil }

func (h *HeadCommit) GetRemoved() []string { _ = "STUB: not implemented"; return nil }

func (h *HeadCommit) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (h *HeadCommit) GetTimestamp() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (h *HeadCommit) GetTreeID() string { _ = "STUB: not implemented"; return "" }

func (h *HeadCommit) GetURL() string { _ = "STUB: not implemented"; return "" }

func (h *HecConfig) GetDomain() string { _ = "STUB: not implemented"; return "" }

func (h *HecConfig) GetEncryptedToken() string { _ = "STUB: not implemented"; return "" }

func (h *HecConfig) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (h *HecConfig) GetPath() string { _ = "STUB: not implemented"; return "" }

func (h *HecConfig) GetPort() uint16 { _ = "STUB: not implemented"; return 0 }

func (h *HecConfig) GetSSLVerify() bool { _ = "STUB: not implemented"; return false }

func (h *Hook) GetActive() bool { _ = "STUB: not implemented"; return false }

func (h *Hook) GetConfig() *HookConfig { _ = "STUB: not implemented"; return nil }

func (h *Hook) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (h *Hook) GetEvents() []string { _ = "STUB: not implemented"; return nil }

func (h *Hook) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (h *Hook) GetLastResponse() map[string]any { _ = "STUB: not implemented"; return nil }

func (h *Hook) GetName() string { _ = "STUB: not implemented"; return "" }

func (h *Hook) GetPingURL() string { _ = "STUB: not implemented"; return "" }

func (h *Hook) GetTestURL() string { _ = "STUB: not implemented"; return "" }

func (h *Hook) GetType() string { _ = "STUB: not implemented"; return "" }

func (h *Hook) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (h *Hook) GetURL() string { _ = "STUB: not implemented"; return "" }

func (h *HookConfig) GetContentType() string { _ = "STUB: not implemented"; return "" }

func (h *HookConfig) GetInsecureSSL() string { _ = "STUB: not implemented"; return "" }

func (h *HookConfig) GetSecret() string { _ = "STUB: not implemented"; return "" }

func (h *HookConfig) GetURL() string { _ = "STUB: not implemented"; return "" }

func (h *HookDelivery) GetAction() string { _ = "STUB: not implemented"; return "" }

func (h *HookDelivery) GetDeliveredAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (h *HookDelivery) GetDuration() float64 { _ = "STUB: not implemented"; return 0 }

func (h *HookDelivery) GetEvent() string { _ = "STUB: not implemented"; return "" }

func (h *HookDelivery) GetGUID() string { _ = "STUB: not implemented"; return "" }

func (h *HookDelivery) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (h *HookDelivery) GetInstallationID() int64 { _ = "STUB: not implemented"; return 0 }

func (h *HookDelivery) GetRedelivery() bool { _ = "STUB: not implemented"; return false }

func (h *HookDelivery) GetRepositoryID() int64 { _ = "STUB: not implemented"; return 0 }

func (h *HookDelivery) GetRequest() *HookRequest { _ = "STUB: not implemented"; return nil }

func (h *HookDelivery) GetResponse() *HookResponse { _ = "STUB: not implemented"; return nil }

func (h *HookDelivery) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (h *HookDelivery) GetStatusCode() int { _ = "STUB: not implemented"; return 0 }

func (h *HookRequest) GetHeaders() map[string]string { _ = "STUB: not implemented"; return nil }

func (h *HookRequest) GetRawPayload() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (h *HookResponse) GetHeaders() map[string]string { _ = "STUB: not implemented"; return nil }

func (h *HookResponse) GetRawPayload() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (h *HookStats) GetActiveHooks() int { _ = "STUB: not implemented"; return 0 }

func (h *HookStats) GetInactiveHooks() int { _ = "STUB: not implemented"; return 0 }

func (h *HookStats) GetTotalHooks() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunner) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunner) GetImageDetails() *HostedRunnerImageDetail {
	_ = "STUB: not implemented"
	return nil
}

func (h *HostedRunner) GetLastActiveOn() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (h *HostedRunner) GetMachineSizeDetails() *HostedRunnerMachineSpec {
	_ = "STUB: not implemented"
	return nil
}

func (h *HostedRunner) GetMaximumRunners() int64 { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunner) GetName() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunner) GetPlatform() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunner) GetPublicIPEnabled() bool { _ = "STUB: not implemented"; return false }

func (h *HostedRunner) GetPublicIPs() []*HostedRunnerPublicIP {
	_ = "STUB: not implemented"
	return nil
}

func (h *HostedRunner) GetRunnerGroupID() int64 { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunner) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerCustomImage) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerCustomImage) GetLatestVersion() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerCustomImage) GetName() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerCustomImage) GetPlatform() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerCustomImage) GetSource() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerCustomImage) GetState() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerCustomImage) GetTotalVersionsSize() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerCustomImage) GetVersionsCount() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerCustomImages) GetImages() []*HostedRunnerCustomImage {
	_ = "STUB: not implemented"
	return nil
}

func (h *HostedRunnerCustomImages) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerCustomImageVersion) GetCreatedOn() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (h *HostedRunnerCustomImageVersion) GetSizeGB() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerCustomImageVersion) GetState() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerCustomImageVersion) GetStateDetails() string {
	_ = "STUB: not implemented"
	return ""
}

func (h *HostedRunnerCustomImageVersion) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerCustomImageVersions) GetImageVersions() []*HostedRunnerCustomImageVersion {
	_ = "STUB: not implemented"
	return nil
}

func (h *HostedRunnerCustomImageVersions) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerImage) GetID() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerImage) GetSource() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerImage) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerImageDetail) GetDisplayName() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerImageDetail) GetID() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerImageDetail) GetSizeGB() int64 { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerImageDetail) GetSource() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerImageDetail) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerImages) GetImages() []*HostedRunnerImageSpecs {
	_ = "STUB: not implemented"
	return nil
}

func (h *HostedRunnerImages) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerImageSpecs) GetDisplayName() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerImageSpecs) GetID() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerImageSpecs) GetPlatform() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerImageSpecs) GetSizeGB() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerImageSpecs) GetSource() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerMachineSpec) GetCPUCores() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerMachineSpec) GetID() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerMachineSpec) GetMemoryGB() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerMachineSpec) GetStorageGB() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerMachineSpecs) GetMachineSpecs() []*HostedRunnerMachineSpec {
	_ = "STUB: not implemented"
	return nil
}

func (h *HostedRunnerMachineSpecs) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerPlatforms) GetPlatforms() []string { _ = "STUB: not implemented"; return nil }

func (h *HostedRunnerPlatforms) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerPublicIP) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (h *HostedRunnerPublicIP) GetLength() int { _ = "STUB: not implemented"; return 0 }

func (h *HostedRunnerPublicIP) GetPrefix() string { _ = "STUB: not implemented"; return "" }

func (h *HostedRunnerPublicIPLimits) GetPublicIPs() *PublicIPUsage {
	_ = "STUB: not implemented"
	return nil
}

func (h *HostedRunners) GetRunners() []*HostedRunner { _ = "STUB: not implemented"; return nil }

func (h *HostedRunners) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (h *Hovercard) GetContexts() []*UserContext { _ = "STUB: not implemented"; return nil }

func (h *HovercardOptions) GetSubjectID() string { _ = "STUB: not implemented"; return "" }

func (h *HovercardOptions) GetSubjectType() string { _ = "STUB: not implemented"; return "" }

func (i *IDPGroup) GetGroupDescription() string { _ = "STUB: not implemented"; return "" }

func (i *IDPGroup) GetGroupID() string { _ = "STUB: not implemented"; return "" }

func (i *IDPGroup) GetGroupName() string { _ = "STUB: not implemented"; return "" }

func (i *IDPGroupList) GetGroups() []*IDPGroup { _ = "STUB: not implemented"; return nil }

func (i *ImmutableReleasePolicy) GetEnforcedRepositories() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *ImmutableReleasePolicy) GetSelectedRepositoryIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (i *ImmutableReleaseSettings) GetEnforcedRepositories() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *ImmutableReleaseSettings) GetSelectedRepositoriesURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *ImpersonateUserOptions) GetScopes() []string { _ = "STUB: not implemented"; return nil }

func (i *Import) GetAuthorsCount() int { _ = "STUB: not implemented"; return 0 }

func (i *Import) GetAuthorsURL() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetCommitCount() int { _ = "STUB: not implemented"; return 0 }

func (i *Import) GetFailedStep() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetHasLargeFiles() bool { _ = "STUB: not implemented"; return false }

func (i *Import) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetHumanName() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetLargeFilesCount() int { _ = "STUB: not implemented"; return 0 }

func (i *Import) GetLargeFilesSize() int { _ = "STUB: not implemented"; return 0 }

func (i *Import) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetPercent() int { _ = "STUB: not implemented"; return 0 }

func (i *Import) GetProjectChoices() []*Import { _ = "STUB: not implemented"; return nil }

func (i *Import) GetPushPercent() int { _ = "STUB: not implemented"; return 0 }

func (i *Import) GetRepositoryURL() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetStatusText() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetTFVCProject() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetURL() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetUseLFS() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetVCS() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetVCSPassword() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetVCSURL() string { _ = "STUB: not implemented"; return "" }

func (i *Import) GetVCSUsername() string { _ = "STUB: not implemented"; return "" }

func (i *InitialConfigOptions) GetLicense() string { _ = "STUB: not implemented"; return "" }

func (i *InitialConfigOptions) GetPassword() string { _ = "STUB: not implemented"; return "" }

func (i *InstallableOrganization) GetAccessibleRepositoriesURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallableOrganization) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *InstallableOrganization) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (i *InstallAppRequest) GetClientID() string { _ = "STUB: not implemented"; return "" }

func (i *InstallAppRequest) GetRepositories() []string { _ = "STUB: not implemented"; return nil }

func (i *InstallAppRequest) GetRepositorySelection() string { _ = "STUB: not implemented"; return "" }

func (i *Installation) GetAccessTokensURL() string { _ = "STUB: not implemented"; return "" }

func (i *Installation) GetAccount() *User { _ = "STUB: not implemented"; return nil }

func (i *Installation) GetAppID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *Installation) GetAppSlug() string { _ = "STUB: not implemented"; return "" }

func (i *Installation) GetClientID() string { _ = "STUB: not implemented"; return "" }

func (i *Installation) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *Installation) GetEvents() []string { _ = "STUB: not implemented"; return nil }

func (i *Installation) GetHasMultipleSingleFiles() bool { _ = "STUB: not implemented"; return false }

func (i *Installation) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (i *Installation) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *Installation) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (i *Installation) GetPermissions() *InstallationPermissions {
	_ = "STUB: not implemented"
	return nil
}

func (i *Installation) GetRepositoriesURL() string { _ = "STUB: not implemented"; return "" }

func (i *Installation) GetRepositorySelection() string { _ = "STUB: not implemented"; return "" }

func (i *Installation) GetSingleFileName() string { _ = "STUB: not implemented"; return "" }

func (i *Installation) GetSingleFilePaths() []string { _ = "STUB: not implemented"; return nil }

func (i *Installation) GetSuspendedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (i *Installation) GetSuspendedBy() *User { _ = "STUB: not implemented"; return nil }

func (i *Installation) GetTargetID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *Installation) GetTargetType() string { _ = "STUB: not implemented"; return "" }

func (i *Installation) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *InstallationChanges) GetLogin() *InstallationLoginChange {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationChanges) GetSlug() *InstallationSlugChange {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (i *InstallationEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (i *InstallationEvent) GetRepositories() []*Repository { _ = "STUB: not implemented"; return nil }

func (i *InstallationEvent) GetRequester() *User { _ = "STUB: not implemented"; return nil }

func (i *InstallationEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (i *InstallationLoginChange) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetActions() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetActionsVariables() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetAdministration() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetAttestations() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetBlocking() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetChecks() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetCodespaces() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetCodespacesLifecycleAdmin() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetCodespacesMetadata() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetCodespacesSecrets() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetCodespacesUserSecrets() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetContentReferences() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetContents() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetCopilotMessages() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetDependabotSecrets() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetDeployments() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetDiscussions() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetEmails() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetEnterpriseAIControls() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetEnterpriseCopilotMetrics() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetEnterpriseCredentials() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetEnterpriseCustomEnterpriseRoles() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetEnterpriseCustomOrgRoles() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetEnterpriseCustomProperties() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetEnterpriseCustomPropertiesForOrgs() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetEnterpriseOrganizationInstallations() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetEnterpriseOrganizations() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetEnterpriseOrgInstallationRepos() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetEnterprisePeople() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetEnterpriseSSO() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetEnterpriseTeams() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetEnvironments() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetFollowers() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetGists() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetGitSigningSSHPublicKeys() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetGPGKeys() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetInteractionLimits() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetIssues() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetKeys() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetMembers() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetMergeQueues() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetMetadata() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetOrganizationActionsVariables() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationAdministration() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationAnnouncementBanners() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationAPIInsights() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationCodespaces() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationCodespacesSecrets() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationCodespacesSettings() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationCopilotMetrics() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationCopilotSeatManagement() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationCustomOrgRoles() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationCustomProperties() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationCustomRoles() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationDependabotSecrets() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationEvents() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationHooks() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationKnowledgeBases() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationPackages() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationPersonalAccessTokenRequests() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationPersonalAccessTokens() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationPlan() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationPreReceiveHooks() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationProjects() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationSecrets() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationSelfHostedRunners() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetOrganizationUserBlocking() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetPackages() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetPages() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetPlan() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetProfile() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetPullRequests() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetRepositoryAdvisories() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetRepositoryCustomProperties() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetRepositoryHooks() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetRepositoryPreReceiveHooks() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetRepositoryProjects() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetSecrets() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetSecretScanningAlerts() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetSecurityEvents() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetSingleFile() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetStarring() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetStatuses() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetTeamDiscussions() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetUserEvents() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetVulnerabilityAlerts() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationPermissions) GetWatching() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationPermissions) GetWorkflows() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationRepositoriesEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationRepositoriesEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationRepositoriesEvent) GetOrg() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationRepositoriesEvent) GetRepositoriesAdded() []*Repository {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationRepositoriesEvent) GetRepositoriesRemoved() []*Repository {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationRepositoriesEvent) GetRepositorySelection() string {
	_ = "STUB: not implemented"
	return ""
}

func (i *InstallationRepositoriesEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (i *InstallationRequest) GetAccount() *User { _ = "STUB: not implemented"; return nil }

func (i *InstallationRequest) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (i *InstallationRequest) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *InstallationRequest) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationRequest) GetRequester() *User { _ = "STUB: not implemented"; return nil }

func (i *InstallationSlugChange) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationTargetEvent) GetAccount() *User { _ = "STUB: not implemented"; return nil }

func (i *InstallationTargetEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationTargetEvent) GetChanges() *InstallationChanges {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationTargetEvent) GetEnterprise() *Enterprise {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationTargetEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationTargetEvent) GetOrganization() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationTargetEvent) GetRepository() *Repository {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationTargetEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (i *InstallationTargetEvent) GetTargetType() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationToken) GetExpiresAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (i *InstallationToken) GetPermissions() *InstallationPermissions {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationToken) GetRepositories() []*Repository { _ = "STUB: not implemented"; return nil }

func (i *InstallationToken) GetToken() string { _ = "STUB: not implemented"; return "" }

func (i *InstallationTokenListRepoOptions) GetPermissions() *InstallationPermissions {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationTokenListRepoOptions) GetRepositories() []string {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationTokenListRepoOptions) GetRepositoryIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationTokenOptions) GetPermissions() *InstallationPermissions {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationTokenOptions) GetRepositories() []string {
	_ = "STUB: not implemented"
	return nil
}

func (i *InstallationTokenOptions) GetRepositoryIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (i *InteractionRestriction) GetExpiresAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (i *InteractionRestriction) GetLimit() string { _ = "STUB: not implemented"; return "" }

func (i *InteractionRestriction) GetOrigin() string { _ = "STUB: not implemented"; return "" }

func (i *Invitation) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *Invitation) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (i *Invitation) GetFailedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *Invitation) GetFailedReason() string { _ = "STUB: not implemented"; return "" }

func (i *Invitation) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *Invitation) GetInvitationTeamURL() string { _ = "STUB: not implemented"; return "" }

func (i *Invitation) GetInviter() *User { _ = "STUB: not implemented"; return nil }

func (i *Invitation) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (i *Invitation) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (i *Invitation) GetRole() string { _ = "STUB: not implemented"; return "" }

func (i *Invitation) GetTeamCount() int { _ = "STUB: not implemented"; return 0 }

func (i *Issue) GetActiveLockReason() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetAssignee() *User { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetAssignees() []*User { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetAuthorAssociation() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetBody() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetClosedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *Issue) GetClosedBy() *User { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetComments() int { _ = "STUB: not implemented"; return 0 }

func (i *Issue) GetCommentsURL() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *Issue) GetDraft() bool { _ = "STUB: not implemented"; return false }

func (i *Issue) GetEventsURL() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *Issue) GetIssueDependenciesSummary() *IssueDependenciesSummary {
	_ = "STUB: not implemented"
	return nil
}

func (i *Issue) GetIssueFieldValues() []*IssueFieldValue { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetLabels() []*Label { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetLabelsURL() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetLocked() bool { _ = "STUB: not implemented"; return false }

func (i *Issue) GetMilestone() *Milestone { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (i *Issue) GetParentIssueURL() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetPerformedViaGithubApp() *App { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetPinnedComment() *IssueComment { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetPullRequestLinks() *PullRequestLinks { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetReactions() *Reactions { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetRepositoryURL() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetState() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetStateReason() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetSubIssuesSummary() *SubIssuesSummary { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetTextMatches() []*TextMatch { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetType() *IssueType { _ = "STUB: not implemented"; return nil }

func (i *Issue) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *Issue) GetURL() string { _ = "STUB: not implemented"; return "" }

func (i *Issue) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (i *IssueComment) GetAuthorAssociation() string { _ = "STUB: not implemented"; return "" }

func (i *IssueComment) GetBody() string { _ = "STUB: not implemented"; return "" }

func (i *IssueComment) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *IssueComment) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (i *IssueComment) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *IssueComment) GetIssueURL() string { _ = "STUB: not implemented"; return "" }

func (i *IssueComment) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (i *IssueComment) GetReactions() *Reactions { _ = "STUB: not implemented"; return nil }

func (i *IssueComment) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *IssueComment) GetURL() string { _ = "STUB: not implemented"; return "" }

func (i *IssueComment) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (i *IssueCommentEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (i *IssueCommentEvent) GetChanges() *EditChange { _ = "STUB: not implemented"; return nil }

func (i *IssueCommentEvent) GetComment() *IssueComment { _ = "STUB: not implemented"; return nil }

func (i *IssueCommentEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (i *IssueCommentEvent) GetIssue() *Issue { _ = "STUB: not implemented"; return nil }

func (i *IssueCommentEvent) GetOrganization() *Organization { _ = "STUB: not implemented"; return nil }

func (i *IssueCommentEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (i *IssueCommentEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (i *IssueDependenciesSummary) GetBlockedBy() int { _ = "STUB: not implemented"; return 0 }

func (i *IssueDependenciesSummary) GetBlocking() int { _ = "STUB: not implemented"; return 0 }

func (i *IssueDependenciesSummary) GetTotalBlockedBy() int { _ = "STUB: not implemented"; return 0 }

func (i *IssueDependenciesSummary) GetTotalBlocking() int { _ = "STUB: not implemented"; return 0 }

func (i *IssueDependencyRequest) GetIssueID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *IssueEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (i *IssueEvent) GetActor() *User { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetAssignee() *User { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetAssigner() *User { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetCommitID() string { _ = "STUB: not implemented"; return "" }

func (i *IssueEvent) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *IssueEvent) GetDismissedReview() *DismissedReview { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetEvent() string { _ = "STUB: not implemented"; return "" }

func (i *IssueEvent) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *IssueEvent) GetIssue() *Issue { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetLabel() *Label { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetLockReason() string { _ = "STUB: not implemented"; return "" }

func (i *IssueEvent) GetMilestone() *Milestone { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetPerformedViaGithubApp() *App { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetRename() *Rename { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetRequestedReviewer() *User { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetRequestedTeam() *Team { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetReviewRequester() *User { _ = "STUB: not implemented"; return nil }

func (i *IssueEvent) GetURL() string { _ = "STUB: not implemented"; return "" }

func (i *IssueFieldValue) GetDataType() string { _ = "STUB: not implemented"; return "" }

func (i *IssueFieldValue) GetIssueFieldID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *IssueFieldValue) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (i *IssueFieldValue) GetSingleSelectOption() *IssueFieldValueSingleSelectOption {
	_ = "STUB: not implemented"
	return nil
}

func (i *IssueFieldValue) GetValue() any { _ = "STUB: not implemented"; return *new(any) }

func (i *IssueFieldValueSingleSelectOption) GetColor() string { _ = "STUB: not implemented"; return "" }

func (i *IssueFieldValueSingleSelectOption) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *IssueFieldValueSingleSelectOption) GetName() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImport) GetAssignee() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImport) GetBody() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImport) GetClosed() bool { _ = "STUB: not implemented"; return false }

func (i *IssueImport) GetClosedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *IssueImport) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *IssueImport) GetLabels() []string { _ = "STUB: not implemented"; return nil }

func (i *IssueImport) GetMilestone() int { _ = "STUB: not implemented"; return 0 }

func (i *IssueImport) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImport) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *IssueImportError) GetCode() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImportError) GetField() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImportError) GetLocation() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImportError) GetResource() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImportError) GetValue() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImportRequest) GetComments() []*Comment { _ = "STUB: not implemented"; return nil }

func (i *IssueImportRequest) GetIssueImport() IssueImport {
	_ = "STUB: not implemented"
	return *new(IssueImport)
}

func (i *IssueImportResponse) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (i *IssueImportResponse) GetDocumentationURL() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImportResponse) GetErrors() []*IssueImportError {
	_ = "STUB: not implemented"
	return nil
}

func (i *IssueImportResponse) GetID() int { _ = "STUB: not implemented"; return 0 }

func (i *IssueImportResponse) GetImportIssuesURL() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImportResponse) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImportResponse) GetRepositoryURL() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImportResponse) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (i *IssueImportResponse) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (i *IssueImportResponse) GetURL() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByOrgOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByOrgOptions) GetFilter() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByOrgOptions) GetLabels() []string { _ = "STUB: not implemented"; return nil }

func (i *IssueListByOrgOptions) GetSince() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (i *IssueListByOrgOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByOrgOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByOrgOptions) GetType() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByRepoOptions) GetAssignee() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByRepoOptions) GetCreator() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByRepoOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByRepoOptions) GetLabels() []string { _ = "STUB: not implemented"; return nil }

func (i *IssueListByRepoOptions) GetMentioned() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByRepoOptions) GetMilestone() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByRepoOptions) GetSince() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (i *IssueListByRepoOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByRepoOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListByRepoOptions) GetType() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListCommentsOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (i *IssueListCommentsOptions) GetSince() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (i *IssueListCommentsOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (i *IssueRequest) GetAssignee() string { _ = "STUB: not implemented"; return "" }

func (i *IssueRequest) GetAssignees() []string { _ = "STUB: not implemented"; return nil }

func (i *IssueRequest) GetBody() string { _ = "STUB: not implemented"; return "" }

func (i *IssueRequest) GetIssueFieldValues() []*IssueRequestFieldValue {
	_ = "STUB: not implemented"
	return nil
}

func (i *IssueRequest) GetLabels() []string { _ = "STUB: not implemented"; return nil }

func (i *IssueRequest) GetMilestone() int { _ = "STUB: not implemented"; return 0 }

func (i *IssueRequest) GetState() string { _ = "STUB: not implemented"; return "" }

func (i *IssueRequest) GetStateReason() string { _ = "STUB: not implemented"; return "" }

func (i *IssueRequest) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (i *IssueRequest) GetType() string { _ = "STUB: not implemented"; return "" }

func (i *IssueRequestFieldValue) GetFieldID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *IssueRequestFieldValue) GetValue() any { _ = "STUB: not implemented"; return *new(any) }

func (i *IssuesEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (i *IssuesEvent) GetAssignee() *User { _ = "STUB: not implemented"; return nil }

func (i *IssuesEvent) GetChanges() *EditChange { _ = "STUB: not implemented"; return nil }

func (i *IssuesEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (i *IssuesEvent) GetIssue() *Issue { _ = "STUB: not implemented"; return nil }

func (i *IssuesEvent) GetLabel() *Label { _ = "STUB: not implemented"; return nil }

func (i *IssuesEvent) GetMilestone() *Milestone { _ = "STUB: not implemented"; return nil }

func (i *IssuesEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (i *IssuesEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (i *IssuesEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (i *IssuesSearchResult) GetIncompleteResults() bool { _ = "STUB: not implemented"; return false }

func (i *IssuesSearchResult) GetIssues() []*Issue { _ = "STUB: not implemented"; return nil }

func (i *IssuesSearchResult) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (i *IssueStats) GetClosedIssues() int { _ = "STUB: not implemented"; return 0 }

func (i *IssueStats) GetOpenIssues() int { _ = "STUB: not implemented"; return 0 }

func (i *IssueStats) GetTotalIssues() int { _ = "STUB: not implemented"; return 0 }

func (i *IssueType) GetColor() string { _ = "STUB: not implemented"; return "" }

func (i *IssueType) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (i *IssueType) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (i *IssueType) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *IssueType) GetName() string { _ = "STUB: not implemented"; return "" }

func (i *IssueType) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (i *IssueType) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (j *JITRunnerConfig) GetEncodedJITConfig() string { _ = "STUB: not implemented"; return "" }

func (j *JITRunnerConfig) GetRunner() *Runner { _ = "STUB: not implemented"; return nil }

func (j *Jobs) GetJobs() []*WorkflowJob { _ = "STUB: not implemented"; return nil }

func (j *Jobs) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (k *Key) GetAddedBy() string { _ = "STUB: not implemented"; return "" }

func (k *Key) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (k *Key) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (k *Key) GetKey() string { _ = "STUB: not implemented"; return "" }

func (k *Key) GetLastUsed() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (k *Key) GetReadOnly() bool { _ = "STUB: not implemented"; return false }

func (k *Key) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (k *Key) GetURL() string { _ = "STUB: not implemented"; return "" }

func (k *Key) GetVerified() bool { _ = "STUB: not implemented"; return false }

func (l *Label) GetColor() string { _ = "STUB: not implemented"; return "" }

func (l *Label) GetDefault() bool { _ = "STUB: not implemented"; return false }

func (l *Label) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (l *Label) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (l *Label) GetName() string { _ = "STUB: not implemented"; return "" }

func (l *Label) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (l *Label) GetURL() string { _ = "STUB: not implemented"; return "" }

func (l *LabelEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (l *LabelEvent) GetChanges() *EditChange { _ = "STUB: not implemented"; return nil }

func (l *LabelEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (l *LabelEvent) GetLabel() *Label { _ = "STUB: not implemented"; return nil }

func (l *LabelEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (l *LabelEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (l *LabelEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (l *LabelResult) GetColor() string { _ = "STUB: not implemented"; return "" }

func (l *LabelResult) GetDefault() bool { _ = "STUB: not implemented"; return false }

func (l *LabelResult) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (l *LabelResult) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (l *LabelResult) GetName() string { _ = "STUB: not implemented"; return "" }

func (l *LabelResult) GetScore() float64 { _ = "STUB: not implemented"; return 0 }

func (l *LabelResult) GetURL() string { _ = "STUB: not implemented"; return "" }

func (l *LabelsSearchResult) GetIncompleteResults() bool { _ = "STUB: not implemented"; return false }

func (l *LabelsSearchResult) GetLabels() []*LabelResult { _ = "STUB: not implemented"; return nil }

func (l *LabelsSearchResult) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (l *LargeFile) GetOID() string { _ = "STUB: not implemented"; return "" }

func (l *LargeFile) GetPath() string { _ = "STUB: not implemented"; return "" }

func (l *LargeFile) GetRefName() string { _ = "STUB: not implemented"; return "" }

func (l *LargeFile) GetSize() int { _ = "STUB: not implemented"; return 0 }

func (l *LastLicenseSync) GetProperties() *LastLicenseSyncProperties {
	_ = "STUB: not implemented"
	return nil
}

func (l *LastLicenseSync) GetType() string { _ = "STUB: not implemented"; return "" }

func (l *LastLicenseSyncProperties) GetDate() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (l *LastLicenseSyncProperties) GetError() string { _ = "STUB: not implemented"; return "" }

func (l *LastLicenseSyncProperties) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (l *License) GetBody() string { _ = "STUB: not implemented"; return "" }

func (l *License) GetConditions() []string { _ = "STUB: not implemented"; return nil }

func (l *License) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (l *License) GetFeatured() bool { _ = "STUB: not implemented"; return false }

func (l *License) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (l *License) GetImplementation() string { _ = "STUB: not implemented"; return "" }

func (l *License) GetKey() string { _ = "STUB: not implemented"; return "" }

func (l *License) GetLimitations() []string { _ = "STUB: not implemented"; return nil }

func (l *License) GetName() string { _ = "STUB: not implemented"; return "" }

func (l *License) GetPermissions() []string { _ = "STUB: not implemented"; return nil }

func (l *License) GetSPDXID() string { _ = "STUB: not implemented"; return "" }

func (l *License) GetURL() string { _ = "STUB: not implemented"; return "" }

func (l *LicenseCheck) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (l *LicenseStatus) GetAdvancedSecurityEnabled() bool { _ = "STUB: not implemented"; return false }

func (l *LicenseStatus) GetAdvancedSecuritySeats() int { _ = "STUB: not implemented"; return 0 }

func (l *LicenseStatus) GetClusterSupport() bool { _ = "STUB: not implemented"; return false }

func (l *LicenseStatus) GetCompany() string { _ = "STUB: not implemented"; return "" }

func (l *LicenseStatus) GetCroquetSupport() bool { _ = "STUB: not implemented"; return false }

func (l *LicenseStatus) GetCustomTerms() bool { _ = "STUB: not implemented"; return false }

func (l *LicenseStatus) GetEvaluation() bool { _ = "STUB: not implemented"; return false }

func (l *LicenseStatus) GetExpireAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (l *LicenseStatus) GetInsightsEnabled() bool { _ = "STUB: not implemented"; return false }

func (l *LicenseStatus) GetInsightsExpireAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (l *LicenseStatus) GetLearningLabEvaluationExpires() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (l *LicenseStatus) GetLearningLabSeats() int { _ = "STUB: not implemented"; return 0 }

func (l *LicenseStatus) GetPerpetual() bool { _ = "STUB: not implemented"; return false }

func (l *LicenseStatus) GetReferenceNumber() string { _ = "STUB: not implemented"; return "" }

func (l *LicenseStatus) GetSeats() int { _ = "STUB: not implemented"; return 0 }

func (l *LicenseStatus) GetSSHAllowed() bool { _ = "STUB: not implemented"; return false }

func (l *LicenseStatus) GetSupportKey() bool { _ = "STUB: not implemented"; return false }

func (l *LicenseStatus) GetUnlimitedSeating() bool { _ = "STUB: not implemented"; return false }

func (l *LinearHistoryRequirementEnforcementLevelChanges) GetFrom() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListAlertsOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (l *ListAlertsOptions) GetEcosystem() string { _ = "STUB: not implemented"; return "" }

func (l *ListAlertsOptions) GetPackage() string { _ = "STUB: not implemented"; return "" }

func (l *ListAlertsOptions) GetScope() string { _ = "STUB: not implemented"; return "" }

func (l *ListAlertsOptions) GetSeverity() string { _ = "STUB: not implemented"; return "" }

func (l *ListAlertsOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (l *ListAlertsOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (l *ListAllIssuesOptions) GetCollab() bool { _ = "STUB: not implemented"; return false }

func (l *ListAllIssuesOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (l *ListAllIssuesOptions) GetFilter() string { _ = "STUB: not implemented"; return "" }

func (l *ListAllIssuesOptions) GetLabels() []string { _ = "STUB: not implemented"; return nil }

func (l *ListAllIssuesOptions) GetOrgs() bool { _ = "STUB: not implemented"; return false }

func (l *ListAllIssuesOptions) GetOwned() bool { _ = "STUB: not implemented"; return false }

func (l *ListAllIssuesOptions) GetPulls() bool { _ = "STUB: not implemented"; return false }

func (l *ListAllIssuesOptions) GetSince() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (l *ListAllIssuesOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (l *ListAllIssuesOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (l *ListArtifactsOptions) GetName() string { _ = "STUB: not implemented"; return "" }

func (l *ListCheckRunsOptions) GetAppID() int64 { _ = "STUB: not implemented"; return 0 }

func (l *ListCheckRunsOptions) GetCheckName() string { _ = "STUB: not implemented"; return "" }

func (l *ListCheckRunsOptions) GetFilter() string { _ = "STUB: not implemented"; return "" }

func (l *ListCheckRunsOptions) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (l *ListCheckRunsResults) GetCheckRuns() []*CheckRun { _ = "STUB: not implemented"; return nil }

func (l *ListCheckRunsResults) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (l *ListCheckSuiteOptions) GetAppID() int64 { _ = "STUB: not implemented"; return 0 }

func (l *ListCheckSuiteOptions) GetCheckName() string { _ = "STUB: not implemented"; return "" }

func (l *ListCheckSuiteResults) GetCheckSuites() []*CheckSuite {
	_ = "STUB: not implemented"
	return nil
}

func (l *ListCheckSuiteResults) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (l *ListCodeQualityFindingsOptions) GetDirection() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListCodeQualityFindingsOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (l *ListCodeSecurityConfigurationRepositoriesOptions) GetAfter() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListCodeSecurityConfigurationRepositoriesOptions) GetBefore() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListCodeSecurityConfigurationRepositoriesOptions) GetPerPage() int {
	_ = "STUB: not implemented"
	return 0
}

func (l *ListCodeSecurityConfigurationRepositoriesOptions) GetStatus() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListCodespaces) GetCodespaces() []*Codespace { _ = "STUB: not implemented"; return nil }

func (l *ListCodespaces) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (l *ListCodespacesOptions) GetRepositoryID() int64 { _ = "STUB: not implemented"; return 0 }

func (l *ListCollaboratorsOptions) GetAffiliation() string { _ = "STUB: not implemented"; return "" }

func (l *ListCollaboratorsOptions) GetPermission() string { _ = "STUB: not implemented"; return "" }

func (l *ListContributorsOptions) GetAnon() string { _ = "STUB: not implemented"; return "" }

func (l *ListCopilotSeatsResponse) GetSeats() []*CopilotSeatDetails {
	_ = "STUB: not implemented"
	return nil
}

func (l *ListCopilotSeatsResponse) GetTotalSeats() int64 { _ = "STUB: not implemented"; return 0 }

func (l *ListCostCenterOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (l *ListCursorOptions) GetAfter() string { _ = "STUB: not implemented"; return "" }

func (l *ListCursorOptions) GetBefore() string { _ = "STUB: not implemented"; return "" }

func (l *ListCursorOptions) GetCursor() string { _ = "STUB: not implemented"; return "" }

func (l *ListCursorOptions) GetFirst() int { _ = "STUB: not implemented"; return 0 }

func (l *ListCursorOptions) GetLast() int { _ = "STUB: not implemented"; return 0 }

func (l *ListCursorOptions) GetPage() string { _ = "STUB: not implemented"; return "" }

func (l *ListCursorOptions) GetPerPage() int { _ = "STUB: not implemented"; return 0 }

func (l *ListCustomDeploymentRuleIntegrationsResponse) GetAvailableIntegrations() []*CustomDeploymentProtectionRuleApp {
	_ = "STUB: not implemented"
	return nil
}

func (l *ListCustomDeploymentRuleIntegrationsResponse) GetTotalCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (l *ListCustomPropertyValuesOptions) GetRepositoryQuery() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListDeploymentProtectionRuleResponse) GetProtectionRules() []*CustomDeploymentProtectionRule {
	_ = "STUB: not implemented"
	return nil
}

func (l *ListDeploymentProtectionRuleResponse) GetTotalCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (l *ListEnterpriseCodeSecurityConfigurationOptions) GetAfter() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListEnterpriseCodeSecurityConfigurationOptions) GetBefore() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListEnterpriseCodeSecurityConfigurationOptions) GetPerPage() int {
	_ = "STUB: not implemented"
	return 0
}

func (l *ListEnterpriseRunnerGroupOptions) GetVisibleToOrganization() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListExternalGroupsOptions) GetDisplayName() string { _ = "STUB: not implemented"; return "" }

func (l *ListFineGrainedPATOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (l *ListFineGrainedPATOptions) GetLastUsedAfter() string { _ = "STUB: not implemented"; return "" }

func (l *ListFineGrainedPATOptions) GetLastUsedBefore() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListFineGrainedPATOptions) GetOwner() []string { _ = "STUB: not implemented"; return nil }

func (l *ListFineGrainedPATOptions) GetPermission() string { _ = "STUB: not implemented"; return "" }

func (l *ListFineGrainedPATOptions) GetRepository() string { _ = "STUB: not implemented"; return "" }

func (l *ListFineGrainedPATOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (l *ListFineGrainedPATOptions) GetTokenID() []int64 { _ = "STUB: not implemented"; return nil }

func (l *ListGlobalSecurityAdvisoriesOptions) GetAffects() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListGlobalSecurityAdvisoriesOptions) GetCVEID() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListGlobalSecurityAdvisoriesOptions) GetCWEs() []string {
	_ = "STUB: not implemented"
	return nil
}

func (l *ListGlobalSecurityAdvisoriesOptions) GetEcosystem() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListGlobalSecurityAdvisoriesOptions) GetGHSAID() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListGlobalSecurityAdvisoriesOptions) GetIsWithdrawn() bool {
	_ = "STUB: not implemented"
	return false
}

func (l *ListGlobalSecurityAdvisoriesOptions) GetModified() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListGlobalSecurityAdvisoriesOptions) GetPublished() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListGlobalSecurityAdvisoriesOptions) GetSeverity() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListGlobalSecurityAdvisoriesOptions) GetType() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListGlobalSecurityAdvisoriesOptions) GetUpdated() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListIDPGroupsOptions) GetQuery() string { _ = "STUB: not implemented"; return "" }

func (l *ListLicensesOptions) GetFeatured() bool { _ = "STUB: not implemented"; return false }

func (l *ListMembersOptions) GetFilter() string { _ = "STUB: not implemented"; return "" }

func (l *ListMembersOptions) GetPublicOnly() bool { _ = "STUB: not implemented"; return false }

func (l *ListMembersOptions) GetRole() string { _ = "STUB: not implemented"; return "" }

func (l *ListOptions) GetPage() int { _ = "STUB: not implemented"; return 0 }

func (l *ListOptions) GetPerPage() int { _ = "STUB: not implemented"; return 0 }

func (l *ListOrganizationCopilotCodingAgentRepositoriesResponse) GetRepositories() []*Repository {
	_ = "STUB: not implemented"
	return nil
}

func (l *ListOrganizationCopilotCodingAgentRepositoriesResponse) GetTotalCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (l *ListOrganizations) GetOrganizations() []*Organization {
	_ = "STUB: not implemented"
	return nil
}

func (l *ListOrganizations) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (l *ListOrgCodeSecurityConfigurationOptions) GetAfter() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListOrgCodeSecurityConfigurationOptions) GetBefore() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListOrgCodeSecurityConfigurationOptions) GetPerPage() int {
	_ = "STUB: not implemented"
	return 0
}

func (l *ListOrgCodeSecurityConfigurationOptions) GetTargetType() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListOrgMembershipsOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (l *ListOrgRunnerGroupOptions) GetVisibleToRepository() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListOutsideCollaboratorsOptions) GetFilter() string { _ = "STUB: not implemented"; return "" }

func (l *ListPackageVersionsOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (l *ListProjectItemsOptions) GetFields() []int64 { _ = "STUB: not implemented"; return nil }

func (l *ListProjectsOptions) GetQuery() string { _ = "STUB: not implemented"; return "" }

func (l *ListProjectsPaginationOptions) GetAfter() string { _ = "STUB: not implemented"; return "" }

func (l *ListProjectsPaginationOptions) GetBefore() string { _ = "STUB: not implemented"; return "" }

func (l *ListProjectsPaginationOptions) GetPerPage() int { _ = "STUB: not implemented"; return 0 }

func (l *ListProvisionedSCIMGroupsEnterpriseOptions) GetCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (l *ListProvisionedSCIMGroupsEnterpriseOptions) GetExcludedAttributes() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListProvisionedSCIMGroupsEnterpriseOptions) GetFilter() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListProvisionedSCIMGroupsEnterpriseOptions) GetStartIndex() int {
	_ = "STUB: not implemented"
	return 0
}

func (l *ListProvisionedSCIMUsersEnterpriseOptions) GetCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (l *ListProvisionedSCIMUsersEnterpriseOptions) GetFilter() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListProvisionedSCIMUsersEnterpriseOptions) GetStartIndex() int {
	_ = "STUB: not implemented"
	return 0
}

func (l *ListReactionOptions) GetContent() string { _ = "STUB: not implemented"; return "" }

func (l *ListRepoMachineTypesOptions) GetClientIP() string { _ = "STUB: not implemented"; return "" }

func (l *ListRepoMachineTypesOptions) GetLocation() string { _ = "STUB: not implemented"; return "" }

func (l *ListRepoMachineTypesOptions) GetRef() string { _ = "STUB: not implemented"; return "" }

func (l *ListRepositories) GetRepositories() []*Repository { _ = "STUB: not implemented"; return nil }

func (l *ListRepositories) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (l *ListRepositoryActivityOptions) GetActivityType() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListRepositoryActivityOptions) GetActor() string { _ = "STUB: not implemented"; return "" }

func (l *ListRepositoryActivityOptions) GetAfter() string { _ = "STUB: not implemented"; return "" }

func (l *ListRepositoryActivityOptions) GetBefore() string { _ = "STUB: not implemented"; return "" }

func (l *ListRepositoryActivityOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (l *ListRepositoryActivityOptions) GetPerPage() int { _ = "STUB: not implemented"; return 0 }

func (l *ListRepositoryActivityOptions) GetRef() string { _ = "STUB: not implemented"; return "" }

func (l *ListRepositoryActivityOptions) GetTimePeriod() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListRepositorySecurityAdvisoriesOptions) GetDirection() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListRepositorySecurityAdvisoriesOptions) GetSort() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListRepositorySecurityAdvisoriesOptions) GetState() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListRunnersOptions) GetName() string { _ = "STUB: not implemented"; return "" }

func (l *ListSCIMProvisionedIdentitiesOptions) GetCount() int { _ = "STUB: not implemented"; return 0 }

func (l *ListSCIMProvisionedIdentitiesOptions) GetFilter() string {
	_ = "STUB: not implemented"
	return ""
}

func (l *ListSCIMProvisionedIdentitiesOptions) GetStartIndex() int {
	_ = "STUB: not implemented"
	return 0
}

func (l *ListUserIssuesOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (l *ListUserIssuesOptions) GetFilter() string { _ = "STUB: not implemented"; return "" }

func (l *ListUserIssuesOptions) GetLabels() []string { _ = "STUB: not implemented"; return nil }

func (l *ListUserIssuesOptions) GetSince() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (l *ListUserIssuesOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (l *ListUserIssuesOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (l *ListWorkflowJobsOptions) GetFilter() string { _ = "STUB: not implemented"; return "" }

func (l *ListWorkflowRunsOptions) GetActor() string { _ = "STUB: not implemented"; return "" }

func (l *ListWorkflowRunsOptions) GetBranch() string { _ = "STUB: not implemented"; return "" }

func (l *ListWorkflowRunsOptions) GetCheckSuiteID() int64 { _ = "STUB: not implemented"; return 0 }

func (l *ListWorkflowRunsOptions) GetCreated() string { _ = "STUB: not implemented"; return "" }

func (l *ListWorkflowRunsOptions) GetEvent() string { _ = "STUB: not implemented"; return "" }

func (l *ListWorkflowRunsOptions) GetExcludePullRequests() bool {
	_ = "STUB: not implemented"
	return false
}

func (l *ListWorkflowRunsOptions) GetHeadSHA() string { _ = "STUB: not implemented"; return "" }

func (l *ListWorkflowRunsOptions) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (l *Location) GetEndColumn() int { _ = "STUB: not implemented"; return 0 }

func (l *Location) GetEndLine() int { _ = "STUB: not implemented"; return 0 }

func (l *Location) GetPath() string { _ = "STUB: not implemented"; return "" }

func (l *Location) GetStartColumn() int { _ = "STUB: not implemented"; return 0 }

func (l *Location) GetStartLine() int { _ = "STUB: not implemented"; return 0 }

func (l *LockBranch) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (l *LockIssueOptions) GetLockReason() string { _ = "STUB: not implemented"; return "" }

func (m *MaintenanceOperationStatus) GetHostname() string { _ = "STUB: not implemented"; return "" }

func (m *MaintenanceOperationStatus) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (m *MaintenanceOperationStatus) GetUUID() string { _ = "STUB: not implemented"; return "" }

func (m *MaintenanceOptions) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (m *MaintenanceOptions) GetIPExceptionList() []string { _ = "STUB: not implemented"; return nil }

func (m *MaintenanceOptions) GetMaintenanceModeMessage() string {
	_ = "STUB: not implemented"
	return ""
}

func (m *MaintenanceOptions) GetUUID() string { _ = "STUB: not implemented"; return "" }

func (m *MaintenanceOptions) GetWhen() string { _ = "STUB: not implemented"; return "" }

func (m *MaintenanceStatus) GetCanUnsetMaintenance() bool { _ = "STUB: not implemented"; return false }

func (m *MaintenanceStatus) GetConnectionServices() []*ConnectionServiceItem {
	_ = "STUB: not implemented"
	return nil
}

func (m *MaintenanceStatus) GetHostname() string { _ = "STUB: not implemented"; return "" }

func (m *MaintenanceStatus) GetIPExceptionList() []string { _ = "STUB: not implemented"; return nil }

func (m *MaintenanceStatus) GetMaintenanceModeMessage() string {
	_ = "STUB: not implemented"
	return ""
}

func (m *MaintenanceStatus) GetScheduledTime() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (m *MaintenanceStatus) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (m *MaintenanceStatus) GetUUID() string { _ = "STUB: not implemented"; return "" }

func (m *MarkdownOptions) GetContext() string { _ = "STUB: not implemented"; return "" }

func (m *MarkdownOptions) GetMode() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePendingChange) GetEffectiveDate() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (m *MarketplacePendingChange) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (m *MarketplacePendingChange) GetPlan() *MarketplacePlan {
	_ = "STUB: not implemented"
	return nil
}

func (m *MarketplacePendingChange) GetUnitCount() int { _ = "STUB: not implemented"; return 0 }

func (m *MarketplacePlan) GetAccountsURL() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePlan) GetBullets() []string { _ = "STUB: not implemented"; return nil }

func (m *MarketplacePlan) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePlan) GetHasFreeTrial() bool { _ = "STUB: not implemented"; return false }

func (m *MarketplacePlan) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (m *MarketplacePlan) GetMonthlyPriceInCents() int { _ = "STUB: not implemented"; return 0 }

func (m *MarketplacePlan) GetName() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePlan) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (m *MarketplacePlan) GetPriceModel() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePlan) GetState() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePlan) GetUnitName() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePlan) GetURL() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePlan) GetYearlyPriceInCents() int { _ = "STUB: not implemented"; return 0 }

func (m *MarketplacePlanAccount) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (m *MarketplacePlanAccount) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePlanAccount) GetMarketplacePendingChange() *MarketplacePendingChange {
	_ = "STUB: not implemented"
	return nil
}

func (m *MarketplacePlanAccount) GetMarketplacePurchase() *MarketplacePurchase {
	_ = "STUB: not implemented"
	return nil
}

func (m *MarketplacePlanAccount) GetOrganizationBillingEmail() string {
	_ = "STUB: not implemented"
	return ""
}

func (m *MarketplacePlanAccount) GetType() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePlanAccount) GetURL() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePurchase) GetAccount() *MarketplacePurchaseAccount {
	_ = "STUB: not implemented"
	return nil
}

func (m *MarketplacePurchase) GetBillingCycle() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePurchase) GetFreeTrialEndsOn() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (m *MarketplacePurchase) GetNextBillingDate() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (m *MarketplacePurchase) GetOnFreeTrial() bool { _ = "STUB: not implemented"; return false }

func (m *MarketplacePurchase) GetPlan() *MarketplacePlan { _ = "STUB: not implemented"; return nil }

func (m *MarketplacePurchase) GetUnitCount() int { _ = "STUB: not implemented"; return 0 }

func (m *MarketplacePurchase) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (m *MarketplacePurchaseAccount) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePurchaseAccount) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (m *MarketplacePurchaseAccount) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePurchaseAccount) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePurchaseAccount) GetOrganizationBillingEmail() string {
	_ = "STUB: not implemented"
	return ""
}

func (m *MarketplacePurchaseAccount) GetType() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePurchaseAccount) GetURL() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePurchaseEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (m *MarketplacePurchaseEvent) GetEffectiveDate() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (m *MarketplacePurchaseEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (m *MarketplacePurchaseEvent) GetMarketplacePurchase() *MarketplacePurchase {
	_ = "STUB: not implemented"
	return nil
}

func (m *MarketplacePurchaseEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (m *MarketplacePurchaseEvent) GetPreviousMarketplacePurchase() *MarketplacePurchase {
	_ = "STUB: not implemented"
	return nil
}

func (m *MarketplacePurchaseEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (m *Match) GetIndices() []int { _ = "STUB: not implemented"; return nil }

func (m *Match) GetText() string { _ = "STUB: not implemented"; return "" }

func (m *MaxFilePathLengthBranchRule) GetParameters() MaxFilePathLengthRuleParameters {
	_ = "STUB: not implemented"
	return *new(MaxFilePathLengthRuleParameters)
}

func (m *MaxFilePathLengthRuleParameters) GetMaxFilePathLength() int {
	_ = "STUB: not implemented"
	return 0
}

func (m *MaxFileSizeBranchRule) GetParameters() MaxFileSizeRuleParameters {
	_ = "STUB: not implemented"
	return *new(MaxFileSizeRuleParameters)
}

func (m *MaxFileSizeRuleParameters) GetMaxFileSize() int64 { _ = "STUB: not implemented"; return 0 }

func (m *MemberChanges) GetPermission() *MemberChangesPermission {
	_ = "STUB: not implemented"
	return nil
}

func (m *MemberChanges) GetRoleName() *MemberChangesRoleName { _ = "STUB: not implemented"; return nil }

func (m *MemberChangesPermission) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (m *MemberChangesPermission) GetTo() string { _ = "STUB: not implemented"; return "" }

func (m *MemberChangesRoleName) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (m *MemberChangesRoleName) GetTo() string { _ = "STUB: not implemented"; return "" }

func (m *MemberEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (m *MemberEvent) GetChanges() *MemberChanges { _ = "STUB: not implemented"; return nil }

func (m *MemberEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (m *MemberEvent) GetMember() *User { _ = "STUB: not implemented"; return nil }

func (m *MemberEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (m *MemberEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (m *MemberEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (m *Membership) GetOrganization() *Organization { _ = "STUB: not implemented"; return nil }

func (m *Membership) GetOrganizationURL() string { _ = "STUB: not implemented"; return "" }

func (m *Membership) GetRole() string { _ = "STUB: not implemented"; return "" }

func (m *Membership) GetState() string { _ = "STUB: not implemented"; return "" }

func (m *Membership) GetURL() string { _ = "STUB: not implemented"; return "" }

func (m *Membership) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (m *MembershipEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (m *MembershipEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (m *MembershipEvent) GetMember() *User { _ = "STUB: not implemented"; return nil }

func (m *MembershipEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (m *MembershipEvent) GetScope() string { _ = "STUB: not implemented"; return "" }

func (m *MembershipEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (m *MembershipEvent) GetTeam() *Team { _ = "STUB: not implemented"; return nil }

func (m *MergeGroup) GetBaseRef() string { _ = "STUB: not implemented"; return "" }

func (m *MergeGroup) GetBaseSHA() string { _ = "STUB: not implemented"; return "" }

func (m *MergeGroup) GetHeadCommit() *Commit { _ = "STUB: not implemented"; return nil }

func (m *MergeGroup) GetHeadRef() string { _ = "STUB: not implemented"; return "" }

func (m *MergeGroup) GetHeadSHA() string { _ = "STUB: not implemented"; return "" }

func (m *MergeGroupEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (m *MergeGroupEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (m *MergeGroupEvent) GetMergeGroup() *MergeGroup { _ = "STUB: not implemented"; return nil }

func (m *MergeGroupEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (m *MergeGroupEvent) GetReason() string { _ = "STUB: not implemented"; return "" }

func (m *MergeGroupEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (m *MergeGroupEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (m *MergeQueueBranchRule) GetParameters() MergeQueueRuleParameters {
	_ = "STUB: not implemented"
	return *new(MergeQueueRuleParameters)
}

func (m *MergeQueueRuleParameters) GetCheckResponseTimeoutMinutes() int {
	_ = "STUB: not implemented"
	return 0
}

func (m *MergeQueueRuleParameters) GetGroupingStrategy() MergeGroupingStrategy {
	_ = "STUB: not implemented"
	return *new(MergeGroupingStrategy)
}

func (m *MergeQueueRuleParameters) GetMaxEntriesToBuild() int { _ = "STUB: not implemented"; return 0 }

func (m *MergeQueueRuleParameters) GetMaxEntriesToMerge() int { _ = "STUB: not implemented"; return 0 }

func (m *MergeQueueRuleParameters) GetMergeMethod() MergeQueueMergeMethod {
	_ = "STUB: not implemented"
	return *new(MergeQueueMergeMethod)
}

func (m *MergeQueueRuleParameters) GetMinEntriesToMerge() int { _ = "STUB: not implemented"; return 0 }

func (m *MergeQueueRuleParameters) GetMinEntriesToMergeWaitMinutes() int {
	_ = "STUB: not implemented"
	return 0
}

func (m *Message) GetText() string { _ = "STUB: not implemented"; return "" }

func (m *MetaEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (m *MetaEvent) GetHook() *Hook { _ = "STUB: not implemented"; return nil }

func (m *MetaEvent) GetHookID() int64 { _ = "STUB: not implemented"; return 0 }

func (m *MetaEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (m *MetaEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (m *MetaEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (m *MetaEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (m *Metric) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (m *Metric) GetKey() string { _ = "STUB: not implemented"; return "" }

func (m *Metric) GetName() string { _ = "STUB: not implemented"; return "" }

func (m *Metric) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (m *Metric) GetSPDXID() string { _ = "STUB: not implemented"; return "" }

func (m *Metric) GetURL() string { _ = "STUB: not implemented"; return "" }

func (m *Migration) GetCreatedAt() string { _ = "STUB: not implemented"; return "" }

func (m *Migration) GetExcludeAttachments() bool { _ = "STUB: not implemented"; return false }

func (m *Migration) GetGUID() string { _ = "STUB: not implemented"; return "" }

func (m *Migration) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (m *Migration) GetLockRepositories() bool { _ = "STUB: not implemented"; return false }

func (m *Migration) GetRepositories() []*Repository { _ = "STUB: not implemented"; return nil }

func (m *Migration) GetState() string { _ = "STUB: not implemented"; return "" }

func (m *Migration) GetUpdatedAt() string { _ = "STUB: not implemented"; return "" }

func (m *Migration) GetURL() string { _ = "STUB: not implemented"; return "" }

func (m *MigrationOptions) GetExclude() []string { _ = "STUB: not implemented"; return nil }

func (m *MigrationOptions) GetExcludeAttachments() bool { _ = "STUB: not implemented"; return false }

func (m *MigrationOptions) GetExcludeReleases() bool { _ = "STUB: not implemented"; return false }

func (m *MigrationOptions) GetLockRepositories() bool { _ = "STUB: not implemented"; return false }

func (m *Milestone) GetClosedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (m *Milestone) GetClosedIssues() int { _ = "STUB: not implemented"; return 0 }

func (m *Milestone) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (m *Milestone) GetCreator() *User { _ = "STUB: not implemented"; return nil }

func (m *Milestone) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (m *Milestone) GetDueOn() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (m *Milestone) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (m *Milestone) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (m *Milestone) GetLabelsURL() string { _ = "STUB: not implemented"; return "" }

func (m *Milestone) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (m *Milestone) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (m *Milestone) GetOpenIssues() int { _ = "STUB: not implemented"; return 0 }

func (m *Milestone) GetState() string { _ = "STUB: not implemented"; return "" }

func (m *Milestone) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (m *Milestone) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (m *Milestone) GetURL() string { _ = "STUB: not implemented"; return "" }

func (m *MilestoneEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (m *MilestoneEvent) GetChanges() *EditChange { _ = "STUB: not implemented"; return nil }

func (m *MilestoneEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (m *MilestoneEvent) GetMilestone() *Milestone { _ = "STUB: not implemented"; return nil }

func (m *MilestoneEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (m *MilestoneEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (m *MilestoneEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (m *MilestoneListOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (m *MilestoneListOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (m *MilestoneListOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (m *MilestoneStats) GetClosedMilestones() int { _ = "STUB: not implemented"; return 0 }

func (m *MilestoneStats) GetOpenMilestones() int { _ = "STUB: not implemented"; return 0 }

func (m *MilestoneStats) GetTotalMilestones() int { _ = "STUB: not implemented"; return 0 }

func (m *MostRecentInstance) GetAnalysisKey() string { _ = "STUB: not implemented"; return "" }

func (m *MostRecentInstance) GetCategory() string { _ = "STUB: not implemented"; return "" }

func (m *MostRecentInstance) GetClassifications() []string { _ = "STUB: not implemented"; return nil }

func (m *MostRecentInstance) GetCommitSHA() string { _ = "STUB: not implemented"; return "" }

func (m *MostRecentInstance) GetEnvironment() string { _ = "STUB: not implemented"; return "" }

func (m *MostRecentInstance) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (m *MostRecentInstance) GetLocation() *Location { _ = "STUB: not implemented"; return nil }

func (m *MostRecentInstance) GetMessage() *Message { _ = "STUB: not implemented"; return nil }

func (m *MostRecentInstance) GetRef() string { _ = "STUB: not implemented"; return "" }

func (m *MostRecentInstance) GetState() string { _ = "STUB: not implemented"; return "" }

func (n *NetworkConfiguration) GetComputeService() *ComputeService {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkConfiguration) GetCreatedOn() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (n *NetworkConfiguration) GetID() string { _ = "STUB: not implemented"; return "" }

func (n *NetworkConfiguration) GetName() string { _ = "STUB: not implemented"; return "" }

func (n *NetworkConfiguration) GetNetworkSettingsIDs() []string {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkConfigurationRequest) GetComputeService() *ComputeService {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkConfigurationRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (n *NetworkConfigurationRequest) GetNetworkSettingsIDs() []string {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkConfigurations) GetNetworkConfigurations() []*NetworkConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkConfigurations) GetTotalCount() int64 { _ = "STUB: not implemented"; return 0 }

func (n *NetworkSettingsResource) GetID() string { _ = "STUB: not implemented"; return "" }

func (n *NetworkSettingsResource) GetName() string { _ = "STUB: not implemented"; return "" }

func (n *NetworkSettingsResource) GetNetworkConfigurationID() string {
	_ = "STUB: not implemented"
	return ""
}

func (n *NetworkSettingsResource) GetRegion() string { _ = "STUB: not implemented"; return "" }

func (n *NetworkSettingsResource) GetSubnetID() string { _ = "STUB: not implemented"; return "" }

func (n *NewPullRequest) GetBase() string { _ = "STUB: not implemented"; return "" }

func (n *NewPullRequest) GetBody() string { _ = "STUB: not implemented"; return "" }

func (n *NewPullRequest) GetDraft() bool { _ = "STUB: not implemented"; return false }

func (n *NewPullRequest) GetHead() string { _ = "STUB: not implemented"; return "" }

func (n *NewPullRequest) GetHeadRepo() string { _ = "STUB: not implemented"; return "" }

func (n *NewPullRequest) GetIssue() int { _ = "STUB: not implemented"; return 0 }

func (n *NewPullRequest) GetMaintainerCanModify() bool { _ = "STUB: not implemented"; return false }

func (n *NewPullRequest) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (n *NewTeam) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (n *NewTeam) GetLDAPDN() string { _ = "STUB: not implemented"; return "" }

func (n *NewTeam) GetMaintainers() []string { _ = "STUB: not implemented"; return nil }

func (n *NewTeam) GetName() string { _ = "STUB: not implemented"; return "" }

func (n *NewTeam) GetNotificationSetting() string { _ = "STUB: not implemented"; return "" }

func (n *NewTeam) GetParentTeamID() int64 { _ = "STUB: not implemented"; return 0 }

func (n *NewTeam) GetPermission() string { _ = "STUB: not implemented"; return "" }

func (n *NewTeam) GetPrivacy() string { _ = "STUB: not implemented"; return "" }

func (n *NewTeam) GetRepoNames() []string { _ = "STUB: not implemented"; return nil }

func (n *NodeDetails) GetClusterRoles() []string { _ = "STUB: not implemented"; return nil }

func (n *NodeDetails) GetHostname() string { _ = "STUB: not implemented"; return "" }

func (n *NodeDetails) GetUUID() string { _ = "STUB: not implemented"; return "" }

func (n *NodeMetadataStatus) GetNodes() []*NodeDetails { _ = "STUB: not implemented"; return nil }

func (n *NodeMetadataStatus) GetTopology() string { _ = "STUB: not implemented"; return "" }

func (n *NodeQueryOptions) GetClusterRoles() string { _ = "STUB: not implemented"; return "" }

func (n *NodeQueryOptions) GetUUID() string { _ = "STUB: not implemented"; return "" }

func (n *NodeReleaseVersion) GetHostname() string { _ = "STUB: not implemented"; return "" }

func (n *NodeReleaseVersion) GetVersion() *ReleaseVersion { _ = "STUB: not implemented"; return nil }

func (n *Notification) GetID() string { _ = "STUB: not implemented"; return "" }

func (n *Notification) GetLastReadAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (n *Notification) GetReason() string { _ = "STUB: not implemented"; return "" }

func (n *Notification) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (n *Notification) GetSubject() *NotificationSubject { _ = "STUB: not implemented"; return nil }

func (n *Notification) GetUnread() bool { _ = "STUB: not implemented"; return false }

func (n *Notification) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (n *Notification) GetURL() string { _ = "STUB: not implemented"; return "" }

func (n *NotificationListOptions) GetAll() bool { _ = "STUB: not implemented"; return false }

func (n *NotificationListOptions) GetBefore() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (n *NotificationListOptions) GetParticipating() bool { _ = "STUB: not implemented"; return false }

func (n *NotificationListOptions) GetSince() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (n *NotificationSubject) GetLatestCommentURL() string { _ = "STUB: not implemented"; return "" }

func (n *NotificationSubject) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (n *NotificationSubject) GetType() string { _ = "STUB: not implemented"; return "" }

func (n *NotificationSubject) GetURL() string { _ = "STUB: not implemented"; return "" }

func (o *OAuthAPP) GetClientID() string { _ = "STUB: not implemented"; return "" }

func (o *OAuthAPP) GetName() string { _ = "STUB: not implemented"; return "" }

func (o *OAuthAPP) GetURL() string { _ = "STUB: not implemented"; return "" }

func (o *OIDCSubjectClaimCustomTemplate) GetIncludeClaimKeys() []string {
	_ = "STUB: not implemented"
	return nil
}

func (o *OIDCSubjectClaimCustomTemplate) GetSubClaimPrefix() string {
	_ = "STUB: not implemented"
	return ""
}

func (o *OIDCSubjectClaimCustomTemplate) GetUseDefault() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *OIDCSubjectClaimCustomTemplate) GetUseImmutableSubject() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetAdvancedSecurityEnabledForNewRepos() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetArchivedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (o *Organization) GetAvatarURL() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetBillingEmail() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetBlog() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetCollaborators() int { _ = "STUB: not implemented"; return 0 }

func (o *Organization) GetCompany() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (o *Organization) GetDefaultRepoPermission() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetDefaultRepoSettings() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetDefaultRepositoryBranch() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetDependabotAlertsEnabledForNewRepos() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetDependabotSecurityUpdatesEnabledForNewRepos() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetDependencyGraphEnabledForNewRepos() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetDeployKeysEnabledForRepositories() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetDiskUsage() int { _ = "STUB: not implemented"; return 0 }

func (o *Organization) GetDisplayCommenterFullNameSettingEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetEventsURL() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetFollowers() int { _ = "STUB: not implemented"; return 0 }

func (o *Organization) GetFollowing() int { _ = "STUB: not implemented"; return 0 }

func (o *Organization) GetHasOrganizationProjects() bool { _ = "STUB: not implemented"; return false }

func (o *Organization) GetHasRepositoryProjects() bool { _ = "STUB: not implemented"; return false }

func (o *Organization) GetHooksURL() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (o *Organization) GetIssuesURL() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetIsVerified() bool { _ = "STUB: not implemented"; return false }

func (o *Organization) GetLocation() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetMembersAllowedRepositoryCreationType() string {
	_ = "STUB: not implemented"
	return ""
}

func (o *Organization) GetMembersCanChangeRepoVisibility() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetMembersCanCreateInternalRepos() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetMembersCanCreatePages() bool { _ = "STUB: not implemented"; return false }

func (o *Organization) GetMembersCanCreatePrivatePages() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetMembersCanCreatePrivateRepos() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetMembersCanCreatePublicPages() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetMembersCanCreatePublicRepos() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetMembersCanCreateRepos() bool { _ = "STUB: not implemented"; return false }

func (o *Organization) GetMembersCanCreateTeams() bool { _ = "STUB: not implemented"; return false }

func (o *Organization) GetMembersCanDeleteIssues() bool { _ = "STUB: not implemented"; return false }

func (o *Organization) GetMembersCanDeleteRepositories() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetMembersCanForkPrivateRepos() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetMembersCanInviteOutsideCollaborators() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetMembersCanViewDependencyInsights() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetMembersURL() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetName() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetOwnedPrivateRepos() int64 { _ = "STUB: not implemented"; return 0 }

func (o *Organization) GetPlan() *Plan { _ = "STUB: not implemented"; return nil }

func (o *Organization) GetPrivateGists() int { _ = "STUB: not implemented"; return 0 }

func (o *Organization) GetPublicGists() int { _ = "STUB: not implemented"; return 0 }

func (o *Organization) GetPublicMembersURL() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetPublicRepos() int { _ = "STUB: not implemented"; return 0 }

func (o *Organization) GetReadersCanCreateDiscussions() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetReposURL() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetSecretScanningEnabledForNewRepos() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetSecretScanningPushProtectionCustomLink() string {
	_ = "STUB: not implemented"
	return ""
}

func (o *Organization) GetSecretScanningPushProtectionCustomLinkEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetSecretScanningPushProtectionEnabledForNewRepos() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetSecretScanningValidityChecksEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetTotalPrivateRepos() int64 { _ = "STUB: not implemented"; return 0 }

func (o *Organization) GetTwitterUsername() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetTwoFactorRequirementEnabled() bool {
	_ = "STUB: not implemented"
	return false
}

func (o *Organization) GetType() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (o *Organization) GetURL() string { _ = "STUB: not implemented"; return "" }

func (o *Organization) GetWebCommitSignoffRequired() bool { _ = "STUB: not implemented"; return false }

func (o *OrganizationCustomPropertyValues) GetProperties() []*CustomPropertyValue {
	_ = "STUB: not implemented"
	return nil
}

func (o *OrganizationCustomRepoRoles) GetCustomRepoRoles() []*CustomRepoRoles {
	_ = "STUB: not implemented"
	return nil
}

func (o *OrganizationCustomRepoRoles) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (o *OrganizationCustomRoles) GetCustomRepoRoles() []*CustomOrgRole {
	_ = "STUB: not implemented"
	return nil
}

func (o *OrganizationCustomRoles) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (o *OrganizationEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (o *OrganizationEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (o *OrganizationEvent) GetInvitation() *Invitation { _ = "STUB: not implemented"; return nil }

func (o *OrganizationEvent) GetMembership() *Membership { _ = "STUB: not implemented"; return nil }

func (o *OrganizationEvent) GetOrganization() *Organization { _ = "STUB: not implemented"; return nil }

func (o *OrganizationEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (o *OrganizationFineGrainedPermission) GetDescription() string {
	_ = "STUB: not implemented"
	return ""
}

func (o *OrganizationFineGrainedPermission) GetName() string { _ = "STUB: not implemented"; return "" }

func (o *OrganizationInstallations) GetInstallations() []*Installation {
	_ = "STUB: not implemented"
	return nil
}

func (o *OrganizationInstallations) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (o *OrganizationsListOptions) GetPerPage() int { _ = "STUB: not implemented"; return 0 }

func (o *OrganizationsListOptions) GetSince() int64 { _ = "STUB: not implemented"; return 0 }

func (o *OrgBlockEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (o *OrgBlockEvent) GetBlockedUser() *User { _ = "STUB: not implemented"; return nil }

func (o *OrgBlockEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (o *OrgBlockEvent) GetOrganization() *Organization { _ = "STUB: not implemented"; return nil }

func (o *OrgBlockEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (o *OrgStats) GetDisabledOrgs() int { _ = "STUB: not implemented"; return 0 }

func (o *OrgStats) GetTotalOrgs() int { _ = "STUB: not implemented"; return 0 }

func (o *OrgStats) GetTotalTeamMembers() int { _ = "STUB: not implemented"; return 0 }

func (o *OrgStats) GetTotalTeams() int { _ = "STUB: not implemented"; return 0 }

func (o *OwnerInfo) GetOrg() *User { _ = "STUB: not implemented"; return nil }

func (o *OwnerInfo) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (p *Package) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *Package) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *Package) GetEcosystem() string { _ = "STUB: not implemented"; return "" }

func (p *Package) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *Package) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *Package) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *Package) GetNamespace() string { _ = "STUB: not implemented"; return "" }

func (p *Package) GetOwner() *User { _ = "STUB: not implemented"; return nil }

func (p *Package) GetPackageType() string { _ = "STUB: not implemented"; return "" }

func (p *Package) GetPackageVersion() *PackageVersion { _ = "STUB: not implemented"; return nil }

func (p *Package) GetRegistry() *PackageRegistry { _ = "STUB: not implemented"; return nil }

func (p *Package) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (p *Package) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *Package) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *Package) GetVersionCount() int64 { _ = "STUB: not implemented"; return 0 }

func (p *Package) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (p *PackageContainerMetadata) GetTags() []string { _ = "STUB: not implemented"; return nil }

func (p *PackageEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (p *PackageEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (p *PackageEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (p *PackageEvent) GetPackage() *Package { _ = "STUB: not implemented"; return nil }

func (p *PackageEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (p *PackageEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *PackageEventContainerMetadata) GetLabels() map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageEventContainerMetadata) GetManifest() map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageEventContainerMetadata) GetTag() *PackageEventContainerMetadataTag {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageEventContainerMetadataTag) GetDigest() string { _ = "STUB: not implemented"; return "" }

func (p *PackageEventContainerMetadataTag) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PackageExternalRef) GetReferenceCategory() string { _ = "STUB: not implemented"; return "" }

func (p *PackageExternalRef) GetReferenceLocator() string { _ = "STUB: not implemented"; return "" }

func (p *PackageExternalRef) GetReferenceType() string { _ = "STUB: not implemented"; return "" }

func (p *PackageFile) GetAuthor() *User { _ = "STUB: not implemented"; return nil }

func (p *PackageFile) GetContentType() string { _ = "STUB: not implemented"; return "" }

func (p *PackageFile) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *PackageFile) GetDownloadURL() string { _ = "STUB: not implemented"; return "" }

func (p *PackageFile) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PackageFile) GetMD5() string { _ = "STUB: not implemented"; return "" }

func (p *PackageFile) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PackageFile) GetSHA1() string { _ = "STUB: not implemented"; return "" }

func (p *PackageFile) GetSHA256() string { _ = "STUB: not implemented"; return "" }

func (p *PackageFile) GetSize() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PackageFile) GetState() string { _ = "STUB: not implemented"; return "" }

func (p *PackageFile) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *PackageListOptions) GetPackageType() string { _ = "STUB: not implemented"; return "" }

func (p *PackageListOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (p *PackageListOptions) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (p *PackageMetadata) GetContainer() *PackageContainerMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageMetadata) GetPackageType() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetAuthor() map[string]string { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetBin() map[string]any { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetBugs() map[string]string { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetCommitOID() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetContributors() []any { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetCPU() []string { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetDeletedByID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PackageNPMMetadata) GetDependencies() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageNPMMetadata) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetDevDependencies() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageNPMMetadata) GetDirectories() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageNPMMetadata) GetDist() map[string]string { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetEngines() map[string]string { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetFiles() []string { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetGitHead() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetHasShrinkwrap() bool { _ = "STUB: not implemented"; return false }

func (p *PackageNPMMetadata) GetHomepage() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetID() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetInstallationCommand() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetKeywords() []string { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetLicense() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetMain() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetMaintainers() []any { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetMan() map[string]any { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetNodeVersion() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetNPMUser() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetNPMVersion() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetOptionalDependencies() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageNPMMetadata) GetOS() []string { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetPeerDependencies() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageNPMMetadata) GetPublishedViaActions() bool { _ = "STUB: not implemented"; return false }

func (p *PackageNPMMetadata) GetReadme() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNPMMetadata) GetReleaseID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PackageNPMMetadata) GetRepository() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageNPMMetadata) GetScripts() map[string]any { _ = "STUB: not implemented"; return nil }

func (p *PackageNPMMetadata) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNugetMetadata) GetID() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (p *PackageNugetMetadata) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PackageNugetMetadata) GetValue() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (p *PackageRegistry) GetAboutURL() string { _ = "STUB: not implemented"; return "" }

func (p *PackageRegistry) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PackageRegistry) GetType() string { _ = "STUB: not implemented"; return "" }

func (p *PackageRegistry) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *PackageRegistry) GetVendor() string { _ = "STUB: not implemented"; return "" }

func (p *PackageRelease) GetAuthor() *User { _ = "STUB: not implemented"; return nil }

func (p *PackageRelease) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PackageRelease) GetDraft() bool { _ = "STUB: not implemented"; return false }

func (p *PackageRelease) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *PackageRelease) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PackageRelease) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PackageRelease) GetPrerelease() bool { _ = "STUB: not implemented"; return false }

func (p *PackageRelease) GetPublishedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PackageRelease) GetTagName() string { _ = "STUB: not implemented"; return "" }

func (p *PackageRelease) GetTargetCommitish() string { _ = "STUB: not implemented"; return "" }

func (p *PackageRelease) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *PackagesBilling) GetIncludedGigabytesBandwidth() int { _ = "STUB: not implemented"; return 0 }

func (p *PackagesBilling) GetTotalGigabytesBandwidthUsed() int { _ = "STUB: not implemented"; return 0 }

func (p *PackagesBilling) GetTotalPaidGigabytesBandwidthUsed() int {
	_ = "STUB: not implemented"
	return 0
}

func (p *PackageVersion) GetAuthor() *User { _ = "STUB: not implemented"; return nil }

func (p *PackageVersion) GetBodyHTML() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetContainerMetadata() *PackageEventContainerMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageVersion) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PackageVersion) GetDeletedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PackageVersion) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetDockerMetadata() []any { _ = "STUB: not implemented"; return nil }

func (p *PackageVersion) GetDraft() bool { _ = "STUB: not implemented"; return false }

func (p *PackageVersion) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PackageVersion) GetInstallationCommand() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetLicense() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetManifest() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetNPMMetadata() *PackageNPMMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageVersion) GetNugetMetadata() []*PackageNugetMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageVersion) GetPackageFiles() []*PackageFile { _ = "STUB: not implemented"; return nil }

func (p *PackageVersion) GetPackageHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetPackageURL() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetPrerelease() bool { _ = "STUB: not implemented"; return false }

func (p *PackageVersion) GetRelease() *PackageRelease { _ = "STUB: not implemented"; return nil }

func (p *PackageVersion) GetRubyMetadata() map[string]any { _ = "STUB: not implemented"; return nil }

func (p *PackageVersion) GetSourceURL() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetSummary() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetTagName() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetTargetCommitish() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetTargetOID() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PackageVersion) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersion) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersionBody) GetInfo() *PackageVersionBodyInfo {
	_ = "STUB: not implemented"
	return nil
}

func (p *PackageVersionBody) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (p *PackageVersionBodyInfo) GetCollection() bool { _ = "STUB: not implemented"; return false }

func (p *PackageVersionBodyInfo) GetMode() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PackageVersionBodyInfo) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersionBodyInfo) GetOID() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersionBodyInfo) GetPath() string { _ = "STUB: not implemented"; return "" }

func (p *PackageVersionBodyInfo) GetSize() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PackageVersionBodyInfo) GetType() string { _ = "STUB: not implemented"; return "" }

func (p *Page) GetAction() string { _ = "STUB: not implemented"; return "" }

func (p *Page) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *Page) GetPageName() string { _ = "STUB: not implemented"; return "" }

func (p *Page) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (p *Page) GetSummary() string { _ = "STUB: not implemented"; return "" }

func (p *Page) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *PageBuildEvent) GetBuild() *PagesBuild { _ = "STUB: not implemented"; return nil }

func (p *PageBuildEvent) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PageBuildEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (p *PageBuildEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (p *PageBuildEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (p *PageBuildEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *Pages) GetBuildType() string { _ = "STUB: not implemented"; return "" }

func (p *Pages) GetCNAME() string { _ = "STUB: not implemented"; return "" }

func (p *Pages) GetCustom404() bool { _ = "STUB: not implemented"; return false }

func (p *Pages) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *Pages) GetHTTPSCertificate() *PagesHTTPSCertificate { _ = "STUB: not implemented"; return nil }

func (p *Pages) GetHTTPSEnforced() bool { _ = "STUB: not implemented"; return false }

func (p *Pages) GetPublic() bool { _ = "STUB: not implemented"; return false }

func (p *Pages) GetSource() *PagesSource { _ = "STUB: not implemented"; return nil }

func (p *Pages) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (p *Pages) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *PagesBuild) GetCommit() string { _ = "STUB: not implemented"; return "" }

func (p *PagesBuild) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *PagesBuild) GetDuration() int { _ = "STUB: not implemented"; return 0 }

func (p *PagesBuild) GetError() *PagesError { _ = "STUB: not implemented"; return nil }

func (p *PagesBuild) GetPusher() *User { _ = "STUB: not implemented"; return nil }

func (p *PagesBuild) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (p *PagesBuild) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *PagesBuild) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *PagesDomain) GetCAAError() string { _ = "STUB: not implemented"; return "" }

func (p *PagesDomain) GetDNSResolves() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetEnforcesHTTPS() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetHasCNAMERecord() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetHasMXRecordsPresent() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetHost() string { _ = "STUB: not implemented"; return "" }

func (p *PagesDomain) GetHTTPSError() string { _ = "STUB: not implemented"; return "" }

func (p *PagesDomain) GetIsApexDomain() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsARecord() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsCloudflareIP() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsCNAMEToFastly() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsCNAMEToGithubUserDomain() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsCNAMEToPagesDotGithubDotCom() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PagesDomain) GetIsFastlyIP() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsHTTPSEligible() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsNonGithubPagesIPPresent() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsOldIPAddress() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsPagesDomain() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsPointedToGithubPagesIP() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsProxied() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsServedByPages() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsValid() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetIsValidDomain() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetNameservers() string { _ = "STUB: not implemented"; return "" }

func (p *PagesDomain) GetReason() string { _ = "STUB: not implemented"; return "" }

func (p *PagesDomain) GetRespondsToHTTPS() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetShouldBeARecord() bool { _ = "STUB: not implemented"; return false }

func (p *PagesDomain) GetURI() string { _ = "STUB: not implemented"; return "" }

func (p *PagesError) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (p *PagesHealthCheckResponse) GetAltDomain() *PagesDomain {
	_ = "STUB: not implemented"
	return nil
}

func (p *PagesHealthCheckResponse) GetDomain() *PagesDomain { _ = "STUB: not implemented"; return nil }

func (p *PagesHTTPSCertificate) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *PagesHTTPSCertificate) GetDomains() []string { _ = "STUB: not implemented"; return nil }

func (p *PagesHTTPSCertificate) GetExpiresAt() string { _ = "STUB: not implemented"; return "" }

func (p *PagesHTTPSCertificate) GetState() string { _ = "STUB: not implemented"; return "" }

func (p *PagesSource) GetBranch() string { _ = "STUB: not implemented"; return "" }

func (p *PagesSource) GetPath() string { _ = "STUB: not implemented"; return "" }

func (p *PageStats) GetTotalPages() int { _ = "STUB: not implemented"; return 0 }

func (p *PagesUpdate) GetBuildType() string { _ = "STUB: not implemented"; return "" }

func (p *PagesUpdate) GetCNAME() string { _ = "STUB: not implemented"; return "" }

func (p *PagesUpdate) GetHTTPSEnforced() bool { _ = "STUB: not implemented"; return false }

func (p *PagesUpdate) GetPublic() bool { _ = "STUB: not implemented"; return false }

func (p *PagesUpdate) GetSource() *PagesSource { _ = "STUB: not implemented"; return nil }

func (p *PagesUpdateWithoutCNAME) GetBuildType() string { _ = "STUB: not implemented"; return "" }

func (p *PagesUpdateWithoutCNAME) GetHTTPSEnforced() bool { _ = "STUB: not implemented"; return false }

func (p *PagesUpdateWithoutCNAME) GetPublic() bool { _ = "STUB: not implemented"; return false }

func (p *PagesUpdateWithoutCNAME) GetSource() *PagesSource { _ = "STUB: not implemented"; return nil }

func (p *PatternBranchRule) GetParameters() PatternRuleParameters {
	_ = "STUB: not implemented"
	return *new(PatternRuleParameters)
}

func (p *PatternRuleParameters) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PatternRuleParameters) GetNegate() bool { _ = "STUB: not implemented"; return false }

func (p *PatternRuleParameters) GetOperator() PatternRuleOperator {
	_ = "STUB: not implemented"
	return *new(PatternRuleOperator)
}

func (p *PatternRuleParameters) GetPattern() string { _ = "STUB: not implemented"; return "" }

func (p *PendingDeployment) GetCurrentUserCanApprove() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PendingDeployment) GetEnvironment() *PendingDeploymentEnvironment {
	_ = "STUB: not implemented"
	return nil
}

func (p *PendingDeployment) GetReviewers() []*RequiredReviewer {
	_ = "STUB: not implemented"
	return nil
}

func (p *PendingDeployment) GetWaitTimer() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PendingDeployment) GetWaitTimerStartedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PendingDeploymentEnvironment) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *PendingDeploymentEnvironment) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PendingDeploymentEnvironment) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PendingDeploymentEnvironment) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *PendingDeploymentEnvironment) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *PendingDeploymentsRequest) GetComment() string { _ = "STUB: not implemented"; return "" }

func (p *PendingDeploymentsRequest) GetEnvironmentIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (p *PendingDeploymentsRequest) GetState() string { _ = "STUB: not implemented"; return "" }

func (p *PersonalAccessToken) GetAccessGrantedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PersonalAccessToken) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PersonalAccessToken) GetOwner() *User { _ = "STUB: not implemented"; return nil }

func (p *PersonalAccessToken) GetPermissions() *PersonalAccessTokenPermissions {
	_ = "STUB: not implemented"
	return nil
}

func (p *PersonalAccessToken) GetRepositoriesURL() string { _ = "STUB: not implemented"; return "" }

func (p *PersonalAccessToken) GetRepositorySelection() string { _ = "STUB: not implemented"; return "" }

func (p *PersonalAccessToken) GetTokenExpired() bool { _ = "STUB: not implemented"; return false }

func (p *PersonalAccessToken) GetTokenExpiresAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PersonalAccessToken) GetTokenID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PersonalAccessToken) GetTokenLastUsedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PersonalAccessToken) GetTokenName() string { _ = "STUB: not implemented"; return "" }

func (p *PersonalAccessTokenPermissions) GetOrg() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *PersonalAccessTokenPermissions) GetOther() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *PersonalAccessTokenPermissions) GetRepo() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (p *PersonalAccessTokenRequest) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PersonalAccessTokenRequest) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PersonalAccessTokenRequest) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (p *PersonalAccessTokenRequest) GetOwner() *User { _ = "STUB: not implemented"; return nil }

func (p *PersonalAccessTokenRequest) GetPermissionsAdded() *PersonalAccessTokenPermissions {
	_ = "STUB: not implemented"
	return nil
}

func (p *PersonalAccessTokenRequest) GetPermissionsResult() *PersonalAccessTokenPermissions {
	_ = "STUB: not implemented"
	return nil
}

func (p *PersonalAccessTokenRequest) GetPermissionsUpgraded() *PersonalAccessTokenPermissions {
	_ = "STUB: not implemented"
	return nil
}

func (p *PersonalAccessTokenRequest) GetRepositories() []*Repository {
	_ = "STUB: not implemented"
	return nil
}

func (p *PersonalAccessTokenRequest) GetRepositoryCount() int64 {
	_ = "STUB: not implemented"
	return 0
}

func (p *PersonalAccessTokenRequest) GetRepositorySelection() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *PersonalAccessTokenRequest) GetTokenExpired() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PersonalAccessTokenRequest) GetTokenExpiresAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PersonalAccessTokenRequest) GetTokenLastUsedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PersonalAccessTokenRequestEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (p *PersonalAccessTokenRequestEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (p *PersonalAccessTokenRequestEvent) GetOrg() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (p *PersonalAccessTokenRequestEvent) GetPersonalAccessTokenRequest() *PersonalAccessTokenRequest {
	_ = "STUB: not implemented"
	return nil
}

func (p *PersonalAccessTokenRequestEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *PingEvent) GetHook() *Hook { _ = "STUB: not implemented"; return nil }

func (p *PingEvent) GetHookID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PingEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (p *PingEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (p *PingEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (p *PingEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *PingEvent) GetZen() string { _ = "STUB: not implemented"; return "" }

func (p *Plan) GetCollaborators() int { _ = "STUB: not implemented"; return 0 }

func (p *Plan) GetFilledSeats() int { _ = "STUB: not implemented"; return 0 }

func (p *Plan) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *Plan) GetPrivateRepos() int64 { _ = "STUB: not implemented"; return 0 }

func (p *Plan) GetSeats() int { _ = "STUB: not implemented"; return 0 }

func (p *Plan) GetSpace() int { _ = "STUB: not implemented"; return 0 }

func (p *PreferenceList) GetAutoTriggerChecks() []*AutoTriggerCheck {
	_ = "STUB: not implemented"
	return nil
}

func (p *PremiumRequestUsageItem) GetDiscountAmount() float64 { _ = "STUB: not implemented"; return 0 }

func (p *PremiumRequestUsageItem) GetDiscountQuantity() float64 {
	_ = "STUB: not implemented"
	return 0
}

func (p *PremiumRequestUsageItem) GetGrossAmount() float64 { _ = "STUB: not implemented"; return 0 }

func (p *PremiumRequestUsageItem) GetGrossQuantity() float64 { _ = "STUB: not implemented"; return 0 }

func (p *PremiumRequestUsageItem) GetModel() string { _ = "STUB: not implemented"; return "" }

func (p *PremiumRequestUsageItem) GetNetAmount() float64 { _ = "STUB: not implemented"; return 0 }

func (p *PremiumRequestUsageItem) GetNetQuantity() float64 { _ = "STUB: not implemented"; return 0 }

func (p *PremiumRequestUsageItem) GetPricePerUnit() float64 { _ = "STUB: not implemented"; return 0 }

func (p *PremiumRequestUsageItem) GetProduct() string { _ = "STUB: not implemented"; return "" }

func (p *PremiumRequestUsageItem) GetSKU() string { _ = "STUB: not implemented"; return "" }

func (p *PremiumRequestUsageItem) GetUnitType() string { _ = "STUB: not implemented"; return "" }

func (p *PremiumRequestUsageReport) GetModel() string { _ = "STUB: not implemented"; return "" }

func (p *PremiumRequestUsageReport) GetOrganization() string { _ = "STUB: not implemented"; return "" }

func (p *PremiumRequestUsageReport) GetProduct() string { _ = "STUB: not implemented"; return "" }

func (p *PremiumRequestUsageReport) GetTimePeriod() PremiumRequestUsageTimePeriod {
	_ = "STUB: not implemented"
	return *new(PremiumRequestUsageTimePeriod)
}

func (p *PremiumRequestUsageReport) GetUsageItems() []*PremiumRequestUsageItem {
	_ = "STUB: not implemented"
	return nil
}

func (p *PremiumRequestUsageReport) GetUser() string { _ = "STUB: not implemented"; return "" }

func (p *PremiumRequestUsageReportOptions) GetDay() int { _ = "STUB: not implemented"; return 0 }

func (p *PremiumRequestUsageReportOptions) GetModel() string { _ = "STUB: not implemented"; return "" }

func (p *PremiumRequestUsageReportOptions) GetMonth() int { _ = "STUB: not implemented"; return 0 }

func (p *PremiumRequestUsageReportOptions) GetProduct() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *PremiumRequestUsageReportOptions) GetUser() string { _ = "STUB: not implemented"; return "" }

func (p *PremiumRequestUsageReportOptions) GetYear() int { _ = "STUB: not implemented"; return 0 }

func (p *PremiumRequestUsageTimePeriod) GetDay() int { _ = "STUB: not implemented"; return 0 }

func (p *PremiumRequestUsageTimePeriod) GetMonth() int { _ = "STUB: not implemented"; return 0 }

func (p *PremiumRequestUsageTimePeriod) GetYear() int { _ = "STUB: not implemented"; return 0 }

func (p *PreReceiveHook) GetConfigURL() string { _ = "STUB: not implemented"; return "" }

func (p *PreReceiveHook) GetEnforcement() string { _ = "STUB: not implemented"; return "" }

func (p *PreReceiveHook) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PreReceiveHook) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistries) GetConfigurations() []*PrivateRegistry {
	_ = "STUB: not implemented"
	return nil
}

func (p *PrivateRegistries) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (p *PrivateRegistry) GetAccountID() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetAudience() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetAuthType() *PrivateRegistryAuthType {
	_ = "STUB: not implemented"
	return nil
}

func (p *PrivateRegistry) GetAWSRegion() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetClientID() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PrivateRegistry) GetDomain() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetDomainOwner() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetIdentityMappingName() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetJFrogOIDCProviderName() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetRegistryType() *PrivateRegistryType {
	_ = "STUB: not implemented"
	return nil
}

func (p *PrivateRegistry) GetReplacesBase() bool { _ = "STUB: not implemented"; return false }

func (p *PrivateRegistry) GetRoleName() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetSelectedRepositoryIDs() []int64 { _ = "STUB: not implemented"; return nil }

func (p *PrivateRegistry) GetTenantID() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PrivateRegistry) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetUsername() string { _ = "STUB: not implemented"; return "" }

func (p *PrivateRegistry) GetVisibility() *PrivateRegistryVisibility {
	_ = "STUB: not implemented"
	return nil
}

func (p *PRLink) GetHRef() string { _ = "STUB: not implemented"; return "" }

func (p *PRLinks) GetComments() *PRLink { _ = "STUB: not implemented"; return nil }

func (p *PRLinks) GetCommits() *PRLink { _ = "STUB: not implemented"; return nil }

func (p *PRLinks) GetHTML() *PRLink { _ = "STUB: not implemented"; return nil }

func (p *PRLinks) GetIssue() *PRLink { _ = "STUB: not implemented"; return nil }

func (p *PRLinks) GetReviewComment() *PRLink { _ = "STUB: not implemented"; return nil }

func (p *PRLinks) GetReviewComments() *PRLink { _ = "STUB: not implemented"; return nil }

func (p *PRLinks) GetSelf() *PRLink { _ = "STUB: not implemented"; return nil }

func (p *PRLinks) GetStatuses() *PRLink { _ = "STUB: not implemented"; return nil }

func (p *ProjectBody) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectCardChange) GetNote() *ProjectCardNote { _ = "STUB: not implemented"; return nil }

func (p *ProjectCardNote) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectChange) GetBody() *ProjectBody { _ = "STUB: not implemented"; return nil }

func (p *ProjectChange) GetName() *ProjectName { _ = "STUB: not implemented"; return nil }

func (p *ProjectColumnChange) GetName() *ProjectColumnName { _ = "STUB: not implemented"; return nil }

func (p *ProjectColumnName) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectName) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2) GetBody() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2) GetClosedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *ProjectV2) GetColumnsURL() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *ProjectV2) GetCreator() *User { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2) GetDeletedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *ProjectV2) GetDeletedBy() *User { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *ProjectV2) GetIsTemplate() bool { _ = "STUB: not implemented"; return false }

func (p *ProjectV2) GetLatestStatusUpdate() *ProjectV2StatusUpdate {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (p *ProjectV2) GetOrganizationPermission() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2) GetOwner() *User { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2) GetOwnerURL() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2) GetPrivate() bool { _ = "STUB: not implemented"; return false }

func (p *ProjectV2) GetPublic() bool { _ = "STUB: not implemented"; return false }

func (p *ProjectV2) GetShortDescription() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2) GetState() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *ProjectV2) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2DraftIssue) GetBody() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2DraftIssue) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *ProjectV2DraftIssue) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *ProjectV2DraftIssue) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2DraftIssue) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2DraftIssue) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *ProjectV2DraftIssue) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2Event) GetAction() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2Event) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2Event) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2Event) GetProjectsV2() *ProjectV2 { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2Event) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2Field) GetConfiguration() *ProjectV2FieldConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2Field) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *ProjectV2Field) GetDataType() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2Field) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *ProjectV2Field) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2Field) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2Field) GetOptions() []*ProjectV2FieldOption {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2Field) GetProjectURL() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2Field) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *ProjectV2FieldConfiguration) GetDuration() int { _ = "STUB: not implemented"; return 0 }

func (p *ProjectV2FieldConfiguration) GetIterations() []*ProjectV2FieldIteration {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2FieldConfiguration) GetStartDay() int { _ = "STUB: not implemented"; return 0 }

func (p *ProjectV2FieldIteration) GetDuration() int { _ = "STUB: not implemented"; return 0 }

func (p *ProjectV2FieldIteration) GetID() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2FieldIteration) GetStartDate() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2FieldIteration) GetTitle() *ProjectV2TextContent {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2FieldIterationConfiguration) GetDuration() int {
	_ = "STUB: not implemented"
	return 0
}

func (p *ProjectV2FieldIterationConfiguration) GetIterations() []*ProjectV2FieldIterationConfigurationIteration {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2FieldIterationConfiguration) GetStartDate() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *ProjectV2FieldIterationConfigurationIteration) GetDuration() int {
	_ = "STUB: not implemented"
	return 0
}

func (p *ProjectV2FieldIterationConfigurationIteration) GetStartDate() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *ProjectV2FieldIterationConfigurationIteration) GetTitle() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *ProjectV2FieldOption) GetColor() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2FieldOption) GetDescription() *ProjectV2TextContent {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2FieldOption) GetID() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2FieldOption) GetName() *ProjectV2TextContent {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2FieldSingleSelectOption) GetColor() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2FieldSingleSelectOption) GetDescription() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *ProjectV2FieldSingleSelectOption) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2Item) GetArchivedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *ProjectV2Item) GetContent() *ProjectV2ItemContent { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2Item) GetContentNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2Item) GetContentType() *ProjectV2ItemContentType {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2Item) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *ProjectV2Item) GetCreator() *User { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2Item) GetFields() []*ProjectV2ItemFieldValue {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2Item) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *ProjectV2Item) GetItemURL() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2Item) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2Item) GetProjectNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2Item) GetProjectURL() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2Item) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *ProjectV2ItemChange) GetArchivedAt() *ArchivedAt { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2ItemChange) GetFieldValue() *FieldValue { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2ItemContent) GetDraftIssue() *ProjectV2DraftIssue {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2ItemContent) GetIssue() *Issue { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2ItemContent) GetPullRequest() *PullRequest { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2ItemEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2ItemEvent) GetChanges() *ProjectV2ItemChange {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2ItemEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2ItemEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2ItemEvent) GetProjectV2Item() *ProjectV2Item {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProjectV2ItemEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2ItemFieldValue) GetDataType() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2ItemFieldValue) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *ProjectV2ItemFieldValue) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2ItemFieldValue) GetValue() any { _ = "STUB: not implemented"; return *new(any) }

func (p *ProjectV2StatusUpdate) GetBody() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2StatusUpdate) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *ProjectV2StatusUpdate) GetCreator() *User { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2StatusUpdate) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *ProjectV2StatusUpdate) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2StatusUpdate) GetProjectNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2StatusUpdate) GetStartDate() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2StatusUpdate) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2StatusUpdate) GetTargetDate() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2StatusUpdate) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *ProjectV2TextContent) GetHTML() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2TextContent) GetRaw() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2View) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *ProjectV2View) GetCreator() User { _ = "STUB: not implemented"; return *new(User) }

func (p *ProjectV2View) GetFilter() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2View) GetGroupBy() []int64 { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2View) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2View) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *ProjectV2View) GetLayout() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2View) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2View) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2View) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (p *ProjectV2View) GetProjectURL() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2View) GetSortBy() []*ProjectV2ViewSortBy { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2View) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *ProjectV2View) GetVerticalGroupBy() []int64 { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2View) GetVisibleFields() []int64 { _ = "STUB: not implemented"; return nil }

func (p *ProjectV2ViewSortBy) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (p *ProjectV2ViewSortBy) GetFieldID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *Protection) GetAllowDeletions() *AllowDeletions { _ = "STUB: not implemented"; return nil }

func (p *Protection) GetAllowForcePushes() *AllowForcePushes { _ = "STUB: not implemented"; return nil }

func (p *Protection) GetAllowForkSyncing() *AllowForkSyncing { _ = "STUB: not implemented"; return nil }

func (p *Protection) GetBlockCreations() *BlockCreations { _ = "STUB: not implemented"; return nil }

func (p *Protection) GetEnforceAdmins() *AdminEnforcement { _ = "STUB: not implemented"; return nil }

func (p *Protection) GetLockBranch() *LockBranch { _ = "STUB: not implemented"; return nil }

func (p *Protection) GetRequiredConversationResolution() *RequiredConversationResolution {
	_ = "STUB: not implemented"
	return nil
}

func (p *Protection) GetRequiredPullRequestReviews() *PullRequestReviewsEnforcement {
	_ = "STUB: not implemented"
	return nil
}

func (p *Protection) GetRequiredSignatures() *SignaturesProtectedBranch {
	_ = "STUB: not implemented"
	return nil
}

func (p *Protection) GetRequiredStatusChecks() *RequiredStatusChecks {
	_ = "STUB: not implemented"
	return nil
}

func (p *Protection) GetRequireLinearHistory() *RequireLinearHistory {
	_ = "STUB: not implemented"
	return nil
}

func (p *Protection) GetRestrictions() *BranchRestrictions { _ = "STUB: not implemented"; return nil }

func (p *Protection) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *ProtectionChanges) GetAdminEnforced() *AdminEnforcedChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetAllowDeletionsEnforcementLevel() *AllowDeletionsEnforcementLevelChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetAuthorizedActorNames() *AuthorizedActorNames {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetAuthorizedActorsOnly() *AuthorizedActorsOnly {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetAuthorizedDismissalActorsOnly() *AuthorizedDismissalActorsOnlyChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetCreateProtected() *CreateProtectedChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetDismissStaleReviewsOnPush() *DismissStaleReviewsOnPushChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetLinearHistoryRequirementEnforcementLevel() *LinearHistoryRequirementEnforcementLevelChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetPullRequestReviewsEnforcementLevel() *PullRequestReviewsEnforcementLevelChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetRequireCodeOwnerReview() *RequireCodeOwnerReviewChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetRequiredConversationResolutionLevel() *RequiredConversationResolutionLevelChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetRequiredDeploymentsEnforcementLevel() *RequiredDeploymentsEnforcementLevelChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetRequiredStatusChecks() *RequiredStatusChecksChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetRequiredStatusChecksEnforcementLevel() *RequiredStatusChecksEnforcementLevelChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetRequireLastPushApproval() *RequireLastPushApprovalChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionChanges) GetSignatureRequirementEnforcementLevel() *SignatureRequirementEnforcementLevelChanges {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionRequest) GetAllowDeletions() bool { _ = "STUB: not implemented"; return false }

func (p *ProtectionRequest) GetAllowForcePushes() bool { _ = "STUB: not implemented"; return false }

func (p *ProtectionRequest) GetAllowForkSyncing() bool { _ = "STUB: not implemented"; return false }

func (p *ProtectionRequest) GetBlockCreations() bool { _ = "STUB: not implemented"; return false }

func (p *ProtectionRequest) GetEnforceAdmins() bool { _ = "STUB: not implemented"; return false }

func (p *ProtectionRequest) GetLockBranch() bool { _ = "STUB: not implemented"; return false }

func (p *ProtectionRequest) GetRequiredConversationResolution() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ProtectionRequest) GetRequiredPullRequestReviews() *PullRequestReviewsEnforcementRequest {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionRequest) GetRequiredStatusChecks() *RequiredStatusChecks {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionRequest) GetRequireLinearHistory() bool { _ = "STUB: not implemented"; return false }

func (p *ProtectionRequest) GetRestrictions() *BranchRestrictionsRequest {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProtectionRule) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *ProtectionRule) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *ProtectionRule) GetPreventSelfReview() bool { _ = "STUB: not implemented"; return false }

func (p *ProtectionRule) GetReviewers() []*RequiredReviewer { _ = "STUB: not implemented"; return nil }

func (p *ProtectionRule) GetType() string { _ = "STUB: not implemented"; return "" }

func (p *ProtectionRule) GetWaitTimer() int { _ = "STUB: not implemented"; return 0 }

func (p *PublicEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (p *PublicEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (p *PublicEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (p *PublicEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *PublicIPUsage) GetCurrentUsage() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PublicIPUsage) GetMaximum() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PublicKey) GetKey() string { _ = "STUB: not implemented"; return "" }

func (p *PublicKey) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (p *PublishCodespaceOptions) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PublishCodespaceOptions) GetPrivate() bool { _ = "STUB: not implemented"; return false }

func (p *PullRequest) GetActiveLockReason() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetAdditions() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequest) GetAssignee() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequest) GetAssignees() []*User { _ = "STUB: not implemented"; return nil }

func (p *PullRequest) GetAuthorAssociation() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetAutoMerge() *PullRequestAutoMerge { _ = "STUB: not implemented"; return nil }

func (p *PullRequest) GetBase() *PullRequestBranch { _ = "STUB: not implemented"; return nil }

func (p *PullRequest) GetBody() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetChangedFiles() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequest) GetClosedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *PullRequest) GetComments() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequest) GetCommentsURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetCommits() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequest) GetCommitsURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *PullRequest) GetDeletions() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequest) GetDiffURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetDraft() bool { _ = "STUB: not implemented"; return false }

func (p *PullRequest) GetHead() *PullRequestBranch { _ = "STUB: not implemented"; return nil }

func (p *PullRequest) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PullRequest) GetIssueURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetLabels() []*Label { _ = "STUB: not implemented"; return nil }

func (p *PullRequest) GetLinks() *PRLinks { _ = "STUB: not implemented"; return nil }

func (p *PullRequest) GetLocked() bool { _ = "STUB: not implemented"; return false }

func (p *PullRequest) GetMaintainerCanModify() bool { _ = "STUB: not implemented"; return false }

func (p *PullRequest) GetMergeable() bool { _ = "STUB: not implemented"; return false }

func (p *PullRequest) GetMergeableState() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetMergeCommitSHA() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetMerged() bool { _ = "STUB: not implemented"; return false }

func (p *PullRequest) GetMergedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *PullRequest) GetMergedBy() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequest) GetMilestone() *Milestone { _ = "STUB: not implemented"; return nil }

func (p *PullRequest) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequest) GetPatchURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetRebaseable() bool { _ = "STUB: not implemented"; return false }

func (p *PullRequest) GetRequestedReviewers() []*User { _ = "STUB: not implemented"; return nil }

func (p *PullRequest) GetRequestedTeams() []*Team { _ = "STUB: not implemented"; return nil }

func (p *PullRequest) GetReviewComments() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequest) GetReviewCommentsURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetReviewCommentURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetState() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetStatusesURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (p *PullRequest) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequest) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestAutoMerge) GetCommitMessage() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestAutoMerge) GetCommitTitle() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestAutoMerge) GetEnabledBy() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestAutoMerge) GetMergeMethod() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestBranch) GetLabel() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestBranch) GetRef() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestBranch) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (p *PullRequestBranch) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestBranch) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestBranchRule) GetParameters() PullRequestRuleParameters {
	_ = "STUB: not implemented"
	return *new(PullRequestRuleParameters)
}

func (p *PullRequestBranchUpdateOptions) GetExpectedHeadSHA() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *PullRequestBranchUpdateResponse) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestBranchUpdateResponse) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetAuthorAssociation() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetBody() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetCommitID() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PullRequestComment) GetDiffHunk() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestComment) GetInReplyTo() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestComment) GetLine() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestComment) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetOriginalCommitID() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetOriginalLine() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestComment) GetOriginalPosition() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestComment) GetOriginalStartLine() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestComment) GetPath() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetPosition() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestComment) GetPullRequestReviewID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestComment) GetPullRequestURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetReactions() *Reactions { _ = "STUB: not implemented"; return nil }

func (p *PullRequestComment) GetSide() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetStartLine() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestComment) GetStartSide() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetSubjectType() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PullRequestComment) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestComment) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestEvent) GetAfter() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestEvent) GetAssignee() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestEvent) GetBefore() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestEvent) GetChanges() *EditChange { _ = "STUB: not implemented"; return nil }

func (p *PullRequestEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (p *PullRequestEvent) GetLabel() *Label { _ = "STUB: not implemented"; return nil }

func (p *PullRequestEvent) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestEvent) GetOrganization() *Organization { _ = "STUB: not implemented"; return nil }

func (p *PullRequestEvent) GetPerformedViaGithubApp() *App { _ = "STUB: not implemented"; return nil }

func (p *PullRequestEvent) GetPullRequest() *PullRequest { _ = "STUB: not implemented"; return nil }

func (p *PullRequestEvent) GetReason() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (p *PullRequestEvent) GetRequestedReviewer() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestEvent) GetRequestedTeam() *Team { _ = "STUB: not implemented"; return nil }

func (p *PullRequestEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestLinks) GetDiffURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestLinks) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestLinks) GetMergedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PullRequestLinks) GetPatchURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestLinks) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestListCommentsOptions) GetDirection() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *PullRequestListCommentsOptions) GetSince() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (p *PullRequestListCommentsOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestListOptions) GetBase() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestListOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestListOptions) GetHead() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestListOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestListOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestMergeResult) GetMerged() bool { _ = "STUB: not implemented"; return false }

func (p *PullRequestMergeResult) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestMergeResult) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestOptions) GetCommitTitle() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestOptions) GetDontDefaultIfBlank() bool { _ = "STUB: not implemented"; return false }

func (p *PullRequestOptions) GetMergeMethod() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestOptions) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReview) GetAuthorAssociation() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReview) GetBody() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReview) GetCommitID() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReview) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReview) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestReview) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReview) GetPullRequestURL() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReview) GetState() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReview) GetSubmittedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PullRequestReview) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestReviewCommentEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReviewCommentEvent) GetChanges() *EditChange {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewCommentEvent) GetComment() *PullRequestComment {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewCommentEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewCommentEvent) GetOrg() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewCommentEvent) GetPullRequest() *PullRequest {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewCommentEvent) GetRepo() *Repository {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewCommentEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestReviewDismissalRequest) GetMessage() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *PullRequestReviewEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReviewEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewEvent) GetOrganization() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewEvent) GetPullRequest() *PullRequest {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (p *PullRequestReviewEvent) GetReview() *PullRequestReview {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestReviewRequest) GetBody() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReviewRequest) GetComments() []*DraftReviewComment {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewRequest) GetCommitID() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReviewRequest) GetEvent() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReviewRequest) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReviewsEnforcement) GetBypassPullRequestAllowances() *BypassPullRequestAllowances {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewsEnforcement) GetDismissalRestrictions() *DismissalRestrictions {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewsEnforcement) GetDismissStaleReviews() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestReviewsEnforcement) GetRequireCodeOwnerReviews() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestReviewsEnforcement) GetRequiredApprovingReviewCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (p *PullRequestReviewsEnforcement) GetRequireLastPushApproval() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestReviewsEnforcementLevelChanges) GetFrom() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *PullRequestReviewsEnforcementRequest) GetBypassPullRequestAllowancesRequest() *BypassPullRequestAllowancesRequest {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewsEnforcementRequest) GetDismissalRestrictionsRequest() *DismissalRestrictionsRequest {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewsEnforcementRequest) GetDismissStaleReviews() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestReviewsEnforcementRequest) GetRequireCodeOwnerReviews() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestReviewsEnforcementRequest) GetRequiredApprovingReviewCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (p *PullRequestReviewsEnforcementRequest) GetRequireLastPushApproval() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestReviewsEnforcementUpdate) GetBypassPullRequestAllowancesRequest() *BypassPullRequestAllowancesRequest {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewsEnforcementUpdate) GetDismissalRestrictionsRequest() *DismissalRestrictionsRequest {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewsEnforcementUpdate) GetDismissStaleReviews() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestReviewsEnforcementUpdate) GetRequireCodeOwnerReviews() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestReviewsEnforcementUpdate) GetRequiredApprovingReviewCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (p *PullRequestReviewsEnforcementUpdate) GetRequireLastPushApproval() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestReviewThreadEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestReviewThreadEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewThreadEvent) GetOrg() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewThreadEvent) GetPullRequest() *PullRequest {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestReviewThreadEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (p *PullRequestReviewThreadEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestReviewThreadEvent) GetThread() *PullRequestThread {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestRuleParameters) GetAllowedMergeMethods() []PullRequestMergeMethod {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestRuleParameters) GetDismissStaleReviewsOnPush() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestRuleParameters) GetRequireCodeOwnerReview() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestRuleParameters) GetRequiredApprovingReviewCount() int {
	_ = "STUB: not implemented"
	return 0
}

func (p *PullRequestRuleParameters) GetRequiredReviewers() []*RulesetRequiredReviewer {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestRuleParameters) GetRequiredReviewThreadResolution() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestRuleParameters) GetRequireLastPushApproval() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullRequestTargetEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestTargetEvent) GetAfter() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestTargetEvent) GetAssignee() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestTargetEvent) GetBefore() string { _ = "STUB: not implemented"; return "" }

func (p *PullRequestTargetEvent) GetChanges() *EditChange { _ = "STUB: not implemented"; return nil }

func (p *PullRequestTargetEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestTargetEvent) GetLabel() *Label { _ = "STUB: not implemented"; return nil }

func (p *PullRequestTargetEvent) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestTargetEvent) GetOrganization() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestTargetEvent) GetPerformedViaGithubApp() *App {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestTargetEvent) GetPullRequest() *PullRequest {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestTargetEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (p *PullRequestTargetEvent) GetRequestedReviewer() *User {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestTargetEvent) GetRequestedTeam() *Team { _ = "STUB: not implemented"; return nil }

func (p *PullRequestTargetEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *PullRequestThread) GetComments() []*PullRequestComment {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullRequestThread) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PullRequestThread) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *PullStats) GetMergeablePulls() int { _ = "STUB: not implemented"; return 0 }

func (p *PullStats) GetMergedPulls() int { _ = "STUB: not implemented"; return 0 }

func (p *PullStats) GetTotalPulls() int { _ = "STUB: not implemented"; return 0 }

func (p *PullStats) GetUnmergeablePulls() int { _ = "STUB: not implemented"; return 0 }

func (p *PunchCard) GetCommits() int { _ = "STUB: not implemented"; return 0 }

func (p *PunchCard) GetDay() int { _ = "STUB: not implemented"; return 0 }

func (p *PunchCard) GetHour() int { _ = "STUB: not implemented"; return 0 }

func (p *PushEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (p *PushEvent) GetAfter() string { _ = "STUB: not implemented"; return "" }

func (p *PushEvent) GetBaseRef() string { _ = "STUB: not implemented"; return "" }

func (p *PushEvent) GetBefore() string { _ = "STUB: not implemented"; return "" }

func (p *PushEvent) GetCommits() []*HeadCommit { _ = "STUB: not implemented"; return nil }

func (p *PushEvent) GetCompare() string { _ = "STUB: not implemented"; return "" }

func (p *PushEvent) GetCreated() bool { _ = "STUB: not implemented"; return false }

func (p *PushEvent) GetDeleted() bool { _ = "STUB: not implemented"; return false }

func (p *PushEvent) GetDistinctSize() int { _ = "STUB: not implemented"; return 0 }

func (p *PushEvent) GetForced() bool { _ = "STUB: not implemented"; return false }

func (p *PushEvent) GetHead() string { _ = "STUB: not implemented"; return "" }

func (p *PushEvent) GetHeadCommit() *HeadCommit { _ = "STUB: not implemented"; return nil }

func (p *PushEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (p *PushEvent) GetOrganization() *Organization { _ = "STUB: not implemented"; return nil }

func (p *PushEvent) GetPusher() *CommitAuthor { _ = "STUB: not implemented"; return nil }

func (p *PushEvent) GetPushID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PushEvent) GetRef() string { _ = "STUB: not implemented"; return "" }

func (p *PushEvent) GetRepo() *PushEventRepository { _ = "STUB: not implemented"; return nil }

func (p *PushEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (p *PushEvent) GetSize() int { _ = "STUB: not implemented"; return 0 }

func (p *PushEventRepoOwner) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepoOwner) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetArchived() bool { _ = "STUB: not implemented"; return false }

func (p *PushEventRepository) GetArchiveURL() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetCloneURL() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PushEventRepository) GetCustomProperties() map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (p *PushEventRepository) GetDefaultBranch() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetDisabled() bool { _ = "STUB: not implemented"; return false }

func (p *PushEventRepository) GetFork() bool { _ = "STUB: not implemented"; return false }

func (p *PushEventRepository) GetForksCount() int { _ = "STUB: not implemented"; return 0 }

func (p *PushEventRepository) GetFullName() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetGitURL() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetHasDownloads() bool { _ = "STUB: not implemented"; return false }

func (p *PushEventRepository) GetHasIssues() bool { _ = "STUB: not implemented"; return false }

func (p *PushEventRepository) GetHasPages() bool { _ = "STUB: not implemented"; return false }

func (p *PushEventRepository) GetHasWiki() bool { _ = "STUB: not implemented"; return false }

func (p *PushEventRepository) GetHomepage() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (p *PushEventRepository) GetLanguage() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetMasterBranch() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetName() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetOpenIssuesCount() int { _ = "STUB: not implemented"; return 0 }

func (p *PushEventRepository) GetOrganization() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetOwner() *User { _ = "STUB: not implemented"; return nil }

func (p *PushEventRepository) GetPrivate() bool { _ = "STUB: not implemented"; return false }

func (p *PushEventRepository) GetPullsURL() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetPushedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PushEventRepository) GetSize() int { _ = "STUB: not implemented"; return 0 }

func (p *PushEventRepository) GetSSHURL() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetStargazersCount() int { _ = "STUB: not implemented"; return 0 }

func (p *PushEventRepository) GetStatusesURL() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetSVNURL() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetTopics() []string { _ = "STUB: not implemented"; return nil }

func (p *PushEventRepository) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PushEventRepository) GetURL() string { _ = "STUB: not implemented"; return "" }

func (p *PushEventRepository) GetWatchersCount() int { _ = "STUB: not implemented"; return 0 }

func (p *PushProtectionBypass) GetExpireAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (p *PushProtectionBypass) GetReason() string { _ = "STUB: not implemented"; return "" }

func (p *PushProtectionBypass) GetTokenType() string { _ = "STUB: not implemented"; return "" }

func (p *PushProtectionBypassRequest) GetPlaceholderID() string {
	_ = "STUB: not implemented"
	return ""
}

func (p *PushProtectionBypassRequest) GetReason() string { _ = "STUB: not implemented"; return "" }

func (r *Rate) GetLimit() int { _ = "STUB: not implemented"; return 0 }

func (r *Rate) GetRemaining() int { _ = "STUB: not implemented"; return 0 }

func (r *Rate) GetReset() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (r *Rate) GetResource() string { _ = "STUB: not implemented"; return "" }

func (r *Rate) GetUsed() int { _ = "STUB: not implemented"; return 0 }

func (r *RateLimitError) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (r *RateLimitError) GetRate() Rate { _ = "STUB: not implemented"; return *new(Rate) }

func (r *RateLimits) GetActionsRunnerRegistration() *Rate { _ = "STUB: not implemented"; return nil }

func (r *RateLimits) GetAuditLog() *Rate { _ = "STUB: not implemented"; return nil }

func (r *RateLimits) GetCodeScanningUpload() *Rate { _ = "STUB: not implemented"; return nil }

func (r *RateLimits) GetCodeSearch() *Rate { _ = "STUB: not implemented"; return nil }

func (r *RateLimits) GetCore() *Rate { _ = "STUB: not implemented"; return nil }

func (r *RateLimits) GetDependencySBOM() *Rate { _ = "STUB: not implemented"; return nil }

func (r *RateLimits) GetDependencySnapshots() *Rate { _ = "STUB: not implemented"; return nil }

func (r *RateLimits) GetGraphQL() *Rate { _ = "STUB: not implemented"; return nil }

func (r *RateLimits) GetIntegrationManifest() *Rate { _ = "STUB: not implemented"; return nil }

func (r *RateLimits) GetSCIM() *Rate { _ = "STUB: not implemented"; return nil }

func (r *RateLimits) GetSearch() *Rate { _ = "STUB: not implemented"; return nil }

func (r *RateLimits) GetSourceImport() *Rate { _ = "STUB: not implemented"; return nil }

func (r *RawOptions) GetType() RawType { _ = "STUB: not implemented"; return *new(RawType) }

func (r *Reaction) GetContent() string { _ = "STUB: not implemented"; return "" }

func (r *Reaction) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (r *Reaction) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Reaction) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (r *Reaction) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (r *Reactions) GetConfused() int { _ = "STUB: not implemented"; return 0 }

func (r *Reactions) GetEyes() int { _ = "STUB: not implemented"; return 0 }

func (r *Reactions) GetHeart() int { _ = "STUB: not implemented"; return 0 }

func (r *Reactions) GetHooray() int { _ = "STUB: not implemented"; return 0 }

func (r *Reactions) GetLaugh() int { _ = "STUB: not implemented"; return 0 }

func (r *Reactions) GetMinusOne() int { _ = "STUB: not implemented"; return 0 }

func (r *Reactions) GetPlusOne() int { _ = "STUB: not implemented"; return 0 }

func (r *Reactions) GetRocket() int { _ = "STUB: not implemented"; return 0 }

func (r *Reactions) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (r *Reactions) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *ReassignedResource) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *ReassignedResource) GetPreviousCostCenter() string { _ = "STUB: not implemented"; return "" }

func (r *ReassignedResource) GetResourceType() string { _ = "STUB: not implemented"; return "" }

func (r *RedirectionError) GetStatusCode() int { _ = "STUB: not implemented"; return 0 }

func (r *Reference) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (r *Reference) GetObject() *GitObject { _ = "STUB: not implemented"; return nil }

func (r *Reference) GetRef() string { _ = "STUB: not implemented"; return "" }

func (r *Reference) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *ReferencedWorkflow) GetPath() string { _ = "STUB: not implemented"; return "" }

func (r *ReferencedWorkflow) GetRef() string { _ = "STUB: not implemented"; return "" }

func (r *ReferencedWorkflow) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (r *RegistrationToken) GetExpiresAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *RegistrationToken) GetToken() string { _ = "STUB: not implemented"; return "" }

func (r *RegistryPackageEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (r *RegistryPackageEvent) GetEnterprise() *Enterprise { _ = "STUB: not implemented"; return nil }

func (r *RegistryPackageEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (r *RegistryPackageEvent) GetOrganization() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (r *RegistryPackageEvent) GetRegistryPackage() *Package { _ = "STUB: not implemented"; return nil }

func (r *RegistryPackageEvent) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (r *RegistryPackageEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (r *ReleaseAsset) GetBrowserDownloadURL() string { _ = "STUB: not implemented"; return "" }

func (r *ReleaseAsset) GetContentType() string { _ = "STUB: not implemented"; return "" }

func (r *ReleaseAsset) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (r *ReleaseAsset) GetDigest() string { _ = "STUB: not implemented"; return "" }

func (r *ReleaseAsset) GetDownloadCount() int { _ = "STUB: not implemented"; return 0 }

func (r *ReleaseAsset) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *ReleaseAsset) GetLabel() string { _ = "STUB: not implemented"; return "" }

func (r *ReleaseAsset) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *ReleaseAsset) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (r *ReleaseAsset) GetSize() int { _ = "STUB: not implemented"; return 0 }

func (r *ReleaseAsset) GetState() string { _ = "STUB: not implemented"; return "" }

func (r *ReleaseAsset) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (r *ReleaseAsset) GetUploader() *User { _ = "STUB: not implemented"; return nil }

func (r *ReleaseAsset) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *ReleaseEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (r *ReleaseEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (r *ReleaseEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (r *ReleaseEvent) GetRelease() *RepositoryRelease { _ = "STUB: not implemented"; return nil }

func (r *ReleaseEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (r *ReleaseEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (r *ReleaseVersion) GetBuildDate() string { _ = "STUB: not implemented"; return "" }

func (r *ReleaseVersion) GetBuildID() string { _ = "STUB: not implemented"; return "" }

func (r *ReleaseVersion) GetPlatform() string { _ = "STUB: not implemented"; return "" }

func (r *ReleaseVersion) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (r *RemoveResourcesFromCostCenterResponse) GetMessage() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RemoveToken) GetExpiresAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (r *RemoveToken) GetToken() string { _ = "STUB: not implemented"; return "" }

func (r *Rename) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (r *Rename) GetTo() string { _ = "STUB: not implemented"; return "" }

func (r *RenameOrgResponse) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (r *RenameOrgResponse) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepoAdvisoryCredit) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (r *RepoAdvisoryCredit) GetType() string { _ = "STUB: not implemented"; return "" }

func (r *RepoAdvisoryCreditDetailed) GetState() string { _ = "STUB: not implemented"; return "" }

func (r *RepoAdvisoryCreditDetailed) GetType() string { _ = "STUB: not implemented"; return "" }

func (r *RepoAdvisoryCreditDetailed) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (r *RepoCustomPropertyValue) GetProperties() []*CustomPropertyValue {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepoCustomPropertyValue) GetRepositoryFullName() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepoCustomPropertyValue) GetRepositoryID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RepoCustomPropertyValue) GetRepositoryName() string { _ = "STUB: not implemented"; return "" }

func (r *RepoDependencies) GetDownloadLocation() string { _ = "STUB: not implemented"; return "" }

func (r *RepoDependencies) GetExternalRefs() []*PackageExternalRef {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepoDependencies) GetFilesAnalyzed() bool { _ = "STUB: not implemented"; return false }

func (r *RepoDependencies) GetLicenseConcluded() string { _ = "STUB: not implemented"; return "" }

func (r *RepoDependencies) GetLicenseDeclared() string { _ = "STUB: not implemented"; return "" }

func (r *RepoDependencies) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *RepoDependencies) GetSPDXID() string { _ = "STUB: not implemented"; return "" }

func (r *RepoDependencies) GetVersionInfo() string { _ = "STUB: not implemented"; return "" }

func (r *RepoFineGrainedPermission) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (r *RepoFineGrainedPermission) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *RepoImmutableReleasesStatus) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (r *RepoImmutableReleasesStatus) GetEnforcedByOwner() bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RepoMergeUpstreamRequest) GetBranch() string { _ = "STUB: not implemented"; return "" }

func (r *RepoMergeUpstreamResult) GetBaseBranch() string { _ = "STUB: not implemented"; return "" }

func (r *RepoMergeUpstreamResult) GetMergeType() string { _ = "STUB: not implemented"; return "" }

func (r *RepoMergeUpstreamResult) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (r *RepoName) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoriesSearchResult) GetIncompleteResults() bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RepositoriesSearchResult) GetRepositories() []*Repository {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoriesSearchResult) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (r *Repository) GetAllowAutoMerge() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetAllowForking() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetAllowMergeCommit() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetAllowRebaseMerge() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetAllowSquashMerge() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetAllowUpdateBranch() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetArchived() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetArchiveURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetAssigneesURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetAutoInit() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetBlobsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetBranchesURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetCloneURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetCodeOfConduct() *CodeOfConduct { _ = "STUB: not implemented"; return nil }

func (r *Repository) GetCollaboratorsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetCommentsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetCommitsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetCompareURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetContentsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetContributorsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (r *Repository) GetCustomProperties() map[string]any { _ = "STUB: not implemented"; return nil }

func (r *Repository) GetDefaultBranch() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetDeleteBranchOnMerge() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetDeploymentsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetDisabled() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetDownloadsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetEventsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetFork() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetForksCount() int { _ = "STUB: not implemented"; return 0 }

func (r *Repository) GetForksURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetFullName() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetGitCommitsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetGitignoreTemplate() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetGitRefsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetGitTagsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetGitURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetHasDiscussions() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetHasDownloads() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetHasIssues() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetHasPages() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetHasProjects() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetHasPullRequests() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetHasWiki() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetHomepage() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetHooksURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Repository) GetIssueCommentURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetIssueEventsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetIssuesURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetIsTemplate() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetKeysURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetLabelsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetLanguage() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetLanguagesURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetLicense() *License { _ = "STUB: not implemented"; return nil }

func (r *Repository) GetLicenseTemplate() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetMasterBranch() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetMergeCommitMessage() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetMergeCommitTitle() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetMergesURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetMilestonesURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetMirrorURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetNetworkCount() int { _ = "STUB: not implemented"; return 0 }

func (r *Repository) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetNotificationsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetOpenIssues() int { _ = "STUB: not implemented"; return 0 }

func (r *Repository) GetOpenIssuesCount() int { _ = "STUB: not implemented"; return 0 }

func (r *Repository) GetOrganization() *Organization { _ = "STUB: not implemented"; return nil }

func (r *Repository) GetOwner() *User { _ = "STUB: not implemented"; return nil }

func (r *Repository) GetParent() *Repository { _ = "STUB: not implemented"; return nil }

func (r *Repository) GetPermissions() *RepositoryPermissions { _ = "STUB: not implemented"; return nil }

func (r *Repository) GetPrivate() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetPullRequestCreationPolicy() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetPullsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetPushedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (r *Repository) GetReleasesURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetRoleName() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetSecurityAndAnalysis() *SecurityAndAnalysis {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) GetSize() int { _ = "STUB: not implemented"; return 0 }

func (r *Repository) GetSource() *Repository { _ = "STUB: not implemented"; return nil }

func (r *Repository) GetSquashMergeCommitMessage() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetSquashMergeCommitTitle() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetSSHURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetStargazersCount() int { _ = "STUB: not implemented"; return 0 }

func (r *Repository) GetStargazersURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetStatusesURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetSubscribersCount() int { _ = "STUB: not implemented"; return 0 }

func (r *Repository) GetSubscribersURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetSubscriptionURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetSVNURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetTagsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetTeamID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Repository) GetTeamsURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetTemplateRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (r *Repository) GetTextMatches() []*TextMatch { _ = "STUB: not implemented"; return nil }

func (r *Repository) GetTopics() []string { _ = "STUB: not implemented"; return nil }

func (r *Repository) GetTreesURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (r *Repository) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetUseSquashPRTitleAsDefault() bool { _ = "STUB: not implemented"; return false }

func (r *Repository) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (r *Repository) GetWatchers() int { _ = "STUB: not implemented"; return 0 }

func (r *Repository) GetWatchersCount() int { _ = "STUB: not implemented"; return 0 }

func (r *Repository) GetWebCommitSignoffRequired() bool { _ = "STUB: not implemented"; return false }

func (r *RepositoryActionsAccessLevel) GetAccessLevel() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryActiveCommitters) GetAdvancedSecurityCommitters() int {
	_ = "STUB: not implemented"
	return 0
}

func (r *RepositoryActiveCommitters) GetAdvancedSecurityCommittersBreakdown() []*AdvancedSecurityCommittersBreakdown {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryActiveCommitters) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActivity) GetActivityType() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActivity) GetActor() *RepositoryActor { _ = "STUB: not implemented"; return nil }

func (r *RepositoryActivity) GetAfter() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActivity) GetBefore() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActivity) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RepositoryActivity) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActivity) GetRef() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActivity) GetTimestamp() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *RepositoryActor) GetAvatarURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetEventsURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetFollowersURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetFollowingURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetGistsURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetGravatarID() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RepositoryActor) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetOrganizationsURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetReceivedEventsURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetReposURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetSiteAdmin() bool { _ = "STUB: not implemented"; return false }

func (r *RepositoryActor) GetStarredURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetSubscriptionsURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetType() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryActor) GetUserViewType() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryAddCollaboratorOptions) GetPermission() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryAttachment) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (r *RepositoryAttachment) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryCodeSecurityConfiguration) GetConfiguration() *CodeSecurityConfiguration {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryCodeSecurityConfiguration) GetState() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryComment) GetBody() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryComment) GetCommitID() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryComment) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *RepositoryComment) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryComment) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RepositoryComment) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryComment) GetPath() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryComment) GetPosition() int { _ = "STUB: not implemented"; return 0 }

func (r *RepositoryComment) GetReactions() *Reactions { _ = "STUB: not implemented"; return nil }

func (r *RepositoryComment) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *RepositoryComment) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryComment) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (r *RepositoryCommit) GetAuthor() *User { _ = "STUB: not implemented"; return nil }

func (r *RepositoryCommit) GetCommentsURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryCommit) GetCommit() *Commit { _ = "STUB: not implemented"; return nil }

func (r *RepositoryCommit) GetCommitter() *User { _ = "STUB: not implemented"; return nil }

func (r *RepositoryCommit) GetFiles() []*CommitFile { _ = "STUB: not implemented"; return nil }

func (r *RepositoryCommit) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryCommit) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryCommit) GetParents() []*Commit { _ = "STUB: not implemented"; return nil }

func (r *RepositoryCommit) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryCommit) GetStats() *CommitStats { _ = "STUB: not implemented"; return nil }

func (r *RepositoryCommit) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContent) GetDownloadURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContent) GetEncoding() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContent) GetGitURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContent) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContent) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContent) GetPath() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContent) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContent) GetSize() int { _ = "STUB: not implemented"; return 0 }

func (r *RepositoryContent) GetSubmoduleGitURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContent) GetTarget() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContent) GetType() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContent) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContentFileOptions) GetAuthor() *CommitAuthor {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryContentFileOptions) GetBranch() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContentFileOptions) GetCommitter() *CommitAuthor {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryContentFileOptions) GetContent() []byte { _ = "STUB: not implemented"; return nil }

func (r *RepositoryContentFileOptions) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContentFileOptions) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContentGetOptions) GetRef() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryContentResponse) GetContent() *RepositoryContent {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryCreateForkOptions) GetDefaultBranchOnly() bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RepositoryCreateForkOptions) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryCreateForkOptions) GetOrganization() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryDispatchEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryDispatchEvent) GetBranch() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryDispatchEvent) GetClientPayload() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (r *RepositoryDispatchEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryDispatchEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (r *RepositoryDispatchEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (r *RepositoryDispatchEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (r *RepositoryEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryEvent) GetChanges() *EditChange { _ = "STUB: not implemented"; return nil }

func (r *RepositoryEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (r *RepositoryEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (r *RepositoryEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (r *RepositoryEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (r *RepositoryImportEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (r *RepositoryImportEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (r *RepositoryImportEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (r *RepositoryImportEvent) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryInvitation) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *RepositoryInvitation) GetExpired() bool { _ = "STUB: not implemented"; return false }

func (r *RepositoryInvitation) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryInvitation) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RepositoryInvitation) GetInvitee() *User { _ = "STUB: not implemented"; return nil }

func (r *RepositoryInvitation) GetInviter() *User { _ = "STUB: not implemented"; return nil }

func (r *RepositoryInvitation) GetPermissions() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryInvitation) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (r *RepositoryInvitation) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryLicense) GetContent() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryLicense) GetDownloadURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryLicense) GetEncoding() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryLicense) GetGitURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryLicense) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryLicense) GetLicense() *License { _ = "STUB: not implemented"; return nil }

func (r *RepositoryLicense) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryLicense) GetPath() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryLicense) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryLicense) GetSize() int { _ = "STUB: not implemented"; return 0 }

func (r *RepositoryLicense) GetType() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryLicense) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListAllOptions) GetSince() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RepositoryListByAuthenticatedUserOptions) GetAffiliation() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryListByAuthenticatedUserOptions) GetDirection() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryListByAuthenticatedUserOptions) GetSort() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryListByAuthenticatedUserOptions) GetType() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryListByAuthenticatedUserOptions) GetVisibility() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryListByOrgOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListByOrgOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListByOrgOptions) GetType() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListByUserOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListByUserOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListByUserOptions) GetType() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListForksOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListOptions) GetAffiliation() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListOptions) GetDirection() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListOptions) GetType() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListOptions) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryListRulesetsOptions) GetIncludesParents() bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RepositoryMergeRequest) GetBase() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryMergeRequest) GetCommitMessage() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryMergeRequest) GetHead() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryParticipation) GetAll() []int { _ = "STUB: not implemented"; return nil }

func (r *RepositoryParticipation) GetOwner() []int { _ = "STUB: not implemented"; return nil }

func (r *RepositoryPermissionLevel) GetPermission() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryPermissionLevel) GetRoleName() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryPermissionLevel) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (r *RepositoryPermissions) GetAdmin() bool { _ = "STUB: not implemented"; return false }

func (r *RepositoryPermissions) GetMaintain() bool { _ = "STUB: not implemented"; return false }

func (r *RepositoryPermissions) GetPull() bool { _ = "STUB: not implemented"; return false }

func (r *RepositoryPermissions) GetPush() bool { _ = "STUB: not implemented"; return false }

func (r *RepositoryPermissions) GetTriage() bool { _ = "STUB: not implemented"; return false }

func (r *RepositoryRelease) GetAssets() []*ReleaseAsset { _ = "STUB: not implemented"; return nil }

func (r *RepositoryRelease) GetAssetsURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetAuthor() *User { _ = "STUB: not implemented"; return nil }

func (r *RepositoryRelease) GetBody() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetBodyHTML() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetBodyText() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *RepositoryRelease) GetDiscussionURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetDraft() bool { _ = "STUB: not implemented"; return false }

func (r *RepositoryRelease) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RepositoryRelease) GetImmutable() bool { _ = "STUB: not implemented"; return false }

func (r *RepositoryRelease) GetMentionsCount() int { _ = "STUB: not implemented"; return 0 }

func (r *RepositoryRelease) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetPrerelease() bool { _ = "STUB: not implemented"; return false }

func (r *RepositoryRelease) GetPublishedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *RepositoryRelease) GetReactions() *Reactions { _ = "STUB: not implemented"; return nil }

func (r *RepositoryRelease) GetTagName() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetTarballURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetTargetCommitish() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *RepositoryRelease) GetUploadURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRelease) GetZipballURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryReleaseNotes) GetBody() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryReleaseNotes) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRule) GetParameters() any { _ = "STUB: not implemented"; return *new(any) }

func (r *RepositoryRule) GetType() RepositoryRuleType {
	_ = "STUB: not implemented"
	return *new(RepositoryRuleType)
}

func (r *RepositoryRuleset) GetBypassActors() []*BypassActor { _ = "STUB: not implemented"; return nil }

func (r *RepositoryRuleset) GetConditions() *RepositoryRulesetConditions {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRuleset) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *RepositoryRuleset) GetCurrentUserCanBypass() *BypassMode {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRuleset) GetEnforcement() RulesetEnforcement {
	_ = "STUB: not implemented"
	return *new(RulesetEnforcement)
}

func (r *RepositoryRuleset) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RepositoryRuleset) GetLinks() *RepositoryRulesetLinks {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRuleset) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRuleset) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRuleset) GetRules() *RepositoryRulesetRules {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRuleset) GetSource() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRuleset) GetSourceType() *RulesetSourceType {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRuleset) GetTarget() *RulesetTarget { _ = "STUB: not implemented"; return nil }

func (r *RepositoryRuleset) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *RepositoryRulesetChangedConditions) GetAdded() []*RepositoryRulesetConditions {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChangedConditions) GetDeleted() []*RepositoryRulesetConditions {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChangedConditions) GetUpdated() []*RepositoryRulesetUpdatedConditions {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChangedRule) GetConfiguration() *RepositoryRulesetChangeSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChangedRule) GetPattern() *RepositoryRulesetChangeSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChangedRule) GetRuleType() *RepositoryRulesetChangeSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChangedRules) GetAdded() []*RepositoryRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChangedRules) GetDeleted() []*RepositoryRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChangedRules) GetUpdated() []*RepositoryRulesetUpdatedRules {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChanges) GetConditions() *RepositoryRulesetChangedConditions {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChanges) GetEnforcement() *RepositoryRulesetChangeSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChanges) GetName() *RepositoryRulesetChangeSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChanges) GetRules() *RepositoryRulesetChangedRules {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetChangeSource) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRulesetChangeSources) GetFrom() []string { _ = "STUB: not implemented"; return nil }

func (r *RepositoryRulesetConditions) GetOrganizationID() *RepositoryRulesetOrganizationIDsConditionParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetConditions) GetOrganizationName() *RepositoryRulesetOrganizationNamesConditionParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetConditions) GetOrganizationProperty() *RepositoryRulesetOrganizationPropertyConditionParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetConditions) GetRefName() *RepositoryRulesetRefConditionParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetConditions) GetRepositoryID() *RepositoryRulesetRepositoryIDsConditionParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetConditions) GetRepositoryName() *RepositoryRulesetRepositoryNamesConditionParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetConditions) GetRepositoryProperty() *RepositoryRulesetRepositoryPropertyConditionParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRulesetEvent) GetChanges() *RepositoryRulesetChanges {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetEvent) GetEnterprise() *Enterprise { _ = "STUB: not implemented"; return nil }

func (r *RepositoryRulesetEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetEvent) GetOrganization() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetEvent) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (r *RepositoryRulesetEvent) GetRepositoryRuleset() *RepositoryRuleset {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (r *RepositoryRulesetLink) GetHRef() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryRulesetLinks) GetHTML() *RepositoryRulesetLink {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetLinks) GetSelf() *RepositoryRulesetLink {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetOrganizationIDsConditionParameters) GetOrganizationIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetOrganizationNamesConditionParameters) GetExclude() []string {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetOrganizationNamesConditionParameters) GetInclude() []string {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetOrganizationPropertyConditionParameters) GetExclude() []*RepositoryRulesetRepositoryPropertyTargetParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetOrganizationPropertyConditionParameters) GetInclude() []*RepositoryRulesetRepositoryPropertyTargetParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRefConditionParameters) GetExclude() []string {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRefConditionParameters) GetInclude() []string {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRepositoryIDsConditionParameters) GetRepositoryIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRepositoryNamesConditionParameters) GetExclude() []string {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRepositoryNamesConditionParameters) GetInclude() []string {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRepositoryNamesConditionParameters) GetProtected() bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RepositoryRulesetRepositoryPropertyConditionParameters) GetExclude() []*RepositoryRulesetRepositoryPropertyTargetParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRepositoryPropertyConditionParameters) GetInclude() []*RepositoryRulesetRepositoryPropertyTargetParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRepositoryPropertyTargetParameters) GetName() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryRulesetRepositoryPropertyTargetParameters) GetPropertyValues() []string {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRepositoryPropertyTargetParameters) GetSource() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryRulesetRules) GetBranchNamePattern() *PatternRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetCodeScanning() *CodeScanningRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetCommitAuthorEmailPattern() *PatternRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetCommitMessagePattern() *PatternRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetCommitterEmailPattern() *PatternRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetCopilotCodeReview() *CopilotCodeReviewRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetCreation() *EmptyRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetDeletion() *EmptyRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetFileExtensionRestriction() *FileExtensionRestrictionRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetFilePathRestriction() *FilePathRestrictionRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetMaxFilePathLength() *MaxFilePathLengthRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetMaxFileSize() *MaxFileSizeRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetMergeQueue() *MergeQueueRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetNonFastForward() *EmptyRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetPullRequest() *PullRequestRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetRepositoryCreate() *EmptyRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetRepositoryDelete() *EmptyRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetRepositoryName() *SimplePatternRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetRepositoryTransfer() *EmptyRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetRepositoryVisibility() *RepositoryVisibilityRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetRequiredDeployments() *RequiredDeploymentsRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetRequiredLinearHistory() *EmptyRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetRequiredSignatures() *EmptyRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetRequiredStatusChecks() *RequiredStatusChecksRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetTagNamePattern() *PatternRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetUpdate() *UpdateRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetRules) GetWorkflows() *WorkflowsRuleParameters {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetUpdatedCondition) GetConditionType() *RepositoryRulesetChangeSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetUpdatedCondition) GetExclude() *RepositoryRulesetChangeSources {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetUpdatedCondition) GetInclude() *RepositoryRulesetChangeSources {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetUpdatedCondition) GetTarget() *RepositoryRulesetChangeSource {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetUpdatedConditions) GetChanges() *RepositoryRulesetUpdatedCondition {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetUpdatedConditions) GetCondition() *RepositoryRulesetConditions {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetUpdatedRules) GetChanges() *RepositoryRulesetChangedRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryRulesetUpdatedRules) GetRule() *RepositoryRule {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryTag) GetCommit() *Commit { _ = "STUB: not implemented"; return nil }

func (r *RepositoryTag) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryTag) GetTarballURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryTag) GetZipballURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryVisibilityRuleParameters) GetInternal() bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RepositoryVisibilityRuleParameters) GetPrivate() bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RepositoryVulnerabilityAlert) GetAffectedPackageName() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryVulnerabilityAlert) GetAffectedRange() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryVulnerabilityAlert) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *RepositoryVulnerabilityAlert) GetDismissedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *RepositoryVulnerabilityAlert) GetDismisser() *User { _ = "STUB: not implemented"; return nil }

func (r *RepositoryVulnerabilityAlert) GetDismissReason() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryVulnerabilityAlert) GetExternalIdentifier() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryVulnerabilityAlert) GetExternalReference() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryVulnerabilityAlert) GetFixedIn() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryVulnerabilityAlert) GetGitHubSecurityAdvisoryID() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryVulnerabilityAlert) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RepositoryVulnerabilityAlert) GetSeverity() string { _ = "STUB: not implemented"; return "" }

func (r *RepositoryVulnerabilityAlertEvent) GetAction() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RepositoryVulnerabilityAlertEvent) GetAlert() *RepositoryVulnerabilityAlert {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVulnerabilityAlertEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVulnerabilityAlertEvent) GetOrg() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVulnerabilityAlertEvent) GetRepository() *Repository {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepositoryVulnerabilityAlertEvent) GetSender() *User {
	_ = "STUB: not implemented"
	return nil
}

func (r *RepoStats) GetForkRepos() int { _ = "STUB: not implemented"; return 0 }

func (r *RepoStats) GetOrgRepos() int { _ = "STUB: not implemented"; return 0 }

func (r *RepoStats) GetRootRepos() int { _ = "STUB: not implemented"; return 0 }

func (r *RepoStats) GetTotalPushes() int { _ = "STUB: not implemented"; return 0 }

func (r *RepoStats) GetTotalRepos() int { _ = "STUB: not implemented"; return 0 }

func (r *RepoStats) GetTotalWikis() int { _ = "STUB: not implemented"; return 0 }

func (r *RepoStatus) GetAvatarURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepoStatus) GetContext() string { _ = "STUB: not implemented"; return "" }

func (r *RepoStatus) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (r *RepoStatus) GetCreator() *User { _ = "STUB: not implemented"; return nil }

func (r *RepoStatus) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (r *RepoStatus) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RepoStatus) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (r *RepoStatus) GetState() string { _ = "STUB: not implemented"; return "" }

func (r *RepoStatus) GetTargetURL() string { _ = "STUB: not implemented"; return "" }

func (r *RepoStatus) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (r *RepoStatus) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *RequestedAction) GetIdentifier() string { _ = "STUB: not implemented"; return "" }

func (r *RequireCodeOwnerReviewChanges) GetFrom() bool { _ = "STUB: not implemented"; return false }

func (r *RequiredConversationResolution) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (r *RequiredConversationResolutionLevelChanges) GetFrom() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RequiredDeploymentsBranchRule) GetParameters() RequiredDeploymentsRuleParameters {
	_ = "STUB: not implemented"
	return *new(RequiredDeploymentsRuleParameters)
}

func (r *RequiredDeploymentsEnforcementLevelChanges) GetFrom() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RequiredDeploymentsRuleParameters) GetRequiredDeploymentEnvironments() []string {
	_ = "STUB: not implemented"
	return nil
}

func (r *RequiredReviewer) GetReviewer() any { _ = "STUB: not implemented"; return *new(any) }

func (r *RequiredReviewer) GetType() string { _ = "STUB: not implemented"; return "" }

func (r *RequiredStatusCheck) GetAppID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RequiredStatusCheck) GetContext() string { _ = "STUB: not implemented"; return "" }

func (r *RequiredStatusChecks) GetChecks() []*RequiredStatusCheck {
	_ = "STUB: not implemented"
	return nil
}

func (r *RequiredStatusChecks) GetContexts() []string { _ = "STUB: not implemented"; return nil }

func (r *RequiredStatusChecks) GetContextsURL() string { _ = "STUB: not implemented"; return "" }

func (r *RequiredStatusChecks) GetStrict() bool { _ = "STUB: not implemented"; return false }

func (r *RequiredStatusChecks) GetURL() string { _ = "STUB: not implemented"; return "" }

func (r *RequiredStatusChecksBranchRule) GetParameters() RequiredStatusChecksRuleParameters {
	_ = "STUB: not implemented"
	return *new(RequiredStatusChecksRuleParameters)
}

func (r *RequiredStatusChecksChanges) GetFrom() []string { _ = "STUB: not implemented"; return nil }

func (r *RequiredStatusChecksEnforcementLevelChanges) GetFrom() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RequiredStatusChecksRequest) GetChecks() []*RequiredStatusCheck {
	_ = "STUB: not implemented"
	return nil
}

func (r *RequiredStatusChecksRequest) GetContexts() []string { _ = "STUB: not implemented"; return nil }

func (r *RequiredStatusChecksRequest) GetStrict() bool { _ = "STUB: not implemented"; return false }

func (r *RequiredStatusChecksRuleParameters) GetDoNotEnforceOnCreate() bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RequiredStatusChecksRuleParameters) GetRequiredStatusChecks() []*RuleStatusCheck {
	_ = "STUB: not implemented"
	return nil
}

func (r *RequiredStatusChecksRuleParameters) GetStrictRequiredStatusChecksPolicy() bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RequireLastPushApprovalChanges) GetFrom() bool { _ = "STUB: not implemented"; return false }

func (r *RequireLinearHistory) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (r *Response) GetAfter() string { _ = "STUB: not implemented"; return "" }

func (r *Response) GetBefore() string { _ = "STUB: not implemented"; return "" }

func (r *Response) GetCursor() string { _ = "STUB: not implemented"; return "" }

func (r *Response) GetFirstPage() int { _ = "STUB: not implemented"; return 0 }

func (r *Response) GetLastPage() int { _ = "STUB: not implemented"; return 0 }

func (r *Response) GetNextPage() int { _ = "STUB: not implemented"; return 0 }

func (r *Response) GetNextPageToken() string { _ = "STUB: not implemented"; return "" }

func (r *Response) GetPrevPage() int { _ = "STUB: not implemented"; return 0 }

func (r *Response) GetRate() Rate { _ = "STUB: not implemented"; return *new(Rate) }

func (r *Response) GetTokenExpiration() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (r *ReviewCustomDeploymentProtectionRuleRequest) GetComment() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *ReviewCustomDeploymentProtectionRuleRequest) GetEnvironmentName() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *ReviewCustomDeploymentProtectionRuleRequest) GetState() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *Reviewers) GetTeams() []*Team { _ = "STUB: not implemented"; return nil }

func (r *Reviewers) GetUsers() []*User { _ = "STUB: not implemented"; return nil }

func (r *ReviewersRequest) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (r *ReviewersRequest) GetReviewers() []string { _ = "STUB: not implemented"; return nil }

func (r *ReviewersRequest) GetTeamReviewers() []string { _ = "STUB: not implemented"; return nil }

func (r *ReviewPersonalAccessTokenRequestOptions) GetAction() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *ReviewPersonalAccessTokenRequestOptions) GetReason() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *Rule) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (r *Rule) GetFullDescription() string { _ = "STUB: not implemented"; return "" }

func (r *Rule) GetHelp() string { _ = "STUB: not implemented"; return "" }

func (r *Rule) GetID() string { _ = "STUB: not implemented"; return "" }

func (r *Rule) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *Rule) GetSecuritySeverityLevel() string { _ = "STUB: not implemented"; return "" }

func (r *Rule) GetSeverity() string { _ = "STUB: not implemented"; return "" }

func (r *Rule) GetTags() []string { _ = "STUB: not implemented"; return nil }

func (r *RuleCodeScanningTool) GetAlertsThreshold() CodeScanningAlertsThreshold {
	_ = "STUB: not implemented"
	return *new(CodeScanningAlertsThreshold)
}

func (r *RuleCodeScanningTool) GetSecurityAlertsThreshold() CodeScanningSecurityAlertsThreshold {
	_ = "STUB: not implemented"
	return *new(CodeScanningSecurityAlertsThreshold)
}

func (r *RuleCodeScanningTool) GetTool() string { _ = "STUB: not implemented"; return "" }

func (r *RulesetRequiredReviewer) GetFilePatterns() []string { _ = "STUB: not implemented"; return nil }

func (r *RulesetRequiredReviewer) GetMinimumApprovals() int { _ = "STUB: not implemented"; return 0 }

func (r *RulesetRequiredReviewer) GetReviewer() *RulesetReviewer {
	_ = "STUB: not implemented"
	return nil
}

func (r *RulesetReviewer) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RulesetReviewer) GetType() *RulesetReviewerType { _ = "STUB: not implemented"; return nil }

func (r *RuleStatusCheck) GetContext() string { _ = "STUB: not implemented"; return "" }

func (r *RuleStatusCheck) GetIntegrationID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RuleWorkflow) GetPath() string { _ = "STUB: not implemented"; return "" }

func (r *RuleWorkflow) GetRef() string { _ = "STUB: not implemented"; return "" }

func (r *RuleWorkflow) GetRepositoryID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RuleWorkflow) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (r *Runner) GetBusy() bool { _ = "STUB: not implemented"; return false }

func (r *Runner) GetEphemeral() bool { _ = "STUB: not implemented"; return false }

func (r *Runner) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Runner) GetLabels() []*RunnerLabels { _ = "STUB: not implemented"; return nil }

func (r *Runner) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *Runner) GetOS() string { _ = "STUB: not implemented"; return "" }

func (r *Runner) GetRunnerGroupID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Runner) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (r *Runner) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (r *RunnerApplicationDownload) GetArchitecture() string { _ = "STUB: not implemented"; return "" }

func (r *RunnerApplicationDownload) GetDownloadURL() string { _ = "STUB: not implemented"; return "" }

func (r *RunnerApplicationDownload) GetFilename() string { _ = "STUB: not implemented"; return "" }

func (r *RunnerApplicationDownload) GetOS() string { _ = "STUB: not implemented"; return "" }

func (r *RunnerApplicationDownload) GetSHA256Checksum() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RunnerApplicationDownload) GetTempDownloadToken() string {
	_ = "STUB: not implemented"
	return ""
}

func (r *RunnerGroup) GetAllowsPublicRepositories() bool { _ = "STUB: not implemented"; return false }

func (r *RunnerGroup) GetDefault() bool { _ = "STUB: not implemented"; return false }

func (r *RunnerGroup) GetHostedRunnersURL() string { _ = "STUB: not implemented"; return "" }

func (r *RunnerGroup) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RunnerGroup) GetInherited() bool { _ = "STUB: not implemented"; return false }

func (r *RunnerGroup) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *RunnerGroup) GetNetworkConfigurationID() string { _ = "STUB: not implemented"; return "" }

func (r *RunnerGroup) GetRestrictedToWorkflows() bool { _ = "STUB: not implemented"; return false }

func (r *RunnerGroup) GetRunnersURL() string { _ = "STUB: not implemented"; return "" }

func (r *RunnerGroup) GetSelectedRepositoriesURL() string { _ = "STUB: not implemented"; return "" }

func (r *RunnerGroup) GetSelectedWorkflows() []string { _ = "STUB: not implemented"; return nil }

func (r *RunnerGroup) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (r *RunnerGroup) GetWorkflowRestrictionsReadOnly() bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RunnerGroups) GetRunnerGroups() []*RunnerGroup { _ = "STUB: not implemented"; return nil }

func (r *RunnerGroups) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (r *RunnerLabels) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (r *RunnerLabels) GetName() string { _ = "STUB: not implemented"; return "" }

func (r *RunnerLabels) GetType() string { _ = "STUB: not implemented"; return "" }

func (r *Runners) GetRunners() []*Runner { _ = "STUB: not implemented"; return nil }

func (r *Runners) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (s *SarifAnalysis) GetCheckoutURI() string { _ = "STUB: not implemented"; return "" }

func (s *SarifAnalysis) GetCommitSHA() string { _ = "STUB: not implemented"; return "" }

func (s *SarifAnalysis) GetRef() string { _ = "STUB: not implemented"; return "" }

func (s *SarifAnalysis) GetSarif() string { _ = "STUB: not implemented"; return "" }

func (s *SarifAnalysis) GetStartedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *SarifAnalysis) GetToolName() string { _ = "STUB: not implemented"; return "" }

func (s *SarifAnalysis) GetValidate() bool { _ = "STUB: not implemented"; return false }

func (s *SarifID) GetID() string { _ = "STUB: not implemented"; return "" }

func (s *SarifID) GetURL() string { _ = "STUB: not implemented"; return "" }

func (s *SARIFUpload) GetAnalysesURL() string { _ = "STUB: not implemented"; return "" }

func (s *SARIFUpload) GetProcessingStatus() string { _ = "STUB: not implemented"; return "" }

func (s *SBOM) GetSBOM() *SBOMInfo { _ = "STUB: not implemented"; return nil }

func (s *SBOMInfo) GetCreationInfo() *CreationInfo { _ = "STUB: not implemented"; return nil }

func (s *SBOMInfo) GetDataLicense() string { _ = "STUB: not implemented"; return "" }

func (s *SBOMInfo) GetDocumentDescribes() []string { _ = "STUB: not implemented"; return nil }

func (s *SBOMInfo) GetDocumentNamespace() string { _ = "STUB: not implemented"; return "" }

func (s *SBOMInfo) GetName() string { _ = "STUB: not implemented"; return "" }

func (s *SBOMInfo) GetPackages() []*RepoDependencies { _ = "STUB: not implemented"; return nil }

func (s *SBOMInfo) GetRelationships() []*SBOMRelationship { _ = "STUB: not implemented"; return nil }

func (s *SBOMInfo) GetSPDXID() string { _ = "STUB: not implemented"; return "" }

func (s *SBOMInfo) GetSPDXVersion() string { _ = "STUB: not implemented"; return "" }

func (s *SBOMRelationship) GetRelatedSPDXElement() string { _ = "STUB: not implemented"; return "" }

func (s *SBOMRelationship) GetRelationshipType() string { _ = "STUB: not implemented"; return "" }

func (s *SBOMRelationship) GetSPDXElementID() string { _ = "STUB: not implemented"; return "" }

func (s *ScanningAnalysis) GetAnalysisKey() string { _ = "STUB: not implemented"; return "" }

func (s *ScanningAnalysis) GetCategory() string { _ = "STUB: not implemented"; return "" }

func (s *ScanningAnalysis) GetCommitSHA() string { _ = "STUB: not implemented"; return "" }

func (s *ScanningAnalysis) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *ScanningAnalysis) GetDeletable() bool { _ = "STUB: not implemented"; return false }

func (s *ScanningAnalysis) GetEnvironment() string { _ = "STUB: not implemented"; return "" }

func (s *ScanningAnalysis) GetError() string { _ = "STUB: not implemented"; return "" }

func (s *ScanningAnalysis) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (s *ScanningAnalysis) GetRef() string { _ = "STUB: not implemented"; return "" }

func (s *ScanningAnalysis) GetResultsCount() int { _ = "STUB: not implemented"; return 0 }

func (s *ScanningAnalysis) GetRulesCount() int { _ = "STUB: not implemented"; return 0 }

func (s *ScanningAnalysis) GetSarifID() string { _ = "STUB: not implemented"; return "" }

func (s *ScanningAnalysis) GetTool() *Tool { _ = "STUB: not implemented"; return nil }

func (s *ScanningAnalysis) GetURL() string { _ = "STUB: not implemented"; return "" }

func (s *ScanningAnalysis) GetWarning() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseAttribute) GetOperations() []*SCIMEnterpriseAttributeOperation {
	_ = "STUB: not implemented"
	return nil
}

func (s *SCIMEnterpriseAttribute) GetSchemas() []string { _ = "STUB: not implemented"; return nil }

func (s *SCIMEnterpriseAttributeOperation) GetOp() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseAttributeOperation) GetPath() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseAttributeOperation) GetValue() any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (s *SCIMEnterpriseDisplayReference) GetDisplay() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseDisplayReference) GetRef() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseDisplayReference) GetValue() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseGroupAttributes) GetDisplayName() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SCIMEnterpriseGroupAttributes) GetExternalID() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SCIMEnterpriseGroupAttributes) GetID() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseGroupAttributes) GetMembers() []*SCIMEnterpriseDisplayReference {
	_ = "STUB: not implemented"
	return nil
}

func (s *SCIMEnterpriseGroupAttributes) GetMeta() *SCIMEnterpriseMeta {
	_ = "STUB: not implemented"
	return nil
}

func (s *SCIMEnterpriseGroupAttributes) GetSchemas() []string {
	_ = "STUB: not implemented"
	return nil
}

func (s *SCIMEnterpriseGroups) GetItemsPerPage() int { _ = "STUB: not implemented"; return 0 }

func (s *SCIMEnterpriseGroups) GetResources() []*SCIMEnterpriseGroupAttributes {
	_ = "STUB: not implemented"
	return nil
}

func (s *SCIMEnterpriseGroups) GetSchemas() []string { _ = "STUB: not implemented"; return nil }

func (s *SCIMEnterpriseGroups) GetStartIndex() int { _ = "STUB: not implemented"; return 0 }

func (s *SCIMEnterpriseGroups) GetTotalResults() int { _ = "STUB: not implemented"; return 0 }

func (s *SCIMEnterpriseMeta) GetCreated() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *SCIMEnterpriseMeta) GetLastModified() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *SCIMEnterpriseMeta) GetLocation() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseMeta) GetResourceType() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUserAttributes) GetActive() bool { _ = "STUB: not implemented"; return false }

func (s *SCIMEnterpriseUserAttributes) GetDisplayName() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SCIMEnterpriseUserAttributes) GetEmails() []*SCIMEnterpriseUserEmail {
	_ = "STUB: not implemented"
	return nil
}

func (s *SCIMEnterpriseUserAttributes) GetExternalID() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUserAttributes) GetGroups() []*SCIMEnterpriseDisplayReference {
	_ = "STUB: not implemented"
	return nil
}

func (s *SCIMEnterpriseUserAttributes) GetID() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUserAttributes) GetMeta() *SCIMEnterpriseMeta {
	_ = "STUB: not implemented"
	return nil
}

func (s *SCIMEnterpriseUserAttributes) GetName() *SCIMEnterpriseUserName {
	_ = "STUB: not implemented"
	return nil
}

func (s *SCIMEnterpriseUserAttributes) GetRoles() []*SCIMEnterpriseUserRole {
	_ = "STUB: not implemented"
	return nil
}

func (s *SCIMEnterpriseUserAttributes) GetSchemas() []string { _ = "STUB: not implemented"; return nil }

func (s *SCIMEnterpriseUserAttributes) GetUserName() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUserEmail) GetPrimary() bool { _ = "STUB: not implemented"; return false }

func (s *SCIMEnterpriseUserEmail) GetType() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUserEmail) GetValue() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUserName) GetFamilyName() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUserName) GetFormatted() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUserName) GetGivenName() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUserName) GetMiddleName() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUserRole) GetDisplay() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUserRole) GetPrimary() bool { _ = "STUB: not implemented"; return false }

func (s *SCIMEnterpriseUserRole) GetType() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUserRole) GetValue() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMEnterpriseUsers) GetItemsPerPage() int { _ = "STUB: not implemented"; return 0 }

func (s *SCIMEnterpriseUsers) GetResources() []*SCIMEnterpriseUserAttributes {
	_ = "STUB: not implemented"
	return nil
}

func (s *SCIMEnterpriseUsers) GetSchemas() []string { _ = "STUB: not implemented"; return nil }

func (s *SCIMEnterpriseUsers) GetStartIndex() int { _ = "STUB: not implemented"; return 0 }

func (s *SCIMEnterpriseUsers) GetTotalResults() int { _ = "STUB: not implemented"; return 0 }

func (s *SCIMMeta) GetCreated() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *SCIMMeta) GetLastModified() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *SCIMMeta) GetLocation() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMMeta) GetResourceType() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMProvisionedIdentities) GetItemsPerPage() int { _ = "STUB: not implemented"; return 0 }

func (s *SCIMProvisionedIdentities) GetResources() []*SCIMUserAttributes {
	_ = "STUB: not implemented"
	return nil
}

func (s *SCIMProvisionedIdentities) GetSchemas() []string { _ = "STUB: not implemented"; return nil }

func (s *SCIMProvisionedIdentities) GetStartIndex() int { _ = "STUB: not implemented"; return 0 }

func (s *SCIMProvisionedIdentities) GetTotalResults() int { _ = "STUB: not implemented"; return 0 }

func (s *SCIMUserAttributes) GetActive() bool { _ = "STUB: not implemented"; return false }

func (s *SCIMUserAttributes) GetDisplayName() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMUserAttributes) GetEmails() []*SCIMUserEmail { _ = "STUB: not implemented"; return nil }

func (s *SCIMUserAttributes) GetExternalID() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMUserAttributes) GetGroups() []string { _ = "STUB: not implemented"; return nil }

func (s *SCIMUserAttributes) GetID() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMUserAttributes) GetMeta() *SCIMMeta { _ = "STUB: not implemented"; return nil }

func (s *SCIMUserAttributes) GetName() SCIMUserName {
	_ = "STUB: not implemented"
	return *new(SCIMUserName)
}

func (s *SCIMUserAttributes) GetRoles() []*SCIMUserRole { _ = "STUB: not implemented"; return nil }

func (s *SCIMUserAttributes) GetSchemas() []string { _ = "STUB: not implemented"; return nil }

func (s *SCIMUserAttributes) GetUserName() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMUserEmail) GetPrimary() bool { _ = "STUB: not implemented"; return false }

func (s *SCIMUserEmail) GetType() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMUserEmail) GetValue() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMUserName) GetFamilyName() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMUserName) GetFormatted() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMUserName) GetGivenName() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMUserRole) GetDisplay() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMUserRole) GetPrimary() bool { _ = "STUB: not implemented"; return false }

func (s *SCIMUserRole) GetType() string { _ = "STUB: not implemented"; return "" }

func (s *SCIMUserRole) GetValue() string { _ = "STUB: not implemented"; return "" }

func (s *SearchOptions) GetAdvancedSearch() bool { _ = "STUB: not implemented"; return false }

func (s *SearchOptions) GetOrder() string { _ = "STUB: not implemented"; return "" }

func (s *SearchOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (s *SearchOptions) GetTextMatch() bool { _ = "STUB: not implemented"; return false }

func (s *SeatAssignments) GetSeatsCreated() int { _ = "STUB: not implemented"; return 0 }

func (s *SeatCancellations) GetSeatsCancelled() int { _ = "STUB: not implemented"; return 0 }

func (s *Secret) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *Secret) GetName() string { _ = "STUB: not implemented"; return "" }

func (s *Secret) GetSelectedRepositoriesURL() string { _ = "STUB: not implemented"; return "" }

func (s *Secret) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *Secret) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (s *SecretOrgRequest) GetEncryptedValue() string { _ = "STUB: not implemented"; return "" }

func (s *SecretOrgRequest) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (s *SecretOrgRequest) GetSelectedRepositoryIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretOrgRequest) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (s *SecretRequest) GetEncryptedValue() string { _ = "STUB: not implemented"; return "" }

func (s *SecretRequest) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (s *Secrets) GetSecrets() []*Secret { _ = "STUB: not implemented"; return nil }

func (s *Secrets) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (s *SecretScanning) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlert) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *SecretScanningAlert) GetFirstLocationDetected() *SecretScanningAlertLocationDetails {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlert) GetHasMoreLocations() bool { _ = "STUB: not implemented"; return false }

func (s *SecretScanningAlert) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlert) GetIsBase64Encoded() bool { _ = "STUB: not implemented"; return false }

func (s *SecretScanningAlert) GetLocationsURL() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlert) GetMultiRepo() bool { _ = "STUB: not implemented"; return false }

func (s *SecretScanningAlert) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (s *SecretScanningAlert) GetPubliclyLeaked() bool { _ = "STUB: not implemented"; return false }

func (s *SecretScanningAlert) GetPushProtectionBypassed() bool {
	_ = "STUB: not implemented"
	return false
}

func (s *SecretScanningAlert) GetPushProtectionBypassedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *SecretScanningAlert) GetPushProtectionBypassedBy() *User {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlert) GetPushProtectionBypassRequestComment() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlert) GetPushProtectionBypassRequestHTMLURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlert) GetPushProtectionBypassRequestReviewer() *User {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlert) GetPushProtectionBypassRequestReviewerComment() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlert) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (s *SecretScanningAlert) GetResolution() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlert) GetResolutionComment() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlert) GetResolvedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *SecretScanningAlert) GetResolvedBy() *User { _ = "STUB: not implemented"; return nil }

func (s *SecretScanningAlert) GetSecret() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlert) GetSecretType() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlert) GetSecretTypeDisplayName() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlert) GetState() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlert) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *SecretScanningAlert) GetURL() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlert) GetValidity() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlertEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlertEvent) GetAlert() *SecretScanningAlert {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlertEvent) GetEnterprise() *Enterprise {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlertEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlertEvent) GetOrganization() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlertEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (s *SecretScanningAlertEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (s *SecretScanningAlertListOptions) GetDirection() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlertListOptions) GetIsMultiRepo() bool {
	_ = "STUB: not implemented"
	return false
}

func (s *SecretScanningAlertListOptions) GetIsPubliclyLeaked() bool {
	_ = "STUB: not implemented"
	return false
}

func (s *SecretScanningAlertListOptions) GetResolution() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlertListOptions) GetSecretType() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlertListOptions) GetSort() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlertListOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlertListOptions) GetValidity() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlertLocation) GetDetails() *SecretScanningAlertLocationDetails {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlertLocation) GetType() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlertLocationDetails) GetBlobSHA() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlertLocationDetails) GetBlobURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlertLocationDetails) GetCommitSHA() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlertLocationDetails) GetCommitURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlertLocationDetails) GetEndColumn() int {
	_ = "STUB: not implemented"
	return 0
}

func (s *SecretScanningAlertLocationDetails) GetEndLine() int { _ = "STUB: not implemented"; return 0 }

func (s *SecretScanningAlertLocationDetails) GetPath() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlertLocationDetails) GetPullRequestCommentURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlertLocationDetails) GetStartColumn() int {
	_ = "STUB: not implemented"
	return 0
}

func (s *SecretScanningAlertLocationDetails) GetStartline() int {
	_ = "STUB: not implemented"
	return 0
}

func (s *SecretScanningAlertLocationEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningAlertLocationEvent) GetAlert() *SecretScanningAlert {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlertLocationEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlertLocationEvent) GetLocation() *SecretScanningAlertLocation {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlertLocationEvent) GetOrganization() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlertLocationEvent) GetRepo() *Repository {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningAlertLocationEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (s *SecretScanningAlertUpdateOptions) GetResolution() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlertUpdateOptions) GetResolutionComment() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningAlertUpdateOptions) GetState() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningCustomPatternSetting) GetCustomPatternVersion() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningCustomPatternSetting) GetPushProtectionSetting() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningCustomPatternSetting) GetTokenType() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningDelegatedBypassOptions) GetReviewers() []*BypassReviewer {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningPatternConfigs) GetCustomPatternOverrides() []*SecretScanningPatternOverride {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningPatternConfigs) GetPatternConfigVersion() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningPatternConfigs) GetProviderPatternOverrides() []*SecretScanningPatternOverride {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningPatternConfigsUpdate) GetPatternConfigVersion() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningPatternConfigsUpdateOptions) GetCustomPatternSettings() []*SecretScanningCustomPatternSetting {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningPatternConfigsUpdateOptions) GetPatternConfigVersion() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningPatternConfigsUpdateOptions) GetProviderPatternSettings() []*SecretScanningProviderPatternSetting {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningPatternOverride) GetAlertTotal() int { _ = "STUB: not implemented"; return 0 }

func (s *SecretScanningPatternOverride) GetAlertTotalPercentage() int {
	_ = "STUB: not implemented"
	return 0
}

func (s *SecretScanningPatternOverride) GetBypassrate() int { _ = "STUB: not implemented"; return 0 }

func (s *SecretScanningPatternOverride) GetCustomPatternVersion() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningPatternOverride) GetDefaultSetting() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningPatternOverride) GetDisplayName() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningPatternOverride) GetEnterpriseSetting() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningPatternOverride) GetFalsePositiveRate() int {
	_ = "STUB: not implemented"
	return 0
}

func (s *SecretScanningPatternOverride) GetFalsePositives() int {
	_ = "STUB: not implemented"
	return 0
}

func (s *SecretScanningPatternOverride) GetSetting() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningPatternOverride) GetSlug() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningPatternOverride) GetTokenType() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningProviderPatternSetting) GetPushProtectionSetting() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningProviderPatternSetting) GetTokenType() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SecretScanningPushProtection) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (s *SecretScanningScanHistory) GetBackfillScans() []*SecretsScan {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningScanHistory) GetCustomPatternBackfillScans() []*CustomPatternBackfillScan {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningScanHistory) GetIncrementalScans() []*SecretsScan {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningScanHistory) GetPatternUpdateScans() []*SecretsScan {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningValidityChecks) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (s *SecretsScan) GetCompletedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *SecretsScan) GetStartedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *SecretsScan) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (s *SecretsScan) GetType() string { _ = "STUB: not implemented"; return "" }

func (s *SecurityAdvisory) GetAuthor() *User { _ = "STUB: not implemented"; return nil }

func (s *SecurityAdvisory) GetClosedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *SecurityAdvisory) GetCollaboratingTeams() []*Team { _ = "STUB: not implemented"; return nil }

func (s *SecurityAdvisory) GetCollaboratingUsers() []*User { _ = "STUB: not implemented"; return nil }

func (s *SecurityAdvisory) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *SecurityAdvisory) GetCredits() []*RepoAdvisoryCredit {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAdvisory) GetCreditsDetailed() []*RepoAdvisoryCreditDetailed {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAdvisory) GetCVEID() string { _ = "STUB: not implemented"; return "" }

func (s *SecurityAdvisory) GetCVSS() *AdvisoryCVSS { _ = "STUB: not implemented"; return nil }

func (s *SecurityAdvisory) GetCWEIDs() []string { _ = "STUB: not implemented"; return nil }

func (s *SecurityAdvisory) GetCWEs() []*AdvisoryCWEs { _ = "STUB: not implemented"; return nil }

func (s *SecurityAdvisory) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (s *SecurityAdvisory) GetGHSAID() string { _ = "STUB: not implemented"; return "" }

func (s *SecurityAdvisory) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (s *SecurityAdvisory) GetIdentifiers() []*AdvisoryIdentifier {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAdvisory) GetPrivateFork() *Repository { _ = "STUB: not implemented"; return nil }

func (s *SecurityAdvisory) GetPublishedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *SecurityAdvisory) GetPublisher() *User { _ = "STUB: not implemented"; return nil }

func (s *SecurityAdvisory) GetReferences() []*AdvisoryReference {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAdvisory) GetSeverity() string { _ = "STUB: not implemented"; return "" }

func (s *SecurityAdvisory) GetState() string { _ = "STUB: not implemented"; return "" }

func (s *SecurityAdvisory) GetSubmission() *SecurityAdvisorySubmission {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAdvisory) GetSummary() string { _ = "STUB: not implemented"; return "" }

func (s *SecurityAdvisory) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *SecurityAdvisory) GetURL() string { _ = "STUB: not implemented"; return "" }

func (s *SecurityAdvisory) GetVulnerabilities() []*AdvisoryVulnerability {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAdvisory) GetWithdrawnAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *SecurityAdvisoryEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (s *SecurityAdvisoryEvent) GetEnterprise() *Enterprise { _ = "STUB: not implemented"; return nil }

func (s *SecurityAdvisoryEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAdvisoryEvent) GetOrganization() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAdvisoryEvent) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (s *SecurityAdvisoryEvent) GetSecurityAdvisory() *SecurityAdvisory {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAdvisoryEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (s *SecurityAdvisorySubmission) GetAccepted() bool { _ = "STUB: not implemented"; return false }

func (s *SecurityAndAnalysis) GetAdvancedSecurity() *AdvancedSecurity {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysis) GetCodeSecurity() *CodeSecurity {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysis) GetDependabotSecurityUpdates() *DependabotSecurityUpdates {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysis) GetSecretScanning() *SecretScanning {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysis) GetSecretScanningPushProtection() *SecretScanningPushProtection {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysis) GetSecretScanningValidityChecks() *SecretScanningValidityChecks {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysisChange) GetFrom() *SecurityAndAnalysisChangeFrom {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysisChangeFrom) GetSecurityAndAnalysis() *SecurityAndAnalysis {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysisEvent) GetChanges() *SecurityAndAnalysisChange {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysisEvent) GetEnterprise() *Enterprise {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysisEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysisEvent) GetOrganization() *Organization {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysisEvent) GetRepository() *Repository {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAndAnalysisEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (s *SelectedReposList) GetRepositories() []*Repository { _ = "STUB: not implemented"; return nil }

func (s *SelectedReposList) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (s *SelfHostedRunnersAllowedRepos) GetRepositories() []*Repository {
	_ = "STUB: not implemented"
	return nil
}

func (s *SelfHostedRunnersAllowedRepos) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (s *SelfHostedRunnersSettingsOrganization) GetEnabledRepositories() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SelfHostedRunnersSettingsOrganization) GetSelectedRepositoriesURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SelfHostedRunnersSettingsOrganizationOpt) GetEnabledRepositories() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SelfHostRunnerPermissionsEnterprise) GetDisableSelfHostedRunnersForAllOrgs() bool {
	_ = "STUB: not implemented"
	return false
}

func (s *ServerInstanceProperties) GetServerInstances() *ServerInstances {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServerInstances) GetItems() *ServiceInstanceItems { _ = "STUB: not implemented"; return nil }

func (s *ServerInstances) GetType() string { _ = "STUB: not implemented"; return "" }

func (s *ServerItemProperties) GetHostname() string { _ = "STUB: not implemented"; return "" }

func (s *ServerItemProperties) GetLastSync() *LastLicenseSync {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServerItemProperties) GetServerID() string { _ = "STUB: not implemented"; return "" }

func (s *ServiceInstanceItems) GetProperties() *ServerItemProperties {
	_ = "STUB: not implemented"
	return nil
}

func (s *ServiceInstanceItems) GetType() string { _ = "STUB: not implemented"; return "" }

func (s *SetOrgAccessRunnerGroupRequest) GetSelectedOrganizationIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (s *SetRepoAccessRunnerGroupRequest) GetSelectedRepositoryIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (s *SetRunnerGroupRunnersRequest) GetRunners() []int64 { _ = "STUB: not implemented"; return nil }

func (s *SignatureRequirementEnforcementLevelChanges) GetFrom() string {
	_ = "STUB: not implemented"
	return ""
}

func (s *SignaturesProtectedBranch) GetEnabled() bool { _ = "STUB: not implemented"; return false }

func (s *SignaturesProtectedBranch) GetURL() string { _ = "STUB: not implemented"; return "" }

func (s *SignatureVerification) GetPayload() string { _ = "STUB: not implemented"; return "" }

func (s *SignatureVerification) GetReason() string { _ = "STUB: not implemented"; return "" }

func (s *SignatureVerification) GetSignature() string { _ = "STUB: not implemented"; return "" }

func (s *SignatureVerification) GetVerified() bool { _ = "STUB: not implemented"; return false }

func (s *SimplePatternRuleParameters) GetNegate() bool { _ = "STUB: not implemented"; return false }

func (s *SimplePatternRuleParameters) GetPattern() string { _ = "STUB: not implemented"; return "" }

func (s *SocialAccount) GetProvider() string { _ = "STUB: not implemented"; return "" }

func (s *SocialAccount) GetURL() string { _ = "STUB: not implemented"; return "" }

func (s *Source) GetActor() *User { _ = "STUB: not implemented"; return nil }

func (s *Source) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (s *Source) GetIssue() *Issue { _ = "STUB: not implemented"; return nil }

func (s *Source) GetType() string { _ = "STUB: not implemented"; return "" }

func (s *Source) GetURL() string { _ = "STUB: not implemented"; return "" }

func (s *SourceImportAuthor) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (s *SourceImportAuthor) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (s *SourceImportAuthor) GetImportURL() string { _ = "STUB: not implemented"; return "" }

func (s *SourceImportAuthor) GetName() string { _ = "STUB: not implemented"; return "" }

func (s *SourceImportAuthor) GetRemoteID() string { _ = "STUB: not implemented"; return "" }

func (s *SourceImportAuthor) GetRemoteName() string { _ = "STUB: not implemented"; return "" }

func (s *SourceImportAuthor) GetURL() string { _ = "STUB: not implemented"; return "" }

func (s *SplunkConfig) GetDomain() string { _ = "STUB: not implemented"; return "" }

func (s *SplunkConfig) GetEncryptedToken() string { _ = "STUB: not implemented"; return "" }

func (s *SplunkConfig) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (s *SplunkConfig) GetPort() uint16 { _ = "STUB: not implemented"; return 0 }

func (s *SplunkConfig) GetSSLVerify() bool { _ = "STUB: not implemented"; return false }

func (s *SponsorshipChanges) GetPrivacyLevel() string { _ = "STUB: not implemented"; return "" }

func (s *SponsorshipChanges) GetTier() *SponsorshipTier { _ = "STUB: not implemented"; return nil }

func (s *SponsorshipEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (s *SponsorshipEvent) GetChanges() *SponsorshipChanges { _ = "STUB: not implemented"; return nil }

func (s *SponsorshipEvent) GetEffectiveDate() string { _ = "STUB: not implemented"; return "" }

func (s *SponsorshipEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (s *SponsorshipEvent) GetOrganization() *Organization { _ = "STUB: not implemented"; return nil }

func (s *SponsorshipEvent) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (s *SponsorshipEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (s *SponsorshipTier) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (s *SSHKeyOptions) GetKey() string { _ = "STUB: not implemented"; return "" }

func (s *SSHKeyStatus) GetHostname() string { _ = "STUB: not implemented"; return "" }

func (s *SSHKeyStatus) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (s *SSHKeyStatus) GetModified() bool { _ = "STUB: not implemented"; return false }

func (s *SSHKeyStatus) GetUUID() string { _ = "STUB: not implemented"; return "" }

func (s *SSHSigningKey) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *SSHSigningKey) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (s *SSHSigningKey) GetKey() string { _ = "STUB: not implemented"; return "" }

func (s *SSHSigningKey) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (s *StarEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (s *StarEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (s *StarEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (s *StarEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (s *StarEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (s *StarEvent) GetStarredAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *Stargazer) GetStarredAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *Stargazer) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (s *StarredRepository) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (s *StarredRepository) GetStarredAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (s *StatusEvent) GetBranches() []*Branch { _ = "STUB: not implemented"; return nil }

func (s *StatusEvent) GetCommit() *RepositoryCommit { _ = "STUB: not implemented"; return nil }

func (s *StatusEvent) GetContext() string { _ = "STUB: not implemented"; return "" }

func (s *StatusEvent) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *StatusEvent) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (s *StatusEvent) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (s *StatusEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (s *StatusEvent) GetName() string { _ = "STUB: not implemented"; return "" }

func (s *StatusEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (s *StatusEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (s *StatusEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (s *StatusEvent) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (s *StatusEvent) GetState() string { _ = "STUB: not implemented"; return "" }

func (s *StatusEvent) GetTargetURL() string { _ = "STUB: not implemented"; return "" }

func (s *StatusEvent) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *StorageBilling) GetDaysLeftInBillingCycle() int { _ = "STUB: not implemented"; return 0 }

func (s *StorageBilling) GetEstimatedPaidStorageForMonth() int { _ = "STUB: not implemented"; return 0 }

func (s *StorageBilling) GetEstimatedStorageForMonth() int { _ = "STUB: not implemented"; return 0 }

func (s *SubIssueRequest) GetAfterID() int64 { _ = "STUB: not implemented"; return 0 }

func (s *SubIssueRequest) GetBeforeID() int64 { _ = "STUB: not implemented"; return 0 }

func (s *SubIssueRequest) GetReplaceParent() bool { _ = "STUB: not implemented"; return false }

func (s *SubIssueRequest) GetSubIssueID() int64 { _ = "STUB: not implemented"; return 0 }

func (s *SubIssuesSummary) GetCompleted() int { _ = "STUB: not implemented"; return 0 }

func (s *SubIssuesSummary) GetPercentCompleted() int { _ = "STUB: not implemented"; return 0 }

func (s *SubIssuesSummary) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (s *Subscription) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (s *Subscription) GetIgnored() bool { _ = "STUB: not implemented"; return false }

func (s *Subscription) GetReason() string { _ = "STUB: not implemented"; return "" }

func (s *Subscription) GetRepositoryURL() string { _ = "STUB: not implemented"; return "" }

func (s *Subscription) GetSubscribed() bool { _ = "STUB: not implemented"; return false }

func (s *Subscription) GetThreadURL() string { _ = "STUB: not implemented"; return "" }

func (s *Subscription) GetURL() string { _ = "STUB: not implemented"; return "" }

func (s *SystemRequirements) GetNodes() []*SystemRequirementsNode {
	_ = "STUB: not implemented"
	return nil
}

func (s *SystemRequirements) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (s *SystemRequirementsNode) GetHostname() string { _ = "STUB: not implemented"; return "" }

func (s *SystemRequirementsNode) GetRolesStatus() []*SystemRequirementsNodeRoleStatus {
	_ = "STUB: not implemented"
	return nil
}

func (s *SystemRequirementsNode) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (s *SystemRequirementsNodeRoleStatus) GetRole() string { _ = "STUB: not implemented"; return "" }

func (s *SystemRequirementsNodeRoleStatus) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (t *Tag) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (t *Tag) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (t *Tag) GetObject() *GitObject { _ = "STUB: not implemented"; return nil }

func (t *Tag) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (t *Tag) GetTag() string { _ = "STUB: not implemented"; return "" }

func (t *Tag) GetTagger() *CommitAuthor { _ = "STUB: not implemented"; return nil }

func (t *Tag) GetURL() string { _ = "STUB: not implemented"; return "" }

func (t *Tag) GetVerification() *SignatureVerification { _ = "STUB: not implemented"; return nil }

func (t *TagProtection) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (t *TagProtection) GetPattern() string { _ = "STUB: not implemented"; return "" }

func (t *TaskStep) GetCompletedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (t *TaskStep) GetConclusion() string { _ = "STUB: not implemented"; return "" }

func (t *TaskStep) GetName() string { _ = "STUB: not implemented"; return "" }

func (t *TaskStep) GetNumber() int64 { _ = "STUB: not implemented"; return 0 }

func (t *TaskStep) GetStartedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (t *TaskStep) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetAccessSource() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetAssignment() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (t *Team) GetLDAPDN() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetMembersCount() int { _ = "STUB: not implemented"; return 0 }

func (t *Team) GetMembersURL() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetName() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetNotificationSetting() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetOrganization() *Organization { _ = "STUB: not implemented"; return nil }

func (t *Team) GetParent() *Team { _ = "STUB: not implemented"; return nil }

func (t *Team) GetPermission() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetPermissions() map[string]bool { _ = "STUB: not implemented"; return nil }

func (t *Team) GetPrivacy() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetReposCount() int { _ = "STUB: not implemented"; return 0 }

func (t *Team) GetRepositoriesURL() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetSlug() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetType() string { _ = "STUB: not implemented"; return "" }

func (t *Team) GetURL() string { _ = "STUB: not implemented"; return "" }

func (t *TeamAddEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (t *TeamAddEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (t *TeamAddEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (t *TeamAddEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (t *TeamAddEvent) GetTeam() *Team { _ = "STUB: not implemented"; return nil }

func (t *TeamAddTeamMembershipOptions) GetRole() string { _ = "STUB: not implemented"; return "" }

func (t *TeamAddTeamRepoOptions) GetPermission() string { _ = "STUB: not implemented"; return "" }

func (t *TeamChange) GetDescription() *TeamDescription { _ = "STUB: not implemented"; return nil }

func (t *TeamChange) GetName() *TeamName { _ = "STUB: not implemented"; return nil }

func (t *TeamChange) GetPrivacy() *TeamPrivacy { _ = "STUB: not implemented"; return nil }

func (t *TeamChange) GetRepository() *TeamRepository { _ = "STUB: not implemented"; return nil }

func (t *TeamDescription) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (t *TeamDiscussion) GetAuthor() *User { _ = "STUB: not implemented"; return nil }

func (t *TeamDiscussion) GetBody() string { _ = "STUB: not implemented"; return "" }

func (t *TeamDiscussion) GetBodyHTML() string { _ = "STUB: not implemented"; return "" }

func (t *TeamDiscussion) GetBodyVersion() string { _ = "STUB: not implemented"; return "" }

func (t *TeamDiscussion) GetCommentsCount() int { _ = "STUB: not implemented"; return 0 }

func (t *TeamDiscussion) GetCommentsURL() string { _ = "STUB: not implemented"; return "" }

func (t *TeamDiscussion) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (t *TeamDiscussion) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (t *TeamDiscussion) GetLastEditedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (t *TeamDiscussion) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (t *TeamDiscussion) GetNumber() int { _ = "STUB: not implemented"; return 0 }

func (t *TeamDiscussion) GetPinned() bool { _ = "STUB: not implemented"; return false }

func (t *TeamDiscussion) GetPrivate() bool { _ = "STUB: not implemented"; return false }

func (t *TeamDiscussion) GetReactions() *Reactions { _ = "STUB: not implemented"; return nil }

func (t *TeamDiscussion) GetTeamURL() string { _ = "STUB: not implemented"; return "" }

func (t *TeamDiscussion) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (t *TeamDiscussion) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (t *TeamDiscussion) GetURL() string { _ = "STUB: not implemented"; return "" }

func (t *TeamEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (t *TeamEvent) GetChanges() *TeamChange { _ = "STUB: not implemented"; return nil }

func (t *TeamEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (t *TeamEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (t *TeamEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (t *TeamEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (t *TeamEvent) GetTeam() *Team { _ = "STUB: not implemented"; return nil }

func (t *TeamLDAPMapping) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (t *TeamLDAPMapping) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (t *TeamLDAPMapping) GetLDAPDN() string { _ = "STUB: not implemented"; return "" }

func (t *TeamLDAPMapping) GetMembersURL() string { _ = "STUB: not implemented"; return "" }

func (t *TeamLDAPMapping) GetName() string { _ = "STUB: not implemented"; return "" }

func (t *TeamLDAPMapping) GetPermission() string { _ = "STUB: not implemented"; return "" }

func (t *TeamLDAPMapping) GetPrivacy() string { _ = "STUB: not implemented"; return "" }

func (t *TeamLDAPMapping) GetRepositoriesURL() string { _ = "STUB: not implemented"; return "" }

func (t *TeamLDAPMapping) GetSlug() string { _ = "STUB: not implemented"; return "" }

func (t *TeamLDAPMapping) GetURL() string { _ = "STUB: not implemented"; return "" }

func (t *TeamListTeamMembersOptions) GetRole() string { _ = "STUB: not implemented"; return "" }

func (t *TeamName) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (t *TeamPermissions) GetFrom() *TeamPermissionsFrom { _ = "STUB: not implemented"; return nil }

func (t *TeamPermissionsFrom) GetAdmin() bool { _ = "STUB: not implemented"; return false }

func (t *TeamPermissionsFrom) GetPull() bool { _ = "STUB: not implemented"; return false }

func (t *TeamPermissionsFrom) GetPush() bool { _ = "STUB: not implemented"; return false }

func (t *TeamPrivacy) GetFrom() string { _ = "STUB: not implemented"; return "" }

func (t *TeamProjectOptions) GetPermission() string { _ = "STUB: not implemented"; return "" }

func (t *TeamRepository) GetPermissions() *TeamPermissions { _ = "STUB: not implemented"; return nil }

func (t *TemplateRepoRequest) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (t *TemplateRepoRequest) GetIncludeAllBranches() bool { _ = "STUB: not implemented"; return false }

func (t *TemplateRepoRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (t *TemplateRepoRequest) GetOwner() string { _ = "STUB: not implemented"; return "" }

func (t *TemplateRepoRequest) GetPrivate() bool { _ = "STUB: not implemented"; return false }

func (t *TextMatch) GetFragment() string { _ = "STUB: not implemented"; return "" }

func (t *TextMatch) GetMatches() []*Match { _ = "STUB: not implemented"; return nil }

func (t *TextMatch) GetObjectType() string { _ = "STUB: not implemented"; return "" }

func (t *TextMatch) GetObjectURL() string { _ = "STUB: not implemented"; return "" }

func (t *TextMatch) GetProperty() string { _ = "STUB: not implemented"; return "" }

func (t *Timeline) GetActor() *User { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetAssignee() *User { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetAssigner() *User { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetAuthor() *CommitAuthor { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetBody() string { _ = "STUB: not implemented"; return "" }

func (t *Timeline) GetCommitID() string { _ = "STUB: not implemented"; return "" }

func (t *Timeline) GetCommitter() *CommitAuthor { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetCommitURL() string { _ = "STUB: not implemented"; return "" }

func (t *Timeline) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (t *Timeline) GetEvent() string { _ = "STUB: not implemented"; return "" }

func (t *Timeline) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (t *Timeline) GetLabel() *Label { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (t *Timeline) GetMilestone() *Milestone { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetParents() []*Commit { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetPerformedViaGithubApp() *App { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetRename() *Rename { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetRequestedTeam() *Team { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetRequester() *User { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetReviewer() *User { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (t *Timeline) GetSource() *Source { _ = "STUB: not implemented"; return nil }

func (t *Timeline) GetState() string { _ = "STUB: not implemented"; return "" }

func (t *Timeline) GetSubmittedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (t *Timeline) GetURL() string { _ = "STUB: not implemented"; return "" }

func (t *Timeline) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (t *Tool) GetGUID() string { _ = "STUB: not implemented"; return "" }

func (t *Tool) GetName() string { _ = "STUB: not implemented"; return "" }

func (t *Tool) GetVersion() string { _ = "STUB: not implemented"; return "" }

func (t *TopicResult) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (t *TopicResult) GetCreatedBy() string { _ = "STUB: not implemented"; return "" }

func (t *TopicResult) GetCurated() bool { _ = "STUB: not implemented"; return false }

func (t *TopicResult) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (t *TopicResult) GetDisplayName() string { _ = "STUB: not implemented"; return "" }

func (t *TopicResult) GetFeatured() bool { _ = "STUB: not implemented"; return false }

func (t *TopicResult) GetName() string { _ = "STUB: not implemented"; return "" }

func (t *TopicResult) GetScore() float64 { _ = "STUB: not implemented"; return 0 }

func (t *TopicResult) GetShortDescription() string { _ = "STUB: not implemented"; return "" }

func (t *TopicResult) GetUpdatedAt() string { _ = "STUB: not implemented"; return "" }

func (t *TopicsSearchResult) GetIncompleteResults() bool { _ = "STUB: not implemented"; return false }

func (t *TopicsSearchResult) GetTopics() []*TopicResult { _ = "STUB: not implemented"; return nil }

func (t *TopicsSearchResult) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (t *TotalCacheUsage) GetTotalActiveCachesCount() int { _ = "STUB: not implemented"; return 0 }

func (t *TotalCacheUsage) GetTotalActiveCachesUsageSizeInBytes() int64 {
	_ = "STUB: not implemented"
	return 0
}

func (t *TrafficBreakdownOptions) GetPer() string { _ = "STUB: not implemented"; return "" }

func (t *TrafficClones) GetClones() []*TrafficData { _ = "STUB: not implemented"; return nil }

func (t *TrafficClones) GetCount() int { _ = "STUB: not implemented"; return 0 }

func (t *TrafficClones) GetUniques() int { _ = "STUB: not implemented"; return 0 }

func (t *TrafficData) GetCount() int { _ = "STUB: not implemented"; return 0 }

func (t *TrafficData) GetTimestamp() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (t *TrafficData) GetUniques() int { _ = "STUB: not implemented"; return 0 }

func (t *TrafficPath) GetCount() int { _ = "STUB: not implemented"; return 0 }

func (t *TrafficPath) GetPath() string { _ = "STUB: not implemented"; return "" }

func (t *TrafficPath) GetTitle() string { _ = "STUB: not implemented"; return "" }

func (t *TrafficPath) GetUniques() int { _ = "STUB: not implemented"; return 0 }

func (t *TrafficReferrer) GetCount() int { _ = "STUB: not implemented"; return 0 }

func (t *TrafficReferrer) GetReferrer() string { _ = "STUB: not implemented"; return "" }

func (t *TrafficReferrer) GetUniques() int { _ = "STUB: not implemented"; return 0 }

func (t *TrafficViews) GetCount() int { _ = "STUB: not implemented"; return 0 }

func (t *TrafficViews) GetUniques() int { _ = "STUB: not implemented"; return 0 }

func (t *TrafficViews) GetViews() []*TrafficData { _ = "STUB: not implemented"; return nil }

func (t *TransferRequest) GetNewName() string { _ = "STUB: not implemented"; return "" }

func (t *TransferRequest) GetNewOwner() string { _ = "STUB: not implemented"; return "" }

func (t *TransferRequest) GetTeamID() []int64 { _ = "STUB: not implemented"; return nil }

func (t *Tree) GetEntries() []*TreeEntry { _ = "STUB: not implemented"; return nil }

func (t *Tree) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (t *Tree) GetTruncated() bool { _ = "STUB: not implemented"; return false }

func (t *TreeEntry) GetContent() string { _ = "STUB: not implemented"; return "" }

func (t *TreeEntry) GetMode() string { _ = "STUB: not implemented"; return "" }

func (t *TreeEntry) GetPath() string { _ = "STUB: not implemented"; return "" }

func (t *TreeEntry) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (t *TreeEntry) GetSize() int { _ = "STUB: not implemented"; return 0 }

func (t *TreeEntry) GetType() string { _ = "STUB: not implemented"; return "" }

func (t *TreeEntry) GetURL() string { _ = "STUB: not implemented"; return "" }

func (u *UnauthenticatedRateLimitedTransport) GetClientID() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UnauthenticatedRateLimitedTransport) GetClientSecret() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateAppInstallationRepositoriesRequest) GetRepositories() []string {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateAppInstallationRepositoriesRequest) GetRepositorySelection() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateAttributeForSCIMUserOperations) GetOp() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateAttributeForSCIMUserOperations) GetPath() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateAttributeForSCIMUserOperations) GetValue() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (u *UpdateAttributeForSCIMUserRequest) GetOperations() []*UpdateAttributeForSCIMUserOperations {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateAttributeForSCIMUserRequest) GetSchemas() []string {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateBranchRule) GetParameters() UpdateRuleParameters {
	_ = "STUB: not implemented"
	return *new(UpdateRuleParameters)
}

func (u *UpdateCheckRunOptions) GetActions() []*CheckRunAction {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateCheckRunOptions) GetCompletedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (u *UpdateCheckRunOptions) GetConclusion() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateCheckRunOptions) GetDetailsURL() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateCheckRunOptions) GetExternalID() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateCheckRunOptions) GetName() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateCheckRunOptions) GetOutput() *CheckRunOutput { _ = "STUB: not implemented"; return nil }

func (u *UpdateCheckRunOptions) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateCodespaceOptions) GetMachine() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateCodespaceOptions) GetRecentFolders() []string { _ = "STUB: not implemented"; return nil }

func (u *UpdateCustomOrgRoleRequest) GetBaseRole() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateCustomOrgRoleRequest) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateCustomOrgRoleRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateCustomOrgRoleRequest) GetPermissions() []string {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateDefaultSetupConfigurationOptions) GetLanguages() []string {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateDefaultSetupConfigurationOptions) GetQuerySuite() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateDefaultSetupConfigurationOptions) GetState() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateDefaultSetupConfigurationResponse) GetRunID() int64 {
	_ = "STUB: not implemented"
	return 0
}

func (u *UpdateDefaultSetupConfigurationResponse) GetRunURL() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateDeploymentBranchPolicyRequest) GetName() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateEnterpriseRunnerGroupRequest) GetAllowsPublicRepositories() bool {
	_ = "STUB: not implemented"
	return false
}

func (u *UpdateEnterpriseRunnerGroupRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateEnterpriseRunnerGroupRequest) GetNetworkConfigurationID() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateEnterpriseRunnerGroupRequest) GetRestrictedToWorkflows() bool {
	_ = "STUB: not implemented"
	return false
}

func (u *UpdateEnterpriseRunnerGroupRequest) GetSelectedWorkflows() []string {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateEnterpriseRunnerGroupRequest) GetVisibility() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateGistCommentRequest) GetBody() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateGistFile) GetContent() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateGistFile) GetFilename() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateGistRequest) GetDescription() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateHostedRunnerRequest) GetEnableStaticIP() bool {
	_ = "STUB: not implemented"
	return false
}

func (u *UpdateHostedRunnerRequest) GetImageID() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateHostedRunnerRequest) GetImageVersion() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateHostedRunnerRequest) GetMaximumRunners() int64 { _ = "STUB: not implemented"; return 0 }

func (u *UpdateHostedRunnerRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateHostedRunnerRequest) GetRunnerGroupID() int64 { _ = "STUB: not implemented"; return 0 }

func (u *UpdateHostedRunnerRequest) GetSize() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateOrganizationPrivateRegistry) GetAccountID() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetAudience() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetAuthType() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetAWSRegion() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetClientID() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetDomain() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetDomainOwner() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetEncryptedValue() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetIdentityMappingName() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetJFrogOIDCProviderName() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetKeyID() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateOrganizationPrivateRegistry) GetRegistryType() *PrivateRegistryType {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateOrganizationPrivateRegistry) GetReplacesBase() bool {
	_ = "STUB: not implemented"
	return false
}

func (u *UpdateOrganizationPrivateRegistry) GetRoleName() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetSelectedRepositoryIDs() []int64 {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateOrganizationPrivateRegistry) GetTenantID() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetURL() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateOrganizationPrivateRegistry) GetUsername() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateOrganizationPrivateRegistry) GetVisibility() *PrivateRegistryVisibility {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateProjectItemOptions) GetArchived() bool { _ = "STUB: not implemented"; return false }

func (u *UpdateProjectItemOptions) GetFields() []*UpdateProjectV2Field {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateProjectV2Field) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (u *UpdateProjectV2Field) GetValue() any { _ = "STUB: not implemented"; return *new(any) }

func (u *UpdateProvisionedOrgMembershipRequest) GetActive() bool {
	_ = "STUB: not implemented"
	return false
}

func (u *UpdateProvisionedOrgMembershipRequest) GetDisplayName() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateProvisionedOrgMembershipRequest) GetEmails() []*SCIMUserEmail {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateProvisionedOrgMembershipRequest) GetExternalID() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateProvisionedOrgMembershipRequest) GetGroups() []string {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateProvisionedOrgMembershipRequest) GetName() SCIMUserName {
	_ = "STUB: not implemented"
	return *new(SCIMUserName)
}

func (u *UpdateProvisionedOrgMembershipRequest) GetSchemas() []string {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateProvisionedOrgMembershipRequest) GetUserName() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateRef) GetForce() bool { _ = "STUB: not implemented"; return false }

func (u *UpdateRef) GetSHA() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateReleaseAssetRequest) GetLabel() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateReleaseAssetRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateReleaseAssetRequest) GetState() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateReleaseRequest) GetBody() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateReleaseRequest) GetDiscussionCategoryName() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateReleaseRequest) GetDraft() bool { _ = "STUB: not implemented"; return false }

func (u *UpdateReleaseRequest) GetMakeLatest() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateReleaseRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateReleaseRequest) GetPrerelease() bool { _ = "STUB: not implemented"; return false }

func (u *UpdateReleaseRequest) GetTagName() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateReleaseRequest) GetTargetCommitish() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateRuleParameters) GetUpdateAllowsFetchAndMerge() bool {
	_ = "STUB: not implemented"
	return false
}

func (u *UpdateRunnerGroupRequest) GetAllowsPublicRepositories() bool {
	_ = "STUB: not implemented"
	return false
}

func (u *UpdateRunnerGroupRequest) GetName() string { _ = "STUB: not implemented"; return "" }

func (u *UpdateRunnerGroupRequest) GetNetworkConfigurationID() string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpdateRunnerGroupRequest) GetRestrictedToWorkflows() bool {
	_ = "STUB: not implemented"
	return false
}

func (u *UpdateRunnerGroupRequest) GetSelectedWorkflows() []string {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpdateRunnerGroupRequest) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (u *UploadLicenseOptions) GetLicense() string { _ = "STUB: not implemented"; return "" }

func (u *UploadOptions) GetLabel() string { _ = "STUB: not implemented"; return "" }

func (u *UploadOptions) GetMediaType() string { _ = "STUB: not implemented"; return "" }

func (u *UploadOptions) GetName() string { _ = "STUB: not implemented"; return "" }

func (u *UsageItem) GetDate() string { _ = "STUB: not implemented"; return "" }

func (u *UsageItem) GetDiscountAmount() float64 { _ = "STUB: not implemented"; return 0 }

func (u *UsageItem) GetGrossAmount() float64 { _ = "STUB: not implemented"; return 0 }

func (u *UsageItem) GetNetAmount() float64 { _ = "STUB: not implemented"; return 0 }

func (u *UsageItem) GetOrganizationName() string { _ = "STUB: not implemented"; return "" }

func (u *UsageItem) GetPricePerUnit() float64 { _ = "STUB: not implemented"; return 0 }

func (u *UsageItem) GetProduct() string { _ = "STUB: not implemented"; return "" }

func (u *UsageItem) GetQuantity() float64 { _ = "STUB: not implemented"; return 0 }

func (u *UsageItem) GetRepositoryName() string { _ = "STUB: not implemented"; return "" }

func (u *UsageItem) GetSKU() string { _ = "STUB: not implemented"; return "" }

func (u *UsageItem) GetUnitType() string { _ = "STUB: not implemented"; return "" }

func (u *UsageReport) GetUsageItems() []*UsageItem { _ = "STUB: not implemented"; return nil }

func (u *UsageReportOptions) GetDay() int { _ = "STUB: not implemented"; return 0 }

func (u *UsageReportOptions) GetHour() int { _ = "STUB: not implemented"; return 0 }

func (u *UsageReportOptions) GetMonth() int { _ = "STUB: not implemented"; return 0 }

func (u *UsageReportOptions) GetYear() int { _ = "STUB: not implemented"; return 0 }

func (u *User) GetAssignment() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetAvatarURL() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetBio() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetBlog() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetBusinessPlus() bool { _ = "STUB: not implemented"; return false }

func (u *User) GetCollaborators() int { _ = "STUB: not implemented"; return 0 }

func (u *User) GetCompany() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (u *User) GetDiskUsage() int { _ = "STUB: not implemented"; return 0 }

func (u *User) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetEventsURL() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetFollowers() int { _ = "STUB: not implemented"; return 0 }

func (u *User) GetFollowersURL() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetFollowing() int { _ = "STUB: not implemented"; return 0 }

func (u *User) GetFollowingURL() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetGistsURL() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetGravatarID() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetHireable() bool { _ = "STUB: not implemented"; return false }

func (u *User) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (u *User) GetInherited() bool { _ = "STUB: not implemented"; return false }

func (u *User) GetInheritedFrom() []*Team { _ = "STUB: not implemented"; return nil }

func (u *User) GetLdapDn() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetLocation() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetName() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetNotificationEmail() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetOrganizationsURL() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetOwnedPrivateRepos() int64 { _ = "STUB: not implemented"; return 0 }

func (u *User) GetPermissions() *RepositoryPermissions { _ = "STUB: not implemented"; return nil }

func (u *User) GetPlan() *Plan { _ = "STUB: not implemented"; return nil }

func (u *User) GetPrivateGists() int { _ = "STUB: not implemented"; return 0 }

func (u *User) GetPublicGists() int { _ = "STUB: not implemented"; return 0 }

func (u *User) GetPublicRepos() int { _ = "STUB: not implemented"; return 0 }

func (u *User) GetReceivedEventsURL() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetReposURL() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetRole() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetRoleName() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetSiteAdmin() bool { _ = "STUB: not implemented"; return false }

func (u *User) GetStarredURL() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetSubscriptionsURL() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetSuspendedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (u *User) GetTextMatches() []*TextMatch { _ = "STUB: not implemented"; return nil }

func (u *User) GetTotalPrivateRepos() int64 { _ = "STUB: not implemented"; return 0 }

func (u *User) GetTwitterUsername() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetTwoFactorAuthentication() bool { _ = "STUB: not implemented"; return false }

func (u *User) GetType() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (u *User) GetURL() string { _ = "STUB: not implemented"; return "" }

func (u *User) GetUserViewType() string { _ = "STUB: not implemented"; return "" }

func (u *UserAuthorization) GetApp() *OAuthAPP { _ = "STUB: not implemented"; return nil }

func (u *UserAuthorization) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (u *UserAuthorization) GetFingerprint() string { _ = "STUB: not implemented"; return "" }

func (u *UserAuthorization) GetHashedToken() string { _ = "STUB: not implemented"; return "" }

func (u *UserAuthorization) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (u *UserAuthorization) GetNote() string { _ = "STUB: not implemented"; return "" }

func (u *UserAuthorization) GetNoteURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserAuthorization) GetScopes() []string { _ = "STUB: not implemented"; return nil }

func (u *UserAuthorization) GetToken() string { _ = "STUB: not implemented"; return "" }

func (u *UserAuthorization) GetTokenLastEight() string { _ = "STUB: not implemented"; return "" }

func (u *UserAuthorization) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (u *UserAuthorization) GetURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserContext) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (u *UserContext) GetOcticon() string { _ = "STUB: not implemented"; return "" }

func (u *UserEmail) GetEmail() string { _ = "STUB: not implemented"; return "" }

func (u *UserEmail) GetPrimary() bool { _ = "STUB: not implemented"; return false }

func (u *UserEmail) GetVerified() bool { _ = "STUB: not implemented"; return false }

func (u *UserEmail) GetVisibility() string { _ = "STUB: not implemented"; return "" }

func (u *UserEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (u *UserEvent) GetEnterprise() *Enterprise { _ = "STUB: not implemented"; return nil }

func (u *UserEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (u *UserEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (u *UserEvent) GetUser() *User { _ = "STUB: not implemented"; return nil }

func (u *UserLDAPMapping) GetAvatarURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetEventsURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetFollowersURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetFollowingURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetGistsURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetGravatarID() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (u *UserLDAPMapping) GetLDAPDN() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetLogin() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetOrganizationsURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetReceivedEventsURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetReposURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetSiteAdmin() bool { _ = "STUB: not implemented"; return false }

func (u *UserLDAPMapping) GetStarredURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetSubscriptionsURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetType() string { _ = "STUB: not implemented"; return "" }

func (u *UserLDAPMapping) GetURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserListOptions) GetPerPage() int { _ = "STUB: not implemented"; return 0 }

func (u *UserListOptions) GetSince() int64 { _ = "STUB: not implemented"; return 0 }

func (u *UserMigration) GetCreatedAt() string { _ = "STUB: not implemented"; return "" }

func (u *UserMigration) GetExcludeAttachments() bool { _ = "STUB: not implemented"; return false }

func (u *UserMigration) GetGUID() string { _ = "STUB: not implemented"; return "" }

func (u *UserMigration) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (u *UserMigration) GetLockRepositories() bool { _ = "STUB: not implemented"; return false }

func (u *UserMigration) GetRepositories() []*Repository { _ = "STUB: not implemented"; return nil }

func (u *UserMigration) GetState() string { _ = "STUB: not implemented"; return "" }

func (u *UserMigration) GetUpdatedAt() string { _ = "STUB: not implemented"; return "" }

func (u *UserMigration) GetURL() string { _ = "STUB: not implemented"; return "" }

func (u *UserMigrationOptions) GetExcludeAttachments() bool {
	_ = "STUB: not implemented"
	return false
}

func (u *UserMigrationOptions) GetLockRepositories() bool { _ = "STUB: not implemented"; return false }

func (u *UsersSearchResult) GetIncompleteResults() bool { _ = "STUB: not implemented"; return false }

func (u *UsersSearchResult) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (u *UsersSearchResult) GetUsers() []*User { _ = "STUB: not implemented"; return nil }

func (u *UserStats) GetAdminUsers() int { _ = "STUB: not implemented"; return 0 }

func (u *UserStats) GetSuspendedUsers() int { _ = "STUB: not implemented"; return 0 }

func (u *UserStats) GetTotalUsers() int { _ = "STUB: not implemented"; return 0 }

func (u *UserSuspendOptions) GetReason() string { _ = "STUB: not implemented"; return "" }

func (v *VulnerabilityPackage) GetEcosystem() string { _ = "STUB: not implemented"; return "" }

func (v *VulnerabilityPackage) GetName() string { _ = "STUB: not implemented"; return "" }

func (w *WatchEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (w *WatchEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (w *WatchEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (w *WatchEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (w *WatchEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (w *WeeklyCommitActivity) GetDays() []int { _ = "STUB: not implemented"; return nil }

func (w *WeeklyCommitActivity) GetTotal() int { _ = "STUB: not implemented"; return 0 }

func (w *WeeklyCommitActivity) GetWeek() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (w *WeeklyStats) GetAdditions() int { _ = "STUB: not implemented"; return 0 }

func (w *WeeklyStats) GetCommits() int { _ = "STUB: not implemented"; return 0 }

func (w *WeeklyStats) GetDeletions() int { _ = "STUB: not implemented"; return 0 }

func (w *WeeklyStats) GetWeek() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (w *Workflow) GetBadgeURL() string { _ = "STUB: not implemented"; return "" }

func (w *Workflow) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (w *Workflow) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (w *Workflow) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (w *Workflow) GetName() string { _ = "STUB: not implemented"; return "" }

func (w *Workflow) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (w *Workflow) GetPath() string { _ = "STUB: not implemented"; return "" }

func (w *Workflow) GetState() string { _ = "STUB: not implemented"; return "" }

func (w *Workflow) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (w *Workflow) GetURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowBill) GetTotalMS() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowDispatchEvent) GetInputs() json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func (w *WorkflowDispatchEvent) GetInstallation() *Installation {
	_ = "STUB: not implemented"
	return nil
}

func (w *WorkflowDispatchEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (w *WorkflowDispatchEvent) GetRef() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowDispatchEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (w *WorkflowDispatchEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (w *WorkflowDispatchEvent) GetWorkflow() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowDispatchRunDetails) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowDispatchRunDetails) GetRunURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowDispatchRunDetails) GetWorkflowRunID() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowJob) GetCheckRunURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJob) GetCompletedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (w *WorkflowJob) GetConclusion() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJob) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (w *WorkflowJob) GetHeadBranch() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJob) GetHeadSHA() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJob) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJob) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowJob) GetLabels() []string { _ = "STUB: not implemented"; return nil }

func (w *WorkflowJob) GetName() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJob) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJob) GetRunAttempt() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowJob) GetRunID() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowJob) GetRunnerGroupID() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowJob) GetRunnerGroupName() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJob) GetRunnerID() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowJob) GetRunnerName() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJob) GetRunURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJob) GetStartedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (w *WorkflowJob) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJob) GetSteps() []*TaskStep { _ = "STUB: not implemented"; return nil }

func (w *WorkflowJob) GetURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJob) GetWorkflowName() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJobEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJobEvent) GetDeployment() *Deployment { _ = "STUB: not implemented"; return nil }

func (w *WorkflowJobEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (w *WorkflowJobEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (w *WorkflowJobEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (w *WorkflowJobEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (w *WorkflowJobEvent) GetWorkflowJob() *WorkflowJob { _ = "STUB: not implemented"; return nil }

func (w *WorkflowJobRun) GetConclusion() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJobRun) GetCreatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (w *WorkflowJobRun) GetEnvironment() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJobRun) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJobRun) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowJobRun) GetName() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJobRun) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowJobRun) GetUpdatedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (w *WorkflowRun) GetActor() *User { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRun) GetArtifactsURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetCancelURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetCheckSuiteID() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowRun) GetCheckSuiteNodeID() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetCheckSuiteURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetConclusion() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetCreatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (w *WorkflowRun) GetDisplayTitle() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetEvent() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetHeadBranch() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetHeadCommit() *HeadCommit { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRun) GetHeadRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRun) GetHeadSHA() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetHTMLURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetID() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowRun) GetJobsURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetLogsURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetName() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetNodeID() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetPath() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetPreviousAttemptURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetPullRequests() []*PullRequest { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRun) GetReferencedWorkflows() []*ReferencedWorkflow {
	_ = "STUB: not implemented"
	return nil
}

func (w *WorkflowRun) GetRepository() *Repository { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRun) GetRerunURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetRunAttempt() int { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowRun) GetRunNumber() int { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowRun) GetRunStartedAt() Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

func (w *WorkflowRun) GetStatus() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetTriggeringActor() *User { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRun) GetUpdatedAt() Timestamp { _ = "STUB: not implemented"; return *new(Timestamp) }

func (w *WorkflowRun) GetURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRun) GetWorkflowID() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowRun) GetWorkflowURL() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRunAttemptOptions) GetExcludePullRequests() bool {
	_ = "STUB: not implemented"
	return false
}

func (w *WorkflowRunBill) GetJobRuns() []*WorkflowRunJobRun { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRunBill) GetJobs() int { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowRunBill) GetTotalMS() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowRunEvent) GetAction() string { _ = "STUB: not implemented"; return "" }

func (w *WorkflowRunEvent) GetInstallation() *Installation { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRunEvent) GetOrg() *Organization { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRunEvent) GetRepo() *Repository { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRunEvent) GetSender() *User { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRunEvent) GetWorkflow() *Workflow { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRunEvent) GetWorkflowRun() *WorkflowRun { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRunJobRun) GetDurationMS() int64 { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowRunJobRun) GetJobID() int { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowRuns) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (w *WorkflowRuns) GetWorkflowRuns() []*WorkflowRun { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRunUsage) GetBillable() *WorkflowRunBillMap { _ = "STUB: not implemented"; return nil }

func (w *WorkflowRunUsage) GetRunDurationMS() int64 { _ = "STUB: not implemented"; return 0 }

func (w *Workflows) GetTotalCount() int { _ = "STUB: not implemented"; return 0 }

func (w *Workflows) GetWorkflows() []*Workflow { _ = "STUB: not implemented"; return nil }

func (w *WorkflowsBranchRule) GetParameters() WorkflowsRuleParameters {
	_ = "STUB: not implemented"
	return *new(WorkflowsRuleParameters)
}

func (w *WorkflowsPermissions) GetRequireApprovalForForkPRWorkflows() bool {
	_ = "STUB: not implemented"
	return false
}

func (w *WorkflowsPermissions) GetRunWorkflowsFromForkPullRequests() bool {
	_ = "STUB: not implemented"
	return false
}

func (w *WorkflowsPermissions) GetSendSecretsAndVariables() bool {
	_ = "STUB: not implemented"
	return false
}

func (w *WorkflowsPermissions) GetSendWriteTokensToWorkflows() bool {
	_ = "STUB: not implemented"
	return false
}

func (w *WorkflowsPermissionsOpt) GetRequireApprovalForForkPRWorkflows() bool {
	_ = "STUB: not implemented"
	return false
}

func (w *WorkflowsPermissionsOpt) GetRunWorkflowsFromForkPullRequests() bool {
	_ = "STUB: not implemented"
	return false
}

func (w *WorkflowsPermissionsOpt) GetSendSecretsAndVariables() bool {
	_ = "STUB: not implemented"
	return false
}

func (w *WorkflowsPermissionsOpt) GetSendWriteTokensToWorkflows() bool {
	_ = "STUB: not implemented"
	return false
}

func (w *WorkflowsRuleParameters) GetDoNotEnforceOnCreate() bool {
	_ = "STUB: not implemented"
	return false
}

func (w *WorkflowsRuleParameters) GetWorkflows() []*RuleWorkflow {
	_ = "STUB: not implemented"
	return nil
}

func (w *WorkflowUsage) GetBillable() *WorkflowBillMap { _ = "STUB: not implemented"; return nil }
