package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// 31-ibc-query sentinel errors
var (
	ErrInvalidVersion          = sdkerrors.Register(ModuleName, 2, "invalid 31-IBC-query version")
	ErrInvalidTimeoutHeight    = sdkerrors.Register(ModuleName, 3, "invalid timeout height")
	ErrCrossChainQueryNotFound = sdkerrors.Register(ModuleName, 4, "no query found for given query id")
	ErrQuerytTimeout           = sdkerrors.Register(ModuleName, 5, "query timeout")
	ErrMaxTransferChannels     = sdkerrors.Register(ModuleName, 6, "max transfer channels")
	ErrUnknownDataType         = sdkerrors.Register(ModuleName, 7, "unknown data type")
)
