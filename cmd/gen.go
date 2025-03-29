package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addrNo string

var genCmd = &cobra.Command{
	Use: "gen",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("addrNo=%s", addrNo)
	},
}

func init() {
	genCmd.Flags().StringVar(&addrNo, "addr-no", "", "Address number")
	_ = genCmd.MarkFlagRequired("addr-no")
	rootCmd.AddCommand(genCmd)
}
