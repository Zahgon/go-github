package github

type PullRequestThread struct {
	ID       *int64                `json:"id,omitempty"`
	NodeID   *string               `json:"node_id,omitempty"`
	Comments []*PullRequestComment `json:"comments,omitempty"`
}

func (p PullRequestThread) String() string { _ = "STUB: not implemented"; return "" }
