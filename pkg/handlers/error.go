package handlers

import (
	log "github.com/sirupsen/logrus"
)

func CheckError(description string, err error) bool {
	b := (err != nil)
	if b {
		log.WithFields(log.Fields{
			"error": err,
		}).Error(description)
	}
	return b
}
