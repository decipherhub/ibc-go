package keeper

import (
	"fmt"

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
	query types.CrossChainQuery,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64) error {
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


	packet := channeltypes.NewPacket(
		packetData.GetBytes(),
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.ics4Wrapper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvPacket processes a cross chain query result.
func (k Keeper) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet) error {
	var packetData types.IBCQueryResultPacketData
	var queryResult types.CrossChainQueryResult

	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &packetData); err != nil {
		return sdkerrors.Wrapf(types.ErrUnknownDataType, "cannot unmarshal ICS-31 interchain query packet packetData")
	}

	if err := packetData.ValidateBasic(); err != nil {
		return err
	}

	// TODO: validate query packetData with proof

	queryResult = types.CrossChainQueryResult{
		Id:     packetData.Id,
		Result: packetData.Result,
		Data:   packetData.Data,
	}

	// remove CrossChainQuery from privateStore
	if _, found := k.GetCrossChainQuery(ctx, queryResult.Id); found {
		k.DeleteCrossChainQuery(ctx, queryResult.Id)
	}

	// store result in privateStore
	k.SetCrossChainQueryResult(ctx, queryResult)

	return nil
}

func (k Keeper) OnAcknowledgementPacket(ctx sdk.Context, ack channeltypes.Acknowledgement) error {
	switch ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		return fmt.Errorf("query packet acknowledgment error")
	default:
		// the acknowledgement succeeded on the receiving chain so nothing
		// needs to be executed and no error needs to be returned
		return nil
	}
}

func (k Keeper) OnTimeoutPacket(ctx sdk.Context) error {
	sdkerrors.Wrapf(types.ErrQuerytTimeout, "query packet timeout")
	return nil
}
