package github

import (
	"encoding/json"
)

type Attestation struct {
	Bundle       json.RawMessage `json:"bundle"`
	RepositoryID int64           `json:"repository_id"`
}

type AttestationsResponse struct {
	Attestations []*Attestation `json:"attestations"`
}
