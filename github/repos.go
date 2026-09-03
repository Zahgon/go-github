package github

import (
	"context"
	"encoding/json"
	"errors"
)

const githubBranchNotProtected string = "Branch not protected"

var ErrBranchNotProtected = errors.New("branch is not protected")

type RepositoriesService service

type Repository struct {
	ID                        *int64                 `json:"id,omitempty"`
	NodeID                    *string                `json:"node_id,omitempty"`
	Owner                     *User                  `json:"owner,omitempty"`
	Name                      *string                `json:"name,omitempty"`
	FullName                  *string                `json:"full_name,omitempty"`
	Description               *string                `json:"description,omitempty"`
	Homepage                  *string                `json:"homepage,omitempty"`
	CodeOfConduct             *CodeOfConduct         `json:"code_of_conduct,omitempty"`
	DefaultBranch             *string                `json:"default_branch,omitempty"`
	MasterBranch              *string                `json:"master_branch,omitempty"`
	CreatedAt                 *Timestamp             `json:"created_at,omitempty"`
	PushedAt                  *Timestamp             `json:"pushed_at,omitempty"`
	UpdatedAt                 *Timestamp             `json:"updated_at,omitempty"`
	HTMLURL                   *string                `json:"html_url,omitempty"`
	CloneURL                  *string                `json:"clone_url,omitempty"`
	GitURL                    *string                `json:"git_url,omitempty"`
	MirrorURL                 *string                `json:"mirror_url,omitempty"`
	SSHURL                    *string                `json:"ssh_url,omitempty"`
	SVNURL                    *string                `json:"svn_url,omitempty"`
	Language                  *string                `json:"language,omitempty"`
	Fork                      *bool                  `json:"fork,omitempty"`
	ForksCount                *int                   `json:"forks_count,omitempty"`
	NetworkCount              *int                   `json:"network_count,omitempty"`
	OpenIssuesCount           *int                   `json:"open_issues_count,omitempty"`
	OpenIssues                *int                   `json:"open_issues,omitempty"`
	StargazersCount           *int                   `json:"stargazers_count,omitempty"`
	SubscribersCount          *int                   `json:"subscribers_count,omitempty"`
	WatchersCount             *int                   `json:"watchers_count,omitempty"`
	Watchers                  *int                   `json:"watchers,omitempty"`
	Size                      *int                   `json:"size,omitempty"`
	AutoInit                  *bool                  `json:"auto_init,omitempty"`
	Parent                    *Repository            `json:"parent,omitempty"`
	Source                    *Repository            `json:"source,omitempty"`
	TemplateRepository        *Repository            `json:"template_repository,omitempty"`
	Organization              *Organization          `json:"organization,omitempty"`
	Permissions               *RepositoryPermissions `json:"permissions,omitempty"`
	AllowRebaseMerge          *bool                  `json:"allow_rebase_merge,omitempty"`
	AllowUpdateBranch         *bool                  `json:"allow_update_branch,omitempty"`
	AllowSquashMerge          *bool                  `json:"allow_squash_merge,omitempty"`
	AllowMergeCommit          *bool                  `json:"allow_merge_commit,omitempty"`
	AllowAutoMerge            *bool                  `json:"allow_auto_merge,omitempty"`
	AllowForking              *bool                  `json:"allow_forking,omitempty"`
	WebCommitSignoffRequired  *bool                  `json:"web_commit_signoff_required,omitempty"`
	DeleteBranchOnMerge       *bool                  `json:"delete_branch_on_merge,omitempty"`
	UseSquashPRTitleAsDefault *bool                  `json:"use_squash_pr_title_as_default,omitempty"`
	SquashMergeCommitTitle    *string                `json:"squash_merge_commit_title,omitempty"`
	SquashMergeCommitMessage  *string                `json:"squash_merge_commit_message,omitempty"`
	MergeCommitTitle          *string                `json:"merge_commit_title,omitempty"`
	MergeCommitMessage        *string                `json:"merge_commit_message,omitempty"`
	Topics                    []string               `json:"topics,omitempty"`
	CustomProperties          map[string]any         `json:"custom_properties,omitempty"`
	Archived                  *bool                  `json:"archived,omitempty"`
	Disabled                  *bool                  `json:"disabled,omitempty"`

	License *License `json:"license,omitempty"`

	Private                   *bool   `json:"private,omitempty"`
	HasIssues                 *bool   `json:"has_issues,omitempty"`
	HasWiki                   *bool   `json:"has_wiki,omitempty"`
	HasPages                  *bool   `json:"has_pages,omitempty"`
	HasProjects               *bool   `json:"has_projects,omitempty"`
	HasDownloads              *bool   `json:"has_downloads,omitempty"`
	HasDiscussions            *bool   `json:"has_discussions,omitempty"`
	HasPullRequests           *bool   `json:"has_pull_requests,omitempty"`
	PullRequestCreationPolicy *string `json:"pull_request_creation_policy,omitempty"`
	IsTemplate                *bool   `json:"is_template,omitempty"`
	LicenseTemplate           *string `json:"license_template,omitempty"`
	GitignoreTemplate         *string `json:"gitignore_template,omitempty"`

	SecurityAndAnalysis *SecurityAndAnalysis `json:"security_and_analysis,omitempty"`

	TeamID *int64 `json:"team_id,omitempty"`

	URL              *string `json:"url,omitempty"`
	ArchiveURL       *string `json:"archive_url,omitempty"`
	AssigneesURL     *string `json:"assignees_url,omitempty"`
	BlobsURL         *string `json:"blobs_url,omitempty"`
	BranchesURL      *string `json:"branches_url,omitempty"`
	CollaboratorsURL *string `json:"collaborators_url,omitempty"`
	CommentsURL      *string `json:"comments_url,omitempty"`
	CommitsURL       *string `json:"commits_url,omitempty"`
	CompareURL       *string `json:"compare_url,omitempty"`
	ContentsURL      *string `json:"contents_url,omitempty"`
	ContributorsURL  *string `json:"contributors_url,omitempty"`
	DeploymentsURL   *string `json:"deployments_url,omitempty"`
	DownloadsURL     *string `json:"downloads_url,omitempty"`
	EventsURL        *string `json:"events_url,omitempty"`
	ForksURL         *string `json:"forks_url,omitempty"`
	GitCommitsURL    *string `json:"git_commits_url,omitempty"`
	GitRefsURL       *string `json:"git_refs_url,omitempty"`
	GitTagsURL       *string `json:"git_tags_url,omitempty"`
	HooksURL         *string `json:"hooks_url,omitempty"`
	IssueCommentURL  *string `json:"issue_comment_url,omitempty"`
	IssueEventsURL   *string `json:"issue_events_url,omitempty"`
	IssuesURL        *string `json:"issues_url,omitempty"`
	KeysURL          *string `json:"keys_url,omitempty"`
	LabelsURL        *string `json:"labels_url,omitempty"`
	LanguagesURL     *string `json:"languages_url,omitempty"`
	MergesURL        *string `json:"merges_url,omitempty"`
	MilestonesURL    *string `json:"milestones_url,omitempty"`
	NotificationsURL *string `json:"notifications_url,omitempty"`
	PullsURL         *string `json:"pulls_url,omitempty"`
	ReleasesURL      *string `json:"releases_url,omitempty"`
	StargazersURL    *string `json:"stargazers_url,omitempty"`
	StatusesURL      *string `json:"statuses_url,omitempty"`
	SubscribersURL   *string `json:"subscribers_url,omitempty"`
	SubscriptionURL  *string `json:"subscription_url,omitempty"`
	TagsURL          *string `json:"tags_url,omitempty"`
	TreesURL         *string `json:"trees_url,omitempty"`
	TeamsURL         *string `json:"teams_url,omitempty"`

	TextMatches []*TextMatch `json:"text_matches,omitempty"`

	Visibility *string `json:"visibility,omitempty"`

	RoleName *string `json:"role_name,omitempty"`
}

func (r Repository) String() string { _ = "STUB: not implemented"; return "" }

type BranchListOptions struct {
	Protected *bool `url:"protected,omitempty"`

	ListOptions
}

type RepositoryListOptions struct {
	Visibility string `url:"visibility,omitempty"`

	Affiliation string `url:"affiliation,omitempty"`

	Type string `url:"type,omitempty"`

	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	ListOptions
}

type SecurityAndAnalysis struct {
	AdvancedSecurity             *AdvancedSecurity             `json:"advanced_security,omitempty"`
	SecretScanning               *SecretScanning               `json:"secret_scanning,omitempty"`
	SecretScanningPushProtection *SecretScanningPushProtection `json:"secret_scanning_push_protection,omitempty"`
	DependabotSecurityUpdates    *DependabotSecurityUpdates    `json:"dependabot_security_updates,omitempty"`
	SecretScanningValidityChecks *SecretScanningValidityChecks `json:"secret_scanning_validity_checks,omitempty"`
	CodeSecurity                 *CodeSecurity                 `json:"code_security,omitempty"`
}

type RepositoryPermissions struct {
	Admin    *bool `json:"admin,omitempty"`
	Maintain *bool `json:"maintain,omitempty"`
	Push     *bool `json:"push,omitempty"`
	Triage   *bool `json:"triage,omitempty"`
	Pull     *bool `json:"pull,omitempty"`
}

func (s SecurityAndAnalysis) String() string { _ = "STUB: not implemented"; return "" }

type AdvancedSecurity struct {
	Status *string `json:"status,omitempty"`
}

func (a AdvancedSecurity) String() string { _ = "STUB: not implemented"; return "" }

type SecretScanning struct {
	Status *string `json:"status,omitempty"`
}

func (s SecretScanning) String() string { _ = "STUB: not implemented"; return "" }

type SecretScanningPushProtection struct {
	Status *string `json:"status,omitempty"`
}

func (s SecretScanningPushProtection) String() string { _ = "STUB: not implemented"; return "" }

type DependabotSecurityUpdates struct {
	Status *string `json:"status,omitempty"`
}

func (d DependabotSecurityUpdates) String() string { _ = "STUB: not implemented"; return "" }

type SecretScanningValidityChecks struct {
	Status *string `json:"status,omitempty"`
}

type CodeSecurity struct {
	Status *string `json:"status,omitempty"`
}

func (c CodeSecurity) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /user/repos
//meta:operation GET /users/{username}/repos
func (s *RepositoriesService) List(ctx context.Context, user string, opts *RepositoryListOptions) ([]*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type RepositoryListByUserOptions struct {
	Type string `url:"type,omitempty"`

	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	ListOptions
}

//meta:operation GET /users/{username}/repos
func (s *RepositoriesService) ListByUser(ctx context.Context, user string, opts *RepositoryListByUserOptions) ([]*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type RepositoryListByAuthenticatedUserOptions struct {
	Visibility string `url:"visibility,omitempty"`

	Affiliation string `url:"affiliation,omitempty"`

	Type string `url:"type,omitempty"`

	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	ListOptions
}

//meta:operation GET /user/repos
func (s *RepositoriesService) ListByAuthenticatedUser(ctx context.Context, opts *RepositoryListByAuthenticatedUserOptions) ([]*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type RepositoryListByOrgOptions struct {
	Type string `url:"type,omitempty"`

	Sort string `url:"sort,omitempty"`

	Direction string `url:"direction,omitempty"`

	ListOptions
}

//meta:operation GET /orgs/{org}/repos
func (s *RepositoriesService) ListByOrg(ctx context.Context, org string, opts *RepositoryListByOrgOptions) ([]*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type RepositoryListAllOptions struct {
	Since int64 `url:"since,omitempty"`
}

//meta:operation GET /repositories
func (s *RepositoriesService) ListAll(ctx context.Context, opts *RepositoryListAllOptions) ([]*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type createRepoRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Homepage    *string `json:"homepage,omitempty"`

	Private        *bool   `json:"private,omitempty"`
	Visibility     *string `json:"visibility,omitempty"`
	HasIssues      *bool   `json:"has_issues,omitempty"`
	HasProjects    *bool   `json:"has_projects,omitempty"`
	HasWiki        *bool   `json:"has_wiki,omitempty"`
	HasDiscussions *bool   `json:"has_discussions,omitempty"`
	IsTemplate     *bool   `json:"is_template,omitempty"`

	TeamID *int64 `json:"team_id,omitempty"`

	AutoInit                  *bool          `json:"auto_init,omitempty"`
	GitignoreTemplate         *string        `json:"gitignore_template,omitempty"`
	LicenseTemplate           *string        `json:"license_template,omitempty"`
	AllowSquashMerge          *bool          `json:"allow_squash_merge,omitempty"`
	AllowMergeCommit          *bool          `json:"allow_merge_commit,omitempty"`
	AllowRebaseMerge          *bool          `json:"allow_rebase_merge,omitempty"`
	AllowUpdateBranch         *bool          `json:"allow_update_branch,omitempty"`
	AllowAutoMerge            *bool          `json:"allow_auto_merge,omitempty"`
	AllowForking              *bool          `json:"allow_forking,omitempty"`
	DeleteBranchOnMerge       *bool          `json:"delete_branch_on_merge,omitempty"`
	UseSquashPRTitleAsDefault *bool          `json:"use_squash_pr_title_as_default,omitempty"`
	SquashMergeCommitTitle    *string        `json:"squash_merge_commit_title,omitempty"`
	SquashMergeCommitMessage  *string        `json:"squash_merge_commit_message,omitempty"`
	MergeCommitTitle          *string        `json:"merge_commit_title,omitempty"`
	MergeCommitMessage        *string        `json:"merge_commit_message,omitempty"`
	CustomProperties          map[string]any `json:"custom_properties,omitempty"`
}

//meta:operation POST /orgs/{org}/repos
//meta:operation POST /user/repos
func (s *RepositoriesService) Create(ctx context.Context, org string, repo *Repository) (*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type TemplateRepoRequest struct {
	Name               string  `json:"name"`
	Owner              *string `json:"owner,omitempty"`
	Description        *string `json:"description,omitempty"`
	IncludeAllBranches *bool   `json:"include_all_branches,omitempty"`
	Private            *bool   `json:"private,omitempty"`
}

//meta:operation POST /repos/{template_owner}/{template_repo}/generate
func (s *RepositoriesService) CreateFromTemplate(ctx context.Context, templateOwner, templateRepo string, body TemplateRepoRequest) (*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}
func (s *RepositoriesService) Get(ctx context.Context, owner, repo string) (*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}
func (s *RepositoriesService) GetCodeOfConduct(ctx context.Context, owner, repo string) (*CodeOfConduct, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repositories/{repository_id}
func (s *RepositoriesService) GetByID(ctx context.Context, id int64) (*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}
func (s *RepositoriesService) Edit(ctx context.Context, owner, repo string, body *Repository) (*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}
func (s *RepositoriesService) Delete(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Contributor struct {
	Login             *string `json:"login,omitempty"`
	ID                *int64  `json:"id,omitempty"`
	NodeID            *string `json:"node_id,omitempty"`
	AvatarURL         *string `json:"avatar_url,omitempty"`
	GravatarID        *string `json:"gravatar_id,omitempty"`
	URL               *string `json:"url,omitempty"`
	HTMLURL           *string `json:"html_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	Type              *string `json:"type,omitempty"`
	SiteAdmin         *bool   `json:"site_admin,omitempty"`
	Contributions     *int    `json:"contributions,omitempty"`
	Name              *string `json:"name,omitempty"`
	Email             *string `json:"email,omitempty"`
}

type ListContributorsOptions struct {
	Anon string `url:"anon,omitempty"`

	ListOptions
}

//meta:operation GET /repos/{owner}/{repo}/vulnerability-alerts
func (s *RepositoriesService) GetVulnerabilityAlerts(ctx context.Context, owner, repository string) (bool, *Response, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/vulnerability-alerts
func (s *RepositoriesService) EnableVulnerabilityAlerts(ctx context.Context, owner, repository string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/vulnerability-alerts
func (s *RepositoriesService) DisableVulnerabilityAlerts(ctx context.Context, owner, repository string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/automated-security-fixes
func (s *RepositoriesService) GetAutomatedSecurityFixes(ctx context.Context, owner, repository string) (*AutomatedSecurityFixes, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/automated-security-fixes
func (s *RepositoriesService) EnableAutomatedSecurityFixes(ctx context.Context, owner, repository string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/automated-security-fixes
func (s *RepositoriesService) DisableAutomatedSecurityFixes(ctx context.Context, owner, repository string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/contributors
func (s *RepositoriesService) ListContributors(ctx context.Context, owner, repository string, opts *ListContributorsOptions) ([]*Contributor, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/languages
func (s *RepositoriesService) ListLanguages(ctx context.Context, owner, repo string) (map[string]int, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/teams
func (s *RepositoriesService) ListTeams(ctx context.Context, owner, repo string, opts *ListOptions) ([]*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type RepositoryTag struct {
	Name       *string `json:"name,omitempty"`
	Commit     *Commit `json:"commit,omitempty"`
	ZipballURL *string `json:"zipball_url,omitempty"`
	TarballURL *string `json:"tarball_url,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/tags
func (s *RepositoriesService) ListTags(ctx context.Context, owner, repo string, opts *ListOptions) ([]*RepositoryTag, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type Branch struct {
	Name      *string           `json:"name,omitempty"`
	Commit    *RepositoryCommit `json:"commit,omitempty"`
	Protected *bool             `json:"protected,omitempty"`

	Protection    *Protection `json:"protection,omitempty"`
	ProtectionURL *string     `json:"protection_url,omitempty"`
}

type Protection struct {
	RequiredStatusChecks           *RequiredStatusChecks           `json:"required_status_checks"`
	RequiredPullRequestReviews     *PullRequestReviewsEnforcement  `json:"required_pull_request_reviews"`
	EnforceAdmins                  *AdminEnforcement               `json:"enforce_admins"`
	Restrictions                   *BranchRestrictions             `json:"restrictions"`
	RequireLinearHistory           *RequireLinearHistory           `json:"required_linear_history"`
	AllowForcePushes               *AllowForcePushes               `json:"allow_force_pushes"`
	AllowDeletions                 *AllowDeletions                 `json:"allow_deletions"`
	RequiredConversationResolution *RequiredConversationResolution `json:"required_conversation_resolution"`
	BlockCreations                 *BlockCreations                 `json:"block_creations,omitempty"`
	LockBranch                     *LockBranch                     `json:"lock_branch,omitempty"`
	AllowForkSyncing               *AllowForkSyncing               `json:"allow_fork_syncing,omitempty"`
	RequiredSignatures             *SignaturesProtectedBranch      `json:"required_signatures,omitempty"`
	URL                            *string                         `json:"url,omitempty"`
}

type BlockCreations struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type LockBranch struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type AllowForkSyncing struct {
	Enabled *bool `json:"enabled,omitempty"`
}

type BranchProtectionRule struct {
	ID                                       *int64     `json:"id,omitempty"`
	RepositoryID                             *int64     `json:"repository_id,omitempty"`
	Name                                     *string    `json:"name,omitempty"`
	CreatedAt                                *Timestamp `json:"created_at,omitempty"`
	UpdatedAt                                *Timestamp `json:"updated_at,omitempty"`
	PullRequestReviewsEnforcementLevel       *string    `json:"pull_request_reviews_enforcement_level,omitempty"`
	RequiredApprovingReviewCount             *int       `json:"required_approving_review_count,omitempty"`
	DismissStaleReviewsOnPush                *bool      `json:"dismiss_stale_reviews_on_push,omitempty"`
	AuthorizedDismissalActorsOnly            *bool      `json:"authorized_dismissal_actors_only,omitempty"`
	IgnoreApprovalsFromContributors          *bool      `json:"ignore_approvals_from_contributors,omitempty"`
	RequireCodeOwnerReview                   *bool      `json:"require_code_owner_review,omitempty"`
	RequiredStatusChecks                     []string   `json:"required_status_checks,omitempty"`
	RequiredStatusChecksEnforcementLevel     *string    `json:"required_status_checks_enforcement_level,omitempty"`
	StrictRequiredStatusChecksPolicy         *bool      `json:"strict_required_status_checks_policy,omitempty"`
	SignatureRequirementEnforcementLevel     *string    `json:"signature_requirement_enforcement_level,omitempty"`
	LinearHistoryRequirementEnforcementLevel *string    `json:"linear_history_requirement_enforcement_level,omitempty"`
	AdminEnforced                            *bool      `json:"admin_enforced,omitempty"`
	AllowForcePushesEnforcementLevel         *string    `json:"allow_force_pushes_enforcement_level,omitempty"`
	AllowDeletionsEnforcementLevel           *string    `json:"allow_deletions_enforcement_level,omitempty"`
	MergeQueueEnforcementLevel               *string    `json:"merge_queue_enforcement_level,omitempty"`
	RequiredDeploymentsEnforcementLevel      *string    `json:"required_deployments_enforcement_level,omitempty"`
	RequiredConversationResolutionLevel      *string    `json:"required_conversation_resolution_level,omitempty"`
	AuthorizedActorsOnly                     *bool      `json:"authorized_actors_only,omitempty"`
	AuthorizedActorNames                     []string   `json:"authorized_actor_names,omitempty"`
	RequireLastPushApproval                  *bool      `json:"require_last_push_approval,omitempty"`
}

type ProtectionChanges struct {
	AdminEnforced                            *AdminEnforcedChanges                            `json:"admin_enforced,omitempty"`
	AllowDeletionsEnforcementLevel           *AllowDeletionsEnforcementLevelChanges           `json:"allow_deletions_enforcement_level,omitempty"`
	AuthorizedActorNames                     *AuthorizedActorNames                            `json:"authorized_actor_names,omitempty"`
	AuthorizedActorsOnly                     *AuthorizedActorsOnly                            `json:"authorized_actors_only,omitempty"`
	AuthorizedDismissalActorsOnly            *AuthorizedDismissalActorsOnlyChanges            `json:"authorized_dismissal_actors_only,omitempty"`
	CreateProtected                          *CreateProtectedChanges                          `json:"create_protected,omitempty"`
	DismissStaleReviewsOnPush                *DismissStaleReviewsOnPushChanges                `json:"dismiss_stale_reviews_on_push,omitempty"`
	LinearHistoryRequirementEnforcementLevel *LinearHistoryRequirementEnforcementLevelChanges `json:"linear_history_requirement_enforcement_level,omitempty"`
	PullRequestReviewsEnforcementLevel       *PullRequestReviewsEnforcementLevelChanges       `json:"pull_request_reviews_enforcement_level,omitempty"`
	RequireCodeOwnerReview                   *RequireCodeOwnerReviewChanges                   `json:"require_code_owner_review,omitempty"`
	RequiredConversationResolutionLevel      *RequiredConversationResolutionLevelChanges      `json:"required_conversation_resolution_level,omitempty"`
	RequiredDeploymentsEnforcementLevel      *RequiredDeploymentsEnforcementLevelChanges      `json:"required_deployments_enforcement_level,omitempty"`
	RequiredStatusChecks                     *RequiredStatusChecksChanges                     `json:"required_status_checks,omitempty"`
	RequiredStatusChecksEnforcementLevel     *RequiredStatusChecksEnforcementLevelChanges     `json:"required_status_checks_enforcement_level,omitempty"`
	SignatureRequirementEnforcementLevel     *SignatureRequirementEnforcementLevelChanges     `json:"signature_requirement_enforcement_level,omitempty"`
	RequireLastPushApproval                  *RequireLastPushApprovalChanges                  `json:"require_last_push_approval,omitempty"`
}

type AdminEnforcedChanges struct {
	From *bool `json:"from,omitempty"`
}

type AllowDeletionsEnforcementLevelChanges struct {
	From *string `json:"from,omitempty"`
}

type AuthorizedActorNames struct {
	From []string `json:"from,omitempty"`
}

type AuthorizedActorsOnly struct {
	From *bool `json:"from,omitempty"`
}

type AuthorizedDismissalActorsOnlyChanges struct {
	From *bool `json:"from,omitempty"`
}

type CreateProtectedChanges struct {
	From *bool `json:"from,omitempty"`
}

type DismissStaleReviewsOnPushChanges struct {
	From *bool `json:"from,omitempty"`
}

type LinearHistoryRequirementEnforcementLevelChanges struct {
	From *string `json:"from,omitempty"`
}

type PullRequestReviewsEnforcementLevelChanges struct {
	From *string `json:"from,omitempty"`
}

type RequireCodeOwnerReviewChanges struct {
	From *bool `json:"from,omitempty"`
}

type RequiredConversationResolutionLevelChanges struct {
	From *string `json:"from,omitempty"`
}

type RequiredDeploymentsEnforcementLevelChanges struct {
	From *string `json:"from,omitempty"`
}

type RequiredStatusChecksChanges struct {
	From []string `json:"from,omitempty"`
}

type RequiredStatusChecksEnforcementLevelChanges struct {
	From *string `json:"from,omitempty"`
}

type SignatureRequirementEnforcementLevelChanges struct {
	From *string `json:"from,omitempty"`
}

type RequireLastPushApprovalChanges struct {
	From *bool `json:"from,omitempty"`
}

type ProtectionRequest struct {
	RequiredStatusChecks       *RequiredStatusChecks                 `json:"required_status_checks"`
	RequiredPullRequestReviews *PullRequestReviewsEnforcementRequest `json:"required_pull_request_reviews"`
	EnforceAdmins              bool                                  `json:"enforce_admins"`
	Restrictions               *BranchRestrictionsRequest            `json:"restrictions"`

	RequireLinearHistory *bool `json:"required_linear_history,omitempty"`

	AllowForcePushes *bool `json:"allow_force_pushes,omitempty"`

	AllowDeletions *bool `json:"allow_deletions,omitempty"`

	RequiredConversationResolution *bool `json:"required_conversation_resolution,omitempty"`

	BlockCreations *bool `json:"block_creations,omitempty"`

	LockBranch *bool `json:"lock_branch,omitempty"`

	AllowForkSyncing *bool `json:"allow_fork_syncing,omitempty"`
}

type RequiredStatusChecks struct {
	Strict bool `json:"strict"`

	Contexts *[]string `json:"contexts,omitempty"`

	Checks      *[]*RequiredStatusCheck `json:"checks,omitempty"`
	ContextsURL *string                 `json:"contexts_url,omitempty"`
	URL         *string                 `json:"url,omitempty"`
}

type RequiredStatusChecksRequest struct {
	Strict *bool `json:"strict,omitempty"`

	Contexts []string               `json:"contexts,omitempty"`
	Checks   []*RequiredStatusCheck `json:"checks,omitempty"`
}

type RequiredStatusCheck struct {
	Context string `json:"context"`

	AppID *int64 `json:"app_id,omitempty"`
}

type PullRequestReviewsEnforcement struct {
	BypassPullRequestAllowances *BypassPullRequestAllowances `json:"bypass_pull_request_allowances,omitempty"`

	DismissalRestrictions *DismissalRestrictions `json:"dismissal_restrictions,omitempty"`

	DismissStaleReviews bool `json:"dismiss_stale_reviews"`

	RequireCodeOwnerReviews bool `json:"require_code_owner_reviews"`

	RequiredApprovingReviewCount int `json:"required_approving_review_count"`

	RequireLastPushApproval bool `json:"require_last_push_approval"`
}

type PullRequestReviewsEnforcementRequest struct {
	BypassPullRequestAllowancesRequest *BypassPullRequestAllowancesRequest `json:"bypass_pull_request_allowances,omitempty"`

	DismissalRestrictionsRequest *DismissalRestrictionsRequest `json:"dismissal_restrictions,omitempty"`

	DismissStaleReviews bool `json:"dismiss_stale_reviews"`

	RequireCodeOwnerReviews bool `json:"require_code_owner_reviews"`

	RequiredApprovingReviewCount int `json:"required_approving_review_count"`

	RequireLastPushApproval *bool `json:"require_last_push_approval,omitempty"`
}

type PullRequestReviewsEnforcementUpdate struct {
	BypassPullRequestAllowancesRequest *BypassPullRequestAllowancesRequest `json:"bypass_pull_request_allowances,omitempty"`

	DismissalRestrictionsRequest *DismissalRestrictionsRequest `json:"dismissal_restrictions,omitempty"`

	DismissStaleReviews *bool `json:"dismiss_stale_reviews,omitempty"`

	RequireCodeOwnerReviews *bool `json:"require_code_owner_reviews,omitempty"`

	RequiredApprovingReviewCount int `json:"required_approving_review_count"`

	RequireLastPushApproval *bool `json:"require_last_push_approval,omitempty"`
}

type RequireLinearHistory struct {
	Enabled bool `json:"enabled"`
}

type AllowDeletions struct {
	Enabled bool `json:"enabled"`
}

type AllowForcePushes struct {
	Enabled bool `json:"enabled"`
}

type RequiredConversationResolution struct {
	Enabled bool `json:"enabled"`
}

type AdminEnforcement struct {
	URL     *string `json:"url,omitempty"`
	Enabled bool    `json:"enabled"`
}

type BranchRestrictions struct {
	Users []*User `json:"users"`

	Teams []*Team `json:"teams"`

	Apps []*App `json:"apps"`
}

type BranchRestrictionsRequest struct {
	Users []string `json:"users"`

	Teams []string `json:"teams"`

	Apps []string `json:"apps"`
}

type BypassPullRequestAllowances struct {
	Users []*User `json:"users"`

	Teams []*Team `json:"teams"`

	Apps []*App `json:"apps"`
}

type BypassPullRequestAllowancesRequest struct {
	Users []string `json:"users"`

	Teams []string `json:"teams"`

	Apps []string `json:"apps"`
}

type DismissalRestrictions struct {
	Users []*User `json:"users"`

	Teams []*Team `json:"teams"`

	Apps []*App `json:"apps"`
}

type DismissalRestrictionsRequest struct {
	Users *[]string `json:"users,omitempty"`

	Teams *[]string `json:"teams,omitempty"`

	Apps *[]string `json:"apps,omitempty"`
}

type SignaturesProtectedBranch struct {
	URL *string `json:"url,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`
}

type AutomatedSecurityFixes struct {
	Enabled *bool `json:"enabled"`
	Paused  *bool `json:"paused"`
}

//meta:operation GET /repos/{owner}/{repo}/branches
func (s *RepositoriesService) ListBranches(ctx context.Context, owner, repo string, opts *BranchListOptions) ([]*Branch, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/branches/{branch}
func (s *RepositoriesService) GetBranch(ctx context.Context, owner, repo, branch string, maxRedirects int) (*Branch, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type renameBranchRequest struct {
	NewName string `json:"new_name"`
}

//meta:operation POST /repos/{owner}/{repo}/branches/{branch}/rename
func (s *RepositoriesService) RenameBranch(ctx context.Context, owner, repo, branch, newName string) (*Branch, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/branches/{branch}/protection
func (s *RepositoriesService) GetBranchProtection(ctx context.Context, owner, repo, branch string) (*Protection, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks
func (s *RepositoriesService) GetRequiredStatusChecks(ctx context.Context, owner, repo, branch string) (*RequiredStatusChecks, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks/contexts
func (s *RepositoriesService) ListRequiredStatusChecksContexts(ctx context.Context, owner, repo, branch string) (contexts []string, resp *Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/branches/{branch}/protection
func (s *RepositoriesService) UpdateBranchProtection(ctx context.Context, owner, repo, branch string, body *ProtectionRequest) (*Protection, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/branches/{branch}/protection
func (s *RepositoriesService) RemoveBranchProtection(ctx context.Context, owner, repo, branch string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/branches/{branch}/protection/required_signatures
func (s *RepositoriesService) GetSignaturesProtectedBranch(ctx context.Context, owner, repo, branch string) (*SignaturesProtectedBranch, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/branches/{branch}/protection/required_signatures
func (s *RepositoriesService) RequireSignaturesOnProtectedBranch(ctx context.Context, owner, repo, branch string) (*SignaturesProtectedBranch, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/branches/{branch}/protection/required_signatures
func (s *RepositoriesService) OptionalSignaturesOnProtectedBranch(ctx context.Context, owner, repo, branch string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks
func (s *RepositoriesService) UpdateRequiredStatusChecks(ctx context.Context, owner, repo, branch string, body *RequiredStatusChecksRequest) (*RequiredStatusChecks, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks
func (s *RepositoriesService) RemoveRequiredStatusChecks(ctx context.Context, owner, repo, branch string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/license
func (s *RepositoriesService) License(ctx context.Context, owner, repo string) (*RepositoryLicense, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/branches/{branch}/protection/required_pull_request_reviews
func (s *RepositoriesService) GetPullRequestReviewEnforcement(ctx context.Context, owner, repo, branch string) (*PullRequestReviewsEnforcement, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/branches/{branch}/protection/required_pull_request_reviews
func (s *RepositoriesService) UpdatePullRequestReviewEnforcement(ctx context.Context, owner, repo, branch string, body *PullRequestReviewsEnforcementUpdate) (*PullRequestReviewsEnforcement, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /repos/{owner}/{repo}/branches/{branch}/protection/required_pull_request_reviews
func (s *RepositoriesService) DisableDismissalRestrictions(ctx context.Context, owner, repo, branch string) (*PullRequestReviewsEnforcement, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/branches/{branch}/protection/required_pull_request_reviews
func (s *RepositoriesService) RemovePullRequestReviewEnforcement(ctx context.Context, owner, repo, branch string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/branches/{branch}/protection/enforce_admins
func (s *RepositoriesService) GetAdminEnforcement(ctx context.Context, owner, repo, branch string) (*AdminEnforcement, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/branches/{branch}/protection/enforce_admins
func (s *RepositoriesService) AddAdminEnforcement(ctx context.Context, owner, repo, branch string) (*AdminEnforcement, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/branches/{branch}/protection/enforce_admins
func (s *RepositoriesService) RemoveAdminEnforcement(ctx context.Context, owner, repo, branch string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type repositoryTopics struct {
	Names []string `json:"names"`
}

//meta:operation GET /repos/{owner}/{repo}/topics
func (s *RepositoriesService) ListAllTopics(ctx context.Context, owner, repo string, opts *ListOptions) ([]string, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/topics
func (s *RepositoriesService) ReplaceAllTopics(ctx context.Context, owner, repo string, topics []string) ([]string, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps
func (s *RepositoriesService) ListApps(ctx context.Context, owner, repo, branch string) ([]*App, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps
func (s *RepositoriesService) ListAppRestrictions(ctx context.Context, owner, repo, branch string) ([]*App, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps
func (s *RepositoriesService) ReplaceAppRestrictions(ctx context.Context, owner, repo, branch string, body []string) ([]*App, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps
func (s *RepositoriesService) AddAppRestrictions(ctx context.Context, owner, repo, branch string, body []string) ([]*App, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps
func (s *RepositoriesService) RemoveAppRestrictions(ctx context.Context, owner, repo, branch string, apps []string) ([]*App, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams
func (s *RepositoriesService) ListTeamRestrictions(ctx context.Context, owner, repo, branch string) ([]*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams
func (s *RepositoriesService) ReplaceTeamRestrictions(ctx context.Context, owner, repo, branch string, body []string) ([]*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams
func (s *RepositoriesService) AddTeamRestrictions(ctx context.Context, owner, repo, branch string, body []string) ([]*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams
func (s *RepositoriesService) RemoveTeamRestrictions(ctx context.Context, owner, repo, branch string, teams []string) ([]*Team, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users
func (s *RepositoriesService) ListUserRestrictions(ctx context.Context, owner, repo, branch string) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PUT /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users
func (s *RepositoriesService) ReplaceUserRestrictions(ctx context.Context, owner, repo, branch string, body []string) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users
func (s *RepositoriesService) AddUserRestrictions(ctx context.Context, owner, repo, branch string, body []string) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users
func (s *RepositoriesService) RemoveUserRestrictions(ctx context.Context, owner, repo, branch string, users []string) ([]*User, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type TransferRequest struct {
	NewOwner string  `json:"new_owner"`
	NewName  *string `json:"new_name,omitempty"`
	TeamID   []int64 `json:"team_ids,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/transfer
func (s *RepositoriesService) Transfer(ctx context.Context, owner, repo string, transfer TransferRequest) (*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type DispatchRequestOptions struct {
	EventType string `json:"event_type"`

	ClientPayload *json.RawMessage `json:"client_payload,omitempty"`
}

//meta:operation POST /repos/{owner}/{repo}/dispatches
func (s *RepositoriesService) Dispatch(ctx context.Context, owner, repo string, opts DispatchRequestOptions) (*Repository, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func isBranchNotProtected(err error) bool { _ = "STUB: not implemented"; return false }

//meta:operation PUT /repos/{owner}/{repo}/private-vulnerability-reporting
func (s *RepositoriesService) EnablePrivateReporting(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation DELETE /repos/{owner}/{repo}/private-vulnerability-reporting
func (s *RepositoriesService) DisablePrivateReporting(ctx context.Context, owner, repo string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type checkPrivateReporting struct {
	Enabled bool `json:"enabled,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/private-vulnerability-reporting
func (s *RepositoriesService) IsPrivateReportingEnabled(ctx context.Context, owner, repo string) (bool, *Response, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

type ListRepositoryActivityOptions struct {
	Direction string `url:"direction,omitempty"`

	PerPage int `url:"per_page,omitempty"`

	Before string `url:"before,omitempty"`

	After string `url:"after,omitempty"`

	Ref string `url:"ref,omitempty"`

	Actor string `url:"actor,omitempty"`

	TimePeriod string `url:"time_period,omitempty"`

	ActivityType string `url:"activity_type,omitempty"`
}

type RepositoryActor struct {
	Login             *string `json:"login,omitempty"`
	ID                *int64  `json:"id,omitempty"`
	NodeID            *string `json:"node_id,omitempty"`
	AvatarURL         *string `json:"avatar_url,omitempty"`
	GravatarID        *string `json:"gravatar_id,omitempty"`
	URL               *string `json:"url,omitempty"`
	HTMLURL           *string `json:"html_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	Type              *string `json:"type,omitempty"`
	UserViewType      *string `json:"user_view_type,omitempty"`
	SiteAdmin         *bool   `json:"site_admin,omitempty"`
}

type RepositoryActivity struct {
	ID           int64            `json:"id"`
	NodeID       string           `json:"node_id"`
	Before       string           `json:"before"`
	After        string           `json:"after"`
	Ref          string           `json:"ref"`
	Timestamp    *Timestamp       `json:"timestamp"`
	ActivityType string           `json:"activity_type"`
	Actor        *RepositoryActor `json:"actor,omitempty"`
}

//meta:operation GET /repos/{owner}/{repo}/activity
func (s *RepositoriesService) ListRepositoryActivities(ctx context.Context, owner, repo string, opts *ListRepositoryActivityOptions) ([]*RepositoryActivity, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
