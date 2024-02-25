package config

import (
	"fmt"

	"github.com/demophoon/shrls/pkg/config"
	"github.com/spf13/cobra"
)

// TODO: Write config Viewing and Modifying CLI

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure shrls service",
	Long:  `View and edit the configuration for the shrls service.`,
	Run:   shrls_config,
}

func shrls_config(cmd *cobra.Command, args []string) {
	config := config.New()
	fmt.Printf("Config: %#v", config)
}
