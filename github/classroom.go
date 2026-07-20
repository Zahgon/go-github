package github

import (
	"context"
)

type ClassroomService service

type ClassroomUser struct {
	ID        *int64  `json:"id,omitempty"`
	Login     *string `json:"login,omitempty"`
	AvatarURL *string `json:"avatar_url,omitempty"`
	HTMLURL   *string `json:"html_url,omitempty"`
}

func (u ClassroomUser) String() string { _ = "STUB: not implemented"; return "" }

type Classroom struct {
	ID           *int64        `json:"id,omitempty"`
	Name         *string       `json:"name,omitempty"`
	Archived     *bool         `json:"archived,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	URL          *string       `json:"url,omitempty"`
}

func (c Classroom) String() string { _ = "STUB: not implemented"; return "" }

type ClassroomAssignment struct {
	ID                          *int64      `json:"id,omitempty"`
	PublicRepo                  *bool       `json:"public_repo,omitempty"`
	Title                       *string     `json:"title,omitempty"`
	Type                        *string     `json:"type,omitempty"`
	InviteLink                  *string     `json:"invite_link,omitempty"`
	InvitationsEnabled          *bool       `json:"invitations_enabled,omitempty"`
	Slug                        *string     `json:"slug,omitempty"`
	StudentsAreRepoAdmins       *bool       `json:"students_are_repo_admins,omitempty"`
	FeedbackPullRequestsEnabled *bool       `json:"feedback_pull_requests_enabled,omitempty"`
	MaxTeams                    *int        `json:"max_teams,omitempty"`
	MaxMembers                  *int        `json:"max_members,omitempty"`
	Editor                      *string     `json:"editor,omitempty"`
	Accepted                    *int        `json:"accepted,omitempty"`
	Submitted                   *int        `json:"submitted,omitempty"`
	Passing                     *int        `json:"passing,omitempty"`
	Language                    *string     `json:"language,omitempty"`
	Deadline                    *Timestamp  `json:"deadline,omitempty"`
	StarterCodeRepository       *Repository `json:"starter_code_repository,omitempty"`
	Classroom                   *Classroom  `json:"classroom,omitempty"`
}

func (a ClassroomAssignment) String() string { _ = "STUB: not implemented"; return "" }

type AcceptedAssignment struct {
	ID          *int64               `json:"id,omitempty"`
	Submitted   *bool                `json:"submitted,omitempty"`
	Passing     *bool                `json:"passing,omitempty"`
	CommitCount *int                 `json:"commit_count,omitempty"`
	Grade       *string              `json:"grade,omitempty"`
	Students    []*ClassroomUser     `json:"students,omitempty"`
	Repository  *Repository          `json:"repository,omitempty"`
	Assignment  *ClassroomAssignment `json:"assignment,omitempty"`
}

func (a AcceptedAssignment) String() string { _ = "STUB: not implemented"; return "" }

type AssignmentGrade struct {
	AssignmentName        *string    `json:"assignment_name,omitempty"`
	AssignmentURL         *string    `json:"assignment_url,omitempty"`
	StarterCodeURL        *string    `json:"starter_code_url,omitempty"`
	GithubUsername        *string    `json:"github_username,omitempty"`
	RosterIdentifier      *string    `json:"roster_identifier,omitempty"`
	StudentRepositoryName *string    `json:"student_repository_name,omitempty"`
	StudentRepositoryURL  *string    `json:"student_repository_url,omitempty"`
	SubmissionTimestamp   *Timestamp `json:"submission_timestamp,omitempty"`
	PointsAwarded         *int       `json:"points_awarded,omitempty"`
	PointsAvailable       *int       `json:"points_available,omitempty"`
	GroupName             *string    `json:"group_name,omitempty"`
}

func (g AssignmentGrade) String() string { _ = "STUB: not implemented"; return "" }

//meta:operation GET /assignments/{assignment_id}
func (s *ClassroomService) GetAssignment(ctx context.Context, assignmentID int64) (*ClassroomAssignment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /classrooms/{classroom_id}
func (s *ClassroomService) GetClassroom(ctx context.Context, classroomID int64) (*Classroom, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /classrooms
func (s *ClassroomService) ListClassrooms(ctx context.Context, opts *ListOptions) ([]*Classroom, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /classrooms/{classroom_id}/assignments
func (s *ClassroomService) ListClassroomAssignments(ctx context.Context, classroomID int64, opts *ListOptions) ([]*ClassroomAssignment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /assignments/{assignment_id}/accepted_assignments
func (s *ClassroomService) ListAcceptedAssignments(ctx context.Context, assignmentID int64, opts *ListOptions) ([]*AcceptedAssignment, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /assignments/{assignment_id}/grades
func (s *ClassroomService) GetAssignmentGrades(ctx context.Context, assignmentID int64) ([]*AssignmentGrade, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
