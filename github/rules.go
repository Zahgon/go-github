package github

import (
	"encoding/json"
)

type RulesetTarget string

const (
	RulesetTargetBranch     RulesetTarget = "branch"
	RulesetTargetTag        RulesetTarget = "tag"
	RulesetTargetPush       RulesetTarget = "push"
	RulesetTargetRepository RulesetTarget = "repository"
)

type RulesetSourceType string

const (
	RulesetSourceTypeRepository   RulesetSourceType = "Repository"
	RulesetSourceTypeOrganization RulesetSourceType = "Organization"
	RulesetSourceTypeEnterprise   RulesetSourceType = "Enterprise"
)

type RulesetEnforcement string

const (
	RulesetEnforcementDisabled RulesetEnforcement = "disabled"
	RulesetEnforcementActive   RulesetEnforcement = "active"
	RulesetEnforcementEvaluate RulesetEnforcement = "evaluate"
)

type BypassActorType string

const (
	BypassActorTypeIntegration       BypassActorType = "Integration"
	BypassActorTypeOrganizationAdmin BypassActorType = "OrganizationAdmin"
	BypassActorTypeRepositoryRole    BypassActorType = "RepositoryRole"
	BypassActorTypeTeam              BypassActorType = "Team"
	BypassActorTypeDeployKey         BypassActorType = "DeployKey"
)

type BypassMode string

const (
	BypassModeAlways      BypassMode = "always"
	BypassModeExempt      BypassMode = "exempt"
	BypassModeNever       BypassMode = "never"
	BypassModePullRequest BypassMode = "pull_request"
)

type RepositoryRuleType string

const (
	RulesetRuleTypeBranchNamePattern        RepositoryRuleType = "branch_name_pattern"
	RulesetRuleTypeCodeScanning             RepositoryRuleType = "code_scanning"
	RulesetRuleTypeCommitAuthorEmailPattern RepositoryRuleType = "commit_author_email_pattern"
	RulesetRuleTypeCommitMessagePattern     RepositoryRuleType = "commit_message_pattern"
	RulesetRuleTypeCommitterEmailPattern    RepositoryRuleType = "committer_email_pattern"
	RulesetRuleTypeCopilotCodeReview        RepositoryRuleType = "copilot_code_review"
	RulesetRuleTypeCreation                 RepositoryRuleType = "creation"
	RulesetRuleTypeDeletion                 RepositoryRuleType = "deletion"
	RulesetRuleTypeMergeQueue               RepositoryRuleType = "merge_queue"
	RulesetRuleTypeNonFastForward           RepositoryRuleType = "non_fast_forward"
	RulesetRuleTypePullRequest              RepositoryRuleType = "pull_request"
	RulesetRuleTypeRequiredDeployments      RepositoryRuleType = "required_deployments"
	RulesetRuleTypeRequiredLinearHistory    RepositoryRuleType = "required_linear_history"
	RulesetRuleTypeRequiredSignatures       RepositoryRuleType = "required_signatures"
	RulesetRuleTypeRequiredStatusChecks     RepositoryRuleType = "required_status_checks"
	RulesetRuleTypeTagNamePattern           RepositoryRuleType = "tag_name_pattern"
	RulesetRuleTypeUpdate                   RepositoryRuleType = "update"
	RulesetRuleTypeWorkflows                RepositoryRuleType = "workflows"

	RulesetRuleTypeFileExtensionRestriction RepositoryRuleType = "file_extension_restriction"
	RulesetRuleTypeFilePathRestriction      RepositoryRuleType = "file_path_restriction"
	RulesetRuleTypeMaxFilePathLength        RepositoryRuleType = "max_file_path_length"
	RulesetRuleTypeMaxFileSize              RepositoryRuleType = "max_file_size"

	RulesetRuleTypeRepositoryCreate     RepositoryRuleType = "repository_create"
	RulesetRuleTypeRepositoryDelete     RepositoryRuleType = "repository_delete"
	RulesetRuleTypeRepositoryName       RepositoryRuleType = "repository_name"
	RulesetRuleTypeRepositoryTransfer   RepositoryRuleType = "repository_transfer"
	RulesetRuleTypeRepositoryVisibility RepositoryRuleType = "repository_visibility"
)

type MergeGroupingStrategy string

const (
	MergeGroupingStrategyAllGreen  MergeGroupingStrategy = "ALLGREEN"
	MergeGroupingStrategyHeadGreen MergeGroupingStrategy = "HEADGREEN"
)

type PullRequestMergeMethod string

const (
	PullRequestMergeMethodMerge  PullRequestMergeMethod = "merge"
	PullRequestMergeMethodRebase PullRequestMergeMethod = "rebase"
	PullRequestMergeMethodSquash PullRequestMergeMethod = "squash"
)

type MergeQueueMergeMethod string

const (
	MergeQueueMergeMethodMerge  MergeQueueMergeMethod = "MERGE"
	MergeQueueMergeMethodRebase MergeQueueMergeMethod = "REBASE"
	MergeQueueMergeMethodSquash MergeQueueMergeMethod = "SQUASH"
)

type RulesetReviewerType string

const (
	RulesetReviewerTypeTeam RulesetReviewerType = "Team"
)

type PatternRuleOperator string

const (
	PatternRuleOperatorStartsWith PatternRuleOperator = "starts_with"
	PatternRuleOperatorEndsWith   PatternRuleOperator = "ends_with"
	PatternRuleOperatorContains   PatternRuleOperator = "contains"
	PatternRuleOperatorRegex      PatternRuleOperator = "regex"
)

type CodeScanningAlertsThreshold string

const (
	CodeScanningAlertsThresholdNone              CodeScanningAlertsThreshold = "none"
	CodeScanningAlertsThresholdErrors            CodeScanningAlertsThreshold = "errors"
	CodeScanningAlertsThresholdErrorsAndWarnings CodeScanningAlertsThreshold = "errors_and_warnings"
	CodeScanningAlertsThresholdAll               CodeScanningAlertsThreshold = "all"
)

type CodeScanningSecurityAlertsThreshold string

const (
	CodeScanningSecurityAlertsThresholdNone           CodeScanningSecurityAlertsThreshold = "none"
	CodeScanningSecurityAlertsThresholdCritical       CodeScanningSecurityAlertsThreshold = "critical"
	CodeScanningSecurityAlertsThresholdHighOrHigher   CodeScanningSecurityAlertsThreshold = "high_or_higher"
	CodeScanningSecurityAlertsThresholdMediumOrHigher CodeScanningSecurityAlertsThreshold = "medium_or_higher"
	CodeScanningSecurityAlertsThresholdAll            CodeScanningSecurityAlertsThreshold = "all"
)

type RepositoryRuleset struct {
	ID                   *int64                       `json:"id,omitempty"`
	Name                 string                       `json:"name"`
	Target               *RulesetTarget               `json:"target,omitempty"`
	SourceType           *RulesetSourceType           `json:"source_type,omitempty"`
	Source               string                       `json:"source"`
	Enforcement          RulesetEnforcement           `json:"enforcement"`
	BypassActors         []*BypassActor               `json:"bypass_actors,omitzero"`
	CurrentUserCanBypass *BypassMode                  `json:"current_user_can_bypass,omitempty"`
	NodeID               *string                      `json:"node_id,omitempty"`
	Links                *RepositoryRulesetLinks      `json:"_links,omitempty"`
	Conditions           *RepositoryRulesetConditions `json:"conditions,omitempty"`
	Rules                *RepositoryRulesetRules      `json:"rules,omitempty"`
	UpdatedAt            *Timestamp                   `json:"updated_at,omitempty"`
	CreatedAt            *Timestamp                   `json:"created_at,omitempty"`
}

type BypassActor struct {
	ActorID    *int64           `json:"actor_id,omitempty"`
	ActorType  *BypassActorType `json:"actor_type,omitempty"`
	BypassMode *BypassMode      `json:"bypass_mode,omitempty"`
}

type RepositoryRulesetLinks struct {
	Self *RepositoryRulesetLink `json:"self,omitempty"`
	HTML *RepositoryRulesetLink `json:"html,omitempty"`
}

type RepositoryRulesetLink struct {
	HRef *string `json:"href,omitempty"`
}

type RepositoryRulesetConditions struct {
	RefName              *RepositoryRulesetRefConditionParameters                  `json:"ref_name,omitempty"`
	RepositoryID         *RepositoryRulesetRepositoryIDsConditionParameters        `json:"repository_id,omitempty"`
	RepositoryName       *RepositoryRulesetRepositoryNamesConditionParameters      `json:"repository_name,omitempty"`
	RepositoryProperty   *RepositoryRulesetRepositoryPropertyConditionParameters   `json:"repository_property,omitempty"`
	OrganizationID       *RepositoryRulesetOrganizationIDsConditionParameters      `json:"organization_id,omitempty"`
	OrganizationName     *RepositoryRulesetOrganizationNamesConditionParameters    `json:"organization_name,omitempty"`
	OrganizationProperty *RepositoryRulesetOrganizationPropertyConditionParameters `json:"organization_property,omitempty"`
}

type RepositoryRulesetRefConditionParameters struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

type RepositoryRulesetOrganizationPropertyConditionParameters struct {
	Include []*RepositoryRulesetRepositoryPropertyTargetParameters `json:"include"`
	Exclude []*RepositoryRulesetRepositoryPropertyTargetParameters `json:"exclude"`
}

type RepositoryRulesetRepositoryIDsConditionParameters struct {
	RepositoryIDs []int64 `json:"repository_ids,omitempty"`
}

type RepositoryRulesetRepositoryNamesConditionParameters struct {
	Include   []string `json:"include"`
	Exclude   []string `json:"exclude"`
	Protected *bool    `json:"protected,omitempty"`
}

type RepositoryRulesetRepositoryPropertyConditionParameters struct {
	Include []*RepositoryRulesetRepositoryPropertyTargetParameters `json:"include"`
	Exclude []*RepositoryRulesetRepositoryPropertyTargetParameters `json:"exclude"`
}

type RepositoryRulesetRepositoryPropertyTargetParameters struct {
	Name           string   `json:"name"`
	PropertyValues []string `json:"property_values"`
	Source         *string  `json:"source,omitempty"`
}

type RepositoryRulesetOrganizationIDsConditionParameters struct {
	OrganizationIDs []int64 `json:"organization_ids,omitempty"`
}

type RepositoryRulesetOrganizationNamesConditionParameters struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

type RepositoryRule struct {
	Type       RepositoryRuleType `json:"type"`
	Parameters any                `json:"parameters,omitempty"`
}

type RepositoryRulesetRules struct {
	Creation                 *EmptyRuleParameters
	Update                   *UpdateRuleParameters
	Deletion                 *EmptyRuleParameters
	RequiredLinearHistory    *EmptyRuleParameters
	MergeQueue               *MergeQueueRuleParameters
	RequiredDeployments      *RequiredDeploymentsRuleParameters
	RequiredSignatures       *EmptyRuleParameters
	PullRequest              *PullRequestRuleParameters
	RequiredStatusChecks     *RequiredStatusChecksRuleParameters
	NonFastForward           *EmptyRuleParameters
	CommitMessagePattern     *PatternRuleParameters
	CommitAuthorEmailPattern *PatternRuleParameters
	CommitterEmailPattern    *PatternRuleParameters
	BranchNamePattern        *PatternRuleParameters
	TagNamePattern           *PatternRuleParameters
	Workflows                *WorkflowsRuleParameters
	CodeScanning             *CodeScanningRuleParameters
	CopilotCodeReview        *CopilotCodeReviewRuleParameters

	FileExtensionRestriction *FileExtensionRestrictionRuleParameters
	FilePathRestriction      *FilePathRestrictionRuleParameters
	MaxFilePathLength        *MaxFilePathLengthRuleParameters
	MaxFileSize              *MaxFileSizeRuleParameters

	RepositoryCreate     *EmptyRuleParameters
	RepositoryDelete     *EmptyRuleParameters
	RepositoryName       *SimplePatternRuleParameters
	RepositoryTransfer   *EmptyRuleParameters
	RepositoryVisibility *RepositoryVisibilityRuleParameters
}

type BranchRules struct {
	Creation                 []*BranchRuleMetadata
	Update                   []*UpdateBranchRule
	Deletion                 []*BranchRuleMetadata
	RequiredLinearHistory    []*BranchRuleMetadata
	MergeQueue               []*MergeQueueBranchRule
	RequiredDeployments      []*RequiredDeploymentsBranchRule
	RequiredSignatures       []*BranchRuleMetadata
	PullRequest              []*PullRequestBranchRule
	RequiredStatusChecks     []*RequiredStatusChecksBranchRule
	NonFastForward           []*BranchRuleMetadata
	CommitMessagePattern     []*PatternBranchRule
	CommitAuthorEmailPattern []*PatternBranchRule
	CommitterEmailPattern    []*PatternBranchRule
	BranchNamePattern        []*PatternBranchRule
	TagNamePattern           []*PatternBranchRule
	Workflows                []*WorkflowsBranchRule
	CodeScanning             []*CodeScanningBranchRule
	CopilotCodeReview        []*CopilotCodeReviewBranchRule

	FileExtensionRestriction []*FileExtensionRestrictionBranchRule
	FilePathRestriction      []*FilePathRestrictionBranchRule
	MaxFilePathLength        []*MaxFilePathLengthBranchRule
	MaxFileSize              []*MaxFileSizeBranchRule
}

type BranchRuleMetadata struct {
	RulesetSourceType RulesetSourceType `json:"ruleset_source_type"`
	RulesetSource     string            `json:"ruleset_source"`
	RulesetID         int64             `json:"ruleset_id"`
}

type UpdateBranchRule struct {
	BranchRuleMetadata
	Parameters UpdateRuleParameters `json:"parameters"`
}

type MergeQueueBranchRule struct {
	BranchRuleMetadata
	Parameters MergeQueueRuleParameters `json:"parameters"`
}

type RequiredDeploymentsBranchRule struct {
	BranchRuleMetadata
	Parameters RequiredDeploymentsRuleParameters `json:"parameters"`
}

type PullRequestBranchRule struct {
	BranchRuleMetadata
	Parameters PullRequestRuleParameters `json:"parameters"`
}

type RequiredStatusChecksBranchRule struct {
	BranchRuleMetadata
	Parameters RequiredStatusChecksRuleParameters `json:"parameters"`
}

type PatternBranchRule struct {
	BranchRuleMetadata
	Parameters PatternRuleParameters `json:"parameters"`
}

type FilePathRestrictionBranchRule struct {
	BranchRuleMetadata
	Parameters FilePathRestrictionRuleParameters `json:"parameters"`
}

type MaxFilePathLengthBranchRule struct {
	BranchRuleMetadata
	Parameters MaxFilePathLengthRuleParameters `json:"parameters"`
}

type FileExtensionRestrictionBranchRule struct {
	BranchRuleMetadata
	Parameters FileExtensionRestrictionRuleParameters `json:"parameters"`
}

type MaxFileSizeBranchRule struct {
	BranchRuleMetadata
	Parameters MaxFileSizeRuleParameters `json:"parameters"`
}

type WorkflowsBranchRule struct {
	BranchRuleMetadata
	Parameters WorkflowsRuleParameters `json:"parameters"`
}

type CodeScanningBranchRule struct {
	BranchRuleMetadata
	Parameters CodeScanningRuleParameters `json:"parameters"`
}

type CopilotCodeReviewBranchRule struct {
	BranchRuleMetadata
	Parameters CopilotCodeReviewRuleParameters `json:"parameters"`
}

type EmptyRuleParameters struct{}

type UpdateRuleParameters struct {
	UpdateAllowsFetchAndMerge bool `json:"update_allows_fetch_and_merge,omitempty"`
}

type MergeQueueRuleParameters struct {
	CheckResponseTimeoutMinutes  int                   `json:"check_response_timeout_minutes"`
	GroupingStrategy             MergeGroupingStrategy `json:"grouping_strategy"`
	MaxEntriesToBuild            int                   `json:"max_entries_to_build"`
	MaxEntriesToMerge            int                   `json:"max_entries_to_merge"`
	MergeMethod                  MergeQueueMergeMethod `json:"merge_method"`
	MinEntriesToMerge            int                   `json:"min_entries_to_merge"`
	MinEntriesToMergeWaitMinutes int                   `json:"min_entries_to_merge_wait_minutes"`
}

type RequiredDeploymentsRuleParameters struct {
	RequiredDeploymentEnvironments []string `json:"required_deployment_environments"`
}

type PullRequestRuleParameters struct {
	AllowedMergeMethods            []PullRequestMergeMethod   `json:"allowed_merge_methods,omitempty"`
	DismissStaleReviewsOnPush      bool                       `json:"dismiss_stale_reviews_on_push"`
	RequireCodeOwnerReview         bool                       `json:"require_code_owner_review"`
	RequireLastPushApproval        bool                       `json:"require_last_push_approval"`
	RequiredApprovingReviewCount   int                        `json:"required_approving_review_count"`
	RequiredReviewers              []*RulesetRequiredReviewer `json:"required_reviewers,omitempty"`
	RequiredReviewThreadResolution bool                       `json:"required_review_thread_resolution"`
}

type RulesetRequiredReviewer struct {
	MinimumApprovals *int             `json:"minimum_approvals,omitempty"`
	FilePatterns     []string         `json:"file_patterns,omitempty"`
	Reviewer         *RulesetReviewer `json:"reviewer,omitempty"`
}

type RulesetReviewer struct {
	ID   *int64               `json:"id,omitempty"`
	Type *RulesetReviewerType `json:"type,omitempty"`
}

func (r *RulesetReviewer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type RequiredStatusChecksRuleParameters struct {
	DoNotEnforceOnCreate             *bool              `json:"do_not_enforce_on_create,omitempty"`
	RequiredStatusChecks             []*RuleStatusCheck `json:"required_status_checks"`
	StrictRequiredStatusChecksPolicy bool               `json:"strict_required_status_checks_policy"`
}

type RuleStatusCheck struct {
	Context       string `json:"context"`
	IntegrationID *int64 `json:"integration_id,omitempty"`
}

type PatternRuleParameters struct {
	Name *string `json:"name,omitempty"`

	Negate   *bool               `json:"negate,omitempty"`
	Operator PatternRuleOperator `json:"operator"`
	Pattern  string              `json:"pattern"`
}

type FilePathRestrictionRuleParameters struct {
	RestrictedFilePaths []string `json:"restricted_file_paths"`
}

type MaxFilePathLengthRuleParameters struct {
	MaxFilePathLength int `json:"max_file_path_length"`
}

type FileExtensionRestrictionRuleParameters struct {
	RestrictedFileExtensions []string `json:"restricted_file_extensions"`
}

type MaxFileSizeRuleParameters struct {
	MaxFileSize int64 `json:"max_file_size"`
}

type WorkflowsRuleParameters struct {
	DoNotEnforceOnCreate *bool           `json:"do_not_enforce_on_create,omitempty"`
	Workflows            []*RuleWorkflow `json:"workflows"`
}

type RuleWorkflow struct {
	Path         string  `json:"path"`
	Ref          *string `json:"ref,omitempty"`
	RepositoryID *int64  `json:"repository_id,omitempty"`
	SHA          *string `json:"sha,omitempty"`
}

type CodeScanningRuleParameters struct {
	CodeScanningTools []*RuleCodeScanningTool `json:"code_scanning_tools"`
}

type CopilotCodeReviewRuleParameters struct {
	ReviewOnPush            bool `json:"review_on_push"`
	ReviewDraftPullRequests bool `json:"review_draft_pull_requests"`
}

type RuleCodeScanningTool struct {
	AlertsThreshold         CodeScanningAlertsThreshold         `json:"alerts_threshold"`
	SecurityAlertsThreshold CodeScanningSecurityAlertsThreshold `json:"security_alerts_threshold"`
	Tool                    string                              `json:"tool"`
}

type SimplePatternRuleParameters struct {
	Negate  bool   `json:"negate"`
	Pattern string `json:"pattern"`
}

type RepositoryVisibilityRuleParameters struct {
	Internal bool `json:"internal"`
	Private  bool `json:"private"`
	Public   bool `json:"public"`
}

type repositoryRulesetRuleWrapper struct {
	Type       RepositoryRuleType `json:"type"`
	Parameters json.RawMessage    `json:"parameters,omitempty"`
}

func (r RepositoryRulesetRules) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func marshalRepositoryRulesetRule[T any](t RepositoryRuleType, params T) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RepositoryRulesetRules) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type branchRuleWrapper struct {
	Type RepositoryRuleType `json:"type"`
	BranchRuleMetadata
	Parameters json.RawMessage `json:"parameters,omitempty"`
}

func (r *BranchRules) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (r *RepositoryRule) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
