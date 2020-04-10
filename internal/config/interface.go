package config

import (
	log "github.com/sirupsen/logrus"
)

// GeneralConfig interface
type GeneralConfig interface {
	GetKey(key string) interface{}
	GetLogLevel() log.Level
}
