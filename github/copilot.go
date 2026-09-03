package github

import (
	"context"
	"io"
	"net/http"
	"time"
)

type CopilotService service

type CopilotOrganizationDetails struct {
	SeatBreakdown         *CopilotSeatBreakdown `json:"seat_breakdown"`
	PublicCodeSuggestions string                `json:"public_code_suggestions"`
	CopilotChat           string                `json:"copilot_chat"`
	SeatManagementSetting string                `json:"seat_management_setting"`
}

type CopilotSeatBreakdown struct {
	Total               int `json:"total"`
	AddedThisCycle      int `json:"added_this_cycle"`
	PendingCancellation int `json:"pending_cancellation"`
	PendingInvitation   int `json:"pending_invitation"`
	ActiveThisCycle     int `json:"active_this_cycle"`
	InactiveThisCycle   int `json:"inactive_this_cycle"`
}

type ListCopilotSeatsResponse struct {
	TotalSeats int64                 `json:"total_seats"`
	Seats      []*CopilotSeatDetails `json:"seats"`
}

type CopilotSeatDetails struct {
	Assignee                any        `json:"assignee"`
	AssigningTeam           *Team      `json:"assigning_team,omitempty"`
	PendingCancellationDate *string    `json:"pending_cancellation_date,omitempty"`
	LastActivityAt          *Timestamp `json:"last_activity_at,omitempty"`
	LastActivityEditor      *string    `json:"last_activity_editor,omitempty"`
	CreatedAt               *Timestamp `json:"created_at"`
	UpdatedAt               *Timestamp `json:"updated_at,omitempty"`
	PlanType                *string    `json:"plan_type,omitempty"`
}

type SeatAssignments struct {
	SeatsCreated int `json:"seats_created"`
}

type SeatCancellations struct {
	SeatsCancelled int `json:"seats_cancelled"`
}

type CopilotMetricsListOptions struct {
	Since *time.Time `url:"since,omitempty"`
	Until *time.Time `url:"until,omitempty"`

	ListOptions
}

type CopilotIDECodeCompletionsLanguage struct {
	Name              string `json:"name"`
	TotalEngagedUsers int    `json:"total_engaged_users"`
}

type CopilotIDECodeCompletionsModelLanguage struct {
	Name                    string `json:"name"`
	TotalEngagedUsers       int    `json:"total_engaged_users"`
	TotalCodeSuggestions    int    `json:"total_code_suggestions"`
	TotalCodeAcceptances    int    `json:"total_code_acceptances"`
	TotalCodeLinesSuggested int    `json:"total_code_lines_suggested"`
	TotalCodeLinesAccepted  int    `json:"total_code_lines_accepted"`
}

type CopilotIDECodeCompletionsModel struct {
	Name                    string                                    `json:"name"`
	IsCustomModel           bool                                      `json:"is_custom_model"`
	CustomModelTrainingDate *string                                   `json:"custom_model_training_date,omitempty"`
	TotalEngagedUsers       int                                       `json:"total_engaged_users"`
	Languages               []*CopilotIDECodeCompletionsModelLanguage `json:"languages"`
}

type CopilotIDECodeCompletionsEditor struct {
	Name              string                            `json:"name"`
	TotalEngagedUsers int                               `json:"total_engaged_users"`
	Models            []*CopilotIDECodeCompletionsModel `json:"models"`
}

type CopilotIDECodeCompletions struct {
	TotalEngagedUsers int                                  `json:"total_engaged_users"`
	Languages         []*CopilotIDECodeCompletionsLanguage `json:"languages"`
	Editors           []*CopilotIDECodeCompletionsEditor   `json:"editors"`
}

type CopilotIDEChatModel struct {
	Name                     string  `json:"name"`
	IsCustomModel            bool    `json:"is_custom_model"`
	CustomModelTrainingDate  *string `json:"custom_model_training_date,omitempty"`
	TotalEngagedUsers        int     `json:"total_engaged_users"`
	TotalChats               int     `json:"total_chats"`
	TotalChatInsertionEvents int     `json:"total_chat_insertion_events"`
	TotalChatCopyEvents      int     `json:"total_chat_copy_events"`
}

type CopilotIDEChatEditor struct {
	Name              string                 `json:"name"`
	TotalEngagedUsers int                    `json:"total_engaged_users"`
	Models            []*CopilotIDEChatModel `json:"models"`
}

type CopilotIDEChat struct {
	TotalEngagedUsers int                     `json:"total_engaged_users"`
	Editors           []*CopilotIDEChatEditor `json:"editors"`
}

type CopilotDotcomChatModel struct {
	Name                    string  `json:"name"`
	IsCustomModel           bool    `json:"is_custom_model"`
	CustomModelTrainingDate *string `json:"custom_model_training_date,omitempty"`
	TotalEngagedUsers       int     `json:"total_engaged_users"`
	TotalChats              int     `json:"total_chats"`
}

type CopilotDotcomChat struct {
	TotalEngagedUsers int                       `json:"total_engaged_users"`
	Models            []*CopilotDotcomChatModel `json:"models"`
}

type CopilotDotcomPullRequestsModel struct {
	Name                    string  `json:"name"`
	IsCustomModel           bool    `json:"is_custom_model"`
	CustomModelTrainingDate *string `json:"custom_model_training_date,omitempty"`
	TotalPRSummariesCreated int     `json:"total_pr_summaries_created"`
	TotalEngagedUsers       int     `json:"total_engaged_users"`
}

type CopilotDotcomPullRequestsRepository struct {
	Name              string                            `json:"name"`
	TotalEngagedUsers int                               `json:"total_engaged_users"`
	Models            []*CopilotDotcomPullRequestsModel `json:"models"`
}

type CopilotDotcomPullRequests struct {
	TotalEngagedUsers int                                    `json:"total_engaged_users"`
	Repositories      []*CopilotDotcomPullRequestsRepository `json:"repositories"`
}

type CopilotMetrics struct {
	Date                      string                     `json:"date"`
	TotalActiveUsers          *int                       `json:"total_active_users,omitempty"`
	TotalEngagedUsers         *int                       `json:"total_engaged_users,omitempty"`
	CopilotIDECodeCompletions *CopilotIDECodeCompletions `json:"copilot_ide_code_completions,omitempty"`
	CopilotIDEChat            *CopilotIDEChat            `json:"copilot_ide_chat,omitempty"`
	CopilotDotcomChat         *CopilotDotcomChat         `json:"copilot_dotcom_chat,omitempty"`
	CopilotDotcomPullRequests *CopilotDotcomPullRequests `json:"copilot_dotcom_pull_requests,omitempty"`
}

type CopilotMetricsReportOptions struct {
	Day string `url:"day"`
}

type CopilotDailyMetricsReport struct {
	DownloadLinks []string `json:"download_links"`
	ReportDay     string   `json:"report_day"`
}

type CopilotMetricsReport struct {
	DownloadLinks  []string `json:"download_links"`
	ReportStartDay string   `json:"report_start_day"`
	ReportEndDay   string   `json:"report_end_day"`
}

func (cp *CopilotSeatDetails) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (cp *CopilotSeatDetails) GetUser() (*User, bool) { _ = "STUB: not implemented"; return nil, false }

func (cp *CopilotSeatDetails) GetTeam() (*Team, bool) { _ = "STUB: not implemented"; return nil, false }

func (cp *CopilotSeatDetails) GetOrganization() (*Organization, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

//meta:operation GET /orgs/{org}/copilot/billing
func (s *CopilotService) GetCopilotBilling(ctx context.Context, org string) (*CopilotOrganizationDetails, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/copilot/billing/seats
func (s *CopilotService) ListCopilotSeats(ctx context.Context, org string, opts *ListOptions) (*ListCopilotSeatsResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/copilot/billing/seats
func (s *CopilotService) ListCopilotEnterpriseSeats(ctx context.Context, enterprise string, opts *ListOptions) (*ListCopilotSeatsResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ListOrganizationCopilotCodingAgentRepositoriesResponse struct {
	TotalCount   int           `json:"total_count"`
	Repositories []*Repository `json:"repositories"`
}

//meta:operation GET /orgs/{org}/copilot/coding-agent/permissions/repositories
func (s *CopilotService) ListOrganizationCodingAgentRepositories(ctx context.Context, org string, opts *ListOptions) (*ListOrganizationCopilotCodingAgentRepositoriesResponse, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type CopilotOrganizationContentExclusionDetails map[string][]string

//meta:operation GET /orgs/{org}/copilot/content_exclusion
func (s *CopilotService) GetOrganizationContentExclusionDetails(ctx context.Context, org string) (CopilotOrganizationContentExclusionDetails, *Response, error) {
	_ = "STUB: not implemented"
	return *new(CopilotOrganizationContentExclusionDetails), nil, nil
}

//meta:operation POST /orgs/{org}/copilot/billing/selected_teams
func (s *CopilotService) AddCopilotTeams(ctx context.Context, org string, teamNames []string) (*SeatAssignments, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/copilot/billing/selected_teams
func (s *CopilotService) RemoveCopilotTeams(ctx context.Context, org string, teamNames []string) (*SeatCancellations, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/copilot/billing/selected_users
func (s *CopilotService) AddCopilotUsers(ctx context.Context, org string, users []string) (*SeatAssignments, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/copilot/billing/selected_users
func (s *CopilotService) RemoveCopilotUsers(ctx context.Context, org string, users []string) (*SeatCancellations, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/members/{username}/copilot
func (s *CopilotService) GetSeatDetails(ctx context.Context, org, user string) (*CopilotSeatDetails, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/copilot/metrics
func (s *CopilotService) GetEnterpriseMetrics(ctx context.Context, enterprise string, opts *CopilotMetricsListOptions) ([]*CopilotMetrics, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/team/{team_slug}/copilot/metrics
func (s *CopilotService) GetEnterpriseTeamMetrics(ctx context.Context, enterprise, team string, opts *CopilotMetricsListOptions) ([]*CopilotMetrics, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/copilot/metrics
func (s *CopilotService) GetOrganizationMetrics(ctx context.Context, org string, opts *CopilotMetricsListOptions) ([]*CopilotMetrics, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/team/{team_slug}/copilot/metrics
func (s *CopilotService) GetOrganizationTeamMetrics(ctx context.Context, org, team string, opts *CopilotMetricsListOptions) ([]*CopilotMetrics, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/copilot/metrics/reports/enterprise-1-day
func (s *CopilotService) GetEnterpriseDailyMetricsReport(ctx context.Context, enterprise string, opts *CopilotMetricsReportOptions) (*CopilotDailyMetricsReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/copilot/metrics/reports/enterprise-28-day/latest
func (s *CopilotService) GetEnterpriseMetricsReport(ctx context.Context, enterprise string) (*CopilotMetricsReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/copilot/metrics/reports/users-1-day
func (s *CopilotService) GetEnterpriseUsersDailyMetricsReport(ctx context.Context, enterprise string, opts *CopilotMetricsReportOptions) (*CopilotDailyMetricsReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /enterprises/{enterprise}/copilot/metrics/reports/users-28-day/latest
func (s *CopilotService) GetEnterpriseUsersMetricsReport(ctx context.Context, enterprise string) (*CopilotMetricsReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/copilot/metrics/reports/organization-1-day
func (s *CopilotService) GetOrganizationDailyMetricsReport(ctx context.Context, org string, opts *CopilotMetricsReportOptions) (*CopilotDailyMetricsReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/copilot/metrics/reports/organization-28-day/latest
func (s *CopilotService) GetOrganizationMetricsReport(ctx context.Context, org string) (*CopilotMetricsReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/copilot/metrics/reports/users-1-day
func (s *CopilotService) GetOrganizationUsersDailyMetricsReport(ctx context.Context, org string, opts *CopilotMetricsReportOptions) (*CopilotDailyMetricsReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/copilot/metrics/reports/users-28-day/latest
func (s *CopilotService) GetOrganizationUsersMetricsReport(ctx context.Context, org string) (*CopilotMetricsReport, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *CopilotService) DownloadCopilotMetrics(ctx context.Context, url string) ([]*CopilotMetrics, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type CopilotMetricsPullRequests struct {
	TotalReviewed                       *int     `json:"total_reviewed,omitempty"`
	TotalCreated                        *int     `json:"total_created,omitempty"`
	TotalCreatedByCopilot               *int     `json:"total_created_by_copilot,omitempty"`
	TotalReviewedByCopilot              *int     `json:"total_reviewed_by_copilot,omitempty"`
	TotalMerged                         *int     `json:"total_merged,omitempty"`
	MedianMinutesToMerge                *float64 `json:"median_minutes_to_merge,omitempty"`
	TotalSuggestions                    *int     `json:"total_suggestions,omitempty"`
	TotalAppliedSuggestions             *int     `json:"total_applied_suggestions,omitempty"`
	TotalMergedCreatedByCopilot         *int     `json:"total_merged_created_by_copilot,omitempty"`
	MedianMinutesToMergeCopilotAuthored *float64 `json:"median_minutes_to_merge_copilot_authored,omitempty"`
	TotalCopilotSuggestions             *int     `json:"total_copilot_suggestions,omitempty"`
	TotalCopilotAppliedSuggestions      *int     `json:"total_copilot_applied_suggestions,omitempty"`
	MedianMinutesToMergeCopilotReviewed *float64 `json:"median_minutes_to_merge_copilot_reviewed,omitempty"`
	TotalMergedReviewedByCopilot        *int     `json:"total_merged_reviewed_by_copilot,omitempty"`
}

type CopilotMetricsCodeActivity struct {
	CodeGenerationActivityCount *int `json:"code_generation_activity_count,omitempty"`
	CodeAcceptanceActivityCount *int `json:"code_acceptance_activity_count,omitempty"`
	LOCSuggestedToAddSum        *int `json:"loc_suggested_to_add_sum,omitempty"`
	LOCSuggestedToDeleteSum     *int `json:"loc_suggested_to_delete_sum,omitempty"`
	LOCAddedSum                 *int `json:"loc_added_sum,omitempty"`
	LOCDeletedSum               *int `json:"loc_deleted_sum,omitempty"`
}

type CopilotMetricsChatPanel struct {
	ChatPanelAgentMode   *int `json:"chat_panel_agent_mode,omitempty"`
	ChatPanelAskMode     *int `json:"chat_panel_ask_mode,omitempty"`
	ChatPanelCustomMode  *int `json:"chat_panel_custom_mode,omitempty"`
	ChatPanelEditMode    *int `json:"chat_panel_edit_mode,omitempty"`
	ChatPanelUnknownMode *int `json:"chat_panel_unknown_mode,omitempty"`
}

type CopilotMetricsIDE struct {
	IDE                           string `json:"ide"`
	UserInitiatedInteractionCount *int   `json:"user_initiated_interaction_count,omitempty"`
	CopilotMetricsCodeActivity
}

type CopilotMetricsFeature struct {
	Feature                       string `json:"feature"`
	UserInitiatedInteractionCount *int   `json:"user_initiated_interaction_count,omitempty"`
	CopilotMetricsCodeActivity
}

type CopilotMetricsLanguageFeature struct {
	Language string `json:"language"`
	Feature  string `json:"feature"`
	CopilotMetricsCodeActivity
}

type CopilotMetricsLanguageModel struct {
	Language string `json:"language"`
	Model    string `json:"model"`
	CopilotMetricsCodeActivity
}

type CopilotMetricsModelFeature struct {
	Model                         string `json:"model"`
	Feature                       string `json:"feature"`
	UserInitiatedInteractionCount *int   `json:"user_initiated_interaction_count,omitempty"`
	CopilotMetricsCodeActivity
}

type CopilotMetricsCLIVersion struct {
	SampledAt  *Timestamp `json:"sampled_at,omitempty"`
	CLIVersion string     `json:"cli_version"`
}

type CopilotMetricsCLITokenUsage struct {
	AvgTokensPerRequest *float64 `json:"avg_tokens_per_request,omitempty"`
	OutputTokensSum     *int     `json:"output_tokens_sum,omitempty"`
	PromptTokensSum     *int     `json:"prompt_tokens_sum,omitempty"`
}

type CopilotMetricsCLI struct {
	SessionCount        *int                         `json:"session_count,omitempty"`
	RequestCount        *int                         `json:"request_count,omitempty"`
	PromptCount         *int                         `json:"prompt_count,omitempty"`
	TokenUsage          *CopilotMetricsCLITokenUsage `json:"token_usage,omitempty"`
	LastKnownCLIVersion *CopilotMetricsCLIVersion    `json:"last_known_cli_version,omitempty"`
}

type CopilotMetricsCopilotAppTokenUsage struct {
	AvgTokensPerRequest *float64 `json:"avg_tokens_per_request,omitempty"`
	OutputTokensSum     *int     `json:"output_tokens_sum,omitempty"`
	PromptTokensSum     *int     `json:"prompt_tokens_sum,omitempty"`
}

type CopilotMetricsCopilotApp struct {
	SessionCount *int                                `json:"session_count,omitempty"`
	RequestCount *int                                `json:"request_count,omitempty"`
	PromptCount  *int                                `json:"prompt_count,omitempty"`
	TokenUsage   *CopilotMetricsCopilotAppTokenUsage `json:"token_usage,omitempty"`
}

type CopilotMetricsThirdPartyAgent struct {
	AgentName                     string `json:"agent_name"`
	AgentID                       string `json:"agent_id"`
	UserInitiatedInteractionCount *int   `json:"user_initiated_interaction_count,omitempty"`
	SessionCount                  *int   `json:"session_count,omitempty"`
}

type CopilotMetricsAIAdoptionPhase struct {
	PhaseNumber int    `json:"phase_number"`
	Phase       string `json:"phase"`
	Version     string `json:"version"`
}

type CopilotMetricsAIAdoptionPhaseTotals struct {
	Phase                               string  `json:"phase"`
	PhaseNumber                         int     `json:"phase_number"`
	TotalEngagedUsers                   int     `json:"total_engaged_users"`
	AvgUserInitiatedInteractions        float64 `json:"avg_user_initiated_interactions"`
	AvgCodeGenerationActivities         float64 `json:"avg_code_generation_activities"`
	AvgCodeAcceptanceActivities         float64 `json:"avg_code_acceptance_activities"`
	AvgLOCAdded                         float64 `json:"avg_loc_added"`
	AvgLOCDeleted                       float64 `json:"avg_loc_deleted"`
	AvgPullRequestsReviewed             float64 `json:"avg_pull_requests_reviewed"`
	AvgPullRequestsCreated              float64 `json:"avg_pull_requests_created"`
	AvgPullRequestsMerged               float64 `json:"avg_pull_requests_merged"`
	TotalPullRequestsMerged             int     `json:"total_pull_requests_merged"`
	AvgPullRequestsMedianMinutesToMerge float64 `json:"avg_pull_requests_median_minutes_to_merge"`
	AvgPullRequestsMinutesToReview      float64 `json:"avg_pull_requests_minutes_to_review"`
	AvgPullRequestsReviewCycles         float64 `json:"avg_pull_requests_review_cycles"`
}

type CopilotDailyMetrics struct {
	Day                                  string  `json:"day"`
	OrganizationID                       *string `json:"organization_id,omitempty"`
	EnterpriseID                         *string `json:"enterprise_id,omitempty"`
	DailyActiveCLIUsers                  *int    `json:"daily_active_cli_users,omitempty"`
	DailyActiveCopilotAppUsers           *int    `json:"daily_active_copilot_app_users,omitempty"`
	DailyActiveUsers                     *int    `json:"daily_active_users,omitempty"`
	DailyActiveCopilotCloudAgentUsers    *int    `json:"daily_active_copilot_cloud_agent_users,omitempty"`
	WeeklyActiveUsers                    *int    `json:"weekly_active_users,omitempty"`
	WeeklyActiveCopilotCloudAgentUsers   *int    `json:"weekly_active_copilot_cloud_agent_users,omitempty"`
	MonthlyActiveUsers                   *int    `json:"monthly_active_users,omitempty"`
	MonthlyActiveChatUsers               *int    `json:"monthly_active_chat_users,omitempty"`
	MonthlyActiveAgentUsers              *int    `json:"monthly_active_agent_users,omitempty"`
	MonthlyActiveCopilotCloudAgentUsers  *int    `json:"monthly_active_copilot_cloud_agent_users,omitempty"`
	DailyActiveCopilotCodeReviewUsers    *int    `json:"daily_active_copilot_code_review_users,omitempty"`
	WeeklyActiveCopilotCodeReviewUsers   *int    `json:"weekly_active_copilot_code_review_users,omitempty"`
	MonthlyActiveCopilotCodeReviewUsers  *int    `json:"monthly_active_copilot_code_review_users,omitempty"`
	DailyPassiveCopilotCodeReviewUsers   *int    `json:"daily_passive_copilot_code_review_users,omitempty"`
	WeeklyPassiveCopilotCodeReviewUsers  *int    `json:"weekly_passive_copilot_code_review_users,omitempty"`
	MonthlyPassiveCopilotCodeReviewUsers *int    `json:"monthly_passive_copilot_code_review_users,omitempty"`
	UserInitiatedInteractionCount        *int    `json:"user_initiated_interaction_count,omitempty"`
	CopilotMetricsChatPanel
	CodeGenerationActivityCount *int                                   `json:"code_generation_activity_count,omitempty"`
	CodeAcceptanceActivityCount *int                                   `json:"code_acceptance_activity_count,omitempty"`
	TotalsByIDE                 []*CopilotMetricsIDE                   `json:"totals_by_ide,omitempty"`
	TotalsByFeature             []*CopilotMetricsFeature               `json:"totals_by_feature,omitempty"`
	TotalsByLanguageFeature     []*CopilotMetricsLanguageFeature       `json:"totals_by_language_feature,omitempty"`
	TotalsByLanguageModel       []*CopilotMetricsLanguageModel         `json:"totals_by_language_model,omitempty"`
	TotalsByModelFeature        []*CopilotMetricsModelFeature          `json:"totals_by_model_feature,omitempty"`
	TotalsByCLI                 *CopilotMetricsCLI                     `json:"totals_by_cli,omitempty"`
	TotalsByCopilotApp          *CopilotMetricsCopilotApp              `json:"totals_by_copilot_app,omitempty"`
	TotalsBy3rdPartyAgent       []*CopilotMetricsThirdPartyAgent       `json:"totals_by_3rd_party_agent,omitempty"`
	TotalsByAIAdoptionPhase     []*CopilotMetricsAIAdoptionPhaseTotals `json:"totals_by_ai_adoption_phase,omitempty"`
	LOCSuggestedToAddSum        *int                                   `json:"loc_suggested_to_add_sum,omitempty"`
	LOCSuggestedToDeleteSum     *int                                   `json:"loc_suggested_to_delete_sum,omitempty"`
	LOCAddedSum                 *int                                   `json:"loc_added_sum,omitempty"`
	LOCDeletedSum               *int                                   `json:"loc_deleted_sum,omitempty"`
	PullRequests                *CopilotMetricsPullRequests            `json:"pull_requests,omitempty"`
}

type CopilotPeriodicMetrics struct {
	ReportStartDay string                 `json:"report_start_day"`
	ReportEndDay   string                 `json:"report_end_day"`
	OrganizationID *string                `json:"organization_id,omitempty"`
	EnterpriseID   *string                `json:"enterprise_id,omitempty"`
	CreatedAt      *Timestamp             `json:"created_at,omitempty"`
	DayTotals      []*CopilotDailyMetrics `json:"day_totals,omitempty"`
}

type CopilotUserMetricsPluginVersion struct {
	SampledAt     *Timestamp `json:"sampled_at,omitempty"`
	Plugin        string     `json:"plugin"`
	PluginVersion string     `json:"plugin_version"`
}

type CopilotUserMetricsIDEVersion struct {
	SampledAt  *Timestamp `json:"sampled_at,omitempty"`
	IDEVersion string     `json:"ide_version"`
}

type CopilotUserMetricsIDE struct {
	IDE                           string `json:"ide"`
	UserInitiatedInteractionCount *int   `json:"user_initiated_interaction_count,omitempty"`
	CopilotMetricsCodeActivity
	LastKnownPluginVersion *CopilotUserMetricsPluginVersion `json:"last_known_plugin_version,omitempty"`
	LastKnownIDEVersion    *CopilotUserMetricsIDEVersion    `json:"last_known_ide_version,omitempty"`
}

type CopilotUserDailyMetrics struct {
	UserID                        int      `json:"user_id"`
	UserLogin                     string   `json:"user_login"`
	Day                           string   `json:"day"`
	OrganizationID                *string  `json:"organization_id,omitempty"`
	EnterpriseID                  *string  `json:"enterprise_id,omitempty"`
	AICreditsUsed                 *float64 `json:"ai_credits_used,omitempty"`
	UserInitiatedInteractionCount *int     `json:"user_initiated_interaction_count,omitempty"`
	CopilotMetricsChatPanel
	CodeGenerationActivityCount  *int                             `json:"code_generation_activity_count,omitempty"`
	CodeAcceptanceActivityCount  *int                             `json:"code_acceptance_activity_count,omitempty"`
	TotalsByIDE                  []*CopilotUserMetricsIDE         `json:"totals_by_ide,omitempty"`
	TotalsByFeature              []*CopilotMetricsFeature         `json:"totals_by_feature,omitempty"`
	TotalsByLanguageFeature      []*CopilotMetricsLanguageFeature `json:"totals_by_language_feature,omitempty"`
	TotalsByLanguageModel        []*CopilotMetricsLanguageModel   `json:"totals_by_language_model,omitempty"`
	TotalsByModelFeature         []*CopilotMetricsModelFeature    `json:"totals_by_model_feature,omitempty"`
	TotalsByCLI                  *CopilotMetricsCLI               `json:"totals_by_cli,omitempty"`
	TotalsByCopilotApp           *CopilotMetricsCopilotApp        `json:"totals_by_copilot_app,omitempty"`
	TotalsBy3rdPartyAgent        []*CopilotMetricsThirdPartyAgent `json:"totals_by_3rd_party_agent,omitempty"`
	AIAdoptionPhase              *CopilotMetricsAIAdoptionPhase   `json:"ai_adoption_phase,omitempty"`
	UsedAgent                    *bool                            `json:"used_agent,omitempty"`
	UsedChat                     *bool                            `json:"used_chat,omitempty"`
	UsedCLI                      *bool                            `json:"used_cli,omitempty"`
	UsedCopilotApp               *bool                            `json:"used_copilot_app,omitempty"`
	UsedCopilotCloudAgent        *bool                            `json:"used_copilot_cloud_agent,omitempty"`
	UsedCopilotCodeReviewActive  *bool                            `json:"used_copilot_code_review_active,omitempty"`
	UsedCopilotCodeReviewPassive *bool                            `json:"used_copilot_code_review_passive,omitempty"`
	UsedCopilotCodingAgent       *bool                            `json:"used_copilot_coding_agent,omitempty"`
	LOCSuggestedToAddSum         *int                             `json:"loc_suggested_to_add_sum,omitempty"`
	LOCSuggestedToDeleteSum      *int                             `json:"loc_suggested_to_delete_sum,omitempty"`
	LOCAddedSum                  *int                             `json:"loc_added_sum,omitempty"`
	LOCDeletedSum                *int                             `json:"loc_deleted_sum,omitempty"`
}

type CopilotUserPeriodicMetrics struct {
	ReportStartDay                string   `json:"report_start_day"`
	ReportEndDay                  string   `json:"report_end_day"`
	Day                           string   `json:"day"`
	OrganizationID                *string  `json:"organization_id,omitempty"`
	EnterpriseID                  *string  `json:"enterprise_id,omitempty"`
	UserID                        int      `json:"user_id"`
	UserLogin                     string   `json:"user_login"`
	AICreditsUsed                 *float64 `json:"ai_credits_used,omitempty"`
	UserInitiatedInteractionCount *int     `json:"user_initiated_interaction_count,omitempty"`
	CopilotMetricsChatPanel
	CodeGenerationActivityCount  *int                             `json:"code_generation_activity_count,omitempty"`
	CodeAcceptanceActivityCount  *int                             `json:"code_acceptance_activity_count,omitempty"`
	TotalsByIDE                  []*CopilotUserMetricsIDE         `json:"totals_by_ide,omitempty"`
	TotalsByFeature              []*CopilotMetricsFeature         `json:"totals_by_feature,omitempty"`
	TotalsByLanguageFeature      []*CopilotMetricsLanguageFeature `json:"totals_by_language_feature,omitempty"`
	TotalsByLanguageModel        []*CopilotMetricsLanguageModel   `json:"totals_by_language_model,omitempty"`
	TotalsByModelFeature         []*CopilotMetricsModelFeature    `json:"totals_by_model_feature,omitempty"`
	TotalsByCLI                  *CopilotMetricsCLI               `json:"totals_by_cli,omitempty"`
	TotalsByCopilotApp           *CopilotMetricsCopilotApp        `json:"totals_by_copilot_app,omitempty"`
	TotalsBy3rdPartyAgent        []*CopilotMetricsThirdPartyAgent `json:"totals_by_3rd_party_agent,omitempty"`
	AIAdoptionPhase              *CopilotMetricsAIAdoptionPhase   `json:"ai_adoption_phase,omitempty"`
	UsedAgent                    *bool                            `json:"used_agent,omitempty"`
	UsedChat                     *bool                            `json:"used_chat,omitempty"`
	UsedCLI                      *bool                            `json:"used_cli,omitempty"`
	UsedCopilotApp               *bool                            `json:"used_copilot_app,omitempty"`
	UsedCopilotCloudAgent        *bool                            `json:"used_copilot_cloud_agent,omitempty"`
	UsedCopilotCodeReviewActive  *bool                            `json:"used_copilot_code_review_active,omitempty"`
	UsedCopilotCodeReviewPassive *bool                            `json:"used_copilot_code_review_passive,omitempty"`
	UsedCopilotCodingAgent       *bool                            `json:"used_copilot_coding_agent,omitempty"`
	LOCSuggestedToAddSum         *int                             `json:"loc_suggested_to_add_sum,omitempty"`
	LOCSuggestedToDeleteSum      *int                             `json:"loc_suggested_to_delete_sum,omitempty"`
	LOCAddedSum                  *int                             `json:"loc_added_sum,omitempty"`
	LOCDeletedSum                *int                             `json:"loc_deleted_sum,omitempty"`
}

func (s *CopilotService) fetchMetricsReport(ctx context.Context, url string) (*http.Response, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func decodeNDJSONMetrics[T any](r io.Reader) ([]*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *CopilotService) DownloadDailyMetrics(ctx context.Context, url string) (*CopilotDailyMetrics, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *CopilotService) DownloadPeriodicMetrics(ctx context.Context, url string) (*CopilotPeriodicMetrics, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *CopilotService) DownloadUserDailyMetrics(ctx context.Context, url string) ([]*CopilotUserDailyMetrics, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *CopilotService) DownloadUserPeriodicMetrics(ctx context.Context, url string) ([]*CopilotUserPeriodicMetrics, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
