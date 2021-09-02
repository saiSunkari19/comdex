package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField = sdkerrors.Register(ModuleName, 101, "invalid field")
)
