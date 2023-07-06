package config

import (
	log "github.com/sirupsen/logrus"
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
	BaseURL          string
	Port             int
	GrpcPort         int `mapstructure:"grpc_port"`
	DefaultRedirect  string
	TerminalRedirect bool

	// TODO: Initialize backends outside of global config object
	// State
	StateBackend          StateBackend
	MongoConnectionString string `mapstructure:"mongo_uri"`

	// Uploads
	UploadBackend   UploadBackend
	UploadDirectory string

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

func InitConfig() {
	viper.SetEnvPrefix("shrls")
	bindConfig()

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

	bindConfig()
}

func bindConfig() {
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

func (c *Config) getConfig() {
	viper.Unmarshal(c)
}
