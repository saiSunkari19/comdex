package keeper

import (
	"github.com/comdex-official/comdex/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetDeposit(ctx sdk.Context, vaultID uint64, depositor sdk.AccAddress) (deposit types.Deposit, found bool) {

	var (
		store = k.Store(ctx)
		key   = types.DepositKey(vaultID, depositor)
		value = store.Get(key)
	)

	if value == nil {
		return deposit, false
	}

	k.cdc.MustUnmarshal(value, &deposit)
	return deposit, true

}

func (k Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {

	var (
		store = k.Store(ctx)
		key   = types.DepositKey(deposit.VaultID, sdk.AccAddress(deposit.Depositor))
		value = store.Get(key)
	)

	store.Set(key, value)
}

func (k Keeper) DeleteDeposit(ctx sdk.Context, vaultID uint64, depositor sdk.AccAddress) {

	var (
		store = k.Store(ctx)
		key   = types.DepositKey(vaultID, depositor)
	)

	store.Delete(key)
}
