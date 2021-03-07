package d_voting

import (
	"go.dedis.ch/dela/core/ordering/cosipbft/types"
	"go.dedis.ch/kyber/v3"
)

type Election struct {
	Name string								// Name of the election
	AdminID string							// ID of the admin
	VotersID []string						// IDs of the voters
	Candidates []string						// List of candidates

	CollectivePublicKey kyber.Point			// DKG public key
	ElectionID types.Digest					// The digest of the first block of the election
}

type Ballot struct {
	UserID string							// ID of the user
	// ELGamal ciphertext pair
	K kyber.Point
	C kyber.Point
}

type ElectionResult map[string]uint32		// The result of the election, stored in a map <Candidate, number of voters>

// ---------------------------------------------------------------------------------------------------------------------
// API :

type OpenElection struct {
	Name string								// Name of the election
	AdminID string							// ID of the admin
	VotersID []string						// IDs of the voters
	Candidates []string						// List of candidates
}

type OpenElectionReply struct {
	ElectionID types.Digest 				// The digest of the first block of the election
	Success bool							// Indicates if the operation was successful
	Error string							// Error message
}

type CastVote struct {
	ElectionID types.Digest 				// The digest of the first block of the election
	Vote Ballot								// The ballot to be casted
}

type CastReply struct {
	BlockID types.Digest 					// The digest of the block storing the transaction
	Success bool							// Indicates if the operation was successful
	Error string							// Error message
}

type EndElection struct {
	ElectionID types.Digest 				// The digest of the first block of the election
	UserID string							// ID of the user
}

type EndElectionReply struct {
	Result ElectionResult					// The result of the election
	Success bool							// Indicates if the operation was successful
	Error string							// Error message
}
