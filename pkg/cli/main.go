package cli

import (
	"fmt"
	"os"

	configcmd "github.com/demophoon/shrls/pkg/cli/config"
	"github.com/demophoon/shrls/pkg/cli/serve"
	"github.com/demophoon/shrls/pkg/cli/shrls"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "shrls",
	Short: "Shrls is a easy to use url shortener",
	Long:  `An easy to use, feature rich url shortner built in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionCmd.Run(cmd, nil)
		cmd.Help()
	},
}

func Execute() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("config", "", "Path to config.yaml file")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	rootCmd.PersistentFlags().Bool("trace", false, "Output trace logging")
	viper.BindPFlag("log_trace", rootCmd.PersistentFlags().Lookup("trace"))

	rootCmd.PersistentFlags().Bool("debug", false, "Output debug logging")
	viper.BindPFlag("log_debug", rootCmd.PersistentFlags().Lookup("debug"))

	viper.BindEnv("experimental", "SHRLS_EXPERIMENTAL")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serve.ServeCmd)
	rootCmd.AddCommand(configcmd.ConfigCmd)

	if viper.GetBool("experimental") {
		rootCmd.AddCommand(shrls.ShrlsCmd)
	}
}
