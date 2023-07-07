package cli

import (
	"fmt"

	"github.com/raphamorim/go-rainbow"
	"github.com/spf13/cobra"
)

var (
	version string
	build   string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Shrls",
	Long:  `All software has versions. This is Shrls's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(rainbow.Bold(rainbow.Hex("#7951A1", "SHRLS")))
		fmt.Print(rainbow.Hex("#C2C2C2", " v"))
		fmt.Println(rainbow.Hex("#FFFFFF", version))
	},
}
