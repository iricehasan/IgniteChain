package selenian_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "selenian/testutil/keeper"
	"selenian/testutil/nullify"
	"selenian/x/selenian"
	"selenian/x/selenian/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SelenianKeeper(t)
	selenian.InitGenesis(ctx, *k, genesisState)
	got := selenian.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
