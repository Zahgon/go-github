package github

import (
	"context"
)

//meta:operation GET /events
func (s *ActivityService) ListEvents(ctx context.Context, opts *ListOptions) ([]*Event, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/events
func (s *ActivityService) ListRepositoryEvents(ctx context.Context, owner, repo string, opts *ListOptions) ([]*Event, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /repos/{owner}/{repo}/issues/events
func (s *ActivityService) ListIssueEventsForRepository(ctx context.Context, owner, repo string, opts *ListOptions) ([]*IssueEvent, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /networks/{owner}/{repo}/events
func (s *ActivityService) ListEventsForRepoNetwork(ctx context.Context, owner, repo string, opts *ListOptions) ([]*Event, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /orgs/{org}/events
func (s *ActivityService) ListEventsForOrganization(ctx context.Context, org string, opts *ListOptions) ([]*Event, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/events
//meta:operation GET /users/{username}/events/public
func (s *ActivityService) ListEventsPerformedByUser(ctx context.Context, user string, publicOnly bool, opts *ListOptions) ([]*Event, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/received_events
//meta:operation GET /users/{username}/received_events/public
func (s *ActivityService) ListEventsReceivedByUser(ctx context.Context, user string, publicOnly bool, opts *ListOptions) ([]*Event, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

//meta:operation GET /users/{username}/events/orgs/{org}
func (s *ActivityService) ListUserEventsForOrganization(ctx context.Context, org, user string, opts *ListOptions) ([]*Event, *Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
