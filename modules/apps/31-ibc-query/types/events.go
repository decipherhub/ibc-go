package types

import (
	fmt "fmt"

	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
)

const (
	EventTypeTimeout = "timeout"
	EventTypePacket  = "ibc_query_packet"
	AttributeKeyAckSuccess       = "success"
	AttributeKeyAckError         = "error"
)

var (
	AttributeValueCategory = fmt.Sprintf("%s_%s", host.ModuleName, ModuleName)
)

func NewEventQuerySubmitted(
	id, path string,
	localTimeoutHeight clienttypes.Height,
	localTimeoutStamp, queryHeight uint64) *EventQuerySubmitted {
	return &EventQuerySubmitted{
		Id:                 id,
		Path:               path,
		LocalTimeoutHeight: &localTimeoutHeight,
		LocalTimeoutStamp:  localTimeoutStamp,
		QueryHeight:        queryHeight,
	}
}
