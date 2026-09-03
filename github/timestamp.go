package github

import (
	"time"
)

type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string { _ = "STUB: not implemented"; return "" }

func (t *Timestamp) GetTime() *time.Time { _ = "STUB: not implemented"; return nil }

func (t *Timestamp) UnmarshalJSON(data []byte) (err error) { _ = "STUB: not implemented"; return nil }

func (t Timestamp) Equal(u Timestamp) bool { _ = "STUB: not implemented"; return false }
