package pikaq

import log "github.com/sirupsen/logrus"

var LogLevel = log.InfoLevel

func init() {
	log.SetLevel(LogLevel)
}
