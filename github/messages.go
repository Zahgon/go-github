package github

import (
	"hash"
	"io"
	"net/http"
	"reflect"
)

const (
	sha1Prefix = "sha1"

	sha256Prefix = "sha256"
	sha512Prefix = "sha512"

	SHA1SignatureHeader = "X-Hub-Signature"

	SHA256SignatureHeader = "X-Hub-Signature-256"

	EventTypeHeader = "X-Github-Event"

	DeliveryIDHeader = "X-Github-Delivery"

	maxPayloadSize = 25 * 1024 * 1024
)

var (
	eventTypeMapping = map[string]any{
		"branch_protection_configuration": &BranchProtectionConfigurationEvent{},
		"branch_protection_rule":          &BranchProtectionRuleEvent{},
		"check_run":                       &CheckRunEvent{},
		"check_suite":                     &CheckSuiteEvent{},
		"code_scanning_alert":             &CodeScanningAlertEvent{},
		"commit_comment":                  &CommitCommentEvent{},
		"content_reference":               &ContentReferenceEvent{},
		"create":                          &CreateEvent{},
		"custom_property":                 &CustomPropertyEvent{},
		"custom_property_values":          &CustomPropertyValuesEvent{},
		"delete":                          &DeleteEvent{},
		"dependabot_alert":                &DependabotAlertEvent{},
		"deploy_key":                      &DeployKeyEvent{},
		"deployment":                      &DeploymentEvent{},
		"deployment_review":               &DeploymentReviewEvent{},
		"deployment_status":               &DeploymentStatusEvent{},
		"deployment_protection_rule":      &DeploymentProtectionRuleEvent{},
		"discussion":                      &DiscussionEvent{},
		"discussion_comment":              &DiscussionCommentEvent{},
		"fork":                            &ForkEvent{},
		"github_app_authorization":        &GitHubAppAuthorizationEvent{},
		"gollum":                          &GollumEvent{},
		"installation":                    &InstallationEvent{},
		"installation_repositories":       &InstallationRepositoriesEvent{},
		"installation_target":             &InstallationTargetEvent{},
		"issue_comment":                   &IssueCommentEvent{},
		"issues":                          &IssuesEvent{},
		"label":                           &LabelEvent{},
		"marketplace_purchase":            &MarketplacePurchaseEvent{},
		"member":                          &MemberEvent{},
		"membership":                      &MembershipEvent{},
		"merge_group":                     &MergeGroupEvent{},
		"meta":                            &MetaEvent{},
		"milestone":                       &MilestoneEvent{},
		"organization":                    &OrganizationEvent{},
		"org_block":                       &OrgBlockEvent{},
		"package":                         &PackageEvent{},
		"page_build":                      &PageBuildEvent{},
		"personal_access_token_request":   &PersonalAccessTokenRequestEvent{},
		"ping":                            &PingEvent{},
		"projects_v2":                     &ProjectV2Event{},
		"projects_v2_item":                &ProjectV2ItemEvent{},
		"public":                          &PublicEvent{},
		"pull_request":                    &PullRequestEvent{},
		"pull_request_review":             &PullRequestReviewEvent{},
		"pull_request_review_comment":     &PullRequestReviewCommentEvent{},
		"pull_request_review_thread":      &PullRequestReviewThreadEvent{},
		"pull_request_target":             &PullRequestTargetEvent{},
		"push":                            &PushEvent{},
		"registry_package":                &RegistryPackageEvent{},
		"repository":                      &RepositoryEvent{},
		"repository_dispatch":             &RepositoryDispatchEvent{},
		"repository_import":               &RepositoryImportEvent{},
		"repository_ruleset":              &RepositoryRulesetEvent{},
		"repository_vulnerability_alert":  &RepositoryVulnerabilityAlertEvent{},
		"release":                         &ReleaseEvent{},
		"secret_scanning_alert":           &SecretScanningAlertEvent{},
		"secret_scanning_alert_location":  &SecretScanningAlertLocationEvent{},
		"security_advisory":               &SecurityAdvisoryEvent{},
		"security_and_analysis":           &SecurityAndAnalysisEvent{},
		"sponsorship":                     &SponsorshipEvent{},
		"star":                            &StarEvent{},
		"status":                          &StatusEvent{},
		"team":                            &TeamEvent{},
		"team_add":                        &TeamAddEvent{},
		"user":                            &UserEvent{},
		"watch":                           &WatchEvent{},
		"workflow_dispatch":               &WorkflowDispatchEvent{},
		"workflow_job":                    &WorkflowJobEvent{},
		"workflow_run":                    &WorkflowRunEvent{},
	}

	messageToTypeName = make(map[string]string, len(eventTypeMapping))

	typeToMessageMapping = make(map[string]string, len(eventTypeMapping))
)

func init() {
	for k, v := range eventTypeMapping {
		typename := reflect.TypeOf(v).Elem().Name()
		messageToTypeName[k] = typename
		typeToMessageMapping[typename] = k
	}
}

func genMAC(message, key []byte, hashFunc func() hash.Hash) []byte {
	_ = "STUB: not implemented"
	return nil
}

func checkMAC(message, messageMAC, key []byte, hashFunc func() hash.Hash) bool {
	_ = "STUB: not implemented"
	return false
}

func readPayloadBody(readable io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func messageMAC(signature string) ([]byte, func() hash.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ValidatePayloadFromBody(contentType string, readable io.Reader, signature string, secretToken []byte) (payload []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidatePayload(r *http.Request, secretToken []byte) (payload []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ValidateSignature(signature string, payload, secretToken []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func WebHookType(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func DeliveryID(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func ParseWebHook(messageType string, payload []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func MessageTypes() []string { _ = "STUB: not implemented"; return nil }

func EventForType(messageType string) any { _ = "STUB: not implemented"; return *new(any) }
