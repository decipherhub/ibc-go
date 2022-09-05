package types

import clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"

const (
	QuerySubmitted = "QuerySubmitted"
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
