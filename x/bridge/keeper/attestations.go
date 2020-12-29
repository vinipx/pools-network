package keeper

import (
	types2 "github.com/bloxapp/pools-network/x/bridge/types"
	types3 "github.com/bloxapp/pools-network/x/poolsnetwork/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) processAttestation(ctx sdk.Context, contract *types2.ClaimAttestation) error {
	return nil // TODO
}

func (k Keeper) attestClaim(ctx sdk.Context, operator types3.Operator, contract types2.EthereumBridgeContact, claim types2.ClaimData) (*types2.ClaimAttestation, error) {
	att, err := k.getAttestation(ctx, contract, claim)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could not get attestation")
	}

	if att == nil {
		att = &types2.ClaimAttestation{
			ClaimId:          types2.GetClaimAttestastionStoreKey(contract, claim),
			ContractAddress:  contract.ContractAddress,
			Votes:            make(map[string]bool),
			AccumulatedPower: 0,
			Finalized:        false,
		}
	}

	// add attestation
	if _, found := att.Votes[operator.ConsensusAddress.Hex()]; !found {
		att.Votes[operator.ConsensusAddress.Hex()] = true
		att.AccumulatedPower += operator.GetPower()
	}

	// If 2/3 of the total staking power voted, mark as finalized
	if att.AccumulatedPower*3 > k.PoolsKeeper.GetLastTotalPower(ctx)*2 {
		att.Finalized = true
	}

	if err := k.saveAttestation(ctx, att); err != nil {
		return nil, sdkerrors.Wrap(err, "could not save attestation")
	}
	return att, nil
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
