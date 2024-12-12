package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type DirectoryUploadBackendOptions struct {
	Path string `yaml:"path"`
}

type MongoStateBackendOptions struct {
	ConnectionString string `mapstructure:"connection_string" yaml:"connection_string"`
}

type BoltStateBackendOptions struct {
	Path string `mapstructure:"path" yaml:"path"`
}

type StateBackend struct {
	Mongo *MongoStateBackendOptions `mapstructure:"mongodb" yaml:"mongodb,omitempty"`
	Bolt  *BoltStateBackendOptions  `mapstructure:"bolt" yaml:"bolt,omitempty"`
}

type StateBackendOptions interface {
	MongoStateBackendOptions | BoltStateBackendOptions
}

type UploadBackend struct {
	Directory *DirectoryUploadBackendOptions `mapstructure:"directory" yaml:"directory,omitempty"`
}

type BasicAuthBackendOptions struct {
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
}

type AuthBackend struct {
	Basic *BasicAuthBackendOptions `yaml:"basic,omitempty"`
}

type Config struct {
	Host               string `mapstructure:"host" yaml:"host,omitempty"`
	Port               int    `yaml:"port"`
	DefaultRedirect    string `mapstructure:"default_redirect" yaml:"default_redirect,omitempty"`
	DefaultRedirectSsl bool   `mapstructure:"default_redirect_ssl" yaml:"default_redirect_ssl,omitempty"`
	TerminalRedirect   bool   `mapstructure:"terminal_redirect" yaml:"-,omitempty"`

	// TODO: Initialize backends outside of global config object
	// State
	StateBackend *StateBackend `mapstructure:"state" yaml:"state"`

	// Uploads
	UploadBackend *UploadBackend `mapstructure:"uploads" yaml:"uploads"`

	// Auth
	AuthBackend *AuthBackend `mapstructure:"auth" yaml:"auth,omitempty"`
}

func New() *Config {
	InitConfig()

	config := Config{}
	config.getConfig()

	if config.UploadBackend == nil {
		config.UploadBackend = &UploadBackend{
			Directory: &DirectoryUploadBackendOptions{
				Path: "uploads",
			},
		}
	}

	if config.StateBackend == nil {
		config.StateBackend = &StateBackend{
			Bolt: &BoltStateBackendOptions{
				Path: "shrls.db",
			},
		}
	}

	return &config
}

func InitConfig() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("shrls")
	bindConfig()

	if viper.GetBool("log_json") {
		log.SetFormatter(&log.JSONFormatter{})
	}

	if viper.GetBool("log_trace") {
		log.SetLevel(log.TraceLevel)
	} else if viper.GetBool("log_debug") {
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

func bindConfig() {
	viper.BindEnv("host")
	viper.SetDefault("host", "localhost")

	viper.BindEnv("port")
	viper.SetDefault("port", 3000)

	viper.BindEnv("default_redirect")
	viper.SetDefault("default_redirect", "/admin")

	viper.BindEnv("default_redirect_ssl")
	viper.SetDefault("default_redirect_ssl", false)

	viper.BindEnv("uploads.directory.path", "SHRLS_UPLOAD_DIRECTORY")

	viper.BindEnv("state.bolt.path")

	viper.BindEnv("state.mongodb.connection_string", "SHRLS_MONGO_CONNECTION_STRING")
	//viper.SetDefault("state.mongodb.connection_string", "mongodb://mongo:password@localhost:27017")

	viper.BindEnv("auth.basic.username", "SHRLS_USERNAME")
	viper.BindEnv("auth.basic.password", "SHRLS_PASSWORD")

	viper.SetDefault("terminal_redirect", false)
}

func (c *Config) getConfig() {
	viper.Unmarshal(c)
}
