package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/types"
)

// InitGenesis initializes the application state from a provided genesis state
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	for _, query := range state.Queries {
		k.SetSubmitCrossChainQuery(ctx, *query)
	}
	for _, result := range state.Results {
		k.SetSubmitCrossChainQueryResult(ctx, *result)
	}
}

// ExportGenesis returns the application exported genesis
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Queries: k.GetAllSubmitCrossChainQueries(ctx),
		Results: k.GetAllSubmitCrossChainQueryResults(ctx),
	}
}
