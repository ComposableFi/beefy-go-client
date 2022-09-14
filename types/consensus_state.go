package types

import (
	types1 "github.com/cosmos/ibc-go/v5/modules/core/23-commitment/types"

	"github.com/ComposableFi/ics11-beefy/exported"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	ibcexported "github.com/cosmos/ibc-go/v5/modules/core/exported"
)

var _ ibcexported.ConsensusState = (*ConsensusState)(nil)

// ClientType returns Beefy
func (ConsensusState) ClientType() string {
	return exported.Beefy
}

// GetRoot returns the commitment Root for the specific
func (cs ConsensusState) GetRoot() ibcexported.Root {
	return types1.MerkleRoot{Hash: cs.Root}
}

// GetTimestamp returns block time in nanoseconds of the header that created consensus state
func (cs ConsensusState) GetTimestamp() uint64 {
	return uint64(cs.Timestamp.UnixNano())
}

// ValidateBasic defines a basic validation for the beefy consensus state.
func (cs ConsensusState) ValidateBasic() error {
	if len(cs.Root) == 0 {
		return sdkerrors.Wrap(clienttypes.ErrInvalidConsensus, "root cannot be empty")
	}

	if cs.GetTimestamp() <= 0 {
		return sdkerrors.Wrap(clienttypes.ErrInvalidConsensus, "timestamp must be a positive Unix time")
	}
	return nil
}
