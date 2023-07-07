package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print the help for Shrls",
	Long:  `Print help text for using Shrls`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Shrls v0.1.0")
	},
}
