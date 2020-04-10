package config

import (
	"flag"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config map configs
type Config struct {
	lviper *viper.Viper
}

// GetConfig returns the configuration map
func GetConfig(flagset *flag.FlagSet) Config {
	var cfg Config
	lviper := cfg.readFile(flagset)
	cfg.lviper = lviper
	return cfg
}

// GetKey returns value by key
func (cfg Config) GetKey(key string) interface{} {
	return cfg.lviper.Get(key)
}

// GetLogLevel finds the integer map to the log level string in config
func (cfg Config) GetLogLevel() log.Level {
	if cfg.lviper.GetString("log.level") == "trace" {
		return log.TraceLevel
	} else if cfg.lviper.GetString("log.level") == "debug" {
		return log.DebugLevel
	} else if cfg.lviper.GetString("log.level") == "info" {
		return log.InfoLevel
	} else if cfg.lviper.GetString("log.level") == "warning" {
		return log.WarnLevel
	} else if cfg.lviper.GetString("log.level") == "error" {
		return log.ErrorLevel
	} else if cfg.lviper.GetString("log.level") == "fatal" {
		return log.FatalLevel
	} else {
		return log.PanicLevel
	}
}

func (cfg Config) readFile(flagset *flag.FlagSet) *viper.Viper {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/gcron/")
	viper.AddConfigPath("./configs/")
	pflag.CommandLine.AddGoFlagSet(flagset)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Can not read the configs: %v", err)
	}
	return viper.GetViper()
}
