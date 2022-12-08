package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "selenian/testutil/keeper"
	"selenian/x/selenian/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SelenianKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
