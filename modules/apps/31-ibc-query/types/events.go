package types

import (
	fmt "fmt"

	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

const (
	EventTypeTimeout = "timeout"
	EventSendQuery   = "sendQuery"
	EventTypePacket  = "ibc_query_packet"

	AttributeQueryData           = "query_data"
	AttributeKeyTimeoutTimestamp = "query_timeout_timestamp"
	AttributeKeyQueryID          = "query_id"
	AttributeKeyTimeoutHeight    = "query_timeout_height"
	AttributeKeyQueryHeight      = "query_height"
	AttributeKeyAckSuccess       = "success"
	AttributeKeyAckError         = "error"

	QuerySubmitted = "QuerySubmitted"
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
		LocalTimeoutHeight: localTimeoutHeight,
		LocalTimeoutStamp:  localTimeoutStamp,
		QueryHeight:        queryHeight,
	}
}
