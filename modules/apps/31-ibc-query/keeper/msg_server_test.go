package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	ibctesting "github.com/cosmos/ibc-go/v4/testing"
)

const (
	portid = "testportid"
	chanid = "channel-0"
)

var (
	timeoutHeight    = clienttypes.NewHeight(0, 100)
	timeoutTimestamp = uint64(0)
	addr             = sdk.AccAddress("testaddr111111111111").String()
	queryHeight      = clienttypes.NewHeight(0, 1)
)

func (suite *KeeperTestSuite) TestSubmitCrossChainQuery() {
	var (
		path *ibctesting.Path
		msg  *types.MsgSubmitCrossChainQuery
	)

	testCases := []struct {
		name     string
		expPass  bool
		malleate func()
	}{
		{
			"success",
			true,
			func() {
				suite.coordinator.CreateChannels(path)
				msg = types.NewMsgSubmitCrossChainQuery("query-1", "test/query_path", timeoutHeight, timeoutTimestamp, queryHeight.RevisionHeight, addr, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
			},
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		path = NewQueryPath(suite.chainA, suite.chainB)
		suite.coordinator.SetupConnections(path)

		tc.malleate()

		if tc.expPass {
			res, err := suite.chainA.GetSimApp().IBCQueryKeeper.SubmitCrossChainQuery(sdk.WrapSDKContext(suite.chainA.GetContext()), msg)

			suite.Require().NoError(err)
			suite.Require().NotNil(res)
			queryResult, found := suite.chainA.GetSimApp().IBCQueryKeeper.GetCrossChainQuery(suite.chainA.GetContext(), "query-1")

			suite.Require().True(found)
			suite.Require().Equal("query-1", queryResult.Id)
			suite.Require().Equal("test/query_path", queryResult.Path)
			suite.Require().Equal(timeoutHeight.RevisionHeight, queryResult.LocalTimeoutHeight.RevisionHeight)
			suite.Require().Equal(timeoutTimestamp, queryResult.LocalTimeoutTimestamp)
			suite.Require().Equal(queryHeight.RevisionHeight, queryResult.QueryHeight)
			suite.Require().Equal(addr, queryResult.ClientId)

		}
	}
}
