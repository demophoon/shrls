package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Shrls",
	Long:  `All software has versions. This is Shrls's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Shrls v0.1.0")
	},
}
