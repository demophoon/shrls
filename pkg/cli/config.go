package cli

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetEnvPrefix("shrls")
	setDefaults()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.TraceLevel)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/shrls/")
	viper.AddConfigPath("$HOME/.config/shrls")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Warn("No config file found, using defaults.")
	}
}

func setDefaults() {
	viper.BindEnv("base_url")

	viper.BindEnv("port")
	viper.SetDefault("port", 3000)

	viper.BindEnv("grpc_port")
	viper.SetDefault("grpc_port", 3001)

	viper.BindEnv("default_redirect")

	viper.BindEnv("upload_directory")
	viper.SetDefault("upload_directory", "./uploads")

	viper.BindEnv("mongo_uri")
	viper.SetDefault("mongo_uri", "mongodb://mongo:password@localhost:27017")

	viper.BindEnv("admin_username")
	viper.SetDefault("admin_username", "")

	viper.BindEnv("admin_password")
	viper.SetDefault("admin_password", "")

	viper.SetDefault("terminal_redirect", false)
}
