package main

import (
	"fmt"
	"log/syslog"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Log struct {
		Level string `yaml:"level"`
		Path  string `yaml:"path"`
	} `yaml:"log"`
	Server struct {
		Port string `yaml:"port", envconfig:"GCRON_SERVER_PORT"`
		Host string `yaml:"host", envconfig:"GCRON_SERVER_HOST"`
	} `yaml:"server"`
}

func (cfg *Config) GetLogLevel() syslog.Priority {
	if cfg.Log.Level == "debug" {
		return syslog.LOG_DEBUG
	} else if cfg.Log.Level == "info" {
		return syslog.LOG_INFO
	} else if cfg.Log.Level == "warning" {
		return syslog.LOG_WARNING
	} else {
		return syslog.LOG_ERR
	}
}

func getConfig() Config {
	var cfg Config
	readFile(&cfg)
	readEnv(&cfg)
	//fmt.Printf("%+v", cfg)
	return cfg
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
