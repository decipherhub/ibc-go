package keeper

import (
	"context"
	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/types"
)

var _ types.QueryServer = Keeper{}

// CrossChainQueryResult implements the Query/CrossChainQueryResult gRPC method
func (k Keeper) CrossChainQueryResult(context context.Context, query *types.QueryCrossChainQueryResult) (*types.QueryCrossChainQueryResultResponse, error) {
	// TODO
	// get queryResult from private store
	return nil, nil
}
