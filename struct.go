package d_voting

import (
	"go.dedis.ch/dela/core/ordering/cosipbft/types"
	"go.dedis.ch/kyber/v3"
)

type ElectionResult map[string]uint32		// The result of the election, stored in a map <Candidate, number of voters>

type ElectionStatus uint32
const (
	OnGoing = 1								//Ballots can be casted
	Closed = 2								//The election has been closed by its admin
	Canceled = 3							//The election has been canceled by its admin
	//shuffle, ballots decryption ... ?
)

type Election struct {
	Name string								// Name of the election
	AdminID string							// ID of the admin
	VotersID []string						// IDs of the voters
	Candidates []string						// List of candidates

	CollectivePublicKey kyber.Point			// DKG public key
	ElectionID types.Digest					// The digest of the first block of the election
	Status ElectionStatus					// The status of the election
}

type Ballot struct {
	UserID string							// ID of the user
	// ELGamal ciphertext pair
	K kyber.Point
	C kyber.Point
}

// ---------------------------------------------------------------------------------------------------------------------
// API :

type Login struct {
}

type LoginReply struct {
	UserID string							// ID of the user
	Success bool							// Indicates if the operation was successful
	Error string							// Error message
}

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

type GetElectionInfo struct {
	ElectionID types.Digest 				// The digest of the first block of the election
	UserID string							// ID of the user
}

type GetElectionInfoReply struct {
	ElectionInfo Election
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

type GetElectionResult struct {
	ElectionID types.Digest 				// The digest of the first block of the election
	UserID string							// ID of the user
}

type GetElectionResultReply struct {
	Result ElectionResult					// The result of the election
	Success bool							// Indicates if the operation was successful
	Error string							// Error message
}

type CancelElection struct {
	ElectionID types.Digest 				// The digest of the first block of the election
}

type CancelElectionReply struct {
	Success bool							// Indicates if the operation was successful
	Error string							// Error message
}
