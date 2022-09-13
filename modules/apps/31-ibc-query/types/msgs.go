package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
)

const (
	TypeMsgSubmitCrossChainQuery       = "submitCrossChainQuery"
	TypeMsgSubmitCrossChainQueryResult = "submitCrossChainQueryResult"
)

// NewMsgSubmitCrossChainQuery creates a new instance of NewMsgSubmitCrossChainQuery
func NewMsgSubmitCrossChainQuery(id string, path string, localTimeoutHeight *clienttypes.Height, localTimeoutStamp uint64, queryHeight uint64, creator string, srcPort string, srcChannel string) *MsgSubmitCrossChainQuery {
	return &MsgSubmitCrossChainQuery{
		Id:                 id,
		Path:               path,
		LocalTimeoutHeight: localTimeoutHeight,
		LocalTimeoutStamp:  localTimeoutStamp,
		QueryHeight:        queryHeight,
		Sender:             creator,
		SourcePort:         srcPort,
		SourceChannel:      srcChannel,
	}
}

func (msg MsgSubmitCrossChainQuery) GetQueryId() string { return msg.Id }

func (msg MsgSubmitCrossChainQuery) GetPath() string { return msg.Path }

func (msg MsgSubmitCrossChainQuery) GetTimeoutHeight() *clienttypes.Height {
	return msg.LocalTimeoutHeight
}

func (msg MsgSubmitCrossChainQuery) GetTimeoutTimestamp() uint64 { return msg.LocalTimeoutStamp }

func (msg MsgSubmitCrossChainQuery) GetQueryHeight() uint64 { return msg.QueryHeight }

// ValidateBasic implements sdk.Msg and performs basic stateless validation
func (msg MsgSubmitCrossChainQuery) ValidateBasic() error {

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgSubmitCrossChainQuery) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

// Route implements sdk.Msg
func (msg MsgSubmitCrossChainQuery) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgSubmitCrossChainQuery) Type() string {
	return TypeMsgSubmitCrossChainQuery
}

// GetSignBytes implements sdk.Msg.
func (msg MsgSubmitCrossChainQuery) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}
