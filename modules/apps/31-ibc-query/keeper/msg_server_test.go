package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	ibctesting "github.com/cosmos/ibc-go/v4/testing"
	"time"
)

const (
	portid = "testportid"
	chanid = "channel-0"
)

func (suite *KeeperTestSuite) TestSubmitCrossChainQuery() {
	var (
		path             *ibctesting.Path
		msg              *types.MsgSubmitCrossChainQuery
		queryHeight      int64
		timeoutHeight    clienttypes.Height
		timeoutTimestamp uint64
		addr             string
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
				queryHeight = suite.chainB.GetContext().BlockHeight()
				timeoutHeight = clienttypes.NewHeight(0, uint64(suite.chainA.GetContext().BlockHeight()+50))
				timeoutTimestamp = uint64(suite.chainA.GetContext().BlockTime().UnixNano() + 5000000)
				msg = types.NewMsgSubmitCrossChainQuery("query-1", "test/query_path", timeoutHeight, timeoutTimestamp, uint64(queryHeight), addr, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
			},
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		path = NewQueryPath(suite.chainA, suite.chainB)
		suite.coordinator.SetupConnections(path)

		tc.malleate()
		res, err := suite.chainA.GetSimApp().IBCQueryKeeper.SubmitCrossChainQuery(sdk.WrapSDKContext(suite.chainA.GetContext()), msg)
		if tc.expPass {
			suite.Require().NoError(err)
			suite.Require().NotNil(res)
			queryResult, found := suite.chainA.GetSimApp().IBCQueryKeeper.GetCrossChainQuery(suite.chainA.GetContext(), "query-1")
			suite.Require().True(found)
			suite.Require().Equal(msg.Id, queryResult.Id)
			suite.Require().Equal(msg.Path, queryResult.Path)
			suite.Require().Equal(msg.LocalTimeoutHeight.RevisionHeight, queryResult.LocalTimeoutHeight.RevisionHeight)
			suite.Require().Equal(msg.LocalTimeoutStamp, queryResult.LocalTimeoutTimestamp)
			suite.Require().Equal(msg.QueryHeight, queryResult.QueryHeight)
			suite.Require().Equal(addr, queryResult.ClientId)
		}
	}
}

func (suite *KeeperTestSuite) TestSubmitCrossChainQueryResult() {
	var (
		path            *ibctesting.Path
		successResult   = types.QueryResult_QUERY_RESULT_SUCCESS
		failResult      = types.QueryResult_QUERY_RESULT_FAILURE
		timeoutResult   = types.QueryResult_QUERY_RESULT_TIMEOUT
		crossChainQuery types.CrossChainQuery
		queryResultData []byte
		chainAHeight    int64
		chainATimestamp uint64
	)

	testCases := []struct {
		name        string
		queryResult types.QueryResult
		malleate    func()
		expPass     bool
		expResult   types.QueryResult
	}{
		{"success receive from queried chain", successResult, func() {
			crossChainQuery = types.CrossChainQuery{
				Id:                    "1",
				Path:                  "test/query-1",
				LocalTimeoutHeight:    clienttypes.NewHeight(0, uint64(chainAHeight+50)),
				LocalTimeoutTimestamp: chainATimestamp + uint64(time.Minute.Nanoseconds()),
				QueryHeight:           uint64(suite.chainB.CurrentHeader.Height),
				ClientId:              suite.chainB.CurrentHeader.ChainID,
			}
			queryResultData = []byte("query result data")
		}, true, successResult},
		{"fail receive from queried chain", failResult, func() {
			crossChainQuery = types.CrossChainQuery{
				Id:                    "2",
				Path:                  "test/query-2",
				LocalTimeoutHeight:    clienttypes.NewHeight(0, uint64(chainAHeight+50)),
				LocalTimeoutTimestamp: chainATimestamp + uint64(time.Minute.Nanoseconds()),
				QueryHeight:           uint64(suite.chainB.CurrentHeader.Height),
				ClientId:              suite.chainB.CurrentHeader.ChainID,
			}
			queryResultData = []byte("query result data2")
		}, true, failResult},
		{"timeout because of block height", successResult, func() {
			crossChainQuery = types.CrossChainQuery{
				Id:                    "3",
				Path:                  "test/query-3",
				LocalTimeoutHeight:    clienttypes.NewHeight(0, uint64(chainAHeight-1)),
				LocalTimeoutTimestamp: chainATimestamp + uint64(time.Minute.Nanoseconds()),
				QueryHeight:           uint64(suite.chainB.CurrentHeader.Height),
				ClientId:              suite.chainB.CurrentHeader.ChainID,
			}
			queryResultData = []byte("query result data3")
		}, true, timeoutResult},
		{"timeout because of block timestamp", failResult, func() {
			crossChainQuery = types.CrossChainQuery{
				Id:                    "4",
				Path:                  "test/query-4",
				LocalTimeoutHeight:    clienttypes.NewHeight(0, uint64(chainAHeight+50)),
				LocalTimeoutTimestamp: chainATimestamp - 30,
				QueryHeight:           uint64(suite.chainB.CurrentHeader.Height),
				ClientId:              suite.chainB.CurrentHeader.ChainID,
			}
			queryResultData = []byte("query result data4")
		}, true, timeoutResult},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		path = NewQueryPath(suite.chainA, suite.chainB)
		suite.coordinator.SetupConnections(path)

		chainAHeight = suite.chainA.GetContext().BlockHeight()
		chainATimestamp = uint64(suite.chainA.GetContext().BlockTime().UnixNano())

		tc.malleate()

		query := crossChainQuery
		suite.chainA.GetSimApp().IBCQueryKeeper.SetCrossChainQuery(suite.chainA.GetContext(), query)

		queryResultMsg := types.NewMsgSubmitCrossChainQueryResult(query.Id, query.Path, query.QueryHeight, tc.queryResult, queryResultData, nil)
		_, err := suite.chainA.GetSimApp().IBCQueryKeeper.SubmitCrossChainQueryResult(sdk.WrapSDKContext(suite.chainA.GetContext()), queryResultMsg)
		suite.Require().NoError(err)
		if tc.expPass {
			result, isUnmarshal := suite.chainA.GetSimApp().IBCQueryKeeper.GetCrossChainQueryResult(suite.chainA.GetContext(), query.Id)
			suite.Require().True(isUnmarshal)
			suite.Require().Equal(queryResultData, result.GetData())
			switch tc.expResult {
			case successResult:
				suite.Require().Equal(successResult, result.GetResult())
			case failResult:
				suite.Require().Equal(failResult, result.GetResult())
			case timeoutResult:
				suite.Require().Equal(timeoutResult, result.GetResult())
			}
		} else {
			suite.Require().Error(err)
		}
	}
}
