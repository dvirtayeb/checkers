package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dvirtayeb/checkers/x/checkers/rules"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1w0rz4r5qsljn96c0ufvzqrr7lcd6qf49vfne36"
	bob   = "cosmos1p6atxgfyxkpw2ulprg93f9dllhx2r84lcrswu8"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func GetStoredGame1() *StoredGame {
	return &StoredGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
		Index:   "1",
		Game:    rules.New().String(),
		Turn:    "b",
	}
}

func TestCanGetAddressCreator(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
	creator, err2 := GetStoredGame1().GetCreatorAddress()
	require.Equal(t, aliceAddress, creator)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

func TestGetAddressWrongCreator(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Creator = "cosmos1w0rz4r5qsljn96c0ufvzqrr7lcd6qf49vfne37" // wrong address
	creator, err := storedGame.GetCreatorAddress()
	require.Nil(t, creator) // should be nil
	// should be the same error:
	require.EqualError(t,
		err,
		"creator address is invalid: cosmos1w0rz4r5qsljn96c0ufvzqrr7lcd6qf49vfne37: decoding bech32 failed: invalid checksum (expected vfne36 got vfne37)")
	require.EqualError(t, storedGame.Validate(), err.Error())
}
