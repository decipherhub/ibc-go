package keeper_test

import (
	"fmt"
	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	ibctesting "github.com/cosmos/ibc-go/v4/testing"
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
					LocalTimeoutHeight:    clienttypes.NewHeight(0, uint64(suite.chainA.CurrentHeader.Height+100)),
					LocalTimeoutTimestamp: uint64(suite.chainA.CurrentHeader.Time.UnixNano() + 10000000),
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
			"timeout from sender as source chain", func() {}, true,
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

			suite.Require().Error(err)
		})
	}
}
