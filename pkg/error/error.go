package error

import (
	log "github.com/sirupsen/logrus"
)

func checkError(description string, err error) {
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error(description)
	}
}
