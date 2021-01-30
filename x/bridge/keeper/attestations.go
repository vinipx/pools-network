package keeper

import (
	bridgeTypes "github.com/bloxapp/pools-network/x/bridge/types"
	poolTypes "github.com/bloxapp/pools-network/x/poolsnetwork/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

// ProcessAttestation process attestation via Pools Keeper delegator and operator
func (k Keeper) ProcessAttestation(ctx sdkTypes.Context, attestation *bridgeTypes.ClaimAttestation) error {
	claim := attestation.Claim
	switch claim.ClaimType {
	case bridgeTypes.ClaimType_Delegate:
		// 1. create delegator
		delegatorAccount := sdkTypes.AccAddress(claim.EthereumAddresses[0][:])
		k.PoolsKeeper.CreateDelegator(ctx, delegatorAccount, claim.Values[0])

		// 2. find operator
		operatorAddress := claim.EthereumAddresses[1]
		operator, found, err := k.PoolsKeeper.GetOperatorByEthereumAddress(ctx, operatorAddress)
		if !found {
			return bridgeTypes.ErrOperatorNotFound
		}
		if err != nil {
			return err
		}
		return k.PoolsKeeper.Delegate(ctx, delegatorAccount, operator, sdkTypes.NewIntFromUint64(claim.Values[0]))
	case bridgeTypes.ClaimType_Undelegate:
		return nil
	case bridgeTypes.ClaimType_CreateOperator:
		//1. Prepare operator's addresses and keys
		operatorAddress := claim.EthereumAddresses[0]
		consensusAddress := claim.ConsensusAddresses[0]
		publicKey := ed25519.GenPrivKey().PubKey() //TODO: find how to get Consensus Pk from ConsensusAddress?
		encoded, err := sdkTypes.Bech32ifyPubKey(sdkTypes.Bech32PubKeyTypeConsPub, publicKey)
		if err != nil {
			return sdkerrors.Wrap(err, "could not encode Consensus public key")
		}
		//2. Structure Operator data
		operator := poolTypes.Operator{
			EthereumAddress:  operatorAddress,
			ConsensusAddress: consensusAddress,
			ConsensusPk:      encoded,
		}
		//3. Perform CreateOperator function via Keeper
		if err := k.PoolsKeeper.CreateOperator(ctx, operator); err != nil {
			return sdkerrors.Wrap(err, "could not create operator")
		}
		return nil
	default:
		return bridgeTypes.ErrUnsupportedClaim
	}
}

// AttestClaim checks claim's validity and return its attestation
func (k Keeper) AttestClaim(
	ctx sdkTypes.Context,
	operator poolTypes.Operator,
	contract bridgeTypes.EthereumBridgeContact,
	claim bridgeTypes.ClaimData,
) (*bridgeTypes.ClaimAttestation, error) {
	att, found, err := k.GetAttestation(ctx, contract, claim)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "could not get attestation")
	}

	if !found {
		att = &bridgeTypes.ClaimAttestation{
			Claim:            claim,
			Contract:         contract,
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
	thresholdPower := k.PoolsKeeper.GetLastTotalPower(ctx) * 2 / 3
	if att.AccumulatedPower >= thresholdPower {
		att.Finalized = true
	}

	if err := k.SaveAttestation(ctx, att); err != nil {
		return nil, sdkerrors.Wrap(err, "could not save attestation")
	}
	return att, nil
}

// SaveAttestation persists claim attestation to the store
func (k Keeper) SaveAttestation(
	ctx sdkTypes.Context,
	claimAttestation *bridgeTypes.ClaimAttestation,
) error {
	byts, err := claimAttestation.Marshal()
	if err != nil {
		return sdkerrors.Wrap(err, "Could not marshal claim attestation")
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(bridgeTypes.GetClaimAttestationStoreKey(claimAttestation.Contract, claimAttestation.Claim), byts)
	return nil
}

// GetAttestation returns claim's attestation from the store
func (k Keeper) GetAttestation(
	ctx sdkTypes.Context,
	contract bridgeTypes.EthereumBridgeContact,
	claim bridgeTypes.ClaimData,
) (*bridgeTypes.ClaimAttestation, bool, error) {
	store := ctx.KVStore(k.storeKey)
	byts := store.Get(bridgeTypes.GetClaimAttestationStoreKey(contract, claim))
	if byts == nil || len(byts) == 0 {
		return nil, false, nil
	}

	// unmarshal
	ret := &bridgeTypes.ClaimAttestation{}
	if err := ret.Unmarshal(byts); err != nil {
		return nil, false, err
	}
	return ret, true, nil
}
