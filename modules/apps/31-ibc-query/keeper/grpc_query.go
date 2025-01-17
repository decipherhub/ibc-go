package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// CrossChainQueryResult implements the Query/CrossChainQueryResult gRPC method
func (k Keeper) CrossChainQueryResult(c context.Context, req *types.QueryCrossChainQueryResult) (*types.QueryCrossChainQueryResultResponse, error) {
	if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }
    ctx := sdk.UnwrapSDKContext(c)
    queryResult, found := k.GetCrossChainQueryResult(ctx, req.Id)
    if !found {
        return nil, sdkerrors.ErrNotFound
    }

    return &types.QueryCrossChainQueryResultResponse{Result: queryResult.Result, Data: queryResult.Data}, nil

}
