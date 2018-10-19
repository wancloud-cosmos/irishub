package record

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
)

// nolint

// Record Keeper
type Keeper struct {

	// The (unexposed) keys used to access the stores from the Context.
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding.
	cdc *wire.Codec

	// Reserved codespace
	codespace sdk.CodespaceType
}

// NewKeeper returns a mapper that uses go-wire to (binary) encode and decode record types.
func NewKeeper(cdc *wire.Codec, key sdk.StoreKey, codespace sdk.CodespaceType) Keeper {
	return Keeper{
		storeKey:  key,
		cdc:       cdc,
		codespace: codespace,
	}
}

// Returns the go-wire codec.
func (keeper Keeper) WireCodec() *wire.Codec {
	return keeper.cdc
}

func KeyRecord(dataHash string) []byte {
	return []byte(fmt.Sprintf("record:%s", dataHash))
}

func (keeper Keeper) AddRecord(ctx sdk.Context, msg MsgSubmitRecord) {

	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinary(msg)
	store.Set(KeyRecord(msg.DataHash), bz)
}
