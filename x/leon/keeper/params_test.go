package keeper_test

import (
	"testing"

	testkeeper "github.com/leonnipp/leon/testutil/keeper"
	"github.com/leonnipp/leon/x/leon/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LeonKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
