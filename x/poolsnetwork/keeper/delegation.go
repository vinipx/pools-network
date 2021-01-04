package keeper

import (
	"github.com/bloxapp/pools-network/shared/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types3 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types2 "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// CreateDelegator creates an account with tokens that can be delegated to a validator.
// We generate just a random
func (k Keeper) CreateDelegator(ctx sdk.Context, address sdk.AccAddress, balance uint64) {
	// set staking power
	initCoins := sdk.NewCoins(sdk.NewCoin(k.StakingKeeper.BondDenom(ctx), sdk.NewIntFromUint64(balance)))

	// create account
	account := k.AccountKeeper.NewAccountWithAddress(ctx, address)
	k.AccountKeeper.SetAccount(ctx, account)
	k.BankKeeper.AddCoins(ctx, address, initCoins)

	// update supply
	prevSupply := k.BankKeeper.GetSupply(ctx)
	newSupply := prevSupply.GetTotal().Add(initCoins...)
	k.BankKeeper.SetSupply(ctx, types2.NewSupply(newSupply))
}

func (k Keeper) GetDelegator(ctx sdk.Context, address sdk.AccAddress) (types3.AccountI, uint64) {
	acc := k.AccountKeeper.GetAccount(ctx, address)
	balance := k.BankKeeper.GetBalance(ctx, address, k.StakingKeeper.BondDenom(ctx))
	return acc, balance.Amount.Uint64()
}

func (k Keeper) Delegate(ctx sdk.Context, from types.EthereumAddress, to types.ConsensusAddress, amount sdk.Int) error {
	return nil
}

//
