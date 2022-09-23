package keeper_test

import (
	"fmt"
	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v4/testing"
	"time"
)

// test querying from chainA to chainB
func (suite *KeeperTestSuite) TestSendQuery() {
	var (
		query types.CrossChainQuery
		path  *ibctesting.Path
		err   error
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"successful query from querying chain",
			func() {
				query = types.CrossChainQuery{
					Id:                    "1",
					Path:                  "test/query-1",
					LocalTimeoutHeight:    clienttypes.NewHeight(0, uint64(suite.chainA.CurrentHeader.Height+1)),
					LocalTimeoutTimestamp: uint64(suite.chainA.CurrentHeader.Time.UnixNano() + 100),
					QueryHeight:           uint64(suite.chainB.CurrentHeader.Height),
					ClientId:              suite.chainA.SenderAccount.GetAddress().String(),
				}
			}, true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset
			path = NewQueryPath(suite.chainA, suite.chainB)
			suite.coordinator.Setup(path)
			//suite.coordinator.SetupConnections(path)
			//suite.coordinator.CreateQueryChannels(path)

			tc.malleate()

			err = suite.chainA.GetSimApp().IBCQueryKeeper.SendQuery(suite.chainA.GetContext(), path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, query)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnRecvPacket() {

	var (
		successResult   = types.QueryResult_QUERY_RESULT_SUCCESS
		failResult      = types.QueryResult_QUERY_RESULT_FAILURE
		timeoutResult   = types.QueryResult_QUERY_RESULT_TIMEOUT
		crossChainQuery types.CrossChainQuery
		queryResultData []byte
		chainAHeight    int64
		chainATimestamp uint64
	)

	testCases := []struct {
		msg         string
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
				LocalTimeoutHeight:    clienttypes.NewHeight(0, uint64(chainAHeight+1)),
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
				LocalTimeoutTimestamp: chainATimestamp + 30,
				QueryHeight:           uint64(suite.chainB.CurrentHeader.Height),
				ClientId:              suite.chainB.CurrentHeader.ChainID,
			}
			queryResultData = []byte("query result data4")
		}, true, timeoutResult},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			path := NewQueryPath(suite.chainA, suite.chainB)
			suite.coordinator.Setup(path)

			chainAHeight = suite.chainA.GetContext().BlockHeight()
			chainATimestamp = uint64(suite.chainA.GetContext().BlockTime().UnixNano())

			tc.malleate()
			query := crossChainQuery
			// send query from chainA to chainB
			queryMsg := types.NewMsgSubmitCrossChainQuery(query.Id, query.Path, query.LocalTimeoutHeight, query.LocalTimeoutTimestamp, query.QueryHeight, suite.chainA.SenderAccount.GetAddress().String(), path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
			_, err := suite.chainA.SendMsgs(queryMsg)
			suite.Require().NoError(err) // message committed
			suite.chainA.GetSimApp().IBCQueryKeeper.SetCrossChainQuery(suite.chainA.GetContext(), query)

			// query result of chain B
			data := types.IBCQueryResultPacketData{
				Id:          query.Id,
				Path:        query.Path,
				QueryHeight: query.QueryHeight,
				Result:      tc.queryResult,
				Data:        queryResultData,
			}

			seq := uint64(1)
			// receive query response packet from chain B
			packet := channeltypes.NewPacket(data.GetBytes(), seq, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.NewHeight(0, 100), 0)
			err = suite.chainA.GetSimApp().IBCQueryKeeper.OnRecvPacket(suite.chainA.GetContext(), packet)
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
		})
	}
}

func (suite *KeeperTestSuite) TestOnAcknowledgementPacket() {
	var (
		successAck = channeltypes.NewResultAcknowledgement([]byte{byte(1)})
		failedAck  = channeltypes.NewErrorAcknowledgement(fmt.Errorf("failed packet query"))
		path       *ibctesting.Path
	)

	testCases := []struct {
		msg      string
		ack      channeltypes.Acknowledgement
		malleate func()
		success  bool // success of ack
		expPass  bool
	}{
		{"success ack causes no-op", successAck, func() {}, true, true},
		{"successful emit error", failedAck, func() {}, false, false},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset
			path = NewQueryPath(suite.chainA, suite.chainB)
			suite.coordinator.Setup(path)

			tc.malleate()

			err := suite.chainA.GetSimApp().IBCQueryKeeper.OnAcknowledgementPacket(suite.chainA.GetContext(), tc.ack)
			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnTimeoutPacket() {
	var (
		path *ibctesting.Path
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"successful timeout from sender as source chain", func() {}, true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			path = NewQueryPath(suite.chainA, suite.chainB)
			suite.coordinator.Setup(path)

			tc.malleate()

			err := suite.chainA.GetSimApp().IBCQueryKeeper.OnTimeoutPacket(suite.chainA.GetContext())

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
