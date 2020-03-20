package configs

import (
	"flag"
	"fmt"
	"log/syslog"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config keep all config file values
type Config struct {
	Log struct {
		Level string
		Path  string
	}
	Out struct {
		Clean  bool
		Notime bool
	}
	Server struct {
		TCP struct {
			Enabled bool
			Host    string
			Port    string
		}
		UDP struct {
			Enabled bool
			Host    string
			Port    string
		}
		Unix struct {
			Enabled bool
			Path    string
		}
	}
}

// GetLogLevel finds the integer map to the log level string in config
func (cfg *Config) GetLogLevel() syslog.Priority {
	if cfg.Log.Level == "debug" {
		return syslog.LOG_DEBUG
	} else if cfg.Log.Level == "info" {
		return syslog.LOG_INFO
	} else if cfg.Log.Level == "notice" {
		return syslog.LOG_NOTICE
	} else if cfg.Log.Level == "warning" {
		return syslog.LOG_WARNING
	} else if cfg.Log.Level == "error" {
		return syslog.LOG_ERR
	} else if cfg.Log.Level == "critical" {
		return syslog.LOG_CRIT
	} else if cfg.Log.Level == "alert" {
		return syslog.LOG_ALERT
	} else {
		return syslog.LOG_EMERG
	}
}

// GetConfig returns the configuration map
func GetConfig(flagset *flag.FlagSet) Config {
	var cfg Config
	lviper := readFile(flagset)
	lviper.Unmarshal(&cfg)
	// fmt.Printf("%+v", cfg)
	return cfg
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(flagset *flag.FlagSet) *viper.Viper {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/gcron/")
	viper.AddConfigPath(".")
	pflag.CommandLine.AddGoFlagSet(flagset)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		processError(err)
	}
	return viper.GetViper()
}
