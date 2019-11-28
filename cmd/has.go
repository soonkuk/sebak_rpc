package cmd

import (
	"fmt"
	jsonrpc "github.com/gorilla/rpc/json"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strconv"
)

var (
	// ReleaseSnapshotCmd is command to release db snapshot
	HasCmd *cobra.Command
)

func init() {
	HasCmd = &cobra.Command{
		Use:  "has <snapshot> <key>",
		Args: cobra.ExactArgs(2),
		Run: func(c *cobra.Command, args []string) {
			var (
				err         error
				snapshot    string
				key         string
				resp        *http.Response
				result      DBHasResult
				blockheight uint64
			)
			snapshot = args[0]
			switch flagKeyType {
			case "account":
				key = GetBlockAccountKey(args[1])
			case "blockheight":
				blockheight, err = strconv.ParseUint(args[1], 10, 64)
				if err != nil {
					fmt.Println(err.Error())
				}
				key = GetBlockKeyPrefixHeight(blockheight)
			case "blockhash":
				key = GetBlockKey(args[1])
			default:
				fmt.Println("flag key-type is missing or incorrect")
				os.Exit(1)
			}

			resp = request("DB.Has", &DBGetArgs{Snapshot: snapshot, Key: key})
			defer resp.Body.Close()

			err = jsonrpc.DecodeClientResponse(resp.Body, &result)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(bool(result))
		},
	}
	HasCmd.Flags().StringVar(&flagKeyType, "key-type", flagKeyType, "key type")
	rootCmd.AddCommand(HasCmd)
}
