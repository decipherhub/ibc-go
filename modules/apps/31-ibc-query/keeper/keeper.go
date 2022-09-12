package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper define 31-ibc-query keeper
type Keeper struct {
	storeKey     sdk.StoreKey
	cdc          codec.BinaryCodec
	paramSpace   paramtypes.Subspace
	scopedKeeper capabilitykeeper.ScopedKeeper

	ics4Wrapper   types.ICS4Wrapper
	channelKeeper types.ChannelKeeper
	portKeeper    types.PortKeeper
}

// NewKeeper creates a new 31-ibc-query Keeper instance
func NewKeeper(cdc codec.BinaryCodec,
	key sdk.StoreKey,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	ics4Wrapper types.ICS4Wrapper,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		scopedKeeper:  scopedKeeper,
		ics4Wrapper:   ics4Wrapper,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+host.ModuleName+"-"+types.ModuleName)
}


// IsBound checks if the IBC query module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the ort Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the transfer module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the transfer module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
}

func (k Keeper) GenerateQueryIdentifier(ctx sdk.Context) (string, error) {
	nextQuerySeq := k.GetNextQuerySequence(ctx)
	if nextQuerySeq == 0 {
		return "", fmt.Errorf("next query sequence is nil")
	}
	queryID := types.FormatQueryIdentifier(nextQuerySeq)

	nextQuerySeq++
	k.SetNextQuerySequence(ctx, nextQuerySeq)
	return queryID, nil
}

func (k Keeper) GetNextQuerySequence(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(types.KeyNextQuerySequence))
	if bz == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetNextQuerySequence(ctx sdk.Context, sequence uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := sdk.Uint64ToBigEndian(sequence)
	store.Set([]byte(types.KeyNextQuerySequence), bz)
}

// SetSubmitCrossChainQuery stores the MsgSubmitCrossChainQuery in state keyed by the query id
func (k Keeper) SetSubmitCrossChainQuery(ctx sdk.Context, query types.MsgSubmitCrossChainQuery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.QueryKey)
	bz := k.MustMarshalQuery(&query)
	store.Set(host.QueryKey(query.Id), bz)
}

// GetSubmitCrossChainQuery retrieve the MsgSubmitCrossChainQuery stored in state given the query id
func (k Keeper) GetSubmitCrossChainQuery(ctx sdk.Context, queryId string) (types.MsgSubmitCrossChainQuery, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.QueryKey)
	key := host.QueryKey(queryId)
	bz := store.Get(key)
	if bz == nil {
		return types.MsgSubmitCrossChainQuery{}, false
	}

	return k.MustUnmarshalQuery(bz), true
}

// GetAllSubmitCrossChainQueries returns a list of all MsgSubmitCrossChainQueries that are stored in state
func (k Keeper) GetAllSubmitCrossChainQueries(ctx sdk.Context) []*types.MsgSubmitCrossChainQuery {
	var crossChainQueries []*types.MsgSubmitCrossChainQuery
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.QueryKey)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		// unmarshal
		query := k.MustUnmarshalQuery(iterator.Value())
		crossChainQueries = append(crossChainQueries, &query)
	}
	return crossChainQueries
}

// DeleteCrossChainQuery deletes MsgSubmitCrossChainQuery associated with the query id
func (k Keeper) DeleteSubmitCrossChainQuery(ctx sdk.Context, queryId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.QueryKey)
	store.Delete(host.QueryKey(queryId))
}

// TODO
// func handleIbcQueryResult
// 1. relayer should be call this func with query result
// 2. save query in private store

// SetCrossChainQueryResult stores the CrossChainQueryResult in state keyed by the query id
func (k Keeper) SetSubmitCrossChainQueryResult(ctx sdk.Context, result types.MsgSubmitCrossChainQueryResult) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.QueryResultKey)
	bz := k.MustMarshalQueryResult(&result)
	store.Set(host.QueryResultKey(result.Id), bz)
}

// GetCrossChainQueryResult retrieve the CrossChainQueryResult stored in state given the query id
func (k Keeper) GetCrossChainQueryResult(ctx sdk.Context, queryId string) (types.MsgSubmitCrossChainQueryResult, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.QueryResultKey)
	key := host.QueryResultKey(queryId)
	bz := store.Get(key)
	if bz == nil {
		return types.MsgSubmitCrossChainQueryResult{}, false
	}

	return k.MustUnmarshalQueryResult(bz), true
}

// GetAllCrossChainQueryResults returns a list of all CrossChainQueryResults that are stored in state
func (k Keeper) GetAllSubmitCrossChainQueryResults(ctx sdk.Context) []*types.MsgSubmitCrossChainQueryResult {
	var crossChainQueryResults []*types.MsgSubmitCrossChainQueryResult
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.QueryResultKey)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		// unmarshal
		result := k.MustUnmarshalQueryResult(iterator.Value())
		crossChainQueryResults = append(crossChainQueryResults, &result)
	}
	return crossChainQueryResults
}

// DeleteCrossChainQueryResult deletes CrossChainQueryResult associated with the query id
func (k Keeper) DeleteSubmitCrossChainQueryResult(ctx sdk.Context, queryId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.QueryResultKey)
	store.Delete(host.QueryResultKey(queryId))
}

// MustMarshalQuery attempts to encode a CrossChainQuery object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalQuery(query *types.MsgSubmitCrossChainQuery) []byte {
	return k.cdc.MustMarshal(query)
}

// MustUnmarshalQuery attempts to decode and return a CrossChainQuery object from
// raw encoded bytes. It panics on error.
func (k Keeper) MustUnmarshalQuery(bz []byte) types.MsgSubmitCrossChainQuery {
	var query types.MsgSubmitCrossChainQuery
	k.cdc.MustUnmarshal(bz, &query)
	return query
}

// MustMarshalQuery attempts to encode a CrossChainQuery object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalQueryResult(result *types.MsgSubmitCrossChainQueryResult) []byte {
	return k.cdc.MustMarshal(result)
}

// MustUnmarshalQuery attempts to decode and return a CrossChainQuery object from
// raw encoded bytes. It panics on error.
func (k Keeper) MustUnmarshalQueryResult(bz []byte) types.MsgSubmitCrossChainQueryResult {
	var result types.MsgSubmitCrossChainQueryResult
	k.cdc.MustUnmarshal(bz, &result)
	return result
}

// ClaimCapability allows the transfer module that can claim a capability that IBC module
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}