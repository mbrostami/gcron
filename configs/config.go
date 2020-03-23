package configs

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config map configs
type Config struct {
	Log struct {
		Enable bool
		Level  string
		Path   string
	}
	Out struct {
		Clean bool
		Tags  bool
		Hide  struct {
			SysTime  bool
			UserTime bool
			Duration bool
			UID      bool
		}
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
func (cfg *Config) GetLogLevel() log.Level {
	if cfg.Log.Level == "trace" {
		return log.TraceLevel
	} else if cfg.Log.Level == "debug" {
		return log.DebugLevel
	} else if cfg.Log.Level == "info" {
		return log.InfoLevel
	} else if cfg.Log.Level == "warning" {
		return log.WarnLevel
	} else if cfg.Log.Level == "error" {
		return log.ErrorLevel
	} else if cfg.Log.Level == "fatal" {
		return log.FatalLevel
	} else {
		return log.PanicLevel
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
