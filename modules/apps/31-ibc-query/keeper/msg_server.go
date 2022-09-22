package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/ibc-go/v3/modules/apps/31-ibc-query/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

//var _ types.MsgServer = Keeper{}

// SubmitCrossChainQuery Handling SubmitCrossChainQuery transaction
func (k Keeper) SubmitCrossChainQuery(goCtx context.Context, msg *types.MsgSubmitCrossChainQuery) (*types.MsgSubmitCrossChainQueryResponse, error) {
	// UnwrapSDKContext
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentTimestamp := uint64(ctx.BlockTime().UnixNano())
	currentHeight := clienttypes.GetSelfHeight(ctx)

	// Sanity-check that localTimeoutHeight is 0 or greater than the current height, otherwise the query will always time out.
	if !(msg.LocalTimeoutHeight.RevisionHeight == 0 || msg.LocalTimeoutHeight.RevisionHeight > currentHeight.RevisionHeight) {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidTimeoutHeight,
			"localTimeoutHeight is not 0 and current height >= localTimeoutHeight(%d >= %d)", currentHeight.RevisionHeight, msg.LocalTimeoutHeight,
		)
	}
	// Sanity-check that localTimeoutTimestamp is 0 or greater than the current timestamp, otherwise the query will always time out.
	if !(msg.LocalTimeoutStamp == 0 || msg.LocalTimeoutStamp > currentTimestamp) {
		return nil, sdkerrors.Wrapf(
			types.ErrQuerytTimeout,
			"localTimeoutTimestamp is not 0 and current timestamp >= localTimeoutTimestamp(%d >= %d)", currentTimestamp, msg.LocalTimeoutStamp,
		)
	}

	// call keeper function
	// keeper func save query in private store
	query := types.CrossChainQuery{
		Id:                    msg.Id,
		Path:                  msg.Path,
		LocalTimeoutHeight:    msg.LocalTimeoutHeight,
		LocalTimeoutTimestamp: msg.LocalTimeoutStamp,
		QueryHeight:           msg.QueryHeight,
		ClientId:              msg.Sender,
	}

	k.SetCrossChainQuery(ctx, query)

	// TODO
	// var data []byte
	// err := query.Unmarshal(data)
	// if err != nil {
	// 	return nil, err
	// }

	//if err := k.SendQuery(ctx, msg.SourcePort, msg.SourceChannel, query.GetSignBytes(),
	//	msg.LocalTimeoutHeight, msg.LocalTimeoutStamp); err != nil {
	//	return nil, err
	//}

	// Log the query request
	k.Logger(ctx).Info("query sent", "query_id", msg.GetQueryId())

	// emit event
	EmitQueryEvent(ctx, msg)

	// It can be used when event emit
	//if err := ctx.EventManager().EmitTypedEvent(
	//	types.NewEventQuerySubmitted(msg.Id, msg.Path, *msg.LocalTimeoutHeight, msg.LocalTimeoutStamp, msg.QueryHeight),
	//); err != nil {
	//	return nil, err
	//}

	return &types.MsgSubmitCrossChainQueryResponse{QueryId: query.Id}, nil
}
