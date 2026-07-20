//go:generate go run gen-accessors.go
//go:generate go run gen-iterators.go
//go:generate go run gen-stringify-test.go
//go:generate sh ../script/metadata.sh update-go

package github

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"
)

const (
	Version = "v89.0.0"

	HeaderRateLimit     = "X-Ratelimit-Limit"
	HeaderRateRemaining = "X-Ratelimit-Remaining"
	HeaderRateReset     = "X-Ratelimit-Reset"
	HeaderRateResource  = "X-Ratelimit-Resource"
	HeaderRateUsed      = "X-Ratelimit-Used"
	HeaderRequestID     = "X-Github-Request-Id"

	api20221128 = "2022-11-28"
	api20260310 = "2026-03-10"

	defaultBaseURL   = "https://api.github.com/"
	defaultUserAgent = "go-github" + "/" + Version
	uploadBaseURL    = "https://uploads.github.com/"

	headerAPIVersion = "X-Github-Api-Version"
	headerOTP        = "X-Github-Otp"
	headerRetryAfter = "Retry-After"

	headerTokenExpiration = "Github-Authentication-Token-Expiration"

	mediaTypeV3                = "application/vnd.github.v3+json"
	defaultMediaType           = "application/octet-stream"
	mediaTypeV3SHA             = "application/vnd.github.v3.sha"
	mediaTypeV3Diff            = "application/vnd.github.v3.diff"
	mediaTypeV3Patch           = "application/vnd.github.v3.patch"
	mediaTypeOrgPermissionRepo = "application/vnd.github.v3.repository+json"
	mediaTypeIssueImportAPI    = "application/vnd.github.golden-comet-preview+json"
	mediaTypeStarring          = "application/vnd.github.star+json"
	mediaTypeSCIM              = "application/scim+json"

	mediaTypeMigrationsPreview = "application/vnd.github.wyandotte-preview+json"

	mediaTypeDeploymentStatusPreview = "application/vnd.github.ant-man-preview+json"

	mediaTypeExpandDeploymentStatusPreview = "application/vnd.github.flash-preview+json"

	mediaTypeReactionsPreview = "application/vnd.github.squirrel-girl-preview"

	mediaTypeTimelinePreview = "application/vnd.github.mockingbird-preview+json"

	mediaTypeProjectsPreview = "application/vnd.github.inertia-preview+json"

	mediaTypeCommitSearchPreview = "application/vnd.github.cloak-preview+json"

	mediaTypeBlockUsersPreview = "application/vnd.github.giant-sentry-fist-preview+json"

	mediaTypeCodesOfConductPreview = "application/vnd.github.scarlet-witch-preview+json"

	mediaTypeTopicsPreview = "application/vnd.github.mercy-preview+json"

	mediaTypeRequiredApprovingReviewsPreview = "application/vnd.github.luke-cage-preview+json"

	mediaTypeCheckRunsPreview = "application/vnd.github.antiope-preview+json"

	mediaTypePreReceiveHooksPreview = "application/vnd.github.eye-scream-preview"

	mediaTypeSignaturePreview = "application/vnd.github.zzzax-preview+json"

	mediaTypeProjectCardDetailsPreview = "application/vnd.github.starfox-preview+json"

	mediaTypeInteractionRestrictionsPreview = "application/vnd.github.sombra-preview+json"

	mediaTypeEnablePagesAPIPreview = "application/vnd.github.switcheroo-preview+json"

	mediaTypeRequiredVulnerabilityAlertsPreview = "application/vnd.github.dorian-preview+json"

	mediaTypeUpdatePullRequestBranchPreview = "application/vnd.github.lydian-preview+json"

	mediaTypeListPullsOrBranchesForCommitPreview = "application/vnd.github.groot-preview+json"

	mediaTypeMemberAllowedRepoCreationTypePreview = "application/vnd.github.surtur-preview+json"

	mediaTypeRepositoryTemplatePreview = "application/vnd.github.baptiste-preview+json"

	mediaTypeMultiLineCommentsPreview = "application/vnd.github.comfort-fade-preview+json"

	mediaTypeOAuthAppPreview = "application/vnd.github.doctor-strange-preview+json"

	mediaTypeRepositoryVisibilityPreview = "application/vnd.github.nebula-preview+json"

	mediaTypeContentAttachmentsPreview = "application/vnd.github.corsair-preview+json"
)

var ErrPathForbidden = errors.New("path must not contain '..' due to auth vulnerability issue")

type Client struct {
	client                *http.Client
	clientIgnoreRedirects *http.Client

	baseURL *url.URL

	uploadURL *url.URL

	apiVersionDefault string

	apiVersionMin string

	apiVersionMax string

	userAgent string

	disableRateLimitCheck bool

	rateMu                  sync.Mutex
	rateLimits              [Categories]Rate
	secondaryRateLimitReset time.Time

	maxSecondaryRateLimitRetryAfterDuration time.Duration

	rateLimitRedirectionalEndpoints bool

	common service

	Actions            *ActionsService
	Activity           *ActivityService
	Admin              *AdminService
	Apps               *AppsService
	Authorizations     *AuthorizationsService
	Billing            *BillingService
	Checks             *ChecksService
	Classroom          *ClassroomService
	CodeQuality        *CodeQualityService
	CodeScanning       *CodeScanningService
	CodesOfConduct     *CodesOfConductService
	Codespaces         *CodespacesService
	Copilot            *CopilotService
	Credentials        *CredentialsService
	Dependabot         *DependabotService
	DependencyGraph    *DependencyGraphService
	Emojis             *EmojisService
	Enterprise         *EnterpriseService
	Gists              *GistsService
	Git                *GitService
	Gitignores         *GitignoresService
	Interactions       *InteractionsService
	IssueImport        *IssueImportService
	Issues             *IssuesService
	Licenses           *LicensesService
	Markdown           *MarkdownService
	Marketplace        *MarketplaceService
	Meta               *MetaService
	Migrations         *MigrationService
	Organizations      *OrganizationsService
	PrivateRegistries  *PrivateRegistriesService
	Projects           *ProjectsService
	PullRequests       *PullRequestsService
	RateLimit          *RateLimitService
	Reactions          *ReactionsService
	Repositories       *RepositoriesService
	SCIM               *SCIMService
	Search             *SearchService
	SecretScanning     *SecretScanningService
	SecurityAdvisories *SecurityAdvisoriesService
	SubIssue           *SubIssueService
	Teams              *TeamsService
	Users              *UsersService
}

type service struct {
	client *Client
}

func (c *Client) Client() *http.Client { _ = "STUB: not implemented"; return nil }

type ListOptions struct {
	Page int `url:"page,omitempty"`

	PerPage int `url:"per_page,omitempty"`
}

type ListCursorOptions struct {
	Page string `url:"page,omitempty"`

	PerPage int `url:"per_page,omitempty"`

	First int `url:"first,omitempty"`

	Last int `url:"last,omitempty"`

	After string `url:"after,omitempty"`

	Before string `url:"before,omitempty"`

	Cursor string `url:"cursor,omitempty"`
}

type UploadOptions struct {
	Name      string `url:"name,omitempty"`
	Label     string `url:"label,omitempty"`
	MediaType string `url:"-"`
}

type RawType uint8

const (
	Diff RawType = 1 + iota

	Patch
)

type RawOptions struct {
	Type RawType
}

type structPtr[T any] interface{ *T }

func addOptions[P structPtr[T], T any](s string, opts P) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

var errUninitialized = errors.New("client is not initialized")

type clientOptions struct {
	httpClient                              *http.Client
	transport                               http.RoundTripper
	timeout                                 *time.Duration
	apiVersionMin                           *string
	apiVersionMax                           *string
	userAgent                               *string
	envProxy                                bool
	token                                   *string
	baseURL                                 *url.URL
	uploadURL                               *url.URL
	disableRateLimitCheck                   bool
	rateLimitRedirectionalEndpoints         bool
	maxSecondaryRateLimitRetryAfterDuration *time.Duration
	marketplaceStubbed                      bool
}

type ClientOptionsFunc func(*clientOptions) error

func WithHTTPClient(httpClient *http.Client) ClientOptionsFunc {
	_ = "STUB: not implemented"
	return *new(ClientOptionsFunc)
}

func WithTransport(transport http.RoundTripper) ClientOptionsFunc {
	_ = "STUB: not implemented"
	return *new(ClientOptionsFunc)
}

func WithTimeout(timeout time.Duration) ClientOptionsFunc {
	_ = "STUB: not implemented"
	return *new(ClientOptionsFunc)
}

func WithUserAgent(userAgent string) ClientOptionsFunc {
	_ = "STUB: not implemented"
	return *new(ClientOptionsFunc)
}

func WithEnvProxy() ClientOptionsFunc { _ = "STUB: not implemented"; return *new(ClientOptionsFunc) }

func WithAuthToken(token string) ClientOptionsFunc {
	_ = "STUB: not implemented"
	return *new(ClientOptionsFunc)
}

func WithURLs(baseURL, uploadURL *string) ClientOptionsFunc {
	_ = "STUB: not implemented"
	return *new(ClientOptionsFunc)
}

func WithEnterpriseURLs(baseURL, uploadURL string) ClientOptionsFunc {
	_ = "STUB: not implemented"
	return *new(ClientOptionsFunc)
}

func WithDisableRateLimitCheck() ClientOptionsFunc {
	_ = "STUB: not implemented"
	return *new(ClientOptionsFunc)
}

func WithRateLimitRedirectionalEndpoints() ClientOptionsFunc {
	_ = "STUB: not implemented"
	return *new(ClientOptionsFunc)
}

func WithMaxSecondaryRateLimitRetryAfterDuration(duration time.Duration) ClientOptionsFunc {
	_ = "STUB: not implemented"
	return *new(ClientOptionsFunc)
}

func NewClient(opts ...ClientOptionsFunc) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newClient(opts clientOptions) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Client) UserAgent() string { _ = "STUB: not implemented"; return "" }

func (c *Client) BaseURL() string { _ = "STUB: not implemented"; return "" }

func (c *Client) UploadURL() string { _ = "STUB: not implemented"; return "" }

func (c *Client) Clone(opts ...ClientOptionsFunc) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type RequestOption func(req *http.Request)

func WithVersion(version string) RequestOption {
	_ = "STUB: not implemented"
	return *new(RequestOption)
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body any, opts ...RequestOption) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) NewFormRequest(ctx context.Context, urlStr string, body io.Reader, opts ...RequestOption) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkURLPathTraversal(urlStr string) error { _ = "STUB: not implemented"; return nil }

func (c *Client) NewUploadRequest(ctx context.Context, urlStr string, reader io.Reader, size int64, mediaType string, opts ...RequestOption) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type uploadRequestBodyReader struct {
	io.Reader
}

type Response struct {
	*http.Response

	NextPage  int
	PrevPage  int
	FirstPage int
	LastPage  int

	NextPageToken string

	Cursor string

	Before string
	After  string

	Rate Rate

	TokenExpiration Timestamp
}

func newResponse(r *http.Response) *Response { _ = "STUB: not implemented"; return nil }

func (r *Response) populatePageValues() { _ = "STUB: not implemented"; return }

func parseRate(r *http.Response) Rate { _ = "STUB: not implemented"; return *new(Rate) }

func parseSecondaryRate(r *http.Response) *time.Duration { _ = "STUB: not implemented"; return nil }

func parseTokenExpiration(r *http.Response) Timestamp {
	_ = "STUB: not implemented"
	return *new(Timestamp)
}

type requestContext uint8

const (
	BypassRateLimitCheck requestContext = iota

	SleepUntilPrimaryRateLimitResetWhenRateLimited
)

const maxErrorBodySize = 1 * 1024 * 1024

var ErrUnsupportedAPIVersion = errors.New("unsupported api version")

func (c *Client) checkRequestAPIVersionBeforeDo(req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) bareDo(caller *http.Client, req *http.Request) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) BareDo(req *http.Request) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) bareDoIgnoreRedirects(req *http.Request) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var errInvalidLocation = errors.New("invalid or empty Location header in redirection response")

func (c *Client) bareDoUntilFound(req *http.Request, maxRedirects int) (*url.URL, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Client) Do(req *http.Request, v any) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) checkRateLimitBeforeDo(req *http.Request, rateLimitCategory RateLimitCategory) *RateLimitError {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) checkSecondaryRateLimitBeforeDo(req *http.Request) *AbuseRateLimitError {
	_ = "STUB: not implemented"
	return nil
}

func compareHTTPResponse(r1, r2 *http.Response) bool { _ = "STUB: not implemented"; return false }

type ErrorResponse struct {
	Response *http.Response `json:"-"`
	Message  string         `json:"message"`
	//nolint:sliceofpointers
	Errors []Error `json:"errors"`

	Block *ErrorBlock `json:"block,omitempty"`

	DocumentationURL string `json:"documentation_url,omitempty"`
}

type ErrorBlock struct {
	Reason    string     `json:"reason,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
}

func (r *ErrorResponse) Error() string { _ = "STUB: not implemented"; return "" }

func (r *ErrorResponse) Is(target error) bool { _ = "STUB: not implemented"; return false }

type TwoFactorAuthError ErrorResponse

func (r *TwoFactorAuthError) Error() string { _ = "STUB: not implemented"; return "" }

type RateLimitError struct {
	Rate     Rate
	Response *http.Response
	Message  string `json:"message"`
}

func (r *RateLimitError) Error() string { _ = "STUB: not implemented"; return "" }

func (r *RateLimitError) Is(target error) bool { _ = "STUB: not implemented"; return false }

type AcceptedError struct {
	Raw []byte
}

func (*AcceptedError) Error() string { _ = "STUB: not implemented"; return "" }

func (ae *AcceptedError) Is(target error) bool { _ = "STUB: not implemented"; return false }

type AbuseRateLimitError struct {
	Response *http.Response
	Message  string `json:"message"`

	RetryAfter *time.Duration
}

func (r *AbuseRateLimitError) Error() string { _ = "STUB: not implemented"; return "" }

func (r *AbuseRateLimitError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func equalDurationPtr(a, b *time.Duration) bool { _ = "STUB: not implemented"; return false }

type RedirectionError struct {
	Response   *http.Response
	StatusCode int
	Location   *url.URL
}

func (r *RedirectionError) Error() string { _ = "STUB: not implemented"; return "" }

func (r *RedirectionError) Is(target error) bool { _ = "STUB: not implemented"; return false }

var sensitiveParams = []string{"client_secret", "access_token", "token"}

func sanitizeURL(uri *url.URL) *url.URL { _ = "STUB: not implemented"; return nil }

type Error struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
	Message  string `json:"message"`
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func CheckResponse(r *http.Response) error { _ = "STUB: not implemented"; return nil }

func parseBoolResponse(err error) (bool, error) { _ = "STUB: not implemented"; return false, nil }

type RateLimitCategory uint8

const (
	CoreCategory RateLimitCategory = iota
	SearchCategory
	GraphqlCategory
	IntegrationManifestCategory
	SourceImportCategory
	CodeScanningUploadCategory
	ActionsRunnerRegistrationCategory
	ScimCategory
	DependencySnapshotsCategory
	CodeSearchCategory
	AuditLogCategory
	DependencySBOMCategory

	Categories
)

func GetRateLimitCategory(method, path string) RateLimitCategory {
	_ = "STUB: not implemented"
	return *new(RateLimitCategory)
}

func (c *Client) RateLimits(ctx context.Context) (*RateLimits, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func setCredentialsAsHeaders(req *http.Request, id, secret string) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

type UnauthenticatedRateLimitedTransport struct {
	ClientID string

	ClientSecret string

	Transport http.RoundTripper
}

func (t *UnauthenticatedRateLimitedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *UnauthenticatedRateLimitedTransport) Client() *http.Client {
	_ = "STUB: not implemented"
	return nil
}

func (t *UnauthenticatedRateLimitedTransport) transport() http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

type BasicAuthTransport struct {
	Username string
	Password string
	OTP      string

	Transport http.RoundTripper
}

func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *BasicAuthTransport) Client() *http.Client { _ = "STUB: not implemented"; return nil }

func (t *BasicAuthTransport) transport() http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

func formatRateReset(d time.Duration) string { _ = "STUB: not implemented"; return "" }

func sleepUntilResetWithBuffer(ctx context.Context, reset time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) roundTripWithOptionalFollowRedirect(ctx context.Context, u string, maxRedirects int, opts ...RequestOption) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) checkRedirectHost(location string) error { _ = "STUB: not implemented"; return nil }

func Ptr[T any](v T) *T { _ = "STUB: not implemented"; return nil }

//go:fix inline
func Bool(v bool) *bool { _ = "STUB: not implemented"; return nil }

//go:fix inline
func Int(v int) *int { _ = "STUB: not implemented"; return nil }

//go:fix inline
func Int64(v int64) *int64 { _ = "STUB: not implemented"; return nil }

//go:fix inline
func String(v string) *string { _ = "STUB: not implemented"; return nil }

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var runIDFromURLRE = regexp.MustCompile(`repos/.*/actions/runs/(\d+)/deployment_protection_rule$`)

func (e *DeploymentProtectionRuleEvent) GetRunID() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
