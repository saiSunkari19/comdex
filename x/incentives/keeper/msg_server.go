package keeper

import (
	"context"
	"github.com/comdex-official/comdex/x/cdp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServiceServer = msgServer{}

func (ms msgServer) MsgClaimRewards(context context.Context, msg *types.MsgCreateCDPRequest) (*types.MsgCreateCDPResponse, error) {
	//ctx := sdk.UnwrapSDKContext(context)
	return nil, nil
}
