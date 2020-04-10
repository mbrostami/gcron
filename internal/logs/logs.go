package logs

import (
	"io"
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/mbrostami/gcron/internal/config"
	log "github.com/sirupsen/logrus"
)

// Initialize logs
func Initialize(cfg config.GeneralConfig) *os.File {

	log.SetLevel(cfg.GetLogLevel())
	log.SetFormatter(&nested.Formatter{
		NoColors: false,
	})

	log.SetOutput(os.Stdout)
	var f *os.File
	var err error
	if cfg.GetKey("log.enable").(bool) {
		logPath := cfg.GetKey("log.path").(string)
		f, err = os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Panicf("error opening file: %v", err)
		}
		writers := io.MultiWriter(
			os.Stdout,
			f,
		)
		log.SetOutput(writers)
	}
	// FIXME handle file close internaly
	return f
}
