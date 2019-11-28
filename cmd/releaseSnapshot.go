package cmd

import (
	"fmt"
	"net/http"

	jsonrpc "github.com/gorilla/rpc/json"
	"github.com/spf13/cobra"
)

var (
	// ReleaseSnapshotCmd is command to release db snapshot
	ReleaseSnapshotCmd *cobra.Command
)

func init() {
	ReleaseSnapshotCmd = &cobra.Command{
		Use:  "release <snapshot>",
		Args: cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			var (
				err      error
				snapshot string
				resp     *http.Response
				result   DBReleaseSnapshotResult
			)
			snapshot = args[0]
			resp = request("DB.ReleaseSnapshot", &DBReleaseSnapshot{Snapshot: snapshot})
			defer resp.Body.Close()

			err = jsonrpc.DecodeClientResponse(resp.Body, &result)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(bool(result))
		},
	}
	rootCmd.AddCommand(ReleaseSnapshotCmd)
}

type DBReleaseSnapshot struct {
	Snapshot string `json:"snapshot"`
}

type DBReleaseSnapshotResult bool
