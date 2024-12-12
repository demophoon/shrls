package cli

import (
	"fmt"

	"github.com/demophoon/shrls/pkg/version"
	"github.com/raphamorim/go-rainbow"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Shrls",
	Long:  `All software has versions. This is Shrls's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(rainbow.Bold(rainbow.Hex("#7951A1", "SHRLS")))
		fmt.Print(rainbow.Hex("#FFFFFF", " "))
		fmt.Println(rainbow.Hex("#FFFFFF", version.Version))
		fmt.Println(rainbow.Dim(fmt.Sprintf("Build: %s", version.Build)))
	},
}
