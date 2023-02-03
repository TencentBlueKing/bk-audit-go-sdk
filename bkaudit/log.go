package bkaudit

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func initLog() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat:   time.RFC3339,
		DisableTimestamp:  false,
		DisableHTMLEscape: true,
		PrettyPrint:       false,
		FieldMap: log.FieldMap{
			log.FieldKeyMsg: "message",
		},
	})
}
