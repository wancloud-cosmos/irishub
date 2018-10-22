package cli

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/irisnet/irishub/client/context"
	"github.com/irisnet/irishub/modules/record"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	recordClient "github.com/irisnet/irishub/client/record"
	"github.com/irisnet/irishub/modules/record/params"
)

type RecordMetadata struct {
	OwnerAddress sdk.AccAddress
	SubmitTime   int64
	DataHash     string
	DataSize     int64
	//PinedNode    string
}

func GetCmdQureyRecord(storeName string, cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query [record ID]",
		Short: "query specified file with record ID",
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			recordID := viper.GetString(flagRecordID)

			para, _ := cliCtx.QueryStore([]byte(recordparams.UploadLimitChangingProcedureParameter.GetStoreKey()), "params")
			var limits recordparams.UploadLimitChangingProcedure
			cdc.MustUnmarshalBinary(para, &limits)
			fmt.Printf("####UploadLimitOnchain : %d", limits.UploadLimitOnchain)
			fmt.Printf("####UploadLimitIpfs : %d", limits.UploadLimitIpfs)
			os.Exit(1)

			res, err := cliCtx.QueryStore([]byte(recordID), storeName)
			if len(res) == 0 || err != nil {
				return fmt.Errorf("Record ID [%s] is not existed", recordID)
			}

			var submitRecord record.MsgSubmitRecord
			cdc.MustUnmarshalBinary(res, &submitRecord)

			recordResponse, err := recordClient.ConvertRecordToRecordOutput(cliCtx, submitRecord)
			if err != nil {
				return err
			}

			output, err := wire.MarshalJSONIndent(cdc, recordResponse)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil

		},
	}

	cmd.Flags().String(flagRecordID, "", "record ID for query")

	return cmd
}
