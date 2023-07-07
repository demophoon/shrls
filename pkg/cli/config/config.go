package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/config"
)

// TODO: Write config Viewing and Modifying CLI

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure shrls service",
	Long:  `View and edit the configuration for the shrls service.`,
	Run:   shrls_config,
}

func shrls_config(cmd *cobra.Command, args []string) {
	config := config.New()
	fmt.Printf("Config: %#v", config)
}
