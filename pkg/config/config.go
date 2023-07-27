package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type StateBackend int64
type UploadBackend int64
type AuthBackend int64

const (
	MongoDB StateBackend = iota
	// TODO: BoltDB
	// TODO: InMemory
)

const (
	Directory UploadBackend = iota
	// TODO: S3
)

const (
	None AuthBackend = iota
	Basic
	// TODO: OAuth
)

type Config struct {
	BaseURL          string `mapstructure:"base_url"`
	Port             int
	DefaultRedirect  string `mapstructure:"default_redirect"`
	TerminalRedirect bool

	// TODO: Initialize backends outside of global config object
	// State
	StateBackend          StateBackend
	MongoConnectionString string `mapstructure:"mongo_uri"`

	BoltPath string `mapstructure:"bolt_path"`

	// Uploads
	UploadBackend   UploadBackend
	UploadDirectory string `mapstructure:"upload_directory"`

	// Auth
	AuthBackend   AuthBackend
	AdminUsername string `mapstructure:"admin_username"`
	AdminPassword string `mapstructure:"admin_password"`
}

func New() *Config {
	config := &Config{}
	config.getConfig()
	return config
}

func InitConfig(rootCmd *cobra.Command) {
	viper.SetEnvPrefix("shrls")
	bindConfig(rootCmd)

	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.ErrorLevel)
	if viper.GetBool("trace") {
		log.SetLevel(log.TraceLevel)
	} else if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	cfgFile := viper.GetString("config")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(viper.GetString("config"))
		viper.AddConfigPath("/etc/shrls/")
		viper.AddConfigPath("$HOME/.config/shrls")
		viper.AddConfigPath(".")
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Debug("Error reading config, Using defaults.", err)
	}
}

func bindConfig(rootCmd *cobra.Command) {
	viper.BindEnv("base_url")

	viper.BindEnv("port")
	viper.SetDefault("port", 3000)

	viper.BindEnv("default_redirect")

	viper.BindEnv("upload_directory")
	viper.SetDefault("upload_directory", "./uploads")

	viper.BindEnv("state_backend")
	viper.SetDefault("state_backend", "mongo")

	viper.BindEnv("mongo_uri")
	viper.SetDefault("mongo_uri", "mongodb://mongo:password@localhost:27017")

	viper.BindEnv("bolt_path")
	viper.SetDefault("bolt_path", "shrls.db")

	viper.BindEnv("admin_username")
	viper.SetDefault("admin_username", "")

	viper.BindEnv("admin_password")
	viper.SetDefault("admin_password", "")

	viper.SetDefault("terminal_redirect", false)

	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("trace", rootCmd.PersistentFlags().Lookup("trace"))
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

func (c *Config) getConfig() {
	viper.Unmarshal(c)
}
