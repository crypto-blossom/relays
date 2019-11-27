package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		// case types.MsgSetLink:
		// 	return handleMsgSetLink(ctx, keeper, msg)
		case types.MsgIngestHeaderChain:
			return handleMsgIngestHeaderChain(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgIngestHeaderChain(ctx sdk.Context, keeper Keeper, msg types.MsgIngestHeaderChain) {
	err := k.IngestHeaderChain(ctx, msg.Headers)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{}
}
