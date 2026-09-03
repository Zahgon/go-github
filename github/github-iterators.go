package github

import (
	"context"
	"iter"
)

func (s *ActionsService) ListArtifactsIter(ctx context.Context, owner string, repo string, opts *ListArtifactsOptions) iter.Seq2[*Artifact, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListCacheUsageByRepoForOrgIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*ActionsCacheUsage, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListCachesIter(ctx context.Context, owner string, repo string, opts *ActionsCacheListOptions) iter.Seq2[*ActionsCache, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListEnabledOrgsInEnterpriseIter(ctx context.Context, owner string, opts *ListOptions) iter.Seq2[*Organization, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListEnabledReposInOrgIter(ctx context.Context, owner string, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListEnvSecretsIter(ctx context.Context, owner string, repo string, env string, opts *ListOptions) iter.Seq2[*Secret, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListEnvVariablesIter(ctx context.Context, owner string, repo string, env string, opts *ListOptions) iter.Seq2[*ActionsVariable, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListHostedRunnersIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*HostedRunner, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListOrgSecretsIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Secret, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListOrgVariablesIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*ActionsVariable, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListOrganizationRunnerGroupsIter(ctx context.Context, org string, opts *ListOrgRunnerGroupOptions) iter.Seq2[*RunnerGroup, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListOrganizationRunnersIter(ctx context.Context, org string, opts *ListRunnersOptions) iter.Seq2[*Runner, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListRepoOrgSecretsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Secret, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListRepoOrgVariablesIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*ActionsVariable, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListRepoSecretsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Secret, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListRepoVariablesIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*ActionsVariable, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListRepositoriesSelfHostedRunnersAllowedInOrganizationIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListRepositoryAccessRunnerGroupIter(ctx context.Context, org string, groupID int64, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListRepositoryWorkflowRunsIter(ctx context.Context, owner string, repo string, opts *ListWorkflowRunsOptions) iter.Seq2[*WorkflowRun, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListRunnerGroupHostedRunnersIter(ctx context.Context, org string, groupID int64, opts *ListOptions) iter.Seq2[*HostedRunner, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListRunnerGroupRunnersIter(ctx context.Context, org string, groupID int64, opts *ListOptions) iter.Seq2[*Runner, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListRunnersIter(ctx context.Context, owner string, repo string, opts *ListRunnersOptions) iter.Seq2[*Runner, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListSelectedReposForOrgSecretIter(ctx context.Context, org string, name string, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListSelectedReposForOrgVariableIter(ctx context.Context, org string, name string, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListWorkflowJobsIter(ctx context.Context, owner string, repo string, runID int64, opts *ListWorkflowJobsOptions) iter.Seq2[*WorkflowJob, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListWorkflowJobsAttemptIter(ctx context.Context, owner string, repo string, runID int64, attemptNumber int64, opts *ListOptions) iter.Seq2[*WorkflowJob, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListWorkflowRunArtifactsIter(ctx context.Context, owner string, repo string, runID int64, opts *ListOptions) iter.Seq2[*Artifact, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListWorkflowRunsByFileNameIter(ctx context.Context, owner string, repo string, workflowFileName string, opts *ListWorkflowRunsOptions) iter.Seq2[*WorkflowRun, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListWorkflowRunsByIDIter(ctx context.Context, owner string, repo string, workflowID int64, opts *ListWorkflowRunsOptions) iter.Seq2[*WorkflowRun, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActionsService) ListWorkflowsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Workflow, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListEventsIter(ctx context.Context, opts *ListOptions) iter.Seq2[*Event, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListEventsForOrganizationIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Event, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListEventsForRepoNetworkIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Event, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListEventsPerformedByUserIter(ctx context.Context, user string, publicOnly bool, opts *ListOptions) iter.Seq2[*Event, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListEventsReceivedByUserIter(ctx context.Context, user string, publicOnly bool, opts *ListOptions) iter.Seq2[*Event, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListIssueEventsForRepositoryIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*IssueEvent, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListNotificationsIter(ctx context.Context, opts *NotificationListOptions) iter.Seq2[*Notification, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListRepositoryEventsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Event, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListRepositoryNotificationsIter(ctx context.Context, owner string, repo string, opts *NotificationListOptions) iter.Seq2[*Notification, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListStargazersIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Stargazer, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListStarredIter(ctx context.Context, user string, opts *ActivityListStarredOptions) iter.Seq2[*StarredRepository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListUserEventsForOrganizationIter(ctx context.Context, org string, user string, opts *ListOptions) iter.Seq2[*Event, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListWatchedIter(ctx context.Context, user string, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ActivityService) ListWatchersIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *AgentTasksService) ListIter(ctx context.Context, opts *AgentTaskListOptions) iter.Seq2[*AgentTask, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *AgentTasksService) ListByRepoIter(ctx context.Context, owner string, repo string, opts *AgentTaskListByRepoOptions) iter.Seq2[*AgentTask, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *AppsService) ListHookDeliveriesIter(ctx context.Context, opts *ListCursorOptions) iter.Seq2[*HookDelivery, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *AppsService) ListInstallationRequestsIter(ctx context.Context, opts *ListOptions) iter.Seq2[*InstallationRequest, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *AppsService) ListInstallationsIter(ctx context.Context, opts *ListOptions) iter.Seq2[*Installation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *AppsService) ListReposIter(ctx context.Context, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *AppsService) ListUserInstallationsIter(ctx context.Context, opts *ListOptions) iter.Seq2[*Installation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *AppsService) ListUserReposIter(ctx context.Context, id int64, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ChecksService) ListCheckRunAnnotationsIter(ctx context.Context, owner string, repo string, checkRunID int64, opts *ListOptions) iter.Seq2[*CheckRunAnnotation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ChecksService) ListCheckRunsCheckSuiteIter(ctx context.Context, owner string, repo string, checkSuiteID int64, opts *ListCheckRunsOptions) iter.Seq2[*CheckRun, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ChecksService) ListCheckRunsForRefIter(ctx context.Context, owner string, repo string, ref string, opts *ListCheckRunsOptions) iter.Seq2[*CheckRun, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ChecksService) ListCheckSuitesForRefIter(ctx context.Context, owner string, repo string, ref string, opts *ListCheckSuiteOptions) iter.Seq2[*CheckSuite, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClassroomService) ListAcceptedAssignmentsIter(ctx context.Context, assignmentID int64, opts *ListOptions) iter.Seq2[*AcceptedAssignment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClassroomService) ListClassroomAssignmentsIter(ctx context.Context, classroomID int64, opts *ListOptions) iter.Seq2[*ClassroomAssignment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClassroomService) ListClassroomsIter(ctx context.Context, opts *ListOptions) iter.Seq2[*Classroom, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodeQualityService) ListFindingsIter(ctx context.Context, owner string, repo string, opts *ListCodeQualityFindingsOptions) iter.Seq2[*CodeQualityFinding, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodeScanningService) ListAlertInstancesIter(ctx context.Context, owner string, repo string, id int64, opts *AlertInstancesListOptions) iter.Seq2[*MostRecentInstance, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodeScanningService) ListAlertsForOrgIter(ctx context.Context, org string, opts *AlertListOptions) iter.Seq2[*Alert, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodeScanningService) ListAlertsForRepoIter(ctx context.Context, owner string, repo string, opts *AlertListOptions) iter.Seq2[*Alert, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodeScanningService) ListAnalysesForRepoIter(ctx context.Context, owner string, repo string, opts *AnalysesListOptions) iter.Seq2[*ScanningAnalysis, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodespacesService) ListIter(ctx context.Context, opts *ListCodespacesOptions) iter.Seq2[*Codespace, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodespacesService) ListDevContainerConfigurationsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*DevContainer, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodespacesService) ListInOrgIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Codespace, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodespacesService) ListInRepoIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Codespace, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodespacesService) ListOrgSecretsIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Secret, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodespacesService) ListRepoSecretsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Secret, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodespacesService) ListSelectedReposForOrgSecretIter(ctx context.Context, org string, name string, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodespacesService) ListSelectedReposForUserSecretIter(ctx context.Context, name string, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodespacesService) ListUserCodespacesInOrgIter(ctx context.Context, org string, username string, opts *ListOptions) iter.Seq2[*Codespace, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CodespacesService) ListUserSecretsIter(ctx context.Context, opts *ListOptions) iter.Seq2[*Secret, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CopilotService) ListCopilotEnterpriseSeatsIter(ctx context.Context, enterprise string, opts *ListOptions) iter.Seq2[*CopilotSeatDetails, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CopilotService) ListCopilotSeatsIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*CopilotSeatDetails, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *CopilotService) ListOrganizationCodingAgentRepositoriesIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *DependabotService) ListOrgAlertsIter(ctx context.Context, org string, opts *ListAlertsOptions) iter.Seq2[*DependabotAlert, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *DependabotService) ListOrgSecretsIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Secret, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *DependabotService) ListRepoAlertsIter(ctx context.Context, owner string, repo string, opts *ListAlertsOptions) iter.Seq2[*DependabotAlert, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *DependabotService) ListRepoSecretsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Secret, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *DependabotService) ListSelectedReposForOrgSecretIter(ctx context.Context, org string, name string, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListAppAccessibleOrganizationRepositoriesIter(ctx context.Context, enterprise string, org string, opts *ListOptions) iter.Seq2[*AccessibleRepository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListAppInstallableOrganizationsIter(ctx context.Context, enterprise string, opts *ListOptions) iter.Seq2[*InstallableOrganization, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListAppInstallationsIter(ctx context.Context, enterprise string, org string, opts *ListOptions) iter.Seq2[*Installation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListAssignmentsIter(ctx context.Context, enterprise string, enterpriseTeam string, opts *ListOptions) iter.Seq2[*Organization, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListCodeSecurityConfigurationRepositoriesIter(ctx context.Context, enterprise string, configurationID int64, opts *ListCodeSecurityConfigurationRepositoriesOptions) iter.Seq2[*RepositoryAttachment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListCodeSecurityConfigurationsIter(ctx context.Context, enterprise string, opts *ListEnterpriseCodeSecurityConfigurationOptions) iter.Seq2[*CodeSecurityConfiguration, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListConsumedLicensesIter(ctx context.Context, enterprise string, opts *ListOptions) iter.Seq2[*EnterpriseLicensedUsers, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListEnterpriseNetworkConfigurationsIter(ctx context.Context, enterprise string, opts *ListOptions) iter.Seq2[*NetworkConfiguration, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListHostedRunnersIter(ctx context.Context, enterprise string, opts *ListOptions) iter.Seq2[*HostedRunner, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListOrganizationAccessRunnerGroupIter(ctx context.Context, enterprise string, groupID int64, opts *ListOptions) iter.Seq2[*Organization, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListOrganizationCustomPropertyValuesIter(ctx context.Context, enterprise string, opts *ListOptions) iter.Seq2[*EnterpriseCustomPropertiesValues, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListRepositoriesForOrgAppInstallationIter(ctx context.Context, enterprise string, org string, installationID int64, opts *ListOptions) iter.Seq2[*AccessibleRepository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListRunnerGroupRunnersIter(ctx context.Context, enterprise string, groupID int64, opts *ListOptions) iter.Seq2[*Runner, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListRunnerGroupsIter(ctx context.Context, enterprise string, opts *ListEnterpriseRunnerGroupOptions) iter.Seq2[*EnterpriseRunnerGroup, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListRunnersIter(ctx context.Context, enterprise string, opts *ListRunnersOptions) iter.Seq2[*Runner, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListTeamMembersIter(ctx context.Context, enterprise string, enterpriseTeam string, opts *ListOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListTeamsIter(ctx context.Context, enterprise string, opts *ListOptions) iter.Seq2[*EnterpriseTeam, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnterpriseService) ListVisualStudioSubscriptionsIter(ctx context.Context, enterprise string, opts *ListVisualStudioSubscriptionsOptions) iter.Seq2[*VisualStudioSubscriptionAssignment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *GistsService) ListIter(ctx context.Context, user string, opts *GistListOptions) iter.Seq2[*Gist, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *GistsService) ListAllIter(ctx context.Context, opts *GistListOptions) iter.Seq2[*Gist, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *GistsService) ListCommentsIter(ctx context.Context, gistID string, opts *ListOptions) iter.Seq2[*GistComment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *GistsService) ListCommitsIter(ctx context.Context, id string, opts *ListOptions) iter.Seq2[*GistCommit, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *GistsService) ListForksIter(ctx context.Context, id string, opts *ListOptions) iter.Seq2[*GistFork, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *GistsService) ListStarredIter(ctx context.Context, opts *GistListOptions) iter.Seq2[*Gist, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListAllIssuesIter(ctx context.Context, opts *ListAllIssuesOptions) iter.Seq2[*Issue, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListAssigneesIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListBlockedByIter(ctx context.Context, owner string, repo string, issueNumber int64, opts *ListOptions) iter.Seq2[*Issue, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListBlockingIter(ctx context.Context, owner string, repo string, issueNumber int64, opts *ListOptions) iter.Seq2[*Issue, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListByOrgIter(ctx context.Context, org string, opts *IssueListByOrgOptions) iter.Seq2[*Issue, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListByRepoIter(ctx context.Context, owner string, repo string, opts *IssueListByRepoOptions) iter.Seq2[*Issue, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListCommentsIter(ctx context.Context, owner string, repo string, number int, opts *IssueListCommentsOptions) iter.Seq2[*IssueComment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListIssueEventsIter(ctx context.Context, owner string, repo string, number int, opts *ListOptions) iter.Seq2[*IssueEvent, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListIssueTimelineIter(ctx context.Context, owner string, repo string, number int, opts *ListOptions) iter.Seq2[*Timeline, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListLabelsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Label, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListLabelsByIssueIter(ctx context.Context, owner string, repo string, number int, opts *ListOptions) iter.Seq2[*Label, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListLabelsForMilestoneIter(ctx context.Context, owner string, repo string, number int, opts *ListOptions) iter.Seq2[*Label, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListMilestonesIter(ctx context.Context, owner string, repo string, opts *MilestoneListOptions) iter.Seq2[*Milestone, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListRepositoryEventsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*IssueEvent, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *IssuesService) ListUserIssuesIter(ctx context.Context, opts *ListUserIssuesOptions) iter.Seq2[*Issue, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *LicensesService) ListIter(ctx context.Context, opts *ListLicensesOptions) iter.Seq2[*License, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *MarketplaceService) ListMarketplacePurchasesForUserIter(ctx context.Context, opts *ListOptions) iter.Seq2[*MarketplacePurchase, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *MarketplaceService) ListPlanAccountsForPlanIter(ctx context.Context, planID int64, opts *ListOptions) iter.Seq2[*MarketplacePlanAccount, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *MarketplaceService) ListPlansIter(ctx context.Context, opts *ListOptions) iter.Seq2[*MarketplacePlan, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *MigrationService) ListMigrationsIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Migration, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *MigrationService) ListUserMigrationsIter(ctx context.Context, opts *ListOptions) iter.Seq2[*UserMigration, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListIter(ctx context.Context, user string, opts *ListOptions) iter.Seq2[*Organization, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListAllRepositoryRulesetsIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*RepositoryRuleset, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListAttestationsIter(ctx context.Context, org string, subjectDigest string, opts *ListOptions) iter.Seq2[*Attestation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListBlockedUsersIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListCodeSecurityConfigurationRepositoriesIter(ctx context.Context, org string, configurationID int64, opts *ListCodeSecurityConfigurationRepositoriesOptions) iter.Seq2[*RepositoryAttachment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListCodeSecurityConfigurationsIter(ctx context.Context, org string, opts *ListOrgCodeSecurityConfigurationOptions) iter.Seq2[*CodeSecurityConfiguration, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListCredentialAuthorizationsIter(ctx context.Context, org string, opts *CredentialAuthorizationsListOptions) iter.Seq2[*CredentialAuthorization, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListCustomPropertyValuesIter(ctx context.Context, org string, opts *ListCustomPropertyValuesOptions) iter.Seq2[*RepoCustomPropertyValue, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListFailedOrgInvitationsIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Invitation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListFineGrainedPersonalAccessTokenRequestsIter(ctx context.Context, org string, opts *ListFineGrainedPATOptions) iter.Seq2[*FineGrainedPersonalAccessTokenRequest, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListFineGrainedPersonalAccessTokensIter(ctx context.Context, org string, opts *ListFineGrainedPATOptions) iter.Seq2[*PersonalAccessToken, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListHookDeliveriesIter(ctx context.Context, org string, id int64, opts *ListCursorOptions) iter.Seq2[*HookDelivery, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListHooksIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Hook, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListImmutableReleaseRepositoriesIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListInstallationsIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Installation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListMembersIter(ctx context.Context, org string, opts *ListMembersOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListNetworkConfigurationsIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*NetworkConfiguration, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListOrgInvitationTeamsIter(ctx context.Context, org string, invitationID string, opts *ListOptions) iter.Seq2[*Team, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListOrgMembershipsIter(ctx context.Context, opts *ListOrgMembershipsOptions) iter.Seq2[*Membership, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListOutsideCollaboratorsIter(ctx context.Context, org string, opts *ListOutsideCollaboratorsOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListPackagesIter(ctx context.Context, org string, opts *PackageListOptions) iter.Seq2[*Package, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListPendingOrgInvitationsIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Invitation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListTeamsAssignedToOrgRoleIter(ctx context.Context, org string, roleID int64, opts *ListOptions) iter.Seq2[*Team, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *OrganizationsService) ListUsersAssignedToOrgRoleIter(ctx context.Context, org string, roleID int64, opts *ListOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *PrivateRegistriesService) ListOrganizationPrivateRegistriesIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*PrivateRegistry, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProjectsService) ListOrganizationProjectFieldsIter(ctx context.Context, org string, projectNumber int, opts *ListProjectsOptions) iter.Seq2[*ProjectV2Field, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProjectsService) ListOrganizationProjectItemsIter(ctx context.Context, org string, projectNumber int, opts *ListProjectItemsOptions) iter.Seq2[*ProjectV2Item, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProjectsService) ListOrganizationProjectViewItemsIter(ctx context.Context, org string, projectNumber int, viewNumber int, opts *ListProjectItemsOptions) iter.Seq2[*ProjectV2Item, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProjectsService) ListOrganizationProjectsIter(ctx context.Context, org string, opts *ListProjectsOptions) iter.Seq2[*ProjectV2, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProjectsService) ListUserProjectFieldsIter(ctx context.Context, user string, projectNumber int, opts *ListProjectsOptions) iter.Seq2[*ProjectV2Field, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProjectsService) ListUserProjectItemsIter(ctx context.Context, username string, projectNumber int, opts *ListProjectItemsOptions) iter.Seq2[*ProjectV2Item, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProjectsService) ListUserProjectViewItemsIter(ctx context.Context, username string, projectNumber int, viewNumber int, opts *ListProjectItemsOptions) iter.Seq2[*ProjectV2Item, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ProjectsService) ListUserProjectsIter(ctx context.Context, username string, opts *ListProjectsOptions) iter.Seq2[*ProjectV2, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *PullRequestsService) ListIter(ctx context.Context, owner string, repo string, opts *PullRequestListOptions) iter.Seq2[*PullRequest, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *PullRequestsService) ListCommentsIter(ctx context.Context, owner string, repo string, number int, opts *PullRequestListCommentsOptions) iter.Seq2[*PullRequestComment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *PullRequestsService) ListCommitsIter(ctx context.Context, owner string, repo string, number int, opts *ListOptions) iter.Seq2[*RepositoryCommit, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *PullRequestsService) ListFilesIter(ctx context.Context, owner string, repo string, number int, opts *ListOptions) iter.Seq2[*CommitFile, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *PullRequestsService) ListPullRequestsWithCommitIter(ctx context.Context, owner string, repo string, sha string, opts *ListOptions) iter.Seq2[*PullRequest, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *PullRequestsService) ListReviewCommentsIter(ctx context.Context, owner string, repo string, number int, reviewID int64, opts *ListOptions) iter.Seq2[*PullRequestComment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *PullRequestsService) ListReviewsIter(ctx context.Context, owner string, repo string, number int, opts *ListOptions) iter.Seq2[*PullRequestReview, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReactionsService) ListCommentReactionsIter(ctx context.Context, owner string, repo string, id int64, opts *ListReactionOptions) iter.Seq2[*Reaction, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReactionsService) ListIssueCommentReactionsIter(ctx context.Context, owner string, repo string, id int64, opts *ListReactionOptions) iter.Seq2[*Reaction, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReactionsService) ListIssueReactionsIter(ctx context.Context, owner string, repo string, number int, opts *ListReactionOptions) iter.Seq2[*Reaction, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReactionsService) ListPullRequestCommentReactionsIter(ctx context.Context, owner string, repo string, id int64, opts *ListReactionOptions) iter.Seq2[*Reaction, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReactionsService) ListReleaseReactionsIter(ctx context.Context, owner string, repo string, releaseID int64, opts *ListReactionOptions) iter.Seq2[*Reaction, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReactionsService) ListTeamDiscussionCommentReactionsIter(ctx context.Context, teamID int64, discussionNumber int, commentNumber int, opts *ListReactionOptions) iter.Seq2[*Reaction, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReactionsService) ListTeamDiscussionReactionsIter(ctx context.Context, teamID int64, discussionNumber int, opts *ListReactionOptions) iter.Seq2[*Reaction, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListCommitComparisonFilesIter(ctx context.Context, owner string, repo string, base string, head string, opts *ListOptions) iter.Seq2[*CommitFile, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListCombinedStatusIter(ctx context.Context, owner string, repo string, ref string, opts *ListOptions) iter.Seq2[*RepoStatus, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListCommitFilesIter(ctx context.Context, owner string, repo string, sha string, opts *ListOptions) iter.Seq2[*CommitFile, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListIter(ctx context.Context, user string, opts *RepositoryListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListAllTopicsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[string, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListAttestationsIter(ctx context.Context, owner string, repo string, subjectDigest string, opts *ListOptions) iter.Seq2[*Attestation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListBranchesIter(ctx context.Context, owner string, repo string, opts *BranchListOptions) iter.Seq2[*Branch, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListByAuthenticatedUserIter(ctx context.Context, opts *RepositoryListByAuthenticatedUserOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListByOrgIter(ctx context.Context, org string, opts *RepositoryListByOrgOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListByUserIter(ctx context.Context, user string, opts *RepositoryListByUserOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListCollaboratorsIter(ctx context.Context, owner string, repo string, opts *ListCollaboratorsOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListCommentsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*RepositoryComment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListCommitCommentsIter(ctx context.Context, owner string, repo string, sha string, opts *ListOptions) iter.Seq2[*RepositoryComment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListCommitsIter(ctx context.Context, owner string, repo string, opts *CommitsListOptions) iter.Seq2[*RepositoryCommit, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListContributorsIter(ctx context.Context, owner string, repository string, opts *ListContributorsOptions) iter.Seq2[*Contributor, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListCustomDeploymentRuleIntegrationsIter(ctx context.Context, owner string, repo string, environment string, opts *ListOptions) iter.Seq2[*CustomDeploymentProtectionRuleApp, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListDeploymentBranchPoliciesIter(ctx context.Context, owner string, repo string, environment string, opts *ListOptions) iter.Seq2[*DeploymentBranchPolicy, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListDeploymentStatusesIter(ctx context.Context, owner string, repo string, deployment int64, opts *ListOptions) iter.Seq2[*DeploymentStatus, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListDeploymentsIter(ctx context.Context, owner string, repo string, opts *DeploymentsListOptions) iter.Seq2[*Deployment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListEnvironmentsIter(ctx context.Context, owner string, repo string, opts *EnvironmentListOptions) iter.Seq2[*Environment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListForksIter(ctx context.Context, owner string, repo string, opts *RepositoryListForksOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListHookDeliveriesIter(ctx context.Context, owner string, repo string, id int64, opts *ListCursorOptions) iter.Seq2[*HookDelivery, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListHooksIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Hook, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListInvitationsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*RepositoryInvitation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListKeysIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Key, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListPagesBuildsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*PagesBuild, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListPreReceiveHooksIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*PreReceiveHook, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListReleaseAssetsIter(ctx context.Context, owner string, repo string, id int64, opts *ListOptions) iter.Seq2[*ReleaseAsset, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListReleasesIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*RepositoryRelease, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListRepositoryActivitiesIter(ctx context.Context, owner string, repo string, opts *ListRepositoryActivityOptions) iter.Seq2[*RepositoryActivity, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListStatusesIter(ctx context.Context, owner string, repo string, ref string, opts *ListOptions) iter.Seq2[*RepoStatus, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListTagsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*RepositoryTag, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *RepositoriesService) ListTeamsIter(ctx context.Context, owner string, repo string, opts *ListOptions) iter.Seq2[*Team, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningService) ListAlertsForEnterpriseIter(ctx context.Context, enterprise string, opts *SecretScanningAlertListOptions) iter.Seq2[*SecretScanningAlert, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningService) ListAlertsForOrgIter(ctx context.Context, org string, opts *SecretScanningAlertListOptions) iter.Seq2[*SecretScanningAlert, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningService) ListAlertsForRepoIter(ctx context.Context, owner string, repo string, opts *SecretScanningAlertListOptions) iter.Seq2[*SecretScanningAlert, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningService) ListCustomPatternsForEnterpriseIter(ctx context.Context, enterprise string, opts *SecretScanningCustomPatternListOptions) iter.Seq2[*SecretScanningCustomPattern, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningService) ListCustomPatternsForOrgIter(ctx context.Context, org string, opts *SecretScanningCustomPatternListOptions) iter.Seq2[*SecretScanningCustomPattern, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningService) ListCustomPatternsForRepoIter(ctx context.Context, owner string, repo string, opts *SecretScanningCustomPatternListOptions) iter.Seq2[*SecretScanningCustomPattern, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecretScanningService) ListLocationsForAlertIter(ctx context.Context, owner string, repo string, number int64, opts *ListOptions) iter.Seq2[*SecretScanningAlertLocation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAdvisoriesService) ListGlobalSecurityAdvisoriesIter(ctx context.Context, opts *ListGlobalSecurityAdvisoriesOptions) iter.Seq2[*GlobalSecurityAdvisory, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAdvisoriesService) ListRepositorySecurityAdvisoriesIter(ctx context.Context, owner string, repo string, opts *ListRepositorySecurityAdvisoriesOptions) iter.Seq2[*SecurityAdvisory, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *SecurityAdvisoriesService) ListRepositorySecurityAdvisoriesForOrgIter(ctx context.Context, org string, opts *ListRepositorySecurityAdvisoriesOptions) iter.Seq2[*SecurityAdvisory, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *SubIssueService) ListByIssueIter(ctx context.Context, owner string, repo string, issueNumber int64, opts *ListOptions) iter.Seq2[*SubIssue, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListChildTeamsByParentIDIter(ctx context.Context, orgID int64, teamID int64, opts *ListOptions) iter.Seq2[*Team, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListChildTeamsByParentSlugIter(ctx context.Context, org string, slug string, opts *ListOptions) iter.Seq2[*Team, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListCommentsByIDIter(ctx context.Context, orgID int64, teamID int64, discussionNumber int, opts *DiscussionCommentListOptions) iter.Seq2[*DiscussionComment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListCommentsBySlugIter(ctx context.Context, org string, slug string, discussionNumber int, opts *DiscussionCommentListOptions) iter.Seq2[*DiscussionComment, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListDiscussionsByIDIter(ctx context.Context, orgID int64, teamID int64, opts *DiscussionListOptions) iter.Seq2[*TeamDiscussion, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListDiscussionsBySlugIter(ctx context.Context, org string, slug string, opts *DiscussionListOptions) iter.Seq2[*TeamDiscussion, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListExternalGroupsIter(ctx context.Context, org string, opts *ListExternalGroupsOptions) iter.Seq2[*ExternalGroup, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListIDPGroupsInOrganizationIter(ctx context.Context, org string, opts *ListIDPGroupsOptions) iter.Seq2[*IDPGroup, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListPendingTeamInvitationsByIDIter(ctx context.Context, orgID int64, teamID int64, opts *ListOptions) iter.Seq2[*Invitation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListPendingTeamInvitationsBySlugIter(ctx context.Context, org string, slug string, opts *ListOptions) iter.Seq2[*Invitation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListTeamMembersByIDIter(ctx context.Context, orgID int64, teamID int64, opts *TeamListTeamMembersOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListTeamMembersBySlugIter(ctx context.Context, org string, slug string, opts *TeamListTeamMembersOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListTeamReposByIDIter(ctx context.Context, orgID int64, teamID int64, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListTeamReposBySlugIter(ctx context.Context, org string, slug string, opts *ListOptions) iter.Seq2[*Repository, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListTeamsIter(ctx context.Context, org string, opts *ListOptions) iter.Seq2[*Team, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *TeamsService) ListUserTeamsIter(ctx context.Context, opts *ListOptions) iter.Seq2[*Team, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListAttestationsIter(ctx context.Context, user string, subjectDigest string, opts *ListOptions) iter.Seq2[*Attestation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListBlockedUsersIter(ctx context.Context, opts *ListOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListEmailsIter(ctx context.Context, opts *ListOptions) iter.Seq2[*UserEmail, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListFollowersIter(ctx context.Context, user string, opts *ListOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListFollowingIter(ctx context.Context, user string, opts *ListOptions) iter.Seq2[*User, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListGPGKeysIter(ctx context.Context, user string, opts *ListOptions) iter.Seq2[*GPGKey, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListInvitationsIter(ctx context.Context, opts *ListOptions) iter.Seq2[*RepositoryInvitation, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListKeysIter(ctx context.Context, user string, opts *ListOptions) iter.Seq2[*Key, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListPackageVersionsIter(ctx context.Context, packageType string, packageName string, opts *ListPackageVersionsOptions) iter.Seq2[*PackageVersion, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListPackagesIter(ctx context.Context, user string, opts *PackageListOptions) iter.Seq2[*Package, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListSSHSigningKeysIter(ctx context.Context, user string, opts *ListOptions) iter.Seq2[*SSHSigningKey, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListSocialAccountsIter(ctx context.Context, opts *ListOptions) iter.Seq2[*SocialAccount, error] {
	_ = "STUB: not implemented"
	return nil
}

func (s *UsersService) ListUserSocialAccountsIter(ctx context.Context, username string, opts *ListOptions) iter.Seq2[*SocialAccount, error] {
	_ = "STUB: not implemented"
	return nil
}
