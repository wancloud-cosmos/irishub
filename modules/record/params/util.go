package recordparams

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Returns the current record upload limit change Procedure from the global param store
func GetUploadLimitChangingProcedure(ctx sdk.Context) UploadLimitChangingProcedure {
	UploadLimitChangingProcedureParameter.LoadValue(ctx)
	return UploadLimitChangingProcedureParameter.Value
}
