package recordparams

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/irisnet/irishub/modules/iparam"
)

const (
	UploadLimitOfOnchain = 1024        //upload limit on chain in bytes(1K currently)
	UploadLimitOfIpfs    = 1024 * 1024 //upload limit on chain in bytes(1M currently)
)

var UploadLimitChangingProcedureParameter UploadLimitChangingProcedureParam
var _ iparam.GovParameter = (*UploadLimitChangingProcedureParam)(nil)

// Procedure around record upload limit changing in governance
type UploadLimitChangingProcedure struct {
	UploadLimitOnchain int64 `json:"onchain_upload_limit"` //  Onchain Upload limit
	UploadLimitIpfs    int64 `json:"ipfs_upload_limit"`    //  Ipfs Upload limit
}

type UploadLimitChangingProcedureParam struct {
	Value   UploadLimitChangingProcedure
	psetter params.Setter
	pgetter params.Getter
}

func (param *UploadLimitChangingProcedureParam) GetValueFromRawData(cdc *wire.Codec, res []byte) interface{} {
	cdc.MustUnmarshalBinary(res, &param.Value)
	return param.Value
}

func (param *UploadLimitChangingProcedureParam) InitGenesis(genesisState interface{}) {
	param.Value = UploadLimitChangingProcedure{
		UploadLimitOnchain: UploadLimitOfOnchain,
		UploadLimitIpfs:    UploadLimitOfIpfs,
	}
}

func (param *UploadLimitChangingProcedureParam) SetReadWriter(setter params.Setter) {
	param.psetter = setter
	param.pgetter = setter.Getter
}

func (param *UploadLimitChangingProcedureParam) GetStoreKey() string {
	return "Gov/record/UploadLimitChangingProcedure"
}

func (param *UploadLimitChangingProcedureParam) SaveValue(ctx sdk.Context) {
	param.psetter.Set(ctx, param.GetStoreKey(), param.Value)
}

func (param *UploadLimitChangingProcedureParam) LoadValue(ctx sdk.Context) bool {
	err := param.pgetter.Get(ctx, param.GetStoreKey(), &param.Value)
	if err != nil {
		return false
	}
	return true
}

func (param *UploadLimitChangingProcedureParam) ToJson(jsonStr string) string {
	var jsonBytes []byte

	if len(jsonStr) == 0 {
		jsonBytes, _ = json.Marshal(param.Value)
		return string(jsonBytes)
	}

	if err := json.Unmarshal([]byte(jsonStr), &param.Value); err == nil {
		jsonBytes, _ = json.Marshal(param.Value)
		return string(jsonBytes)
	}
	return string(jsonBytes)
}

func (param *UploadLimitChangingProcedureParam) Update(ctx sdk.Context, jsonStr string) {
	if err := json.Unmarshal([]byte(jsonStr), &param.Value); err == nil {
		param.SaveValue(ctx)
	}
}

func (param *UploadLimitChangingProcedureParam) Valid(jsonStr string) sdk.Error {
	return nil
}
