package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/keeper"
)


func GetQueryIdentifier() (string,error) {
	return keeper.Keeper.GenerateQueryIdentifier(keeper.Keeper{}, sdk.Context{})
}

