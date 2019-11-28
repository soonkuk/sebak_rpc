package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   os.Args[0],
		Short: "jsonrpc",
		Run: func(c *cobra.Command, args []string) {
			if len(args) < 1 {
				c.Usage()
			}
		},
	}
	flagKeyType string
)

const (
	BlockPrefixHash                       = string(0x00)
	BlockPrefixConfirmed                  = string(0x01)
	BlockPrefixHeight                     = string(0x02)
	BlockTransactionPrefixHash            = string(0x10)
	BlockTransactionPrefixSource          = string(0x11)
	BlockTransactionPrefixConfirmed       = string(0x12)
	BlockTransactionPrefixAccount         = string(0x13)
	BlockTransactionPrefixBlock           = string(0x14)
	BlockOperationPrefixHash              = string(0x20)
	BlockOperationPrefixTxHash            = string(0x21)
	BlockOperationPrefixSource            = string(0x22)
	BlockOperationPrefixTarget            = string(0x23)
	BlockOperationPrefixPeers             = string(0x24)
	BlockOperationPrefixTypeSource        = string(0x25)
	BlockOperationPrefixTypeTarget        = string(0x26)
	BlockOperationPrefixTypePeers         = string(0x27)
	BlockOperationPrefixCreateFrozen      = string(0x28)
	BlockOperationPrefixFrozenLinked      = string(0x29)
	BlockOperationPrefixBlockHeight       = string(0x2A)
	BlockAccountPrefixAddress             = string(0x30)
	BlockAccountPrefixCreated             = string(0x31)
	BlockAccountSequenceIDPrefix          = string(0x32)
	BlockAccountSequenceIDByAddressPrefix = string(0x33)
	TransactionPoolPrefix                 = string(0x40)
	InternalPrefix                        = string(0x50) // internal data
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(rootCmd, "", err)
	}
}

func SetArgs(s []string) {
	rootCmd.SetArgs(s)
}

type DBHasResult bool

type DBGetResult IterItem

type DBGetArgs struct {
	Snapshot string `json:"snapshot"`
	Key      string `json:"key"`
}

type IterItem struct {
	N     uint64
	Key   []byte
	Value []byte
}

type DBOpenSnapshotResult struct {
	Snapshot string `json:"snapshot"`
}

func GetBlockAccountKey(address string) string {
	return fmt.Sprintf("%s%s", BlockAccountPrefixAddress, address)
}

func GetBlockKey(hash string) string {
	return fmt.Sprintf("%s%s", BlockPrefixHash, hash)
}

func GetBlockKeyPrefixHeight(height uint64) string {
	return fmt.Sprintf("%s%020d", BlockPrefixHeight, height)
}
