package leon_test

import (
	"testing"

	keepertest "github.com/leonnipp/leon/testutil/keeper"
	"github.com/leonnipp/leon/testutil/nullify"
	"github.com/leonnipp/leon/x/leon"
	"github.com/leonnipp/leon/x/leon/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LeonKeeper(t)
	leon.InitGenesis(ctx, *k, genesisState)
	got := leon.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
