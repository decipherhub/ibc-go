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
				msg = types.NewMsgSubmitCrossChainQuery("query-1", "test/query_path", &timeoutHeight, timeoutTimestamp, queryHeight.RevisionHeight, addr, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
			},
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		path = NewQueryrPath(suite.chainA, suite.chainB)
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

func (suite *KeeperTestSuite) TestSubmitCrossChainQueryResult() {
	var (
		query  types.CrossChainQuery
		msg    *types.MsgSubmitCrossChainQueryResult
		result types.QueryResult
		data   []byte
	)

	// checkList
	// 1. retrieve the query from privateStore
	// 2. remove query from privateStore
	// 3. store result in privateStore

	testCases := []struct {
		name     string
		expPass  bool
		malleate func()
	}{
		{
			"success",
			true,
			func() {
				query = types.CrossChainQuery{Id: "queryId"}
				result = types.QueryResult_QUERY_RESULT_SUCCESS
				data = []byte("query data")
				msg = types.NewMsgSubmitCrossChainQueryResult("queryId", result, data)
			},
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()

		tc.malleate()

		suite.chainA.GetSimApp().IBCQueryKeeper.SetCrossChainQuery(suite.chainA.GetContext(), query)

		res, err := suite.chainA.GetSimApp().IBCQueryKeeper.SubmitCrossChainQueryResult(sdk.WrapSDKContext(suite.chainA.GetContext()), msg)

		if tc.expPass {
			suite.Require().NoError(err)
			suite.Require().NotNil(res)
			queryResult, found := suite.chainA.GetSimApp().IBCQueryKeeper.GetCrossChainQueryResult(suite.chainA.GetContext(), "queryId")

			suite.Require().True(found)
			suite.Require().Equal(query.Id, queryResult.Id)
			suite.Require().Equal(result, queryResult.Result)
			suite.Require().Equal(data, queryResult.Data)
		} else {
			suite.Require().Error(err)
		}
	}
}
