package config

import (
	"fmt"

	"github.com/demophoon/shrls/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// TODO: Write config Viewing and Modifying CLI

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure shrls service",
	Long:  `View and edit the configuration for the shrls service.`,
	Run:   shrls_config,
}

var ConfigGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a sample config",
	Long:  `Generate a sample configuration for shrls to run`,
	Run:   shrls_config_generate,
}

func init() {
	ConfigCmd.AddCommand(ConfigGenerateCmd)
}

func shrls_config(cmd *cobra.Command, args []string) {
	config := config.New()

	yaml_config, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("Unable to print config. %s", err)
	}
	fmt.Print(string(yaml_config))
}

func shrls_config_generate(cmd *cobra.Command, args []string) {
	config.New()

	err := viper.WriteConfigAs("./config.yaml")
	if err != nil {
		log.Fatalf("Couldn't write generated config. %s", err)
	}
}
