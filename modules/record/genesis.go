package record

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irishub/modules/gov"
	"github.com/irisnet/irishub/modules/iparam"
	"github.com/irisnet/irishub/modules/record/params"
)

func InitGenesis(ctx sdk.Context, k gov.Keeper) {
	iparam.InitGenesisParameter(&recordparams.UploadLimitChangingProcedureParameter, ctx, nil)
}
