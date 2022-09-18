package keeper_test

import (
	"fmt"

	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibctesting "github.com/cosmos/ibc-go/v4/testing"
)

const (
	queryID = "query-1"
)

// test querying from chainA to chainB using both coin that orignate on
// chainA and coin that orignate on chainB
func (suite *KeeperTestSuite) TestSendQuery() {
	var (
		query types.CrossChainQuery
		path   *ibctesting.Path
		err    error
	)

	testCases := []struct {
		msg            string
		malleate       func()
		sendFromSource bool
		expPass        bool
	}{
		{
			"successful query from source chain",
			func() {
				suite.coordinator.CreateQueryChannels(path)
				query = types.CrossChainQuery{
					Id:                    queryID,
					Path:                  "test/query-1",
					LocalTimeoutHeight:    clienttypes.NewHeight(0, 110),
					LocalTimeoutTimestamp: 0,
					QueryHeight:           uint64(11),
					ClientId:              suite.chainA.SenderAccount.GetAddress().String(),
				}
			}, true, true,
		},
		{
			"successful query from counterparty chain",
			func() {
				suite.coordinator.CreateQueryChannels(path)
			}, false, true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset
			path = NewQueryPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)

			tc.malleate()

			if !tc.sendFromSource {
				// send query from chainB to chainA
				transferMsg := types.NewMsgSubmitCrossChainQuery(queryID, "path/query-1", clienttypes.NewHeight(0, 110), 0, uint64(11), suite.chainB.SenderAccount.GetAddress().String(), path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID)
				_, err = suite.chainB.SendMsgs(transferMsg)
				suite.Require().NoError(err) // message committed

				// receive query on chainA from chainB
				queryPacket := types.NewIBCQueryPacketData(queryID, "path/query-1", uint64(11))
				packet := channeltypes.NewPacket(queryPacket.GetBytes(), 1, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, clienttypes.NewHeight(0, 110), 0)

				// get proof of packet commitment from chainB
				err = path.EndpointA.UpdateClient()
				suite.Require().NoError(err)
				packetKey := host.PacketCommitmentKey(packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetSequence())
				proof, proofHeight := path.EndpointB.QueryProof(packetKey)

				recvMsg := channeltypes.NewMsgRecvPacket(packet, proof, proofHeight, suite.chainA.SenderAccount.GetAddress().String())
				_, err = suite.chainA.SendMsgs(recvMsg)
				suite.Require().NoError(err) // message committed
			}

			err = suite.chainA.GetSimApp().IBCQueryKeeper.SendQuery(
				suite.chainA.GetContext(), path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID,
				query,  clienttypes.NewHeight(0, 110), 0,
			)

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
		queryResultData []byte
	)

	testCases := []struct {
		msg          string
		queryResult  types.QueryResult
		malleate     func()
		success      bool // success of query
		expPass      bool
	}{
		{"success receive on source chain", successResult, func() {
			queryResultData = []byte("query result data")
		}, true, true},
		{"fail receive on source chain", failResult, func() {
			queryResultData = nil
		}, false, true},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			suite.SetupTest() // reset

			path := NewQueryPath(suite.chainA, suite.chainB)
			suite.coordinator.Setup(path)

			seq := uint64(1)

			// send query from chainA to chainB
			queryMsg := types.NewMsgSubmitCrossChainQuery(queryID, "test/query-1", clienttypes.NewHeight(0, 110), 0, uint64(11), suite.chainA.SenderAccount.GetAddress().String(), path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
			_, err := suite.chainA.SendMsgs(queryMsg)
			query := types.CrossChainQuery{
				Id:                    queryMsg.Id,
				Path:                  queryMsg.Path,
				LocalTimeoutHeight:    queryMsg.LocalTimeoutHeight,
				LocalTimeoutTimestamp: queryMsg.LocalTimeoutStamp,
				QueryHeight:           queryMsg.QueryHeight,
				ClientId:              queryMsg.Sender,
			}
			suite.chainA.GetSimApp().IBCQueryKeeper.SetCrossChainQuery(suite.chainA.GetContext(), query)
			suite.Require().NoError(err) // message committed

			tc.malleate()
			
			// query result of chain B
			data := (types.IBCQueryResultPacketData{
				Id:           queryID,
				Path:         "test/query-1",
				QueryHeight:  uint64(11),
				Result:       tc.queryResult,
				Data:         queryResultData,
			})
			
			// receive query response packet from chain B 
			packet := channeltypes.NewPacket(data.GetBytes(), seq, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.NewHeight(0, 100), 0)
			err = suite.chainA.GetSimApp().IBCQueryKeeper.OnRecvPacket(suite.chainA.GetContext(), packet)

			if tc.expPass {
				suite.Require().NoError(err)
				result, _ := suite.chainA.GetSimApp().IBCQueryKeeper.GetCrossChainQueryResult(suite.chainA.GetContext(), queryID)

				if tc.success {
					suite.Require().Equal("query result data", string(result.GetData()))
					suite.Require().Equal("QUERY_RESULT_SUCCESS", result.GetResult().String())
				} else {
					suite.Require().Equal("", string(result.GetData()))
					suite.Require().Equal("QUERY_RESULT_FAILURE", result.GetResult().String())
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
		{"success ack causes no-op", successAck, func() {	}, true, true},
		{"successful emit error", failedAck, func() {   }, false, false},
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
		path   *ibctesting.Path
	)

	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
	}{
		{
			"successful timeout from sender as source chain", func() { }, true,
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
