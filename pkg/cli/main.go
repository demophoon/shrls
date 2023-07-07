package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/cli/serve"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/cli/shrls"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/config"
)

var rootCmd = &cobra.Command{
	Use:   "shrls",
	Short: "Shrls is a easy to use url shortener",
	Long:  `An easy to use, feature rich url shortner built in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(helpCmd)
	rootCmd.AddCommand(shrls.ShrlsCmd)
	rootCmd.AddCommand(serve.ServeCmd)
}
