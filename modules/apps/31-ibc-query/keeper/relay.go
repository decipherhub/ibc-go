package keeper

import (
	"github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	coretypes "github.com/cosmos/ibc-go/v4/modules/core/types"
)

func (k Keeper) SendQuery(ctx sdk.Context,
	sourcePort,
	sourceChannel string,
	data []byte,
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

	packet := channeltypes.NewPacket(
		data,
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


func (k Keeper) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet, data types.MsgSubmitCrossChainQuery) error {

	labels := []metrics.Label{
		telemetry.NewLabel(coretypes.LabelSourcePort, packet.GetSourcePort()),
		telemetry.NewLabel(coretypes.LabelSourceChannel, packet.GetSourceChannel()),
	}

	defer func() {

		telemetry.IncrCounterWithLabels(
			[]string{"ibc", types.ModuleName, "receive"},
			1,
			append(
				labels, telemetry.NewLabel(coretypes.LabelSource, "false"),
			),
		)
	}()

	return nil
}
 
func (k Keeper) OnAcknowledgementPacket(ctx sdk.Context, packet channeltypes.Packet, data types.MsgSubmitCrossChainQuery, ack channeltypes.Acknowledgement) error {
	switch ack.Response.(type) {
	default:
		// the acknowledgement succeeded on the receiving chain so nothing
		// needs to be executed and no error needs to be returned
		return nil
	}
}

func (k Keeper) OnTimeoutPacket(ctx sdk.Context, packet channeltypes.Packet, data types.MsgSubmitCrossChainQuery) error {
	return nil
}
