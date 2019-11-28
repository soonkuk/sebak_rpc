package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	jsonrpc "github.com/gorilla/rpc/json"
	"github.com/spf13/cobra"
)

var (
	// OpenSnapshotCmd is command to open db snapshot
	OpenSnapshotCmd *cobra.Command
)

func init() {
	OpenSnapshotCmd = &cobra.Command{
		Use: "open",
		Run: func(c *cobra.Command, args []string) {
			var (
				err      error
				resp     *http.Response
				result   DBOpenSnapshotResult
				snapshot string
			)
			resp = request("DB.OpenSnapshot", &DBOpenSnapshotResult{})
			defer resp.Body.Close()

			err = jsonrpc.DecodeClientResponse(resp.Body, &result)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			snapshot = result.Snapshot
			fmt.Println(snapshot)
		},
	}
	rootCmd.AddCommand(OpenSnapshotCmd)
}

func request(method string, args interface{}) (resp *http.Response) {
	var (
		err      error
		rawURL   = "http://0.0.0.0:54321"
		endpoint *url.URL
		message  []byte
		req      *http.Request
		client   *http.Client
	)

	endpoint, err = url.Parse(rawURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	message, err = jsonrpc.EncodeClientRequest(method, &args)
	req, err = http.NewRequest("POST", endpoint.String(), bytes.NewBuffer(message))
	req.Header.Set("Content-Type", "application/json")
	client = new(http.Client)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
