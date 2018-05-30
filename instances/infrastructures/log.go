package infrastructures

import (
	"os"

	"bitbucket.org/burhanmubarok/microservice/structures/infrastructures"
	log "github.com/sirupsen/logrus"
)

// Logger doc
type Logger struct {
	log *log.Logger
}

// Debug doc
func (l *Logger) Debug(data structures.Log) {
	l.write("debug", data)
}

// Info doc
func (l *Logger) Info(data structures.Log) {
	l.write("info", data)
}

// Warning doc
func (l *Logger) Warning(data structures.Log) {
	l.write("warning", data)
}

// Error doc
func (l *Logger) Error(data structures.Log) {
	l.write("error", data)
}

func (l *Logger) write(level string, data structures.Log) {
	l.init()

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		l.log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	switch level {
	case "info", "warning":
		l.log.WithFields(log.Fields{
			"data": data,
		}).Info(data.Message)
	case "error":
		l.log.WithFields(log.Fields{
			"stacks": data,
		}).Error(data.Message)
	default:
		l.log.WithFields(log.Fields{
			"data": data,
		}).Debug(data.Message)
	}
}

func (l *Logger) init() {
	l.log = log.New()
	l.log.SetLevel(log.DebugLevel)
	l.log.Formatter = &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: log.FieldMap{
			"time": "timestamp",
			"msg":  "message",
		},
	}
}