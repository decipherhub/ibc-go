package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/31-ibc-query/types"
)

func (suite *KeeperTestSuite) TestQueryResult() {
    var (
        req      *types.QueryCrossChainQueryResult
    )

    testCases := []struct {
        msg      string
        malleate func()
        expPass  bool
    }{
        {
            "success: correct query id",
            func() {
                queryResult := types.CrossChainQueryResult{
                    Id: "query-1",
                    Result: types.QueryResult_QUERY_RESULT_SUCCESS,
                    Data: []byte("test data"),
                }
                suite.chainA.GetSimApp().IBCQueryKeeper.SetCrossChainQueryResult(suite.chainA.GetContext(), queryResult)

                req = &types.QueryCrossChainQueryResult{
                    Id: "query-1",
                }
            },
            true,
        },
        {
            "failure: invalid query id",
            func() {
                queryResult := types.CrossChainQueryResult{
                    Id: "query-1",
                    Result: types.QueryResult_QUERY_RESULT_SUCCESS,
                    Data: []byte("test data"),
                }
                suite.chainA.GetSimApp().IBCQueryKeeper.SetCrossChainQueryResult(suite.chainA.GetContext(), queryResult)
                
                req = &types.QueryCrossChainQueryResult{
                    Id: "query-2",
                }
            },
            false,
        },
    }

    for _, tc := range testCases {
        suite.Run(fmt.Sprintf("Case %s", tc.msg), func() {
            suite.SetupTest() // reset

            tc.malleate()
            ctx := sdk.WrapSDKContext(suite.chainA.GetContext())
            res, err := suite.queryClient.CrossChainQueryResult(ctx, req)

            if tc.expPass {
                suite.Require().NoError(err)
                suite.Require().Equal([]byte("test data"), res.Data)
                suite.Require().Equal(types.QueryResult_QUERY_RESULT_SUCCESS, res.Result)
            } else {
                suite.Require().Error(err)
            }
        })
    }
}

