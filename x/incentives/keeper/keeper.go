package keeper

import (
	"github.com/comdex-official/comdex/x/incentives/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	cdc           codec.BinaryMarshaler
	storeKey      sdk.StoreKey
	memKey        sdk.StoreKey
	paramSpace    types.ParamSubspace
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}
