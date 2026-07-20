package github

import (
	"context"
)

type ProjectsService service

type ProjectV2ItemContentType string

const (
	ProjectV2ItemContentTypeDraftIssue  ProjectV2ItemContentType = "DraftIssue"
	ProjectV2ItemContentTypeIssue       ProjectV2ItemContentType = "Issue"
	ProjectV2ItemContentTypePullRequest ProjectV2ItemContentType = "PullRequest"
)

type ProjectV2StatusUpdate struct {
	ID            *int64     `json:"id,omitempty"`
	NodeID        *string    `json:"node_id,omitempty"`
	ProjectNodeID *string    `json:"project_node_id,omitempty"`
	Creator       *User      `json:"creator,omitempty"`
	CreatedAt     *Timestamp `json:"created_at,omitempty"`
	UpdatedAt     *Timestamp `json:"updated_at,omitempty"`

	Status     *string `json:"status,omitempty"`
	StartDate  *string `json:"start_date,omitempty"`
	TargetDate *string `json:"target_date,omitempty"`
	Body       *string `json:"body,omitempty"`
}

type ProjectV2DraftIssue struct {
	ID        *int64     `json:"id,omitempty"`
	NodeID    *string    `json:"node_id,omitempty"`
	Title     *string    `json:"title,omitempty"`
	Body      *string    `json:"body,omitempty"`
	User      *User      `json:"user,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
	UpdatedAt *Timestamp `json:"updated_at,omitempty"`
}

type ProjectV2 struct {
	ID                 *int64                 `json:"id,omitempty"`
	NodeID             *string                `json:"node_id,omitempty"`
	Owner              *User                  `json:"owner,omitempty"`
	Creator            *User                  `json:"creator,omitempty"`
	Title              *string                `json:"title,omitempty"`
	Description        *string                `json:"description,omitempty"`
	Public             *bool                  `json:"public,omitempty"`
	ClosedAt           *Timestamp             `json:"closed_at,omitempty"`
	CreatedAt          *Timestamp             `json:"created_at,omitempty"`
	UpdatedAt          *Timestamp             `json:"updated_at,omitempty"`
	DeletedAt          *Timestamp             `json:"deleted_at,omitempty"`
	Number             *int                   `json:"number,omitempty"`
	ShortDescription   *string                `json:"short_description,omitempty"`
	DeletedBy          *User                  `json:"deleted_by,omitempty"`
	State              *string                `json:"state,omitempty"`
	LatestStatusUpdate *ProjectV2StatusUpdate `json:"latest_status_update,omitempty"`
	IsTemplate         *bool                  `json:"is_template,omitempty"`

	URL                    *string `json:"url,omitempty"`
	HTMLURL                *string `json:"html_url,omitempty"`
	ColumnsURL             *string `json:"columns_url,omitempty"`
	OwnerURL               *string `json:"owner_url,omitempty"`
	Name                   *string `json:"name,omitempty"`
	Body                   *string `json:"body,omitempty"`
	OrganizationPermission *string `json:"organization_permission,omitempty"`
	Private                *bool   `json:"private,omitempty"`
}

func (p ProjectV2) String() string { _ = "STUB: not implemented"; return "" }

type ListProjectsPaginationOptions struct {
	Before string `url:"before,omitempty"`

	After string `url:"after,omitempty"`

	PerPage int `url:"per_page,omitempty"`
}

type ListProjectsOptions struct {
	ListProjectsPaginationOptions

	Query string `url:"q,omitempty"`
}

type ProjectV2TextContent struct {
	HTML *string `json:"html,omitempty"`
	Raw  *string `json:"raw,omitempty"`
}

type ProjectV2FieldOption struct {
	ID          *string               `json:"id,omitempty"`
	Color       *string               `json:"color,omitempty"`
	Description *ProjectV2TextContent `json:"description,omitempty"`
	Name        *ProjectV2TextContent `json:"name,omitempty"`
}

type ProjectV2FieldIteration struct {
	ID        *string               `json:"id,omitempty"`
	Title     *ProjectV2TextContent `json:"title,omitempty"`
	StartDate *string               `json:"start_date,omitempty"`
	Duration  *int                  `json:"duration,omitempty"`
}

type ProjectV2FieldConfiguration struct {
	Duration   *int                       `json:"duration,omitempty"`
	StartDay   *int                       `json:"start_day,omitempty"`
	Iterations []*ProjectV2FieldIteration `json:"iterations,omitempty"`
}

type ProjectV2ItemContent struct {
	Issue       *Issue               `json:"-"`
	PullRequest *PullRequest         `json:"-"`
	DraftIssue  *ProjectV2DraftIssue `json:"-"`
}

func (c ProjectV2ItemContent) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ProjectV2Item struct {
	ArchivedAt  *Timestamp                 `json:"archived_at,omitempty"`
	Content     *ProjectV2ItemContent      `json:"content,omitempty"`
	ContentType *ProjectV2ItemContentType  `json:"content_type,omitempty"`
	CreatedAt   *Timestamp                 `json:"created_at,omitempty"`
	Creator     *User                      `json:"creator,omitempty"`
	Fields      []*ProjectV2ItemFieldValue `json:"fields,omitempty"`
	ID          *int64                     `json:"id,omitempty"`
	ItemURL     *string                    `json:"item_url,omitempty"`
	NodeID      *string                    `json:"node_id,omitempty"`
	ProjectURL  *string                    `json:"project_url,omitempty"`
	UpdatedAt   *Timestamp                 `json:"updated_at,omitempty"`

	ProjectNodeID *string `json:"project_node_id,omitempty"`
	ContentNodeID *string `json:"content_node_id,omitempty"`
}

func (p *ProjectV2Item) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type ProjectV2Field struct {
	ID            *int64                       `json:"id,omitempty"`
	NodeID        *string                      `json:"node_id,omitempty"`
	Name          *string                      `json:"name,omitempty"`
	DataType      *string                      `json:"data_type,omitempty"`
	ProjectURL    *string                      `json:"project_url,omitempty"`
	Options       []*ProjectV2FieldOption      `json:"options,omitempty"`
	Configuration *ProjectV2FieldConfiguration `json:"configuration,omitempty"`
	CreatedAt     *Timestamp                   `json:"created_at,omitempty"`
	UpdatedAt     *Timestamp                   `json:"updated_at,omitempty"`
}

type ProjectV2ItemFieldValue struct {
	ID       *int64  `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	DataType *string `json:"data_type,omitempty"`

	Value any `json:"value,omitempty"`
}

//meta:operation GET /orgs/{org}/projectsV2
func (s *ProjectsService) ListOrganizationProjects(ctx context.Context, org string, opts *ListProjectsOptions) ([]*ProjectV2, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/projectsV2/{project_number}
func (s *ProjectsService) GetOrganizationProject(ctx context.Context, org string, projectNumber int) (*ProjectV2, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/projectsV2
func (s *ProjectsService) ListUserProjects(ctx context.Context, username string, opts *ListProjectsOptions) ([]*ProjectV2, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/projectsV2/{project_number}
func (s *ProjectsService) GetUserProject(ctx context.Context, username string, projectNumber int) (*ProjectV2, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/projectsV2/{project_number}/fields
func (s *ProjectsService) ListOrganizationProjectFields(ctx context.Context, org string, projectNumber int, opts *ListProjectsOptions) ([]*ProjectV2Field, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/projectsV2/{project_number}/fields
func (s *ProjectsService) ListUserProjectFields(ctx context.Context, user string, projectNumber int, opts *ListProjectsOptions) ([]*ProjectV2Field, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/projectsV2/{project_number}/fields/{field_id}
func (s *ProjectsService) GetOrganizationProjectField(ctx context.Context, org string, projectNumber int, fieldID int64) (*ProjectV2Field, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/projectsV2/{project_number}/fields/{field_id}
func (s *ProjectsService) GetUserProjectField(ctx context.Context, user string, projectNumber int, fieldID int64) (*ProjectV2Field, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ListProjectItemsOptions struct {
	ListProjectsOptions

	Fields []int64 `url:"fields,omitempty,comma"`
}

type GetProjectItemOptions struct {
	Fields []int64 `url:"fields,omitempty,comma"`
}

type AddProjectItemOptions struct {
	Type *ProjectV2ItemContentType `json:"type,omitempty"`
	ID   *int64                    `json:"id,omitempty"`
}

type UpdateProjectV2Field struct {
	ID int64 `json:"id"`

	Value any `json:"value"`
}

type UpdateProjectItemOptions struct {
	Archived *bool `json:"archived,omitempty"`

	Fields []*UpdateProjectV2Field `json:"fields,omitempty"`
}

//meta:operation GET /orgs/{org}/projectsV2/{project_number}/items
func (s *ProjectsService) ListOrganizationProjectItems(ctx context.Context, org string, projectNumber int, opts *ListProjectItemsOptions) ([]*ProjectV2Item, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /orgs/{org}/projectsV2/{project_number}/items
func (s *ProjectsService) AddOrganizationProjectItem(ctx context.Context, org string, projectNumber int, body *AddProjectItemOptions) (*ProjectV2Item, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/projectsV2/{project_number}/items/{item_id}
func (s *ProjectsService) GetOrganizationProjectItem(ctx context.Context, org string, projectNumber int, itemID int64, opts *GetProjectItemOptions) (*ProjectV2Item, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /orgs/{org}/projectsV2/{project_number}/items/{item_id}
func (s *ProjectsService) UpdateOrganizationProjectItem(ctx context.Context, org string, projectNumber int, itemID int64, body *UpdateProjectItemOptions) (*ProjectV2Item, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /orgs/{org}/projectsV2/{project_number}/items/{item_id}
func (s *ProjectsService) DeleteOrganizationProjectItem(ctx context.Context, org string, projectNumber int, itemID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//meta:operation GET /users/{username}/projectsV2/{project_number}/items
func (s *ProjectsService) ListUserProjectItems(ctx context.Context, username string, projectNumber int, opts *ListProjectItemsOptions) ([]*ProjectV2Item, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /users/{username}/projectsV2/{project_number}/items
func (s *ProjectsService) AddUserProjectItem(ctx context.Context, username string, projectNumber int, body *AddProjectItemOptions) (*ProjectV2Item, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/projectsV2/{project_number}/items/{item_id}
func (s *ProjectsService) GetUserProjectItem(ctx context.Context, username string, projectNumber int, itemID int64, opts *GetProjectItemOptions) (*ProjectV2Item, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation PATCH /users/{username}/projectsV2/{project_number}/items/{item_id}
func (s *ProjectsService) UpdateUserProjectItem(ctx context.Context, username string, projectNumber int, itemID int64, body *UpdateProjectItemOptions) (*ProjectV2Item, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation DELETE /users/{username}/projectsV2/{project_number}/items/{item_id}
func (s *ProjectsService) DeleteUserProjectItem(ctx context.Context, username string, projectNumber int, itemID int64) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CreateProjectV2DraftItemRequest struct {
	Title string `json:"title"`

	Body *string `json:"body,omitempty"`
}

//meta:operation POST /orgs/{org}/projectsV2/{project_number}/drafts
func (s *ProjectsService) CreateOrganizationProjectDraftItem(ctx context.Context, org string, projectNumber int, body CreateProjectV2DraftItemRequest) (*ProjectV2Item, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /user/{user_id}/projectsV2/{project_number}/drafts
func (s *ProjectsService) CreateUserProjectDraftItem(ctx context.Context, userID int64, projectNumber int, body CreateProjectV2DraftItemRequest) (*ProjectV2Item, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ProjectV2FieldSingleSelectOption struct {
	Name string `json:"name"`

	Color *string `json:"color,omitempty"`

	Description *string `json:"description,omitempty"`
}

type ProjectV2FieldIterationConfiguration struct {
	StartDate *string `json:"start_date,omitempty"`

	Duration *int `json:"duration,omitempty"`

	Iterations []*ProjectV2FieldIterationConfigurationIteration `json:"iterations,omitempty"`
}

type ProjectV2FieldIterationConfigurationIteration struct {
	Title *string `json:"title,omitempty"`

	StartDate *string `json:"start_date,omitempty"`

	Duration *int `json:"duration,omitempty"`
}

type AddProjectV2FieldRequest struct {
	Name *string `json:"name,omitempty"`

	DataType *string `json:"data_type,omitempty"`

	SingleSelectOptions []*ProjectV2FieldSingleSelectOption `json:"single_select_options,omitempty"`

	IterationConfiguration *ProjectV2FieldIterationConfiguration `json:"iteration_configuration,omitempty"`

	IssueFieldID *int64 `json:"issue_field_id,omitempty"`
}

//meta:operation POST /orgs/{org}/projectsV2/{project_number}/fields
func (s *ProjectsService) AddOrganizationProjectField(ctx context.Context, org string, projectNumber int, body AddProjectV2FieldRequest) (*ProjectV2Field, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /users/{username}/projectsV2/{project_number}/fields
func (s *ProjectsService) AddUserProjectField(ctx context.Context, username string, projectNumber int, body AddProjectV2FieldRequest) (*ProjectV2Field, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type ProjectV2View struct {
	ID              int64                  `json:"id"`
	Number          int                    `json:"number"`
	Name            string                 `json:"name"`
	Layout          string                 `json:"layout"`
	NodeID          string                 `json:"node_id"`
	ProjectURL      string                 `json:"project_url"`
	HTMLURL         string                 `json:"html_url"`
	Creator         User                   `json:"creator"`
	Filter          *string                `json:"filter,omitempty"`
	VisibleFields   []int64                `json:"visible_fields"`
	SortBy          []*ProjectV2ViewSortBy `json:"sort_by"`
	GroupBy         []int64                `json:"group_by"`
	VerticalGroupBy []int64                `json:"vertical_group_by"`
	CreatedAt       Timestamp              `json:"created_at"`
	UpdatedAt       Timestamp              `json:"updated_at"`
}

type ProjectV2ViewSortBy struct {
	FieldID *int64 `json:"-"`

	Direction *string `json:"-"`
}

func (s *ProjectV2ViewSortBy) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ProjectV2ViewSortBy) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type CreateProjectV2ViewRequest struct {
	Name string `json:"name"`

	Layout string `json:"layout"`

	Filter *string `json:"filter,omitempty"`

	VisibleFields []int64 `json:"visible_fields,omitempty"`
}

//meta:operation POST /orgs/{org}/projectsV2/{project_number}/views
func (s *ProjectsService) CreateOrganizationProjectView(ctx context.Context, org string, projectNumber int, body CreateProjectV2ViewRequest) (*ProjectV2View, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation POST /users/{user_id}/projectsV2/{project_number}/views
func (s *ProjectsService) CreateUserProjectView(ctx context.Context, userID int64, projectNumber int, body CreateProjectV2ViewRequest) (*ProjectV2View, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/projectsV2/{project_number}/views/{view_number}/items
func (s *ProjectsService) ListOrganizationProjectViewItems(ctx context.Context, org string, projectNumber, viewNumber int, opts *ListProjectItemsOptions) ([]*ProjectV2Item, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/projectsV2/{project_number}/views/{view_number}/items
func (s *ProjectsService) ListUserProjectViewItems(ctx context.Context, username string, projectNumber, viewNumber int, opts *ListProjectItemsOptions) ([]*ProjectV2Item, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
