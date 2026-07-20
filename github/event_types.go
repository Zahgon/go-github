package github

import "encoding/json"

type RequestedAction struct {
	Identifier string `json:"identifier"`
}

type BranchProtectionRuleEvent struct {
	Action       *string               `json:"action,omitempty"`
	Rule         *BranchProtectionRule `json:"rule,omitempty"`
	Changes      *ProtectionChanges    `json:"changes,omitempty"`
	Repo         *Repository           `json:"repository,omitempty"`
	Org          *Organization         `json:"organization,omitempty"`
	Sender       *User                 `json:"sender,omitempty"`
	Installation *Installation         `json:"installation,omitempty"`
}

type BranchProtectionConfigurationEvent struct {
	Action       *string       `json:"action,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Enterprise   *Enterprise   `json:"enterprise,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type CheckRunEvent struct {
	CheckRun *CheckRun `json:"check_run,omitempty"`

	Action *string `json:"action,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	RequestedAction *RequestedAction `json:"requested_action,omitempty"`
}

type CheckSuiteEvent struct {
	CheckSuite *CheckSuite `json:"check_suite,omitempty"`

	Action *string `json:"action,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type CommitCommentEvent struct {
	Comment *RepositoryComment `json:"comment,omitempty"`

	Action       *string       `json:"action,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type ContentReferenceEvent struct {
	Action           *string           `json:"action,omitempty"`
	ContentReference *ContentReference `json:"content_reference,omitempty"`
	Repo             *Repository       `json:"repository,omitempty"`
	Sender           *User             `json:"sender,omitempty"`
	Installation     *Installation     `json:"installation,omitempty"`
}

type CreateEvent struct {
	Ref *string `json:"ref,omitempty"`

	RefType      *string `json:"ref_type,omitempty"`
	MasterBranch *string `json:"master_branch,omitempty"`
	Description  *string `json:"description,omitempty"`
	PusherType   *string `json:"pusher_type,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type CustomPropertyEvent struct {
	Action     *string         `json:"action,omitempty"`
	Definition *CustomProperty `json:"definition,omitempty"`

	Enterprise   *Enterprise   `json:"enterprise,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
}

type CustomPropertyValuesEvent struct {
	Action            *string                `json:"action,omitempty"`
	NewPropertyValues []*CustomPropertyValue `json:"new_property_values,omitempty"`
	OldPropertyValues []*CustomPropertyValue `json:"old_property_values,omitempty"`

	Enterprise   *Enterprise   `json:"enterprise,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
}

type DeleteEvent struct {
	Ref *string `json:"ref,omitempty"`

	RefType *string `json:"ref_type,omitempty"`

	PusherType   *string       `json:"pusher_type,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type DependabotAlertEvent struct {
	Action *string          `json:"action,omitempty"`
	Alert  *DependabotAlert `json:"alert,omitempty"`

	Installation *Installation `json:"installation,omitempty"`
	Enterprise   *Enterprise   `json:"enterprise,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`

	Organization *Organization `json:"organization,omitempty"`
}

type DeployKeyEvent struct {
	Action *string `json:"action,omitempty"`

	Key *Key `json:"key,omitempty"`

	Repo *Repository `json:"repository,omitempty"`

	Organization *Organization `json:"organization,omitempty"`

	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type DeploymentEvent struct {
	Deployment  *Deployment  `json:"deployment,omitempty"`
	Repo        *Repository  `json:"repository,omitempty"`
	Workflow    *Workflow    `json:"workflow,omitempty"`
	WorkflowRun *WorkflowRun `json:"workflow_run,omitempty"`

	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type DeploymentProtectionRuleEvent struct {
	Action      *string `json:"action,omitempty"`
	Environment *string `json:"environment,omitempty"`
	Event       *string `json:"event,omitempty"`

	DeploymentCallbackURL *string        `json:"deployment_callback_url,omitempty"`
	Deployment            *Deployment    `json:"deployment,omitempty"`
	Repo                  *Repository    `json:"repository,omitempty"`
	Organization          *Organization  `json:"organization,omitempty"`
	PullRequests          []*PullRequest `json:"pull_requests,omitempty"`
	Sender                *User          `json:"sender,omitempty"`
	Installation          *Installation  `json:"installation,omitempty"`
}

type DeploymentReviewEvent struct {
	Action *string `json:"action,omitempty"`

	Requester   *User   `json:"requester,omitempty"`
	Environment *string `json:"environment,omitempty"`

	Approver        *User             `json:"approver,omitempty"`
	Comment         *string           `json:"comment,omitempty"`
	WorkflowJobRuns []*WorkflowJobRun `json:"workflow_job_runs,omitempty"`

	Enterprise     *Enterprise         `json:"enterprise,omitempty"`
	Installation   *Installation       `json:"installation,omitempty"`
	Organization   *Organization       `json:"organization,omitempty"`
	Repo           *Repository         `json:"repository,omitempty"`
	Reviewers      []*RequiredReviewer `json:"reviewers,omitempty"`
	Sender         *User               `json:"sender,omitempty"`
	Since          *string             `json:"since,omitempty"`
	WorkflowJobRun *WorkflowJobRun     `json:"workflow_job_run,omitempty"`
	WorkflowRun    *WorkflowRun        `json:"workflow_run,omitempty"`
}

type WorkflowJobRun struct {
	Conclusion  *string    `json:"conclusion,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	Environment *string    `json:"environment,omitempty"`
	HTMLURL     *string    `json:"html_url,omitempty"`
	ID          *int64     `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Status      *string    `json:"status,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
}

type DeploymentStatusEvent struct {
	Action           *string           `json:"action,omitempty"`
	Deployment       *Deployment       `json:"deployment,omitempty"`
	DeploymentStatus *DeploymentStatus `json:"deployment_status,omitempty"`
	Repo             *Repository       `json:"repository,omitempty"`

	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type DiscussionCommentEvent struct {
	Action       *string            `json:"action,omitempty"`
	Discussion   *Discussion        `json:"discussion,omitempty"`
	Comment      *CommentDiscussion `json:"comment,omitempty"`
	Repo         *Repository        `json:"repository,omitempty"`
	Org          *Organization      `json:"organization,omitempty"`
	Sender       *User              `json:"sender,omitempty"`
	Installation *Installation      `json:"installation,omitempty"`
}

type CommentDiscussion struct {
	AuthorAssociation *string    `json:"author_association,omitempty"`
	Body              *string    `json:"body,omitempty"`
	ChildCommentCount *int       `json:"child_comment_count,omitempty"`
	CreatedAt         *Timestamp `json:"created_at,omitempty"`
	DiscussionID      *int64     `json:"discussion_id,omitempty"`
	HTMLURL           *string    `json:"html_url,omitempty"`
	ID                *int64     `json:"id,omitempty"`
	NodeID            *string    `json:"node_id,omitempty"`
	ParentID          *int64     `json:"parent_id,omitempty"`
	Reactions         *Reactions `json:"reactions,omitempty"`
	RepositoryURL     *string    `json:"repository_url,omitempty"`
	UpdatedAt         *Timestamp `json:"updated_at,omitempty"`
	User              *User      `json:"user,omitempty"`
}

type DiscussionEvent struct {
	Action       *string       `json:"action,omitempty"`
	Discussion   *Discussion   `json:"discussion,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type Discussion struct {
	RepositoryURL      *string             `json:"repository_url,omitempty"`
	DiscussionCategory *DiscussionCategory `json:"category,omitempty"`
	AnswerHTMLURL      *string             `json:"answer_html_url,omitempty"`
	AnswerChosenAt     *Timestamp          `json:"answer_chosen_at,omitempty"`
	AnswerChosenBy     *string             `json:"answer_chosen_by,omitempty"`
	HTMLURL            *string             `json:"html_url,omitempty"`
	ID                 *int64              `json:"id,omitempty"`
	NodeID             *string             `json:"node_id,omitempty"`
	Number             *int                `json:"number,omitempty"`
	Title              *string             `json:"title,omitempty"`
	User               *User               `json:"user,omitempty"`
	State              *string             `json:"state,omitempty"`
	Locked             *bool               `json:"locked,omitempty"`
	Comments           *int                `json:"comments,omitempty"`
	CreatedAt          *Timestamp          `json:"created_at,omitempty"`
	UpdatedAt          *Timestamp          `json:"updated_at,omitempty"`

	AuthorAssociation *string `json:"author_association,omitempty"`
	ActiveLockReason  *string `json:"active_lock_reason,omitempty"`
	Body              *string `json:"body,omitempty"`
}

type DiscussionCategory struct {
	ID           *int64     `json:"id,omitempty"`
	NodeID       *string    `json:"node_id,omitempty"`
	RepositoryID *int64     `json:"repository_id,omitempty"`
	Emoji        *string    `json:"emoji,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Description  *string    `json:"description,omitempty"`
	CreatedAt    *Timestamp `json:"created_at,omitempty"`
	UpdatedAt    *Timestamp `json:"updated_at,omitempty"`
	Slug         *string    `json:"slug,omitempty"`
	IsAnswerable *bool      `json:"is_answerable,omitempty"`
}

type ForkEvent struct {
	Forkee *Repository `json:"forkee,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type GitHubAppAuthorizationEvent struct {
	Action *string `json:"action,omitempty"`

	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type Page struct {
	PageName *string `json:"page_name,omitempty"`
	Title    *string `json:"title,omitempty"`
	Summary  *string `json:"summary,omitempty"`
	Action   *string `json:"action,omitempty"`
	SHA      *string `json:"sha,omitempty"`
	HTMLURL  *string `json:"html_url,omitempty"`
}

type GollumEvent struct {
	Pages []*Page `json:"pages,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type EditChange struct {
	Title         *EditTitle         `json:"title,omitempty"`
	Body          *EditBody          `json:"body,omitempty"`
	Base          *EditBase          `json:"base,omitempty"`
	Repo          *EditRepo          `json:"repository,omitempty"`
	Owner         *EditOwner         `json:"owner,omitempty"`
	DefaultBranch *EditDefaultBranch `json:"default_branch,omitempty"`
	Topics        *EditTopics        `json:"topics,omitempty"`
}

type EditTitle struct {
	From *string `json:"from,omitempty"`
}

type EditBody struct {
	From *string `json:"from,omitempty"`
}

type EditBase struct {
	Ref *EditRef `json:"ref,omitempty"`
	SHA *EditSHA `json:"sha,omitempty"`
}

type EditRef struct {
	From *string `json:"from,omitempty"`
}

type EditRepo struct {
	Name *RepoName `json:"name,omitempty"`
}

type EditOwner struct {
	OwnerInfo *OwnerInfo `json:"from,omitempty"`
}

type OwnerInfo struct {
	User *User `json:"user,omitempty"`
	Org  *User `json:"organization,omitempty"`
}

type RepoName struct {
	From *string `json:"from,omitempty"`
}

type EditTopics struct {
	From []string `json:"from,omitempty"`
}

type EditSHA struct {
	From *string `json:"from,omitempty"`
}

type EditDefaultBranch struct {
	From *string `json:"from,omitempty"`
}

type ProjectChange struct {
	Name *ProjectName `json:"name,omitempty"`
	Body *ProjectBody `json:"body,omitempty"`
}

type ProjectName struct {
	From *string `json:"from,omitempty"`
}

type ProjectBody struct {
	From *string `json:"from,omitempty"`
}

type ProjectCardChange struct {
	Note *ProjectCardNote `json:"note,omitempty"`
}

type ProjectCardNote struct {
	From *string `json:"from,omitempty"`
}

type ProjectColumnChange struct {
	Name *ProjectColumnName `json:"name,omitempty"`
}

type ProjectColumnName struct {
	From *string `json:"from,omitempty"`
}

type TeamChange struct {
	Description *TeamDescription `json:"description,omitempty"`
	Name        *TeamName        `json:"name,omitempty"`
	Privacy     *TeamPrivacy     `json:"privacy,omitempty"`
	Repository  *TeamRepository  `json:"repository,omitempty"`
}

type TeamDescription struct {
	From *string `json:"from,omitempty"`
}

type TeamName struct {
	From *string `json:"from,omitempty"`
}

type TeamPrivacy struct {
	From *string `json:"from,omitempty"`
}

type TeamRepository struct {
	Permissions *TeamPermissions `json:"permissions,omitempty"`
}

type TeamPermissions struct {
	From *TeamPermissionsFrom `json:"from,omitempty"`
}

type TeamPermissionsFrom struct {
	Admin *bool `json:"admin,omitempty"`
	Pull  *bool `json:"pull,omitempty"`
	Push  *bool `json:"push,omitempty"`
}

type InstallationEvent struct {
	Action       *string       `json:"action,omitempty"`
	Repositories []*Repository `json:"repositories,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
	Requester    *User         `json:"requester,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type InstallationRepositoriesEvent struct {
	Action              *string       `json:"action,omitempty"`
	RepositoriesAdded   []*Repository `json:"repositories_added,omitempty"`
	RepositoriesRemoved []*Repository `json:"repositories_removed,omitempty"`
	RepositorySelection *string       `json:"repository_selection,omitempty"`
	Sender              *User         `json:"sender,omitempty"`
	Installation        *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type InstallationLoginChange struct {
	From *string `json:"from,omitempty"`
}

type InstallationSlugChange struct {
	From *string `json:"from,omitempty"`
}

type InstallationChanges struct {
	Login *InstallationLoginChange `json:"login,omitempty"`
	Slug  *InstallationSlugChange  `json:"slug,omitempty"`
}

type InstallationTargetEvent struct {
	Account      *User                `json:"account,omitempty"`
	Action       *string              `json:"action,omitempty"`
	Changes      *InstallationChanges `json:"changes,omitempty"`
	Enterprise   *Enterprise          `json:"enterprise,omitempty"`
	Installation *Installation        `json:"installation,omitempty"`
	Organization *Organization        `json:"organization,omitempty"`
	Repository   *Repository          `json:"repository,omitempty"`
	Sender       *User                `json:"sender,omitempty"`
	TargetType   *string              `json:"target_type,omitempty"`
}

type IssueCommentEvent struct {
	Action  *string       `json:"action,omitempty"`
	Issue   *Issue        `json:"issue,omitempty"`
	Comment *IssueComment `json:"comment,omitempty"`

	Changes      *EditChange   `json:"changes,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Organization *Organization `json:"organization,omitempty"`
}

type IssuesEvent struct {
	Action   *string `json:"action,omitempty"`
	Issue    *Issue  `json:"issue,omitempty"`
	Assignee *User   `json:"assignee,omitempty"`
	Label    *Label  `json:"label,omitempty"`

	Changes      *EditChange   `json:"changes,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
	Milestone    *Milestone    `json:"milestone,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type LabelEvent struct {
	Action  *string     `json:"action,omitempty"`
	Label   *Label      `json:"label,omitempty"`
	Changes *EditChange `json:"changes,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type MarketplacePurchaseEvent struct {
	Action *string `json:"action,omitempty"`

	EffectiveDate               *Timestamp           `json:"effective_date,omitempty"`
	MarketplacePurchase         *MarketplacePurchase `json:"marketplace_purchase,omitempty"`
	PreviousMarketplacePurchase *MarketplacePurchase `json:"previous_marketplace_purchase,omitempty"`
	Sender                      *User                `json:"sender,omitempty"`
	Installation                *Installation        `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type MemberChangesPermission struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
}

type MemberChangesRoleName struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
}

type MemberChanges struct {
	Permission *MemberChangesPermission `json:"permission,omitempty"`
	RoleName   *MemberChangesRoleName   `json:"role_name,omitempty"`
}

type MemberEvent struct {
	Action  *string        `json:"action,omitempty"`
	Member  *User          `json:"member,omitempty"`
	Changes *MemberChanges `json:"changes,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type MembershipEvent struct {
	Action *string `json:"action,omitempty"`

	Scope  *string `json:"scope,omitempty"`
	Member *User   `json:"member,omitempty"`
	Team   *Team   `json:"team,omitempty"`

	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type MergeGroup struct {
	HeadSHA *string `json:"head_sha,omitempty"`

	HeadRef *string `json:"head_ref,omitempty"`

	BaseSHA *string `json:"base_sha,omitempty"`

	BaseRef *string `json:"base_ref,omitempty"`

	HeadCommit *Commit `json:"head_commit,omitempty"`
}

type MergeGroupEvent struct {
	Action *string `json:"action,omitempty"`

	Reason *string `json:"reason,omitempty"`

	MergeGroup *MergeGroup `json:"merge_group,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
}

type MetaEvent struct {
	Action *string `json:"action,omitempty"`

	HookID *int64 `json:"hook_id,omitempty"`

	Hook *Hook `json:"hook,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type MilestoneEvent struct {
	Action    *string    `json:"action,omitempty"`
	Milestone *Milestone `json:"milestone,omitempty"`

	Changes      *EditChange   `json:"changes,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type OrganizationEvent struct {
	Action *string `json:"action,omitempty"`

	Invitation *Invitation `json:"invitation,omitempty"`

	Membership *Membership `json:"membership,omitempty"`

	Organization *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type OrgBlockEvent struct {
	Action       *string       `json:"action,omitempty"`
	BlockedUser  *User         `json:"blocked_user,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`

	Installation *Installation `json:"installation,omitempty"`
}

type PackageEvent struct {
	Action  *string       `json:"action,omitempty"`
	Package *Package      `json:"package,omitempty"`
	Repo    *Repository   `json:"repository,omitempty"`
	Org     *Organization `json:"organization,omitempty"`
	Sender  *User         `json:"sender,omitempty"`

	Installation *Installation `json:"installation,omitempty"`
}

type PageBuildEvent struct {
	Build *PagesBuild `json:"build,omitempty"`

	ID           *int64        `json:"id,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type PersonalAccessTokenRequestEvent struct {
	Action                     *string                     `json:"action,omitempty"`
	PersonalAccessTokenRequest *PersonalAccessTokenRequest `json:"personal_access_token_request,omitempty"`
	Org                        *Organization               `json:"organization,omitempty"`
	Sender                     *User                       `json:"sender,omitempty"`
	Installation               *Installation               `json:"installation,omitempty"`
}

type PersonalAccessTokenRequest struct {
	ID    *int64 `json:"id,omitempty"`
	Owner *User  `json:"owner,omitempty"`

	PermissionsAdded *PersonalAccessTokenPermissions `json:"permissions_added,omitempty"`

	PermissionsUpgraded *PersonalAccessTokenPermissions `json:"permissions_upgraded,omitempty"`

	PermissionsResult *PersonalAccessTokenPermissions `json:"permissions_result,omitempty"`

	RepositorySelection *string `json:"repository_selection,omitempty"`

	RepositoryCount *int64 `json:"repository_count,omitempty"`

	Repositories []*Repository `json:"repositories,omitempty"`

	CreatedAt *Timestamp `json:"created_at,omitempty"`

	TokenExpired *bool `json:"token_expired,omitempty"`

	TokenExpiresAt *Timestamp `json:"token_expires_at,omitempty"`

	TokenLastUsedAt *Timestamp `json:"token_last_used_at,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type PersonalAccessTokenPermissions struct {
	Org   map[string]string `json:"organization,omitempty"`
	Repo  map[string]string `json:"repository,omitempty"`
	Other map[string]string `json:"other,omitempty"`
}

type PingEvent struct {
	Zen *string `json:"zen,omitempty"`

	HookID *int64 `json:"hook_id,omitempty"`

	Hook *Hook `json:"hook,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type ProjectV2Event struct {
	Action     *string    `json:"action,omitempty"`
	ProjectsV2 *ProjectV2 `json:"projects_v2,omitempty"`

	Installation *Installation `json:"installation,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
}

type ProjectV2ItemEvent struct {
	Action        *string              `json:"action,omitempty"`
	Changes       *ProjectV2ItemChange `json:"changes,omitempty"`
	ProjectV2Item *ProjectV2Item       `json:"projects_v2_item,omitempty"`

	Installation *Installation `json:"installation,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
}

type ProjectV2ItemChange struct {
	ArchivedAt *ArchivedAt `json:"archived_at,omitempty"`
	FieldValue *FieldValue `json:"field_value,omitempty"`
}

type ArchivedAt struct {
	From *Timestamp `json:"from,omitempty"`
	To   *Timestamp `json:"to,omitempty"`
}

type FieldValue struct {
	FieldNodeID   *string         `json:"field_node_id,omitempty"`
	FieldType     *string         `json:"field_type,omitempty"`
	FieldName     *string         `json:"field_name,omitempty"`
	ProjectNumber *int64          `json:"project_number,omitempty"`
	From          json.RawMessage `json:"from,omitempty"`
	To            json.RawMessage `json:"to,omitempty"`
}

type PublicEvent struct {
	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type PullRequestEvent struct {
	Action      *string      `json:"action,omitempty"`
	Assignee    *User        `json:"assignee,omitempty"`
	Number      *int         `json:"number,omitempty"`
	PullRequest *PullRequest `json:"pull_request,omitempty"`

	Changes *EditChange `json:"changes,omitempty"`

	RequestedReviewer *User `json:"requested_reviewer,omitempty"`

	RequestedTeam *Team         `json:"requested_team,omitempty"`
	Repo          *Repository   `json:"repository,omitempty"`
	Sender        *User         `json:"sender,omitempty"`
	Installation  *Installation `json:"installation,omitempty"`
	Label         *Label        `json:"label,omitempty"`
	Reason        *string       `json:"reason,omitempty"`

	Organization *Organization `json:"organization,omitempty"`

	Before *string `json:"before,omitempty"`
	After  *string `json:"after,omitempty"`

	PerformedViaGithubApp *App `json:"performed_via_github_app,omitempty"`
}

type PullRequestReviewEvent struct {
	Action      *string            `json:"action,omitempty"`
	Review      *PullRequestReview `json:"review,omitempty"`
	PullRequest *PullRequest       `json:"pull_request,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Organization *Organization `json:"organization,omitempty"`
}

type PullRequestReviewCommentEvent struct {
	Action      *string             `json:"action,omitempty"`
	PullRequest *PullRequest        `json:"pull_request,omitempty"`
	Comment     *PullRequestComment `json:"comment,omitempty"`

	Changes      *EditChange   `json:"changes,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type PullRequestReviewThreadEvent struct {
	Action      *string            `json:"action,omitempty"`
	Thread      *PullRequestThread `json:"thread,omitempty"`
	PullRequest *PullRequest       `json:"pull_request,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type PullRequestTargetEvent struct {
	Action      *string      `json:"action,omitempty"`
	Assignee    *User        `json:"assignee,omitempty"`
	Number      *int         `json:"number,omitempty"`
	PullRequest *PullRequest `json:"pull_request,omitempty"`

	Changes *EditChange `json:"changes,omitempty"`

	RequestedReviewer *User `json:"requested_reviewer,omitempty"`

	RequestedTeam *Team         `json:"requested_team,omitempty"`
	Repo          *Repository   `json:"repository,omitempty"`
	Sender        *User         `json:"sender,omitempty"`
	Installation  *Installation `json:"installation,omitempty"`
	Label         *Label        `json:"label,omitempty"`

	Organization *Organization `json:"organization,omitempty"`

	Before *string `json:"before,omitempty"`
	After  *string `json:"after,omitempty"`

	PerformedViaGithubApp *App `json:"performed_via_github_app,omitempty"`
}

type PushEvent struct {
	PushID *int64  `json:"push_id,omitempty"`
	Head   *string `json:"head,omitempty"`
	Ref    *string `json:"ref,omitempty"`

	Size *int `json:"size,omitempty"`

	Commits []*HeadCommit `json:"commits,omitempty"`
	Before  *string       `json:"before,omitempty"`

	DistinctSize *int `json:"distinct_size,omitempty"`

	Action       *string              `json:"action,omitempty"`
	After        *string              `json:"after,omitempty"`
	Created      *bool                `json:"created,omitempty"`
	Deleted      *bool                `json:"deleted,omitempty"`
	Forced       *bool                `json:"forced,omitempty"`
	BaseRef      *string              `json:"base_ref,omitempty"`
	Compare      *string              `json:"compare,omitempty"`
	Repo         *PushEventRepository `json:"repository,omitempty"`
	HeadCommit   *HeadCommit          `json:"head_commit,omitempty"`
	Pusher       *CommitAuthor        `json:"pusher,omitempty"`
	Sender       *User                `json:"sender,omitempty"`
	Installation *Installation        `json:"installation,omitempty"`

	Organization *Organization `json:"organization,omitempty"`
}

func (p PushEvent) String() string { _ = "STUB: not implemented"; return "" }

type HeadCommit struct {
	Message  *string       `json:"message,omitempty"`
	Author   *CommitAuthor `json:"author,omitempty"`
	URL      *string       `json:"url,omitempty"`
	Distinct *bool         `json:"distinct,omitempty"`

	SHA *string `json:"sha,omitempty"`

	ID        *string       `json:"id,omitempty"`
	TreeID    *string       `json:"tree_id,omitempty"`
	Timestamp *Timestamp    `json:"timestamp,omitempty"`
	Committer *CommitAuthor `json:"committer,omitempty"`
	Added     []string      `json:"added,omitempty"`
	Removed   []string      `json:"removed,omitempty"`
	Modified  []string      `json:"modified,omitempty"`
}

func (h HeadCommit) String() string { _ = "STUB: not implemented"; return "" }

type PushEventRepository struct {
	ID               *int64         `json:"id,omitempty"`
	NodeID           *string        `json:"node_id,omitempty"`
	Name             *string        `json:"name,omitempty"`
	FullName         *string        `json:"full_name,omitempty"`
	Owner            *User          `json:"owner,omitempty"`
	Private          *bool          `json:"private,omitempty"`
	Description      *string        `json:"description,omitempty"`
	Fork             *bool          `json:"fork,omitempty"`
	CreatedAt        *Timestamp     `json:"created_at,omitempty"`
	PushedAt         *Timestamp     `json:"pushed_at,omitempty"`
	UpdatedAt        *Timestamp     `json:"updated_at,omitempty"`
	Homepage         *string        `json:"homepage,omitempty"`
	PullsURL         *string        `json:"pulls_url,omitempty"`
	Size             *int           `json:"size,omitempty"`
	StargazersCount  *int           `json:"stargazers_count,omitempty"`
	WatchersCount    *int           `json:"watchers_count,omitempty"`
	Language         *string        `json:"language,omitempty"`
	HasIssues        *bool          `json:"has_issues,omitempty"`
	HasDownloads     *bool          `json:"has_downloads,omitempty"`
	HasWiki          *bool          `json:"has_wiki,omitempty"`
	HasPages         *bool          `json:"has_pages,omitempty"`
	ForksCount       *int           `json:"forks_count,omitempty"`
	Archived         *bool          `json:"archived,omitempty"`
	Disabled         *bool          `json:"disabled,omitempty"`
	OpenIssuesCount  *int           `json:"open_issues_count,omitempty"`
	DefaultBranch    *string        `json:"default_branch,omitempty"`
	MasterBranch     *string        `json:"master_branch,omitempty"`
	Organization     *string        `json:"organization,omitempty"`
	URL              *string        `json:"url,omitempty"`
	ArchiveURL       *string        `json:"archive_url,omitempty"`
	HTMLURL          *string        `json:"html_url,omitempty"`
	StatusesURL      *string        `json:"statuses_url,omitempty"`
	GitURL           *string        `json:"git_url,omitempty"`
	SSHURL           *string        `json:"ssh_url,omitempty"`
	CloneURL         *string        `json:"clone_url,omitempty"`
	SVNURL           *string        `json:"svn_url,omitempty"`
	Topics           []string       `json:"topics,omitempty"`
	CustomProperties map[string]any `json:"custom_properties,omitempty"`
}

type PushEventRepoOwner struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type RegistryPackageEvent struct {
	Action          *string       `json:"action,omitempty"`
	RegistryPackage *Package      `json:"registry_package,omitempty"`
	Repository      *Repository   `json:"repository,omitempty"`
	Organization    *Organization `json:"organization,omitempty"`
	Enterprise      *Enterprise   `json:"enterprise,omitempty"`
	Sender          *User         `json:"sender,omitempty"`

	Installation *Installation `json:"installation,omitempty"`
}

type ReleaseEvent struct {
	Action  *string            `json:"action,omitempty"`
	Release *RepositoryRelease `json:"release,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type RepositoryEvent struct {
	Action *string     `json:"action,omitempty"`
	Repo   *Repository `json:"repository,omitempty"`

	Changes      *EditChange   `json:"changes,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type RepositoryDispatchEvent struct {
	Action        *string         `json:"action,omitempty"`
	Branch        *string         `json:"branch,omitempty"`
	ClientPayload json.RawMessage `json:"client_payload,omitempty"`
	Repo          *Repository     `json:"repository,omitempty"`

	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type RepositoryImportEvent struct {
	Status *string       `json:"status,omitempty"`
	Repo   *Repository   `json:"repository,omitempty"`
	Org    *Organization `json:"organization,omitempty"`
	Sender *User         `json:"sender,omitempty"`
}

type RepositoryRulesetEvent struct {
	Action            *string                   `json:"action,omitempty"`
	Enterprise        *Enterprise               `json:"enterprise,omitempty"`
	Installation      *Installation             `json:"installation,omitempty"`
	Organization      *Organization             `json:"organization,omitempty"`
	Repository        *Repository               `json:"repository,omitempty"`
	RepositoryRuleset *RepositoryRuleset        `json:"repository_ruleset"`
	Changes           *RepositoryRulesetChanges `json:"changes,omitempty"`
	Sender            *User                     `json:"sender"`
}

type RepositoryRulesetChanges struct {
	Name        *RepositoryRulesetChangeSource      `json:"name,omitempty"`
	Enforcement *RepositoryRulesetChangeSource      `json:"enforcement,omitempty"`
	Conditions  *RepositoryRulesetChangedConditions `json:"conditions,omitempty"`
	Rules       *RepositoryRulesetChangedRules      `json:"rules,omitempty"`
}

type RepositoryRulesetChangeSource struct {
	From *string `json:"from,omitempty"`
}

type RepositoryRulesetChangeSources struct {
	From []string `json:"from,omitempty"`
}

type RepositoryRulesetChangedConditions struct {
	Added   []*RepositoryRulesetConditions        `json:"added,omitempty"`
	Deleted []*RepositoryRulesetConditions        `json:"deleted,omitempty"`
	Updated []*RepositoryRulesetUpdatedConditions `json:"updated,omitempty"`
}

type RepositoryRulesetUpdatedConditions struct {
	Condition *RepositoryRulesetConditions       `json:"condition,omitempty"`
	Changes   *RepositoryRulesetUpdatedCondition `json:"changes,omitempty"`
}

type RepositoryRulesetUpdatedCondition struct {
	ConditionType *RepositoryRulesetChangeSource  `json:"condition_type,omitempty"`
	Target        *RepositoryRulesetChangeSource  `json:"target,omitempty"`
	Include       *RepositoryRulesetChangeSources `json:"include,omitempty"`
	Exclude       *RepositoryRulesetChangeSources `json:"exclude,omitempty"`
}

type RepositoryRulesetChangedRules struct {
	Added   []*RepositoryRule                `json:"added,omitempty"`
	Deleted []*RepositoryRule                `json:"deleted,omitempty"`
	Updated []*RepositoryRulesetUpdatedRules `json:"updated,omitempty"`
}

type RepositoryRulesetUpdatedRules struct {
	Rule    *RepositoryRule               `json:"rule,omitempty"`
	Changes *RepositoryRulesetChangedRule `json:"changes,omitempty"`
}

type RepositoryRulesetChangedRule struct {
	Configuration *RepositoryRulesetChangeSource `json:"configuration,omitempty"`
	RuleType      *RepositoryRulesetChangeSource `json:"rule_type,omitempty"`
	Pattern       *RepositoryRulesetChangeSource `json:"pattern,omitempty"`
}

type RepositoryVulnerabilityAlertEvent struct {
	Action *string `json:"action,omitempty"`

	Alert *RepositoryVulnerabilityAlert `json:"alert,omitempty"`

	Repository *Repository `json:"repository,omitempty"`

	Installation *Installation `json:"installation,omitempty"`

	Sender *User `json:"sender,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type RepositoryVulnerabilityAlert struct {
	ID                       *int64     `json:"id,omitempty"`
	AffectedRange            *string    `json:"affected_range,omitempty"`
	AffectedPackageName      *string    `json:"affected_package_name,omitempty"`
	ExternalReference        *string    `json:"external_reference,omitempty"`
	ExternalIdentifier       *string    `json:"external_identifier,omitempty"`
	GitHubSecurityAdvisoryID *string    `json:"ghsa_id,omitempty"`
	Severity                 *string    `json:"severity,omitempty"`
	CreatedAt                *Timestamp `json:"created_at,omitempty"`
	FixedIn                  *string    `json:"fixed_in,omitempty"`
	Dismisser                *User      `json:"dismisser,omitempty"`
	DismissReason            *string    `json:"dismiss_reason,omitempty"`
	DismissedAt              *Timestamp `json:"dismissed_at,omitempty"`
}

type SecretScanningAlertEvent struct {
	Action *string `json:"action,omitempty"`

	Alert *SecretScanningAlert `json:"alert,omitempty"`

	Sender *User `json:"sender,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	Enterprise   *Enterprise   `json:"enterprise,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type SecretScanningAlertLocationEvent struct {
	Action       *string                      `json:"action,omitempty"`
	Alert        *SecretScanningAlert         `json:"alert,omitempty"`
	Installation *Installation                `json:"installation,omitempty"`
	Location     *SecretScanningAlertLocation `json:"location,omitempty"`
	Organization *Organization                `json:"organization,omitempty"`
	Repo         *Repository                  `json:"repository,omitempty"`
	Sender       *User                        `json:"sender,omitempty"`
}

type SecurityAndAnalysisEvent struct {
	Changes      *SecurityAndAnalysisChange `json:"changes,omitempty"`
	Enterprise   *Enterprise                `json:"enterprise,omitempty"`
	Installation *Installation              `json:"installation,omitempty"`
	Organization *Organization              `json:"organization,omitempty"`
	Repository   *Repository                `json:"repository,omitempty"`
	Sender       *User                      `json:"sender,omitempty"`
}

type SecurityAndAnalysisChange struct {
	From *SecurityAndAnalysisChangeFrom `json:"from,omitempty"`
}

type SecurityAndAnalysisChangeFrom struct {
	SecurityAndAnalysis *SecurityAndAnalysis `json:"security_and_analysis,omitempty"`
}

type StarEvent struct {
	Action *string `json:"action,omitempty"`

	StarredAt *Timestamp `json:"starred_at,omitempty"`

	Org          *Organization `json:"organization,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type StatusEvent struct {
	SHA *string `json:"sha,omitempty"`

	State       *string   `json:"state,omitempty"`
	Description *string   `json:"description,omitempty"`
	TargetURL   *string   `json:"target_url,omitempty"`
	Branches    []*Branch `json:"branches,omitempty"`

	ID           *int64            `json:"id,omitempty"`
	Name         *string           `json:"name,omitempty"`
	Context      *string           `json:"context,omitempty"`
	Commit       *RepositoryCommit `json:"commit,omitempty"`
	CreatedAt    *Timestamp        `json:"created_at,omitempty"`
	UpdatedAt    *Timestamp        `json:"updated_at,omitempty"`
	Repo         *Repository       `json:"repository,omitempty"`
	Sender       *User             `json:"sender,omitempty"`
	Installation *Installation     `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type TeamEvent struct {
	Action  *string     `json:"action,omitempty"`
	Team    *Team       `json:"team,omitempty"`
	Changes *TeamChange `json:"changes,omitempty"`
	Repo    *Repository `json:"repository,omitempty"`

	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type TeamAddEvent struct {
	Team *Team       `json:"team,omitempty"`
	Repo *Repository `json:"repository,omitempty"`

	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type UserEvent struct {
	User *User `json:"user,omitempty"`

	Action     *string     `json:"action,omitempty"`
	Enterprise *Enterprise `json:"enterprise,omitempty"`
	Sender     *User       `json:"sender,omitempty"`

	Installation *Installation `json:"installation,omitempty"`
}

type WatchEvent struct {
	Action *string `json:"action,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`

	Org *Organization `json:"organization,omitempty"`
}

type WorkflowDispatchEvent struct {
	Inputs   json.RawMessage `json:"inputs,omitempty"`
	Ref      *string         `json:"ref,omitempty"`
	Workflow *string         `json:"workflow,omitempty"`

	Repo         *Repository   `json:"repository,omitempty"`
	Org          *Organization `json:"organization,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type WorkflowJobEvent struct {
	WorkflowJob *WorkflowJob `json:"workflow_job,omitempty"`

	Action *string `json:"action,omitempty"`

	Org          *Organization `json:"organization,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
	Deployment   *Deployment   `json:"deployment,omitempty"`
}

type WorkflowRunEvent struct {
	Action      *string      `json:"action,omitempty"`
	Workflow    *Workflow    `json:"workflow,omitempty"`
	WorkflowRun *WorkflowRun `json:"workflow_run,omitempty"`

	Org          *Organization `json:"organization,omitempty"`
	Repo         *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
}

type SecurityAdvisory struct {
	CVSS               *AdvisoryCVSS                 `json:"cvss,omitempty"`
	CWEs               []*AdvisoryCWEs               `json:"cwes,omitempty"`
	GHSAID             *string                       `json:"ghsa_id,omitempty"`
	Summary            *string                       `json:"summary,omitempty"`
	Description        *string                       `json:"description,omitempty"`
	Severity           *string                       `json:"severity,omitempty"`
	Identifiers        []*AdvisoryIdentifier         `json:"identifiers,omitempty"`
	References         []*AdvisoryReference          `json:"references,omitempty"`
	PublishedAt        *Timestamp                    `json:"published_at,omitempty"`
	UpdatedAt          *Timestamp                    `json:"updated_at,omitempty"`
	WithdrawnAt        *Timestamp                    `json:"withdrawn_at,omitempty"`
	Vulnerabilities    []*AdvisoryVulnerability      `json:"vulnerabilities,omitempty"`
	CVEID              *string                       `json:"cve_id,omitempty"`
	URL                *string                       `json:"url,omitempty"`
	HTMLURL            *string                       `json:"html_url,omitempty"`
	Author             *User                         `json:"author,omitempty"`
	Publisher          *User                         `json:"publisher,omitempty"`
	State              *string                       `json:"state,omitempty"`
	CreatedAt          *Timestamp                    `json:"created_at,omitempty"`
	ClosedAt           *Timestamp                    `json:"closed_at,omitempty"`
	Submission         *SecurityAdvisorySubmission   `json:"submission,omitempty"`
	CWEIDs             []string                      `json:"cwe_ids,omitempty"`
	Credits            []*RepoAdvisoryCredit         `json:"credits,omitempty"`
	CreditsDetailed    []*RepoAdvisoryCreditDetailed `json:"credits_detailed,omitempty"`
	CollaboratingUsers []*User                       `json:"collaborating_users,omitempty"`
	CollaboratingTeams []*Team                       `json:"collaborating_teams,omitempty"`
	PrivateFork        *Repository                   `json:"private_fork,omitempty"`
}

type AdvisoryIdentifier struct {
	Value *string `json:"value,omitempty"`
	Type  *string `json:"type,omitempty"`
}

type AdvisoryReference struct {
	URL *string `json:"url,omitempty"`
}

type AdvisoryVulnerability struct {
	Package                *VulnerabilityPackage `json:"package,omitempty"`
	Severity               *string               `json:"severity,omitempty"`
	VulnerableVersionRange *string               `json:"vulnerable_version_range,omitempty"`
	FirstPatchedVersion    *FirstPatchedVersion  `json:"first_patched_version,omitempty"`

	PatchedVersions     *string  `json:"patched_versions,omitempty"`
	VulnerableFunctions []string `json:"vulnerable_functions,omitempty"`
}

type VulnerabilityPackage struct {
	Ecosystem *string `json:"ecosystem,omitempty"`
	Name      *string `json:"name,omitempty"`
}

type FirstPatchedVersion struct {
	Identifier *string `json:"identifier,omitempty"`
}

type SecurityAdvisoryEvent struct {
	Action           *string           `json:"action,omitempty"`
	SecurityAdvisory *SecurityAdvisory `json:"security_advisory,omitempty"`

	Enterprise   *Enterprise   `json:"enterprise,omitempty"`
	Installation *Installation `json:"installation,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	Repository   *Repository   `json:"repository,omitempty"`
	Sender       *User         `json:"sender,omitempty"`
}

type CodeScanningAlertEvent struct {
	Action *string `json:"action,omitempty"`
	Alert  *Alert  `json:"alert,omitempty"`
	Ref    *string `json:"ref,omitempty"`

	CommitOID *string       `json:"commit_oid,omitempty"`
	Repo      *Repository   `json:"repository,omitempty"`
	Org       *Organization `json:"organization,omitempty"`
	Sender    *User         `json:"sender,omitempty"`

	Installation *Installation `json:"installation,omitempty"`
}

type SponsorshipEvent struct {
	Action        *string             `json:"action,omitempty"`
	EffectiveDate *string             `json:"effective_date,omitempty"`
	Changes       *SponsorshipChanges `json:"changes,omitempty"`
	Repository    *Repository         `json:"repository,omitempty"`
	Organization  *Organization       `json:"organization,omitempty"`
	Sender        *User               `json:"sender,omitempty"`
	Installation  *Installation       `json:"installation,omitempty"`
}

type SponsorshipChanges struct {
	Tier         *SponsorshipTier `json:"tier,omitempty"`
	PrivacyLevel *string          `json:"privacy_level,omitempty"`
}

type SponsorshipTier struct {
	From *string `json:"from,omitempty"`
}
