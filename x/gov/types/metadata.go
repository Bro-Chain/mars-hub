package types

import (
	"encoding/json"
)

// ProposalMetadata defines the required schema for proposal metadata.
type ProposalMetadata struct {
	Title             string   `json:"title"`
	Authors           []string `json:"authors"`
	Summary           string   `json:"summary,omitempty"`
	Details           string   `json:"details"`
	ProposalForumURL  string   `json:"proposal_forum_url,omitempty"`
	VoteOptionContext string   `json:"vote_option_context,omitempty"`
}

// VoteMetadata defines the required schema for vote metadata.
type VoteMetadata struct {
	Justification string `json:"justification,omitempty"`
}

// UnmarshalProposalMetadata unmarshals a string into ProposalMetadata.
//
// Golang's JSON unmarshal function is doesn't check for missing fields.
// Instead, for example, if the "title" field here in ProposalMetadata is
// missing, the json.Unmarshal simply returns metadata.Title = "" instead of
// throwing an error.
//
// Here's the equivalent Rust code for comparison, which properly throws an
// error is a required field is missing:
// https://play.rust-lang.org/?version=stable&mode=debug&edition=2015&gist=0e2eadad38b7cd212962b1a0e7a6da44
//
// Therefore, we have to implement our own unmarshal function.
func UnmarshalProposalMetadata(metadataStr string) (*ProposalMetadata, error) {
	var metadata ProposalMetadata

	if err := json.Unmarshal([]byte(metadataStr), &metadata); err != nil {
		return nil, ErrInvalidMetadata.Wrap(err.Error())
	}

	if metadata.Title == "" {
		return nil, ErrInvalidMetadata.Wrap("missing field `title`")
	}

	if metadata.Authors == nil {
		return nil, ErrInvalidMetadata.Wrap("missing field `authors`")
	}

	if metadata.Details == "" {
		return nil, ErrInvalidMetadata.Wrap("missing field `details`")
	}

	return &metadata, nil
}

// UnmarshalVoteMetadata unmarshals a string into VoteMetdata.
//
// See the comments for UnmarshalProposalMetadata on why we need to define this
// function instead of using Go's native json.Unmarshal function.
func UnmarshalVoteMetadata(metadataStr string) (*VoteMetadata, error) {
	var metadata VoteMetadata

	if err := json.Unmarshal([]byte(metadataStr), &metadata); err != nil {
		return nil, ErrInvalidMetadata.Wrap(err.Error())
	}

	return &metadata, nil
}
