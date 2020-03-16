package main

import (
	"fmt"
	"log/syslog"
	"os"

	"github.com/spf13/viper"
)

// Config keep all config file values
type Config struct {
	Log struct {
		Level string
		Path  string
	}
	Server struct {
		TCP struct {
			Host string
			Port string
		}
		UDP struct {
			Host string
			Port string
		}
		Unix struct {
			Path string
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
func GetConfig(cfgPath string) Config {
	var cfg Config
	lviper := readFile(cfgPath)
	lviper.Unmarshal(&cfg)
	// fmt.Printf("%+v", cfg)
	return cfg
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfgPath string) *viper.Viper {
	// f, err := os.Open(cfgPath)
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/gcron/")
	viper.AddConfigPath(cfgPath)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		processError(err)
	}
	return viper.GetViper()
}
