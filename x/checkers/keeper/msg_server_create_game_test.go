package keeper_test

import (
	// "context"
	"testing"

	// sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dvirtayeb/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
	// "github.com/dvirtayeb/checkers/x/checkers"
	// "github.com/dvirtayeb/checkers/x/checkers/keeper"
)

const (
	alice = "cosmos1w0rz4r5qsljn96c0ufvzqrr7lcd6qf49vfne36"
	bob   = "cosmos1p6atxgfyxkpw2ulprg93f9dllhx2r84lcrswu8"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

// func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
//     k, ctx := setupKeeper(t)
//     checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
//     return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
// }

func TestCreateGame(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Red:     bob,
		Black:   carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		IdValue: "", // TODO: update with a proper value when updated
	}, *createResponse)
}
