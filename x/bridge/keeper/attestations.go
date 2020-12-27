package keeper

import (
	types2 "github.com/bloxapp/pools-network/x/bridge/types"
	types3 "github.com/bloxapp/pools-network/x/poolsnetwork/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) attestClaim(ctx sdk.Context, operator types3.Operator, contract types2.EthereumBridgeContact, claim types2.ClaimData) error {
	att, err := k.getAttestation(ctx, contract, claim)
	if err != nil {
		return err
	}

	if att == nil {
		att = &types2.ClaimAttestation{
			ClaimId:          types2.GetClaimAttestationStoreKey(contract, claim),
			Votes:            make(map[string]bool),
			AccumulatedPower: 0,
			Finalized:        false,
		}
	}

	if _, found := att.Votes[operator.ConsensusAddress.Hex()]; !found {
		att.Votes[operator.ConsensusAddress.Hex()] = true
		att.AccumulatedPower += operator.GetPower()
	}

	return k.saveAttestation(ctx, att)
}

func (k Keeper) saveAttestation(
	ctx sdk.Context,
	claimAttestation *types2.ClaimAttestation,
) error {
	byts, err := claimAttestation.Marshal()
	if err != nil {
		return sdkerrors.Wrap(err, "Could not marshal claim attestation")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(claimAttestation.ClaimId, byts)
	return nil
}

func (k Keeper) getAttestation(
	ctx sdk.Context,
	contract types2.EthereumBridgeContact,
	claim types2.ClaimData,
) (*types2.ClaimAttestation, error) {
	store := ctx.KVStore(k.storeKey)
	byts := store.Get(types2.GetClaimAttestationStoreKey(contract, claim))
	if byts == nil || len(byts) == 0 {
		return nil, nil
	}

	// unmarshal
	ret := &types2.ClaimAttestation{}
	if err := ret.Unmarshal(byts); err != nil {
		return nil, err
	}
	return ret, nil
}
