package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
)

func (k Keeper) SendQuery(ctx sdk.Context,
	sourcePort,
	sourceChannel string,
	query types.CrossChainQuery) error {
	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel)
	}

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetData := types.NewIBCQueryPacketData(
		query.Id, query.Path, query.QueryHeight,
	)

	// timeoutHeight == 0 pass checking timeoutHeight on receiving chain
	packet := channeltypes.NewPacket(
		packetData.GetBytes(),
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		clienttypes.ZeroHeight(),
		query.LocalTimeoutTimestamp,
	)

	if err := k.ics4Wrapper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvPacket processes a cross chain query result.
func (k Keeper) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet) error {
	var packetData types.IBCQueryResultPacketData

	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &packetData); err != nil {
		return sdkerrors.Wrapf(types.ErrUnknownDataType, "cannot unmarshal ICS-31 cross chain query packetData")
	}

	if err := packetData.ValidateBasic(); err != nil {
		return err
	}

	queryResult := types.CrossChainQueryResult{
		Id:     packetData.Id,
		Result: packetData.Result,
		Data:   packetData.Data,
	}

	query, found := k.GetCrossChainQuery(ctx, queryResult.Id)
	// if CrossChainQuery of queryId doesn't exist in store, other relayer already submitted CrossChainQueryResult
	if !found {
		return sdkerrors.Wrapf(types.ErrCrossChainQueryNotFound, "cannot find ICS-31 cross chain query id %s", queryResult.Id)
	}

	k.DeleteCrossChainQuery(ctx, queryResult.Id)

	queryResult = types.CrossChainQueryResult{
		Id:     packetData.Id,
		Result: packetData.Result,
		Data:   packetData.Data,
	}

	// check Timeout by comparing the latest height of chain, latest timestamp
	selfHeight := clienttypes.GetSelfHeight(ctx)
	selfBlockTime := uint64(ctx.BlockTime().UnixNano())
	if selfHeight.GTE(query.LocalTimeoutHeight) {
		queryResult.Result = types.QueryResult_QUERY_RESULT_TIMEOUT
	}
	if selfBlockTime >= query.LocalTimeoutTimestamp {
		queryResult.Result = types.QueryResult_QUERY_RESULT_TIMEOUT
	}

	// store result in privateStore
	k.SetCrossChainQueryResult(ctx, queryResult)

	return nil
}


func (k Keeper) OnTimeoutPacket(ctx sdk.Context) error {
	return sdkerrors.Wrapf(types.ErrTimeout, "query packet timeout")
}
